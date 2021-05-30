// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/Skarlso/gtui/models"
	"github.com/Skarlso/gtui/pkg/providers"
	"github.com/rivo/tview"
)

type FakeGithub struct {
	GetProjectStub        func(context.Context, int64) (*models.Project, error)
	getProjectMutex       sync.RWMutex
	getProjectArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	getProjectReturns struct {
		result1 *models.Project
		result2 error
	}
	getProjectReturnsOnCall map[int]struct {
		result1 *models.Project
		result2 error
	}
	GetProjectDataStub        func(context.Context, int64) (*models.ProjectData, error)
	getProjectDataMutex       sync.RWMutex
	getProjectDataArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	getProjectDataReturns struct {
		result1 *models.ProjectData
		result2 error
	}
	getProjectDataReturnsOnCall map[int]struct {
		result1 *models.ProjectData
		result2 error
	}
	ListOrganizationProjectsStub        func(context.Context, string, *models.ListOptions) ([]*models.Project, error)
	listOrganizationProjectsMutex       sync.RWMutex
	listOrganizationProjectsArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 *models.ListOptions
	}
	listOrganizationProjectsReturns struct {
		result1 []*models.Project
		result2 error
	}
	listOrganizationProjectsReturnsOnCall map[int]struct {
		result1 []*models.Project
		result2 error
	}
	ListRepositoryProjectsStub        func(context.Context, string, string, *models.ListOptions) ([]*models.Project, error)
	listRepositoryProjectsMutex       sync.RWMutex
	listRepositoryProjectsArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 *models.ListOptions
	}
	listRepositoryProjectsReturns struct {
		result1 []*models.Project
		result2 error
	}
	listRepositoryProjectsReturnsOnCall map[int]struct {
		result1 []*models.Project
		result2 error
	}
	LoadRestStub        func(context.Context, int64, *tview.List) error
	loadRestMutex       sync.RWMutex
	loadRestArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 *tview.List
	}
	loadRestReturns struct {
		result1 error
	}
	loadRestReturnsOnCall map[int]struct {
		result1 error
	}
	MoveAnIssueStub        func(context.Context, int64, int64) error
	moveAnIssueMutex       sync.RWMutex
	moveAnIssueArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 int64
	}
	moveAnIssueReturns struct {
		result1 error
	}
	moveAnIssueReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGithub) GetProject(arg1 context.Context, arg2 int64) (*models.Project, error) {
	fake.getProjectMutex.Lock()
	ret, specificReturn := fake.getProjectReturnsOnCall[len(fake.getProjectArgsForCall)]
	fake.getProjectArgsForCall = append(fake.getProjectArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	stub := fake.GetProjectStub
	fakeReturns := fake.getProjectReturns
	fake.recordInvocation("GetProject", []interface{}{arg1, arg2})
	fake.getProjectMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithub) GetProjectCallCount() int {
	fake.getProjectMutex.RLock()
	defer fake.getProjectMutex.RUnlock()
	return len(fake.getProjectArgsForCall)
}

func (fake *FakeGithub) GetProjectCalls(stub func(context.Context, int64) (*models.Project, error)) {
	fake.getProjectMutex.Lock()
	defer fake.getProjectMutex.Unlock()
	fake.GetProjectStub = stub
}

func (fake *FakeGithub) GetProjectArgsForCall(i int) (context.Context, int64) {
	fake.getProjectMutex.RLock()
	defer fake.getProjectMutex.RUnlock()
	argsForCall := fake.getProjectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGithub) GetProjectReturns(result1 *models.Project, result2 error) {
	fake.getProjectMutex.Lock()
	defer fake.getProjectMutex.Unlock()
	fake.GetProjectStub = nil
	fake.getProjectReturns = struct {
		result1 *models.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeGithub) GetProjectReturnsOnCall(i int, result1 *models.Project, result2 error) {
	fake.getProjectMutex.Lock()
	defer fake.getProjectMutex.Unlock()
	fake.GetProjectStub = nil
	if fake.getProjectReturnsOnCall == nil {
		fake.getProjectReturnsOnCall = make(map[int]struct {
			result1 *models.Project
			result2 error
		})
	}
	fake.getProjectReturnsOnCall[i] = struct {
		result1 *models.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeGithub) GetProjectData(arg1 context.Context, arg2 int64) (*models.ProjectData, error) {
	fake.getProjectDataMutex.Lock()
	ret, specificReturn := fake.getProjectDataReturnsOnCall[len(fake.getProjectDataArgsForCall)]
	fake.getProjectDataArgsForCall = append(fake.getProjectDataArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	stub := fake.GetProjectDataStub
	fakeReturns := fake.getProjectDataReturns
	fake.recordInvocation("GetProjectData", []interface{}{arg1, arg2})
	fake.getProjectDataMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithub) GetProjectDataCallCount() int {
	fake.getProjectDataMutex.RLock()
	defer fake.getProjectDataMutex.RUnlock()
	return len(fake.getProjectDataArgsForCall)
}

func (fake *FakeGithub) GetProjectDataCalls(stub func(context.Context, int64) (*models.ProjectData, error)) {
	fake.getProjectDataMutex.Lock()
	defer fake.getProjectDataMutex.Unlock()
	fake.GetProjectDataStub = stub
}

func (fake *FakeGithub) GetProjectDataArgsForCall(i int) (context.Context, int64) {
	fake.getProjectDataMutex.RLock()
	defer fake.getProjectDataMutex.RUnlock()
	argsForCall := fake.getProjectDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGithub) GetProjectDataReturns(result1 *models.ProjectData, result2 error) {
	fake.getProjectDataMutex.Lock()
	defer fake.getProjectDataMutex.Unlock()
	fake.GetProjectDataStub = nil
	fake.getProjectDataReturns = struct {
		result1 *models.ProjectData
		result2 error
	}{result1, result2}
}

func (fake *FakeGithub) GetProjectDataReturnsOnCall(i int, result1 *models.ProjectData, result2 error) {
	fake.getProjectDataMutex.Lock()
	defer fake.getProjectDataMutex.Unlock()
	fake.GetProjectDataStub = nil
	if fake.getProjectDataReturnsOnCall == nil {
		fake.getProjectDataReturnsOnCall = make(map[int]struct {
			result1 *models.ProjectData
			result2 error
		})
	}
	fake.getProjectDataReturnsOnCall[i] = struct {
		result1 *models.ProjectData
		result2 error
	}{result1, result2}
}

func (fake *FakeGithub) ListOrganizationProjects(arg1 context.Context, arg2 string, arg3 *models.ListOptions) ([]*models.Project, error) {
	fake.listOrganizationProjectsMutex.Lock()
	ret, specificReturn := fake.listOrganizationProjectsReturnsOnCall[len(fake.listOrganizationProjectsArgsForCall)]
	fake.listOrganizationProjectsArgsForCall = append(fake.listOrganizationProjectsArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 *models.ListOptions
	}{arg1, arg2, arg3})
	stub := fake.ListOrganizationProjectsStub
	fakeReturns := fake.listOrganizationProjectsReturns
	fake.recordInvocation("ListOrganizationProjects", []interface{}{arg1, arg2, arg3})
	fake.listOrganizationProjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithub) ListOrganizationProjectsCallCount() int {
	fake.listOrganizationProjectsMutex.RLock()
	defer fake.listOrganizationProjectsMutex.RUnlock()
	return len(fake.listOrganizationProjectsArgsForCall)
}

func (fake *FakeGithub) ListOrganizationProjectsCalls(stub func(context.Context, string, *models.ListOptions) ([]*models.Project, error)) {
	fake.listOrganizationProjectsMutex.Lock()
	defer fake.listOrganizationProjectsMutex.Unlock()
	fake.ListOrganizationProjectsStub = stub
}

func (fake *FakeGithub) ListOrganizationProjectsArgsForCall(i int) (context.Context, string, *models.ListOptions) {
	fake.listOrganizationProjectsMutex.RLock()
	defer fake.listOrganizationProjectsMutex.RUnlock()
	argsForCall := fake.listOrganizationProjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGithub) ListOrganizationProjectsReturns(result1 []*models.Project, result2 error) {
	fake.listOrganizationProjectsMutex.Lock()
	defer fake.listOrganizationProjectsMutex.Unlock()
	fake.ListOrganizationProjectsStub = nil
	fake.listOrganizationProjectsReturns = struct {
		result1 []*models.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeGithub) ListOrganizationProjectsReturnsOnCall(i int, result1 []*models.Project, result2 error) {
	fake.listOrganizationProjectsMutex.Lock()
	defer fake.listOrganizationProjectsMutex.Unlock()
	fake.ListOrganizationProjectsStub = nil
	if fake.listOrganizationProjectsReturnsOnCall == nil {
		fake.listOrganizationProjectsReturnsOnCall = make(map[int]struct {
			result1 []*models.Project
			result2 error
		})
	}
	fake.listOrganizationProjectsReturnsOnCall[i] = struct {
		result1 []*models.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeGithub) ListRepositoryProjects(arg1 context.Context, arg2 string, arg3 string, arg4 *models.ListOptions) ([]*models.Project, error) {
	fake.listRepositoryProjectsMutex.Lock()
	ret, specificReturn := fake.listRepositoryProjectsReturnsOnCall[len(fake.listRepositoryProjectsArgsForCall)]
	fake.listRepositoryProjectsArgsForCall = append(fake.listRepositoryProjectsArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 *models.ListOptions
	}{arg1, arg2, arg3, arg4})
	stub := fake.ListRepositoryProjectsStub
	fakeReturns := fake.listRepositoryProjectsReturns
	fake.recordInvocation("ListRepositoryProjects", []interface{}{arg1, arg2, arg3, arg4})
	fake.listRepositoryProjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithub) ListRepositoryProjectsCallCount() int {
	fake.listRepositoryProjectsMutex.RLock()
	defer fake.listRepositoryProjectsMutex.RUnlock()
	return len(fake.listRepositoryProjectsArgsForCall)
}

func (fake *FakeGithub) ListRepositoryProjectsCalls(stub func(context.Context, string, string, *models.ListOptions) ([]*models.Project, error)) {
	fake.listRepositoryProjectsMutex.Lock()
	defer fake.listRepositoryProjectsMutex.Unlock()
	fake.ListRepositoryProjectsStub = stub
}

func (fake *FakeGithub) ListRepositoryProjectsArgsForCall(i int) (context.Context, string, string, *models.ListOptions) {
	fake.listRepositoryProjectsMutex.RLock()
	defer fake.listRepositoryProjectsMutex.RUnlock()
	argsForCall := fake.listRepositoryProjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeGithub) ListRepositoryProjectsReturns(result1 []*models.Project, result2 error) {
	fake.listRepositoryProjectsMutex.Lock()
	defer fake.listRepositoryProjectsMutex.Unlock()
	fake.ListRepositoryProjectsStub = nil
	fake.listRepositoryProjectsReturns = struct {
		result1 []*models.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeGithub) ListRepositoryProjectsReturnsOnCall(i int, result1 []*models.Project, result2 error) {
	fake.listRepositoryProjectsMutex.Lock()
	defer fake.listRepositoryProjectsMutex.Unlock()
	fake.ListRepositoryProjectsStub = nil
	if fake.listRepositoryProjectsReturnsOnCall == nil {
		fake.listRepositoryProjectsReturnsOnCall = make(map[int]struct {
			result1 []*models.Project
			result2 error
		})
	}
	fake.listRepositoryProjectsReturnsOnCall[i] = struct {
		result1 []*models.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeGithub) LoadRest(arg1 context.Context, arg2 int64, arg3 *tview.List) error {
	fake.loadRestMutex.Lock()
	ret, specificReturn := fake.loadRestReturnsOnCall[len(fake.loadRestArgsForCall)]
	fake.loadRestArgsForCall = append(fake.loadRestArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 *tview.List
	}{arg1, arg2, arg3})
	stub := fake.LoadRestStub
	fakeReturns := fake.loadRestReturns
	fake.recordInvocation("LoadRest", []interface{}{arg1, arg2, arg3})
	fake.loadRestMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGithub) LoadRestCallCount() int {
	fake.loadRestMutex.RLock()
	defer fake.loadRestMutex.RUnlock()
	return len(fake.loadRestArgsForCall)
}

func (fake *FakeGithub) LoadRestCalls(stub func(context.Context, int64, *tview.List) error) {
	fake.loadRestMutex.Lock()
	defer fake.loadRestMutex.Unlock()
	fake.LoadRestStub = stub
}

func (fake *FakeGithub) LoadRestArgsForCall(i int) (context.Context, int64, *tview.List) {
	fake.loadRestMutex.RLock()
	defer fake.loadRestMutex.RUnlock()
	argsForCall := fake.loadRestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGithub) LoadRestReturns(result1 error) {
	fake.loadRestMutex.Lock()
	defer fake.loadRestMutex.Unlock()
	fake.LoadRestStub = nil
	fake.loadRestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGithub) LoadRestReturnsOnCall(i int, result1 error) {
	fake.loadRestMutex.Lock()
	defer fake.loadRestMutex.Unlock()
	fake.LoadRestStub = nil
	if fake.loadRestReturnsOnCall == nil {
		fake.loadRestReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loadRestReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGithub) MoveAnIssue(arg1 context.Context, arg2 int64, arg3 int64) error {
	fake.moveAnIssueMutex.Lock()
	ret, specificReturn := fake.moveAnIssueReturnsOnCall[len(fake.moveAnIssueArgsForCall)]
	fake.moveAnIssueArgsForCall = append(fake.moveAnIssueArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 int64
	}{arg1, arg2, arg3})
	stub := fake.MoveAnIssueStub
	fakeReturns := fake.moveAnIssueReturns
	fake.recordInvocation("MoveAnIssue", []interface{}{arg1, arg2, arg3})
	fake.moveAnIssueMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGithub) MoveAnIssueCallCount() int {
	fake.moveAnIssueMutex.RLock()
	defer fake.moveAnIssueMutex.RUnlock()
	return len(fake.moveAnIssueArgsForCall)
}

func (fake *FakeGithub) MoveAnIssueCalls(stub func(context.Context, int64, int64) error) {
	fake.moveAnIssueMutex.Lock()
	defer fake.moveAnIssueMutex.Unlock()
	fake.MoveAnIssueStub = stub
}

func (fake *FakeGithub) MoveAnIssueArgsForCall(i int) (context.Context, int64, int64) {
	fake.moveAnIssueMutex.RLock()
	defer fake.moveAnIssueMutex.RUnlock()
	argsForCall := fake.moveAnIssueArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGithub) MoveAnIssueReturns(result1 error) {
	fake.moveAnIssueMutex.Lock()
	defer fake.moveAnIssueMutex.Unlock()
	fake.MoveAnIssueStub = nil
	fake.moveAnIssueReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGithub) MoveAnIssueReturnsOnCall(i int, result1 error) {
	fake.moveAnIssueMutex.Lock()
	defer fake.moveAnIssueMutex.Unlock()
	fake.MoveAnIssueStub = nil
	if fake.moveAnIssueReturnsOnCall == nil {
		fake.moveAnIssueReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.moveAnIssueReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGithub) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getProjectMutex.RLock()
	defer fake.getProjectMutex.RUnlock()
	fake.getProjectDataMutex.RLock()
	defer fake.getProjectDataMutex.RUnlock()
	fake.listOrganizationProjectsMutex.RLock()
	defer fake.listOrganizationProjectsMutex.RUnlock()
	fake.listRepositoryProjectsMutex.RLock()
	defer fake.listRepositoryProjectsMutex.RUnlock()
	fake.loadRestMutex.RLock()
	defer fake.loadRestMutex.RUnlock()
	fake.moveAnIssueMutex.RLock()
	defer fake.moveAnIssueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGithub) recordInvocation(key string, args []interface{}) {
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

var _ providers.Github = new(FakeGithub)
