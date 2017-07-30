//+build ignore

package main

import (
	"fmt"
	"log"

	"github.com/mmogo/popup"
)

func main() {
	popup.Message("", "Hello World!")
	popup.Error(fmt.Errorf("Oh No! A never ending popup!"))
	b := popup.Bool("", "Just kidding...! Would you like to quit?")
	log.Println(b)
}
