package cfg

// Validations represent a collection of config setting validations
type Validations struct {
	validations map[string]Validatable
}

// NewValidations creates and returns an instance of Validations
func NewValidations() *Validations {
	vals := &Validations{
		validations: make(map[string]Validatable),
	}

	return vals
}

func (vals *Validations) append(key string, posVal Validatable) {
	vals.validations[key] = posVal
}

func (vals *Validations) valueFor(key string) int {
	val := vals.validations[key]
	if val != nil {
		return val.IntValue()
	}

	return 0
}
