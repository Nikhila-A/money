package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMoney(t *testing.T) {
	t.Run("check whether money is initialized", func(t *testing.T) {
		assert.IsType(t, Money{}, NewMoney(29))
	})

	t.Run("check initialization of money with negative paise", func(t *testing.T) {
		assert.Panics(t, func() {
			NewMoney(-2)
		})
	})

	t.Run("check initialization of money with positive paise", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewMoney(3)
		})
	})

	t.Run("check initialization of money with zero paise", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewMoney(0)
		})
	})
}

func TestAddMoney(t *testing.T) {
	t.Run("check whether the money is correct when paise is negative", func(t *testing.T) {
		assert.Panics(t, func() {
			money := NewMoney(5)
			money.AddMoney(-8)
		})
	})

	money := NewMoney(1000)
	t.Run("check whether the money is correct when paise is positive", func(t *testing.T) {
		money.AddMoney(120)
		assert.Equal(t, 1120, money.paise)
	})

	money = NewMoney(1000)
	t.Run("check value when paise is zero", func(t *testing.T) {
		money.AddMoney(0)
		assert.Equal(t, 1000, money.paise)
	})
}
