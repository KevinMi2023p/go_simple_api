package amountValidation

var totalCentValue int = 0

// ResetCentValueCounter resets the total cent value counter to zero.
func ResetCentValueCounter() {
	totalCentValue = 0
}

// AddCentValueToTotal adds the input cent value to the total cent value counter.
func AddCentValueToTotal(input int) {
	totalCentValue += input
}

// GetTotalCentValue returns the current total cent value.
func GetTotalCentValue() int {
	return totalCentValue
}
