package lib

import (
	"errors"
	//"fmt"
)

type Music struct {
	Name, Path, Type string
}
type MusicManager struct {
	musics []Music
}

func MusicInit() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}
func (self *MusicManager) Add(music *Music) {
	self.musics = append(self.musics, *music)
}
func (self *MusicManager) Len() int {
	return len(self.musics)
}
func (self *MusicManager) Find(musicname string) (index *Music, err error) {
	if self.Len() == 0 {
		return nil, errors.New("Music List is Empty")
	}
	for _, m := range self.musics {
		if musicname == m.Name {
			return &m, nil
		}
	}
	return nil, errors.New("Music not found")
}
func (self *MusicManager) Remove(musicname string) error {
	if self.Len() == 0 {
		return errors.New("Nothing to remove")
	}
	for i, m := range self.musics {
		if m.Name == musicname {
			self.musics = append(self.musics[:i], self.musics[i+1:]...)
			return nil
		}
	}
	return errors.New("Music Not found")
}
func (self *MusicManager) Get(index int) (music *Music, err error) {
	if index < 0 && index >= len(self.musics) {
		return nil, errors.New("Out of index")
	}
	return &self.musics[index], nil
}
