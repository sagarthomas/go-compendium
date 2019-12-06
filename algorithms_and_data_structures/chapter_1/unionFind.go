package main

type UF struct {
	id    []int
	count int
}

func (uf *UF) Union(p int, q int) {

}

func (uf UF) Find(p int) int {

}

func (uf UF) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf UF) Count() int {
	return uf.count
}
