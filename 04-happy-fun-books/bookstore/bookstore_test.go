package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {Title: "For the Love of Go"},
		2: {Title: "My First Book"},
	}

	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "My First Book"},
	}

	got := catalog.GetAllBooks()

	// the third value ignores any unexported fields on the struct, ie, private fields
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {
			ID:    1,
			Title: "For the Love of Go",
		},
		2: {
			ID:    2,
			Title: "My First Book",
		},
	}

	want := bookstore.Book{
		ID:    2,
		Title: "My First Book",
	}

	got := catalog.GetBook(2)

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestBuy(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "My First Book",
		Author: "Craig R Broughton",
		Copies: 2,
	}

	want := 1

	result, err := bookstore.Buy(b)

	if err != nil {
		t.Fatal(err)
	}

	got := result.Copies

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "My First Book",
		Author: "Craig R Broughton",
		Copies: 0,
	}

	_, err := bookstore.Buy(b)

	if err == nil {
		t.Error("want error buying from zero copies, got nils")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:           "My First Book",
		PriceCents:      4000,
		DiscountPercent: 25,
	}

	want := 3000

	got := b.NetPriceCents()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:      "My First Book",
		PriceCents: 4000,
	}

	want := 3000

	err := b.SetPriceCents(want)

	if err != nil {
		t.Fatal(err)
	}

	got := b.PriceCents

	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title: "My First Book",
	}

	categories := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}

	for _, category := range categories {
		err := b.SetCategory(category)

		if err != nil {
			t.Fatal(err)
		}
		got := b.GetCategory()

		if category != got {
			t.Errorf("want category %q, got %q", category, got)
		}
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory(999)
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}
