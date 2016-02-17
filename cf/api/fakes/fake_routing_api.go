// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/fujitsu-cf/cli/cf/api"
	"github.com/fujitsu-cf/cli/cf/models"
)

type FakeRoutingApiRepository struct {
	ListRouterGroupsStub        func(cb func(models.RouterGroup) bool) (apiErr error)
	listRouterGroupsMutex       sync.RWMutex
	listRouterGroupsArgsForCall []struct {
		cb func(models.RouterGroup) bool
	}
	listRouterGroupsReturns struct {
		result1 error
	}
}

func (fake *FakeRoutingApiRepository) ListRouterGroups(cb func(models.RouterGroup) bool) (apiErr error) {
	fake.listRouterGroupsMutex.Lock()
	fake.listRouterGroupsArgsForCall = append(fake.listRouterGroupsArgsForCall, struct {
		cb func(models.RouterGroup) bool
	}{cb})
	fake.listRouterGroupsMutex.Unlock()
	if fake.ListRouterGroupsStub != nil {
		return fake.ListRouterGroupsStub(cb)
	} else {
		return fake.listRouterGroupsReturns.result1
	}
}

func (fake *FakeRoutingApiRepository) ListRouterGroupsCallCount() int {
	fake.listRouterGroupsMutex.RLock()
	defer fake.listRouterGroupsMutex.RUnlock()
	return len(fake.listRouterGroupsArgsForCall)
}

func (fake *FakeRoutingApiRepository) ListRouterGroupsArgsForCall(i int) func(models.RouterGroup) bool {
	fake.listRouterGroupsMutex.RLock()
	defer fake.listRouterGroupsMutex.RUnlock()
	return fake.listRouterGroupsArgsForCall[i].cb
}

func (fake *FakeRoutingApiRepository) ListRouterGroupsReturns(result1 error) {
	fake.ListRouterGroupsStub = nil
	fake.listRouterGroupsReturns = struct {
		result1 error
	}{result1}
}

var _ api.RoutingApiRepository = new(FakeRoutingApiRepository)
