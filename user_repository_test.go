package ctfsendai2018

import (
	"reflect"
	"testing"
)

func TestCreateInstance(t *testing.T) {
	type args struct {
	}
	tests := map[string]struct {
		args args
		want UserRepository
		err  error
	}{
		"UserRepositoryのインスタンスが作成できること": {
			want: &userRepository{},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewUserRepository()
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}
		})
	}
}
