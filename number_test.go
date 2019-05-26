package r18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_int2digits(t *testing.T) {
	for i := 0; i < 10; i++ {
		assert.Equal(t, []int{i}, int2digits(i))
	}
	for i := 10; i < 20; i++ {
		assert.Equal(t, []int{i - 10, 1}, int2digits(i))
	}
	for i := 100; i < 110; i++ {
		assert.Equal(t, []int{i - 100, 0, 1}, int2digits(i))
	}
}

func TestFormatNumber(t *testing.T) {
	assert.Equal(t, "R$ 123,99", _FormatNumber("R$ {.,2}", 12399))
	assert.Equal(t, "R$ 123.000,99", _FormatNumber("R$ {.,2}", 12300099))
	assert.Equal(t, "R$ 123.000,987", _FormatNumber("R$ {.,3}", 123000987))
	assert.Equal(t, "R$ 123.000", _FormatNumber("R$ {.}", 123000))
	assert.Equal(t, "R$ {}", _FormatNumber("R$ \\{\\}", 0))
}

func Test_PT_NumberIntCardinal(t *testing.T) {
	assert.Equal(t, "zero", _NumberIntCardinal(PT, GENDER_MALE, 0))
	assert.Equal(t, "um", _NumberIntCardinal(PT, GENDER_MALE, 1))
	assert.Equal(t, "dois", _NumberIntCardinal(PT, GENDER_MALE, 2))
	assert.Equal(t, "doze", _NumberIntCardinal(PT, GENDER_MALE, 12))
	assert.Equal(t, "dezenove", _NumberIntCardinal(PT, GENDER_MALE, 19))
	assert.Equal(t, "vinte", _NumberIntCardinal(PT, GENDER_MALE, 20))
	assert.Equal(t, "vinte e um", _NumberIntCardinal(PT, GENDER_MALE, 21))
	assert.Equal(t, "vinte e dois", _NumberIntCardinal(PT, GENDER_MALE, 22))
	assert.Equal(t, "cem", _NumberIntCardinal(PT, GENDER_MALE, 100))
	assert.Equal(t, "cento e vinte e um", _NumberIntCardinal(PT, GENDER_MALE, 121))
	assert.Equal(t, "cento e vinte e dois", _NumberIntCardinal(PT, GENDER_MALE, 122))
	assert.Equal(t, "mil cento e vinte e três", _NumberIntCardinal(PT, GENDER_MALE, 1123))
	assert.Equal(t, "mil duzentos e trinta e quatro", _NumberIntCardinal(PT, GENDER_MALE, 1234))
	assert.Equal(t, "dois mil duzentos e trinta e quatro", _NumberIntCardinal(PT, GENDER_MALE, 2234))
	assert.Equal(t, "vinte e um mil duzentos e trinta e quatro", _NumberIntCardinal(PT, GENDER_MALE, 21234))

	assert.Equal(t, "duas", _NumberIntCardinal(PT, GENDER_FEMALE, 2))
	assert.Equal(t, "doux", _NumberIntCardinal(PT, GENDER_NON_BINARY, 2))
	assert.Equal(t, "um negativo", _NumberIntCardinal(PT, GENDER_MALE, -1))

	assert.Equal(t, "vinte e três", _NumberIntCardinal(PT, GENDER_MALE, 23))
	assert.Equal(t, "cento e vinte e três", _NumberIntCardinal(PT, GENDER_MALE, 123))
	assert.Equal(t, "mil duzentos e trinta e quatro", _NumberIntCardinal(PT, GENDER_MALE, 1234))

	assert.Equal(t, "duzentos", _NumberIntCardinal(PT, GENDER_MALE, 200))
	assert.Equal(t, "duzentas", _NumberIntCardinal(PT, GENDER_FEMALE, 200))
	assert.Equal(t, "duzentxs", _NumberIntCardinal(PT, GENDER_NON_BINARY, 200))
	assert.Equal(t, "cem milhões", _NumberIntCardinal(PT, GENDER_NON_BINARY, 100000000))
	assert.Equal(t, "três milhões quinhentos e vinte e sete mil duzentos e trinta e dois", _NumberIntCardinal(PT, GENDER_MALE, 3527232))
	assert.Equal(t, "três milhões quinhentas e vinte e sete mil duzentas e trinta e duas", _NumberIntCardinal(PT, GENDER_FEMALE, 3527232))
	assert.Equal(t, "três milhões quinhentxs e vinte e sete mil duzentxs e trinta e doux", _NumberIntCardinal(PT, GENDER_NON_BINARY, 3527232))

	assert.Equal(t, "cento e três milhões quinhentas e vinte e sete mil duzentas e trinta e duas", _NumberIntCardinal(PT, GENDER_FEMALE, 103527232))

	assert.Equal(t, "mil", _NumberIntCardinal(PT, GENDER_FEMALE, 1000))
	assert.Equal(t, "um milhão", _NumberIntCardinal(PT, GENDER_FEMALE, 1000000))
	assert.Equal(t, "um bilhão", _NumberIntCardinal(PT, GENDER_FEMALE, 1000000000))
	assert.Equal(t, "cem bilhões", _NumberIntCardinal(PT, GENDER_FEMALE, 100000000000))
	assert.Equal(t, "um trilhão", _NumberIntCardinal(PT, GENDER_FEMALE, 1000000000000))
	assert.Equal(t, "cem trilhões", _NumberIntCardinal(PT, GENDER_FEMALE, 100000000000000))
	assert.Equal(t, "um quadrilhão", _NumberIntCardinal(PT, GENDER_FEMALE, 1000000000000000))
	assert.Equal(t, "cem quadrilhões", _NumberIntCardinal(PT, GENDER_FEMALE, 100000000000000000))
	// assert.Equal(t, "um quintilhão", _NumberIntCardinal(PT, GENDER_FEMALE, 1000000000000000000))
	assert.Equal(t, "um zero zero zero zero zero zero zero zero zero zero zero zero zero zero zero zero zero zero", _NumberIntCardinal(PT, GENDER_FEMALE, 1000000000000000000))
}

func Test_PT_NumberFloatCardinal(t *testing.T) {
	assert.Equal(t, "zero", _NumberFloatCardinal(PT, GENDER_FEMALE, 0))
	assert.Equal(t, "doze", _NumberFloatCardinal(PT, GENDER_FEMALE, 12))
	assert.Equal(t, "um décimo", _NumberFloatCardinal(PT, GENDER_FEMALE, 0.1))
	assert.Equal(t, "um centavo", _NumberFloatCardinal(PT, GENDER_FEMALE, 0.01))
	assert.Equal(t, "um milésimo", _NumberFloatCardinal(PT, GENDER_FEMALE, 0.001))
	assert.Equal(t, "dois décimos", _NumberFloatCardinal(PT, GENDER_FEMALE, 0.2))
	assert.Equal(t, "doze centavos", _NumberFloatCardinal(PT, GENDER_FEMALE, 0.12))
	assert.Equal(t, "cento e vinte e três milésimos", _NumberFloatCardinal(PT, GENDER_FEMALE, 0.123))
	assert.Equal(t, "zero separador decimal um dois três quatro", _NumberFloatCardinal(PT, GENDER_FEMALE, 0.1234))
	assert.Equal(t, "três separador decimal um quatro um cinco nove dois seis cinco três cinco nove", _NumberFloatCardinal(PT, GENDER_FEMALE, 3.14159265359))
}

func Test_EN_NumberFloatCardinal(t *testing.T) {
	assert.Equal(t, "zero", _NumberFloatCardinal(EN, GENDER_FEMALE, 0))
	assert.Equal(t, "twelve", _NumberFloatCardinal(EN, GENDER_FEMALE, 12))
	assert.Equal(t, "one tenth", _NumberFloatCardinal(EN, GENDER_FEMALE, 0.1))
	assert.Equal(t, "one hundredth", _NumberFloatCardinal(EN, GENDER_FEMALE, 0.01))
	assert.Equal(t, "one thousandth", _NumberFloatCardinal(EN, GENDER_FEMALE, 0.001))
	assert.Equal(t, "two tenths", _NumberFloatCardinal(EN, GENDER_FEMALE, 0.2))
	assert.Equal(t, "twelve hundredths", _NumberFloatCardinal(EN, GENDER_FEMALE, 0.12))
	assert.Equal(t, "one hundred and twenty three thousandths", _NumberFloatCardinal(EN, GENDER_FEMALE, 0.123))
	assert.Equal(t, "zero decimal separator one two three four", _NumberFloatCardinal(EN, GENDER_FEMALE, 0.1234))
	assert.Equal(t, "three decimal separator one four one five nine two six five three five nine", _NumberFloatCardinal(EN, GENDER_FEMALE, 3.14159265359))
}

func Test_EN_NumberIntCardinal(t *testing.T) {
	assert.Equal(t, "zero", _NumberIntCardinal(EN, GENDER_MALE, 0))
	assert.Equal(t, "one", _NumberIntCardinal(EN, GENDER_MALE, 1))
	assert.Equal(t, "two", _NumberIntCardinal(EN, GENDER_MALE, 2))
	assert.Equal(t, "twelve", _NumberIntCardinal(EN, GENDER_MALE, 12))
	assert.Equal(t, "nineteen", _NumberIntCardinal(EN, GENDER_MALE, 19))
	assert.Equal(t, "twenty", _NumberIntCardinal(EN, GENDER_MALE, 20))
	assert.Equal(t, "twenty one", _NumberIntCardinal(EN, GENDER_MALE, 21))
	assert.Equal(t, "twenty two", _NumberIntCardinal(EN, GENDER_MALE, 22))
	assert.Equal(t, "one hundred", _NumberIntCardinal(EN, GENDER_MALE, 100))
	assert.Equal(t, "one hundred and twenty one", _NumberIntCardinal(EN, GENDER_MALE, 121))
	assert.Equal(t, "one hundred and twenty two", _NumberIntCardinal(EN, GENDER_MALE, 122))
	assert.Equal(t, "one thousand one hundred and twenty three", _NumberIntCardinal(EN, GENDER_MALE, 1123))
	assert.Equal(t, "twenty one thousand two hundred and thirty four", _NumberIntCardinal(EN, GENDER_MALE, 21234))

	assert.Equal(t, "two", _NumberIntCardinal(EN, GENDER_FEMALE, 2))
	assert.Equal(t, "two", _NumberIntCardinal(EN, GENDER_NON_BINARY, 2))
	assert.Equal(t, "negative one", _NumberIntCardinal(EN, GENDER_MALE, -1))

	assert.Equal(t, "three million five hundred and twenty seven thousand two hundred and thirty two", _NumberIntCardinal(EN, GENDER_MALE, 3527232))
}

func Test_PT_NumberOrdinal(t *testing.T) {
	assert.Equal(t, "zerésima", _NumberOrdinal(PT, GENDER_FEMALE, 0))
	assert.Equal(t, "primeira", _NumberOrdinal(PT, GENDER_FEMALE, 1))
	assert.Equal(t, "segunda", _NumberOrdinal(PT, GENDER_FEMALE, 2))
	assert.Equal(t, "terceira", _NumberOrdinal(PT, GENDER_FEMALE, 3))
	assert.Equal(t, "quarta", _NumberOrdinal(PT, GENDER_FEMALE, 4))
	assert.Equal(t, "quinta", _NumberOrdinal(PT, GENDER_FEMALE, 5))
	assert.Equal(t, "sexta", _NumberOrdinal(PT, GENDER_FEMALE, 6))
	assert.Equal(t, "sétima", _NumberOrdinal(PT, GENDER_FEMALE, 7))
	assert.Equal(t, "oitava", _NumberOrdinal(PT, GENDER_FEMALE, 8))
	assert.Equal(t, "nona", _NumberOrdinal(PT, GENDER_FEMALE, 9))
	assert.Equal(t, "décima", _NumberOrdinal(PT, GENDER_FEMALE, 10))
	assert.Equal(t, "décima primeira", _NumberOrdinal(PT, GENDER_FEMALE, 11))
	assert.Equal(t, "décima segunda", _NumberOrdinal(PT, GENDER_FEMALE, 12))
	assert.Equal(t, "décima terceira", _NumberOrdinal(PT, GENDER_FEMALE, 13))
	assert.Equal(t, "vigésima", _NumberOrdinal(PT, GENDER_FEMALE, 20))
	assert.Equal(t, "trigésima", _NumberOrdinal(PT, GENDER_FEMALE, 30))
	assert.Equal(t, "quadragésimo", _NumberOrdinal(PT, GENDER_MALE, 40))
	assert.Equal(t, "quinquagésimo", _NumberOrdinal(PT, GENDER_MALE, 50))
	assert.Equal(t, "sexagésimo", _NumberOrdinal(PT, GENDER_MALE, 60))
	assert.Equal(t, "septuagésimo", _NumberOrdinal(PT, GENDER_MALE, 70))
	assert.Equal(t, "octogésimo", _NumberOrdinal(PT, GENDER_MALE, 80))
	assert.Equal(t, "nonagésimo", _NumberOrdinal(PT, GENDER_MALE, 90))
	assert.Equal(t, "centésimo", _NumberOrdinal(PT, GENDER_MALE, 100))
	assert.Equal(t, "centésimx vigésimx terceirx", _NumberOrdinal(PT, GENDER_NON_BINARY, 123))
	assert.Equal(t, "ducentésimo", _NumberOrdinal(PT, GENDER_MALE, 200))
	assert.Equal(t, "tricentésimo", _NumberOrdinal(PT, GENDER_MALE, 300))
	assert.Equal(t, "quadringentésimo", _NumberOrdinal(PT, GENDER_MALE, 400))
	assert.Equal(t, "quingentésimo", _NumberOrdinal(PT, GENDER_MALE, 500))
	assert.Equal(t, "sexcentésimo", _NumberOrdinal(PT, GENDER_MALE, 600))
	assert.Equal(t, "septingentésimo", _NumberOrdinal(PT, GENDER_MALE, 700))
	assert.Equal(t, "octingentésimo", _NumberOrdinal(PT, GENDER_MALE, 800))
	assert.Equal(t, "noningentésimo", _NumberOrdinal(PT, GENDER_MALE, 900))
	assert.Equal(t, "milésimo", _NumberOrdinal(PT, GENDER_MALE, 1000))
	assert.Equal(t, "décimo milésimo", _NumberOrdinal(PT, GENDER_MALE, 10000))
	assert.Equal(t, "centésimo milésimo", _NumberOrdinal(PT, GENDER_MALE, 100000))
	assert.Equal(t, "milionésimo", _NumberOrdinal(PT, GENDER_MALE, 1000000))
	assert.Equal(t, "bilionésimo", _NumberOrdinal(PT, GENDER_MALE, 1000000000))
	assert.Equal(t, "trilionésimo", _NumberOrdinal(PT, GENDER_MALE, 1000000000000))
}
