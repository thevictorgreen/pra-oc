// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: EventParsing)

package mocks

import (
	github "github.com/google/go-github/v31/github"
	azuredevops "github.com/mcdafydd/go-azuredevops/azuredevops"
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	go_gitlab "github.com/xanzy/go-gitlab"
	"reflect"
	"time"
)

type MockEventParsing struct {
	fail func(message string, callerSkip ...int)
}

func NewMockEventParsing(options ...pegomock.Option) *MockEventParsing {
	mock := &MockEventParsing{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockEventParsing) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockEventParsing) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockEventParsing) ParseGithubIssueCommentEvent(comment *github.IssueCommentEvent) (models.Repo, models.User, int, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{comment}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseGithubIssueCommentEvent", params, []reflect.Type{reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.User)(nil)).Elem(), reflect.TypeOf((*int)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.Repo
	var ret1 models.User
	var ret2 int
	var ret3 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.Repo)
		}
		if result[1] != nil {
			ret1 = result[1].(models.User)
		}
		if result[2] != nil {
			ret2 = result[2].(int)
		}
		if result[3] != nil {
			ret3 = result[3].(error)
		}
	}
	return ret0, ret1, ret2, ret3
}

func (mock *MockEventParsing) ParseGithubPull(ghPull *github.PullRequest) (models.PullRequest, models.Repo, models.Repo, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{ghPull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseGithubPull", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullRequest
	var ret1 models.Repo
	var ret2 models.Repo
	var ret3 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(models.Repo)
		}
		if result[2] != nil {
			ret2 = result[2].(models.Repo)
		}
		if result[3] != nil {
			ret3 = result[3].(error)
		}
	}
	return ret0, ret1, ret2, ret3
}

func (mock *MockEventParsing) ParseGithubPullEvent(pullEvent *github.PullRequestEvent) (models.PullRequest, models.PullRequestEventType, models.Repo, models.Repo, models.User, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{pullEvent}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseGithubPullEvent", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem(), reflect.TypeOf((*models.PullRequestEventType)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.User)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullRequest
	var ret1 models.PullRequestEventType
	var ret2 models.Repo
	var ret3 models.Repo
	var ret4 models.User
	var ret5 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(models.PullRequestEventType)
		}
		if result[2] != nil {
			ret2 = result[2].(models.Repo)
		}
		if result[3] != nil {
			ret3 = result[3].(models.Repo)
		}
		if result[4] != nil {
			ret4 = result[4].(models.User)
		}
		if result[5] != nil {
			ret5 = result[5].(error)
		}
	}
	return ret0, ret1, ret2, ret3, ret4, ret5
}

func (mock *MockEventParsing) ParseGithubRepo(ghRepo *github.Repository) (models.Repo, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{ghRepo}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseGithubRepo", params, []reflect.Type{reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.Repo
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.Repo)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEventParsing) ParseGitlabMergeRequestEvent(event go_gitlab.MergeEvent) (models.PullRequest, models.PullRequestEventType, models.Repo, models.Repo, models.User, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{event}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseGitlabMergeRequestEvent", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem(), reflect.TypeOf((*models.PullRequestEventType)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.User)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullRequest
	var ret1 models.PullRequestEventType
	var ret2 models.Repo
	var ret3 models.Repo
	var ret4 models.User
	var ret5 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(models.PullRequestEventType)
		}
		if result[2] != nil {
			ret2 = result[2].(models.Repo)
		}
		if result[3] != nil {
			ret3 = result[3].(models.Repo)
		}
		if result[4] != nil {
			ret4 = result[4].(models.User)
		}
		if result[5] != nil {
			ret5 = result[5].(error)
		}
	}
	return ret0, ret1, ret2, ret3, ret4, ret5
}

func (mock *MockEventParsing) ParseGitlabMergeRequestCommentEvent(event go_gitlab.MergeCommentEvent) (models.Repo, models.Repo, models.User, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{event}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseGitlabMergeRequestCommentEvent", params, []reflect.Type{reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.User)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.Repo
	var ret1 models.Repo
	var ret2 models.User
	var ret3 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.Repo)
		}
		if result[1] != nil {
			ret1 = result[1].(models.Repo)
		}
		if result[2] != nil {
			ret2 = result[2].(models.User)
		}
		if result[3] != nil {
			ret3 = result[3].(error)
		}
	}
	return ret0, ret1, ret2, ret3
}

func (mock *MockEventParsing) ParseGitlabMergeRequest(mr *go_gitlab.MergeRequest, baseRepo models.Repo) models.PullRequest {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{mr, baseRepo}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseGitlabMergeRequest", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem()})
	var ret0 models.PullRequest
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
	}
	return ret0
}

func (mock *MockEventParsing) ParseBitbucketCloudPullEvent(body []byte) (models.PullRequest, models.Repo, models.Repo, models.User, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{body}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseBitbucketCloudPullEvent", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.User)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullRequest
	var ret1 models.Repo
	var ret2 models.Repo
	var ret3 models.User
	var ret4 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(models.Repo)
		}
		if result[2] != nil {
			ret2 = result[2].(models.Repo)
		}
		if result[3] != nil {
			ret3 = result[3].(models.User)
		}
		if result[4] != nil {
			ret4 = result[4].(error)
		}
	}
	return ret0, ret1, ret2, ret3, ret4
}

func (mock *MockEventParsing) ParseBitbucketCloudPullCommentEvent(body []byte) (models.PullRequest, models.Repo, models.Repo, models.User, string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{body}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseBitbucketCloudPullCommentEvent", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.User)(nil)).Elem(), reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullRequest
	var ret1 models.Repo
	var ret2 models.Repo
	var ret3 models.User
	var ret4 string
	var ret5 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(models.Repo)
		}
		if result[2] != nil {
			ret2 = result[2].(models.Repo)
		}
		if result[3] != nil {
			ret3 = result[3].(models.User)
		}
		if result[4] != nil {
			ret4 = result[4].(string)
		}
		if result[5] != nil {
			ret5 = result[5].(error)
		}
	}
	return ret0, ret1, ret2, ret3, ret4, ret5
}

func (mock *MockEventParsing) GetBitbucketCloudPullEventType(eventTypeHeader string) models.PullRequestEventType {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{eventTypeHeader}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetBitbucketCloudPullEventType", params, []reflect.Type{reflect.TypeOf((*models.PullRequestEventType)(nil)).Elem()})
	var ret0 models.PullRequestEventType
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequestEventType)
		}
	}
	return ret0
}

func (mock *MockEventParsing) ParseBitbucketServerPullEvent(body []byte) (models.PullRequest, models.Repo, models.Repo, models.User, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{body}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseBitbucketServerPullEvent", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.User)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullRequest
	var ret1 models.Repo
	var ret2 models.Repo
	var ret3 models.User
	var ret4 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(models.Repo)
		}
		if result[2] != nil {
			ret2 = result[2].(models.Repo)
		}
		if result[3] != nil {
			ret3 = result[3].(models.User)
		}
		if result[4] != nil {
			ret4 = result[4].(error)
		}
	}
	return ret0, ret1, ret2, ret3, ret4
}

func (mock *MockEventParsing) ParseBitbucketServerPullCommentEvent(body []byte) (models.PullRequest, models.Repo, models.Repo, models.User, string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{body}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseBitbucketServerPullCommentEvent", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.User)(nil)).Elem(), reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullRequest
	var ret1 models.Repo
	var ret2 models.Repo
	var ret3 models.User
	var ret4 string
	var ret5 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(models.Repo)
		}
		if result[2] != nil {
			ret2 = result[2].(models.Repo)
		}
		if result[3] != nil {
			ret3 = result[3].(models.User)
		}
		if result[4] != nil {
			ret4 = result[4].(string)
		}
		if result[5] != nil {
			ret5 = result[5].(error)
		}
	}
	return ret0, ret1, ret2, ret3, ret4, ret5
}

func (mock *MockEventParsing) GetBitbucketServerPullEventType(eventTypeHeader string) models.PullRequestEventType {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{eventTypeHeader}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetBitbucketServerPullEventType", params, []reflect.Type{reflect.TypeOf((*models.PullRequestEventType)(nil)).Elem()})
	var ret0 models.PullRequestEventType
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequestEventType)
		}
	}
	return ret0
}

func (mock *MockEventParsing) ParseAzureDevopsPull(adPull *azuredevops.GitPullRequest) (models.PullRequest, models.Repo, models.Repo, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{adPull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseAzureDevopsPull", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullRequest
	var ret1 models.Repo
	var ret2 models.Repo
	var ret3 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(models.Repo)
		}
		if result[2] != nil {
			ret2 = result[2].(models.Repo)
		}
		if result[3] != nil {
			ret3 = result[3].(error)
		}
	}
	return ret0, ret1, ret2, ret3
}

func (mock *MockEventParsing) ParseAzureDevopsPullEvent(pullEvent azuredevops.Event) (models.PullRequest, models.PullRequestEventType, models.Repo, models.Repo, models.User, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{pullEvent}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseAzureDevopsPullEvent", params, []reflect.Type{reflect.TypeOf((*models.PullRequest)(nil)).Elem(), reflect.TypeOf((*models.PullRequestEventType)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*models.User)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullRequest
	var ret1 models.PullRequestEventType
	var ret2 models.Repo
	var ret3 models.Repo
	var ret4 models.User
	var ret5 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(models.PullRequestEventType)
		}
		if result[2] != nil {
			ret2 = result[2].(models.Repo)
		}
		if result[3] != nil {
			ret3 = result[3].(models.Repo)
		}
		if result[4] != nil {
			ret4 = result[4].(models.User)
		}
		if result[5] != nil {
			ret5 = result[5].(error)
		}
	}
	return ret0, ret1, ret2, ret3, ret4, ret5
}

func (mock *MockEventParsing) ParseAzureDevopsRepo(adRepo *azuredevops.GitRepository) (models.Repo, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEventParsing().")
	}
	params := []pegomock.Param{adRepo}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ParseAzureDevopsRepo", params, []reflect.Type{reflect.TypeOf((*models.Repo)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.Repo
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.Repo)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEventParsing) VerifyWasCalledOnce() *VerifierMockEventParsing {
	return &VerifierMockEventParsing{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockEventParsing) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockEventParsing {
	return &VerifierMockEventParsing{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockEventParsing) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockEventParsing {
	return &VerifierMockEventParsing{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockEventParsing) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockEventParsing {
	return &VerifierMockEventParsing{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockEventParsing struct {
	mock                   *MockEventParsing
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockEventParsing) ParseGithubIssueCommentEvent(comment *github.IssueCommentEvent) *MockEventParsing_ParseGithubIssueCommentEvent_OngoingVerification {
	params := []pegomock.Param{comment}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseGithubIssueCommentEvent", params, verifier.timeout)
	return &MockEventParsing_ParseGithubIssueCommentEvent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseGithubIssueCommentEvent_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseGithubIssueCommentEvent_OngoingVerification) GetCapturedArguments() *github.IssueCommentEvent {
	comment := c.GetAllCapturedArguments()
	return comment[len(comment)-1]
}

func (c *MockEventParsing_ParseGithubIssueCommentEvent_OngoingVerification) GetAllCapturedArguments() (_param0 []*github.IssueCommentEvent) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*github.IssueCommentEvent, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*github.IssueCommentEvent)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseGithubPull(ghPull *github.PullRequest) *MockEventParsing_ParseGithubPull_OngoingVerification {
	params := []pegomock.Param{ghPull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseGithubPull", params, verifier.timeout)
	return &MockEventParsing_ParseGithubPull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseGithubPull_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseGithubPull_OngoingVerification) GetCapturedArguments() *github.PullRequest {
	ghPull := c.GetAllCapturedArguments()
	return ghPull[len(ghPull)-1]
}

func (c *MockEventParsing_ParseGithubPull_OngoingVerification) GetAllCapturedArguments() (_param0 []*github.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*github.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*github.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseGithubPullEvent(pullEvent *github.PullRequestEvent) *MockEventParsing_ParseGithubPullEvent_OngoingVerification {
	params := []pegomock.Param{pullEvent}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseGithubPullEvent", params, verifier.timeout)
	return &MockEventParsing_ParseGithubPullEvent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseGithubPullEvent_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseGithubPullEvent_OngoingVerification) GetCapturedArguments() *github.PullRequestEvent {
	pullEvent := c.GetAllCapturedArguments()
	return pullEvent[len(pullEvent)-1]
}

func (c *MockEventParsing_ParseGithubPullEvent_OngoingVerification) GetAllCapturedArguments() (_param0 []*github.PullRequestEvent) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*github.PullRequestEvent, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*github.PullRequestEvent)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseGithubRepo(ghRepo *github.Repository) *MockEventParsing_ParseGithubRepo_OngoingVerification {
	params := []pegomock.Param{ghRepo}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseGithubRepo", params, verifier.timeout)
	return &MockEventParsing_ParseGithubRepo_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseGithubRepo_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseGithubRepo_OngoingVerification) GetCapturedArguments() *github.Repository {
	ghRepo := c.GetAllCapturedArguments()
	return ghRepo[len(ghRepo)-1]
}

func (c *MockEventParsing_ParseGithubRepo_OngoingVerification) GetAllCapturedArguments() (_param0 []*github.Repository) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*github.Repository, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*github.Repository)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseGitlabMergeRequestEvent(event go_gitlab.MergeEvent) *MockEventParsing_ParseGitlabMergeRequestEvent_OngoingVerification {
	params := []pegomock.Param{event}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseGitlabMergeRequestEvent", params, verifier.timeout)
	return &MockEventParsing_ParseGitlabMergeRequestEvent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseGitlabMergeRequestEvent_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseGitlabMergeRequestEvent_OngoingVerification) GetCapturedArguments() go_gitlab.MergeEvent {
	event := c.GetAllCapturedArguments()
	return event[len(event)-1]
}

func (c *MockEventParsing_ParseGitlabMergeRequestEvent_OngoingVerification) GetAllCapturedArguments() (_param0 []go_gitlab.MergeEvent) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]go_gitlab.MergeEvent, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(go_gitlab.MergeEvent)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseGitlabMergeRequestCommentEvent(event go_gitlab.MergeCommentEvent) *MockEventParsing_ParseGitlabMergeRequestCommentEvent_OngoingVerification {
	params := []pegomock.Param{event}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseGitlabMergeRequestCommentEvent", params, verifier.timeout)
	return &MockEventParsing_ParseGitlabMergeRequestCommentEvent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseGitlabMergeRequestCommentEvent_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseGitlabMergeRequestCommentEvent_OngoingVerification) GetCapturedArguments() go_gitlab.MergeCommentEvent {
	event := c.GetAllCapturedArguments()
	return event[len(event)-1]
}

func (c *MockEventParsing_ParseGitlabMergeRequestCommentEvent_OngoingVerification) GetAllCapturedArguments() (_param0 []go_gitlab.MergeCommentEvent) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]go_gitlab.MergeCommentEvent, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(go_gitlab.MergeCommentEvent)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseGitlabMergeRequest(mr *go_gitlab.MergeRequest, baseRepo models.Repo) *MockEventParsing_ParseGitlabMergeRequest_OngoingVerification {
	params := []pegomock.Param{mr, baseRepo}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseGitlabMergeRequest", params, verifier.timeout)
	return &MockEventParsing_ParseGitlabMergeRequest_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseGitlabMergeRequest_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseGitlabMergeRequest_OngoingVerification) GetCapturedArguments() (*go_gitlab.MergeRequest, models.Repo) {
	mr, baseRepo := c.GetAllCapturedArguments()
	return mr[len(mr)-1], baseRepo[len(baseRepo)-1]
}

func (c *MockEventParsing_ParseGitlabMergeRequest_OngoingVerification) GetAllCapturedArguments() (_param0 []*go_gitlab.MergeRequest, _param1 []models.Repo) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*go_gitlab.MergeRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*go_gitlab.MergeRequest)
		}
		_param1 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.Repo)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseBitbucketCloudPullEvent(body []byte) *MockEventParsing_ParseBitbucketCloudPullEvent_OngoingVerification {
	params := []pegomock.Param{body}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseBitbucketCloudPullEvent", params, verifier.timeout)
	return &MockEventParsing_ParseBitbucketCloudPullEvent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseBitbucketCloudPullEvent_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseBitbucketCloudPullEvent_OngoingVerification) GetCapturedArguments() []byte {
	body := c.GetAllCapturedArguments()
	return body[len(body)-1]
}

func (c *MockEventParsing_ParseBitbucketCloudPullEvent_OngoingVerification) GetAllCapturedArguments() (_param0 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]byte, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.([]byte)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseBitbucketCloudPullCommentEvent(body []byte) *MockEventParsing_ParseBitbucketCloudPullCommentEvent_OngoingVerification {
	params := []pegomock.Param{body}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseBitbucketCloudPullCommentEvent", params, verifier.timeout)
	return &MockEventParsing_ParseBitbucketCloudPullCommentEvent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseBitbucketCloudPullCommentEvent_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseBitbucketCloudPullCommentEvent_OngoingVerification) GetCapturedArguments() []byte {
	body := c.GetAllCapturedArguments()
	return body[len(body)-1]
}

func (c *MockEventParsing_ParseBitbucketCloudPullCommentEvent_OngoingVerification) GetAllCapturedArguments() (_param0 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]byte, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.([]byte)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) GetBitbucketCloudPullEventType(eventTypeHeader string) *MockEventParsing_GetBitbucketCloudPullEventType_OngoingVerification {
	params := []pegomock.Param{eventTypeHeader}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetBitbucketCloudPullEventType", params, verifier.timeout)
	return &MockEventParsing_GetBitbucketCloudPullEventType_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_GetBitbucketCloudPullEventType_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_GetBitbucketCloudPullEventType_OngoingVerification) GetCapturedArguments() string {
	eventTypeHeader := c.GetAllCapturedArguments()
	return eventTypeHeader[len(eventTypeHeader)-1]
}

func (c *MockEventParsing_GetBitbucketCloudPullEventType_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseBitbucketServerPullEvent(body []byte) *MockEventParsing_ParseBitbucketServerPullEvent_OngoingVerification {
	params := []pegomock.Param{body}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseBitbucketServerPullEvent", params, verifier.timeout)
	return &MockEventParsing_ParseBitbucketServerPullEvent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseBitbucketServerPullEvent_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseBitbucketServerPullEvent_OngoingVerification) GetCapturedArguments() []byte {
	body := c.GetAllCapturedArguments()
	return body[len(body)-1]
}

func (c *MockEventParsing_ParseBitbucketServerPullEvent_OngoingVerification) GetAllCapturedArguments() (_param0 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]byte, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.([]byte)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseBitbucketServerPullCommentEvent(body []byte) *MockEventParsing_ParseBitbucketServerPullCommentEvent_OngoingVerification {
	params := []pegomock.Param{body}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseBitbucketServerPullCommentEvent", params, verifier.timeout)
	return &MockEventParsing_ParseBitbucketServerPullCommentEvent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseBitbucketServerPullCommentEvent_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseBitbucketServerPullCommentEvent_OngoingVerification) GetCapturedArguments() []byte {
	body := c.GetAllCapturedArguments()
	return body[len(body)-1]
}

func (c *MockEventParsing_ParseBitbucketServerPullCommentEvent_OngoingVerification) GetAllCapturedArguments() (_param0 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]byte, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.([]byte)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) GetBitbucketServerPullEventType(eventTypeHeader string) *MockEventParsing_GetBitbucketServerPullEventType_OngoingVerification {
	params := []pegomock.Param{eventTypeHeader}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetBitbucketServerPullEventType", params, verifier.timeout)
	return &MockEventParsing_GetBitbucketServerPullEventType_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_GetBitbucketServerPullEventType_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_GetBitbucketServerPullEventType_OngoingVerification) GetCapturedArguments() string {
	eventTypeHeader := c.GetAllCapturedArguments()
	return eventTypeHeader[len(eventTypeHeader)-1]
}

func (c *MockEventParsing_GetBitbucketServerPullEventType_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseAzureDevopsPull(adPull *azuredevops.GitPullRequest) *MockEventParsing_ParseAzureDevopsPull_OngoingVerification {
	params := []pegomock.Param{adPull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseAzureDevopsPull", params, verifier.timeout)
	return &MockEventParsing_ParseAzureDevopsPull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseAzureDevopsPull_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseAzureDevopsPull_OngoingVerification) GetCapturedArguments() *azuredevops.GitPullRequest {
	adPull := c.GetAllCapturedArguments()
	return adPull[len(adPull)-1]
}

func (c *MockEventParsing_ParseAzureDevopsPull_OngoingVerification) GetAllCapturedArguments() (_param0 []*azuredevops.GitPullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*azuredevops.GitPullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*azuredevops.GitPullRequest)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseAzureDevopsPullEvent(pullEvent azuredevops.Event) *MockEventParsing_ParseAzureDevopsPullEvent_OngoingVerification {
	params := []pegomock.Param{pullEvent}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseAzureDevopsPullEvent", params, verifier.timeout)
	return &MockEventParsing_ParseAzureDevopsPullEvent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseAzureDevopsPullEvent_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseAzureDevopsPullEvent_OngoingVerification) GetCapturedArguments() azuredevops.Event {
	pullEvent := c.GetAllCapturedArguments()
	return pullEvent[len(pullEvent)-1]
}

func (c *MockEventParsing_ParseAzureDevopsPullEvent_OngoingVerification) GetAllCapturedArguments() (_param0 []azuredevops.Event) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]azuredevops.Event, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(azuredevops.Event)
		}
	}
	return
}

func (verifier *VerifierMockEventParsing) ParseAzureDevopsRepo(adRepo *azuredevops.GitRepository) *MockEventParsing_ParseAzureDevopsRepo_OngoingVerification {
	params := []pegomock.Param{adRepo}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ParseAzureDevopsRepo", params, verifier.timeout)
	return &MockEventParsing_ParseAzureDevopsRepo_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEventParsing_ParseAzureDevopsRepo_OngoingVerification struct {
	mock              *MockEventParsing
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEventParsing_ParseAzureDevopsRepo_OngoingVerification) GetCapturedArguments() *azuredevops.GitRepository {
	adRepo := c.GetAllCapturedArguments()
	return adRepo[len(adRepo)-1]
}

func (c *MockEventParsing_ParseAzureDevopsRepo_OngoingVerification) GetAllCapturedArguments() (_param0 []*azuredevops.GitRepository) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*azuredevops.GitRepository, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*azuredevops.GitRepository)
		}
	}
	return
}
