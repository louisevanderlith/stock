package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
	"time"

	"errors"
)

type Car struct {
	StockItem
	VehicleKey    *keys.TimeKey
	Info          string `hsk:"size(128)"`
	Year          int    `orm:"null"`
	Mileage       int    `orm:"null"`
	HasNatis      bool   `hsk:"default(false)"`
	EstValue      int64
	LicenseExpiry time.Time
}

func (o Car) Valid() error {
	err := validation.Struct(o)
	if err != nil {
		return err
	}

	if o.Year > 0 && o.Year > time.Now().Year() {
		errors.New("year can't be in the future")
	}

	if o.Mileage < 0 {
		return errors.New("mileage can't be negative")
	}

	if o.HasNatis && o.LicenseExpiry.Before(time.Now()) {
		return errors.New("license has already expired")
	}

	//Price compare - Fair Price?
	//Estimate Value should be populated with the Average price of the same types of vehicles.
	return PriceInBounds(o.Price, o.EstValue)
}

func (c Car) Create() (hsk.Key, error) {
	return ctx.Cars.Create(c)
}

func (c Car) Update(key hsk.Key) error {
	return ctx.Cars.Update(key, c)
}
