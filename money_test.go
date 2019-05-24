package r18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatMoney(t *testing.T) {
	assert.Equal(t, "R$ 123.000,99", FormatMoney("R$ {.,2}", 12300099))
	assert.Equal(t, "R$ 123.000,987", FormatMoney("R$ {.,3}", 123000987))
	assert.Equal(t, "R$ 123.000", FormatMoney("R$ {.}", 123000))
}
