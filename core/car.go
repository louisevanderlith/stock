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
	EstValue      int64
	LicenseExpiry time.Time
}

func (o Car) Valid() (bool, error) {
	var issues []string

	valid, common := husk.ValidateStruct(&o)
	if !valid {
		issues = append(issues, common.Error())
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
	//Estimate Value should be populated with the Average price of similiar vehicles.
	if err := PriceInBounds(o.Price, o.EstValue); err != nil {
		return false, err
	}

	isValid := len(issues) < 1
	finErr := errors.New(strings.Join(issues, "\r\n"))

	return isValid, finErr
}

func GetCar(key husk.Key) (*Car, error) {
	rec, err := ctx.Cars.FindByKey(key)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*Car), nil
}

func GetLatestCars(page, size int) husk.Collection {
	return ctx.Cars.Find(page, size, husk.Everything())
}

func (c Car) Create() husk.CreateSet {
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
