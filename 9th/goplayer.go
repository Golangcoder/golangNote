package main

import (
	"./lib"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var m *lib.MusicManager = lib.MusicInit()

func handleLibCommand(token []string) {
	switch token[1] {
	case "list":
		for i := 0; i < m.Len(); i++ {
			e, _ := m.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Path, e.Type)
		}
	case "add":
		if len(token) == 5 {
			m.Add(&lib.Music{token[2], token[3], token[4]})
		} else {
			fmt.Println("Useage: lib add <name><Path><Type>")
		}
	case "remove":
		if len(token) == 3 {
			m.Remove(token[2])
		} else {
			fmt.Println("Useage: lib remove <name>")
		}
	default:
		fmt.Println("Uknown lib useage", token[1])
	}
}
func handlePlayCommand(token []string) {
	if len(token) != 2 {
		fmt.Println("Useage: Play <name>")
		return
	}
	f, e := m.Find(token[1])
	if e != nil {
		fmt.Println("the music", token[1], "does not exist")
		return
	}
	lib.Play(f.Name, f.Type)
}
func main() {
	fmt.Println(`
Enter the following commands to control the player
lib list -- view the music list
lib add <name><type><path> -- add music to the list
lib remove <name> -- remove the music from the list
play <name> -- play the music`)

	r := bufio.NewReader(os.Stdin)
	for {
		rawline, _, _ := r.ReadLine()
		line := string(rawline)
		if line == "q" {
			break
		}
		token := strings.Split(line, " ")
		if token[0] == "lib" && len(token) > 1 {
			handleLibCommand(token)
		} else if token[0] == "play" {
			handlePlayCommand(token)
		} else {
			fmt.Println("Uknown command", token)
		}
	}
}
