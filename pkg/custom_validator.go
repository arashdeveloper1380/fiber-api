package pkg

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Rule  string
	Msg   string
}

type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	var sb strings.Builder
	for _, e := range ve {
		sb.WriteString(fmt.Sprintf("Field '%s' failed on '%s': %s\n", e.Field, e.Rule, e.Msg))
	}
	return sb.String()
}

func ValidationStruct(s interface{}) error {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}

	var errs ValidationErrors

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("arash_validate")

		if tag == "" {
			continue
		}

		rules := strings.Split(tag, ",")

		for _, rule := range rules {

			if strings.Contains(rule, "=") {
				parts := strings.SplitN(rule, "=", 2)
				key := parts[0]
				param := parts[1]

				switch key {
				case "min":
					if err := checkMin(value, param, field.Name); err != nil {
						errs = append(errs, ValidationError{field.Name, "min", err.Error()})
					}
				case "max":
					if err := checkMax(value, param, field.Name); err != nil {
						errs = append(errs, ValidationError{field.Name, "max", err.Error()})
					}
				case "len":
					if err := checkLen(value, param, field.Name); err != nil {
						errs = append(errs, ValidationError{field.Name, "len", err.Error()})
					}
				}
				continue
			}

			switch rule {
			case "required":
				if isEmpty(value) {
					errs = append(errs, ValidationError{field.Name, "required", "field is required"})
				}
			case "positive":
				if value.Kind() == reflect.Int && value.Int() <= 0 {
					errs = append(errs, ValidationError{field.Name, "positive", "must be positive"})
				}
			case "email":
				if value.Kind() == reflect.String && !strings.Contains(value.String(), "@") {
					errs = append(errs, ValidationError{field.Name, "email", "must be a valid email"})
				}
			}
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	}
	return false
}

func checkMin(v reflect.Value, param, fieldName string) error {
	min, _ := strconv.Atoi(param)
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		if v.Len() < min {
			return fmt.Errorf("field '%s' must have at least %d characters/items", fieldName, min)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() < int64(min) {
			return fmt.Errorf("field '%s' must be >= %d", fieldName, min)
		}
	}
	return nil
}

func checkMax(v reflect.Value, param, fieldName string) error {
	max, _ := strconv.Atoi(param)
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		if v.Len() > max {
			return fmt.Errorf("field '%s' must have at most %d characters/items", fieldName, max)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() > int64(max) {
			return fmt.Errorf("field '%s' must be <= %d", fieldName, max)
		}
	}
	return nil
}

func checkLen(v reflect.Value, param, fieldName string) error {
	l, _ := strconv.Atoi(param)
	if v.Kind() == reflect.String || v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if v.Len() != l {
			return fmt.Errorf("field '%s' must have length %d", fieldName, l)
		}
	}
	return nil
}

type InfoUser struct {
	Name  string   `arash_validate:"required"`
	Age   int      `arash_validate:"required,positive,min=18,max=100"`
	Email string   `arash_validate:"required,email"`
	Tags  []string `arash_validate:"min=3,max=5"`
}

func CustomValidate() {
	info := InfoUser{
		Name:  "arash",
		Age:   15,
		Email: "arash@gmail.com",
		Tags:  []string{"go", "reflect"},
	}

	if err := ValidationStruct(info); err != nil {
		fmt.Println("❌ Validation error:\n", err)
	} else {
		fmt.Println("✅ info is valid")
	}
}
