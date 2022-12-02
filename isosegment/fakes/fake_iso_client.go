// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/cloudfoundry-community/go-cfclient/v3/client"
	"github.com/cloudfoundry-community/go-cfclient/v3/resource"
	"github.com/vmwarepivotallabs/cf-mgmt/isosegment"
)

type FakeCFIsolationSegmentClient struct {
	CreateStub        func(context.Context, *resource.IsolationSegmentCreate) (*resource.IsolationSegment, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 *resource.IsolationSegmentCreate
	}
	createReturns struct {
		result1 *resource.IsolationSegment
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *resource.IsolationSegment
		result2 error
	}
	DeleteStub        func(context.Context, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	EntitleOrganizationStub        func(context.Context, string, string) (*resource.IsolationSegmentRelationship, error)
	entitleOrganizationMutex       sync.RWMutex
	entitleOrganizationArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	entitleOrganizationReturns struct {
		result1 *resource.IsolationSegmentRelationship
		result2 error
	}
	entitleOrganizationReturnsOnCall map[int]struct {
		result1 *resource.IsolationSegmentRelationship
		result2 error
	}
	GetStub        func(context.Context, string) (*resource.IsolationSegment, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getReturns struct {
		result1 *resource.IsolationSegment
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *resource.IsolationSegment
		result2 error
	}
	ListAllStub        func(context.Context, *client.IsolationSegmentListOptions) ([]*resource.IsolationSegment, error)
	listAllMutex       sync.RWMutex
	listAllArgsForCall []struct {
		arg1 context.Context
		arg2 *client.IsolationSegmentListOptions
	}
	listAllReturns struct {
		result1 []*resource.IsolationSegment
		result2 error
	}
	listAllReturnsOnCall map[int]struct {
		result1 []*resource.IsolationSegment
		result2 error
	}
	RevokeOrganizationStub        func(context.Context, string, string) error
	revokeOrganizationMutex       sync.RWMutex
	revokeOrganizationArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	revokeOrganizationReturns struct {
		result1 error
	}
	revokeOrganizationReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFIsolationSegmentClient) Create(arg1 context.Context, arg2 *resource.IsolationSegmentCreate) (*resource.IsolationSegment, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 *resource.IsolationSegmentCreate
	}{arg1, arg2})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFIsolationSegmentClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeCFIsolationSegmentClient) CreateCalls(stub func(context.Context, *resource.IsolationSegmentCreate) (*resource.IsolationSegment, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeCFIsolationSegmentClient) CreateArgsForCall(i int) (context.Context, *resource.IsolationSegmentCreate) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFIsolationSegmentClient) CreateReturns(result1 *resource.IsolationSegment, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *resource.IsolationSegment
		result2 error
	}{result1, result2}
}

func (fake *FakeCFIsolationSegmentClient) CreateReturnsOnCall(i int, result1 *resource.IsolationSegment, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *resource.IsolationSegment
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *resource.IsolationSegment
		result2 error
	}{result1, result2}
}

func (fake *FakeCFIsolationSegmentClient) Delete(arg1 context.Context, arg2 string) error {
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
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCFIsolationSegmentClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCFIsolationSegmentClient) DeleteCalls(stub func(context.Context, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeCFIsolationSegmentClient) DeleteArgsForCall(i int) (context.Context, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFIsolationSegmentClient) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFIsolationSegmentClient) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFIsolationSegmentClient) EntitleOrganization(arg1 context.Context, arg2 string, arg3 string) (*resource.IsolationSegmentRelationship, error) {
	fake.entitleOrganizationMutex.Lock()
	ret, specificReturn := fake.entitleOrganizationReturnsOnCall[len(fake.entitleOrganizationArgsForCall)]
	fake.entitleOrganizationArgsForCall = append(fake.entitleOrganizationArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.EntitleOrganizationStub
	fakeReturns := fake.entitleOrganizationReturns
	fake.recordInvocation("EntitleOrganization", []interface{}{arg1, arg2, arg3})
	fake.entitleOrganizationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFIsolationSegmentClient) EntitleOrganizationCallCount() int {
	fake.entitleOrganizationMutex.RLock()
	defer fake.entitleOrganizationMutex.RUnlock()
	return len(fake.entitleOrganizationArgsForCall)
}

func (fake *FakeCFIsolationSegmentClient) EntitleOrganizationCalls(stub func(context.Context, string, string) (*resource.IsolationSegmentRelationship, error)) {
	fake.entitleOrganizationMutex.Lock()
	defer fake.entitleOrganizationMutex.Unlock()
	fake.EntitleOrganizationStub = stub
}

func (fake *FakeCFIsolationSegmentClient) EntitleOrganizationArgsForCall(i int) (context.Context, string, string) {
	fake.entitleOrganizationMutex.RLock()
	defer fake.entitleOrganizationMutex.RUnlock()
	argsForCall := fake.entitleOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCFIsolationSegmentClient) EntitleOrganizationReturns(result1 *resource.IsolationSegmentRelationship, result2 error) {
	fake.entitleOrganizationMutex.Lock()
	defer fake.entitleOrganizationMutex.Unlock()
	fake.EntitleOrganizationStub = nil
	fake.entitleOrganizationReturns = struct {
		result1 *resource.IsolationSegmentRelationship
		result2 error
	}{result1, result2}
}

func (fake *FakeCFIsolationSegmentClient) EntitleOrganizationReturnsOnCall(i int, result1 *resource.IsolationSegmentRelationship, result2 error) {
	fake.entitleOrganizationMutex.Lock()
	defer fake.entitleOrganizationMutex.Unlock()
	fake.EntitleOrganizationStub = nil
	if fake.entitleOrganizationReturnsOnCall == nil {
		fake.entitleOrganizationReturnsOnCall = make(map[int]struct {
			result1 *resource.IsolationSegmentRelationship
			result2 error
		})
	}
	fake.entitleOrganizationReturnsOnCall[i] = struct {
		result1 *resource.IsolationSegmentRelationship
		result2 error
	}{result1, result2}
}

func (fake *FakeCFIsolationSegmentClient) Get(arg1 context.Context, arg2 string) (*resource.IsolationSegment, error) {
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

func (fake *FakeCFIsolationSegmentClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeCFIsolationSegmentClient) GetCalls(stub func(context.Context, string) (*resource.IsolationSegment, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeCFIsolationSegmentClient) GetArgsForCall(i int) (context.Context, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFIsolationSegmentClient) GetReturns(result1 *resource.IsolationSegment, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *resource.IsolationSegment
		result2 error
	}{result1, result2}
}

func (fake *FakeCFIsolationSegmentClient) GetReturnsOnCall(i int, result1 *resource.IsolationSegment, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *resource.IsolationSegment
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *resource.IsolationSegment
		result2 error
	}{result1, result2}
}

func (fake *FakeCFIsolationSegmentClient) ListAll(arg1 context.Context, arg2 *client.IsolationSegmentListOptions) ([]*resource.IsolationSegment, error) {
	fake.listAllMutex.Lock()
	ret, specificReturn := fake.listAllReturnsOnCall[len(fake.listAllArgsForCall)]
	fake.listAllArgsForCall = append(fake.listAllArgsForCall, struct {
		arg1 context.Context
		arg2 *client.IsolationSegmentListOptions
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

func (fake *FakeCFIsolationSegmentClient) ListAllCallCount() int {
	fake.listAllMutex.RLock()
	defer fake.listAllMutex.RUnlock()
	return len(fake.listAllArgsForCall)
}

func (fake *FakeCFIsolationSegmentClient) ListAllCalls(stub func(context.Context, *client.IsolationSegmentListOptions) ([]*resource.IsolationSegment, error)) {
	fake.listAllMutex.Lock()
	defer fake.listAllMutex.Unlock()
	fake.ListAllStub = stub
}

func (fake *FakeCFIsolationSegmentClient) ListAllArgsForCall(i int) (context.Context, *client.IsolationSegmentListOptions) {
	fake.listAllMutex.RLock()
	defer fake.listAllMutex.RUnlock()
	argsForCall := fake.listAllArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFIsolationSegmentClient) ListAllReturns(result1 []*resource.IsolationSegment, result2 error) {
	fake.listAllMutex.Lock()
	defer fake.listAllMutex.Unlock()
	fake.ListAllStub = nil
	fake.listAllReturns = struct {
		result1 []*resource.IsolationSegment
		result2 error
	}{result1, result2}
}

func (fake *FakeCFIsolationSegmentClient) ListAllReturnsOnCall(i int, result1 []*resource.IsolationSegment, result2 error) {
	fake.listAllMutex.Lock()
	defer fake.listAllMutex.Unlock()
	fake.ListAllStub = nil
	if fake.listAllReturnsOnCall == nil {
		fake.listAllReturnsOnCall = make(map[int]struct {
			result1 []*resource.IsolationSegment
			result2 error
		})
	}
	fake.listAllReturnsOnCall[i] = struct {
		result1 []*resource.IsolationSegment
		result2 error
	}{result1, result2}
}

func (fake *FakeCFIsolationSegmentClient) RevokeOrganization(arg1 context.Context, arg2 string, arg3 string) error {
	fake.revokeOrganizationMutex.Lock()
	ret, specificReturn := fake.revokeOrganizationReturnsOnCall[len(fake.revokeOrganizationArgsForCall)]
	fake.revokeOrganizationArgsForCall = append(fake.revokeOrganizationArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.RevokeOrganizationStub
	fakeReturns := fake.revokeOrganizationReturns
	fake.recordInvocation("RevokeOrganization", []interface{}{arg1, arg2, arg3})
	fake.revokeOrganizationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCFIsolationSegmentClient) RevokeOrganizationCallCount() int {
	fake.revokeOrganizationMutex.RLock()
	defer fake.revokeOrganizationMutex.RUnlock()
	return len(fake.revokeOrganizationArgsForCall)
}

func (fake *FakeCFIsolationSegmentClient) RevokeOrganizationCalls(stub func(context.Context, string, string) error) {
	fake.revokeOrganizationMutex.Lock()
	defer fake.revokeOrganizationMutex.Unlock()
	fake.RevokeOrganizationStub = stub
}

func (fake *FakeCFIsolationSegmentClient) RevokeOrganizationArgsForCall(i int) (context.Context, string, string) {
	fake.revokeOrganizationMutex.RLock()
	defer fake.revokeOrganizationMutex.RUnlock()
	argsForCall := fake.revokeOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCFIsolationSegmentClient) RevokeOrganizationReturns(result1 error) {
	fake.revokeOrganizationMutex.Lock()
	defer fake.revokeOrganizationMutex.Unlock()
	fake.RevokeOrganizationStub = nil
	fake.revokeOrganizationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFIsolationSegmentClient) RevokeOrganizationReturnsOnCall(i int, result1 error) {
	fake.revokeOrganizationMutex.Lock()
	defer fake.revokeOrganizationMutex.Unlock()
	fake.RevokeOrganizationStub = nil
	if fake.revokeOrganizationReturnsOnCall == nil {
		fake.revokeOrganizationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.revokeOrganizationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFIsolationSegmentClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.entitleOrganizationMutex.RLock()
	defer fake.entitleOrganizationMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listAllMutex.RLock()
	defer fake.listAllMutex.RUnlock()
	fake.revokeOrganizationMutex.RLock()
	defer fake.revokeOrganizationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFIsolationSegmentClient) recordInvocation(key string, args []interface{}) {
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

var _ isosegment.CFIsolationSegmentClient = new(FakeCFIsolationSegmentClient)
