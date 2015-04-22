func someFunc(name string) bool { }

func someFunc(name string, title string) bool { }

func someFunc(name, title string) bool { }

func someFunc(name, title string) (bool, int) { }

func someMessyFunc(
	name string,
	age int,
	title string,
	source_id int,
) (bool, int, string) {
}

func funcWithDefer(s socket) {
	defer s.Close()
	// do stuff
}
