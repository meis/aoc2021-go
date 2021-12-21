package main

type Universe struct {
	Player1, Player2 Player
	Player1Turn      bool
}

type DiracDiceResult struct {
	Player1, Player2 int
}

func PlayDiracDice(p1 Player, p2 Player, maxScore int) DiracDiceResult {
	return playDiracDice(Universe{p1, p2, true}, maxScore, make(map[Universe]DiracDiceResult))
}

func playDiracDice(u Universe, maxScore int, cache map[Universe]DiracDiceResult) DiracDiceResult {
	result, found := cache[u]
	if found {
		return result
	}

	result = DiracDiceResult{0, 0}
	p1 := u.Player1
	p2 := u.Player2

	if p1.score >= maxScore {
		result.Player1++
	} else if p2.score >= maxScore {
		result.Player2++
	} else {
		for roll1 := 1; roll1 <= 3; roll1++ {
			for roll2 := 1; roll2 <= 3; roll2++ {
				for roll3 := 1; roll3 <= 3; roll3++ {
					if u.Player1Turn {
						p1 = Player{u.Player1.position, u.Player1.score}
						p1.move(roll1 + roll2 + roll3)
					} else {
						p2 = Player{u.Player2.position, u.Player2.score}
						p2.move(roll1 + roll2 + roll3)

					}
					r := playDiracDice(Universe{p1, p2, !u.Player1Turn}, maxScore, cache)
					result.Player1 += r.Player1
					result.Player2 += r.Player2
				}
			}
		}
	}

	cache[u] = result
	return result
}
