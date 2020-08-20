package core

import "github.com/louisevanderlith/husk"

type Part struct {
	StockItem
	Number string
}

func (o Part) Valid() error {
	return husk.ValidateStruct(o)
}

func (c Part) Create() (husk.Recorder, error) {
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
