package main

type UF struct {
	id    []int // all parent links are here, indexed by site
	sz    []int // size of component for each root, indexed by site
	count int
}

// Constructor for our union-find data structure
func NewUF(n int) *UF {
	id := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i // every site is initalized to be its own root
		sz[i] = 1 // every root is initialized to be of size 1
	}
	return &UF{id: id, sz: sz, count: n}
}

func (uf *UF) Union(p int, q int) {
	// Find the components of p and q respectively
	i := uf.Find(p)
	j := uf.Find(q)

	if i == j { // they are in the same component
		return
	}

	if uf.sz[i] < uf.sz[j] {
		uf.id[i] = j         // component i will now point to j, becoming part of the j component
		uf.sz[j] += uf.sz[i] // the size increases accordingly
	} else {
		uf.id[j] = i
		uf.sz[i] += uf.sz[j]
	}
	uf.count-- // decrement the total number of components sice we just did a union between the two

}

func (uf UF) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf UF) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf UF) Count() int {
	return uf.count
}
