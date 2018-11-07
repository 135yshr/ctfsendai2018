package ctfsendai2018

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		name  string
		age   uint8
		email string
		note  string
		pass  string
		auth  uint8
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
			want: &User{Name: "hoge fuga"},
		},
		"氏名が空のときにインスタンスがnilになり、エラーが返ってくること": {
			err: errors.New("name can not be empty"),
		},
		"氏名にhoge fugaを指定してNameからhoge fugaが取得できること": {
			args: args{name: "hoge fuga"},
			want: &User{Name: "hoge fuga"},
		},
		"氏名にfoo barを指定してNameからfoo barが取得できること": {
			args: args{name: "foo bar"},
			want: &User{Name: "foo bar"},
		},
		"年齢に18を指定してAgeから18が取得できること": {
			args: args{name: "foo bar", age: 18},
			want: &User{Name: "foo bar", Age: 18},
		},
		"年齢に20を指定してAgeから20が取得できること": {
			args: args{name: "foo bar", age: 20},
			want: &User{Name: "foo bar", Age: 20},
		},
		"メールアドレスにhoge@hoge.comを指定してEmailからhoge@hoge.comを取得できること": {
			args: args{name: "foo bar", email: "hoge@hoge.com"},
			want: &User{Name: "foo bar", EMail: "hoge@hoge.com"},
		},
		"メールアドレスにfoo@foo.co.jpを指定してEmailからfoo@foo.co.jpを取得できること": {
			args: args{name: "foo bar", email: "foo@foo.co.jp"},
			want: &User{Name: "foo bar", EMail: "foo@foo.co.jp"},
		},
		"備考にxxxxxxを指定してNoteからxxxxxが取得できること": {
			args: args{name: "foo bar", note: "xxxxxx"},
			want: &User{Name: "foo bar", Note: "xxxxxx"},
		},
		"備考にyyyyyyを指定してNoteからyyyyyが取得できること": {
			args: args{name: "foo bar", note: "yyyyyy"},
			want: &User{Name: "foo bar", Note: "yyyyyy"},
		},
		"パスワードにpasswordを指定してPasswordからpasswordが取得できること": {
			args: args{name: "foo bar", pass: "password"},
			want: &User{Name: "foo bar", Password: "password"},
		},
		"パスワードにpasspassを指定してPasswordからpasspassが取得できること": {
			args: args{name: "foo bar", pass: "passpass"},
			want: &User{Name: "foo bar", Password: "passpass"},
		},
		"権限にAdmin(1)を指定してAuthorityからAdmin(1)が取得できること": {
			args: args{name: "foo bar", auth: 1},
			want: &User{Name: "foo bar", Auth: 1},
		},
		"権限にWorker(2)を指定してAuthorityからWorker(2)が取得できること": {
			args: args{name: "foo bar", auth: 2},
			want: &User{Name: "foo bar", Auth: 2},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut, err := NewUser(arg.args.name, arg.args.age, arg.args.email, arg.args.note, arg.args.pass, arg.args.auth)
			if reflect.DeepEqual(err, arg.err) == false {
				t.Errorf("Error actual: %v, expected: %v", err, arg.err)
			}
			if reflect.DeepEqual(sut, arg.want) == false {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
			}
		})
	}
}
