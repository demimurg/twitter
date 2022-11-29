## Тестирование

Для мокирования зависимостей мы используем библиотеку minimock. 
Стандартный шаблон для табличных тестов выглядит следующим образом:
```go
func Test_Template(t *testing.T) {
	type mocks struct {
		// *mock.OneMock
		// *mock.TwoMock
	}
	// you need to create feed manager and mocks for each testcase
	setup := func(t *testing.T, expect func(mocks)) Some {
		mc := minimock.NewController(t)
		t.Cleanup(mc.Finish)

		m := mocks{
			// mock.NewOneMock(mc), mock.NewTwoMock(mc),
		}
		expect(m)

		return &some{
			// one: m.OneMock,
			// two: m.TwoMock,
		}
	}

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

			someUsecase := setup(t, tc.expect)
			err := someUsecase.CallMethod(tc.arg)
			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
		})
	}
}
```