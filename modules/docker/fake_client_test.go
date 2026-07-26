package docker

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/system"
)

// fakeDockerClient is a test double for dockerAPIClient. It lets tests drive
// getSystemInfo/getContainerStates/refreshDisplayBuffer without a running
// docker daemon by returning configurable values or errors.
type fakeDockerClient struct {
	info      system.Info
	infoErr   error
	diskUsage types.DiskUsage
	diskErr   error
	cntrs     []container.Summary
	listErr   error
}

func (f *fakeDockerClient) Info(_ context.Context) (system.Info, error) {
	return f.info, f.infoErr
}

func (f *fakeDockerClient) DiskUsage(_ context.Context, _ types.DiskUsageOptions) (types.DiskUsage, error) {
	return f.diskUsage, f.diskErr
}

func (f *fakeDockerClient) ContainerList(_ context.Context, _ container.ListOptions) ([]container.Summary, error) {
	return f.cntrs, f.listErr
}

var errBoom = errors.New("boom")
