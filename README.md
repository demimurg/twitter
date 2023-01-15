[![codecov](https://codecov.io/github/demimurg/twitter/branch/main/graph/badge.svg?token=NMZND94SUY)](https://codecov.io/github/demimurg/twitter)
[![Go Report Card](https://goreportcard.com/badge/github.com/demimurg/twitter)](https://goreportcard.com/report/github.com/demimurg/twitter)

## Testing

You should mock dependencies using following way with minimock library:
```go
func Test_Template(t *testing.T) {
    var (
        _ = "fake variable"
    )

    type mocks struct {
        // *mock.OneMock
        // *mock.TwoMock
    }
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
            defer mc.Finish()
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