package daos

import (
	"fmt"
	"piotrek813/goth/models"
	"testing"
)

func TestSaveInput(t *testing.T) {
	input := models.TextInput{
		Value: "cdscs",
		BaseInput: models.BaseInput{
			Name: "dupa",
		},
	}

	err := SaveInput(input)
	fmt.Printf("err: %v\n", err)
}
