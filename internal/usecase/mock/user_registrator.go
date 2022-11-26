package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/gojuno/minimock/v3"
)

// UserRegistratorMock implements usecase.UserRegistrator
type UserRegistratorMock struct {
	t minimock.Tester

	funcDeactivate          func(ctx context.Context, userID int) (err error)
	inspectFuncDeactivate   func(ctx context.Context, userID int)
	afterDeactivateCounter  uint64
	beforeDeactivateCounter uint64
	DeactivateMock          mUserRegistratorMockDeactivate

	funcRegister          func(ctx context.Context, name string, email string, birthDate string) (up1 *entity.User, err error)
	inspectFuncRegister   func(ctx context.Context, name string, email string, birthDate string)
	afterRegisterCounter  uint64
	beforeRegisterCounter uint64
	RegisterMock          mUserRegistratorMockRegister
}

// NewUserRegistratorMock returns a mock for usecase.UserRegistrator
func NewUserRegistratorMock(t minimock.Tester) *UserRegistratorMock {
	m := &UserRegistratorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeactivateMock = mUserRegistratorMockDeactivate{mock: m}
	m.DeactivateMock.callArgs = []*UserRegistratorMockDeactivateParams{}

	m.RegisterMock = mUserRegistratorMockRegister{mock: m}
	m.RegisterMock.callArgs = []*UserRegistratorMockRegisterParams{}

	return m
}

type mUserRegistratorMockDeactivate struct {
	mock               *UserRegistratorMock
	defaultExpectation *UserRegistratorMockDeactivateExpectation
	expectations       []*UserRegistratorMockDeactivateExpectation

	callArgs []*UserRegistratorMockDeactivateParams
	mutex    sync.RWMutex
}

// UserRegistratorMockDeactivateExpectation specifies expectation struct of the UserRegistrator.Deactivate
type UserRegistratorMockDeactivateExpectation struct {
	mock    *UserRegistratorMock
	params  *UserRegistratorMockDeactivateParams
	results *UserRegistratorMockDeactivateResults
	Counter uint64
}

// UserRegistratorMockDeactivateParams contains parameters of the UserRegistrator.Deactivate
type UserRegistratorMockDeactivateParams struct {
	ctx    context.Context
	userID int
}

// UserRegistratorMockDeactivateResults contains results of the UserRegistrator.Deactivate
type UserRegistratorMockDeactivateResults struct {
	err error
}

// Expect sets up expected params for UserRegistrator.Deactivate
func (mmDeactivate *mUserRegistratorMockDeactivate) Expect(ctx context.Context, userID int) *mUserRegistratorMockDeactivate {
	if mmDeactivate.mock.funcDeactivate != nil {
		mmDeactivate.mock.t.Fatalf("UserRegistratorMock.Deactivate mock is already set by Set")
	}

	if mmDeactivate.defaultExpectation == nil {
		mmDeactivate.defaultExpectation = &UserRegistratorMockDeactivateExpectation{}
	}

	mmDeactivate.defaultExpectation.params = &UserRegistratorMockDeactivateParams{ctx, userID}
	for _, e := range mmDeactivate.expectations {
		if minimock.Equal(e.params, mmDeactivate.defaultExpectation.params) {
			mmDeactivate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeactivate.defaultExpectation.params)
		}
	}

	return mmDeactivate
}

// Inspect accepts an inspector function that has same arguments as the UserRegistrator.Deactivate
func (mmDeactivate *mUserRegistratorMockDeactivate) Inspect(f func(ctx context.Context, userID int)) *mUserRegistratorMockDeactivate {
	if mmDeactivate.mock.inspectFuncDeactivate != nil {
		mmDeactivate.mock.t.Fatalf("Inspect function is already set for UserRegistratorMock.Deactivate")
	}

	mmDeactivate.mock.inspectFuncDeactivate = f

	return mmDeactivate
}

// Return sets up results that will be returned by UserRegistrator.Deactivate
func (mmDeactivate *mUserRegistratorMockDeactivate) Return(err error) *UserRegistratorMock {
	if mmDeactivate.mock.funcDeactivate != nil {
		mmDeactivate.mock.t.Fatalf("UserRegistratorMock.Deactivate mock is already set by Set")
	}

	if mmDeactivate.defaultExpectation == nil {
		mmDeactivate.defaultExpectation = &UserRegistratorMockDeactivateExpectation{mock: mmDeactivate.mock}
	}
	mmDeactivate.defaultExpectation.results = &UserRegistratorMockDeactivateResults{err}
	return mmDeactivate.mock
}

//Set uses given function f to mock the UserRegistrator.Deactivate method
func (mmDeactivate *mUserRegistratorMockDeactivate) Set(f func(ctx context.Context, userID int) (err error)) *UserRegistratorMock {
	if mmDeactivate.defaultExpectation != nil {
		mmDeactivate.mock.t.Fatalf("Default expectation is already set for the UserRegistrator.Deactivate method")
	}

	if len(mmDeactivate.expectations) > 0 {
		mmDeactivate.mock.t.Fatalf("Some expectations are already set for the UserRegistrator.Deactivate method")
	}

	mmDeactivate.mock.funcDeactivate = f
	return mmDeactivate.mock
}

// When sets expectation for the UserRegistrator.Deactivate which will trigger the result defined by the following
// Then helper
func (mmDeactivate *mUserRegistratorMockDeactivate) When(ctx context.Context, userID int) *UserRegistratorMockDeactivateExpectation {
	if mmDeactivate.mock.funcDeactivate != nil {
		mmDeactivate.mock.t.Fatalf("UserRegistratorMock.Deactivate mock is already set by Set")
	}

	expectation := &UserRegistratorMockDeactivateExpectation{
		mock:   mmDeactivate.mock,
		params: &UserRegistratorMockDeactivateParams{ctx, userID},
	}
	mmDeactivate.expectations = append(mmDeactivate.expectations, expectation)
	return expectation
}

// Then sets up UserRegistrator.Deactivate return parameters for the expectation previously defined by the When method
func (e *UserRegistratorMockDeactivateExpectation) Then(err error) *UserRegistratorMock {
	e.results = &UserRegistratorMockDeactivateResults{err}
	return e.mock
}

// Deactivate implements usecase.UserRegistrator
func (mmDeactivate *UserRegistratorMock) Deactivate(ctx context.Context, userID int) (err error) {
	mm_atomic.AddUint64(&mmDeactivate.beforeDeactivateCounter, 1)
	defer mm_atomic.AddUint64(&mmDeactivate.afterDeactivateCounter, 1)

	if mmDeactivate.inspectFuncDeactivate != nil {
		mmDeactivate.inspectFuncDeactivate(ctx, userID)
	}

	mm_params := &UserRegistratorMockDeactivateParams{ctx, userID}

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
		mm_got := UserRegistratorMockDeactivateParams{ctx, userID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeactivate.t.Errorf("UserRegistratorMock.Deactivate got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeactivate.DeactivateMock.defaultExpectation.results
		if mm_results == nil {
			mmDeactivate.t.Fatal("No results are set for the UserRegistratorMock.Deactivate")
		}
		return (*mm_results).err
	}
	if mmDeactivate.funcDeactivate != nil {
		return mmDeactivate.funcDeactivate(ctx, userID)
	}
	mmDeactivate.t.Fatalf("Unexpected call to UserRegistratorMock.Deactivate. %v %v", ctx, userID)
	return
}

// DeactivateAfterCounter returns a count of finished UserRegistratorMock.Deactivate invocations
func (mmDeactivate *UserRegistratorMock) DeactivateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeactivate.afterDeactivateCounter)
}

// DeactivateBeforeCounter returns a count of UserRegistratorMock.Deactivate invocations
func (mmDeactivate *UserRegistratorMock) DeactivateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeactivate.beforeDeactivateCounter)
}

// Calls returns a list of arguments used in each call to UserRegistratorMock.Deactivate.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeactivate *mUserRegistratorMockDeactivate) Calls() []*UserRegistratorMockDeactivateParams {
	mmDeactivate.mutex.RLock()

	argCopy := make([]*UserRegistratorMockDeactivateParams, len(mmDeactivate.callArgs))
	copy(argCopy, mmDeactivate.callArgs)

	mmDeactivate.mutex.RUnlock()

	return argCopy
}

// MinimockDeactivateDone returns true if the count of the Deactivate invocations corresponds
// the number of defined expectations
func (m *UserRegistratorMock) MinimockDeactivateDone() bool {
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
func (m *UserRegistratorMock) MinimockDeactivateInspect() {
	for _, e := range m.DeactivateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRegistratorMock.Deactivate with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeactivateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeactivateCounter) < 1 {
		if m.DeactivateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRegistratorMock.Deactivate")
		} else {
			m.t.Errorf("Expected call to UserRegistratorMock.Deactivate with params: %#v", *m.DeactivateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeactivate != nil && mm_atomic.LoadUint64(&m.afterDeactivateCounter) < 1 {
		m.t.Error("Expected call to UserRegistratorMock.Deactivate")
	}
}

type mUserRegistratorMockRegister struct {
	mock               *UserRegistratorMock
	defaultExpectation *UserRegistratorMockRegisterExpectation
	expectations       []*UserRegistratorMockRegisterExpectation

	callArgs []*UserRegistratorMockRegisterParams
	mutex    sync.RWMutex
}

// UserRegistratorMockRegisterExpectation specifies expectation struct of the UserRegistrator.Register
type UserRegistratorMockRegisterExpectation struct {
	mock    *UserRegistratorMock
	params  *UserRegistratorMockRegisterParams
	results *UserRegistratorMockRegisterResults
	Counter uint64
}

// UserRegistratorMockRegisterParams contains parameters of the UserRegistrator.Register
type UserRegistratorMockRegisterParams struct {
	ctx       context.Context
	name      string
	email     string
	birthDate string
}

// UserRegistratorMockRegisterResults contains results of the UserRegistrator.Register
type UserRegistratorMockRegisterResults struct {
	up1 *entity.User
	err error
}

// Expect sets up expected params for UserRegistrator.Register
func (mmRegister *mUserRegistratorMockRegister) Expect(ctx context.Context, name string, email string, birthDate string) *mUserRegistratorMockRegister {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserRegistratorMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &UserRegistratorMockRegisterExpectation{}
	}

	mmRegister.defaultExpectation.params = &UserRegistratorMockRegisterParams{ctx, name, email, birthDate}
	for _, e := range mmRegister.expectations {
		if minimock.Equal(e.params, mmRegister.defaultExpectation.params) {
			mmRegister.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRegister.defaultExpectation.params)
		}
	}

	return mmRegister
}

// Inspect accepts an inspector function that has same arguments as the UserRegistrator.Register
func (mmRegister *mUserRegistratorMockRegister) Inspect(f func(ctx context.Context, name string, email string, birthDate string)) *mUserRegistratorMockRegister {
	if mmRegister.mock.inspectFuncRegister != nil {
		mmRegister.mock.t.Fatalf("Inspect function is already set for UserRegistratorMock.Register")
	}

	mmRegister.mock.inspectFuncRegister = f

	return mmRegister
}

// Return sets up results that will be returned by UserRegistrator.Register
func (mmRegister *mUserRegistratorMockRegister) Return(up1 *entity.User, err error) *UserRegistratorMock {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserRegistratorMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &UserRegistratorMockRegisterExpectation{mock: mmRegister.mock}
	}
	mmRegister.defaultExpectation.results = &UserRegistratorMockRegisterResults{up1, err}
	return mmRegister.mock
}

//Set uses given function f to mock the UserRegistrator.Register method
func (mmRegister *mUserRegistratorMockRegister) Set(f func(ctx context.Context, name string, email string, birthDate string) (up1 *entity.User, err error)) *UserRegistratorMock {
	if mmRegister.defaultExpectation != nil {
		mmRegister.mock.t.Fatalf("Default expectation is already set for the UserRegistrator.Register method")
	}

	if len(mmRegister.expectations) > 0 {
		mmRegister.mock.t.Fatalf("Some expectations are already set for the UserRegistrator.Register method")
	}

	mmRegister.mock.funcRegister = f
	return mmRegister.mock
}

// When sets expectation for the UserRegistrator.Register which will trigger the result defined by the following
// Then helper
func (mmRegister *mUserRegistratorMockRegister) When(ctx context.Context, name string, email string, birthDate string) *UserRegistratorMockRegisterExpectation {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("UserRegistratorMock.Register mock is already set by Set")
	}

	expectation := &UserRegistratorMockRegisterExpectation{
		mock:   mmRegister.mock,
		params: &UserRegistratorMockRegisterParams{ctx, name, email, birthDate},
	}
	mmRegister.expectations = append(mmRegister.expectations, expectation)
	return expectation
}

// Then sets up UserRegistrator.Register return parameters for the expectation previously defined by the When method
func (e *UserRegistratorMockRegisterExpectation) Then(up1 *entity.User, err error) *UserRegistratorMock {
	e.results = &UserRegistratorMockRegisterResults{up1, err}
	return e.mock
}

// Register implements usecase.UserRegistrator
func (mmRegister *UserRegistratorMock) Register(ctx context.Context, name string, email string, birthDate string) (up1 *entity.User, err error) {
	mm_atomic.AddUint64(&mmRegister.beforeRegisterCounter, 1)
	defer mm_atomic.AddUint64(&mmRegister.afterRegisterCounter, 1)

	if mmRegister.inspectFuncRegister != nil {
		mmRegister.inspectFuncRegister(ctx, name, email, birthDate)
	}

	mm_params := &UserRegistratorMockRegisterParams{ctx, name, email, birthDate}

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
		mm_got := UserRegistratorMockRegisterParams{ctx, name, email, birthDate}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRegister.t.Errorf("UserRegistratorMock.Register got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRegister.RegisterMock.defaultExpectation.results
		if mm_results == nil {
			mmRegister.t.Fatal("No results are set for the UserRegistratorMock.Register")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmRegister.funcRegister != nil {
		return mmRegister.funcRegister(ctx, name, email, birthDate)
	}
	mmRegister.t.Fatalf("Unexpected call to UserRegistratorMock.Register. %v %v %v %v", ctx, name, email, birthDate)
	return
}

// RegisterAfterCounter returns a count of finished UserRegistratorMock.Register invocations
func (mmRegister *UserRegistratorMock) RegisterAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegister.afterRegisterCounter)
}

// RegisterBeforeCounter returns a count of UserRegistratorMock.Register invocations
func (mmRegister *UserRegistratorMock) RegisterBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegister.beforeRegisterCounter)
}

// Calls returns a list of arguments used in each call to UserRegistratorMock.Register.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRegister *mUserRegistratorMockRegister) Calls() []*UserRegistratorMockRegisterParams {
	mmRegister.mutex.RLock()

	argCopy := make([]*UserRegistratorMockRegisterParams, len(mmRegister.callArgs))
	copy(argCopy, mmRegister.callArgs)

	mmRegister.mutex.RUnlock()

	return argCopy
}

// MinimockRegisterDone returns true if the count of the Register invocations corresponds
// the number of defined expectations
func (m *UserRegistratorMock) MinimockRegisterDone() bool {
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
func (m *UserRegistratorMock) MinimockRegisterInspect() {
	for _, e := range m.RegisterMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRegistratorMock.Register with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RegisterMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		if m.RegisterMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRegistratorMock.Register")
		} else {
			m.t.Errorf("Expected call to UserRegistratorMock.Register with params: %#v", *m.RegisterMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegister != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		m.t.Error("Expected call to UserRegistratorMock.Register")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserRegistratorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeactivateInspect()

		m.MinimockRegisterInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserRegistratorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *UserRegistratorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeactivateDone() &&
		m.MinimockRegisterDone()
}