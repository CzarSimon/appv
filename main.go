package main

import (
	"fmt"
	"os"

	"github.com/CzarSimon/appv/pkg/cli"
)

func main() {
	cli := cli.New()

	err := cli.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
