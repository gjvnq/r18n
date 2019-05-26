package r18n

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	_FLOAT_EPSILON = 1E-5
)

/* Format: {[thousand separator][decimal separator][digits after separator]}
Ex:  {.} 1.000
Ex:  {.,2} 1.000,00
Ex:  {,.2} 1,000.00 */
func _FormatNumber(fmt string, amount int) string {
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

func getBestNum(val *int, m map[int]string) string {
	best_k := -1
	for k, _ := range m {
		if best_k < k && k <= *val {
			best_k = k
		}
	}
	if best_k < 0 {
		return ""
	}
	*val -= best_k
	return m[best_k]
}

func inMap(val string, m map[int]string) bool {
	for _, v := range m {
		if v == val {
			return true
		}
	}
	return false
}

// gives in reverse order. Ex: 123 -> []int{3 2 1}
func int2digits(val int) []int {
	ans := make([]int, 0)

	// Special case
	if val == 0 {
		return []int{0}
	}
	// Ignore sign
	if val < 0 {
		val *= -1
	}

	// Do it!
	for val > 0 {
		ans = append(ans, val%10)
		val /= 10
	}

	return ans
}

func float2parts(val float64) (int, int, int) {
	if val < 0 {
		val *= -1
	}

	int_part := math.Trunc(val)
	dec_part := val - int_part

	// Find "exponent"
	exp := 0
	if dec_part < _FLOAT_EPSILON {
		return int(int_part), exp, 0
	}
	for ; dec_part-math.Trunc(dec_part) > _FLOAT_EPSILON; exp-- {
		dec_part *= 10
	}

	return int(int_part), exp, int(dec_part)
}

// Warning: this only goes
func _NumberFloatCardinal(language string, gender string, val float64) string {
	if val == 0 {
		return "zero"
	}

	int_part, exp, decimal_part := float2parts(val)
	int_part_str := ""
	if int_part != 0 {
		int_part_str = _NumberIntCardinal(language, gender, int_part)
	}

	if exp == 0 {
		return int_part_str
	}
	ans := int_part_str

	if exp >= -3 {
		cardinalsScale := enCardinalsScale
		if language == PT {
			cardinalsScale = ptCardinalsScale
			gender = GENDER_MALE
		}
		if language == PT && decimal_part > 1 {
			cardinalsScale = ptCardinalsScalePlural
		}
		if language == EN && decimal_part > 1 {
			cardinalsScale = enCardinalsScalePlural
		}

		if ans != "" {
			if language == PT {
				ans += " " + ptAnd + " "
			} else {
				ans += " " + enAnd + " "
			}
		}
		ans += _NumberIntCardinal(language, gender, decimal_part)
		ans += " "
		ans += cardinalsScale[exp]
		return ans
	}

	if ans == "" {
		ans += "zero"
	}

	if language == PT {
		ans += " separador decimal "
	} else {
		ans += " decimal separator "
	}
	ans += _NumberIntCardinalInDigits(language, gender, decimal_part)
	return ans
}

func digits2triadseq(digits []int) []int {
	for len(digits)%3 != 0 {
		digits = append(digits, 0)
	}

	triads := make([]int, 0)
	for i := 0; i < len(digits); i += 3 {
		triad := digits[i] + 10*digits[i+1] + 100*digits[i+2]
		triads = append(triads, triad)
	}

	// Reverse
	for i := 0; i < len(triads)/2; i++ {
		a := triads[i]
		b := triads[len(triads)-i-1]
		triads[i] = b
		triads[len(triads)-i-1] = a
	}

	return triads
}

func _NumberIntCardinalInDigits(language string, gender string, val int) string {
	var cardinalsSmall map[int]string
	if language == PT {
		cardinalsSmall = ptCardinalsSmall
	} else if language == EN {
		cardinalsSmall = enCardinalsSmall
	} else {
		panic(errors.New("unknown language: " + language))
	}

	digits := int2digits(val)
	words := make([]string, 0)
	for i := len(digits) - 1; i >= 0; i-- {
		w := cardinalsSmall[digits[i]]
		if digits[i] == 0 {
			w = "zero"
		}
		words = append(words, w)
	}

	return strings.Join(words, " ")
}

func _NumberIntCardinal(language string, gender string, val int) string {
	var cardinalsScale, cardinalsScalePlural, cardinalsHundreds, cardinalsTens, cardinalsSmall map[int]string
	var cardinalsScaleSet map[string]bool
	var and string
	words := make([]string, 0)
	negative := false

	if val == 0 {
		return "zero"
	}
	if val < 0 {
		val *= -1
		negative = true
	}
	if negative && language == EN {
		words = append(words, "negative")
	}
	if language == PT {
		cardinalsScale = ptCardinalsScale
		cardinalsScalePlural = ptCardinalsScalePlural
		cardinalsScaleSet = ptCardinalsScaleSet
		cardinalsHundreds = ptCardinalsHundreds
		cardinalsTens = ptCardinalsTens
		cardinalsSmall = ptCardinalsSmall
		and = ptAnd
	} else if language == EN {
		cardinalsScale = enCardinalsScale
		cardinalsScaleSet = enCardinalsScaleSet
		cardinalsHundreds = enCardinalsHundreds
		cardinalsTens = enCardinalsTens
		cardinalsSmall = enCardinalsSmall
		and = enAnd
	} else {
		panic(errors.New("unknown language: " + language))
	}
	if val >= 1E18 {
		return _NumberIntCardinalInDigits(language, gender, val)
	} else {
		digits := int2digits(val)
		triads := digits2triadseq(digits)

		for i, triad := range triads {
			if language == PT && triad == 100 {
				words = append(words, "cem")
			} else {
				tmp := triad
				words = append(words, getBestNum(&tmp, cardinalsHundreds))
				words = append(words, getBestNum(&tmp, cardinalsTens))
				words = append(words, getBestNum(&tmp, cardinalsSmall))
			}

			scale := (len(triads) - i - 1) * 3
			scaleWord := cardinalsScale[scale]
			if triad > 1 && cardinalsScalePlural != nil {
				scaleWord = cardinalsScalePlural[scale]
			}
			if triad > 0 {
				words = append(words, scaleWord)
			}
		}
	}

	// Filter
	words2 := words
	words = make([]string, 0)
	for _, w := range words2 {
		if w != "" {
			words = append(words, w)
		}
	}

	// Join words
	ans := ""
	if and != " " {
		and = " " + and + " "
	}
	for i, w := range words {
		if language == PT && i == 0 && i+1 < len(words) && w == "um" && words[i+1] == "mil" {
			continue
		}
		if ans != "" {
			switch language {
			case PT:
				if inMap(words[i-1], cardinalsSmall) || cardinalsScaleSet[words[i-1]] || words[i-1] == "cem" {
					ans += " "
				} else {
					ans += and
				}
			default:
				if i != 0 && inMap(words[i-1], cardinalsHundreds) {
					ans += and
				} else {
					ans += " "
				}
			}

		}
		ans += w
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

	// Remebember to tell is the number was negative
	if negative && language == PT {
		ans += " negativ@"
		if val > 1 {
			ans += "s"
		}
	}
	// Fix gender
	ans = strings.Replace(ans, "@", gender_marker, -1)
	return strings.TrimSpace(ans)
}

// not working properly
func _NumberOrdinal(language string, gender string, val int) string {
	negative := false
	words := make([]string, 0)

	if val < 0 {
		val *= -1
		negative = true
	}
	if negative && language == EN {
		words = append(words, "negative")
	}

	if val == 0 {
		words = append(words, ptOrdinal[0])
	}

	for val > 0 {
		words = append(words, getBestNum(&val, ptOrdinal))
	}

	ans := strings.Join(words, " ")

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

	// Remebember to tell is the number was negative
	if negative && language == PT {
		ans += " negativ@"
		if val > 1 {
			ans += "s"
		}
	}
	// Fix gender
	ans = strings.Replace(ans, "@", gender_marker, -1)
	return strings.TrimSpace(ans)
}
