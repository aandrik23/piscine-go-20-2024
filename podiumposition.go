package piscine

func PodiumPosition(podium [][]string) [][]string {
	if len(podium) > 1 {
		swapped := true
		for swapped {
			swapped = false
			for i := 0; i < len(podium)-1; i++ {
				if podium[i][0] > podium[i+1][0] {
					podium[i], podium[i+1] = podium[i+1], podium[i]
					swapped = true
				}
			}
		}
	}
	return podium
}
