package models

type BaseModel struct {
	Id        int
	CreatedAt string
}

type Model interface {
	TableName() string
}
