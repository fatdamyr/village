package components

//PlayerColor represents the different color piece a player is controlling
type PlayerColor uint

const (
	//Yellow player color
	Yellow PlayerColor = iota
	//Red player color
	Red
	//Blue player color
	Blue
	//White player color
	White
)

// Number of each generation in the players family at the start of the game
const (
	generation1 = 4
	generation2 = 3
	generation3 = 2
	generation4 = 2
)

//Player is a representation of the player state. Holds information about all resources the player has.
type Player struct {
	//Player color
	Color PlayerColor `json:"color"`

	//Bags of Grain
	Grain int `json:"grain"`

	//Life time track location (0-9)
	Lifetime int `json:"lifetime"`

	//Accumulated points
	Prestige int `json:"prestige"`
}

//Setup the player state for the start of the game.
func (p *Player) Setup(color PlayerColor) {
	//Allocate family members
}
