package object

import ( "fmt")

type Integer struct {
    Value int64
}
func (i *Integer) Inspect() string {
    return fmt.Sprintf("%d", i.Value) 
}

type ObjectType string
type Object interface {
    Type() ObjectType
    Inspect() string
}
