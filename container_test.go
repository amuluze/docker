// Package docker
// Date: 2024/07/09 14:13:44
// Author: Amu
// Description:
package docker

import (
	"context"
	"testing"
)

func TestListContainer(t *testing.T) {
	manager, _ := NewManager()
	containers, _ := manager.ListContainer(context.Background())
	for _, c := range containers {
		t.Logf("container name: %s, container ports: %#v, container labels: %#v, container network: %#v\n", c.Name, c.Ports, c.Labels, c.Network)
	}
}

func TestContainerCreate(t *testing.T) {
	manager, _ := NewManager()
	cid, err := manager.CreateContainer(
		context.Background(),
		"redis",
		"redis:7.0.5",
		"test",
		[]string{"6379:6379"},
		[]string{"/Users/amu/Desktop/common.scss:/app/common.scss:rw"},
		[]string{},
		[]string{"redis-server", "--requirepass", "coreblox123"},
		map[string]string{CreatedByProbe: "true", ServerTypeLabel: WebServer},
	)
	if err != nil {
		t.Error("create container error: ", err)
	}
	t.Logf("container id: %#v", cid)
}

func TestContainerMem(t *testing.T) {
	manager, _ := NewManager()
	percent, used, limit, err := manager.GetContainerMem(context.Background(), "5c28bf6e16be")
	if err != nil {
		panic(err)
	}
	t.Logf("container mem percent: %v, used: %v, limit: %v \n", percent, used, limit)
}

func TestContainerCPU(t *testing.T) {
	manager, _ := NewManager()
	cpu, err := manager.GetContainerCpu(context.Background(), "5c28bf6e16be")
	if err != nil {
		panic(err)
	}
	t.Logf("cpu percent: %v\n", cpu)
}

func TestRenameContainer(t *testing.T) {
	manager, _ := NewManager()
	err := manager.RenameContainer(context.Background(), "5c28bf6e16be", "tt")
	if err != nil {
		t.Error("rename container error: ", err)
	}
}

func TestContainerStart(t *testing.T) {
	manager, _ := NewManager()
	err := manager.StartContainer(context.Background(), "5c28bf6e16be")
	if err != nil {
		t.Error("start container error: ", err)
	}
}

func TestContainerStop(t *testing.T) {
	manager, _ := NewManager()
	err := manager.StopContainer(context.Background(), "eedaf881e6c8")
	if err != nil {
		t.Error("stop container error: ", err)
	}
}

func TestContainerDelete(t *testing.T) {
	manager, _ := NewManager()
	err := manager.DeleteContainer(context.Background(), "eedaf881e6c8")
	if err != nil {
		t.Error("delete container error: ", err)
	}
}

func TestContainerExists(t *testing.T) {
	manager, _ := NewManager()
	exists, err := manager.ContainerExists(context.Background(), "e79ceb874917")
	if err != nil {
		t.Error("container exists error: ", err)
		return
	}
	t.Log(exists)
}
