package main

func operatorsPriority() map[string]int {
	return map[string]int{
		"(": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}
}

func validOperators() []string {
	return []string{
		"+", "-", "/", "*", "^",
	}
}
