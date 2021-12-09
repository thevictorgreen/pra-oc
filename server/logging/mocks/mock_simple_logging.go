// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/logging (interfaces: SimpleLogging)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	logging "github.com/runatlantis/atlantis/server/logging"
	"reflect"
	"time"
)

type MockSimpleLogging struct {
	fail func(message string, callerSkip ...int)
}

func NewMockSimpleLogging(options ...pegomock.Option) *MockSimpleLogging {
	mock := &MockSimpleLogging{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockSimpleLogging) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockSimpleLogging) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockSimpleLogging) Debug(format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Debug", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) Info(format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Info", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) Warn(format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Warn", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) Err(format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Err", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) Log(level logging.LogLevel, format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{level, format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Log", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) SetLevel(lvl logging.LogLevel) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{lvl}
	pegomock.GetGenericMockFrom(mock).Invoke("SetLevel", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) With(a ...interface{}) logging.SimpleLogging {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{}
	for _, param := range a {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("With", params, []reflect.Type{reflect.TypeOf((*logging.SimpleLogging)(nil)).Elem()})
	var ret0 logging.SimpleLogging
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(logging.SimpleLogging)
		}
	}
	return ret0
}

func (mock *MockSimpleLogging) WithHistory(a ...interface{}) logging.SimpleLogging {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{}
	for _, param := range a {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("WithHistory", params, []reflect.Type{reflect.TypeOf((*logging.SimpleLogging)(nil)).Elem()})
	var ret0 logging.SimpleLogging
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(logging.SimpleLogging)
		}
	}
	return ret0
}

func (mock *MockSimpleLogging) GetHistory() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetHistory", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockSimpleLogging) Flush() error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Flush", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSimpleLogging) VerifyWasCalledOnce() *VerifierMockSimpleLogging {
	return &VerifierMockSimpleLogging{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockSimpleLogging) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockSimpleLogging {
	return &VerifierMockSimpleLogging{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockSimpleLogging) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockSimpleLogging {
	return &VerifierMockSimpleLogging{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockSimpleLogging) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockSimpleLogging {
	return &VerifierMockSimpleLogging{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockSimpleLogging struct {
	mock                   *MockSimpleLogging
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockSimpleLogging) Debug(format string, a ...interface{}) *MockSimpleLogging_Debug_OngoingVerification {
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Debug", params, verifier.timeout)
	return &MockSimpleLogging_Debug_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Debug_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Debug_OngoingVerification) GetCapturedArguments() (string, []interface{}) {
	format, a := c.GetAllCapturedArguments()
	return format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Debug_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param1[u] = make([]interface{}, len(params)-1)
			for x := 1; x < len(params); x++ {
				if params[x][u] != nil {
					_param1[u][x-1] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) Info(format string, a ...interface{}) *MockSimpleLogging_Info_OngoingVerification {
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Info", params, verifier.timeout)
	return &MockSimpleLogging_Info_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Info_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Info_OngoingVerification) GetCapturedArguments() (string, []interface{}) {
	format, a := c.GetAllCapturedArguments()
	return format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Info_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param1[u] = make([]interface{}, len(params)-1)
			for x := 1; x < len(params); x++ {
				if params[x][u] != nil {
					_param1[u][x-1] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) Warn(format string, a ...interface{}) *MockSimpleLogging_Warn_OngoingVerification {
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Warn", params, verifier.timeout)
	return &MockSimpleLogging_Warn_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Warn_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Warn_OngoingVerification) GetCapturedArguments() (string, []interface{}) {
	format, a := c.GetAllCapturedArguments()
	return format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Warn_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param1[u] = make([]interface{}, len(params)-1)
			for x := 1; x < len(params); x++ {
				if params[x][u] != nil {
					_param1[u][x-1] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) Err(format string, a ...interface{}) *MockSimpleLogging_Err_OngoingVerification {
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Err", params, verifier.timeout)
	return &MockSimpleLogging_Err_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Err_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Err_OngoingVerification) GetCapturedArguments() (string, []interface{}) {
	format, a := c.GetAllCapturedArguments()
	return format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Err_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param1[u] = make([]interface{}, len(params)-1)
			for x := 1; x < len(params); x++ {
				if params[x][u] != nil {
					_param1[u][x-1] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) Log(level logging.LogLevel, format string, a ...interface{}) *MockSimpleLogging_Log_OngoingVerification {
	params := []pegomock.Param{level, format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Log", params, verifier.timeout)
	return &MockSimpleLogging_Log_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Log_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Log_OngoingVerification) GetCapturedArguments() (logging.LogLevel, string, []interface{}) {
	level, format, a := c.GetAllCapturedArguments()
	return level[len(level)-1], format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Log_OngoingVerification) GetAllCapturedArguments() (_param0 []logging.LogLevel, _param1 []string, _param2 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]logging.LogLevel, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(logging.LogLevel)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param2[u] = make([]interface{}, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) SetLevel(lvl logging.LogLevel) *MockSimpleLogging_SetLevel_OngoingVerification {
	params := []pegomock.Param{lvl}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetLevel", params, verifier.timeout)
	return &MockSimpleLogging_SetLevel_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_SetLevel_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_SetLevel_OngoingVerification) GetCapturedArguments() logging.LogLevel {
	lvl := c.GetAllCapturedArguments()
	return lvl[len(lvl)-1]
}

func (c *MockSimpleLogging_SetLevel_OngoingVerification) GetAllCapturedArguments() (_param0 []logging.LogLevel) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]logging.LogLevel, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(logging.LogLevel)
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) With(a ...interface{}) *MockSimpleLogging_With_OngoingVerification {
	params := []pegomock.Param{}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "With", params, verifier.timeout)
	return &MockSimpleLogging_With_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_With_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_With_OngoingVerification) GetCapturedArguments() []interface{} {
	a := c.GetAllCapturedArguments()
	return a[len(a)-1]
}

func (c *MockSimpleLogging_With_OngoingVerification) GetAllCapturedArguments() (_param0 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param0[u] = make([]interface{}, len(params)-0)
			for x := 0; x < len(params); x++ {
				if params[x][u] != nil {
					_param0[u][x-0] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) WithHistory(a ...interface{}) *MockSimpleLogging_WithHistory_OngoingVerification {
	params := []pegomock.Param{}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "WithHistory", params, verifier.timeout)
	return &MockSimpleLogging_WithHistory_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_WithHistory_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_WithHistory_OngoingVerification) GetCapturedArguments() []interface{} {
	a := c.GetAllCapturedArguments()
	return a[len(a)-1]
}

func (c *MockSimpleLogging_WithHistory_OngoingVerification) GetAllCapturedArguments() (_param0 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param0[u] = make([]interface{}, len(params)-0)
			for x := 0; x < len(params); x++ {
				if params[x][u] != nil {
					_param0[u][x-0] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) GetHistory() *MockSimpleLogging_GetHistory_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetHistory", params, verifier.timeout)
	return &MockSimpleLogging_GetHistory_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_GetHistory_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_GetHistory_OngoingVerification) GetCapturedArguments() {
}

func (c *MockSimpleLogging_GetHistory_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockSimpleLogging) Flush() *MockSimpleLogging_Flush_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Flush", params, verifier.timeout)
	return &MockSimpleLogging_Flush_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Flush_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Flush_OngoingVerification) GetCapturedArguments() {
}

func (c *MockSimpleLogging_Flush_OngoingVerification) GetAllCapturedArguments() {
}