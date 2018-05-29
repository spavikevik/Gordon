package object

import (
	"fmt"
	"gordon/helpers"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	REAL_OBJ    = "REAL"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Real struct {
	Value float64
}

func (r *Real) Inspect() string  { return helpers.Realf(r.Value) }
func (r *Real) Type() ObjectType { return REAL_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

func CastReal(o Object) *Real {
	integer := o.(*Integer)
	return &Real{Value: float64(integer.Value)}
}

func CastInt(o Object) *Integer {
	real := o.(*Real)
	return &Integer{Value: int64(real.Value)}
}
