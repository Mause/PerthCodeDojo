package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	// "github.com/pquerna/ffjson"
)

type Position struct {
	pos_x, pos_y int
}

type SoilType struct {
	name string
}

type Hydration struct {
	last_watered time.Time
}

type Garden struct {
	max_x, MaxY, min_x, min_y int
	soil                      *SoilType
	hyd                       *Hydration
	plantings                 []interface{}
}

type Planting struct {
	position   *Position
	plant_type string
}

type Tree struct {
	position   *Position
	plant_type string
}

type Flower struct {
	position   *Position
	plant_type string
}

func NewTree(x int, y int) *Tree {
	return &Tree{position: &Position{x, y}}
}

func NewFlower(x int, y int) *Flower {
	return &Flower{position: &Position{x, y}}
}

func NewGarden() *Garden {
	items := make([]interface{}, 3)
	items = append(items, NewTree(0, 0))
	items = append(items, NewFlower(1, 1))
	items = append(items, NewFlower(2, 1))

	return &Garden{
		max_x: 50,
		MaxY:  10,
		min_x: 0,
		min_y: 0,

		plantings: items,
	}
}

func main() {
	b, err := ioutil.ReadFile("garden.json")
	if err != nil {
		panic("Argh!")
	}

	gd := &Garden{}
	err = json.Unmarshal(b, gd)
	if err != nil {
		fmt.Println(err)
		panic("...damn")
	}

	fmt.Printf("Done: %d\n", gd.MaxY)
}
