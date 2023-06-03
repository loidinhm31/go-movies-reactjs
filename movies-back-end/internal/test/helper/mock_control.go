package helper

import "github.com/stretchr/testify/mock"

// MockManagementCtrl is a mock implementation of the management control interface for testing.
type MockManagementCtrl struct {
	mock.Mock
}

func (m *MockManagementCtrl) CheckPrivilege(username string) bool {
	args := m.Called(username)
	return args.Bool(0)
}

func (m *MockManagementCtrl) CheckAdminPrivilege(username string) bool {
	args := m.Called(username)
	return args.Bool(0)
}

func (m *MockManagementCtrl) CheckUser(username string) (bool, bool) {
	args := m.Called(username)
	return args.Bool(0), args.Bool(1)
}
