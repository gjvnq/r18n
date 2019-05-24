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
	for _, c := range fmt {
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

	int_part := []rune(strconv.Itoa(amount / intPow(10, digits)))
	frac_part := strconv.Itoa(amount % intPow(10, digits))

	int_part2 := make([]rune, 0)
	for i := len(int_part) - 1; i >= 0; i-- {
		int_part2 = append(int_part2, int_part[i])
		if i%3 == 0 {
			int_part2 = append(int_part2, sep_1000)
		}
	}

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

var ptCardinalsSmall map[int]string = map[int]string{
	0:  "zero",
	1:  "um",
	2:  "dois",
	3:  "três",
	4:  "quatro",
	5:  "cinco",
	6:  "seis",
	7:  "sete",
	8:  "oito",
	9:  "nove",
	10: "dez",
	11: "onze",
	12: "doze",
	13: "treze",
	14: "catorze",
	15: "quinze",
	16: "dezesseis",
	17: "dezessete",
	18: "dezoito",
	19: "dezenove",
}

var ptCardinalsTens map[int]string = map[int]string{
	20: "vinte",
	30: "trinta",
	40: "quarenta",
	50: "cinquenta",
	60: "sessenta",
	70: "setenta",
	80: "oitenta",
	90: "noventa",
}

var ptCardinalsHundreds map[int]string = map[int]string{
	100: "cem",
	200: "duzent@s",
	300: "trezent@s",
	400: "quatrocent@s",
	500: "quinhent@s",
	600: "seiscent@s",
	700: "setecent@s",
	800: "oitocent@s",
	900: "novecent@s",
}

var ptCardinalsScale map[int]string = map[int]string{
	3:  "mil",
	6:  "milhão",
	9:  "bilhão",
	12: "trilhão",
	15: "quadrilhão",
}

func ptNumberIntCardinal(gender string, val int) string {
	ans := ""
	negative := false

	if val < 0 {
		val *= -1
		negative = true
	}
	// Special case
	if val == 100 {
		return "cem"
	}

	// Process gender
	gender_marker := "o"
	if gender == GENDER_FEMALE {
		gender_marker = "a"
		ans = strings.Replace(ans, "dois", "duas", -1)
	}
	if gender == GENDER_NON_BINARY {
		gender_marker = "x"
		ans = strings.Replace(ans, "dois", "doux", -1)
	}

	// I should read http://www.blackwasp.co.uk/NumberToWords.aspx

	// General case (do largest to smallest)
	for val > 0 {
		if ans != "" {
			ans += " e "
		}
		if got, ok := ptCardinals[val]; ok {
			ans += got
			break
		}
		best := 0
		for _, offer := range ptCardinalsOrder {
			if best < val && offer <= val {
				best = offer
			}
		}
		ans += ptCardinals[best]
		val -= best
	}

	// Remebember to tell is the number was negative
	if negative {
		ans += " negativ@"
		if val > 1 {
			ans += "s"
		}
	}
	// Fix gender
	ans = strings.Replace(ans, "@", gender_marker, -1)
	return ans
}
