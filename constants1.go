const name = "Glenn"

const (
	flagFoo = 1 << 0
	flagBar = 1 << 1
	flagBaz = 1 << 2
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
