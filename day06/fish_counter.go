package main

type FishCounter [9]int

func NewFishCounter(fishes []int) FishCounter {
	var counter FishCounter

	for _, num := range fishes {
		counter[num]++
	}

	return counter
}

func (c *FishCounter) Iterate() FishCounter {
	var counter FishCounter

	for i := range c {
		if i == 0 {
			counter[8] += c[i]
			counter[6] += c[i]
		} else {
			counter[i-1] += c[i]
		}
	}

	return counter
}

func (c *FishCounter) TotalFish() int {
	total := 0
	for _, num := range c {
		total += num
	}

	return total
}
