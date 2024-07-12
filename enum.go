package goldtk

import (
	"goldtk/maybe"
	"goldtk/quicktype"
)

// Enum represents one value among many possible options.
type Enum interface {
	// Identifier returns the unique string ident for the enum.
	Identifier() Identifier

	// Uid returns the unique integer ID for the enum.
	Uid() Uid

	// Tags returns a slice of zero-or-more tags associated to the enum.
	Tags() []string

	// Values returns a list of all possible values for the enum.
	Values() []EnumValue
}

type enum struct {
	def    quicktype.EnumDefinition
	values []EnumValue
}

func (e enum) Identifier() Identifier {
	return Identifier(e.def.Identifier)
}

func (e enum) Uid() Uid {
	return Uid(e.def.Uid)
}

func (e enum) Tags() []string {
	return e.def.Tags
}

func (e enum) Values() []EnumValue {
	return e.values
}

func NewEnum(def quicktype.EnumDefinition) Enum {
	values := make([]EnumValue, len(def.Values))
	for _, v := range def.Values {
		values = append(values, NewEnumValue(v))
	}

	return enum{
		def:    def,
		values: values,
	}
}

var _ Enum = enum{}

// EnumValue represents a specific value for an Enum.
type EnumValue interface {
	// Id returns the enum-unique string ident for the current value.
	Id() Identifier

	Color() Color

	Tile() maybe.Value[Tile]
	TileRect() maybe.Value[quicktype.TilesetRectangle]
}

type enumValue struct {
	def quicktype.EnumValueDefinition
}

func (e enumValue) Id() Identifier {
	return Identifier(e.def.ID)
}

func (e enumValue) Color() Color {
	return ColorFromInt64(e.def.Color)
}

func (e enumValue) Tile() maybe.Value[Tile] {
	// TODO return nullable tile using TileRect
	panic("implement me")
}

func (e enumValue) TileRect() maybe.Value[quicktype.TilesetRectangle] {
	return maybe.From[quicktype.TilesetRectangle](e.def.TileRect)
}

func NewEnumValue(def quicktype.EnumValueDefinition) EnumValue {
	return enumValue{
		def: def,
	}
}

var _ EnumValue = enumValue{}
