package main

import (
	"time"
)

func say(s string) {

	for i := 0; i < 5; i++ {
		time.sleep(100 * time.Milli)
	}
}

func main() {

}
