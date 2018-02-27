package components

import "math/rand"

//Board represents the main game board and its locations
type Board struct {
	//Action Locations
	Harvest,
	Family,
	Crafts,
	Market,
	Travel,
	Council,
	Church,
	Well Location
}

//Location represents a location on the board in which a player can take an action
type Location struct {
	Family    []FamilyMember
	Influence []InfluenceType
}

func (l Location) hasInfluence() bool {
	return (len(l.Influence) != 0)
}

//InfluenceBag represents the cloth bag that contains the influence cubes used to seed the board
type InfluenceBag struct {
	Influence []InfluenceType
}

func (bag *InfluenceBag) shuffle() {
	for i := len(bag.Influence) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		bag.Influence[i], bag.Influence[j] = bag.Influence[j], bag.Influence[i]
	}
}

func (bag *InfluenceBag) setupRound(numPlayers int) {

}
