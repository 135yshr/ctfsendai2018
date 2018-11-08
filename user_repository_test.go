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

func TestAddUser(t *testing.T) {
	type args struct {
		users []*User
	}
	tests := map[string]struct {
		args args
		want []*User
		err  error
	}{
		"アカウント登録を１回実行してユーザー情報を１件取得できること": {
			args: args{users: []*User{&User{}}},
			want: []*User{
				&User{},
			},
		},
		"アカウント登録を２回実行してユーザー情報を２件取得できること": {
			args: args{
				users: []*User{
					&User{Name: "hoge hoge"},
					&User{Name: "huga huga"},
				},
			},
			want: []*User{
				&User{Name: "hoge hoge"},
				&User{Name: "huga huga"},
			},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewUserRepository()
			for _, u := range arg.args.users {
				err := sut.Add(u)
				if reflect.DeepEqual(err, arg.err) == false {
					t.Errorf("Error actual: %v, expected: %v", err, arg.err)
				}
			}
			list, _ := sut.List()
			if reflect.DeepEqual(list, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", list, arg.want)
			}
		})
	}
}

func TestUserRepository_FetchByEMail(t *testing.T) {
	type args struct {
		email string
		users []*User
	}
	tests := map[string]struct {
		args args
		want *User
		err  error
	}{
		"メールアドレスがhoge@hoge.comのユーザー情報を登録しhoge@hoge.comを指定してユーザー情報を取得できること": {
			args: args{
				email: "hoge@hoge.com",
				users: []*User{
					&User{EMail: "hoge@hoge.com"},
					&User{EMail: "fuga@fuga.com"},
				},
			},
			want: &User{EMail: "hoge@hoge.com"},
		},
		"メールアドレスがfuga@fuga.comのユーザー情報を登録しfuga@fuga.comを指定してユーザー情報を取得できること": {
			args: args{
				email: "fuga@fuga.com",
				users: []*User{
					&User{EMail: "hoge@hoge.com"},
					&User{EMail: "fuga@fuga.com"},
				},
			},
			want: &User{EMail: "fuga@fuga.com"},
		},
	}
	for testName, arg := range tests {
		sut := NewUserRepository()
		for _, u := range arg.args.users {
			sut.Add(u)
		}
		t.Run(testName, func(t *testing.T) {
			u, err := sut.FetchByEMail(arg.args.email)
			if reflect.DeepEqual(err, arg.err) == false {
				t.Errorf("Error actual: %v, expected: %v", err, arg.err)
			}
			if reflect.DeepEqual(u, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", u, arg.want)
			}
		})
	}
}
