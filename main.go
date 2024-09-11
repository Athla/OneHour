package main

import (
	lg "github.com/charmbracelet/log"
)

func main() {
	lg.Info("\n\tThis project isn't meant to be used with 'go run .' or similar commands.\n\tOnly with 'go test {dsa}', where dsa is the algorithm or data struct you might want to test.")
}
