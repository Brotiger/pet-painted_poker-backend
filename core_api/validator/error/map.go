package validatorError

type ValidatorError interface {
	GetError(param string) string
}

var Map = map[string]ValidatorError{
	"min":      &Min{},
	"required": &Required{},
	"email":    &Email{},
}
