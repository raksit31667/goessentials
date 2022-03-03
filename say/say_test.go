package say

import "testing"

func TestSay(t *testing.T) {
	
	s, _ := Say2("Yo")

	if s != "Yo Hello, Gopher" {
		t.Error("Expected 'Yo Hello Gopher 2' but got", s)
	}
}