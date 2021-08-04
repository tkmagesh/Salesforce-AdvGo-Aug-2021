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

//sub tests
/* func TestUtils(t *testing.T) {
	t.Run("Number Tests", func(t *testing.T) {
		t.Run("Prime Test", func(t *testing.T) {
			t.Run("Testing Prime for 11", func(t *testing.T) {
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

			t.Run("Testing Prime for 13", func(t *testing.T) {
				//Arrage
				no := 11
				expectedResult := true

				//Act
				result := IsPrime(no)

				//Assert
				if result != expectedResult {
					t.Fail()
				}

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
	})
} */

//table drive tests

type PrimeTestCase struct {
	name           string
	no             int
	expectedResult bool
}

func TestUtils(t *testing.T) {
	t.Run("Number Tests", func(t *testing.T) {
		t.Run("Prime Test", func(t *testing.T) {
			testCases := []PrimeTestCase{
				{"Testing Prime for 11", 11, true},
				{"Testing Prime for 13", 13, true},
				{"Testing Prime for 17", 17, true},
				{"Testing Prime for 19", 19, true},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					//Arrage
					no := testCase.no

					//Act
					result := IsPrime(no)

					//Assert
					if result != testCase.expectedResult {
						t.Fail()
					}
				})
			}
		})
	})
}
