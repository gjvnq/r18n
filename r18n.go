package r18n

const (
	GENDER_MALE       = "M" // Ex: aluno
	GENDER_FEMALE     = "F" // Ex: aluna
	GENDER_NON_BINARY = "X" // Ex: alunx
	SINGULAR          = "S" // Ex: aluno
	PLURAL            = "P" // Ex: alunos
)

const (
	EN = "en"
	PT = "pt"
)

type MsgPart interface {
	GenderAndNumber() string
	String() string
}

type MsgPartNumerical struct {
	Amount int
}

type MsgPartMoney struct {
	Amount   int // in cents
	Currency string
}

type R18N struct {
	Messages            map[string]map[string]string
	FallbackGenderOrder []string
	FallbackPluralOrder []string
}

func (this *R18N) init() {
	if this.Messages == nil {
		this.Messages = make(map[string]map[string]string)
	}
}

func NewR18N() *R18N {
	ans := new(R18N)
	ans.init()
	return ans
}
