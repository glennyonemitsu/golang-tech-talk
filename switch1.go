switch name {
	case "Glenn":
		return 1
	case "foo", "bar", "baz":
		return 2
	default:
		return 0
}

switch result := doWork(); result + 100 {
	case 100:
		doThis(result)
	case 200:
		doThat(result)
}

switch {
	case age < 21:
		return "You can't drink here"
	case name = "Glenn":
		return "No Technomancers allowed"
}
