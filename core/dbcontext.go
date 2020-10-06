package core

import (
	"encoding/json"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"os"
	"reflect"
)

type StockContext interface {
	GetService(key hsk.Key) (Service, error)
	FindServices(page, size int, profile string) (records.Page, error)
	CreateService(obj Service) (hsk.Key, error)
	UpdateService(key hsk.Key, obj Service) error
	GetClothing(key hsk.Key) (Clothing, error)
	FindClothing(page, size int, profile string) (records.Page, error)
	CreateClothing(obj Clothing) (hsk.Key, error)
	UpdateClothing(key hsk.Key, obj Clothing) error
	GetCar(key hsk.Key) (Car, error)
	FindLatestCars(page, size int, profile string) (records.Page, error)
	CreateCar(obj Car) (hsk.Key, error)
	UpdateCar(key hsk.Key, obj Car) error
	GetPart(key hsk.Key) (Part, error)
	FindLatestParts(page, size int, profile string) (records.Page, error)
	CreatePart(obj Part) (hsk.Key, error)
	UpdatePart(key hsk.Key, obj Part) error
	GetProperty(key hsk.Key) (Property, error)
	FindLatestProperties(page, size int, profile string) (records.Page, error)
	CreateProperty(obj Property) (hsk.Key, error)
	UpdateProperty(key hsk.Key, obj Property) error
}

type context struct {
	Cars       husk.Table
	Services   husk.Table
	Parts      husk.Table
	Properties husk.Table
	Clothing   husk.Table
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
		Clothing:   husk.NewTable(Clothing{}),
	}

	seed()
}

func seed() {
	services, err := serviceSeeds()

	if err != nil {
		panic(err)
	}

	err = ctx.Services.Seed(services)

	if err != nil {
		panic(err)
	}

	clothes, err := clothingSeeds()

	if err != nil {
		panic(err)
	}

	err = ctx.Clothing.Seed(clothes)

	if err != nil {
		panic(err)
	}
}

func serviceSeeds() (collections.Enumerable, error) {
	f, err := os.Open("db/services.seed.json")

	if err != nil {
		return nil, err
	}

	var items []Service
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}

func clothingSeeds() (collections.Enumerable, error) {
	f, err := os.Open("db/clothing.seed.json")

	if err != nil {
		return nil, err
	}

	var items []Clothing
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}

func Shutdown() {
	ctx.Cars.Save()
	ctx.Services.Save()
	ctx.Parts.Save()
	ctx.Properties.Save()
}

func (c context) GetService(key hsk.Key) (Service, error) {
	rec, err := c.Services.FindByKey(key)

	if err != nil {
		return Service{}, err
	}

	return rec.GetValue().(Service), nil
}

func (c context) FindServices(page, size int, profile string) (records.Page, error) {
	return c.Services.Find(page, size, byServiceProfile(profile))
}

func (c context) CreateService(obj Service) (hsk.Key, error) {
	return c.Services.Create(obj)
}

func (c context) UpdateService(key hsk.Key, obj Service) error {
	return c.Services.Update(key, obj)
}

func (c context) GetClothing(key hsk.Key) (Clothing, error) {
	rec, err := c.Clothing.FindByKey(key)

	if err != nil {
		return Clothing{}, err
	}

	return rec.GetValue().(Clothing), nil
}

func (c context) FindClothing(page, size int, profile string) (records.Page, error) {
	return c.Clothing.Find(page, size, byClothingCategory(profile))
}

func (c context) CreateClothing(obj Clothing) (hsk.Key, error) {
	return c.Clothing.Create(obj)
}

func (c context) UpdateClothing(key hsk.Key, obj Clothing) error {
	return c.Clothing.Update(key, obj)
}

func (c context) GetCar(key hsk.Key) (Car, error) {
	rec, err := c.Cars.FindByKey(key)

	if err != nil {
		return Car{}, err
	}

	return rec.GetValue().(Car), nil
}

func (c context) FindLatestCars(page, size int, profile string) (records.Page, error) {
	return c.Cars.Find(page, size, byProfile(profile))
}

func (c context) CreateCar(obj Car) (hsk.Key, error) {
	return c.Cars.Create(obj)
}

func (c context) UpdateCar(key hsk.Key, obj Car) error {
	return c.Cars.Update(key, obj)
}

func (c context) GetPart(key hsk.Key) (Part, error) {
	rec, err := c.Parts.FindByKey(key)

	if err != nil {
		return Part{}, err
	}

	return rec.GetValue().(Part), nil
}

func (c context) FindLatestParts(page, size int, profile string) (records.Page, error) {
	return c.Parts.Find(page, size, byProfile(profile))
}

func (c context) CreatePart(obj Part) (hsk.Key, error) {
	return c.Parts.Create(obj)
}

func (c context) UpdatePart(key hsk.Key, obj Part) error {
	return c.Parts.Update(key, obj)
}

func (c context) GetProperty(key hsk.Key) (Property, error) {
	rec, err := c.Properties.FindByKey(key)

	if err != nil {
		return Property{}, err
	}

	return rec.GetValue().(Property), nil
}

func (c context) FindLatestProperties(page, size int, profile string) (records.Page, error) {
	return c.Properties.Find(page, size, byProfile(profile))
}

func (c context) CreateProperty(obj Property) (hsk.Key, error) {
	return c.Properties.Create(obj)
}

func (c context) UpdateProperty(key hsk.Key, obj Property) error {
	return c.Properties.Update(key, obj)
}
