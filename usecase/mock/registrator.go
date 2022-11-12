package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"
	"github.com/demimurg/twitter/entity"

	"github.com/gojuno/minimock/v3"
)

// RegistratorMock implements twitter.Registrator
type RegistratorMock struct {
	t minimock.Tester

	funcDeactivateUser          func(userID string) (err error)
	inspectFuncDeactivateUser   func(userID string)
	afterDeactivateUserCounter  uint64
	beforeDeactivateUserCounter uint64
	DeactivateUserMock          mRegistratorMockDeactivateUser

	funcRegisterUser          func(name string, email string, birthDate *time.Time) (up1 *entity.User, err error)
	inspectFuncRegisterUser   func(name string, email string, birthDate *time.Time)
	afterRegisterUserCounter  uint64
	beforeRegisterUserCounter uint64
	RegisterUserMock          mRegistratorMockRegisterUser
}

// NewRegistratorMock returns a mock for twitter.Registrator
func NewRegistratorMock(t minimock.Tester) *RegistratorMock {
	m := &RegistratorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeactivateUserMock = mRegistratorMockDeactivateUser{mock: m}
	m.DeactivateUserMock.callArgs = []*RegistratorMockDeactivateUserParams{}

	m.RegisterUserMock = mRegistratorMockRegisterUser{mock: m}
	m.RegisterUserMock.callArgs = []*RegistratorMockRegisterUserParams{}

	return m
}

type mRegistratorMockDeactivateUser struct {
	mock               *RegistratorMock
	defaultExpectation *RegistratorMockDeactivateUserExpectation
	expectations       []*RegistratorMockDeactivateUserExpectation

	callArgs []*RegistratorMockDeactivateUserParams
	mutex    sync.RWMutex
}

// RegistratorMockDeactivateUserExpectation specifies expectation struct of the Registrator.DeactivateUser
type RegistratorMockDeactivateUserExpectation struct {
	mock    *RegistratorMock
	params  *RegistratorMockDeactivateUserParams
	results *RegistratorMockDeactivateUserResults
	Counter uint64
}

// RegistratorMockDeactivateUserParams contains parameters of the Registrator.DeactivateUser
type RegistratorMockDeactivateUserParams struct {
	userID string
}

// RegistratorMockDeactivateUserResults contains results of the Registrator.DeactivateUser
type RegistratorMockDeactivateUserResults struct {
	err error
}

// Expect sets up expected params for Registrator.DeactivateUser
func (mmDeactivateUser *mRegistratorMockDeactivateUser) Expect(userID string) *mRegistratorMockDeactivateUser {
	if mmDeactivateUser.mock.funcDeactivateUser != nil {
		mmDeactivateUser.mock.t.Fatalf("RegistratorMock.DeactivateUser mock is already set by Set")
	}

	if mmDeactivateUser.defaultExpectation == nil {
		mmDeactivateUser.defaultExpectation = &RegistratorMockDeactivateUserExpectation{}
	}

	mmDeactivateUser.defaultExpectation.params = &RegistratorMockDeactivateUserParams{userID}
	for _, e := range mmDeactivateUser.expectations {
		if minimock.Equal(e.params, mmDeactivateUser.defaultExpectation.params) {
			mmDeactivateUser.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeactivateUser.defaultExpectation.params)
		}
	}

	return mmDeactivateUser
}

// Inspect accepts an inspector function that has same arguments as the Registrator.DeactivateUser
func (mmDeactivateUser *mRegistratorMockDeactivateUser) Inspect(f func(userID string)) *mRegistratorMockDeactivateUser {
	if mmDeactivateUser.mock.inspectFuncDeactivateUser != nil {
		mmDeactivateUser.mock.t.Fatalf("Inspect function is already set for RegistratorMock.DeactivateUser")
	}

	mmDeactivateUser.mock.inspectFuncDeactivateUser = f

	return mmDeactivateUser
}

// Return sets up results that will be returned by Registrator.DeactivateUser
func (mmDeactivateUser *mRegistratorMockDeactivateUser) Return(err error) *RegistratorMock {
	if mmDeactivateUser.mock.funcDeactivateUser != nil {
		mmDeactivateUser.mock.t.Fatalf("RegistratorMock.DeactivateUser mock is already set by Set")
	}

	if mmDeactivateUser.defaultExpectation == nil {
		mmDeactivateUser.defaultExpectation = &RegistratorMockDeactivateUserExpectation{mock: mmDeactivateUser.mock}
	}
	mmDeactivateUser.defaultExpectation.results = &RegistratorMockDeactivateUserResults{err}
	return mmDeactivateUser.mock
}

//Set uses given function f to mock the Registrator.DeactivateUser method
func (mmDeactivateUser *mRegistratorMockDeactivateUser) Set(f func(userID string) (err error)) *RegistratorMock {
	if mmDeactivateUser.defaultExpectation != nil {
		mmDeactivateUser.mock.t.Fatalf("Default expectation is already set for the Registrator.DeactivateUser method")
	}

	if len(mmDeactivateUser.expectations) > 0 {
		mmDeactivateUser.mock.t.Fatalf("Some expectations are already set for the Registrator.DeactivateUser method")
	}

	mmDeactivateUser.mock.funcDeactivateUser = f
	return mmDeactivateUser.mock
}

// When sets expectation for the Registrator.DeactivateUser which will trigger the result defined by the following
// Then helper
func (mmDeactivateUser *mRegistratorMockDeactivateUser) When(userID string) *RegistratorMockDeactivateUserExpectation {
	if mmDeactivateUser.mock.funcDeactivateUser != nil {
		mmDeactivateUser.mock.t.Fatalf("RegistratorMock.DeactivateUser mock is already set by Set")
	}

	expectation := &RegistratorMockDeactivateUserExpectation{
		mock:   mmDeactivateUser.mock,
		params: &RegistratorMockDeactivateUserParams{userID},
	}
	mmDeactivateUser.expectations = append(mmDeactivateUser.expectations, expectation)
	return expectation
}

// Then sets up Registrator.DeactivateUser return parameters for the expectation previously defined by the When method
func (e *RegistratorMockDeactivateUserExpectation) Then(err error) *RegistratorMock {
	e.results = &RegistratorMockDeactivateUserResults{err}
	return e.mock
}

// DeactivateUser implements twitter.Registrator
func (mmDeactivateUser *RegistratorMock) DeactivateUser(userID string) (err error) {
	mm_atomic.AddUint64(&mmDeactivateUser.beforeDeactivateUserCounter, 1)
	defer mm_atomic.AddUint64(&mmDeactivateUser.afterDeactivateUserCounter, 1)

	if mmDeactivateUser.inspectFuncDeactivateUser != nil {
		mmDeactivateUser.inspectFuncDeactivateUser(userID)
	}

	mm_params := &RegistratorMockDeactivateUserParams{userID}

	// Record call args
	mmDeactivateUser.DeactivateUserMock.mutex.Lock()
	mmDeactivateUser.DeactivateUserMock.callArgs = append(mmDeactivateUser.DeactivateUserMock.callArgs, mm_params)
	mmDeactivateUser.DeactivateUserMock.mutex.Unlock()

	for _, e := range mmDeactivateUser.DeactivateUserMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeactivateUser.DeactivateUserMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeactivateUser.DeactivateUserMock.defaultExpectation.Counter, 1)
		mm_want := mmDeactivateUser.DeactivateUserMock.defaultExpectation.params
		mm_got := RegistratorMockDeactivateUserParams{userID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeactivateUser.t.Errorf("RegistratorMock.DeactivateUser got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeactivateUser.DeactivateUserMock.defaultExpectation.results
		if mm_results == nil {
			mmDeactivateUser.t.Fatal("No results are set for the RegistratorMock.DeactivateUser")
		}
		return (*mm_results).err
	}
	if mmDeactivateUser.funcDeactivateUser != nil {
		return mmDeactivateUser.funcDeactivateUser(userID)
	}
	mmDeactivateUser.t.Fatalf("Unexpected call to RegistratorMock.DeactivateUser. %v", userID)
	return
}

// DeactivateUserAfterCounter returns a count of finished RegistratorMock.DeactivateUser invocations
func (mmDeactivateUser *RegistratorMock) DeactivateUserAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeactivateUser.afterDeactivateUserCounter)
}

// DeactivateUserBeforeCounter returns a count of RegistratorMock.DeactivateUser invocations
func (mmDeactivateUser *RegistratorMock) DeactivateUserBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeactivateUser.beforeDeactivateUserCounter)
}

// Calls returns a list of arguments used in each call to RegistratorMock.DeactivateUser.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeactivateUser *mRegistratorMockDeactivateUser) Calls() []*RegistratorMockDeactivateUserParams {
	mmDeactivateUser.mutex.RLock()

	argCopy := make([]*RegistratorMockDeactivateUserParams, len(mmDeactivateUser.callArgs))
	copy(argCopy, mmDeactivateUser.callArgs)

	mmDeactivateUser.mutex.RUnlock()

	return argCopy
}

// MinimockDeactivateUserDone returns true if the count of the DeactivateUser invocations corresponds
// the number of defined expectations
func (m *RegistratorMock) MinimockDeactivateUserDone() bool {
	for _, e := range m.DeactivateUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeactivateUserMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeactivateUserCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeactivateUser != nil && mm_atomic.LoadUint64(&m.afterDeactivateUserCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeactivateUserInspect logs each unmet expectation
func (m *RegistratorMock) MinimockDeactivateUserInspect() {
	for _, e := range m.DeactivateUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RegistratorMock.DeactivateUser with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeactivateUserMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeactivateUserCounter) < 1 {
		if m.DeactivateUserMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RegistratorMock.DeactivateUser")
		} else {
			m.t.Errorf("Expected call to RegistratorMock.DeactivateUser with params: %#v", *m.DeactivateUserMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeactivateUser != nil && mm_atomic.LoadUint64(&m.afterDeactivateUserCounter) < 1 {
		m.t.Error("Expected call to RegistratorMock.DeactivateUser")
	}
}

type mRegistratorMockRegisterUser struct {
	mock               *RegistratorMock
	defaultExpectation *RegistratorMockRegisterUserExpectation
	expectations       []*RegistratorMockRegisterUserExpectation

	callArgs []*RegistratorMockRegisterUserParams
	mutex    sync.RWMutex
}

// RegistratorMockRegisterUserExpectation specifies expectation struct of the Registrator.RegisterUser
type RegistratorMockRegisterUserExpectation struct {
	mock    *RegistratorMock
	params  *RegistratorMockRegisterUserParams
	results *RegistratorMockRegisterUserResults
	Counter uint64
}

// RegistratorMockRegisterUserParams contains parameters of the Registrator.RegisterUser
type RegistratorMockRegisterUserParams struct {
	name      string
	email     string
	birthDate *time.Time
}

// RegistratorMockRegisterUserResults contains results of the Registrator.RegisterUser
type RegistratorMockRegisterUserResults struct {
	up1 *entity.User
	err error
}

// Expect sets up expected params for Registrator.RegisterUser
func (mmRegisterUser *mRegistratorMockRegisterUser) Expect(name string, email string, birthDate *time.Time) *mRegistratorMockRegisterUser {
	if mmRegisterUser.mock.funcRegisterUser != nil {
		mmRegisterUser.mock.t.Fatalf("RegistratorMock.RegisterUser mock is already set by Set")
	}

	if mmRegisterUser.defaultExpectation == nil {
		mmRegisterUser.defaultExpectation = &RegistratorMockRegisterUserExpectation{}
	}

	mmRegisterUser.defaultExpectation.params = &RegistratorMockRegisterUserParams{name, email, birthDate}
	for _, e := range mmRegisterUser.expectations {
		if minimock.Equal(e.params, mmRegisterUser.defaultExpectation.params) {
			mmRegisterUser.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRegisterUser.defaultExpectation.params)
		}
	}

	return mmRegisterUser
}

// Inspect accepts an inspector function that has same arguments as the Registrator.RegisterUser
func (mmRegisterUser *mRegistratorMockRegisterUser) Inspect(f func(name string, email string, birthDate *time.Time)) *mRegistratorMockRegisterUser {
	if mmRegisterUser.mock.inspectFuncRegisterUser != nil {
		mmRegisterUser.mock.t.Fatalf("Inspect function is already set for RegistratorMock.RegisterUser")
	}

	mmRegisterUser.mock.inspectFuncRegisterUser = f

	return mmRegisterUser
}

// Return sets up results that will be returned by Registrator.RegisterUser
func (mmRegisterUser *mRegistratorMockRegisterUser) Return(up1 *entity.User, err error) *RegistratorMock {
	if mmRegisterUser.mock.funcRegisterUser != nil {
		mmRegisterUser.mock.t.Fatalf("RegistratorMock.RegisterUser mock is already set by Set")
	}

	if mmRegisterUser.defaultExpectation == nil {
		mmRegisterUser.defaultExpectation = &RegistratorMockRegisterUserExpectation{mock: mmRegisterUser.mock}
	}
	mmRegisterUser.defaultExpectation.results = &RegistratorMockRegisterUserResults{up1, err}
	return mmRegisterUser.mock
}

//Set uses given function f to mock the Registrator.RegisterUser method
func (mmRegisterUser *mRegistratorMockRegisterUser) Set(f func(name string, email string, birthDate *time.Time) (up1 *entity.User, err error)) *RegistratorMock {
	if mmRegisterUser.defaultExpectation != nil {
		mmRegisterUser.mock.t.Fatalf("Default expectation is already set for the Registrator.RegisterUser method")
	}

	if len(mmRegisterUser.expectations) > 0 {
		mmRegisterUser.mock.t.Fatalf("Some expectations are already set for the Registrator.RegisterUser method")
	}

	mmRegisterUser.mock.funcRegisterUser = f
	return mmRegisterUser.mock
}

// When sets expectation for the Registrator.RegisterUser which will trigger the result defined by the following
// Then helper
func (mmRegisterUser *mRegistratorMockRegisterUser) When(name string, email string, birthDate *time.Time) *RegistratorMockRegisterUserExpectation {
	if mmRegisterUser.mock.funcRegisterUser != nil {
		mmRegisterUser.mock.t.Fatalf("RegistratorMock.RegisterUser mock is already set by Set")
	}

	expectation := &RegistratorMockRegisterUserExpectation{
		mock:   mmRegisterUser.mock,
		params: &RegistratorMockRegisterUserParams{name, email, birthDate},
	}
	mmRegisterUser.expectations = append(mmRegisterUser.expectations, expectation)
	return expectation
}

// Then sets up Registrator.RegisterUser return parameters for the expectation previously defined by the When method
func (e *RegistratorMockRegisterUserExpectation) Then(up1 *entity.User, err error) *RegistratorMock {
	e.results = &RegistratorMockRegisterUserResults{up1, err}
	return e.mock
}

// RegisterUser implements twitter.Registrator
func (mmRegisterUser *RegistratorMock) RegisterUser(name string, email string, birthDate *time.Time) (up1 *entity.User, err error) {
	mm_atomic.AddUint64(&mmRegisterUser.beforeRegisterUserCounter, 1)
	defer mm_atomic.AddUint64(&mmRegisterUser.afterRegisterUserCounter, 1)

	if mmRegisterUser.inspectFuncRegisterUser != nil {
		mmRegisterUser.inspectFuncRegisterUser(name, email, birthDate)
	}

	mm_params := &RegistratorMockRegisterUserParams{name, email, birthDate}

	// Record call args
	mmRegisterUser.RegisterUserMock.mutex.Lock()
	mmRegisterUser.RegisterUserMock.callArgs = append(mmRegisterUser.RegisterUserMock.callArgs, mm_params)
	mmRegisterUser.RegisterUserMock.mutex.Unlock()

	for _, e := range mmRegisterUser.RegisterUserMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmRegisterUser.RegisterUserMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRegisterUser.RegisterUserMock.defaultExpectation.Counter, 1)
		mm_want := mmRegisterUser.RegisterUserMock.defaultExpectation.params
		mm_got := RegistratorMockRegisterUserParams{name, email, birthDate}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRegisterUser.t.Errorf("RegistratorMock.RegisterUser got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRegisterUser.RegisterUserMock.defaultExpectation.results
		if mm_results == nil {
			mmRegisterUser.t.Fatal("No results are set for the RegistratorMock.RegisterUser")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmRegisterUser.funcRegisterUser != nil {
		return mmRegisterUser.funcRegisterUser(name, email, birthDate)
	}
	mmRegisterUser.t.Fatalf("Unexpected call to RegistratorMock.RegisterUser. %v %v %v", name, email, birthDate)
	return
}

// RegisterUserAfterCounter returns a count of finished RegistratorMock.RegisterUser invocations
func (mmRegisterUser *RegistratorMock) RegisterUserAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegisterUser.afterRegisterUserCounter)
}

// RegisterUserBeforeCounter returns a count of RegistratorMock.RegisterUser invocations
func (mmRegisterUser *RegistratorMock) RegisterUserBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegisterUser.beforeRegisterUserCounter)
}

// Calls returns a list of arguments used in each call to RegistratorMock.RegisterUser.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRegisterUser *mRegistratorMockRegisterUser) Calls() []*RegistratorMockRegisterUserParams {
	mmRegisterUser.mutex.RLock()

	argCopy := make([]*RegistratorMockRegisterUserParams, len(mmRegisterUser.callArgs))
	copy(argCopy, mmRegisterUser.callArgs)

	mmRegisterUser.mutex.RUnlock()

	return argCopy
}

// MinimockRegisterUserDone returns true if the count of the RegisterUser invocations corresponds
// the number of defined expectations
func (m *RegistratorMock) MinimockRegisterUserDone() bool {
	for _, e := range m.RegisterUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RegisterUserMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRegisterUserCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegisterUser != nil && mm_atomic.LoadUint64(&m.afterRegisterUserCounter) < 1 {
		return false
	}
	return true
}

// MinimockRegisterUserInspect logs each unmet expectation
func (m *RegistratorMock) MinimockRegisterUserInspect() {
	for _, e := range m.RegisterUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RegistratorMock.RegisterUser with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RegisterUserMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRegisterUserCounter) < 1 {
		if m.RegisterUserMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RegistratorMock.RegisterUser")
		} else {
			m.t.Errorf("Expected call to RegistratorMock.RegisterUser with params: %#v", *m.RegisterUserMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegisterUser != nil && mm_atomic.LoadUint64(&m.afterRegisterUserCounter) < 1 {
		m.t.Error("Expected call to RegistratorMock.RegisterUser")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RegistratorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeactivateUserInspect()

		m.MinimockRegisterUserInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RegistratorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RegistratorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeactivateUserDone() &&
		m.MinimockRegisterUserDone()
}
