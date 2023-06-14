package bookstore

import (
	"errors"
	"fmt"
)

type Catalog map[int]Book

type Book struct {
	ID              int
	PriceCents      int
	DiscountPercent int
	Title           string
	Author          string
	Copies          int
}

func (b Book) NetPriceCents() int {
	savings := b.PriceCents * b.DiscountPercent / 100

	return b.PriceCents - savings
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("negative price %d", price)
	}
	b.PriceCents = price
	return nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}

	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(ID int) Book {
	for _, b := range c {
		if b.ID == ID {
			return b
		}
	}
	return Book{}
}
