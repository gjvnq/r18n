package r18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatNumber(t *testing.T) {
	assert.Equal(t, "R$ 123,99", FormatNumber("R$ {.,2}", 12399))
	assert.Equal(t, "R$ 123.000,99", FormatNumber("R$ {.,2}", 12300099))
	assert.Equal(t, "R$ 123.000,987", FormatNumber("R$ {.,3}", 123000987))
	assert.Equal(t, "R$ 123.000", FormatNumber("R$ {.}", 123000))
	assert.Equal(t, "R$ {}", FormatNumber("R$ \\{\\}", 0))
}

func Test_PT_NumberIntCardinal(t *testing.T) {
	assert.Equal(t, "zero", NumberIntCardinal(PT, GENDER_MALE, 0))
	assert.Equal(t, "um", NumberIntCardinal(PT, GENDER_MALE, 1))
	assert.Equal(t, "dois", NumberIntCardinal(PT, GENDER_MALE, 2))
	assert.Equal(t, "doze", NumberIntCardinal(PT, GENDER_MALE, 12))
	assert.Equal(t, "dezenove", NumberIntCardinal(PT, GENDER_MALE, 19))
	assert.Equal(t, "vinte", NumberIntCardinal(PT, GENDER_MALE, 20))
	assert.Equal(t, "vinte e um", NumberIntCardinal(PT, GENDER_MALE, 21))
	assert.Equal(t, "vinte e dois", NumberIntCardinal(PT, GENDER_MALE, 22))
	assert.Equal(t, "cem", NumberIntCardinal(PT, GENDER_MALE, 100))
	assert.Equal(t, "cento e vinte e um", NumberIntCardinal(PT, GENDER_MALE, 121))
	assert.Equal(t, "cento e vinte e dois", NumberIntCardinal(PT, GENDER_MALE, 122))
	assert.Equal(t, "mil cento e vinte e três", NumberIntCardinal(PT, GENDER_MALE, 1123))
	assert.Equal(t, "mil duzentos e trinta e quatro", NumberIntCardinal(PT, GENDER_MALE, 1234))
	assert.Equal(t, "dois mil duzentos e trinta e quatro", NumberIntCardinal(PT, GENDER_MALE, 2234))
	assert.Equal(t, "vinte e um mil duzentos e trinta e quatro", NumberIntCardinal(PT, GENDER_MALE, 21234))

	assert.Equal(t, "duas", NumberIntCardinal(PT, GENDER_FEMALE, 2))
	assert.Equal(t, "doux", NumberIntCardinal(PT, GENDER_NON_BINARY, 2))
	assert.Equal(t, "um negativo", NumberIntCardinal(PT, GENDER_MALE, -1))

	assert.Equal(t, "vinte e três", NumberIntCardinal(PT, GENDER_MALE, 23))
	assert.Equal(t, "cento e vinte e três", NumberIntCardinal(PT, GENDER_MALE, 123))
	assert.Equal(t, "mil duzentos e trinta e quatro", NumberIntCardinal(PT, GENDER_MALE, 1234))

	assert.Equal(t, "duzentos", NumberIntCardinal(PT, GENDER_MALE, 200))
	assert.Equal(t, "duzentas", NumberIntCardinal(PT, GENDER_FEMALE, 200))
	assert.Equal(t, "duzentxs", NumberIntCardinal(PT, GENDER_NON_BINARY, 200))
	assert.Equal(t, "três milhões quinhentos e vinte e sete mil duzentos e trinta e dois", NumberIntCardinal(PT, GENDER_MALE, 3527232))
	assert.Equal(t, "três milhões quinhentas e vinte e sete mil duzentas e trinta e duas", NumberIntCardinal(PT, GENDER_FEMALE, 3527232))
	assert.Equal(t, "três milhões quinhentxs e vinte e sete mil duzentxs e trinta e doux", NumberIntCardinal(PT, GENDER_NON_BINARY, 3527232))
}

func Test_EN_NumberIntCardinal(t *testing.T) {
	assert.Equal(t, "zero", NumberIntCardinal(EN, GENDER_MALE, 0))
	assert.Equal(t, "one", NumberIntCardinal(EN, GENDER_MALE, 1))
	assert.Equal(t, "two", NumberIntCardinal(EN, GENDER_MALE, 2))
	assert.Equal(t, "twelve", NumberIntCardinal(EN, GENDER_MALE, 12))
	assert.Equal(t, "nineteen", NumberIntCardinal(EN, GENDER_MALE, 19))
	assert.Equal(t, "twenty", NumberIntCardinal(EN, GENDER_MALE, 20))
	assert.Equal(t, "twenty one", NumberIntCardinal(EN, GENDER_MALE, 21))
	assert.Equal(t, "twenty two", NumberIntCardinal(EN, GENDER_MALE, 22))
	assert.Equal(t, "one hundred", NumberIntCardinal(EN, GENDER_MALE, 100))
	assert.Equal(t, "one hundred and twenty one", NumberIntCardinal(EN, GENDER_MALE, 121))
	assert.Equal(t, "one hundred and twenty two", NumberIntCardinal(EN, GENDER_MALE, 122))
	assert.Equal(t, "one thousand one hundred and twenty three", NumberIntCardinal(EN, GENDER_MALE, 1123))
	assert.Equal(t, "twenty one thousand two hundred and thirty four", NumberIntCardinal(EN, GENDER_MALE, 21234))

	assert.Equal(t, "two", NumberIntCardinal(EN, GENDER_FEMALE, 2))
	assert.Equal(t, "two", NumberIntCardinal(EN, GENDER_NON_BINARY, 2))
	assert.Equal(t, "negative one", NumberIntCardinal(EN, GENDER_MALE, -1))

	assert.Equal(t, "three million five hundred and twenty seven thousand two hundred and thirty two", NumberIntCardinal(EN, GENDER_MALE, 3527232))
}
