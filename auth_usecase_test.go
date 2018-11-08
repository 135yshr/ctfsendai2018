package ctfsendai2018

import (
	"reflect"
	"testing"
)

func TestNewAuthUsecase(t *testing.T) {
	rep := NewAuthRepository()
	type args struct {
	}
	tests := map[string]struct {
		args args
		want AuthUsecase
		err  error
	}{
		"AuthUsecaseのインスタンスが作成できること": {
			want: &authService{rep},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewAuthService(rep)
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}

		})
	}
}
