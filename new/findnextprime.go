package piscine

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func FindNextPrime(n int) int {
	if n <= 1 {
		return 2
	}
	prime := n
	for !isPrime(prime) {
		prime++
	}
	return prime
}
