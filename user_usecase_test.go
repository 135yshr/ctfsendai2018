package ctfsendai2018

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewUserUsecase(t *testing.T) {
	rep := NewUserRepository()
	type args struct {
	}
	tests := map[string]struct {
		args args
		want UserUsecase
		err  error
	}{
		"UserUsecaseのインスタンスが作成できること": {
			want: &userUsecase{rep},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewUserUsecase(NewUserRepository())
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}
		})
	}
}

func TestUserUsecase_Add(t *testing.T) {
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
					ID:       "aG9nZUBob2dlLmNvbQ",
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
			sut := NewUserUsecase(NewUserRepository())
			err := sut.Add(arg.args.auth, arg.args.user)
			if reflect.DeepEqual(err, arg.err) == false {
				t.Errorf("Error actual: %v, expected: %v", err, arg.err)
			}
			list, _ := sut.List(arg.args.auth)
			if len(list) != len(arg.want) {
				t.Errorf("Not equals count actual: %d, expected: %d", len(list), len(arg.want))
			}
			for i := 0; i < len(list); i++ {
				if reflect.DeepEqual(list[i], arg.want[i]) == false {
					t.Errorf("Not equals actual: %v, expected: %v", list[i], arg.want[i])
				}
			}
		})
	}
}
