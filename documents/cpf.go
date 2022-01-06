package documents

import (
	"encoding/json"
	"fmt"
	"regexp"
)

var (
	regexFormattedCpf  = regexp.MustCompile(`^(\d{3})\.(\d{3})\.(\d{3})-(\d{2})$`)
	regexOnlyDigitsCpf = regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)
)

type CPF struct {
	value string
}

// NewCPF creates a CPF type, asserting that it matches a valid CPF, with or without formatting.
// This means that the cpf parameter sent to this constructor must match either \d{11} or \d{3}.\d{3}.\d{3}-\d{2}.
// Note: it still does not verify if the CPF is valid in any way besides it's length.
func NewCPF(cpf string) (CPF, error) {
	cpf, err := format(cpf)
	if err != nil {
		return CPF{}, err
	}
	if err := ValidateCPF(cpf); err != nil {
		return CPF{}, fmt.Errorf("%w: cpf %s", err, cpf)
	}
	return CPF{value: cpf}, nil
}

// format ensures that the state of the cpf string is correctly formatted
// if better performance becomes a problem in the future we can skip the regexes and use a for-solution istead
func format(cpf string) (string, error) {
	onlyDigitsCPF := regexOnlyDigitsCpf.MatchString(cpf)
	if !regexFormattedCpf.MatchString(cpf) && !onlyDigitsCPF {
		return "", ErrInvalidCPF
	}
	if onlyDigitsCPF {
		cpf = regexOnlyDigitsCpf.ReplaceAllString(cpf, "$1.$2.$3-$4")
	}
	return cpf, nil
}

// String returns the formatted string representation of the CPF
func (c CPF) String() string {
	return c.value
}

// Formatted returns the formatted string representation of the CPF
func (c CPF) Formatted() string {
	return c.value
}

// DigitsOnly returns only the numerical digits the of the CPF
func (c CPF) DigitsOnly() string {
	// as internal state cannot be changed from the outside world it is safe to access the slice by index
	return c.value[:3] + c.value[4:7] + c.value[8:11] + c.value[12:]
}

func (c *CPF) UnmarshalJSON(bytes []byte) error {
	var cpf string
	if err := json.Unmarshal(bytes, &cpf); err != nil {
		return fmt.Errorf(`failed to unmarshal CPF: %w`, err)
	}
	cpf, err := format(cpf)
	if err != nil {
		return err
	}
	c.value = cpf
	return nil
}

func (c CPF) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(c.value)
	if err != nil {
		return nil, fmt.Errorf(`faield to marshal CPF: %w`, err)
	}
	return b, nil
}
