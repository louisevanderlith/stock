package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/stock/core/categories"
)

type categoryFilter func(obj Category) bool

func (f categoryFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Category))
}

func byName(name string) categoryFilter {
	return func(obj Category) bool {
		return obj.Name == name
	}
}

func byClient(clientid string) categoryFilter {
	return func(obj Category) bool {
		return obj.ClientID == clientid
	}
}

func byBaseCategory(enum categories.Enum) categoryFilter {
	return func(obj Category) bool {
		return obj.BaseCategory == enum
	}
}

func byOwner(owner hsk.Key) categoryFilter {
	return func(obj Category) bool {
		return obj.OwnerKey.Compare(owner) == 0
	}
}

func byCategorySearch(in Category) categoryFilter {
	var mustPass []categoryFilter

	if in.OwnerKey.Compare(keys.CrazyKey()) != 0 {
		mustPass = append(mustPass, byOwner(in.OwnerKey))
	}

	if len(in.Name) != 0 {
		mustPass = append(mustPass, byName(in.Name))
	}

	if len(in.ClientID) != 0 {
		mustPass = append(mustPass, byClient(in.ClientID))
	}

	if in.BaseCategory > 0 {
		mustPass = append(mustPass, byBaseCategory(in.BaseCategory))
	}

	return func(obj Category) bool {
		for _, test := range mustPass {
			if !test(obj) {
				return false
			}
		}

		return true
	}
}
