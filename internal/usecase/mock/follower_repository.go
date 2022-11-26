package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// FollowerRepositoryMock implements usecase.FollowerRepository
type FollowerRepositoryMock struct {
	t minimock.Tester

	funcAdd          func(ctx context.Context, followerID int, toUserID int) (err error)
	inspectFuncAdd   func(ctx context.Context, followerID int, toUserID int)
	afterAddCounter  uint64
	beforeAddCounter uint64
	AddMock          mFollowerRepositoryMockAdd

	funcGetFollowers          func(ctx context.Context, userID int, topN int) (ia1 []int, err error)
	inspectFuncGetFollowers   func(ctx context.Context, userID int, topN int)
	afterGetFollowersCounter  uint64
	beforeGetFollowersCounter uint64
	GetFollowersMock          mFollowerRepositoryMockGetFollowers

	funcGetFollowing          func(ctx context.Context, userID int, topN int) (ia1 []int, err error)
	inspectFuncGetFollowing   func(ctx context.Context, userID int, topN int)
	afterGetFollowingCounter  uint64
	beforeGetFollowingCounter uint64
	GetFollowingMock          mFollowerRepositoryMockGetFollowing

	funcRemove          func(ctx context.Context, followerID int, fromUserID int) (err error)
	inspectFuncRemove   func(ctx context.Context, followerID int, fromUserID int)
	afterRemoveCounter  uint64
	beforeRemoveCounter uint64
	RemoveMock          mFollowerRepositoryMockRemove
}

// NewFollowerRepositoryMock returns a mock for usecase.FollowerRepository
func NewFollowerRepositoryMock(t minimock.Tester) *FollowerRepositoryMock {
	m := &FollowerRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddMock = mFollowerRepositoryMockAdd{mock: m}
	m.AddMock.callArgs = []*FollowerRepositoryMockAddParams{}

	m.GetFollowersMock = mFollowerRepositoryMockGetFollowers{mock: m}
	m.GetFollowersMock.callArgs = []*FollowerRepositoryMockGetFollowersParams{}

	m.GetFollowingMock = mFollowerRepositoryMockGetFollowing{mock: m}
	m.GetFollowingMock.callArgs = []*FollowerRepositoryMockGetFollowingParams{}

	m.RemoveMock = mFollowerRepositoryMockRemove{mock: m}
	m.RemoveMock.callArgs = []*FollowerRepositoryMockRemoveParams{}

	return m
}

type mFollowerRepositoryMockAdd struct {
	mock               *FollowerRepositoryMock
	defaultExpectation *FollowerRepositoryMockAddExpectation
	expectations       []*FollowerRepositoryMockAddExpectation

	callArgs []*FollowerRepositoryMockAddParams
	mutex    sync.RWMutex
}

// FollowerRepositoryMockAddExpectation specifies expectation struct of the FollowerRepository.Add
type FollowerRepositoryMockAddExpectation struct {
	mock    *FollowerRepositoryMock
	params  *FollowerRepositoryMockAddParams
	results *FollowerRepositoryMockAddResults
	Counter uint64
}

// FollowerRepositoryMockAddParams contains parameters of the FollowerRepository.Add
type FollowerRepositoryMockAddParams struct {
	ctx        context.Context
	followerID int
	toUserID   int
}

// FollowerRepositoryMockAddResults contains results of the FollowerRepository.Add
type FollowerRepositoryMockAddResults struct {
	err error
}

// Expect sets up expected params for FollowerRepository.Add
func (mmAdd *mFollowerRepositoryMockAdd) Expect(ctx context.Context, followerID int, toUserID int) *mFollowerRepositoryMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("FollowerRepositoryMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &FollowerRepositoryMockAddExpectation{}
	}

	mmAdd.defaultExpectation.params = &FollowerRepositoryMockAddParams{ctx, followerID, toUserID}
	for _, e := range mmAdd.expectations {
		if minimock.Equal(e.params, mmAdd.defaultExpectation.params) {
			mmAdd.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAdd.defaultExpectation.params)
		}
	}

	return mmAdd
}

// Inspect accepts an inspector function that has same arguments as the FollowerRepository.Add
func (mmAdd *mFollowerRepositoryMockAdd) Inspect(f func(ctx context.Context, followerID int, toUserID int)) *mFollowerRepositoryMockAdd {
	if mmAdd.mock.inspectFuncAdd != nil {
		mmAdd.mock.t.Fatalf("Inspect function is already set for FollowerRepositoryMock.Add")
	}

	mmAdd.mock.inspectFuncAdd = f

	return mmAdd
}

// Return sets up results that will be returned by FollowerRepository.Add
func (mmAdd *mFollowerRepositoryMockAdd) Return(err error) *FollowerRepositoryMock {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("FollowerRepositoryMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &FollowerRepositoryMockAddExpectation{mock: mmAdd.mock}
	}
	mmAdd.defaultExpectation.results = &FollowerRepositoryMockAddResults{err}
	return mmAdd.mock
}

//Set uses given function f to mock the FollowerRepository.Add method
func (mmAdd *mFollowerRepositoryMockAdd) Set(f func(ctx context.Context, followerID int, toUserID int) (err error)) *FollowerRepositoryMock {
	if mmAdd.defaultExpectation != nil {
		mmAdd.mock.t.Fatalf("Default expectation is already set for the FollowerRepository.Add method")
	}

	if len(mmAdd.expectations) > 0 {
		mmAdd.mock.t.Fatalf("Some expectations are already set for the FollowerRepository.Add method")
	}

	mmAdd.mock.funcAdd = f
	return mmAdd.mock
}

// When sets expectation for the FollowerRepository.Add which will trigger the result defined by the following
// Then helper
func (mmAdd *mFollowerRepositoryMockAdd) When(ctx context.Context, followerID int, toUserID int) *FollowerRepositoryMockAddExpectation {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("FollowerRepositoryMock.Add mock is already set by Set")
	}

	expectation := &FollowerRepositoryMockAddExpectation{
		mock:   mmAdd.mock,
		params: &FollowerRepositoryMockAddParams{ctx, followerID, toUserID},
	}
	mmAdd.expectations = append(mmAdd.expectations, expectation)
	return expectation
}

// Then sets up FollowerRepository.Add return parameters for the expectation previously defined by the When method
func (e *FollowerRepositoryMockAddExpectation) Then(err error) *FollowerRepositoryMock {
	e.results = &FollowerRepositoryMockAddResults{err}
	return e.mock
}

// Add implements usecase.FollowerRepository
func (mmAdd *FollowerRepositoryMock) Add(ctx context.Context, followerID int, toUserID int) (err error) {
	mm_atomic.AddUint64(&mmAdd.beforeAddCounter, 1)
	defer mm_atomic.AddUint64(&mmAdd.afterAddCounter, 1)

	if mmAdd.inspectFuncAdd != nil {
		mmAdd.inspectFuncAdd(ctx, followerID, toUserID)
	}

	mm_params := &FollowerRepositoryMockAddParams{ctx, followerID, toUserID}

	// Record call args
	mmAdd.AddMock.mutex.Lock()
	mmAdd.AddMock.callArgs = append(mmAdd.AddMock.callArgs, mm_params)
	mmAdd.AddMock.mutex.Unlock()

	for _, e := range mmAdd.AddMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAdd.AddMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAdd.AddMock.defaultExpectation.Counter, 1)
		mm_want := mmAdd.AddMock.defaultExpectation.params
		mm_got := FollowerRepositoryMockAddParams{ctx, followerID, toUserID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAdd.t.Errorf("FollowerRepositoryMock.Add got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAdd.AddMock.defaultExpectation.results
		if mm_results == nil {
			mmAdd.t.Fatal("No results are set for the FollowerRepositoryMock.Add")
		}
		return (*mm_results).err
	}
	if mmAdd.funcAdd != nil {
		return mmAdd.funcAdd(ctx, followerID, toUserID)
	}
	mmAdd.t.Fatalf("Unexpected call to FollowerRepositoryMock.Add. %v %v %v", ctx, followerID, toUserID)
	return
}

// AddAfterCounter returns a count of finished FollowerRepositoryMock.Add invocations
func (mmAdd *FollowerRepositoryMock) AddAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.afterAddCounter)
}

// AddBeforeCounter returns a count of FollowerRepositoryMock.Add invocations
func (mmAdd *FollowerRepositoryMock) AddBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.beforeAddCounter)
}

// Calls returns a list of arguments used in each call to FollowerRepositoryMock.Add.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAdd *mFollowerRepositoryMockAdd) Calls() []*FollowerRepositoryMockAddParams {
	mmAdd.mutex.RLock()

	argCopy := make([]*FollowerRepositoryMockAddParams, len(mmAdd.callArgs))
	copy(argCopy, mmAdd.callArgs)

	mmAdd.mutex.RUnlock()

	return argCopy
}

// MinimockAddDone returns true if the count of the Add invocations corresponds
// the number of defined expectations
func (m *FollowerRepositoryMock) MinimockAddDone() bool {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddInspect logs each unmet expectation
func (m *FollowerRepositoryMock) MinimockAddInspect() {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FollowerRepositoryMock.Add with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		if m.AddMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FollowerRepositoryMock.Add")
		} else {
			m.t.Errorf("Expected call to FollowerRepositoryMock.Add with params: %#v", *m.AddMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		m.t.Error("Expected call to FollowerRepositoryMock.Add")
	}
}

type mFollowerRepositoryMockGetFollowers struct {
	mock               *FollowerRepositoryMock
	defaultExpectation *FollowerRepositoryMockGetFollowersExpectation
	expectations       []*FollowerRepositoryMockGetFollowersExpectation

	callArgs []*FollowerRepositoryMockGetFollowersParams
	mutex    sync.RWMutex
}

// FollowerRepositoryMockGetFollowersExpectation specifies expectation struct of the FollowerRepository.GetFollowers
type FollowerRepositoryMockGetFollowersExpectation struct {
	mock    *FollowerRepositoryMock
	params  *FollowerRepositoryMockGetFollowersParams
	results *FollowerRepositoryMockGetFollowersResults
	Counter uint64
}

// FollowerRepositoryMockGetFollowersParams contains parameters of the FollowerRepository.GetFollowers
type FollowerRepositoryMockGetFollowersParams struct {
	ctx    context.Context
	userID int
	topN   int
}

// FollowerRepositoryMockGetFollowersResults contains results of the FollowerRepository.GetFollowers
type FollowerRepositoryMockGetFollowersResults struct {
	ia1 []int
	err error
}

// Expect sets up expected params for FollowerRepository.GetFollowers
func (mmGetFollowers *mFollowerRepositoryMockGetFollowers) Expect(ctx context.Context, userID int, topN int) *mFollowerRepositoryMockGetFollowers {
	if mmGetFollowers.mock.funcGetFollowers != nil {
		mmGetFollowers.mock.t.Fatalf("FollowerRepositoryMock.GetFollowers mock is already set by Set")
	}

	if mmGetFollowers.defaultExpectation == nil {
		mmGetFollowers.defaultExpectation = &FollowerRepositoryMockGetFollowersExpectation{}
	}

	mmGetFollowers.defaultExpectation.params = &FollowerRepositoryMockGetFollowersParams{ctx, userID, topN}
	for _, e := range mmGetFollowers.expectations {
		if minimock.Equal(e.params, mmGetFollowers.defaultExpectation.params) {
			mmGetFollowers.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetFollowers.defaultExpectation.params)
		}
	}

	return mmGetFollowers
}

// Inspect accepts an inspector function that has same arguments as the FollowerRepository.GetFollowers
func (mmGetFollowers *mFollowerRepositoryMockGetFollowers) Inspect(f func(ctx context.Context, userID int, topN int)) *mFollowerRepositoryMockGetFollowers {
	if mmGetFollowers.mock.inspectFuncGetFollowers != nil {
		mmGetFollowers.mock.t.Fatalf("Inspect function is already set for FollowerRepositoryMock.GetFollowers")
	}

	mmGetFollowers.mock.inspectFuncGetFollowers = f

	return mmGetFollowers
}

// Return sets up results that will be returned by FollowerRepository.GetFollowers
func (mmGetFollowers *mFollowerRepositoryMockGetFollowers) Return(ia1 []int, err error) *FollowerRepositoryMock {
	if mmGetFollowers.mock.funcGetFollowers != nil {
		mmGetFollowers.mock.t.Fatalf("FollowerRepositoryMock.GetFollowers mock is already set by Set")
	}

	if mmGetFollowers.defaultExpectation == nil {
		mmGetFollowers.defaultExpectation = &FollowerRepositoryMockGetFollowersExpectation{mock: mmGetFollowers.mock}
	}
	mmGetFollowers.defaultExpectation.results = &FollowerRepositoryMockGetFollowersResults{ia1, err}
	return mmGetFollowers.mock
}

//Set uses given function f to mock the FollowerRepository.GetFollowers method
func (mmGetFollowers *mFollowerRepositoryMockGetFollowers) Set(f func(ctx context.Context, userID int, topN int) (ia1 []int, err error)) *FollowerRepositoryMock {
	if mmGetFollowers.defaultExpectation != nil {
		mmGetFollowers.mock.t.Fatalf("Default expectation is already set for the FollowerRepository.GetFollowers method")
	}

	if len(mmGetFollowers.expectations) > 0 {
		mmGetFollowers.mock.t.Fatalf("Some expectations are already set for the FollowerRepository.GetFollowers method")
	}

	mmGetFollowers.mock.funcGetFollowers = f
	return mmGetFollowers.mock
}

// When sets expectation for the FollowerRepository.GetFollowers which will trigger the result defined by the following
// Then helper
func (mmGetFollowers *mFollowerRepositoryMockGetFollowers) When(ctx context.Context, userID int, topN int) *FollowerRepositoryMockGetFollowersExpectation {
	if mmGetFollowers.mock.funcGetFollowers != nil {
		mmGetFollowers.mock.t.Fatalf("FollowerRepositoryMock.GetFollowers mock is already set by Set")
	}

	expectation := &FollowerRepositoryMockGetFollowersExpectation{
		mock:   mmGetFollowers.mock,
		params: &FollowerRepositoryMockGetFollowersParams{ctx, userID, topN},
	}
	mmGetFollowers.expectations = append(mmGetFollowers.expectations, expectation)
	return expectation
}

// Then sets up FollowerRepository.GetFollowers return parameters for the expectation previously defined by the When method
func (e *FollowerRepositoryMockGetFollowersExpectation) Then(ia1 []int, err error) *FollowerRepositoryMock {
	e.results = &FollowerRepositoryMockGetFollowersResults{ia1, err}
	return e.mock
}

// GetFollowers implements usecase.FollowerRepository
func (mmGetFollowers *FollowerRepositoryMock) GetFollowers(ctx context.Context, userID int, topN int) (ia1 []int, err error) {
	mm_atomic.AddUint64(&mmGetFollowers.beforeGetFollowersCounter, 1)
	defer mm_atomic.AddUint64(&mmGetFollowers.afterGetFollowersCounter, 1)

	if mmGetFollowers.inspectFuncGetFollowers != nil {
		mmGetFollowers.inspectFuncGetFollowers(ctx, userID, topN)
	}

	mm_params := &FollowerRepositoryMockGetFollowersParams{ctx, userID, topN}

	// Record call args
	mmGetFollowers.GetFollowersMock.mutex.Lock()
	mmGetFollowers.GetFollowersMock.callArgs = append(mmGetFollowers.GetFollowersMock.callArgs, mm_params)
	mmGetFollowers.GetFollowersMock.mutex.Unlock()

	for _, e := range mmGetFollowers.GetFollowersMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ia1, e.results.err
		}
	}

	if mmGetFollowers.GetFollowersMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetFollowers.GetFollowersMock.defaultExpectation.Counter, 1)
		mm_want := mmGetFollowers.GetFollowersMock.defaultExpectation.params
		mm_got := FollowerRepositoryMockGetFollowersParams{ctx, userID, topN}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetFollowers.t.Errorf("FollowerRepositoryMock.GetFollowers got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetFollowers.GetFollowersMock.defaultExpectation.results
		if mm_results == nil {
			mmGetFollowers.t.Fatal("No results are set for the FollowerRepositoryMock.GetFollowers")
		}
		return (*mm_results).ia1, (*mm_results).err
	}
	if mmGetFollowers.funcGetFollowers != nil {
		return mmGetFollowers.funcGetFollowers(ctx, userID, topN)
	}
	mmGetFollowers.t.Fatalf("Unexpected call to FollowerRepositoryMock.GetFollowers. %v %v %v", ctx, userID, topN)
	return
}

// GetFollowersAfterCounter returns a count of finished FollowerRepositoryMock.GetFollowers invocations
func (mmGetFollowers *FollowerRepositoryMock) GetFollowersAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFollowers.afterGetFollowersCounter)
}

// GetFollowersBeforeCounter returns a count of FollowerRepositoryMock.GetFollowers invocations
func (mmGetFollowers *FollowerRepositoryMock) GetFollowersBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFollowers.beforeGetFollowersCounter)
}

// Calls returns a list of arguments used in each call to FollowerRepositoryMock.GetFollowers.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetFollowers *mFollowerRepositoryMockGetFollowers) Calls() []*FollowerRepositoryMockGetFollowersParams {
	mmGetFollowers.mutex.RLock()

	argCopy := make([]*FollowerRepositoryMockGetFollowersParams, len(mmGetFollowers.callArgs))
	copy(argCopy, mmGetFollowers.callArgs)

	mmGetFollowers.mutex.RUnlock()

	return argCopy
}

// MinimockGetFollowersDone returns true if the count of the GetFollowers invocations corresponds
// the number of defined expectations
func (m *FollowerRepositoryMock) MinimockGetFollowersDone() bool {
	for _, e := range m.GetFollowersMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFollowersMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFollowersCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFollowers != nil && mm_atomic.LoadUint64(&m.afterGetFollowersCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetFollowersInspect logs each unmet expectation
func (m *FollowerRepositoryMock) MinimockGetFollowersInspect() {
	for _, e := range m.GetFollowersMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FollowerRepositoryMock.GetFollowers with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFollowersMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFollowersCounter) < 1 {
		if m.GetFollowersMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FollowerRepositoryMock.GetFollowers")
		} else {
			m.t.Errorf("Expected call to FollowerRepositoryMock.GetFollowers with params: %#v", *m.GetFollowersMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFollowers != nil && mm_atomic.LoadUint64(&m.afterGetFollowersCounter) < 1 {
		m.t.Error("Expected call to FollowerRepositoryMock.GetFollowers")
	}
}

type mFollowerRepositoryMockGetFollowing struct {
	mock               *FollowerRepositoryMock
	defaultExpectation *FollowerRepositoryMockGetFollowingExpectation
	expectations       []*FollowerRepositoryMockGetFollowingExpectation

	callArgs []*FollowerRepositoryMockGetFollowingParams
	mutex    sync.RWMutex
}

// FollowerRepositoryMockGetFollowingExpectation specifies expectation struct of the FollowerRepository.GetFollowing
type FollowerRepositoryMockGetFollowingExpectation struct {
	mock    *FollowerRepositoryMock
	params  *FollowerRepositoryMockGetFollowingParams
	results *FollowerRepositoryMockGetFollowingResults
	Counter uint64
}

// FollowerRepositoryMockGetFollowingParams contains parameters of the FollowerRepository.GetFollowing
type FollowerRepositoryMockGetFollowingParams struct {
	ctx    context.Context
	userID int
	topN   int
}

// FollowerRepositoryMockGetFollowingResults contains results of the FollowerRepository.GetFollowing
type FollowerRepositoryMockGetFollowingResults struct {
	ia1 []int
	err error
}

// Expect sets up expected params for FollowerRepository.GetFollowing
func (mmGetFollowing *mFollowerRepositoryMockGetFollowing) Expect(ctx context.Context, userID int, topN int) *mFollowerRepositoryMockGetFollowing {
	if mmGetFollowing.mock.funcGetFollowing != nil {
		mmGetFollowing.mock.t.Fatalf("FollowerRepositoryMock.GetFollowing mock is already set by Set")
	}

	if mmGetFollowing.defaultExpectation == nil {
		mmGetFollowing.defaultExpectation = &FollowerRepositoryMockGetFollowingExpectation{}
	}

	mmGetFollowing.defaultExpectation.params = &FollowerRepositoryMockGetFollowingParams{ctx, userID, topN}
	for _, e := range mmGetFollowing.expectations {
		if minimock.Equal(e.params, mmGetFollowing.defaultExpectation.params) {
			mmGetFollowing.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetFollowing.defaultExpectation.params)
		}
	}

	return mmGetFollowing
}

// Inspect accepts an inspector function that has same arguments as the FollowerRepository.GetFollowing
func (mmGetFollowing *mFollowerRepositoryMockGetFollowing) Inspect(f func(ctx context.Context, userID int, topN int)) *mFollowerRepositoryMockGetFollowing {
	if mmGetFollowing.mock.inspectFuncGetFollowing != nil {
		mmGetFollowing.mock.t.Fatalf("Inspect function is already set for FollowerRepositoryMock.GetFollowing")
	}

	mmGetFollowing.mock.inspectFuncGetFollowing = f

	return mmGetFollowing
}

// Return sets up results that will be returned by FollowerRepository.GetFollowing
func (mmGetFollowing *mFollowerRepositoryMockGetFollowing) Return(ia1 []int, err error) *FollowerRepositoryMock {
	if mmGetFollowing.mock.funcGetFollowing != nil {
		mmGetFollowing.mock.t.Fatalf("FollowerRepositoryMock.GetFollowing mock is already set by Set")
	}

	if mmGetFollowing.defaultExpectation == nil {
		mmGetFollowing.defaultExpectation = &FollowerRepositoryMockGetFollowingExpectation{mock: mmGetFollowing.mock}
	}
	mmGetFollowing.defaultExpectation.results = &FollowerRepositoryMockGetFollowingResults{ia1, err}
	return mmGetFollowing.mock
}

//Set uses given function f to mock the FollowerRepository.GetFollowing method
func (mmGetFollowing *mFollowerRepositoryMockGetFollowing) Set(f func(ctx context.Context, userID int, topN int) (ia1 []int, err error)) *FollowerRepositoryMock {
	if mmGetFollowing.defaultExpectation != nil {
		mmGetFollowing.mock.t.Fatalf("Default expectation is already set for the FollowerRepository.GetFollowing method")
	}

	if len(mmGetFollowing.expectations) > 0 {
		mmGetFollowing.mock.t.Fatalf("Some expectations are already set for the FollowerRepository.GetFollowing method")
	}

	mmGetFollowing.mock.funcGetFollowing = f
	return mmGetFollowing.mock
}

// When sets expectation for the FollowerRepository.GetFollowing which will trigger the result defined by the following
// Then helper
func (mmGetFollowing *mFollowerRepositoryMockGetFollowing) When(ctx context.Context, userID int, topN int) *FollowerRepositoryMockGetFollowingExpectation {
	if mmGetFollowing.mock.funcGetFollowing != nil {
		mmGetFollowing.mock.t.Fatalf("FollowerRepositoryMock.GetFollowing mock is already set by Set")
	}

	expectation := &FollowerRepositoryMockGetFollowingExpectation{
		mock:   mmGetFollowing.mock,
		params: &FollowerRepositoryMockGetFollowingParams{ctx, userID, topN},
	}
	mmGetFollowing.expectations = append(mmGetFollowing.expectations, expectation)
	return expectation
}

// Then sets up FollowerRepository.GetFollowing return parameters for the expectation previously defined by the When method
func (e *FollowerRepositoryMockGetFollowingExpectation) Then(ia1 []int, err error) *FollowerRepositoryMock {
	e.results = &FollowerRepositoryMockGetFollowingResults{ia1, err}
	return e.mock
}

// GetFollowing implements usecase.FollowerRepository
func (mmGetFollowing *FollowerRepositoryMock) GetFollowing(ctx context.Context, userID int, topN int) (ia1 []int, err error) {
	mm_atomic.AddUint64(&mmGetFollowing.beforeGetFollowingCounter, 1)
	defer mm_atomic.AddUint64(&mmGetFollowing.afterGetFollowingCounter, 1)

	if mmGetFollowing.inspectFuncGetFollowing != nil {
		mmGetFollowing.inspectFuncGetFollowing(ctx, userID, topN)
	}

	mm_params := &FollowerRepositoryMockGetFollowingParams{ctx, userID, topN}

	// Record call args
	mmGetFollowing.GetFollowingMock.mutex.Lock()
	mmGetFollowing.GetFollowingMock.callArgs = append(mmGetFollowing.GetFollowingMock.callArgs, mm_params)
	mmGetFollowing.GetFollowingMock.mutex.Unlock()

	for _, e := range mmGetFollowing.GetFollowingMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ia1, e.results.err
		}
	}

	if mmGetFollowing.GetFollowingMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetFollowing.GetFollowingMock.defaultExpectation.Counter, 1)
		mm_want := mmGetFollowing.GetFollowingMock.defaultExpectation.params
		mm_got := FollowerRepositoryMockGetFollowingParams{ctx, userID, topN}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetFollowing.t.Errorf("FollowerRepositoryMock.GetFollowing got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetFollowing.GetFollowingMock.defaultExpectation.results
		if mm_results == nil {
			mmGetFollowing.t.Fatal("No results are set for the FollowerRepositoryMock.GetFollowing")
		}
		return (*mm_results).ia1, (*mm_results).err
	}
	if mmGetFollowing.funcGetFollowing != nil {
		return mmGetFollowing.funcGetFollowing(ctx, userID, topN)
	}
	mmGetFollowing.t.Fatalf("Unexpected call to FollowerRepositoryMock.GetFollowing. %v %v %v", ctx, userID, topN)
	return
}

// GetFollowingAfterCounter returns a count of finished FollowerRepositoryMock.GetFollowing invocations
func (mmGetFollowing *FollowerRepositoryMock) GetFollowingAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFollowing.afterGetFollowingCounter)
}

// GetFollowingBeforeCounter returns a count of FollowerRepositoryMock.GetFollowing invocations
func (mmGetFollowing *FollowerRepositoryMock) GetFollowingBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFollowing.beforeGetFollowingCounter)
}

// Calls returns a list of arguments used in each call to FollowerRepositoryMock.GetFollowing.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetFollowing *mFollowerRepositoryMockGetFollowing) Calls() []*FollowerRepositoryMockGetFollowingParams {
	mmGetFollowing.mutex.RLock()

	argCopy := make([]*FollowerRepositoryMockGetFollowingParams, len(mmGetFollowing.callArgs))
	copy(argCopy, mmGetFollowing.callArgs)

	mmGetFollowing.mutex.RUnlock()

	return argCopy
}

// MinimockGetFollowingDone returns true if the count of the GetFollowing invocations corresponds
// the number of defined expectations
func (m *FollowerRepositoryMock) MinimockGetFollowingDone() bool {
	for _, e := range m.GetFollowingMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFollowingMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFollowingCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFollowing != nil && mm_atomic.LoadUint64(&m.afterGetFollowingCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetFollowingInspect logs each unmet expectation
func (m *FollowerRepositoryMock) MinimockGetFollowingInspect() {
	for _, e := range m.GetFollowingMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FollowerRepositoryMock.GetFollowing with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFollowingMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFollowingCounter) < 1 {
		if m.GetFollowingMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FollowerRepositoryMock.GetFollowing")
		} else {
			m.t.Errorf("Expected call to FollowerRepositoryMock.GetFollowing with params: %#v", *m.GetFollowingMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFollowing != nil && mm_atomic.LoadUint64(&m.afterGetFollowingCounter) < 1 {
		m.t.Error("Expected call to FollowerRepositoryMock.GetFollowing")
	}
}

type mFollowerRepositoryMockRemove struct {
	mock               *FollowerRepositoryMock
	defaultExpectation *FollowerRepositoryMockRemoveExpectation
	expectations       []*FollowerRepositoryMockRemoveExpectation

	callArgs []*FollowerRepositoryMockRemoveParams
	mutex    sync.RWMutex
}

// FollowerRepositoryMockRemoveExpectation specifies expectation struct of the FollowerRepository.Remove
type FollowerRepositoryMockRemoveExpectation struct {
	mock    *FollowerRepositoryMock
	params  *FollowerRepositoryMockRemoveParams
	results *FollowerRepositoryMockRemoveResults
	Counter uint64
}

// FollowerRepositoryMockRemoveParams contains parameters of the FollowerRepository.Remove
type FollowerRepositoryMockRemoveParams struct {
	ctx        context.Context
	followerID int
	fromUserID int
}

// FollowerRepositoryMockRemoveResults contains results of the FollowerRepository.Remove
type FollowerRepositoryMockRemoveResults struct {
	err error
}

// Expect sets up expected params for FollowerRepository.Remove
func (mmRemove *mFollowerRepositoryMockRemove) Expect(ctx context.Context, followerID int, fromUserID int) *mFollowerRepositoryMockRemove {
	if mmRemove.mock.funcRemove != nil {
		mmRemove.mock.t.Fatalf("FollowerRepositoryMock.Remove mock is already set by Set")
	}

	if mmRemove.defaultExpectation == nil {
		mmRemove.defaultExpectation = &FollowerRepositoryMockRemoveExpectation{}
	}

	mmRemove.defaultExpectation.params = &FollowerRepositoryMockRemoveParams{ctx, followerID, fromUserID}
	for _, e := range mmRemove.expectations {
		if minimock.Equal(e.params, mmRemove.defaultExpectation.params) {
			mmRemove.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRemove.defaultExpectation.params)
		}
	}

	return mmRemove
}

// Inspect accepts an inspector function that has same arguments as the FollowerRepository.Remove
func (mmRemove *mFollowerRepositoryMockRemove) Inspect(f func(ctx context.Context, followerID int, fromUserID int)) *mFollowerRepositoryMockRemove {
	if mmRemove.mock.inspectFuncRemove != nil {
		mmRemove.mock.t.Fatalf("Inspect function is already set for FollowerRepositoryMock.Remove")
	}

	mmRemove.mock.inspectFuncRemove = f

	return mmRemove
}

// Return sets up results that will be returned by FollowerRepository.Remove
func (mmRemove *mFollowerRepositoryMockRemove) Return(err error) *FollowerRepositoryMock {
	if mmRemove.mock.funcRemove != nil {
		mmRemove.mock.t.Fatalf("FollowerRepositoryMock.Remove mock is already set by Set")
	}

	if mmRemove.defaultExpectation == nil {
		mmRemove.defaultExpectation = &FollowerRepositoryMockRemoveExpectation{mock: mmRemove.mock}
	}
	mmRemove.defaultExpectation.results = &FollowerRepositoryMockRemoveResults{err}
	return mmRemove.mock
}

//Set uses given function f to mock the FollowerRepository.Remove method
func (mmRemove *mFollowerRepositoryMockRemove) Set(f func(ctx context.Context, followerID int, fromUserID int) (err error)) *FollowerRepositoryMock {
	if mmRemove.defaultExpectation != nil {
		mmRemove.mock.t.Fatalf("Default expectation is already set for the FollowerRepository.Remove method")
	}

	if len(mmRemove.expectations) > 0 {
		mmRemove.mock.t.Fatalf("Some expectations are already set for the FollowerRepository.Remove method")
	}

	mmRemove.mock.funcRemove = f
	return mmRemove.mock
}

// When sets expectation for the FollowerRepository.Remove which will trigger the result defined by the following
// Then helper
func (mmRemove *mFollowerRepositoryMockRemove) When(ctx context.Context, followerID int, fromUserID int) *FollowerRepositoryMockRemoveExpectation {
	if mmRemove.mock.funcRemove != nil {
		mmRemove.mock.t.Fatalf("FollowerRepositoryMock.Remove mock is already set by Set")
	}

	expectation := &FollowerRepositoryMockRemoveExpectation{
		mock:   mmRemove.mock,
		params: &FollowerRepositoryMockRemoveParams{ctx, followerID, fromUserID},
	}
	mmRemove.expectations = append(mmRemove.expectations, expectation)
	return expectation
}

// Then sets up FollowerRepository.Remove return parameters for the expectation previously defined by the When method
func (e *FollowerRepositoryMockRemoveExpectation) Then(err error) *FollowerRepositoryMock {
	e.results = &FollowerRepositoryMockRemoveResults{err}
	return e.mock
}

// Remove implements usecase.FollowerRepository
func (mmRemove *FollowerRepositoryMock) Remove(ctx context.Context, followerID int, fromUserID int) (err error) {
	mm_atomic.AddUint64(&mmRemove.beforeRemoveCounter, 1)
	defer mm_atomic.AddUint64(&mmRemove.afterRemoveCounter, 1)

	if mmRemove.inspectFuncRemove != nil {
		mmRemove.inspectFuncRemove(ctx, followerID, fromUserID)
	}

	mm_params := &FollowerRepositoryMockRemoveParams{ctx, followerID, fromUserID}

	// Record call args
	mmRemove.RemoveMock.mutex.Lock()
	mmRemove.RemoveMock.callArgs = append(mmRemove.RemoveMock.callArgs, mm_params)
	mmRemove.RemoveMock.mutex.Unlock()

	for _, e := range mmRemove.RemoveMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmRemove.RemoveMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRemove.RemoveMock.defaultExpectation.Counter, 1)
		mm_want := mmRemove.RemoveMock.defaultExpectation.params
		mm_got := FollowerRepositoryMockRemoveParams{ctx, followerID, fromUserID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRemove.t.Errorf("FollowerRepositoryMock.Remove got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRemove.RemoveMock.defaultExpectation.results
		if mm_results == nil {
			mmRemove.t.Fatal("No results are set for the FollowerRepositoryMock.Remove")
		}
		return (*mm_results).err
	}
	if mmRemove.funcRemove != nil {
		return mmRemove.funcRemove(ctx, followerID, fromUserID)
	}
	mmRemove.t.Fatalf("Unexpected call to FollowerRepositoryMock.Remove. %v %v %v", ctx, followerID, fromUserID)
	return
}

// RemoveAfterCounter returns a count of finished FollowerRepositoryMock.Remove invocations
func (mmRemove *FollowerRepositoryMock) RemoveAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRemove.afterRemoveCounter)
}

// RemoveBeforeCounter returns a count of FollowerRepositoryMock.Remove invocations
func (mmRemove *FollowerRepositoryMock) RemoveBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRemove.beforeRemoveCounter)
}

// Calls returns a list of arguments used in each call to FollowerRepositoryMock.Remove.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRemove *mFollowerRepositoryMockRemove) Calls() []*FollowerRepositoryMockRemoveParams {
	mmRemove.mutex.RLock()

	argCopy := make([]*FollowerRepositoryMockRemoveParams, len(mmRemove.callArgs))
	copy(argCopy, mmRemove.callArgs)

	mmRemove.mutex.RUnlock()

	return argCopy
}

// MinimockRemoveDone returns true if the count of the Remove invocations corresponds
// the number of defined expectations
func (m *FollowerRepositoryMock) MinimockRemoveDone() bool {
	for _, e := range m.RemoveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RemoveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRemoveCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRemove != nil && mm_atomic.LoadUint64(&m.afterRemoveCounter) < 1 {
		return false
	}
	return true
}

// MinimockRemoveInspect logs each unmet expectation
func (m *FollowerRepositoryMock) MinimockRemoveInspect() {
	for _, e := range m.RemoveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FollowerRepositoryMock.Remove with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RemoveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRemoveCounter) < 1 {
		if m.RemoveMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FollowerRepositoryMock.Remove")
		} else {
			m.t.Errorf("Expected call to FollowerRepositoryMock.Remove with params: %#v", *m.RemoveMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRemove != nil && mm_atomic.LoadUint64(&m.afterRemoveCounter) < 1 {
		m.t.Error("Expected call to FollowerRepositoryMock.Remove")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *FollowerRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddInspect()

		m.MinimockGetFollowersInspect()

		m.MinimockGetFollowingInspect()

		m.MinimockRemoveInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *FollowerRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *FollowerRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddDone() &&
		m.MinimockGetFollowersDone() &&
		m.MinimockGetFollowingDone() &&
		m.MinimockRemoveDone()
}