package daos

import (
	"fmt"
	"piotrek813/goth/db"
	"piotrek813/goth/models"
	"reflect"
	"strings"
)

func getSelectArgs(obj any, args *[]string, values *[]any) {
	rtype := reflect.TypeOf(obj)
	rval := reflect.ValueOf(obj)

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		value := rval.Field(i)

		if value.Kind() == reflect.Struct {
			getNamedArgs(value.Interface(), args, values)
		}

		tag := field.Tag.Get("db")

		if tag != "" {
			*args = append(*args, tag)
			*values = append(*values, value)
		}
	}
}

func Select(input models.Model) error {
	db := db.GetDB()

	args := []string{}
	values := []any{}

	getSelectArgs(input, &args, &values)

	argsStr := strings.Join(args, ",")

	q := fmt.Sprintf("SELECT %v FROM %v", argsStr, input.TableName())

	row := db.QueryRow(q)

	row.Scan(values...)

	return nil
}
