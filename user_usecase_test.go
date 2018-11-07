package ctfsendai2018

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewUserService(t *testing.T) {
	rep := NewUserRepository()
	type args struct {
	}
	tests := map[string]struct {
		args args
		want UserService
		err  error
	}{
		"UserServiceのインスタンスが作成できること": {
			want: &userService{rep},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewUserService(NewUserRepository())
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}
		})
	}
}

func TestUserService_Add(t *testing.T) {
	type args struct {
		auth uint8
		user *User
	}
	tests := map[string]struct {
		args args
		want []*User
		err  error
	}{
		"管理者アカウントのトークンを使ってユーザーの登録ができること": {
			args: args{
				auth: 1,
				user: &User{
					Name:     "hoge hoge",
					Age:      18,
					EMail:    "hoge@hoge.com",
					Password: "password",
					Auth:     1,
				},
			},
			want: []*User{
				&User{
					Name:     "hoge hoge",
					Age:      18,
					EMail:    "hoge@hoge.com",
					Password: "password",
					Auth:     1,
				},
			},
		},
		"一般アカウントのトークンを使ってユーザーの登録ができないこと": {
			args: args{
				auth: 2,
				user: &User{
					Name:     "hoge hoge",
					Age:      18,
					EMail:    "hoge@hoge.com",
					Password: "password",
					Auth:     1,
				},
			},
			want: []*User{},
			err:  errors.New("Don't have permission"),
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewUserService(NewUserRepository())
			err := sut.Add(arg.args.auth, arg.args.user)
			if reflect.DeepEqual(err, arg.err) == false {
				t.Errorf("Error actual: %v, expected: %v", err, arg.err)
			}
			list, _ := sut.List(arg.args.auth)
			if len(list) != len(arg.want) && reflect.DeepEqual(list, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", list, arg.want)
			}
		})
	}
}
