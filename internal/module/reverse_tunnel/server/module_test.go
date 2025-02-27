package server

import (
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/module/modserver"
)

var (
	_ modserver.Module  = &module{}
	_ modserver.Factory = &Factory{}
)
