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

func TestLogin(t *testing.T) {
	type args struct {
		user *User
	}
	tests := map[string]struct {
		args args
		want *Authentication
		err  error
	}{
		"アカウントを登録した後登録したユーザーでログインできること": {
			args: args{
				user: &User{
					EMail:    "hoge@hoge.com",
					Password: "pass",
				},
			},
			want: &Authentication{Token: "eyJleHAiOjE0OTQ1MTQyNjksInVzZXIiOiLnrqHnkIbogIUifQo="},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewAuthRepository()
			auth, err := sut.Login(arg.args.user)
			if err != arg.err {
				t.Errorf("Error actual: %v, expected: %v", err, arg.err)
			}
			if reflect.DeepEqual(auth, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", auth, arg.want)
			}
		})
	}
}
