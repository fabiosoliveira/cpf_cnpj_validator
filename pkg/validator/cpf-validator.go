package validator

func IsValidCPF(cpf *string) bool {

	if hasInvalidLength(cpf, 11) {
		return false
	}

	cpfSlice, ok := stringToSliceInt(cpf)
	if !ok {
		return false
	}

	if intIsRepeated(cpfSlice) {
		return false
	}

	firstVerifierDigit := calculateFirstVerifierDigit(cpfSlice)

	if cpfSlice[9] != firstVerifierDigit {
		return false
	}

	secondVerifierDigit := calculateSecondVerifierDigit(cpfSlice, firstVerifierDigit)

	return cpfSlice[10] == secondVerifierDigit
}

func hasInvalidLength(cpf *string, size int) bool {
	return len(*cpf) != size
}

func stringToSliceInt(s *string) ([]int, bool) {
	var sliceInt []int

	for _, r := range *s {
		if !runeIsInt(r) {
			return nil, false
		}

		sliceInt = append(sliceInt, runeToInt(r))
	}

	return sliceInt, true
}

func runeToInt(r rune) int {
	return int(r - '0')
}

func runeIsInt(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}

	return false
}

func intIsRepeated(sliceInt []int) bool {
	for i := 1; i < len(sliceInt); i++ {
		if sliceInt[i] != sliceInt[i-1] {
			return false
		}
	}

	return true
}

func calculateFirstVerifierDigit(cpf []int) int {
	var sum int
	for i, v := range cpf[:9] {
		sum += v * (i + 1)
	}

	result := sum % 11

	if result == 10 {
		return 0
	}

	return result
}

func calculateSecondVerifierDigit(cpf []int, firstVerifierDigit int) int {
	var sum int
	for i, v := range cpf[:9] {
		sum += v * i
	}

	sum += (firstVerifierDigit * 9)

	result := sum % 11

	if result == 10 {
		return 0
	}

	return result
}
