package r18n

var ptAnd string = "e"

var ptOrdinal map[int]string = map[int]string{
	0:             "zerésim@",
	1:             "primeir@",
	2:             "segund@",
	3:             "terceir@",
	4:             "quart@",
	5:             "quint@",
	6:             "sext@",
	7:             "sétim@",
	8:             "oitav@",
	9:             "non@",
	10:            "décim@",
	20:            "vigésim@",
	30:            "trigésim@",
	40:            "quadragésim@",
	50:            "quinquagésim@",
	60:            "sexagésim@",
	70:            "septuagésim@",
	80:            "octogésim@",
	90:            "nonagésim@",
	100:           "centésim@",
	200:           "ducentésimo",
	300:           "tricentésimo",
	400:           "quadringentésimo",
	500:           "quingentésimo",
	600:           "sexcentésimo",
	700:           "septingentésimo",
	800:           "octingentésimo",
	900:           "noningentésimo",
	1000:          "milésimo",
	2000:          "dumilésimo",
	3000:          "trimilésimo",
	4000:          "quadrimilésimo",
	5000:          "quinquemilésimo",
	6000:          "sexmilésimo",
	7000:          "heptamilésimo",
	8000:          "octamilésimo",
	9000:          "nonamilésimo",
	10000:         "décimo milésimo",
	100000:        "centésimo milésimo",
	1000000:       "milionésimo",
	1000000000:    "bilionésimo",
	1000000000000: "trilionésimo",
}

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
	-3: "milésimo",
	-2: "centavo",
	-1: "décimo",
	3:  "mil",
	6:  "milhão",
	9:  "bilhão",
	12: "trilhão",
	15: "quadrilhão",
}

var ptCardinalsScalePlural map[int]string = map[int]string{
	-3: "milésimos",
	-2: "centavos",
	-1: "décimos",
	3:  "mil",
	6:  "milhões",
	9:  "bilhões",
	12: "trilhões",
	15: "quadrilhões",
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
