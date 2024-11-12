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
	TilesIDs.Monastery: {
		Dots: [5][5]Terrain{
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
		},

		Monastery: true,
	},
	TilesIDs.MonasteryRoad: {
		Dots: [5][5]Terrain{
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
		},

		Monastery: true,
	},
	TilesIDs.TownFull: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
		},
	},
	TilesIDs.TownCenter: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
			{Town, Town, Land, Town, Town},
			{Town, Land, Land, Land, Town},
		},
	},
	TilesIDs.TownCenterRoad: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
			{Town, Town, Road, Town, Town},
			{Town, Land, Road, Land, Town},
		},
	},
	TilesIDs.TownEdge: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Town, Land, Land, Land, Land},
			{Town, Land, Land, Land, Land},
			{Town, Land, Land, Land, Land},
			{Town, Land, Land, Land, Land},
		},
	},
	TilesIDs.TownEdgeRoad: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Town, Land, Land, Land, Land},
			{Town, Land, Road, Road, Road},
			{Town, Land, Road, Land, Land},
			{Town, Land, Road, Land, Land},
		},
	},
	TilesIDs.TownOpposite: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Town, Town, Town, Town, Town},
		},
	},
	TilesIDs.TownSides: {
		Dots: [5][5]Terrain{
			{Land, Town, Town, Town, Town},
			{Town, Land, Land, Land, Land},
			{Town, Land, Land, Land, Land},
			{Town, Land, Land, Land, Land},
			{Town, Land, Land, Land, Land},
		},
	},
	TilesIDs.TownTube: {
		Dots: [5][5]Terrain{
			{Town, Land, Land, Land, Town},
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
			{Town, Town, Town, Town, Town},
			{Town, Land, Land, Land, Town},
		},
	},
	TilesIDs.TownWall: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
		},
	},
	TilesIDs.TownWallLeft: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Land, Land, Land, Land, Land},
			{Road, Road, Road, Land, Land},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
		},
	},
	TilesIDs.TownWallRight: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Land, Land, Land, Land, Land},
			{Land, Land, Road, Road, Road},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
		},
	},
	TilesIDs.TownWallCenter: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Land, Land, Land, Land, Land},
			{Road, Road, Road, Road, Road},
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
		},
	},
	TilesIDs.TownWallJunction: {
		Dots: [5][5]Terrain{
			{Town, Town, Town, Town, Town},
			{Land, Land, Land, Land, Land},
			{Road, Road, Land, Road, Road},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
		},
	},
	TilesIDs.RoadJunctionAll: {
		Dots: [5][5]Terrain{
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
			{Road, Road, Land, Road, Road},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
		},
	},
	TilesIDs.RoadJunction: {
		Dots: [5][5]Terrain{
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Road, Road, Land, Road, Road},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
		},
	},
	TilesIDs.RoadStraight: {
		Dots: [5][5]Terrain{
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
		},
	},
	TilesIDs.RoadCurve: {
		Dots: [5][5]Terrain{
			{Land, Land, Land, Land, Land},
			{Land, Land, Land, Land, Land},
			{Road, Road, Road, Land, Land},
			{Land, Land, Road, Land, Land},
			{Land, Land, Road, Land, Land},
		},
	},
}
