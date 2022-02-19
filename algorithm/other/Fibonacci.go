package other

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciFor(n int) int {
	var mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	fn1, fn2, fn := 0, 0, 1
	for i := 2; i <= n; i++ {
		fn1 = fn2
		fn2 = fn
		fn = (fn1 + fn2) % mod
	}
	return fn
}
