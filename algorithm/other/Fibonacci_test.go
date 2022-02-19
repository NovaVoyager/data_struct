package other

import "testing"

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{n: 5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciFor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "3",
			args: args{n: 45},
			want: 134903163,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciFor(tt.args.n); got != tt.want {
				t.Errorf("FibonacciFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
