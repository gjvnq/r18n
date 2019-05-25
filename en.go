package r18n

var enAnd string = "and"

var enCardinalsSmall map[int]string = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifthteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var enCardinalsTens map[int]string = map[int]string{
	20: "twenty",
	30: "thirty",
	40: "fourty",
	50: "fifity",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "nighty",
}

var enCardinalsHundreds map[int]string = map[int]string{
	100: "one hundred",
	200: "two hundred",
	300: "three hundred",
	400: "four hundred",
	500: "five hundred",
	600: "six hundred",
	700: "seven hundred",
	800: "eight hundred",
	900: "nine hundred",
}

var enCardinalsScale map[int]string = map[int]string{
	1000:             "thousand",
	1000000:          "million",
	1000000000:       "billion",
	1000000000000:    "trillion",
	1000000000000000: "quadrillion",
}

var enCardinalsScalePlural map[int]string = map[int]string{
	1000:             "thousands",
	1000000:          "millions",
	1000000000:       "billions",
	1000000000000:    "trillions",
	1000000000000000: "quadrillions",
}

var enCardinalsScaleSet map[string]bool

func init() {
	enCardinalsScaleSet = make(map[string]bool)
	for _, v := range enCardinalsScale {
		enCardinalsScaleSet[v] = true
	}
	for _, v := range enCardinalsScalePlural {
		enCardinalsScaleSet[v] = true
	}
}
