package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type stockFilter func(obj StockItem) bool

func (f stockFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(StockItem))
}

//byProfile filter will filter by stock Owner
func byProfile(name string) stockFilter {
	return func(obj StockItem) bool {
		return obj.Profile == name
	}
}

type serviceFilter func(obj Service) bool

func (f serviceFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Service))
}

//byProfile filter will filter by stock Owner
func byServiceProfile(name string) serviceFilter {
	return func(obj Service) bool {
		return obj.Profile == name
	}
}

type clothingFilter func(obj Clothing) bool

func (f clothingFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Clothing))
}

//byProfile filter will filter by stock Owner
func byClothingCategory(profile string) clothingFilter {
	return func(obj Clothing) bool {
		return obj.Profile == profile
	}
}
