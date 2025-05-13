package models

import (
	"errors"
	"fmt"
)

type Pokemon struct {
	Name      string
	Type      []string
	Level     int
	EvolvesTo string
}

func (p *Pokemon) Evolve() error {
	if p.Level < 16 {
		return errors.New("pokemon ainda não pode evoluir")
	}

	fmt.Printf("¡%s evolui para %s!", p.Name, p.EvolvesTo)
	p.Name = p.EvolvesTo

	return nil
}
