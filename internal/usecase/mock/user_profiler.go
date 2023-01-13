package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/demimurg/twitter/internal/usecase.UserProfiler -o ./internal/usecase/mock/user_profiler.go -n UserProfilerMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/gojuno/minimock/v3"
)

// UserProfilerMock implements usecase.UserProfiler
type UserProfilerMock struct {
	t minimock.Tester

	funcDeactivate          func(ctx context.Context, userID int) (err error)
	inspectFuncDeactivate   func(ctx context.Context, userID int)
	afterDeactivateCounter  uint64
	beforeDeactivateCounter uint64
	DeactivateMock          mUserProfilerMockDeactivate

	funcLogin          func(ctx context.Context, email string) (up1 *entity.User, err error)
	inspectFuncLogin   func(ctx context.Context, email string)
	afterLoginCounter  uint64
	beforeLoginCounter uint64
	LoginMock          mUserProfilerMockLogin

	funcRegister          func(ctx context.Context, name string, email string, caption string, birthDate time.Time) (up1 *entity.User, err error)
	inspectFuncRegister   func(ctx context.Context, name string, email string, caption string, birthDate time.Time)
	afterRegisterCounter  uint64
	beforeRegisterCounter uint64
	RegisterMock          mUserProfilerMockRegister

	funcUpdateCaption          func(ctx context.Context, userID int, newCaption string) (err error)
	inspectFuncUpdateCaption   func(ctx context.Context, userID int, newCaption string)
	afterUpdateCaptionCounter  uint64
	beforeUpdateCaptionCounter uint64
	UpdateCaptionMock          mUserProfilerMockUpdateCaption
}

// NewUserProfilerMock returns a mock for usecase.UserProfiler
func NewUserProfilerMock(t minimock.Tester) *UserProfilerMock {
	m := &UserProfilerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeactivateMock = mUserProfilerMockDeactivate{mock: m}
	m.DeactivateMock.callArgs = []*UserProfilerMockDeactivateParams{}

	m.LoginMock = mUserProfilerMockLogin{mock: m}
	m.LoginMock.callArgs = []*UserProfilerMockLoginParams{}

	m.RegisterMock = mUserProfilerMockRegister{mock: m}
	m.RegisterMock.callArgs = []*UserProfilerMockRegisterParams{}

	m.UpdateCaptionMock = mUserProfilerMockUpdateCaption{mock: m}
	m.UpdateCaptionMock.callArgs = []*UserProfilerMockUpdateCaptionParams{}

	return m
}

type mUserProfilerMockDeactivate struct {
	mock               *UserProfilerMock
	defaultExpectation *UserProfilerMockDeactivateExpectation
	expectations       []*UserProfilerMockDeactivateExpectation

	callArgs []*UserProfilerMockDeactivateParams
	mutex    sync.RWMutex
}

// UserProfilerMockDeactivateExpectation specifies expectation struct of the UserProfiler.Deactivate
type UserProfilerMockDeactivateExpectation struct {
	mock    *UserProfilerMock
	params  *UserProfilerMockDeactivateParams
	results *UserProfilerMockDeactivateResults
	Counter uint64
}

// UserProfilerMockDeactivateParams contains parameters of the UserProfiler.Deactivate
type UserProfilerMockDeactivateParams struct {
	ctx    context.Context
	userID int
}

// UserProfilerMockDeactivateResults contains results of the UserProfiler.Deactivate
type UserProfilerMockDeactivateResults struct {
	err error
}

// Expect sets up expected params for UserProfiler.Deactivate
func (mmDeactivate *mUserProfilerMockDeactivate) Expect(ctx context.Context, userID int) *mUserProfilerMockDeactivate {
	if mmDeactivate.mock.funcDeactivate != nil {
		mmDeactivate.mock.t.Fatalf("UserProfilerMock.Deactivate mock is already set by Set")
	}

	if mmDeactivate.defaultExpectation == nil {
		mmDeactivate.defaultExpectation = &UserProfilerMockDeactivateExpectation{}
	}

	mmDeactivate.defaultExpectation.params = &UserProfilerMockDeactivateParams{ctx, userID}
	for _, e := range mmDeactivate.expectations {
		if minimock.Equal(e.params, mmDeactivate.defaultExpectation.params) {
			mmDeactivate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeactivate.defaultExpectation.params)
		}
	}

	return mmDeactivate
}

// Inspect accepts an inspector function that has same arguments as the UserProfiler.Deactivate
func (mmDeactivate *mUserProfilerMockDeactivate) Inspect(f func(ctx context.Context, userID int)) *mUserProfilerMockDeactivate {
	if mmDeactivate.mock.inspectFuncDeactivate != nil {
		mmDeactivate.mock.t.Fatalf("Inspect function is already set for UserProfilerMock.Deactivate")
	}

	mmDeactivate.mock.inspectFuncDeactivate = f

	return mmDeactivate
}

// Return sets up results that will be returned by UserProfiler.Deactivate
func (mmDeactivate *mUserProfilerMockDeactivate) Return(err error) *UserProfilerMock {
	if mmDeactivate.mock.funcDeactivate != nil {
		mmDeactivate.mock.t.Fatalf("UserProfilerMock.Deactivate mock is already set by Set")
	}

	if mmDeactivate.defaultExpectation == nil {
		mmDeactivate.defaultExpectation = &UserProfilerMockDeactivateExpectation{mock: mmDeactivate.mock}
	}
	mmDeactivate.defaultExpectation.results = &UserProfilerMockDeactivateResults{err}
	return mmDeactivate.mock
}

// Set uses given function f to mock the UserProfiler.Deactivate method
func (mmDeactivate *mUserProfilerMockDeactivate) Set(f func(ctx context.Context, userID int) (err error)) *UserProfilerMock {
	if mmDeactivate.defaultExpectation != nil {
		mmDeactivate.mock.t.Fatalf("Default expectation is already set for the UserProfiler.Deactivate method")
	}

	if len(mmDeactivate.expectations) > 0 {
		mmDeactivate.mock.t.Fatalf("Some expectations are already set for the UserProfiler.Deactivate method")
	}

	mmDeactivate.mock.funcDeactivate = f
	return mmDeactivate.mock
}

// When sets expectation for the UserProfiler.Deactivate which will trigger the result defined by the following
// Then helper
func (mmDeactivate *mUserProfilerMockDeactivate) When(ctx context.Context, userID int) *UserProfilerMockDeactivateExpectation {
	if mmDeactivate.mock.funcDeactivate != nil {
		mmDeactivate.mock.t.Fatalf("UserProfilerMock.Deactivate mock is already set by Set")
	}

	expectation := &UserProfilerMockDeactivateExpectation{
		mock:   mmDeactivate.mock,
		params: &UserProfilerMockDeactivateParams{ctx, userID},
	}
	mmDeactivate.expectations = append(mmDeactivate.expectations, expectation)
	return expectation
}

// Then sets up UserProfiler.Deactivate return parameters for the expectation previously defined by the When method
func (e *UserProfilerMockDeactivateExpectation) Then(err error) *UserProfilerMock {
	e.results = &UserProfilerMockDeactivateResults{err}
	return e.mock
}

// Deactivate implements usecase.UserProfiler
func (mmDeactivate *UserProfilerMock) Deactivate(ctx context.Context, userID int) (err error) {
	mm_atomic.AddUint64(&mmDeactivate.beforeDeactivateCounter, 1)
	defer mm_atomic.AddUint64(&mmDeactivate.afterDeactivateCounter, 1)

	if mmDeactivate.inspectFuncDeactivate != nil {
		mmDeactivate.inspectFuncDeactivate(ctx, userID)
	}

	mm_params := &UserProfilerMockDeactivateParams{ctx, userID}

	// Record call args
	mmDeactivate.DeactivateMock.mutex.Lock()
	mmDeactivate.DeactivateMock.callArgs = append(mmDeactivate.DeactivateMock.callArgs, mm_params)
	mmDeactivate.DeactivateMock.mutex.Unlock()

	for _, e := range mmDeactivate.DeactivateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeactivate.DeactivateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeactivate.DeactivateMock.defaultExpectation.Counter, 1)
		mm_want := mmDeactivate.DeactivateMock.defaultExpectation.params
		mm_got := UserProfilerMockDeactivateParams{ctx, userID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeactivate.t.Errorf("UserProfilerMock.Deactivate got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeactivate.DeactivateMock.defaultExpectation.results
		if mm_results == nil {
			mmDeactivate.t.Fatal("No results are set for the UserProfilerMock.Deactivate")
		}
		return (*mm_results).err
	}
	if mmDeactivate.funcDeactivate != nil {
		return mmDeactivate.funcDeactivate(ctx, userID)
	}
	mmDeactivate.t.Fatalf("Unexpected call to UserProfilerMock.Deactivate. %v %v", ctx, userID)
	return
}

// DeactivateAfterCounter returns a count of finished UserProfilerMock.Deactivate invocations
func (mmDeactivate *UserProfilerMock) DeactivateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeactivate.afterDeactivateCounter)
}

// DeactivateBeforeCounter returns a count of UserProfilerMock.Deactivate invocations
func (mmDeactivate *UserProfilerMock) DeactivateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeactivate.beforeDeactivateCounter)
}

// Calls returns a list of arguments used in each call to UserProfilerMock.Deactivate.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeactivate *mUserProfilerMockDeactivate) Calls() []*UserProfilerMockDeactivateParams {
	mmDeactivate.mutex.RLock()

	argCopy := make([]*UserProfilerMockDeactivateParams, len(mmDeactivate.callArgs))
	copy(argCopy, mmDeactivate.callArgs)

	mmDeactivate.mutex.RUnlock()

	return argCopy
}

// MinimockDeactivateDone returns true if the count of the Deactivate invocations corresponds
// the number of defined expectations
func (m *UserProfilerMock) MinimockDeactivateDone() bool {
	for _, e := range m.DeactivateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeactivateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeactivateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeactivate != nil && mm_atomic.LoadUint64(&m.afterDeactivateCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeactivateInspect logs each unmet expectation
func (m *UserProfilerMock) MinimockDeactivateInspect() {
	for _, e := range m.DeactivateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserProfilerMock.Deactivate with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeactivateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeactivateCounter) < 1 {
		if m.DeactivateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserProfilerMock.Deactivate")
		} else {
			m.t.Errorf("Expected call to UserProfilerMock.Deactivate with params: %#v", *m.DeactivateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeactivate != nil && mm_atomic.LoadUint64(&m.afterDeactivateCounter) < 1 {
		m.t.Error("Expected call to UserProfilerMock.Deactivate")
	}
}

type mUserProfilerMockLogin struct {
	mock               *UserProfilerMock
	defaultExpectation *UserProfilerMockLoginExpectation
	expectations       []*UserProfilerMockLoginExpectation

	callArgs []*UserProfilerMockLoginParams
	mutex    sync.RWMutex
}

// UserProfilerMockLoginExpectation specifies expectation struct of the UserProfiler.Login
type UserProfilerMockLoginExpectation struct {
	mock    *UserProfilerMock
	params  *UserProfilerMockLoginParams
	results *UserProfilerMockLoginResults
	Counter uint64
}

// UserProfilerMockLoginParams contains parameters of the UserProfiler.Login
type UserProfilerMockLoginParams struct {
	ctx   context.Context
	email string
}

// UserProfilerMockLoginResults contains results of the UserProfiler.Login
type UserProfilerMockLoginResults struct {
	up1 *entity.User
	err error
}

// Expect sets up expected params for UserProfiler.Login
func (mmLogin *mUserProfilerMockLogin) Expect(ctx context.Context, email string) *mUserProfilerMockLogin {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("UserProfilerMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &UserProfilerMockLoginExpectation{}
	}

	mmLogin.defaultExpectation.params = &UserProfilerMockLoginParams{ctx, email}
	for _, e := range mmLogin.expectations {
		if minimock.Equal(e.params, mmLogin.defaultExpectation.params) {
			mmLogin.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLogin.defaultExpectation.params)
		}
	}

	return mmLogin
}

// Inspect accepts an inspector function that has same arguments as the UserProfiler.Login
func (mmLogin *mUserProfilerMockLogin) Inspect(f func(ctx context.Context, email string)) *mUserProfilerMockLogin {
	if mmLogin.mock.inspectFuncLogin != nil {
		mmLogin.mock.t.Fatalf("Inspect function is already set for UserProfilerMock.Login")
	}

	mmLogin.mock.inspectFuncLogin = f

	return mmLogin
}

// Return sets up results that will be returned by UserProfiler.Login
func (mmLogin *mUserProfilerMockLogin) Return(up1 *entity.User, err error) *UserProfilerMock {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("UserProfilerMock.Login mock is already set by Set")
	}

	if mmLogin.defaultExpectation == nil {
		mmLogin.defaultExpectation = &UserProfilerMockLoginExpectation{mock: mmLogin.mock}
	}
	mmLogin.defaultExpectation.results = &UserProfilerMockLoginResults{up1, err}
	return mmLogin.mock
}

// Set uses given function f to mock the UserProfiler.Login method
func (mmLogin *mUserProfilerMockLogin) Set(f func(ctx context.Context, email string) (up1 *entity.User, err error)) *UserProfilerMock {
	if mmLogin.defaultExpectation != nil {
		mmLogin.mock.t.Fatalf("Default expectation is already set for the UserProfiler.Login method")
	}

	if len(mmLogin.expectations) > 0 {
		mmLogin.mock.t.Fatalf("Some expectations are already set for the UserProfiler.Login method")
	}

	mmLogin.mock.funcLogin = f
	return mmLogin.mock
}

// When sets expectation for the UserProfiler.Login which will trigger the result defined by the following
// Then helper
func (mmLogin *mUserProfilerMockLogin) When(ctx context.Context, email string) *UserProfilerMockLoginExpectation {
	if mmLogin.mock.funcLogin != nil {
		mmLogin.mock.t.Fatalf("UserProfilerMock.Login mock is already set by Set")
	}

	expectation := &UserProfilerMockLoginExpectation{
		mock:   mmLogin.mock,
		params: &UserProfilerMockLoginParams{ctx, email},
	}
	mmLogin.expectations = append(mmLogin.expectations, expectation)
	return expectation
}

// Then sets up UserProfiler.Login return parameters for the expectation previously defined by the When method
func (e *UserProfilerMockLoginExpectation) Then(up1 *entity.User, err error) *UserProfilerMock {
	e.results = &UserProfilerMockLoginResults{up1, err}
	return e.mock
}

// Login implements usecase.UserProfiler
func (mmLogin *UserProfilerMock) Login(ctx context.Context, email string) (up1 *entity.User, err error) {
	mm_atomic.AddUint64(&mmLogin.beforeLoginCounter, 1)
	defer mm_atomic.AddUint64(&mmLogin.afterLoginCounter, 1)

	if mmLogin.inspectFuncLogin != nil {
		mmLogin.inspectFuncLogin(ctx, email)
	}

	mm_params := &UserProfilerMockLoginParams{ctx, email}

	// Record call args
	mmLogin.LoginMock.mutex.Lock()
	mmLogin.LoginMock.callArgs = append(mmLogin.LoginMock.callArgs, mm_params)
	mmLogin.LoginMock.mutex.Unlock()

	for _, e := range mmLogin.LoginMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmLogin.LoginMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLogin.LoginMock.defaultExpectation.Counter, 1)
		mm_want := mmLogin.LoginMock.defaultExpectation.params
		mm_got := UserProfilerMockLoginParams{ctx, email}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLogin.t.Errorf("UserProfilerMock.Login got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLogin.LoginMock.defaultExpectation.results
		if mm_results == nil {
			mmLogin.t.Fatal("No results are set for the UserProfilerMock.Login")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmLogin.funcLogin != nil {
		return mmLogin.funcLogin(ctx, email)
	}
	mmLogin.t.Fatalf("Unexpected call to UserProfilerMock.Login. %v %v", ctx, email)
	return
}

// LoginAfterCounter returns a count of finished UserProfilerMock.Login invocations
func (mmLogin *UserProfilerMock) LoginAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLogin.afterLoginCounter)
}

// LoginBeforeCounter returns a count of UserProfilerMock.Login invocations
func (mmLogin *UserProfilerMock) LoginBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLogin.beforeLoginCounter)
}

// Calls returns a list of arguments used in each call to UserProfilerMock.Login.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLogin *mUserProfilerMockLogin) Calls() []*UserProfilerMockLoginParams {
	mmLogin.mutex.RLock()

	argCopy := make([]*UserProfilerMockLoginParams, len(mmLogin.callArgs))
	copy(argCopy, mmLogin.callArgs)

	mmLogin.mutex.RUnlock()

	return argCopy
}

// MinimockLoginDone returns true if the count of the Login invocations corresponds
// the number of defined expectations
func (m *UserProfilerMock) MinimockLoginDone() bool {
	for _, e := range m.LoginMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoginMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLogin != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		return false
	}
	return true
}

// MinimockLoginInspect logs each unmet expectation
func (m *UserProfilerMock) MinimockLoginInspect() {
	for _, e := range m.LoginMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserProfilerMock.Login with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LoginMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		if m.LoginMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserProfilerMock.Login")
		} else {
			m.t.Errorf("Expected call to UserProfilerMock.Login with params: %#v", *m.LoginMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLogin != nil && mm_atomic.LoadUint64(&m.afterLoginCounter) < 1 {
		m.t.Error("Expected call to UserProfilerMock.Login")
	}
}

type mUserProfilerMockRegister struct {
	mock               *UserProfilerMock
	defaultExpectation *UserProfilerMockRegisterExpectation
	expectations       []*UserProfilerMockRegisterExpectation

	callArgs []*UserProfilerMockRegisterParams
	mutex    sync.RWMutex
}

// UserProfilerMockRegisterExpectation specifies expectation struct of the UserProfiler.Register
type UserProfilerMockRegisterExpectation struct {
	mock    *UserProfilerMock
	params  *UserProfilerMockRegisterParams
	results *UserProfilerMockRegisterResults
	Counter uint64
}

// UserProfilerMockRegisterParams contains parameters of the UserProfiler.Register
type UserProfilerMockRegisterParams struct {
	ctx       context.Context
	name      string
	email     string
	caption   string
	birthDate time.Time
}

// UserProfilerMockRegisterResults contains results of the UserProfiler.Register
type UserProfilerMockRegisterResults struct {
	up1 *entity.User
	err error
}

// Expect sets up expected params for UserProfiler.Register
func (mmRegister *mUserProfilerMockRegister) Expect(ctx context.Context, name string, email string, caption string, birthDate time.Time) *mUserProfilerMockRegister {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserProfilerMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &UserProfilerMockRegisterExpectation{}
	}

	mmRegister.defaultExpectation.params = &UserProfilerMockRegisterParams{ctx, name, email, caption, birthDate}
	for _, e := range mmRegister.expectations {
		if minimock.Equal(e.params, mmRegister.defaultExpectation.params) {
			mmRegister.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRegister.defaultExpectation.params)
		}
	}

	return mmRegister
}

// Inspect accepts an inspector function that has same arguments as the UserProfiler.Register
func (mmRegister *mUserProfilerMockRegister) Inspect(f func(ctx context.Context, name string, email string, caption string, birthDate time.Time)) *mUserProfilerMockRegister {
	if mmRegister.mock.inspectFuncRegister != nil {
		mmRegister.mock.t.Fatalf("Inspect function is already set for UserProfilerMock.Register")
	}

	mmRegister.mock.inspectFuncRegister = f

	return mmRegister
}

// Return sets up results that will be returned by UserProfiler.Register
func (mmRegister *mUserProfilerMockRegister) Return(up1 *entity.User, err error) *UserProfilerMock {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserProfilerMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &UserProfilerMockRegisterExpectation{mock: mmRegister.mock}
	}
	mmRegister.defaultExpectation.results = &UserProfilerMockRegisterResults{up1, err}
	return mmRegister.mock
}

// Set uses given function f to mock the UserProfiler.Register method
func (mmRegister *mUserProfilerMockRegister) Set(f func(ctx context.Context, name string, email string, caption string, birthDate time.Time) (up1 *entity.User, err error)) *UserProfilerMock {
	if mmRegister.defaultExpectation != nil {
		mmRegister.mock.t.Fatalf("Default expectation is already set for the UserProfiler.Register method")
	}

	if len(mmRegister.expectations) > 0 {
		mmRegister.mock.t.Fatalf("Some expectations are already set for the UserProfiler.Register method")
	}

	mmRegister.mock.funcRegister = f
	return mmRegister.mock
}

// When sets expectation for the UserProfiler.Register which will trigger the result defined by the following
// Then helper
func (mmRegister *mUserProfilerMockRegister) When(ctx context.Context, name string, email string, caption string, birthDate time.Time) *UserProfilerMockRegisterExpectation {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserProfilerMock.Register mock is already set by Set")
	}

	expectation := &UserProfilerMockRegisterExpectation{
		mock:   mmRegister.mock,
		params: &UserProfilerMockRegisterParams{ctx, name, email, caption, birthDate},
	}
	mmRegister.expectations = append(mmRegister.expectations, expectation)
	return expectation
}

// Then sets up UserProfiler.Register return parameters for the expectation previously defined by the When method
func (e *UserProfilerMockRegisterExpectation) Then(up1 *entity.User, err error) *UserProfilerMock {
	e.results = &UserProfilerMockRegisterResults{up1, err}
	return e.mock
}

// Register implements usecase.UserProfiler
func (mmRegister *UserProfilerMock) Register(ctx context.Context, name string, email string, caption string, birthDate time.Time) (up1 *entity.User, err error) {
	mm_atomic.AddUint64(&mmRegister.beforeRegisterCounter, 1)
	defer mm_atomic.AddUint64(&mmRegister.afterRegisterCounter, 1)

	if mmRegister.inspectFuncRegister != nil {
		mmRegister.inspectFuncRegister(ctx, name, email, caption, birthDate)
	}

	mm_params := &UserProfilerMockRegisterParams{ctx, name, email, caption, birthDate}

	// Record call args
	mmRegister.RegisterMock.mutex.Lock()
	mmRegister.RegisterMock.callArgs = append(mmRegister.RegisterMock.callArgs, mm_params)
	mmRegister.RegisterMock.mutex.Unlock()

	for _, e := range mmRegister.RegisterMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmRegister.RegisterMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRegister.RegisterMock.defaultExpectation.Counter, 1)
		mm_want := mmRegister.RegisterMock.defaultExpectation.params
		mm_got := UserProfilerMockRegisterParams{ctx, name, email, caption, birthDate}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRegister.t.Errorf("UserProfilerMock.Register got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRegister.RegisterMock.defaultExpectation.results
		if mm_results == nil {
			mmRegister.t.Fatal("No results are set for the UserProfilerMock.Register")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmRegister.funcRegister != nil {
		return mmRegister.funcRegister(ctx, name, email, caption, birthDate)
	}
	mmRegister.t.Fatalf("Unexpected call to UserProfilerMock.Register. %v %v %v %v %v", ctx, name, email, caption, birthDate)
	return
}

// RegisterAfterCounter returns a count of finished UserProfilerMock.Register invocations
func (mmRegister *UserProfilerMock) RegisterAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegister.afterRegisterCounter)
}

// RegisterBeforeCounter returns a count of UserProfilerMock.Register invocations
func (mmRegister *UserProfilerMock) RegisterBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegister.beforeRegisterCounter)
}

// Calls returns a list of arguments used in each call to UserProfilerMock.Register.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRegister *mUserProfilerMockRegister) Calls() []*UserProfilerMockRegisterParams {
	mmRegister.mutex.RLock()

	argCopy := make([]*UserProfilerMockRegisterParams, len(mmRegister.callArgs))
	copy(argCopy, mmRegister.callArgs)

	mmRegister.mutex.RUnlock()

	return argCopy
}

// MinimockRegisterDone returns true if the count of the Register invocations corresponds
// the number of defined expectations
func (m *UserProfilerMock) MinimockRegisterDone() bool {
	for _, e := range m.RegisterMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RegisterMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegister != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		return false
	}
	return true
}

// MinimockRegisterInspect logs each unmet expectation
func (m *UserProfilerMock) MinimockRegisterInspect() {
	for _, e := range m.RegisterMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserProfilerMock.Register with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RegisterMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		if m.RegisterMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserProfilerMock.Register")
		} else {
			m.t.Errorf("Expected call to UserProfilerMock.Register with params: %#v", *m.RegisterMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegister != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		m.t.Error("Expected call to UserProfilerMock.Register")
	}
}

type mUserProfilerMockUpdateCaption struct {
	mock               *UserProfilerMock
	defaultExpectation *UserProfilerMockUpdateCaptionExpectation
	expectations       []*UserProfilerMockUpdateCaptionExpectation

	callArgs []*UserProfilerMockUpdateCaptionParams
	mutex    sync.RWMutex
}

// UserProfilerMockUpdateCaptionExpectation specifies expectation struct of the UserProfiler.UpdateCaption
type UserProfilerMockUpdateCaptionExpectation struct {
	mock    *UserProfilerMock
	params  *UserProfilerMockUpdateCaptionParams
	results *UserProfilerMockUpdateCaptionResults
	Counter uint64
}

// UserProfilerMockUpdateCaptionParams contains parameters of the UserProfiler.UpdateCaption
type UserProfilerMockUpdateCaptionParams struct {
	ctx        context.Context
	userID     int
	newCaption string
}

// UserProfilerMockUpdateCaptionResults contains results of the UserProfiler.UpdateCaption
type UserProfilerMockUpdateCaptionResults struct {
	err error
}

// Expect sets up expected params for UserProfiler.UpdateCaption
func (mmUpdateCaption *mUserProfilerMockUpdateCaption) Expect(ctx context.Context, userID int, newCaption string) *mUserProfilerMockUpdateCaption {
	if mmUpdateCaption.mock.funcUpdateCaption != nil {
		mmUpdateCaption.mock.t.Fatalf("UserProfilerMock.UpdateCaption mock is already set by Set")
	}

	if mmUpdateCaption.defaultExpectation == nil {
		mmUpdateCaption.defaultExpectation = &UserProfilerMockUpdateCaptionExpectation{}
	}

	mmUpdateCaption.defaultExpectation.params = &UserProfilerMockUpdateCaptionParams{ctx, userID, newCaption}
	for _, e := range mmUpdateCaption.expectations {
		if minimock.Equal(e.params, mmUpdateCaption.defaultExpectation.params) {
			mmUpdateCaption.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdateCaption.defaultExpectation.params)
		}
	}

	return mmUpdateCaption
}

// Inspect accepts an inspector function that has same arguments as the UserProfiler.UpdateCaption
func (mmUpdateCaption *mUserProfilerMockUpdateCaption) Inspect(f func(ctx context.Context, userID int, newCaption string)) *mUserProfilerMockUpdateCaption {
	if mmUpdateCaption.mock.inspectFuncUpdateCaption != nil {
		mmUpdateCaption.mock.t.Fatalf("Inspect function is already set for UserProfilerMock.UpdateCaption")
	}

	mmUpdateCaption.mock.inspectFuncUpdateCaption = f

	return mmUpdateCaption
}

// Return sets up results that will be returned by UserProfiler.UpdateCaption
func (mmUpdateCaption *mUserProfilerMockUpdateCaption) Return(err error) *UserProfilerMock {
	if mmUpdateCaption.mock.funcUpdateCaption != nil {
		mmUpdateCaption.mock.t.Fatalf("UserProfilerMock.UpdateCaption mock is already set by Set")
	}

	if mmUpdateCaption.defaultExpectation == nil {
		mmUpdateCaption.defaultExpectation = &UserProfilerMockUpdateCaptionExpectation{mock: mmUpdateCaption.mock}
	}
	mmUpdateCaption.defaultExpectation.results = &UserProfilerMockUpdateCaptionResults{err}
	return mmUpdateCaption.mock
}

// Set uses given function f to mock the UserProfiler.UpdateCaption method
func (mmUpdateCaption *mUserProfilerMockUpdateCaption) Set(f func(ctx context.Context, userID int, newCaption string) (err error)) *UserProfilerMock {
	if mmUpdateCaption.defaultExpectation != nil {
		mmUpdateCaption.mock.t.Fatalf("Default expectation is already set for the UserProfiler.UpdateCaption method")
	}

	if len(mmUpdateCaption.expectations) > 0 {
		mmUpdateCaption.mock.t.Fatalf("Some expectations are already set for the UserProfiler.UpdateCaption method")
	}

	mmUpdateCaption.mock.funcUpdateCaption = f
	return mmUpdateCaption.mock
}

// When sets expectation for the UserProfiler.UpdateCaption which will trigger the result defined by the following
// Then helper
func (mmUpdateCaption *mUserProfilerMockUpdateCaption) When(ctx context.Context, userID int, newCaption string) *UserProfilerMockUpdateCaptionExpectation {
	if mmUpdateCaption.mock.funcUpdateCaption != nil {
		mmUpdateCaption.mock.t.Fatalf("UserProfilerMock.UpdateCaption mock is already set by Set")
	}

	expectation := &UserProfilerMockUpdateCaptionExpectation{
		mock:   mmUpdateCaption.mock,
		params: &UserProfilerMockUpdateCaptionParams{ctx, userID, newCaption},
	}
	mmUpdateCaption.expectations = append(mmUpdateCaption.expectations, expectation)
	return expectation
}

// Then sets up UserProfiler.UpdateCaption return parameters for the expectation previously defined by the When method
func (e *UserProfilerMockUpdateCaptionExpectation) Then(err error) *UserProfilerMock {
	e.results = &UserProfilerMockUpdateCaptionResults{err}
	return e.mock
}

// UpdateCaption implements usecase.UserProfiler
func (mmUpdateCaption *UserProfilerMock) UpdateCaption(ctx context.Context, userID int, newCaption string) (err error) {
	mm_atomic.AddUint64(&mmUpdateCaption.beforeUpdateCaptionCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdateCaption.afterUpdateCaptionCounter, 1)

	if mmUpdateCaption.inspectFuncUpdateCaption != nil {
		mmUpdateCaption.inspectFuncUpdateCaption(ctx, userID, newCaption)
	}

	mm_params := &UserProfilerMockUpdateCaptionParams{ctx, userID, newCaption}

	// Record call args
	mmUpdateCaption.UpdateCaptionMock.mutex.Lock()
	mmUpdateCaption.UpdateCaptionMock.callArgs = append(mmUpdateCaption.UpdateCaptionMock.callArgs, mm_params)
	mmUpdateCaption.UpdateCaptionMock.mutex.Unlock()

	for _, e := range mmUpdateCaption.UpdateCaptionMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdateCaption.UpdateCaptionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdateCaption.UpdateCaptionMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdateCaption.UpdateCaptionMock.defaultExpectation.params
		mm_got := UserProfilerMockUpdateCaptionParams{ctx, userID, newCaption}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdateCaption.t.Errorf("UserProfilerMock.UpdateCaption got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdateCaption.UpdateCaptionMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdateCaption.t.Fatal("No results are set for the UserProfilerMock.UpdateCaption")
		}
		return (*mm_results).err
	}
	if mmUpdateCaption.funcUpdateCaption != nil {
		return mmUpdateCaption.funcUpdateCaption(ctx, userID, newCaption)
	}
	mmUpdateCaption.t.Fatalf("Unexpected call to UserProfilerMock.UpdateCaption. %v %v %v", ctx, userID, newCaption)
	return
}

// UpdateCaptionAfterCounter returns a count of finished UserProfilerMock.UpdateCaption invocations
func (mmUpdateCaption *UserProfilerMock) UpdateCaptionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateCaption.afterUpdateCaptionCounter)
}

// UpdateCaptionBeforeCounter returns a count of UserProfilerMock.UpdateCaption invocations
func (mmUpdateCaption *UserProfilerMock) UpdateCaptionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateCaption.beforeUpdateCaptionCounter)
}

// Calls returns a list of arguments used in each call to UserProfilerMock.UpdateCaption.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdateCaption *mUserProfilerMockUpdateCaption) Calls() []*UserProfilerMockUpdateCaptionParams {
	mmUpdateCaption.mutex.RLock()

	argCopy := make([]*UserProfilerMockUpdateCaptionParams, len(mmUpdateCaption.callArgs))
	copy(argCopy, mmUpdateCaption.callArgs)

	mmUpdateCaption.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateCaptionDone returns true if the count of the UpdateCaption invocations corresponds
// the number of defined expectations
func (m *UserProfilerMock) MinimockUpdateCaptionDone() bool {
	for _, e := range m.UpdateCaptionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateCaptionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCaptionCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateCaption != nil && mm_atomic.LoadUint64(&m.afterUpdateCaptionCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateCaptionInspect logs each unmet expectation
func (m *UserProfilerMock) MinimockUpdateCaptionInspect() {
	for _, e := range m.UpdateCaptionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserProfilerMock.UpdateCaption with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateCaptionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCaptionCounter) < 1 {
		if m.UpdateCaptionMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserProfilerMock.UpdateCaption")
		} else {
			m.t.Errorf("Expected call to UserProfilerMock.UpdateCaption with params: %#v", *m.UpdateCaptionMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateCaption != nil && mm_atomic.LoadUint64(&m.afterUpdateCaptionCounter) < 1 {
		m.t.Error("Expected call to UserProfilerMock.UpdateCaption")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserProfilerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeactivateInspect()

		m.MinimockLoginInspect()

		m.MinimockRegisterInspect()

		m.MinimockUpdateCaptionInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserProfilerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserProfilerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeactivateDone() &&
		m.MinimockLoginDone() &&
		m.MinimockRegisterDone() &&
		m.MinimockUpdateCaptionDone()
}
