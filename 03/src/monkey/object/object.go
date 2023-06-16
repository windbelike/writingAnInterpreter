package object

import (
	"fmt"
)

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
)

type ObjectType string

// everything in monkey language is an object
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer object
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}
