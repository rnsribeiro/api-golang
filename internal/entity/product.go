package entity

import (
	"errors"
	"time"

	"github.com/rnsribeiro/api-golang/pkg/entity"
)

var (
	errIDIsRequired    = errors.New("the id is required")
	errIvalidID        = errors.New("the id is invalid")
	errNameIsRequired  = errors.New("the name is required")
	errPriceIsRequired = errors.New("the price is required")
	errInValidPrice    = errors.New("the price must be greater than zero")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return errIvalidID
	}
	if p.Name == "" {
		return errNameIsRequired
	}
	if p.Price == 0 {
		return errPriceIsRequired
	}
	if p.Price < 0 {
		return errInValidPrice
	}
	return nil
}
