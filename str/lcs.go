package str

// LCS is longest common subsequence
func LCS(s1, s2 string) string {
	r1, r2 := []rune(s1), []rune(s2)
	t := make([][]int, len(s1)+1)
	for i := range t {
		t[i] = make([]int, len(s2)+1)
	}

	for i, a := range s1 {
		for j, b := range s2 {
			ans, bs := t[i][j+1], t[i+1][j]
			ds := t[i][j]
			if a == b {
				ds++
			}

			if bs > ans {
				ans = bs
			}
			if ds > ans {
				ans = ds
			}

			t[i+1][j+1] = ans
		}
	}

	i, j := len(s1), len(s2)
	index := t[i][j] - 1
	s := make([]rune, t[i][j])
	for t[i][j] != 0 {
		a, b, c := t[i-1][j], t[i][j-1], t[i-1][j-1]

		if a < t[i][j] {
			s[index] = r1[i-1]
			i--
			index--
			continue
		}

		if b < t[i][j] {
			s[index] = r2[j-1]
			j--
			index--
			continue
		}

		if c < t[i][j] {
			s[index] = r1[i-1]
			index--
		}
		i--
		j--
	}
	return string(s)
}
