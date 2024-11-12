package logic

type Terrain uint8

const (
	None Terrain = iota
	Town
	Road
	Land
)

type Tile struct {
	Dots      [5][5]Terrain
	Monastery bool
}
