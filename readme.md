## Тестирование

Для мокирования зависимостей мы используем библиотеку minimock. 
Стандартный шаблон для табличных тестов выглядит следующим образом:
```go
func Test_Template(t *testing.T) {
	t.Parallel()

	var (
		_ = "fake"
	)

	type mocks struct {
		// *mock.SomeInterfaceMock
	}
	testCases := []struct {
		name string
		// your args
		setup func(mocks)
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
				// mock.NewSomeInterfaceMock(mc)
			}
			tc.setup(m)

			// init your usecase using mock interfaces
			// someUsecase := usecase{m.SomeInterface}
			// err := someUsecase.CallMethod(tc.arg)
			var err error
			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
		})
	}
}
```