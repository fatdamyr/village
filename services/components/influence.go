package components

//InfluenceType is an enumeration of the different types of influence a player can get in the game
type InfluenceType uint

const (
	//Skill is the orange influence cube
	Skill InfluenceType = iota

	//Persuasivness is the green influence cube
	Persuasivness

	//Faith is the brown influence cube
	Faith

	//Knowledge is the pink influence cube
	Knowledge

	//Plague is the black cube
	Plague
)
