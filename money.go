package r18n

import (
	"strconv"
	"strings"
	"unicode/utf8"
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
			actualFormatNumber(tmp.String(), amount, &ans)
		default:
			panic("unexpected mode in FormatMoney")
		}
	}
	if mode == 3 {
		actualFormatNumber(tmp.String(), amount, &ans)
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

func actualFormatNumber(fmt_ string, amount int, ans *strings.Builder) {
	sep_1000 := '\''
	sep_dec := utf8.RuneError
	var err error
	digits := 0

	fmt2 := []rune(fmt_)
	sep_1000 = fmt2[0]
	if len(fmt_) >= 3 {
		sep_dec = fmt2[1]
		digits, err = strconv.Atoi(fmt_[2:])
		if err != nil {
			panic(err)
		}
	}

	println(">", sep_1000, sep_dec, digits)

	int_part := []rune(strconv.Itoa(amount / intPow(10, digits)))
	println("int_part=", string(int_part))
	frac_part := strconv.Itoa(amount % intPow(10, digits))
	println("frac_part=", string(frac_part))

	int_part2 := make([]rune, 0)
	for i := len(int_part) - 1; i >= 0; i-- {
		int_part2 = append(int_part2, int_part[i])
		if i%3 == 0 {
			int_part2 = append(int_part2, sep_1000)
		}
	}
	println(string(int_part2))

	for i := len(int_part2) - 1; i >= 0; i-- {
		r := int_part2[i]
		if r == '.' && i == len(int_part2)-1 {
			continue
		}
		ans.WriteRune(r)
	}

	if sep_dec != utf8.RuneError {
		ans.WriteRune(sep_dec)
		ans.WriteString(frac_part)
	}
}
