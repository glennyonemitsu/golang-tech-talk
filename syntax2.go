package main

func main() {
	name := "Glenn"

	sayHi := func() int {
		if name == "Glenn" {
			println("Hello, Glenn the Technomancer!")
			return 1
		} else {
			return 0
		}
	}

	for i, j := 0, sayHi(); i < 8; i, j = i + 1, j * 2 {
		println(j)
	}
}
