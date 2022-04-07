package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	pluginDir = flag.String("plugindir", "default", "flag to set")
)

func main() {
	flag.Parse()
	fmt.Printf("flag.Parsed(): %v\nos.Args: %s\npluginDir: %s\n", flag.Parsed(), os.Args, *pluginDir)
}
