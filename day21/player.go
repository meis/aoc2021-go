package main

type Player struct {
	position int
	score    int
}

func (p *Player) move(distance int) {
	(*p).position = (((*p).position + distance - 1) % 10) + 1
	(*p).score += (*p).position
}
