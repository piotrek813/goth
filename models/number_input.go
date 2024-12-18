package models

type NumberInput struct {
	BaseInput

	Value int
}

func (i *NumberInput) TableName() string {
	return "NumberInputs"
}
