package r18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatNumber(t *testing.T) {
	assert.Equal(t, "R$ 123,99", FormatNumber("R$ {.,2}", 12399))
	assert.Equal(t, "R$ 123.000,99", FormatNumber("R$ {.,2}", 12300099))
	assert.Equal(t, "R$ 123.000,987", FormatNumber("R$ {.,3}", 123000987))
	assert.Equal(t, "R$ 123.000", FormatNumber("R$ ç {.}", 123000))
	assert.Equal(t, "R$ {}", FormatNumber("R$ \\{\\}", 0))
}
