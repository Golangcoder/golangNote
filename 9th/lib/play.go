package lib

import (
	"fmt"
)

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player
	switch mtype {
	case "MP3":
		p = &Mp3Player{}
	// case "WMV":
	// 	p = &WMVPlayer{}
	default:
		fmt.Println("Music Type not supported")
		return
	}
	p.Play(source)
}
