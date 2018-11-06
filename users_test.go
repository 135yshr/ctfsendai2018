package ctfsendai2018

import (
	"errors"
	"reflect"
	"testing"
)

func TestCreateInstance(t *testing.T) {
	type args struct {
		name string
	}
	tests := map[string]struct {
		args args
		want *User
		err  error
	}{
		"氏名にhoge fugaを指定してユーザーのインスタンスが作成されること": {
			args: args{
				name: "hoge fuga",
			},
			want: &User{},
		},
		"氏名が空のときにインスタンスがnilになり、エラーが返ってくること": {
			err: errors.New("name can not be empty"),
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut, err := NewUser(arg.args.name)
			if reflect.DeepEqual(err, arg.err) == false {
				t.Errorf("Error actual: %v, expected: %v", err, arg.err)
			}
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
				return
			}
		})
	}
}
