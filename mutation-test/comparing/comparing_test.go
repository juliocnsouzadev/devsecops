package comparing

import (
	"testing"
)

func Test_isEqual(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1 == 1", args{1, 1}, true},
		{"2 == 1", args{2, 1}, false},
		{"2 == 2", args{2, 2}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isGreater(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1 > 1", args{1, 1}, false},
		{"2 > 1", args{2, 1}, true},
		{"2 > 3", args{2, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGreater(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isGreater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLess(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1 < 1", args{1, 1}, false},
		{"2 < 1", args{2, 1}, false},
		{"2 < 3", args{2, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLess(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isLess() = %v, want %v", got, tt.want)
			}
		})
	}
}
