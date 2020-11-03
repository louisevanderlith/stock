package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type carFilter func(obj Car) bool

func (f carFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Car))
}

//byProfile filter will filter by stock Owner
func byCarProfile(name string) carFilter {
	return func(obj Car) bool {
		return obj.Profile == name
	}
}

type partFilter func(obj Part) bool

func (f partFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Part))
}

//byProfile filter will filter by stock Owner
func byPartProfile(name string) partFilter {
	return func(obj Part) bool {
		return obj.Profile == name
	}
}

type propertyFilter func(obj Property) bool

func (f propertyFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Property))
}

//byProfile filter will filter by stock Owner
func byPropertyProfile(name string) propertyFilter {
	return func(obj Property) bool {
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
