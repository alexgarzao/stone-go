package documents

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCPF(t *testing.T) {
	validCPF := "841.494.050-18"
	onlyDigitsCPF := "84149405018"

	tests := []struct {
		name    string
		args    string
		want    CPF
		wantErr bool
	}{
		{
			name:    "formatted cpf keeps formatting after creation",
			args:    validCPF,
			want:    CPF{value: validCPF},
			wantErr: false,
		},
		{
			name:    "cpf containing only digits gets formatted correctly",
			args:    onlyDigitsCPF,
			want:    CPF{value: validCPF},
			wantErr: false,
		},
		{
			name:    "invalid CPF returns an err",
			args:    "12345 5032",
			want:    CPF{},
			wantErr: true,
		},
		{
			name:    "empty string returns an err",
			args:    "",
			want:    CPF{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCPF(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCPF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("NewCPF() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCPFMarshalling(t *testing.T) {
	validCPF := "841.494.050-18"
	type pessoa struct {
		Nome string `json:"nome"`
		CPF  CPF    `json:"cpf"`
	}
	t.Run("marshalling", func(t *testing.T) {
		cpf, _ := NewCPF(validCPF)
		p := pessoa{
			Nome: "Daskommunistischemanifest",
			CPF:  cpf,
		}

		m, err := json.Marshal(p)
		assert.Nil(t, err)
		assert.Equal(t, `{"nome":"Daskommunistischemanifest","cpf":"841.494.050-18"}`, string(m))
	})
	t.Run("unmarshalling", func(t *testing.T) {
		t.Run("valid formatted cpf", func(t *testing.T) {
			pJson := `{"nome":"Daskommunistischemanifest","cpf":"841.494.050-18"}`
			var p pessoa
			err := json.Unmarshal([]byte(pJson), &p)
			assert.Nil(t, err)
			assert.Equal(t, "Daskommunistischemanifest", p.Nome)
			assert.Equal(t, "841.494.050-18", p.CPF.String())
		})
		t.Run("valid unformatted cpf", func(t *testing.T) {
			pJson := `{"nome":"Daskommunistischemanifest","cpf":"84149405018"}`
			var p pessoa
			err := json.Unmarshal([]byte(pJson), &p)
			assert.Nil(t, err)
			assert.Equal(t, "Daskommunistischemanifest", p.Nome)
			assert.Equal(t, "841.494.050-18", p.CPF.String())
		})
		t.Run("invalid unformatted cpf", func(t *testing.T) {
			pJson := `{"nome":"Daskommunistischemanifest","cpf":"841"}`
			var p pessoa
			err := json.Unmarshal([]byte(pJson), &p)
			assert.NotNil(t, err)
			assert.Equal(t, ErrInvalidCPF, err)
		})
	})
}

func TestCPF_DigitsOnly(t *testing.T) {
	tests := []struct {
		name string
		cpf  string
		want string
	}{
		{
			name: "DigitsOnly returns exactly the digits of the cpf",
			cpf:  "023.999.023-99",
			want: "02399902399",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CPF{
				value: tt.cpf,
			}
			if got := c.DigitsOnly(); got != tt.want {
				t.Errorf("DigitsOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}
