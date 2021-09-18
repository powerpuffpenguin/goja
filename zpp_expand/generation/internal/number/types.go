package number

import (
	"strings"
)

type Type struct {
	Name string
	Type string
}

var types = []Type{
	{Name: `Int`, Type: `int`},
	{Name: `Int64`, Type: `int64`},
	{Name: `Int32`, Type: `int32`},
	{Name: `Int16`, Type: `int16`},
	{Name: `Int8`, Type: `int8`},
	{Name: `Uint`, Type: `uint`},
	{Name: `Uint64`, Type: `uint64`},
	{Name: `Uint32`, Type: `uint32`},
	{Name: `Uint16`, Type: `uint16`},
	{Name: `Uint8`, Type: `uint8`},
	{Name: `Float64`, Type: `float64`},
	{Name: `Float32`, Type: `float32`},
}

func (t *Type) Max() bool {
	return t.Type != `int` &&
		t.Type != `uint`
}
func (t *Type) Min() bool {
	return t.Type != `int` &&
		!strings.HasPrefix(t.Type, `u`) &&
		!strings.HasPrefix(t.Type, `f`)
}
func (t *Type) Integer() bool {
	return strings.HasPrefix(t.Type, `u`) || strings.HasPrefix(t.Type, `i`)
}
func (t *Type) Uint() bool {
	return strings.HasPrefix(t.Type, `u`)
}
func (t *Type) BitSize() int {
	result := 0
	switch t.Type {
	case `int`:
		result = 64
	case `int8`:
		result = 8
	case `int16`:
		result = 16
	case `int32`:
		result = 32
	case `int64`:
		result = 64
	case `uint`:
		result = 64
	case `uint8`:
		result = 8
	case `uint16`:
		result = 16
	case `uint32`:
		result = 32
	case `uint64`:
		result = 64
	case `float32`:
		result = 32
	case `float64`:
		result = 64
	}
	return result
}
func (t *Type) Float() bool {
	return strings.HasPrefix(t.Type, `f`)
}
