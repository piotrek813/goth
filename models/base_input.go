package models

const (
	TextInputType = iota
	NumberInputType
)

type BaseInput struct {
	BaseModel

	Name     string `db:"name"`
	TypeCode int
}

type Input interface {
	Model
}
