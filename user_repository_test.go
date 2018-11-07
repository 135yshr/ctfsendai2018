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
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}
		})
	}
}
