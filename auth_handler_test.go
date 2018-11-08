package ctfsendai2018

import (
	"reflect"
	"testing"
)

func TestNewHander(t *testing.T) {
	type args struct {
		uc AuthUsecase
	}
	tests := map[string]struct {
		args args
		want AuthHandler
		err  error
	}{
		"AuthHanlderのインスタンスが作成できること": {
			args: args{
				uc: NewAuthUsecase(NewAuthRepository(), NewUserRepository()),
			},
			want: &authHandler{NewAuthUsecase(NewAuthRepository(), NewUserRepository())},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewAuthHandler(arg.args.uc)
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}
		})
	}
}
