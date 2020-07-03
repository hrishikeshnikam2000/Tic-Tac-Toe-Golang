package components

type Player struct {
	Mark string
	Name string
}

func NewPlayer(name string, mark string) *Player {
	newPlayer := &Player{Mark: mark, Name: name}
	return newPlayer
}
