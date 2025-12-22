package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i := 0; i < len(l); i++ {
		if l[i].Name == name {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(reader io.Reader) (league []Player, err error) {
	err = json.NewDecoder(reader).Decode(&league)
	if err != nil {
		if err == io.EOF {
			return []Player{}, nil
		}
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return
}
