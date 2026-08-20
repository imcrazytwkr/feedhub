package jsontree

import "fmt"

type Type int

const (
	TypeNull Type = iota
	TypeObject
	TypeArray
	TypeString
	TypeNumber
	TypeTrue
	TypeFalse
)

func (t Type) String() string {
	switch t {
	case TypeNull:
		return "null"
	case TypeObject:
		return "object"
	case TypeArray:
		return "array"
	case TypeString:
		return "string"
	case TypeNumber:
		return "number"
	case TypeTrue:
		return "true"
	case TypeFalse:
		return "false"
	default:
		panic(fmt.Errorf("BUG: unknown Value type: %d", t))
	}
}
