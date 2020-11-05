package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type categoryFilter func(obj Category) bool

func (f categoryFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Category))
}

func byName(name string) categoryFilter {
	return func(obj Category) bool {
		return obj.Text == name
	}
}

func byClient(clientid string) categoryFilter {
	return func(obj Category) bool {
		return obj.ClientID == clientid
	}
}

func byOwner(owner hsk.Key) categoryFilter {
	return func(obj Category) bool {
		for _, itm := range obj.Items {
			if itm.OwnerKey.Compare(owner) == 0 {
				return true
			}
		}

		return false
	}
}
