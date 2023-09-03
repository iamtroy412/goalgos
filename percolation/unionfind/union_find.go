package unionfind

import "fmt"

type UnionFind struct {
	id []int
	sz []int
}

func New(n int) UnionFind {
    id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return UnionFind{
		id: id,
		sz: make([]int, n),
	}
}

func (u *UnionFind) Root(n int) int {
	i := u.id[n]

	for i != u.id[i] {
		u.id[i] = u.id[u.id[i]]
		i = u.id[i]
	}
	return i
}

func (u *UnionFind) Conneccted(p, q int) bool {
	return u.Root(p) == u.Root(q)
}

func (u *UnionFind) Union(p, q int) {
	i := u.Root(p)
	j := u.Root(q)

	if i == j {
		return
	}

	if u.sz[i] < u.sz[j] {
		u.id[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.id[j] = i
		u.sz[i] += u.sz[j]
	}
}

func (u *UnionFind) Print() {
	for i := range u.id {
		fmt.Printf("%v", i)
	}
}
