package core

import (
	"strings"
	"time"

	"errors"

	"github.com/louisevanderlith/husk"
)

type CarAdvert struct {
	Advert
	VehicleKey    husk.Key
	Info          string `hsk:"size(128)"`
	Year          int    `orm:"null"`
	Mileage       int    `orm:"null"`
	HasNatis      bool   `hsk:"default(false)"`
	LicenseExpiry time.Time
}

func (o CarAdvert) Valid() (bool, error) {
	var issues []string

	valid, common := husk.ValidateStruct(&o)
	if !valid {
		issues = append(issues, common.Error())
	}

	if o.Year > 0 && o.Year > time.Now().Year() {
		issues = append(issues, "Year can't be in the future.")
	}

	if o.Mileage < 0 {
		issues = append(issues, "Odometer can't be negative.")
	}

	if o.HasNatis && o.LicenseExpiry.Before(time.Now()) {
		issues = append(issues, "License has already expired.")
	}

	isValid := len(issues) < 1
	finErr := errors.New(strings.Join(issues, "\r\n"))

	return isValid, finErr
}

func GetCarAdvert(key husk.Key) (*CarAdvert, error) {
	rec, err := ctx.Cars.FindByKey(key)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*CarAdvert), nil
}

func GetLatestCars(page, size int) husk.Collection {
	return ctx.Cars.Find(page, size, husk.Everything())
}

func (c CarAdvert) Create() husk.CreateSet {
	return ctx.Cars.Create(c)
}

func (c CarAdvert) Update(key husk.Key) error {
	obj, err := ctx.Cars.FindByKey(key)

	if err != nil {
		return err
	}

	obj.Set(c)

	defer ctx.Cars.Save()
	return ctx.Cars.Update(obj)
}
