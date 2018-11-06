package ctfsendai2018

import "testing"

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
		"氏名が空のときにインスタンスがnilになり、エラーが返ってくること": {},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewUser(arg.args.name)
			if sut != arg.want {
				t.Errorf("Not equals actual: %v, expected: %v", sut, arg.want)
				return
			}
		})
	}
}
