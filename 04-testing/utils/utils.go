package utils

func IsPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsEven(no int) bool {
	return no%2 == 0
}
