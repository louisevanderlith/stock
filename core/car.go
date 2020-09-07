package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
	"strings"
	"time"

	"errors"
)

type Car struct {
	StockItem
	VehicleKey    hsk.Key
	Info          string `hsk:"size(128)"`
	Year          int    `orm:"null"`
	Mileage       int    `orm:"null"`
	HasNatis      bool   `hsk:"default(false)"`
	EstValue      uint64
	LicenseExpiry time.Time
}

func (o Car) Valid() error {
	var issues []string

	err := validation.Struct(o)
	if err != nil {
		issues = append(issues, err.Error())
	}

	if o.Year > 0 && o.Year > time.Now().Year() {
		issues = append(issues, "Year can't be in the future.")
	}

	if o.Mileage < 0 {
		issues = append(issues, "Mileage can't be negative.")
	}

	if o.HasNatis && o.LicenseExpiry.Before(time.Now()) {
		issues = append(issues, "License has already expired.")
	}

	//Price compare - Fair Price?
	//Estimate Value should be populated with the Average price of the same types of vehicles.
	if err := PriceInBounds(o.Price, o.EstValue); err != nil {
		return err
	}

	finErr := errors.New(strings.Join(issues, "\r\n"))

	return finErr
}

func (c Car) Create() (hsk.Key, error) {
	return ctx.Cars.Create(c)
}

func (c Car) Update(key hsk.Key) error {
	return ctx.Cars.Update(key, c)
}
