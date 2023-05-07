//go:generate go run git.rootprojects.org/root/go-gitver/v2
package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	commit  = "0000000"
	version = "0.0.0-pre0+0000000"
	date    = "0000-00-00T00:00:00+0000"
)

func main() {
	if len(os.Args) > 1 && "version" == strings.TrimLeft(os.Args[1], "-") {
		fmt.Printf("Foobar v%s (%s) %s\n", version, commit[:7], date)
	}
}
