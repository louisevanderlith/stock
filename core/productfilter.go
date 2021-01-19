package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
)

type productFilter func(obj Product) bool

func (f productFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Product))
}

func byCategory(categoryKey hsk.Key) productFilter {
	return func(obj Product) bool {
		return obj.CategoryKey.Compare(categoryKey) == 0
	}
}

func byShortName(shortname string) productFilter {
	return func(obj Product) bool {
		return obj.ShortName == shortname
	}
}

func byItemKey(itemKey hsk.Key) productFilter {
	return func(obj Product) bool {
		for i := 0; i < len(obj.ItemKeys); i++ {
			if obj.ItemKeys[i].Compare(itemKey) == 0 {
				return true
			}
		}

		return false
	}
}

func byProductSearch(in Product) productFilter {
	var mustPass []productFilter

	if in.CategoryKey.Compare(keys.CrazyKey()) != 0 {
		mustPass = append(mustPass, byCategory(in.CategoryKey))
	}

	if len(in.ShortName) > 0 {
		mustPass = append(mustPass, byShortName(in.ShortName))
	}

	return func(obj Product) bool {
		for _, test := range mustPass {
			if !test(obj) {
				return false
			}
		}

		return true
	}
}
