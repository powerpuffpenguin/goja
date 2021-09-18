package main

import (
	"generation/internal/number"
	"log"
)

type Generation struct {
	dir string
}

func NewGeneration(dir string) *Generation {
	return &Generation{
		dir: dir,
	}
}
func (g *Generation) Serve() {
	dir := g.dir
	log.Println(`work on`, dir)
	number.NewTemplate(dir).Serve()
}
