package goldtk

import "goldtk/quicktype"

type Level interface {
	Identifier() Identifier
	Iid() InstanceIdentifier
	Uid() Uid

	BgColor() Color

	PxWidth() int
	PxHeight() int

	WorldX() int
	WorldY() int
	WorldDepth() int

	Layers() []Layer
	Neighbours() []Neighbor

	ExternalRelPath() // TODO

	Fields() []Field
}

type level struct {
	inst quicktype.Level
}

func (l level) Identifier() Identifier {
	return Identifier(l.inst.Identifier)
}

func (l level) Iid() InstanceIdentifier {
	return InstanceIdentifier(l.inst.Identifier)
}

func (l level) Uid() Uid {
	return Uid(l.inst.Uid)
}

func (l level) BgColor() Color {
	return ColorFromHex(l.inst.BgColor)
}

func (l level) PxWidth() int {
	return int(l.inst.PxWid)
}

func (l level) PxHeight() int {
	return int(l.inst.PxHei)
}

func (l level) WorldX() int {
	return int(l.inst.WorldX)
}

func (l level) WorldY() int {
	return int(l.inst.WorldY)
}

func (l level) WorldDepth() int {
	return int(l.inst.WorldDepth)
}

func (l level) Layers() []Layer {
	layers := make([]Layer, 0)
	for _, lyr := range l.inst.LayerInstances {
		layers = append(layers, NewLayer(lyr))
	}

	return layers
}

func (l level) Neighbours() []Neighbor {
	neighbors := make([]Neighbor, 0)
	for _, n := range l.inst.Neighbours {
		neighbors = append(neighbors, NewNeighbor(n))
	}

	return neighbors
}

func (l level) ExternalRelPath() {
	//TODO implement me
	panic("implement me")
}

func (l level) Fields() []Field {
	fields := make([]Field, 0)
	for _, f := range l.inst.FieldInstances {
		fields = append(fields, NewField(f))
	}

	return fields
}

func NewLevel(inst quicktype.Level) Level {
	return level{inst}
}

var _ Level = level{}

type Neighbor interface {
	Dir() LevelDirection
	LevelIid() InstanceIdentifier
}

type neighbor struct {
	inst quicktype.NeighbourLevel
}

func (n neighbor) Dir() LevelDirection {
	return LevelDirection(n.inst.Dir)
}

func (n neighbor) LevelIid() InstanceIdentifier {
	return InstanceIdentifier(n.inst.LevelIid)
}

func NewNeighbor(inst quicktype.NeighbourLevel) Neighbor {
	return neighbor{inst: inst}
}

var _ Neighbor = neighbor{}

type LevelDirection string

const (
	North     LevelDirection = "n"
	NorthEast LevelDirection = "ne"
	East      LevelDirection = "e"
	SouthEast LevelDirection = "se"
	South     LevelDirection = "s"
	SouthWest LevelDirection = "sw"
	West      LevelDirection = "w"
	NorthWest LevelDirection = "nw"

	Above LevelDirection = "<"
	Below LevelDirection = ">"

	Overlap LevelDirection = "o"
)
