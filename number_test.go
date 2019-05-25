package r18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatNumber(t *testing.T) {
	assert.Equal(t, "R$ 123,99", FormatNumber("R$ {.,2}", 12399))
	assert.Equal(t, "R$ 123.000,99", FormatNumber("R$ {.,2}", 12300099))
	assert.Equal(t, "R$ 123.000,987", FormatNumber("R$ {.,3}", 123000987))
	assert.Equal(t, "R$ 123.000", FormatNumber("R$ {.}", 123000))
	assert.Equal(t, "R$ {}", FormatNumber("R$ \\{\\}", 0))
}

func Test_ptNumberIntCardinal(t *testing.T) {
	require.Equal(t, "zero", ptNumberIntCardinal(GENDER_MALE, 0))
	require.Equal(t, "um", ptNumberIntCardinal(GENDER_MALE, 1))
	require.Equal(t, "dois", ptNumberIntCardinal(GENDER_MALE, 2))
	require.Equal(t, "doze", ptNumberIntCardinal(GENDER_MALE, 12))
	require.Equal(t, "dezenove", ptNumberIntCardinal(GENDER_MALE, 19))
	require.Equal(t, "vinte", ptNumberIntCardinal(GENDER_MALE, 20))
	require.Equal(t, "vinte e um", ptNumberIntCardinal(GENDER_MALE, 21))
	require.Equal(t, "vinte e dois", ptNumberIntCardinal(GENDER_MALE, 22))
	require.Equal(t, "cem", ptNumberIntCardinal(GENDER_MALE, 100))
	require.Equal(t, "cento e vinte e um", ptNumberIntCardinal(GENDER_MALE, 121))
	require.Equal(t, "cento e vinte e dois", ptNumberIntCardinal(GENDER_MALE, 122))
	require.Equal(t, "mil cento e vinte e três", ptNumberIntCardinal(GENDER_MALE, 1123))
	require.Equal(t, "mil duzentos e trinta e quatro", ptNumberIntCardinal(GENDER_MALE, 1234))
	require.Equal(t, "dois mil duzentos e trinta e quatro", ptNumberIntCardinal(GENDER_MALE, 2234))
	require.Equal(t, "vinte e um mil duzentos e trinta e quatro", ptNumberIntCardinal(GENDER_MALE, 21234))

	assert.Equal(t, "duas", ptNumberIntCardinal(GENDER_FEMALE, 2))
	assert.Equal(t, "doux", ptNumberIntCardinal(GENDER_NON_BINARY, 2))
	assert.Equal(t, "um negativo", ptNumberIntCardinal(GENDER_MALE, -1))

	assert.Equal(t, "vinte e três", ptNumberIntCardinal(GENDER_MALE, 23))
	assert.Equal(t, "cento e vinte e três", ptNumberIntCardinal(GENDER_MALE, 123))
	assert.Equal(t, "mil duzentos e trinta e quatro", ptNumberIntCardinal(GENDER_MALE, 1234))

	assert.Equal(t, "duzentos", ptNumberIntCardinal(GENDER_MALE, 200))
	assert.Equal(t, "duzentas", ptNumberIntCardinal(GENDER_FEMALE, 200))
	assert.Equal(t, "duzentxs", ptNumberIntCardinal(GENDER_NON_BINARY, 200))
	assert.Equal(t, "três milhões quinhentos e vinte e sete mil duzentos e trinta e dois", ptNumberIntCardinal(GENDER_MALE, 3527232))
	assert.Equal(t, "três milhões quinhentas e vinte e sete mil duzentas e trinta e duas", ptNumberIntCardinal(GENDER_FEMALE, 3527232))
	assert.Equal(t, "três milhões quinhentxs e vinte e sete mil duzentxs e trinta e doux", ptNumberIntCardinal(GENDER_NON_BINARY, 3527232))
}
