// Code generated by counterfeiter. DO NOT EDIT.
package fake_controllers

import (
	"context"
	"sync"

	"code.cloudfoundry.org/bbs/handlers"
	"code.cloudfoundry.org/bbs/models"
	lager "code.cloudfoundry.org/lager/v3"
)

type FakeEvacuationController struct {
	EvacuateClaimedActualLRPStub        func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) (bool, error)
	evacuateClaimedActualLRPMutex       sync.RWMutex
	evacuateClaimedActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
	}
	evacuateClaimedActualLRPReturns struct {
		result1 bool
		result2 error
	}
	evacuateClaimedActualLRPReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	EvacuateCrashedActualLRPStub        func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) error
	evacuateCrashedActualLRPMutex       sync.RWMutex
	evacuateCrashedActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
		arg5 string
	}
	evacuateCrashedActualLRPReturns struct {
		result1 error
	}
	evacuateCrashedActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	EvacuateRunningActualLRPStub        func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo, []*models.ActualLRPInternalRoute, map[string]string) (bool, error)
	evacuateRunningActualLRPMutex       sync.RWMutex
	evacuateRunningActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
		arg5 *models.ActualLRPNetInfo
		arg6 []*models.ActualLRPInternalRoute
		arg7 map[string]string
	}
	evacuateRunningActualLRPReturns struct {
		result1 bool
		result2 error
	}
	evacuateRunningActualLRPReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	EvacuateStoppedActualLRPStub        func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) error
	evacuateStoppedActualLRPMutex       sync.RWMutex
	evacuateStoppedActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
	}
	evacuateStoppedActualLRPReturns struct {
		result1 error
	}
	evacuateStoppedActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveEvacuatingActualLRPStub        func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) error
	removeEvacuatingActualLRPMutex       sync.RWMutex
	removeEvacuatingActualLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
	}
	removeEvacuatingActualLRPReturns struct {
		result1 error
	}
	removeEvacuatingActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.ActualLRPKey, arg4 *models.ActualLRPInstanceKey) (bool, error) {
	fake.evacuateClaimedActualLRPMutex.Lock()
	ret, specificReturn := fake.evacuateClaimedActualLRPReturnsOnCall[len(fake.evacuateClaimedActualLRPArgsForCall)]
	fake.evacuateClaimedActualLRPArgsForCall = append(fake.evacuateClaimedActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
	}{arg1, arg2, arg3, arg4})
	stub := fake.EvacuateClaimedActualLRPStub
	fakeReturns := fake.evacuateClaimedActualLRPReturns
	fake.recordInvocation("EvacuateClaimedActualLRP", []interface{}{arg1, arg2, arg3, arg4})
	fake.evacuateClaimedActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRPCallCount() int {
	fake.evacuateClaimedActualLRPMutex.RLock()
	defer fake.evacuateClaimedActualLRPMutex.RUnlock()
	return len(fake.evacuateClaimedActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRPCalls(stub func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) (bool, error)) {
	fake.evacuateClaimedActualLRPMutex.Lock()
	defer fake.evacuateClaimedActualLRPMutex.Unlock()
	fake.EvacuateClaimedActualLRPStub = stub
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRPArgsForCall(i int) (context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) {
	fake.evacuateClaimedActualLRPMutex.RLock()
	defer fake.evacuateClaimedActualLRPMutex.RUnlock()
	argsForCall := fake.evacuateClaimedActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRPReturns(result1 bool, result2 error) {
	fake.evacuateClaimedActualLRPMutex.Lock()
	defer fake.evacuateClaimedActualLRPMutex.Unlock()
	fake.EvacuateClaimedActualLRPStub = nil
	fake.evacuateClaimedActualLRPReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRPReturnsOnCall(i int, result1 bool, result2 error) {
	fake.evacuateClaimedActualLRPMutex.Lock()
	defer fake.evacuateClaimedActualLRPMutex.Unlock()
	fake.EvacuateClaimedActualLRPStub = nil
	if fake.evacuateClaimedActualLRPReturnsOnCall == nil {
		fake.evacuateClaimedActualLRPReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.evacuateClaimedActualLRPReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.ActualLRPKey, arg4 *models.ActualLRPInstanceKey, arg5 string) error {
	fake.evacuateCrashedActualLRPMutex.Lock()
	ret, specificReturn := fake.evacuateCrashedActualLRPReturnsOnCall[len(fake.evacuateCrashedActualLRPArgsForCall)]
	fake.evacuateCrashedActualLRPArgsForCall = append(fake.evacuateCrashedActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.EvacuateCrashedActualLRPStub
	fakeReturns := fake.evacuateCrashedActualLRPReturns
	fake.recordInvocation("EvacuateCrashedActualLRP", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.evacuateCrashedActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRPCallCount() int {
	fake.evacuateCrashedActualLRPMutex.RLock()
	defer fake.evacuateCrashedActualLRPMutex.RUnlock()
	return len(fake.evacuateCrashedActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRPCalls(stub func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) error) {
	fake.evacuateCrashedActualLRPMutex.Lock()
	defer fake.evacuateCrashedActualLRPMutex.Unlock()
	fake.EvacuateCrashedActualLRPStub = stub
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRPArgsForCall(i int) (context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) {
	fake.evacuateCrashedActualLRPMutex.RLock()
	defer fake.evacuateCrashedActualLRPMutex.RUnlock()
	argsForCall := fake.evacuateCrashedActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRPReturns(result1 error) {
	fake.evacuateCrashedActualLRPMutex.Lock()
	defer fake.evacuateCrashedActualLRPMutex.Unlock()
	fake.EvacuateCrashedActualLRPStub = nil
	fake.evacuateCrashedActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRPReturnsOnCall(i int, result1 error) {
	fake.evacuateCrashedActualLRPMutex.Lock()
	defer fake.evacuateCrashedActualLRPMutex.Unlock()
	fake.EvacuateCrashedActualLRPStub = nil
	if fake.evacuateCrashedActualLRPReturnsOnCall == nil {
		fake.evacuateCrashedActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evacuateCrashedActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.ActualLRPKey, arg4 *models.ActualLRPInstanceKey, arg5 *models.ActualLRPNetInfo, arg6 []*models.ActualLRPInternalRoute, arg7 map[string]string) (bool, error) {
	var arg6Copy []*models.ActualLRPInternalRoute
	if arg6 != nil {
		arg6Copy = make([]*models.ActualLRPInternalRoute, len(arg6))
		copy(arg6Copy, arg6)
	}
	fake.evacuateRunningActualLRPMutex.Lock()
	ret, specificReturn := fake.evacuateRunningActualLRPReturnsOnCall[len(fake.evacuateRunningActualLRPArgsForCall)]
	fake.evacuateRunningActualLRPArgsForCall = append(fake.evacuateRunningActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
		arg5 *models.ActualLRPNetInfo
		arg6 []*models.ActualLRPInternalRoute
		arg7 map[string]string
	}{arg1, arg2, arg3, arg4, arg5, arg6Copy, arg7})
	stub := fake.EvacuateRunningActualLRPStub
	fakeReturns := fake.evacuateRunningActualLRPReturns
	fake.recordInvocation("EvacuateRunningActualLRP", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6Copy, arg7})
	fake.evacuateRunningActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRPCallCount() int {
	fake.evacuateRunningActualLRPMutex.RLock()
	defer fake.evacuateRunningActualLRPMutex.RUnlock()
	return len(fake.evacuateRunningActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRPCalls(stub func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo, []*models.ActualLRPInternalRoute, map[string]string) (bool, error)) {
	fake.evacuateRunningActualLRPMutex.Lock()
	defer fake.evacuateRunningActualLRPMutex.Unlock()
	fake.EvacuateRunningActualLRPStub = stub
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRPArgsForCall(i int) (context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo, []*models.ActualLRPInternalRoute, map[string]string) {
	fake.evacuateRunningActualLRPMutex.RLock()
	defer fake.evacuateRunningActualLRPMutex.RUnlock()
	argsForCall := fake.evacuateRunningActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRPReturns(result1 bool, result2 error) {
	fake.evacuateRunningActualLRPMutex.Lock()
	defer fake.evacuateRunningActualLRPMutex.Unlock()
	fake.EvacuateRunningActualLRPStub = nil
	fake.evacuateRunningActualLRPReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRPReturnsOnCall(i int, result1 bool, result2 error) {
	fake.evacuateRunningActualLRPMutex.Lock()
	defer fake.evacuateRunningActualLRPMutex.Unlock()
	fake.EvacuateRunningActualLRPStub = nil
	if fake.evacuateRunningActualLRPReturnsOnCall == nil {
		fake.evacuateRunningActualLRPReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.evacuateRunningActualLRPReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.ActualLRPKey, arg4 *models.ActualLRPInstanceKey) error {
	fake.evacuateStoppedActualLRPMutex.Lock()
	ret, specificReturn := fake.evacuateStoppedActualLRPReturnsOnCall[len(fake.evacuateStoppedActualLRPArgsForCall)]
	fake.evacuateStoppedActualLRPArgsForCall = append(fake.evacuateStoppedActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
	}{arg1, arg2, arg3, arg4})
	stub := fake.EvacuateStoppedActualLRPStub
	fakeReturns := fake.evacuateStoppedActualLRPReturns
	fake.recordInvocation("EvacuateStoppedActualLRP", []interface{}{arg1, arg2, arg3, arg4})
	fake.evacuateStoppedActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRPCallCount() int {
	fake.evacuateStoppedActualLRPMutex.RLock()
	defer fake.evacuateStoppedActualLRPMutex.RUnlock()
	return len(fake.evacuateStoppedActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRPCalls(stub func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) error) {
	fake.evacuateStoppedActualLRPMutex.Lock()
	defer fake.evacuateStoppedActualLRPMutex.Unlock()
	fake.EvacuateStoppedActualLRPStub = stub
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRPArgsForCall(i int) (context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) {
	fake.evacuateStoppedActualLRPMutex.RLock()
	defer fake.evacuateStoppedActualLRPMutex.RUnlock()
	argsForCall := fake.evacuateStoppedActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRPReturns(result1 error) {
	fake.evacuateStoppedActualLRPMutex.Lock()
	defer fake.evacuateStoppedActualLRPMutex.Unlock()
	fake.EvacuateStoppedActualLRPStub = nil
	fake.evacuateStoppedActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRPReturnsOnCall(i int, result1 error) {
	fake.evacuateStoppedActualLRPMutex.Lock()
	defer fake.evacuateStoppedActualLRPMutex.Unlock()
	fake.EvacuateStoppedActualLRPStub = nil
	if fake.evacuateStoppedActualLRPReturnsOnCall == nil {
		fake.evacuateStoppedActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evacuateStoppedActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.ActualLRPKey, arg4 *models.ActualLRPInstanceKey) error {
	fake.removeEvacuatingActualLRPMutex.Lock()
	ret, specificReturn := fake.removeEvacuatingActualLRPReturnsOnCall[len(fake.removeEvacuatingActualLRPArgsForCall)]
	fake.removeEvacuatingActualLRPArgsForCall = append(fake.removeEvacuatingActualLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.ActualLRPKey
		arg4 *models.ActualLRPInstanceKey
	}{arg1, arg2, arg3, arg4})
	stub := fake.RemoveEvacuatingActualLRPStub
	fakeReturns := fake.removeEvacuatingActualLRPReturns
	fake.recordInvocation("RemoveEvacuatingActualLRP", []interface{}{arg1, arg2, arg3, arg4})
	fake.removeEvacuatingActualLRPMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRPCallCount() int {
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	return len(fake.removeEvacuatingActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRPCalls(stub func(context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) error) {
	fake.removeEvacuatingActualLRPMutex.Lock()
	defer fake.removeEvacuatingActualLRPMutex.Unlock()
	fake.RemoveEvacuatingActualLRPStub = stub
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRPArgsForCall(i int) (context.Context, lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) {
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	argsForCall := fake.removeEvacuatingActualLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRPReturns(result1 error) {
	fake.removeEvacuatingActualLRPMutex.Lock()
	defer fake.removeEvacuatingActualLRPMutex.Unlock()
	fake.RemoveEvacuatingActualLRPStub = nil
	fake.removeEvacuatingActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRPReturnsOnCall(i int, result1 error) {
	fake.removeEvacuatingActualLRPMutex.Lock()
	defer fake.removeEvacuatingActualLRPMutex.Unlock()
	fake.RemoveEvacuatingActualLRPStub = nil
	if fake.removeEvacuatingActualLRPReturnsOnCall == nil {
		fake.removeEvacuatingActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeEvacuatingActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.evacuateClaimedActualLRPMutex.RLock()
	defer fake.evacuateClaimedActualLRPMutex.RUnlock()
	fake.evacuateCrashedActualLRPMutex.RLock()
	defer fake.evacuateCrashedActualLRPMutex.RUnlock()
	fake.evacuateRunningActualLRPMutex.RLock()
	defer fake.evacuateRunningActualLRPMutex.RUnlock()
	fake.evacuateStoppedActualLRPMutex.RLock()
	defer fake.evacuateStoppedActualLRPMutex.RUnlock()
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEvacuationController) recordInvocation(key string, args []interface{}) {
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

var _ handlers.EvacuationController = new(FakeEvacuationController)
