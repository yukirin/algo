package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	x, y := make([]float64, n), make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = nextFloat()
		y[i] = nextFloat()
	}

	dist := make([][]float64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dist[i][j] = math.Hypot(x[i]-x[j], y[i]-y[j])
		}
	}

	inf := 1e100
	dp := make([][]float64, 1<<n)
	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], inf)
		}
	}
	dp[1][0] = 0

	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == inf {
				continue
			}

			for k := 0; k < n; k++ {
				if (i>>k)%2 == 1 {
					continue
				}

				nextI := i | (1 << k)
				nextD := dp[i][j] + dist[j][k]
				if nextD < dp[nextI][k] {
					dp[nextI][k] = nextD
				}
			}
		}
	}

	all := (1 << n) - 1
	ret := inf

	for i := 0; i < n; i++ {
		if dp[all][i] == inf {
			continue
		}

		temp := dp[all][i] + dist[i][0]
		if temp < ret {
			ret = temp
		}
	}
	fmt.Println(ret)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextLine())
	return i
}

func nextInt64() int64 {
	i, _ := strconv.ParseInt(nextLine(), 10, 64)
	return i
}

func nextUint64() uint64 {
	i, _ := strconv.ParseUint(nextLine(), 10, 64)
	return i
}

func nextFloat() float64 {
	f, _ := strconv.ParseFloat(nextLine(), 64)
	return f
}

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func factor(n uint64) []uint64 {
	if n <= 1 {
		return []uint64{n}
	}

	ps := make([]uint64, 0, 100)

	for i := uint64(2); i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			ps = append(ps, i)
		}
	}
	if n > 1 {
		ps = append(ps, n)
	}
	return ps
}

func gcd(a, b uint64) uint64 {
	for ; b != 0; b, a = a%b, b {
	}
	return a
}

func exGcd(x, y int64) (a, b, c int64) {
	if x <= 0 || y <= 0 {
		return
	}

	r0, r1 := x, y
	a0, a1 := int64(1), int64(0)
	b0, b1 := int64(0), int64(1)

	for r1 > 0 {
		q1, r2 := r0/r1, r0%r1
		a2, b2 := a0-q1*a1, b0-q1*b1
		r0, r1 = r1, r2
		a0, a1 = a1, a2
		b0, b1 = b1, b2
	}
	a, b, c = a0, b0, r0
	return
}

func chinese(a1, m1, a2, m2 int64) int64 {
	a, _, c := exGcd(m1, m2)
	if c != 1 {
		return -1
	}

	x := a1 + (a2-a1)*a*m1
	for x < 0 {
		x += m1 * m2
	}
	return x
}

func lcm(a, b uint64) uint64 {
	return (a * b) / gcd(a, b)
}

func round(f float64, place int) float64 {
	shift := math.Pow10(place)
	return math.Floor(f*shift+.5) / shift
}

func round1(f float64) float64 {
	return math.Floor(f + .5)
}
