package core

import (
	"github.com/louisevanderlith/husk"
)

type StockContext interface {
	GetService(key husk.Key) (Service, error)
	FindServices(page, size int, profile string) (husk.Collection, error)
	CreateService(obj Service) (husk.Recorder, error)
	UpdateService(key husk.Key, obj Service) error
	GetCar(key husk.Key) (Car, error)
	FindLatestCars(page, size int, profile string) (husk.Collection, error)
	CreateCar(obj Car) (husk.Recorder, error)
	UpdateCar(key husk.Key, obj Car) error
	GetPart(key husk.Key) (Part, error)
	FindLatestParts(page, size int, profile string) (husk.Collection, error)
	CreatePart(obj Part) (husk.Recorder, error)
	UpdatePart(key husk.Key, obj Part) error
	GetProperty(key husk.Key) (Property, error)
	FindLatestProperties(page, size int, profile string) (husk.Collection, error)
	CreateProperty(obj Property) (husk.Recorder, error)
	UpdateProperty(key husk.Key, obj Property) error
}

type context struct {
	Cars       husk.Tabler
	Services   husk.Tabler
	Parts      husk.Tabler
	Properties husk.Tabler
}

var ctx context

func Context() StockContext {
	return ctx
}

func CreateContext() {
	ctx = context{
		Cars:       husk.NewTable(Car{}),
		Services:   husk.NewTable(Service{}),
		Parts:      husk.NewTable(Part{}),
		Properties: husk.NewTable(Property{}),
	}
}

func Shutdown() {
	ctx.Cars.Save()
	ctx.Services.Save()
	ctx.Parts.Save()
	ctx.Properties.Save()
}

func (c context) GetService(key husk.Key) (Service, error) {
	rec, err := c.Services.FindByKey(key)

	if err != nil {
		return Service{}, err
	}

	return rec.Data().(Service), nil
}

func (c context) FindServices(page, size int, profile string) (husk.Collection, error) {
	return c.Services.Find(page, size, byProfile(profile))
}

func (c context) CreateService(obj Service) (husk.Recorder, error) {
	return c.Services.Create(obj)
}

func (c context) UpdateService(key husk.Key, obj Service) error {
	rec, err := c.Services.FindByKey(key)

	if err != nil {
		return err
	}

	err = rec.Set(obj)

	if err != nil {
		return nil
	}

	err = c.Services.Update(rec)

	if err != nil {
		return err
	}

	return c.Services.Save()
}

func (c context) GetCar(key husk.Key) (Car, error) {
	rec, err := c.Cars.FindByKey(key)

	if err != nil {
		return Car{}, err
	}

	return rec.Data().(Car), nil
}

func (c context) FindLatestCars(page, size int, profile string) (husk.Collection, error) {
	return c.Cars.Find(page, size, byProfile(profile))
}

func (c context) CreateCar(obj Car) (husk.Recorder, error) {
	return c.Cars.Create(obj)
}

func (c context) UpdateCar(key husk.Key, obj Car) error {
	rec, err := c.Cars.FindByKey(key)

	if err != nil {
		return err
	}

	err = rec.Set(obj)

	if err != nil {
		return nil
	}

	err = c.Cars.Update(rec)

	if err != nil {
		return err
	}

	return c.Cars.Save()
}

func (c context) GetPart(key husk.Key) (Part, error) {
	rec, err := c.Parts.FindByKey(key)

	if err != nil {
		return Part{}, err
	}

	return rec.Data().(Part), nil
}

func (c context) FindLatestParts(page, size int, profile string) (husk.Collection, error) {
	return c.Parts.Find(page, size, byProfile(profile))
}

func (c context) CreatePart(obj Part) (husk.Recorder, error) {
	return c.Parts.Create(obj)
}

func (c context) UpdatePart(key husk.Key, obj Part) error {
	rec, err := c.Parts.FindByKey(key)

	if err != nil {
		return err
	}

	err = rec.Set(obj)

	if err != nil {
		return nil
	}

	err = c.Parts.Update(rec)

	if err != nil {
		return err
	}

	return c.Parts.Save()
}

func (c context) GetProperty(key husk.Key) (Property, error) {
	rec, err := c.Properties.FindByKey(key)

	if err != nil {
		return Property{}, err
	}

	return rec.Data().(Property), nil
}

func (c context) FindLatestProperties(page, size int, profile string) (husk.Collection, error) {
	return c.Properties.Find(page, size, byProfile(profile))
}

func (c context) CreateProperty(obj Property) (husk.Recorder, error) {
	return c.Properties.Create(obj)
}

func (c context) UpdateProperty(key husk.Key, obj Property) error {
	rec, err := c.Properties.FindByKey(key)

	if err != nil {
		return err
	}

	err = rec.Set(obj)

	if err != nil {
		return nil
	}

	err = c.Properties.Update(rec)

	if err != nil {
		return err
	}

	return c.Properties.Save()
}
