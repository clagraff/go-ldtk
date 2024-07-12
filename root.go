package goldtk

import (
	"fmt"
	"goldtk/quicktype"
	"io/fs"
)

type Root interface {
	Iid() InstanceIdentifier
	BgColor() Color
	Levels() []Level
	Worlds() []World
	Tilesets() []Tileset
}

type root struct {
	inst quicktype.LdtkJSON
	ts   []Tileset
}

func (r root) Tilesets() []Tileset {
	//TODO implement me
	panic("implement me")
}

func (r root) Iid() InstanceIdentifier {
	return InstanceIdentifier(r.inst.Iid)
}

func (r root) BgColor() Color {
	return ColorFromHex(r.inst.BgColor)
}

func (r root) Levels() []Level {
	lvls := make([]Level, len(r.inst.Levels))
	for i, l := range r.inst.Levels {
		lvls[i] = NewLevel(l)
	}

	return lvls
}

func (r root) Worlds() []World {
	for range r.inst.Worlds {
		// do thing
	}
	//TODO implement me
	panic("implement me")
}

func NewRoot(ldtk quicktype.LdtkJSON, sys fs.FS) (Root, error) {
	tilesets := make([]Tileset, 0)
	for _, def := range ldtk.Defs.Tilesets {
		ts, err := NewTileset(def, sys)
		if err != nil {
			return nil, fmt.Errorf("error creating tileset: %v", err)
		}

		tilesets = append(tilesets, ts)
	}
	return root{
		inst: ldtk,
		ts:   nil,
	}, nil
}

var _ Root = root{}
