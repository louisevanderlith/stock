package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type stockFilter func(obj StockItem) bool

func (f stockFilter) Filter(obj hsk.Record) bool {
	return f(obj.Data().(StockItem))
}

//byProfile filter will filter by stock Owner
func byProfile(name string) stockFilter {
	return func(obj StockItem) bool {
		return obj.Profile == name
	}
}

type serviceFilter func(obj Service) bool

func (f serviceFilter) Filter(obj hsk.Record) bool {
	return f(obj.Data().(Service))
}

//byProfile filter will filter by stock Owner
func byServiceProfile(name string) serviceFilter {
	return func(obj Service) bool {
		return obj.Profile == name
	}
}
