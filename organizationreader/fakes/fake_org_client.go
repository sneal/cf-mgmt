// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/cloudfoundry-community/go-cfclient/v3/client"
	"github.com/cloudfoundry-community/go-cfclient/v3/resource"
	"github.com/vmwarepivotallabs/cf-mgmt/organizationreader"
)

type FakeCFOrganizationClient struct {
	DeleteStub        func(context.Context, string) (string, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	deleteReturns struct {
		result1 string
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetStub        func(context.Context, string) (*resource.Organization, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getReturns struct {
		result1 *resource.Organization
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *resource.Organization
		result2 error
	}
	ListAllStub        func(context.Context, *client.OrganizationListOptions) ([]*resource.Organization, error)
	listAllMutex       sync.RWMutex
	listAllArgsForCall []struct {
		arg1 context.Context
		arg2 *client.OrganizationListOptions
	}
	listAllReturns struct {
		result1 []*resource.Organization
		result2 error
	}
	listAllReturnsOnCall map[int]struct {
		result1 []*resource.Organization
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFOrganizationClient) Delete(arg1 context.Context, arg2 string) (string, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFOrganizationClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCFOrganizationClient) DeleteCalls(stub func(context.Context, string) (string, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeCFOrganizationClient) DeleteArgsForCall(i int) (context.Context, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFOrganizationClient) DeleteReturns(result1 string, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFOrganizationClient) DeleteReturnsOnCall(i int, result1 string, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFOrganizationClient) Get(arg1 context.Context, arg2 string) (*resource.Organization, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFOrganizationClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeCFOrganizationClient) GetCalls(stub func(context.Context, string) (*resource.Organization, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeCFOrganizationClient) GetArgsForCall(i int) (context.Context, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFOrganizationClient) GetReturns(result1 *resource.Organization, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeCFOrganizationClient) GetReturnsOnCall(i int, result1 *resource.Organization, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *resource.Organization
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeCFOrganizationClient) ListAll(arg1 context.Context, arg2 *client.OrganizationListOptions) ([]*resource.Organization, error) {
	fake.listAllMutex.Lock()
	ret, specificReturn := fake.listAllReturnsOnCall[len(fake.listAllArgsForCall)]
	fake.listAllArgsForCall = append(fake.listAllArgsForCall, struct {
		arg1 context.Context
		arg2 *client.OrganizationListOptions
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

func (fake *FakeCFOrganizationClient) ListAllCallCount() int {
	fake.listAllMutex.RLock()
	defer fake.listAllMutex.RUnlock()
	return len(fake.listAllArgsForCall)
}

func (fake *FakeCFOrganizationClient) ListAllCalls(stub func(context.Context, *client.OrganizationListOptions) ([]*resource.Organization, error)) {
	fake.listAllMutex.Lock()
	defer fake.listAllMutex.Unlock()
	fake.ListAllStub = stub
}

func (fake *FakeCFOrganizationClient) ListAllArgsForCall(i int) (context.Context, *client.OrganizationListOptions) {
	fake.listAllMutex.RLock()
	defer fake.listAllMutex.RUnlock()
	argsForCall := fake.listAllArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFOrganizationClient) ListAllReturns(result1 []*resource.Organization, result2 error) {
	fake.listAllMutex.Lock()
	defer fake.listAllMutex.Unlock()
	fake.ListAllStub = nil
	fake.listAllReturns = struct {
		result1 []*resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeCFOrganizationClient) ListAllReturnsOnCall(i int, result1 []*resource.Organization, result2 error) {
	fake.listAllMutex.Lock()
	defer fake.listAllMutex.Unlock()
	fake.ListAllStub = nil
	if fake.listAllReturnsOnCall == nil {
		fake.listAllReturnsOnCall = make(map[int]struct {
			result1 []*resource.Organization
			result2 error
		})
	}
	fake.listAllReturnsOnCall[i] = struct {
		result1 []*resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeCFOrganizationClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listAllMutex.RLock()
	defer fake.listAllMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFOrganizationClient) recordInvocation(key string, args []interface{}) {
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

var _ organizationreader.CFOrganizationClient = new(FakeCFOrganizationClient)
