package goldtk

import (
	"goldtk/quicktype"
)

//type Serializer[T any] interface {
//	To() T
//	From(T) error
//}

type Identifier string
type InstanceIdentifier string
type Uid int

type World interface {
}

type Reference interface {
	Reference() quicktype.ReferenceToAnEntityInstance

	EntityIid() string
	LayerIid() string
	LevelIid() string
	WorldIid() string

	Entity() interface{} // TODO: Entity here
	Layer() interface{}  // TODO: Entity here
	Level() interface{}  // TODO: Entity here
	World() interface{}  // TODO: Entity here
}

//type Tile interface {
//	LayerDefUid() int64
//	TilesetRectangle() quicktype.TilesetRectangle
//
//	Width() int
//	Height() int
//	Size() (width, height int)
//
//	X() int
//	Y() int
//	Pos() (x, y int)
//
//	Image() image.Image
//}
//
//type tile struct {
//	tsr quicktype.TilesetRectangle
//}
//
//func (t tile) LayerDefUid() int64 {
//	return t.tsr.TilesetUid
//}
//
//func (t tile) TilesetRectangle() quicktype.TilesetRectangle {
//	return t.tsr
//}
//
//func (t tile) Width() int {
//	return int(t.tsr.W)
//}
//
//func (t tile) Height() int {
//	return int(t.tsr.H)
//}
//
//func (t tile) Size() (width, height int) {
//	return t.Width(), t.Height()
//}
//
//func (t tile) X() int {
//	return int(t.tsr.X)
//}
//
//func (t tile) Y() int {
//	return int(t.tsr.Y)
//}
//
//func (t tile) Pos() Point {
//	return Point{
//		X: t.X(),
//		Y: t.Y(),
//	}
//}
//
//func (t tile) Image() image.Image {
//	panic("not implemented")
//}
//
//var _ Tile = tile{}

type IntGrid interface{}
