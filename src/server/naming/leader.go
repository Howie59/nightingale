package naming

import (
	"context"
	"sort"

	"github.com/didi/nightingale/v5/src/server/config"
	"github.com/toolkits/pkg/logger"
)

// IamLeader 取第一个作为leader
func IamLeader() (bool, error) {
	servers, err := ActiveServers(context.Background(), config.C.ClusterName)
	if err != nil {
		logger.Errorf("failed to get active servers: %v", err)
		return false, err
	}

	if len(servers) == 0 {
		logger.Errorf("active servers empty")
		return false, err
	}

	sort.Strings(servers)

	return config.C.Heartbeat.Endpoint == servers[0], nil
}
