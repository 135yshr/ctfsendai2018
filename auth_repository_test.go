package ctfsendai2018

import (
	"reflect"
	"testing"
)

func TestNewAuthRepository(t *testing.T) {
	type args struct {
	}
	tests := map[string]struct {
		args args
		want AuthRepository
		err  error
	}{
		"AuthRepositoriesのインスタンスが作成できること": {
			args: args{},
			want: &authRepository{},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewAuthRepository()
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}
		})
	}
}
