package decimalvalidators

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestDecimalValidatorGt(t *testing.T) {
	v := validator.New()
	RegisterDecimalValidators(v)

	rec := struct {
		Val decimal.Decimal `validate:"dgt=10"`
	}{}

	// Less Than
	rec.Val = decimal.NewFromInt(9)
	err := v.Struct(rec)
	require.Error(t, err)

	// Equal
	rec.Val = decimal.NewFromInt(10)
	err = v.Struct(rec)
	require.Error(t, err)

	// Greater than
	rec.Val = decimal.NewFromInt(11)
	err = v.Struct(rec)
	require.NoError(t, err)

}

func TestDecimalValidatorGte(t *testing.T) {
	v := validator.New()
	RegisterDecimalValidators(v)

	rec := struct {
		Val decimal.Decimal `validate:"dgte=10"`
	}{}

	// Less Than
	rec.Val = decimal.NewFromInt(9)
	err := v.Struct(rec)
	require.Error(t, err)

	// Equal
	rec.Val = decimal.NewFromInt(10)
	err = v.Struct(rec)
	require.NoError(t, err)

	// Greater than
	rec.Val = decimal.NewFromInt(11)
	err = v.Struct(rec)
	require.NoError(t, err)

}

func TestDecimalValidatorLt(t *testing.T) {
	v := validator.New()
	RegisterDecimalValidators(v)

	rec := struct {
		Val decimal.Decimal `validate:"dlt=10"`
	}{}

	// Less Than
	rec.Val = decimal.NewFromInt(9)
	err := v.Struct(rec)
	require.NoError(t, err)

	// Equal
	rec.Val = decimal.NewFromInt(10)
	err = v.Struct(rec)
	require.Error(t, err)

	// Greater than
	rec.Val = decimal.NewFromInt(11)
	err = v.Struct(rec)
	require.Error(t, err)

}

func TestDecimalValidatorLte(t *testing.T) {
	v := validator.New()
	RegisterDecimalValidators(v)

	rec := struct {
		Val decimal.Decimal `validate:"dlte=10"`
	}{}

	// Less Than
	rec.Val = decimal.NewFromInt(9)
	err := v.Struct(rec)
	require.NoError(t, err)

	// Equal
	rec.Val = decimal.NewFromInt(10)
	err = v.Struct(rec)
	require.NoError(t, err)

	// Greater than
	rec.Val = decimal.NewFromInt(11)
	err = v.Struct(rec)
	require.Error(t, err)

}

func TestDecimalValidatorEq(t *testing.T) {
	v := validator.New()
	RegisterDecimalValidators(v)

	rec := struct {
		Val decimal.Decimal `validate:"deq=10"`
	}{}

	// Less Than
	rec.Val = decimal.NewFromInt(9)
	err := v.Struct(rec)
	require.Error(t, err)

	// Equal
	rec.Val = decimal.NewFromInt(10)
	err = v.Struct(rec)
	require.NoError(t, err)

	// Greater than
	rec.Val = decimal.NewFromInt(11)
	err = v.Struct(rec)
	require.Error(t, err)

}

func TestDecimalValidatorNeq(t *testing.T) {
	v := validator.New()
	RegisterDecimalValidators(v)

	rec := struct {
		Val decimal.Decimal `validate:"dneq=10"`
	}{}

	// Less Than
	rec.Val = decimal.NewFromInt(9)
	err := v.Struct(rec)
	require.NoError(t, err)

	// Equal
	rec.Val = decimal.NewFromInt(10)
	err = v.Struct(rec)
	require.Error(t, err)

	// Greater than
	rec.Val = decimal.NewFromInt(11)
	err = v.Struct(rec)
	require.NoError(t, err)

}
