package redistool

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

const (
	redisURLEnvName = "REDIS_URL"
	ttl             = 2 * time.Second
)

var (
	_ ExpiringHashInterface[int, int] = (*ExpiringHash[int, int])(nil)
)

func TestExpiringHash_Set(t *testing.T) {
	client, hash, key, value := setupHash(t)

	require.NoError(t, hash.Set(key, 123, value)(context.Background()))

	equalHash(t, client, key, 123, value)
}

func TestExpiringHash_Unset(t *testing.T) {
	client, hash, key, value := setupHash(t)

	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	require.NoError(t, hash.Unset(key, 123)(context.Background()))

	require.Empty(t, getHash(t, client, key))
}

func TestExpiringHash_Forget(t *testing.T) {
	client, hash, key, value := setupHash(t)

	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	hash.Forget(key, 123)

	equalHash(t, client, key, 123, value)
	require.Empty(t, hash.data)
}

func TestExpiringHash_Expires(t *testing.T) {
	client, hash, key, value := setupHash(t)

	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	time.Sleep(ttl + 100*time.Millisecond)

	require.Empty(t, getHash(t, client, key))
}

func TestExpiringHash_GC(t *testing.T) {
	client, hash, key, value := setupHash(t)

	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	_, err := client.Pipelined(context.Background(), func(p redis.Pipeliner) error {
		newExpireIn := 3 * ttl
		p.PExpire(context.Background(), key, newExpireIn)
		return nil
	})
	require.NoError(t, err)
	time.Sleep(ttl + time.Second)
	require.NoError(t, hash.Set(key, 321, value)(context.Background()))

	keysDeleted, err := hash.GC()(context.Background())
	require.NoError(t, err)
	assert.EqualValues(t, 1, keysDeleted)

	equalHash(t, client, key, 321, value)
}

func TestExpiringHash_Refresh_ToExpireSoonerThanNextRefresh(t *testing.T) {
	client, hash, key, value := setupHash(t)

	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	registrationTime := time.Now()
	time.Sleep(ttl / 2)
	require.NoError(t, hash.Refresh(registrationTime.Add(ttl*2))(context.Background()))

	expireAfter := registrationTime.Add(ttl)
	valuesExpireAfter(t, client, key, expireAfter)
}

func TestExpiringHash_Refresh_ToExpireAfterNextRefresh(t *testing.T) {
	client, hash, key, value := setupHash(t)

	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	h1 := getHash(t, client, key)
	require.NoError(t, hash.Refresh(time.Now().Add(ttl/10))(context.Background()))
	h2 := getHash(t, client, key)
	assert.Equal(t, h1, h2)
}

func TestExpiringHash_ScanEmpty(t *testing.T) {
	_, hash, key, _ := setupHash(t)

	keysDeleted, err := hash.Scan(context.Background(), key, func(rawHashKey string, value []byte, err error) (bool, error) {
		require.NoError(t, err)
		assert.FailNow(t, "unexpected callback invocation")
		return false, nil
	})
	require.NoError(t, err)
	assert.Zero(t, keysDeleted)
}

func TestExpiringHash_Scan(t *testing.T) {
	_, hash, key, value := setupHash(t)
	cbCalled := false

	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	keysDeleted, err := hash.Scan(context.Background(), key, func(rawHashKey string, v []byte, err error) (bool, error) {
		cbCalled = true
		require.NoError(t, err)
		assert.Equal(t, value, v)
		assert.Equal(t, "123", rawHashKey)
		return false, nil
	})
	require.NoError(t, err)
	assert.Zero(t, keysDeleted)
	assert.True(t, cbCalled)
}

func TestExpiringHash_Len(t *testing.T) {
	_, hash, key, value := setupHash(t)
	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	size, err := hash.Len(context.Background(), key)
	require.NoError(t, err)
	assert.EqualValues(t, 1, size)
}

func TestExpiringHash_ScanGC(t *testing.T) {
	client, hash, key, value := setupHash(t)
	cbCalled := false

	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	_, err := client.Pipelined(context.Background(), func(p redis.Pipeliner) error {
		cbCalled = true
		newExpireIn := 3 * ttl
		p.PExpire(context.Background(), key, newExpireIn)
		return nil
	})
	require.NoError(t, err)
	time.Sleep(ttl + time.Second)
	require.NoError(t, hash.Set(key, 321, value)(context.Background()))
	assert.True(t, cbCalled)

	cbCalled = false
	keysDeleted, err := hash.Scan(context.Background(), key, func(rawHashKey string, v []byte, err error) (bool, error) {
		cbCalled = true
		require.NoError(t, err)
		assert.Equal(t, "321", rawHashKey)
		assert.Equal(t, value, v)
		return false, nil
	})
	require.NoError(t, err)
	assert.EqualValues(t, 1, keysDeleted)
	assert.True(t, cbCalled)
}

func TestExpiringHash_Clear(t *testing.T) {
	client, hash, key, value := setupHash(t)
	require.NoError(t, hash.Set(key, 123, value)(context.Background()))
	require.NoError(t, hash.Set(key+"123", 321, value)(context.Background()))
	size, err := hash.Clear(context.Background())
	require.NoError(t, err)
	assert.EqualValues(t, 2, size)
	assert.Empty(t, hash.data)
	h := getHash(t, client, key)
	assert.Empty(t, h)
	size, err = hash.Clear(context.Background())
	require.NoError(t, err)
	assert.Zero(t, size)
}

func BenchmarkExpiringValue_Unmarshal(b *testing.B) {
	d, err := proto.Marshal(&ExpiringValue{
		ExpiresAt: 123123123,
		Value:     []byte("1231231231232313"),
	})
	require.NoError(b, err)
	b.Run("ExpiringValue", func(b *testing.B) {
		b.ReportAllocs()
		var val ExpiringValue
		for i := 0; i < b.N; i++ {
			err = proto.Unmarshal(d, &val)
		}
	})
	b.Run("ExpiringValueTimestamp", func(b *testing.B) {
		b.ReportAllocs()
		var val ExpiringValueTimestamp
		for i := 0; i < b.N; i++ {
			err = proto.Unmarshal(d, &val)
		}
	})
	b.Run("ExpiringValueTimestamp DiscardUnknown", func(b *testing.B) {
		b.ReportAllocs()
		var val ExpiringValueTimestamp
		for i := 0; i < b.N; i++ {
			err = proto.UnmarshalOptions{
				DiscardUnknown: true,
			}.Unmarshal(d, &val)
		}
	})
}

func setupHash(t *testing.T) (redis.UniversalClient, *ExpiringHash[string, int64], string, []byte) {
	t.Parallel()
	client := redisClient(t)
	t.Cleanup(func() {
		assert.NoError(t, client.Close())
	})
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prefix := make([]byte, 32)
	_, err := r.Read(prefix)
	require.NoError(t, err)
	key := string(prefix)
	hash := NewExpiringHash[string, int64](client, func(key string) string {
		return key
	}, int64ToStr, ttl)
	return client, hash, key, []byte{1, 2, 3}
}

func redisClient(t *testing.T) redis.UniversalClient {
	redisURL := os.Getenv(redisURLEnvName)
	if redisURL == "" {
		t.Skipf("%s environment variable not set, skipping test", redisURLEnvName)
	}

	opts, err := redis.ParseURL(redisURL)
	require.NoError(t, err)
	return redis.NewClient(opts)
}

func getHash(t *testing.T, client redis.UniversalClient, key string) map[string]string {
	reply, err := client.HGetAll(context.Background(), key).Result()
	require.NoError(t, err)
	return reply
}

func equalHash(t *testing.T, client redis.UniversalClient, key string, hashKey int64, value []byte) {
	hash := getHash(t, client, key)
	require.Len(t, hash, 1)
	connectionIdStr := strconv.FormatInt(hashKey, 10)
	require.Contains(t, hash, connectionIdStr)
	val := hash[connectionIdStr]
	var msg ExpiringValue
	err := proto.Unmarshal([]byte(val), &msg)
	require.NoError(t, err)
	assert.Equal(t, value, msg.Value)
}

func valuesExpireAfter(t *testing.T, client redis.UniversalClient, key string, expireAfter time.Time) {
	hash := getHash(t, client, key)
	require.NotEmpty(t, hash)
	for _, val := range hash {
		var msg ExpiringValue
		err := proto.Unmarshal([]byte(val), &msg)
		require.NoError(t, err)
		assert.Greater(t, msg.ExpiresAt, expireAfter.Unix())
	}
}

func int64ToStr(key int64) string {
	return strconv.FormatInt(key, 10)
}
