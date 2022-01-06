package documents

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessValidateCPF(t *testing.T) {
	successTests := []struct {
		name string
		in   string
	}{
		{"Document type CPF valid formated ", "883.500.570-17"},
		{"Document type CPF valid without format", "80545919002"},
		{"Document type CPF valid with padding", " 93388834008  "},
	}

	// Apply scenarios successful tests.
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCPF(tt.in)
			assert.Nil(t, err)
		})
	}
}

func TestFailScenariosValidateCPF(t *testing.T) {
	failTests := []struct {
		name string
		in   string
	}{
		{"Empty value ", ""},
		{"Document type CPF invalid size", "601"},
		{"Document type CPF invalid size", "714.330.560"},
		{"Document type CPF all equals", "111.111.111-11"},
		{"Document type CPF invalid DV, correct is 03 ", "714.330.560-99"},
		{"Document type CPF invalid DV, correct is 40", "60140404049"},
		{"Document type CPF invalid DV, with padding, correct is 12", " 68167355088  "},
	}

	for _, tt := range failTests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCPF(tt.in)
			assert.True(t, errors.Is(err, ErrInvalidCPF))
		})
	}
}

func TestSuccessValidateCNPJ(t *testing.T) {
	successTests := []struct {
		name string
		in   string
	}{
		{"Document type CNPJ valid formated ", "19.783.246/0001-05"},
		{"Document type CNPJ valid without format", "95040409000148"},
		{"Document type CNPJ valid with padding", "  26053781000176  "},
	}

	// Apply scenarios successful tests.
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCNPJ(tt.in)
			assert.Nil(t, err)
		})
	}
}

func TestFailScenariosValidateCNPJ(t *testing.T) {
	failTests := []struct {
		name string
		in   string
	}{
		{"Empty value ", ""},
		{"Document type CNPJ invalid size", "24.2"},
		{"Document type CNPJ invalid size", "24.247.999/0001"},
		{"Document type CNPJ all equals", "22.222.222/2222-22"},
		{"Document type CNPJ invalid DV, correct is 36 ", "24.247.999/0001-99"},
		{"Document type CNPJ invalid DV, correct is 08", "01941097000199"},
		{"Document type CNPJ invalid DV, with padding, correct is 86", " 65728975000109  "},
	}

	for _, tt := range failTests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCNPJ(tt.in)
			assert.True(t, errors.Is(err, ErrInvalidCNPJ))
		})
	}
}

func TestSuccessGenerateCPF(t *testing.T) {
	successTests := []struct {
		name string
		in   string
	}{
		{"Generates a valid cpf with only the numerical digits", GenerateCPF()},
		{"Generates a valid cpf formatted", GenerateCPFFormatted()},
	}

	// Apply scenarios successful tests.
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCPF(tt.in)
			assert.Nil(t, err)
		})
	}
}
