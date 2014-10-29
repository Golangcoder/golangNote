package lib

import (
	"fmt"
	"time"
)

type Mp3Player struct {
	stat    int
	process int
}

func (self *Mp3Player) Play(source string) {
	fmt.Println("Playing Music:", source)
	self.process = 0
	for self.process < 100 {
		time.Sleep(time.Millisecond * 100)
		fmt.Print(".")
		self.process += 10
	}
	fmt.Println("\nPlaying Music Finished", source)
}
func Hello() {
	fmt.Println("hello")
}
