package service

import (
	"example.com/go-gin-todolist/domain/entity"
	"example.com/go-gin-todolist/domain/repository"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestTaskService_GetAllTasks(t *testing.T) {
	tests := []struct {
		name          string
		prepareMockFn func(m *repository.MockTaskRepository)
		want          []entity.Task
	}{
		{
			name: "Success",
			prepareMockFn: func(m *repository.MockTaskRepository) {
				m.EXPECT().FindAllTasks().Return([]entity.Task{{0, "text0"}, {1, "text1"}})
			},
			want: []entity.Task{{0, "text0"}, {1, "text1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			// mock generation
			mock := repository.NewMockTaskRepository(ctrl)
			// Arrange
			tt.prepareMockFn(mock)
			ts := &TaskService{
				repository: mock,
			}

			// Act & Assert
			if got := ts.GetAllTasks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
