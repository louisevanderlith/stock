package core

import "github.com/louisevanderlith/husk"

type Part struct {
	StockItem
	Number string
}

func (o Part) Valid() error {
	return husk.ValidateStruct(&o)
}

func GetPart(key husk.Key) (Part, error) {
	rec, err := ctx.Parts.FindByKey(key)

	if err != nil {
		return Part{}, err
	}

	return rec.Data().(Part), nil
}

func GetLatestParts(page, size int) (husk.Collection, error) {
	return ctx.Parts.Find(page, size, husk.Everything())
}

func (c Part) Create() husk.CreateSet {
	return ctx.Parts.Create(c)
}

func (c Part) Update(key husk.Key) error {
	obj, err := ctx.Parts.FindByKey(key)

	if err != nil {
		return err
	}

	err = obj.Set(c)

	if err != nil {
		return nil
	}

	defer ctx.Parts.Save()
	return ctx.Parts.Update(obj)
}
