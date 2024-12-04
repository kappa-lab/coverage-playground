package cmd

import "testing"

func Test_hoge(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{v: ""},
			want: "empty",
		},
		{
			name: "s",
			args: args{v: "s"},
			want: "s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hoge(tt.args.v); got != tt.want {
				t.Errorf("hoge() = %v, want %v", got, tt.want)
			}
		})
	}
}
