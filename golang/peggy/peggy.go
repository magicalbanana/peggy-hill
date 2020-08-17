package peggy

func calculateGears(firstGearSize, divisor int, pegs []int) []int {
	prevRadius := firstGearSize / divisor
	for i := 0; i < len(pegs)-2; i++ {
		width := pegs[i+1] - pegs[i]
		if prevRadius < 0 || prevRadius > (width-1) {
			return []int{-1, -1}
		}
		prevRadius = width - prevRadius
	}

	if firstGearSize > 1 {
		return []int{firstGearSize, divisor}
	}

	return []int{-1, -1}

}

func calculateLastGearSize(pegs []int) int {
	gearSizes := make([]int, 0)
	for i := 0; i < len(pegs)-1; i++ {
		gearSizes = append(gearSizes, pegs[i+1]-pegs[i])
	}

	sign := -1
	if len(pegs)%2 == 0 {
		sign = 1
	}

	result := 0
	counter := 1

	// https://gist.github.com/1lann/be45311db1bd8cbbe6650b0a3e9d1977#gistcomment-2079729
	for i := 0; i < len(gearSizes); i++ {
		if i == len(gearSizes)-1 {
			result = result + (sign * gearSizes[i])
		} else if counter%2 == 0 {
			result -= gearSizes[i]
		} else {
			result += gearSizes[i]
		}
		counter++
	}

	return result
}

// Calculate ...
func Calculate(pegs []int) []int {
	firstGearSize := calculateLastGearSize(pegs) * 2
	divisor := 1
	if len(pegs)%2 == 0 {
		if firstGearSize%3 == 0 {
			firstGearSize /= 3
		} else {
			divisor = 3
		}
	}

	return calculateGears(firstGearSize, divisor, pegs)
}
