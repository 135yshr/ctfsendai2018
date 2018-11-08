package ctfsendai2018

import (
	"reflect"
	"testing"
)

func TestNewUserHander(t *testing.T) {
	type args struct {
		uc UserUsecase
	}
	tests := map[string]struct {
		args args
		want UserHandler
		err  error
	}{
		"UserHanlderのインスタンスが作成できること": {
			args: args{
				uc: NewUserUsecase(NewUserRepository()),
			},
			want: &userHandler{NewUserUsecase(NewUserRepository())},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewUserHandler(arg.args.uc)
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}
		})
	}
}
