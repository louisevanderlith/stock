package core

import "github.com/louisevanderlith/husk"

type serviceFilter func(obj Service) bool

func (f serviceFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(Service))
}

//byOwner filter will filter by stock Owner
func byOwner(ownerKey husk.Key) serviceFilter {
	return func(obj Service) bool {
		return obj.OwnerKey == ownerKey
	}
}
