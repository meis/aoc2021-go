package main

type DeterministicDiceResult struct {
	Player1Points, Player2Points int
	Rolls                        int
}

func PlayDeterministicDice(p1 Player, p2 Player, maxScore int) DeterministicDiceResult {
	dieValue := 0
	rolls := 0
	roll := func() int {
		dieValue++
		rolls++

		return dieValue
	}

	for {
		p1.move(roll() + roll() + roll())
		if p1.score >= maxScore {
			break
		}
		p2.move(roll() + roll() + roll())
		if p2.score >= maxScore {
			break
		}
	}

	return DeterministicDiceResult{p1.score, p2.score, rolls}
}
