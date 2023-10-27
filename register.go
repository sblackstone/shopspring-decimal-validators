package decimalvalidators

import (
	"errors"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	validator "github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func RegisterDecimalValidators(v *validator.Validate) {
	registerDecimalGreaterThan(v)
	registerDecimalGreaterEqual(v)
	registerDecimalLessThan(v)
	registerDecimalLessThanOrEqual(v)
	registerDecimalEqual(v)
	registerDecimalNotEqual(v)
	registerDecimalType(v)
}

func registerDecimalType(v *validator.Validate) {
	v.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		if valuer, ok := field.Interface().(decimal.Decimal); ok {
			return valuer.String()
		}
		return nil
	}, decimal.Decimal{})
}

func getValues(fl validator.FieldLevel) (*decimal.Decimal, *decimal.Decimal, error) {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return nil, nil, errors.New("")
	}
	value, err := decimal.NewFromString(data)
	if err != nil {
		return nil, nil, errors.New("")
	}
	baseValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return nil, nil, errors.New("")
	}
	return &value, &baseValue, nil
}

func registerDecimalGreaterThan(v *validator.Validate) {
	err := v.RegisterValidation("dgt", func(fl validator.FieldLevel) bool {
		if value, baseValue, err := getValues(fl); err != nil {
			spew.Dump(err)
			return false
		} else {
			spew.Dump(value)
			spew.Dump(*baseValue)
			return value.GreaterThan(*baseValue)
		}
	})
	spew.Dump(err)
}

func registerDecimalGreaterEqual(v *validator.Validate) {
	v.RegisterValidation("dgte", func(fl validator.FieldLevel) bool {
		if value, baseValue, err := getValues(fl); err != nil {
			return false
		} else {
			return value.GreaterThanOrEqual(*baseValue)
		}
	})
}

func registerDecimalLessThan(v *validator.Validate) {
	v.RegisterValidation("dlt", func(fl validator.FieldLevel) bool {
		if value, baseValue, err := getValues(fl); err != nil {
			return false
		} else {
			return value.LessThan(*baseValue)
		}
	})
}

func registerDecimalLessThanOrEqual(v *validator.Validate) {
	v.RegisterValidation("dlte", func(fl validator.FieldLevel) bool {
		if value, baseValue, err := getValues(fl); err != nil {
			return false
		} else {
			return value.LessThanOrEqual(*baseValue)
		}
	})
}

func registerDecimalEqual(v *validator.Validate) {
	v.RegisterValidation("deq", func(fl validator.FieldLevel) bool {
		if value, baseValue, err := getValues(fl); err != nil {
			return false
		} else {
			return value.Equal(*baseValue)
		}
	})
}

func registerDecimalNotEqual(v *validator.Validate) {
	v.RegisterValidation("dneq", func(fl validator.FieldLevel) bool {
		if value, baseValue, err := getValues(fl); err != nil {
			return false
		} else {
			return !value.Equal(*baseValue)
		}
	})
}
