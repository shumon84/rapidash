package rapidash

type IndexTree interface {
	searchEq(key *Value) Leaf
	searchLt(key *Value) []Leaf
	searchGt(key *Value) []Leaf
	searchLte(key *Value) []Leaf
	searchGte(key *Value) []Leaf
	all() []Leaf
	insert(key *Value, value Leaf)
	dump()
	isEmpty() bool
}

type Leaf interface{}
