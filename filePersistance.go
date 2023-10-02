package main

import (
	"encoding/json"
	"os"
)

func (s *FileStorage) SaveState() error {
	marshalled, err := json.Marshal(s.CurrentState)
	if err != nil {
		return err
	}

	stateFile, err := os.Create(s.Filename)
	if err != nil {
		return err
	}
	defer stateFile.Close()

	stateFile.Write(marshalled)

	return nil
}

func (s *FileStorage) GetState() error {
	stateContent, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(stateContent, &s.CurrentState)
	if err != nil {
		return err
	}

	return nil
}
