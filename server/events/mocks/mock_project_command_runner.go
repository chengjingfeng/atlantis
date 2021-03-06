// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: ProjectCommandRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockProjectCommandRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectCommandRunner(options ...pegomock.Option) *MockProjectCommandRunner {
	mock := &MockProjectCommandRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockProjectCommandRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockProjectCommandRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockProjectCommandRunner) Plan(ctx models.ProjectCommandContext) models.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Plan", params, []reflect.Type{reflect.TypeOf((*models.ProjectResult)(nil)).Elem()})
	var ret0 models.ProjectResult
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.ProjectResult)
		}
	}
	return ret0
}

func (mock *MockProjectCommandRunner) Apply(ctx models.ProjectCommandContext) models.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Apply", params, []reflect.Type{reflect.TypeOf((*models.ProjectResult)(nil)).Elem()})
	var ret0 models.ProjectResult
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.ProjectResult)
		}
	}
	return ret0
}

func (mock *MockProjectCommandRunner) VerifyWasCalledOnce() *VerifierProjectCommandRunner {
	return &VerifierProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectCommandRunner) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierProjectCommandRunner {
	return &VerifierProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectCommandRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierProjectCommandRunner {
	return &VerifierProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectCommandRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierProjectCommandRunner {
	return &VerifierProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierProjectCommandRunner struct {
	mock                   *MockProjectCommandRunner
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierProjectCommandRunner) Plan(ctx models.ProjectCommandContext) *ProjectCommandRunner_Plan_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Plan", params, verifier.timeout)
	return &ProjectCommandRunner_Plan_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ProjectCommandRunner_Plan_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *ProjectCommandRunner_Plan_OngoingVerification) GetCapturedArguments() models.ProjectCommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *ProjectCommandRunner_Plan_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
	}
	return
}

func (verifier *VerifierProjectCommandRunner) Apply(ctx models.ProjectCommandContext) *ProjectCommandRunner_Apply_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Apply", params, verifier.timeout)
	return &ProjectCommandRunner_Apply_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ProjectCommandRunner_Apply_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *ProjectCommandRunner_Apply_OngoingVerification) GetCapturedArguments() models.ProjectCommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *ProjectCommandRunner_Apply_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
	}
	return
}
