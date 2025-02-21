package entity

import (
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/gitaly/vendored/gitalypb"
)

func (r *GitalyRepository) ToGitalyRepository() *gitalypb.Repository {
	return &gitalypb.Repository{
		StorageName:                   r.StorageName,
		RelativePath:                  r.RelativePath,
		GitObjectDirectory:            r.GitObjectDirectory,
		GitAlternateObjectDirectories: r.GitAlternateObjectDirectories,
		GlRepository:                  r.GlRepository,
		GlProjectPath:                 r.GlProjectPath,
	}
}
