package usetestify

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetData(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *Data
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Test 1",
			args:    args{id: "123"},
			want:    &Data{Value: "Hello, 123"},
			wantErr: assert.NoError,
		},
		{
			name:    "Test 2",
			args:    args{id: "123"},
			want:    &Data{Value: "Hello, 123"},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := Service{}
			got, err := se.GetData(tt.args.id)
			tt.wantErr(t, err)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestGetData(t *testing.T) {
	mockService := new(MockDataService)
	mockService.On("GetData", "123").Return(&Data{Value: "test data"}, nil)

	result, err := mockService.GetData("123")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test data", result.Value)

	mockService.AssertExpectations(t)
}

// ---- HELPER FUNCTIONS ----

type MockDataService struct {
	mock.Mock
}

func (m *MockDataService) GetData(id string) (*Data, error) {
	args := m.Called(id)
	return args.Get(0).(*Data), args.Error(1)
}
