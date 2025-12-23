package app

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

type FileDatabase interface {
	io.ReadWriteSeeker
	Truncate(size int64) error
}

type FileSystemPlayerStore struct {
	database FileDatabase
	league   League
	encoder  *json.Encoder
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player == nil {
		f.league = append(f.league, Player{Name: name, Wins: 0})
		player = &f.league[len(f.league)-1]
	}
	player.Wins++
	f.database.Seek(0, io.SeekStart)
	f.database.Truncate(0)
	f.encoder.Encode(f.league)
}

func NewFileSystemPlayerStore(database FileDatabase) (*FileSystemPlayerStore, error) {
	database.Seek(0, io.SeekStart)

	league, err := NewLeague(database)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %v", err)
	}

	encoder := json.NewEncoder(database)

	return &FileSystemPlayerStore{
		database: database,
		league:   league,
		encoder:  encoder,
	}, nil
}
