// Package pkg ...
package pkg

var (
	a = ""  // want `Unicode control character U\+0007`
	b = "" // want `Unicode control characters`
	c = "Test	test"
	d = `T
est`
	e = `ZeroβWidth` // want `Unicode format character U\+200B`
	f = "\u200b"
	g = "π©π½βπ¬" //  want `Unicode control character U\+0007`
	h = "π©π½βπ¬β" // want `Unicode format and control characters`
)
