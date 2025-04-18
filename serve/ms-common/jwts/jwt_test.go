package jwts

import "testing"

func TestParseToken(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU1NjIwMzAsInRva2VuIjoiMTAwMCJ9.prj-anv12rM-WpkNLqVGUNSBYh16pZZK9DulnU9U7BY"
	ParseToken(tokenStr, "msproject")
}
