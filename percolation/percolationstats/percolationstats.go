package percolationstats

import (
	"math"
	"math/rand"

	"github.com/iamtroy412/goalgos/percolation/percolation"
)

type PercolationStats struct {
	t  int
	xs []float64
}

func New(n, t int) PercolationStats {
    var xs []float64
    var n_site float64 = float64(n * n)

    for x := 0; x < t; x ++ {
        sites := percolation.New(n)
        var count float64 = 0

        for {
            random := rand.Intn(n * n)
            i := random / n + 1
            j := random % n + 1

            if sites.IsOpen(i, j) {
                continue
            }
            sites.Open(i, j)
            count += 1
            if sites.Percolates() {
                break
            }
        }
        xs = append(xs, count / n_site)
    }

    return PercolationStats{
        t: t,
        xs: xs,
    }
}

func (p *PercolationStats) Mean() float64 {
    var sum float64
    for _, x := range p.xs {
        sum += x
    }
    return sum / float64(len(p.xs))
}

func (p *PercolationStats) StdDev() float64 {
    var sd float64
    mean := p.Mean()
    for j := 0; j < len(p.xs); j++ {
        sd += math.Pow(p.xs[j] - mean, 2)
    }
    return math.Sqrt(sd/float64(len(p.xs)))
}

func (p *PercolationStats) ConfidenceHigh() float64 {
    return p.Mean() + 1.96 * p.StdDev() / math.Sqrt(float64(p.t))
}

func (p *PercolationStats) ConfidenceLow() float64 {
    return p.Mean() - 1.96 * p.StdDev() / math.Sqrt(float64(p.t))
}
