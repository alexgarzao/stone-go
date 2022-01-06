package documents

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/stone-payments/stone-go/formatter"
)

const (
	minValueWeight         = 2
	initialValueWeightCNPJ = 9
)

// Default errors message.
var (
	ErrInvalidCPF  = errors.New("invalid CPF")
	ErrInvalidCNPJ = errors.New("invalid CNPJ")
)

//ValidateCPF - Checks if the document informed is valid CPF.
func ValidateCPF(doc string) error {
	const (
		firstIndexDigitVerification = 9
		initialWeight               = 10
	)

	if !valid(doc, firstIndexDigitVerification, initialWeight) {
		return ErrInvalidCPF
	}

	return nil
}

//ValidateCNPJ - Checks if the document informed is valid CNPJ.
func ValidateCNPJ(doc string) error {
	const (
		firstIndexDigitVerification = 12
		initialWeight               = 5
	)

	if !valid(doc, firstIndexDigitVerification, initialWeight) {
		return ErrInvalidCNPJ
	}

	return nil
}

//validate - Checks if the document informed is valid, removing special characters and calculating
// the check digits, according to the rules for CPF and CNPJ.
//
// The informed document must contain all the characters:
// 		E.g: CPF - 11 digits, CNPJ - 14 digits
// 		Any informed number containing a different size of these will be returned as invalid.
//
// based on: https://github.com/Nhanderu/brdoc
func valid(doc string, firstIndexDigitVerification int, initialWeight int) bool {
	// Removes non-numeric characters to validate the positions of each digit.
	docNormalized := formatter.OnlyNumbers(doc)

	// When the document is empty
	// or the size is smaller than the first digit verification index, it is not valid.
	if len(docNormalized) <= 0 || len(docNormalized) < firstIndexDigitVerification {
		return false
	}

	// Document is not valid when all digits are equal.
	if allEquals(docNormalized) {
		return false
	}

	// Calculates the first check digit.
	data := docNormalized[:firstIndexDigitVerification]
	digit := calculateDigit(data, initialWeight)

	// Calculates the second check digit.
	data = data + digit
	digit = calculateDigit(data, initialWeight+1)

	return docNormalized == data+digit
}

// allEquals - Checks if all digits are equal
// e.g: 111.111.111-11 or 22.222.222/2222-22
func allEquals(value string) bool {
	base := value[0]
	for i := 1; i < len(value); i++ {
		if base != value[i] {
			return false
		}
	}

	return true
}

// calculateDigit - Calculates the verification digit of the informed document according
// to its type, obtaining the base 11 verification digit.
//
// value - document digits
// initialWeight - represents the weight for the check digit calculation rule.
//
// Rule of CPF weight: 10, 9, 8, 7, 6, 5, 4, 3, 2
// Rule of CNPJ weight: 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2
func calculateDigit(value string, initialWeight int) string {
	var sum int
	for _, num := range value {
		digit := int(num - '0')
		sum += digit * initialWeight
		initialWeight--

		// Conditional to satisfy the weight rule for CNPJ type documents
		if initialWeight < minValueWeight {
			initialWeight = initialValueWeightCNPJ
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}

// GenerateCPF generate a valid random cpf with only the numerical digits
func GenerateCPF() string {
	const firstIndexDigitVerification = 9

	rand.Seed(time.Now().UTC().UnixNano())
	cpf := rand.Perm(firstIndexDigitVerification)
	cpf = append(cpf, generateVerificationDigit(cpf, len(cpf)))
	cpf = append(cpf, generateVerificationDigit(cpf, len(cpf)))

	var cpfString string
	for _, c := range cpf {
		cpfString += strconv.Itoa(c)
	}
	return cpfString
}

// GenerateCPFFormatted generate a valid random formatted cpf
func GenerateCPFFormatted() (cpf string) {
	cpfFormatted, _ := format(GenerateCPF())
	return cpfFormatted
}

// generateVerificationDigit generate the cpf verification digit
func generateVerificationDigit(data []int, n int) int {
	var total int

	for i := 0; i < n; i++ {
		total += data[i] * (n + 1 - i)
	}

	total = total % 11
	if total < 2 {
		return 0
	}
	return 11 - total
}
