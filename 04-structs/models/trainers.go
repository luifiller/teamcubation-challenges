package models

import (
	"errors"
)

type Trainer struct {
	Name  string
	Party []Pokemon
}

func (t *Trainer) AddToParty(p Pokemon) error {
	if len(t.Party) >= 6 {
		return errors.New("o time do treinador já está cheio")
	}

	t.Party = append(t.Party, p)

	return nil
}
