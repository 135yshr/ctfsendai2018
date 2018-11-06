package ctfsendai2018

import "testing"

func TestCreateInstance(t *testing.T) {
	type args struct {
		name string
	}
	tests := map[string]struct {
		args args
		err  error
	}{
		"氏名にhoge fugaを指定してユーザーのインスタンスが作成されること": {
			args: args{
				name: "hoge fuga",
			},
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			sut := NewUser(arg.args.name)
			if sut == nil {
				t.Errorf("Instance is nil")
			}
		})
	}
}

