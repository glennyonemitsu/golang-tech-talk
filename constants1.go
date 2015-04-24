const name = "Glenn"

const (
	foo = 1 << 0
	bar = 1 << 1
	baz = 1 << 2
)

const (
	foo = 1 << iota
	bar
	baz
)

const (
	_ = iota
	alpha
	bravo
	charlie
	delta
)
