package pktline

// Utility functions for working with the Git pkt-line format. See
// https://github.com/git/git/blob/master/Documentation/technical/protocol-common.txt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"unsafe"
)

const (
	// MaxPktSize is the maximum size of content of a Git pktline side-band-64k
	// packet, excluding size of length and band number
	// https://gitlab.com/gitlab-org/git/-/blob/v2.30.0/pkt-line.h#L216
	MaxPktSize = 65520
	pktDelim   = "0001"
)

// NewScanner returns a bufio.Scanner that splits on Git pktline boundaries
// Buf must be at least MaxPktSize large.
func NewScanner(r io.Reader, buf []byte) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(buf[0:MaxPktSize:MaxPktSize], MaxPktSize)
	scanner.Split(pktLineSplitter)
	return scanner
}

// Data returns the packet pkt without its length header. The length
// header is not validated. Returns an empty slice when pkt is a magic packet such
// as '0000'.
func Data(pkt []byte) []byte {
	return pkt[4:]
}

// IsFlush detects the special flush packet '0000'
func IsFlush(pkt []byte) bool {
	return bytes.Equal(pkt, PktFlush())
}

// WriteString writes a string with pkt-line framing
func WriteString(w io.Writer, str string) (int, error) {
	pktLen := len(str) + 4
	if pktLen > MaxPktSize {
		return 0, fmt.Errorf("string too large: %d bytes", len(str))
	}

	_, err := fmt.Fprintf(w, "%04x%s", pktLen, str)
	return len(str), err
}

// WriteFlush writes a pkt flush packet.
func WriteFlush(w io.Writer) error {
	_, err := w.Write(PktFlush())
	return err
}

// WriteDelim writes a pkt delim packet.
func WriteDelim(w io.Writer) error {
	_, err := fmt.Fprint(w, pktDelim)
	return err
}

// PktDone returns the bytes for a "done" packet.
func PktDone() []byte {
	return []byte("0009done\n")
}

// PktFlush returns the bytes for a "flush" packet.
func PktFlush() []byte {
	return []byte("0000")
}

func pktLineSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) < 4 {
		if atEOF && len(data) > 0 {
			return 0, nil, fmt.Errorf("pktLineSplitter: incomplete length prefix on %q", data)
		}
		return 0, nil, nil // want more data
	}

	// We have at least 4 bytes available so we can decode the 4-hex digit
	// length prefix of the packet line. Avoid allocating memory for parsing.
	lenSliceStr := unsafe.String(unsafe.SliceData(data), 4)

	pktLength64, err := strconv.ParseUint(lenSliceStr, 16, 0)
	if err != nil {
		return 0, nil, fmt.Errorf("pktLineSplitter: decode length: %w", err)
	}

	// Cast is safe because we requested an int-size number from strconv.ParseInt
	pktLength := int(pktLength64)

	if pktLength < 0 {
		return 0, nil, fmt.Errorf("pktLineSplitter: invalid length: %d", pktLength)
	}

	if pktLength < 4 {
		// Special case: magic empty packet 0000, 0001, 0002 or 0003.
		return 4, data[:4], nil
	}

	if len(data) < pktLength {
		// data contains incomplete packet

		if atEOF {
			return 0, nil, io.ErrUnexpectedEOF
		}

		return 0, nil, nil // want more data
	}

	return pktLength, data[:pktLength], nil
}
