```
package main

import (
	"errors"
	"fmt"
)

type Music struct {
	Name, Artist, Path, Type string
}
type MusicManager struct {
	musics []MusicEntry//这里的musics类型没有定义，在NewMusicManager执行中动态创建可以吗？
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}//这里动态创建了MusicEntry这个slice
}
func (self *MusicManager) Len() int {
	return len(self.musics)
}
func (self *MusicManager) Get(index int) (music *Musicentry, err error) {
	if index < 0 || index >= len(self.musics) {
		return nil, errors.New("index of out range")
	}
	return &self.musics[index], nil
}
func (self *MusicManager) Find(name string) *MusicEntry {
	if len(self.musics) == 0 {
		return nil
	}
	for _, m := range self.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}
func (self *MusicManager)Add(music *MusicEntry){
	m.musics = append(m.musics, *music)
}
func (self *MusicManager) Remove(index int) *MusicEntry{
	if index < 0 || index >= len(self.musics){
		return nil
	}
	removedMusics:=&self.musics[index]
	if index < len(self.musics - 1){
		self.musics = append(self.musics[:index], self.musics[index + 1:]...)
	}
	elseif index == 0{
		self.musics = make([]MusicEntry,0)
	}
	else{
		self.musics = self.musics[:index-1]
	}
}
```