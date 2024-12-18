package models

type TextInput struct {
	BaseInput

	Value string `db:"value"`
}

func (i TextInput) TableName() string {
	return "TextInputs"
}
