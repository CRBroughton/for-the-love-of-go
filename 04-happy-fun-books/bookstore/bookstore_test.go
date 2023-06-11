package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
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

	if !cmp.Equal(want, got) {
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

	if !cmp.Equal(want, got) {
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
