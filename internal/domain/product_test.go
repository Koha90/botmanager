package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestProduct_AddVariant(t *testing.T) {
	p, _ := NewProduct("Test Product", 1, "Description", "")

	require.Equal(t, 1, p.Version())

	err := p.AddVariant("1p", 1, 100)
	require.NoError(t, err)
	require.Equal(t, 2, p.Version())
}

func TestNewProduct(t *testing.T) {
	_, err := NewProduct("", 1, "", "")
	require.ErrorIs(t, err, ErrInvalidProductName)

	_, err = NewProduct("Tea", 0, "", "")
	require.ErrorIs(t, err, ErrInvalidCategoryID)

	p, err := NewProduct("Tea", 1, "desc", "img")
	require.NoError(t, err)
	require.Equal(t, 1, p.Version())
}

func TestProduct_Rename(t *testing.T) {
	p, _ := NewProduct("Tea", 1, "", "")

	err := p.Rename("")
	require.ErrorIs(t, err, ErrInvalidProductName)

	require.NoError(t, p.Rename("Coffee"))
	require.Equal(t, "Coffee", p.Name())
	require.Equal(t, 2, p.Version())
}

func TestProduct_AddVariant_Duplicate(t *testing.T) {
	p, _ := NewProduct("Tea", 1, "", "")

	require.NoError(t, p.AddVariant("250g", 1, 100))
	err := p.AddVariant("250g", 1, 200)

	require.ErrorIs(t, err, ErrVariantAlreadyExists)
}

func TestProduct_ArchiveVariant(t *testing.T) {
	p, _ := NewProduct("Tea", 1, "", "")

	_ = p.AddVariant("250g", 1, 100)
	_ = p.AddVariant("500g", 1, 200)

	variants := p.ActiveVariants()
	require.Len(t, variants, 2)

	err := p.ArchiveVariant(variants[0].ID(), time.Now())
	require.NoError(t, err)

	require.Len(t, p.ActiveVariants(), 1)
}

func TestProduct_Archive_LastVariant(t *testing.T) {
	p, _ := NewProduct("Tea", 1, "", "")

	_ = p.AddVariant("250g", 1, 100)
	v := p.ActiveVariants()[0]

	err := p.ArchiveVariant(v.ID(), time.Now())
	require.ErrorIs(t, err, ErrCannotArchiveLastVariant)
}

func TestProduct_VariantByID(t *testing.T) {
	p, _ := NewProduct("Tea", 1, "", "")
	_ = p.AddVariant("250g", 1, 100)

	v := p.ActiveVariants()[0]

	found, err := p.VariantByID(v.ID())
	require.NoError(t, err)
	require.Equal(t, v.ID(), found.ID())

	_, err = p.VariantByID(999)
	require.ErrorIs(t, err, ErrVariantNotFound)
}

func TestProduct_PullEvents(t *testing.T) {
	p, _ := NewProduct("Tea", 1, "", "")
	_ = p.AddVariant("250g", 1, 100)

	ev := p.PullEvents()
	require.NotEmpty(t, ev)

	require.Empty(t, p.PullEvents())
}
