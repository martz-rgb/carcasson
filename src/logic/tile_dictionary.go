package logic

type TileID string

var TilesIDs = struct {
	Monastery     TileID
	MonasteryRoad TileID

	// 4-sides
	TownFull TileID

	// 3-sides
	TownCenter     TileID
	TownCenterRoad TileID

	// 2-sides
	TownEdge     TileID
	TownEdgeRoad TileID
	TownOpposite TileID
	TownSides    TileID
	TownTube     TileID

	// 1-side
	TownWall         TileID
	TownWallLeft     TileID
	TownWallRight    TileID
	TownWallCenter   TileID
	TownWallJunction TileID

	// roads only
	// 4-sides
	RoadJunctionAll TileID

	// 3-sides
	RoadJunction TileID

	// 2-sides
	RoadStraight TileID
	RoadCurve    TileID
}{
	Monastery:        "monastery",
	MonasteryRoad:    "monastery_road",
	TownFull:         "town_full",
	TownCenter:       "town_center",
	TownCenterRoad:   "town_center_road",
	TownEdge:         "town_edge",
	TownEdgeRoad:     "town_edge_road",
	TownOpposite:     "town_opposite",
	TownSides:        "town_sides",
	TownTube:         "town_tube",
	TownWall:         "town_wall",
	TownWallLeft:     "town_wall_left",
	TownWallRight:    "town_wall_right",
	TownWallCenter:   "town_wall_center",
	TownWallJunction: "town_wall_junction",
	RoadJunctionAll:  "road_junction_all",
	RoadJunction:     "road_junction",
	RoadStraight:     "road_straight",
	RoadCurve:        "road_curve",
}

var Tiles = map[TileID]Tile{
	TilesIDs.TownWallCenter: {
		Center: Terrains.Road,
		Down:   Terrains.Meadow,
		Right:  Terrains.Road,
		Up:     Terrains.Town,
		Left:   Terrains.Road,
	},
	TilesIDs.Monastery: {
		Center: Terrains.Meadow,
		Down:   Terrains.Meadow,
		Right:  Terrains.Meadow,
		Up:     Terrains.Meadow,
		Left:   Terrains.Meadow,

		Monastery: true,
	},
	TilesIDs.MonasteryRoad: {
		Center: Terrains.Road,
		Down:   Terrains.Road,
		Right:  Terrains.Meadow,
		Up:     Terrains.Meadow,
		Left:   Terrains.Meadow,

		Monastery: true,
	},
	TilesIDs.TownFull: {
		Center: Terrains.Town,
		Down:   Terrains.Town,
		Right:  Terrains.Town,
		Up:     Terrains.Town,
		Left:   Terrains.Town,
	},
	TilesIDs.TownCenter: {
		Center: Terrains.Town,
		Down:   Terrains.Meadow,
		Right:  Terrains.Town,
		Up:     Terrains.Town,
		Left:   Terrains.Town,
	},
	TilesIDs.TownCenterRoad: {
		Center: Terrains.Town,
		Down:   Terrains.Road,
		Right:  Terrains.Town,
		Up:     Terrains.Town,
		Left:   Terrains.Town,
	},
	TilesIDs.TownEdge: {
		Center: Terrains.Meadow,
		Down:   Terrains.Meadow,
		Right:  Terrains.Meadow,
		Up:     Terrains.Town,
		Left:   Terrains.Town,
	},
	TilesIDs.TownEdgeRoad: {
		Center: Terrains.Road,
		Down:   Terrains.Road,
		Right:  Terrains.Road,
		Up:     Terrains.Town,
		Left:   Terrains.Town,
	},
	TilesIDs.TownOpposite: {
		Center: Terrains.Meadow,
		Down:   Terrains.Town,
		Right:  Terrains.Meadow,
		Up:     Terrains.Town,
		Left:   Terrains.Meadow,
	},
	TilesIDs.TownSides: {
		Center: Terrains.Meadow,
		Down:   Terrains.Meadow,
		Right:  Terrains.Meadow,
		Up:     Terrains.Town,
		Left:   Terrains.Town,
	},
	TilesIDs.TownTube: {
		Center: Terrains.Town,
		Down:   Terrains.Meadow,
		Right:  Terrains.Town,
		Up:     Terrains.Meadow,
		Left:   Terrains.Town,
	},
	TilesIDs.TownWall: {
		Center: Terrains.Meadow,
		Down:   Terrains.Meadow,
		Right:  Terrains.Meadow,
		Up:     Terrains.Town,
		Left:   Terrains.Meadow,
	},
	TilesIDs.TownWallLeft: {
		Center: Terrains.Road,
		Down:   Terrains.Road,
		Right:  Terrains.Meadow,
		Up:     Terrains.Town,
		Left:   Terrains.Road,
	},
	TilesIDs.TownWallRight: {
		Center: Terrains.Road,
		Down:   Terrains.Road,
		Right:  Terrains.Road,
		Up:     Terrains.Town,
		Left:   Terrains.Meadow,
	},
	TilesIDs.TownWallCenter: {
		Center: Terrains.Road,
		Down:   Terrains.Meadow,
		Right:  Terrains.Road,
		Up:     Terrains.Town,
		Left:   Terrains.Road,
	},
	TilesIDs.TownWallJunction: {
		Center: Terrains.Meadow,
		Down:   Terrains.Road,
		Right:  Terrains.Road,
		Up:     Terrains.Town,
		Left:   Terrains.Road,
	},
	TilesIDs.RoadJunctionAll: {
		Center: Terrains.Meadow,
		Down:   Terrains.Road,
		Right:  Terrains.Road,
		Up:     Terrains.Road,
		Left:   Terrains.Road,
	},
	TilesIDs.RoadJunction: {
		Center: Terrains.Meadow,
		Down:   Terrains.Road,
		Right:  Terrains.Road,
		Up:     Terrains.Meadow,
		Left:   Terrains.Road,
	},
	TilesIDs.RoadStraight: {
		Center: Terrains.Road,
		Down:   Terrains.Road,
		Right:  Terrains.Meadow,
		Up:     Terrains.Road,
		Left:   Terrains.Meadow,
	},
	TilesIDs.RoadCurve: {
		Center: Terrains.Road,
		Down:   Terrains.Road,
		Right:  Terrains.Meadow,
		Up:     Terrains.Meadow,
		Left:   Terrains.Road,
	},
}
