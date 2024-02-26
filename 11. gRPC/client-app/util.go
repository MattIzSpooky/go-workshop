package main

import (
	"github.com/fatih/color"
	"os"
)

func exitOnError(err error) {
	if err != nil {
		color.Red(err.Error())
		os.Exit(-1)
	}
}
