package percolation

import "github.com/iamtroy412/goalgos/percolation/unionfind"

type Percolation struct {
	n      int
	uf     unionfind.UnionFind
	opened []bool
}

func New(n int) Percolation {
	return Percolation{
		n:      n,
		uf:     unionfind.New(n * n),
		opened: make([]bool, n*n),
	}
}

func (p *Percolation) IndexOf(row, col int) int {
	return (row-1)*p.n + col - 1
}

func (p *Percolation) GetNeighbors(row, col int) []int {
	var neighbors []int

	if row != 1 {
		neighbors = append(neighbors, p.IndexOf(row-1, col))
	}
	if row != p.n {
		neighbors = append(neighbors, p.IndexOf(row+1, col))
	}
	if col != 1 {
		neighbors = append(neighbors, p.IndexOf(row, col-1))
	}
	if col != p.n {
		neighbors = append(neighbors, p.IndexOf(row, col+1))
	}
	return neighbors
}

func (p *Percolation) Open(row, col int) {
	index := p.IndexOf(row, col)
	p.opened[index] = true

	for _, i := range p.GetNeighbors(row, col) {
		if p.opened[i] {
			p.uf.Union(index, i)
		}
	}
}

func (p *Percolation) IsOpen(row, col int) bool {
	return p.opened[p.IndexOf(row, col)]
}

func (p *Percolation) IsFull(row, col int) bool {
	index := p.IndexOf(row, col)
	for column := 1; column <= p.n; column++ {
		if p.IsOpen(1, column) && p.uf.Conneccted(p.IndexOf(1, column), index) {
			return true
		}
	}
	return false
}

func (p *Percolation) NumOpenSites() int {
	open := 0
	for _, i := range p.opened {
		if i {
			open++
		}
	}
	return open
}

func (p *Percolation) Percolates() bool {
	for col := 1; col <= p.n; col++ {
		if p.IsOpen(p.n, col) && p.IsFull(p.n, col) {
			return true
		}
	}
	return false
}
