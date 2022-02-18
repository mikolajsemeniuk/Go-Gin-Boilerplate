package main

import (
	"check/application"
)

func main() {
	application.
		NewServer().
		Listen()
}
