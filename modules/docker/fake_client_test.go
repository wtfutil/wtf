package docker

import (
	"context"
	"errors"

	"github.com/moby/moby/client"
)

// fakeDockerClient is a test double for dockerAPIClient. It lets tests drive
// getSystemInfo/getContainerStates/refreshDisplayBuffer without a running
// docker daemon by returning configurable values or errors.
type fakeDockerClient struct {
	info      client.SystemInfoResult
	infoErr   error
	diskUsage client.DiskUsageResult
	diskErr   error
	cntrs     client.ContainerListResult
	listErr   error
}

func (f *fakeDockerClient) Info(_ context.Context, _ client.InfoOptions) (client.SystemInfoResult, error) {
	return f.info, f.infoErr
}

func (f *fakeDockerClient) DiskUsage(_ context.Context, _ client.DiskUsageOptions) (client.DiskUsageResult, error) {
	return f.diskUsage, f.diskErr
}

func (f *fakeDockerClient) ContainerList(_ context.Context, _ client.ContainerListOptions) (client.ContainerListResult, error) {
	return f.cntrs, f.listErr
}

var errBoom = errors.New("boom")
