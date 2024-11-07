package logic

var standard_start = TilesIDs.TownWallCenter
var standart_count = 72

var standard_plain_cards = map[TileID]int{
	TilesIDs.Monastery:        4,
	TilesIDs.MonasteryRoad:    2,
	TilesIDs.TownCenter:       3,
	TilesIDs.TownCenterRoad:   1,
	TilesIDs.TownEdge:         3,
	TilesIDs.TownEdgeRoad:     3,
	TilesIDs.TownTube:         1,
	TilesIDs.TownSides:        2,
	TilesIDs.TownOpposite:     3,
	TilesIDs.TownWall:         5,
	TilesIDs.TownWallLeft:     3,
	TilesIDs.TownWallRight:    3,
	TilesIDs.TownWallJunction: 3,
	TilesIDs.TownWallCenter:   4,
	TilesIDs.RoadStraight:     8,
	TilesIDs.RoadCurve:        9,
	TilesIDs.RoadJunction:     4,
	TilesIDs.RoadJunctionAll:  1,
}

var standard_emblem_cards = map[TileID]int{
	TilesIDs.TownFull:       1,
	TilesIDs.TownCenter:     1,
	TilesIDs.TownCenterRoad: 2,
	TilesIDs.TownEdge:       2,
	TilesIDs.TownEdgeRoad:   2,
	TilesIDs.TownTube:       2,
}
