package core

import "github.com/louisevanderlith/husk"

type stockFilter func(obj StockItem) bool

func (f stockFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(StockItem))
}

//byProfile filter will filter by stock Owner
func byProfile(name string) stockFilter {
	return func(obj StockItem) bool {
		return obj.Profile == name
	}
}