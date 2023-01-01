## Testing

You should mock dependencies using following way with minimock library:
```go
func Test_Template(t *testing.T) {
	type mocks struct {
		// *mock.OneMock
		// *mock.TwoMock
	}
	// you need to create feed manager and mocks for each testcase


	var (
		_ = "fake variable"
	)

	testCases := []struct {
		name string
		// your args
		expect func(mocks)
		// want value
		wantError bool
	}{
		{},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
            mc := minimock.NewController(t)
            defer mc.Finish
            m := mocks{
                // mock.NewOneMock(mc), mock.NewTwoMock(mc),
            }
            tc.expect(m)

            s := &some{
                // one: m.OneMock, // two: m.TwoMock
            }
			err := s.CallMethod(tc.arg)
			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
		})
	}
}
```