package str

// LevenShtein is LevenShtein Distance
func LevenShtein(s1, s2 string) int {
	t := make([][]int, len(s1)+1)
	for i := range t {
		t[i] = make([]int, len(s2)+1)
	}

	for i := range t {
		t[i][0] = i
	}

	for i := range t[0] {
		t[0][i] = i
	}

	for i, a := range s1 {
		for j, b := range s2 {
			cost, del, rep := t[i][j+1]+1, t[i+1][j]+1, t[i][j]
			if a != b {
				rep++
			}

			if del < cost {
				cost = del
			}
			if rep < cost {
				cost = rep
			}
			t[i+1][j+1] = cost
		}
	}

	return t[len(t)-1][len(t[0])-1]
}
