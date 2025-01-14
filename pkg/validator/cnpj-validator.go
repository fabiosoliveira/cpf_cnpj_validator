package validator

var v_CNPJ = [13]int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

func IsValidCNPJ(cnpj *string) bool {
	if hasInvalidLength(cnpj, 14) {
		return false
	}

	cnpjSlice, ok := stringToSliceInt(cnpj)
	if !ok {
		return false
	}

	if intIsRepeated(cnpjSlice) {
		return false
	}

	firstVerifierDigit := calculateFirstVerifierDigitCnpj(cnpjSlice)

	if cnpjSlice[12] != firstVerifierDigit {
		return false
	}

	secondVerifierDigit := calculateSecondVerifierDigitCnpj(cnpjSlice, firstVerifierDigit)

	return cnpjSlice[13] == secondVerifierDigit
}

func calculateFirstVerifierDigitCnpj(cnpj []int) int {
	var sum int
	_CNPJ := v_CNPJ[1:]
	for i, v := range cnpj[:12] {
		sum += v * _CNPJ[i]
	}

	result := sum % 11

	if result < 2 {
		return 0
	}

	return 11 - result
}

func calculateSecondVerifierDigitCnpj(cnpj []int, firstVerifierDigit int) int {
	var sum int

	for i, v := range cnpj[:12] {
		sum += v * v_CNPJ[i]
	}

	sum += (firstVerifierDigit * v_CNPJ[12])

	result := sum % 11

	if result < 2 {
		return 0
	}

	return 11 - result
}
