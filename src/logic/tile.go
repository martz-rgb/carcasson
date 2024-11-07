package logic

type Terrain uint8

var Terrains = struct {
	Town   Terrain
	Road   Terrain
	Meadow Terrain
}{
	Town:   1,
	Road:   2,
	Meadow: 3,
}

type Tile struct {
	Center, Down, Right, Up, Left Terrain
	Monastery                     bool
}

type TilePosition uint8

var TilePositions = struct {
	Center    TilePosition
	DownLeft  TilePosition
	Down      TilePosition
	DownRight TilePosition
	Right     TilePosition
	RightUp   TilePosition
	Up        TilePosition
	UpLeft    TilePosition
	Left      TilePosition
}{
	Center:    1,
	DownLeft:  2,
	Down:      3,
	DownRight: 4,
	Right:     5,
	RightUp:   6,
	Up:        7,
	UpLeft:    8,
	Left:      9,
}

type TileRotation uint8

var TileRotations = struct {
	Down  TileRotation
	Right TileRotation
	Up    TileRotation
	Left  TileRotation
}{
	Down:  1,
	Right: 2,
	Up:    3,
	Left:  4,
}
