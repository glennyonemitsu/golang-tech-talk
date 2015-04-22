package main

var name string

func sayHi() int {
	name = "Glenn"
	if name == "Glenn" {
		println("Hello, Glenn the Technomancer!")
		return 1
	} else {
		return 0
	}
}

func main() {
	var i, j int
	j = sayHi()
	for i = 0; i < 8; i += 1 {
		println(j)
		j *= 2
	}
}
