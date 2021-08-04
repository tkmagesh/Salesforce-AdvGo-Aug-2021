package utils

import "testing"

/* func TestIsPrime(t *testing.T) {
	//Arrage
	no := 11
	expectedResult := true

	//Act
	result := IsPrime(no)

	//Assert
	if result != expectedResult {
		t.Fail()
	}
}

func TestIsNotPrime(t *testing.T) {
	//Arrange
	no := 100
	expectedResult := false

	//Act
	result := IsPrime(no)

	//Assert
	if result != expectedResult {
		t.Fail()
	}
} */

func TestUtils(t *testing.T) {
	t.Run("Number Tests", func(t *testing.T) {
		t.Run("Prime Test", func(t *testing.T) {
			t.Run("Testing Prime", func(t *testing.T) {
				//Arrage
				no := 11
				expectedResult := true

				//Act
				result := IsPrime(no)

				//Assert
				if result != expectedResult {
					t.Fail()
				}
			})

			t.Run("Testing Not Prime", func(t *testing.T) {
				no := 100
				expectedResult := false

				//Act
				result := IsPrime(no)

				//Assert
				if result != expectedResult {
					t.Fail()
				}
			})
		})

		t.Run("Even Tests", func(t *testing.T) {
			if IsEven(10) != true {
				t.Fail()
			}
		})
	})
}
