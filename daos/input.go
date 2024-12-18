package daos

import (
	"fmt"
	"piotrek813/goth/db"
	"piotrek813/goth/models"
	"reflect"
	"strings"
)

func getNamedArgs(obj any, args *[]string, values *[]any) {
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

func getQuestionMakrs(n int) string {
	questionMarks := []string{}
	for i := 0; i < n; i++ {
		questionMarks = append(questionMarks, "?")
	}

	return strings.Join(questionMarks, ",")
}

func SaveInput(input models.Input) error {
	db := db.GetDB()

	args := []string{}
	values := []any{}

	getNamedArgs(input, &args, &values)

	argsStr := strings.Join(args, ",")
	questionmarks := getQuestionMakrs(len(args))

	q := fmt.Sprintf("INSERT INTO %v(%v) VALUES(%v)", input.TableName(), argsStr, questionmarks)

	stmt, err := db.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(values...)

	if err != nil {
		return err
	}

	return nil
}
