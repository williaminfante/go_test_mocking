package mocks

import "github.com/stretchr/testify/mock"

type MockedMaintenanceWrapper struct {
	mock.Mock
}

func (w *MockedMaintenanceWrapper) NeedsMaintenance(days int) (decision bool, err error) {
	args := w.Called(days)
	if len(args) > 1 {
		return args.Get(0).(bool), args.Get(1).(error)
	}
	return args.Get(0).(bool), nil
}
