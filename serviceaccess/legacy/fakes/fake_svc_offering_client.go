// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/cloudfoundry-community/go-cfclient/v3/client"
	"github.com/cloudfoundry-community/go-cfclient/v3/resource"
	"github.com/vmwarepivotallabs/cf-mgmt/serviceaccess/legacy"
)

type FakeCFServiceOfferingClient struct {
	ListAllStub        func(context.Context, *client.ServiceOfferingListOptions) ([]*resource.ServiceOffering, error)
	listAllMutex       sync.RWMutex
	listAllArgsForCall []struct {
		arg1 context.Context
		arg2 *client.ServiceOfferingListOptions
	}
	listAllReturns struct {
		result1 []*resource.ServiceOffering
		result2 error
	}
	listAllReturnsOnCall map[int]struct {
		result1 []*resource.ServiceOffering
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFServiceOfferingClient) ListAll(arg1 context.Context, arg2 *client.ServiceOfferingListOptions) ([]*resource.ServiceOffering, error) {
	fake.listAllMutex.Lock()
	ret, specificReturn := fake.listAllReturnsOnCall[len(fake.listAllArgsForCall)]
	fake.listAllArgsForCall = append(fake.listAllArgsForCall, struct {
		arg1 context.Context
		arg2 *client.ServiceOfferingListOptions
	}{arg1, arg2})
	stub := fake.ListAllStub
	fakeReturns := fake.listAllReturns
	fake.recordInvocation("ListAll", []interface{}{arg1, arg2})
	fake.listAllMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFServiceOfferingClient) ListAllCallCount() int {
	fake.listAllMutex.RLock()
	defer fake.listAllMutex.RUnlock()
	return len(fake.listAllArgsForCall)
}

func (fake *FakeCFServiceOfferingClient) ListAllCalls(stub func(context.Context, *client.ServiceOfferingListOptions) ([]*resource.ServiceOffering, error)) {
	fake.listAllMutex.Lock()
	defer fake.listAllMutex.Unlock()
	fake.ListAllStub = stub
}

func (fake *FakeCFServiceOfferingClient) ListAllArgsForCall(i int) (context.Context, *client.ServiceOfferingListOptions) {
	fake.listAllMutex.RLock()
	defer fake.listAllMutex.RUnlock()
	argsForCall := fake.listAllArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFServiceOfferingClient) ListAllReturns(result1 []*resource.ServiceOffering, result2 error) {
	fake.listAllMutex.Lock()
	defer fake.listAllMutex.Unlock()
	fake.ListAllStub = nil
	fake.listAllReturns = struct {
		result1 []*resource.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeCFServiceOfferingClient) ListAllReturnsOnCall(i int, result1 []*resource.ServiceOffering, result2 error) {
	fake.listAllMutex.Lock()
	defer fake.listAllMutex.Unlock()
	fake.ListAllStub = nil
	if fake.listAllReturnsOnCall == nil {
		fake.listAllReturnsOnCall = make(map[int]struct {
			result1 []*resource.ServiceOffering
			result2 error
		})
	}
	fake.listAllReturnsOnCall[i] = struct {
		result1 []*resource.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeCFServiceOfferingClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listAllMutex.RLock()
	defer fake.listAllMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFServiceOfferingClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ legacy.CFServiceOfferingClient = new(FakeCFServiceOfferingClient)
