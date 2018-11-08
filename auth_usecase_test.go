package ctfsendai2018

import (
	"reflect"
	"testing"
)

func TestNewAuthUsecase(t *testing.T) {
	authRep := NewAuthRepository()
	userRep := NewUserRepository()
	type args struct {
	}
	tests := map[string]struct {
		args args
		want AuthUsecase
		err  error
	}{
		"AuthUsecaseのインスタンスが作成できること": {
			want: &authUsecase{authRep, userRep},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewAuthUsecase(authRep, userRep)
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}

		})
	}
}

func TestAuthUsecase_Register(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := map[string]struct {
		args args
		want []*User
		err  error
	}{
		"メールアドレスhoge@hoge.comで新しくアカウントが登録できること": {
			args: args{"hoge@hoge.com", "pass"},
			want: []*User{
				&User{
					EMail:    "hoge@hoge.com",
					Password: "pass",
					ID:       "aG9nZUBob2dlLmNvbQ",
					Auth:     2,
				},
			},
		},
	}
	for testName, arg := range tests {
		userRep := NewUserRepository()
		t.Run(testName, func(t *testing.T) {
			sut := NewAuthUsecase(NewAuthRepository(), userRep)
			if err := sut.Register(arg.args.email, arg.args.password); reflect.DeepEqual(err, arg.err) == false {
				t.Errorf("Error actual: %v, expected: %v", err, arg.err)
			}
			list, _ := userRep.List()
			if reflect.DeepEqual(list, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", list, arg.want)
			}
		})
	}
}
