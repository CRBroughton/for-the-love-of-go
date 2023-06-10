package bookstore_test

import (
	"bookstore"
	"sort"
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

	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "My First Book"},
	}

	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "My First Book"},
	}

	got := bookstore.GetAllBooks(catalog)

	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	catalog := map[int]bookstore.Book{
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

	got, err := bookstore.GetBook(catalog, 2)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()

	catalog := map[int]bookstore.Book{}

	_, err := bookstore.GetBook(catalog, 999)

	if err == nil {
		t.Fatal("want error for non-existant ID, got nil")
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
