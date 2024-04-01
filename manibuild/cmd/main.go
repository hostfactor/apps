package main

import (
	"fmt"
	"manibuild"
	"os"
)

func main() {
	err := manibuild.NewApp().Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
