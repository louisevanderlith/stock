package core

import "github.com/louisevanderlith/husk"

type Service struct {
	StockItem
	Url string
}

func (o Service) Valid() error {
	return husk.ValidateStruct(&o)
}

func GetService(key husk.Key) (Service, error) {
	rec, err := ctx.Services.FindByKey(key)

	if err != nil {
		return Service{}, err
	}

	return rec.Data().(Service), nil
}

func GetServices(page, size int) (husk.Collection, error) {
	return ctx.Services.Find(page, size, husk.Everything())
}

func (c Service) Create() husk.CreateSet {
	return ctx.Services.Create(c)
}

func (c Service) Update(key husk.Key) error {
	obj, err := ctx.Services.FindByKey(key)

	if err != nil {
		return err
	}

	err = obj.Set(c)

	if err != nil {
		return nil
	}

	defer ctx.Services.Save()
	return ctx.Services.Update(obj)
}
