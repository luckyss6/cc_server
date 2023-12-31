package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerService struct {
}

func NewDockerService() *DockerService {
	return &DockerService{}
}

func (d *DockerService) GetImageList(ctx context.Context) ([]types.ImageSummary, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		slog.Error(fmt.Sprintf("create docker client error: %s", err.Error()))
		return nil, err
	}
	defer cli.Close()

	imageInfo, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		slog.Error(fmt.Sprintf("get image list error: %s", err.Error()))
		return nil, err
	}
	return imageInfo, nil
}

func (d *DockerService) GetContainerList(ctx context.Context) ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		slog.Error(fmt.Sprintf("create docker client error: %s", err.Error()))
		return nil, err
	}
	defer cli.Close()

	containerInfo, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		slog.Error(fmt.Sprintf("get container list error: %s", err.Error()))
		return nil, err
	}
	return containerInfo, nil
}

func (d *DockerService) GetContainerInfo(ctx context.Context, id string) (*types.ContainerJSON, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		slog.Error(fmt.Sprintf("create docker client error: %s", err.Error()))
		return nil, err
	}
	defer cli.Close()

	containerInfo, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		slog.Error(fmt.Sprintf("get container info error: %s", err.Error()))
		return nil, err
	}
	return &containerInfo, nil
}


