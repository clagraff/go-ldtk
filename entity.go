package goldtk

import (
	"goldtk/maybe"
	"goldtk/quicktype"
)

type Entity interface {
	Identifier() Identifier
	Iid() InstanceIdentifier

	Tags() []string
	Tile() maybe.Value[Tile]
	Fields() []Field

	WorldX() maybe.Value[int64]
	WorldY() maybe.Value[int64]

	// TODO something something layer offsets
	LocalX() int
	LocalY() int

	Height() int64
	Width() int64
	Size() (width, height int64)
}

type entity struct {
	inst quicktype.EntityInstance
}

func (e entity) Identifier() Identifier {
	return Identifier(e.inst.Identifier)
}

func (e entity) Iid() InstanceIdentifier {
	return InstanceIdentifier(e.inst.Iid)
}

func (e entity) Tags() []string {
	return e.inst.Tags
}

func (e entity) Tile() maybe.Value[Tile] {
	//TODO implement me
	panic("implement me")
}

func (e entity) Fields() []Field {
	//TODO implement me
	panic("implement me")
}

func (e entity) WorldX() maybe.Value[int64] {
	return maybe.From[int64](e.inst.WorldX)
}

func (e entity) WorldY() maybe.Value[int64] {
	return maybe.From[int64](e.inst.WorldY)
}

func (e entity) LocalX() int {
	panic("need to consider layer offset??")
	return int(e.inst.Px[0])
}

func (e entity) LocalY() int {
	panic("need to consider layer offset??")
	return int(e.inst.Px[1])
}

func (e entity) Height() int64 {
	return e.inst.Height
}

func (e entity) Width() int64 {
	return e.inst.Width
}

func (e entity) Size() (width, height int64) {
	return e.inst.Width, e.inst.Height
}

func NewEntity(inst quicktype.EntityInstance) Entity {
	return entity{
		inst: inst,
	}
}

var _ Entity = entity{}
