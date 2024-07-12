package goldtk

import (
	"goldtk/quicktype"
)

type LayerType string

const (
	IntGridLayer LayerType = "IntGrid"
	EntityLayer  LayerType = "Entity"
	TilesLayer   LayerType = "Tiles"
	AutoLayer    LayerType = "Auto"
)

type Layer interface {
	Identifier() Identifier
	Iid() InstanceIdentifier
	LayerDefUid() Uid

	Type() LayerType

	GridWidth() int
	GridHeight() int

	GridSizeInPx() int

	Opacity() float32
	IsVisible() bool

	Tileset() quicktype.TilesetDefinition // TODO need a tileset

	AutoLayerTiles()
	Entities() []Entity
	GridTiles() []Tile
	IntGrid() []IntGrid
}

type layer struct {
	inst quicktype.LayerInstance
}

func (l layer) Tileset() quicktype.TilesetDefinition {
	//TODO implement me
	panic("implement me")
}

func (l layer) Identifier() Identifier {
	return Identifier(l.inst.Identifier)
}

func (l layer) Iid() InstanceIdentifier {
	return InstanceIdentifier(l.inst.Identifier)
}

func (l layer) LayerDefUid() Uid {
	return Uid(l.inst.LayerDefUid)
}

func (l layer) Type() LayerType {
	return LayerType(l.inst.Type)
}

func (l layer) GridWidth() int {
	return int(l.inst.CWid)
}

func (l layer) GridHeight() int {
	return int(l.inst.CHei)
}

func (l layer) GridSizeInPx() int {
	return int(l.inst.GridSize)
}

func (l layer) Opacity() float32 {
	return float32(l.inst.Opacity)
}

func (l layer) IsVisible() bool {
	return l.inst.Visible
}

func (l layer) AutoLayerTiles() {
	//TODO implement me
	panic("implement me")
}

func (l layer) Entities() []Entity {
	entities := make([]Entity, 0)
	for _, e := range l.inst.EntityInstances {
		entities = append(entities, NewEntity(e))
	}

	return entities
}

func (l layer) GridTiles() []Tile {
	//TODO implement me
	panic("implement me")
}

func (l layer) IntGrid() []IntGrid {
	//TODO implement me
	panic("implement me")
}

func NewLayer(inst quicktype.LayerInstance) Layer {
	return layer{
		inst: inst,
	}
}

var _ Layer = layer{}
