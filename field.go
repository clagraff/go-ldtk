package goldtk

import (
	"fmt"
	"goldtk/quicktype"
	"image/color"
	"os"
)

// Field represents an LDtk property on an entity, layer, level, or world.
type Field interface {
	// Identifier is a unique identifier for the current field.
	Identifier() Identifier

	// FieldValue is an interface for interact with the underlying value of the field.
	Value() FieldValue

	// Type returns the string name of the underlying value.
	Type() string
}

type field struct {
	inst quicktype.FieldInstance
	val  FieldValue
}

func (f field) Identifier() Identifier {
	return Identifier(f.inst.Identifier)
}

func (f field) Value() FieldValue {
	return f.val
}

func (f field) Type() string {
	return f.inst.Type
}

func NewField(inst quicktype.FieldInstance) Field {
	return field{inst, NewFieldValue(inst.Value)}
}

var _ Field = field{}

// FieldValue is an interface to represent the underlying value of a field.
type FieldValue interface {
	Int() int
	Int32() int32
	Int64() int64

	Float64() float64

	Bool() bool

	String() string
	Multilines() Multilines

	ColorHex() string
	Color() color.Color

	FilePath() string
	File() (contents []byte, err error)

	Tile() Tile

	EntityRef() Reference

	Point() Point

	Array() []FieldValue
}

type Multilines struct {
	contents       string
	classification string
}

type File struct {
	Path string
}

// Point represents local coordinates within the level.
type Point struct {
	X int
	Y int
}

type value struct {
	data interface{}
}

func (v value) Int() int {
	var zeroValue int

	if data, ok := v.data.(int); ok {
		return data
	}

	return zeroValue
}

func (v value) Int32() int32 {
	var zeroValue int32

	if data, ok := v.data.(int32); ok {
		return data
	}

	return zeroValue
}

func (v value) Int64() int64 {
	var zeroValue int64

	if data, ok := v.data.(int64); ok {
		return data
	}

	return zeroValue
}

func (v value) Float64() float64 {
	var zeroValue float64

	if data, ok := v.data.(float64); ok {
		return data
	}

	return zeroValue
}

func (v value) Bool() bool {
	var zeroValue bool

	if data, ok := v.data.(bool); ok {
		return data
	}

	return zeroValue
}

func (v value) String() string {
	var zeroValue string

	if data, ok := v.data.(string); ok {
		return data
	}

	return zeroValue
}

func (v value) Multilines() Multilines {
	var zeroValue Multilines

	if data, ok := v.data.(Multilines); ok {
		return data
	}

	return zeroValue
}

func (v value) ColorHex() string {
	if data, ok := v.data.(color.Color); ok {
		return colorToHex(data)
	}

	return "#000000"
}

func (v value) Color() color.Color {
	var zeroValue color.Color

	if data, ok := v.data.(color.Color); ok {
		return data
	}

	return zeroValue
}

func (v value) FilePath() string {
	var zeroValue string

	if data, ok := v.data.(File); ok {
		return data.Path
	}

	return zeroValue
}

func (v value) File() (contents []byte, err error) {
	if data, ok := v.data.(File); ok {
		return os.ReadFile(data.Path) // TODO use a fs?
	}

	return []byte{}, nil
}

func (v value) Tile() Tile {
	var zeroValue Tile

	if data, ok := v.data.(Tile); ok {
		return data
	}

	return zeroValue
}

func (v value) EntityRef() Reference {
	var zeroValue Reference

	if data, ok := v.data.(Reference); ok {
		return data
	}

	return zeroValue
}

func (v value) Point() Point {
	var zeroValue Point

	if data, ok := v.data.(Point); ok {
		return data
	}

	return zeroValue
}

func (v value) Array() []FieldValue {
	var zeroValue []FieldValue

	if data, ok := v.data.([]FieldValue); ok {
		return data
	}

	return zeroValue
}

func NewFieldValue(v any) FieldValue {
	return value{
		data: v,
	}
}

var _ FieldValue = value{}

func colorToHex(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%02x%02x%02x%02x", rgba.R, rgba.G, rgba.B, rgba.A)
}
