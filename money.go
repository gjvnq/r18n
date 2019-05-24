package r18n

import (
	"fmt"
	"strconv"
	"strings"
)

// Format: {[thousand separator][decimal separator][digits after separator]}
// Ex:  {.} 1.000
// Ex:  {.,2} 1.000,00
// Ex:  {,.2} 1,000.00
func FormatNumber(fmt string, amount int) string {
	ans := strings.Builder{}
	tmp := strings.Builder{}
	mode := 0 // 0 - copying | 1 - escape | 2 - processing | 3 - finalizing
	for i, c := range fmt {
		println(mode, i, string(c))
		switch mode {
		case 0:
			switch c {
			case '\\':
				mode = 1
			case '{':
				tmp.Reset()
				mode = 2
			default:
				ans.WriteRune(c)
			}
		case 1:
			ans.WriteRune(c)
			mode = 0
		case 2:
			switch c {
			case '}':
				mode = 3
			default:
				tmp.WriteRune(c)
			}
		case 3:
			mode = 0
			ans.WriteString(actualFormatNumber(tmp.String(), amount))
		default:
			panic("unexpected mode in FormatMoney")
		}
	}
	if mode == 3 {
		ans.WriteString(actualFormatNumber(tmp.String(), amount))
	}
	println()

	return ans.String()
}

func intPow(base, exp int) int {
	ans := 1
	for i := 0; i < exp; i++ {
		ans *= base
	}
	return ans
}

func actualFormatNumber(fmt_ string, amount int) string {
	sep_1000 := fmt_[0]
	sep_dec := fmt_[1]
	digits, err := strconv.Atoi(fmt_[2:])
	if err != nil {
		panic(err)
	}

	println(">", sep_1000, sep_dec, digits)

	int_part := amount / intPow(10, digits)
	frac_part := amount % intPow(10, digits)

	println(">", int_part, frac_part)
	ans := fmt.Sprintf("%d%c%d", int_part, sep_dec, frac_part)
	println(">", ans)
	return ans
}
