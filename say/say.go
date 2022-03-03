package say

func Say() string {
	return "Hello, Gopher"
}

func Say2(s string) (string, int) {
	return s + " Hello, Gopher", 2
}