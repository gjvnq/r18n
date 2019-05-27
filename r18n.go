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

type R18N struct {
	Messages            map[string]map[string]Message
	FallbackGenderOrder []string
	FallbackPluralOrder []string
}

func (this *R18N) init() {
	if this.Messages == nil {
		this.Messages = make(map[string]map[string]string)
	}
}

func (this *R18N) T(lang string, msg_id string, values ...interface{}) string {
	return ""
}

func NewR18N() *R18N {
	ans := new(R18N)
	ans.init()
	return ans
}
