package categories

import "strings"

type Enum int

const (
	Cars Enum = iota
	Clothing
	Spares
	Properties
	Utilities
	Tokens
)

var vals = [...]string{
	"Cars",
	"Clothing",
	"Spares",
	"Properties",
	"Utilities",
	"Tokens",
}

func (r Enum) String() string {
	return vals[r]
}

func GetEnum(name string) Enum {
	var result Enum

	for k, v := range vals {
		if strings.ToUpper(name) == strings.ToUpper(v) {
			result = Enum(k)
			break
		}
	}

	return result
}
