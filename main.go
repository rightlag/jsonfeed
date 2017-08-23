package jsonfeed

import (
	"flag"
	"fmt"
)

const (
	Version string = "0.1.0"
)

var (
	version string
)

func init() {
	flag.StringVar(&version, "version", Version, "")
}

func main() {
	flag.Parse()
	fmt.Println(version)
}
