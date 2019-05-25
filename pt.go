package r18n

var ptAnd string = "e"

var ptCardinalsSmall map[int]string = map[int]string{
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
	100: "cento",
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
	1000:             "mil",
	1000000:          "milhão",
	1000000000:       "bilhão",
	1000000000000:    "trilhão",
	1000000000000000: "quadrilhão",
}

var ptCardinalsScalePlural map[int]string = map[int]string{
	1000:             "mil",
	1000000:          "milhões",
	1000000000:       "bilhões",
	1000000000000:    "trilhões",
	1000000000000000: "quadrilhões",
}

var ptCardinalsScaleSet map[string]bool

func init() {
	ptCardinalsScaleSet = make(map[string]bool)
	for _, v := range ptCardinalsScale {
		ptCardinalsScaleSet[v] = true
	}
	for _, v := range ptCardinalsScalePlural {
		ptCardinalsScaleSet[v] = true
	}
}
