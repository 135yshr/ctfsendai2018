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
		"メールアドレスがhoge@hoge.comのログイン情報を作成できること": {
			args: args{
				user: &User{
					Name:     "hoge hoge",
					Auth:     2,
					EMail:    "hoge@hoge.com",
					Password: "pass",
				},
			},
			want: &Authentication{Token: "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJhdXRoIjoyLCJleHAiOjE5MjQ5NTIzOTksInVzZXIiOiJob2dlIGhvZ2UifQ."},
		},
		"メールアドレスがfuga@fuga.comのユーザー情報でログイン情報を作成できること": {
			args: args{
				user: &User{
					Name:     "fuga fuga",
					Auth:     2,
					EMail:    "fuga@fuga.com",
					Password: "pass",
				},
			},
			want: &Authentication{Token: "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJhdXRoIjoyLCJleHAiOjE5MjQ5NTIzOTksInVzZXIiOiJmdWdhIGZ1Z2EifQ."},
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
