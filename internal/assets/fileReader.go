package assets

import (
    "fmt"
    "os"
)

func check (e error) {
	if e != nil {
		panic(e)
	}
}

func printFile (path string) {
	dat, err := os.ReadFile(path)
    check(err)
    fmt.Print(string(dat))
}