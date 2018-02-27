package components

//Board represents the main game board and its locations
type Board struct {
}

//Location represents a location on the board in which a player can take an action
type Location struct {
	Family    []FamilyMember
	Influence []int
}

func (l Location) hasInfluence() bool {
	return (len(l.Influence) != 0)
}
