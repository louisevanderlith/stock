package core

import (
	"strings"
	"time"

	"errors"

	"github.com/louisevanderlith/husk"
)

type Car struct {
	StockItem
	VehicleKey    husk.Key
	Info          string `hsk:"size(128)"`
	Year          int    `orm:"null"`
	Mileage       int    `orm:"null"`
	HasNatis      bool   `hsk:"default(false)"`
	EstValue      uint64
	LicenseExpiry time.Time
}

func (o Car) Valid() error {
	var issues []string

	err := husk.ValidateStruct(o)
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

func (c Car) Create() (husk.Recorder, error) {
	return ctx.Cars.Create(c)
}

func (c Car) Update(key husk.Key) error {
	obj, err := ctx.Cars.FindByKey(key)

	if err != nil {
		return err
	}

	err = obj.Set(c)

	if err != nil {
		return nil
	}

	defer ctx.Cars.Save()
	return ctx.Cars.Update(obj)
}
