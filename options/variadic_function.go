package main

import "fmt"

type FishedHouse struct {
	style                  int
	centralAirConditioning bool
	floorMaterial          string
	wallMaterial           string
}

type Option func(*FishedHouse)

func NewFishedHouse(options ...Option) *FishedHouse {
	h := &FishedHouse{
		style:                  0,
		centralAirConditioning: true,
		floorMaterial:          "wood",
		wallMaterial:           "paper",
	}

	for _, option := range options {
		option(h)
	}

	return h
}

func WithStyle(style int) Option {
	return func(h *FishedHouse) {
		h.style = style
	}
}

func WithFloorMaterial(material string) Option {
	return func(h *FishedHouse) {
		h.floorMaterial = material
	}
}

func WithWallMaterial(material string) Option {
	return func(h *FishedHouse) {
		h.wallMaterial = material
	}
}

func WithCentralAirConditioning(centralAirConditioning bool) Option {
	return func(h *FishedHouse) {
		h.centralAirConditioning = centralAirConditioning
	}
}

func main() {
	fmt.Printf("%+v\n", NewFishedHouse())
	// 很像Java中的Builder模式有木有？
	fmt.Printf("%+v\n", NewFishedHouse(
		WithStyle(1),
		WithWallMaterial("ground-tile"),
		WithFloorMaterial("Paper"),
		WithCentralAirConditioning(true)))
}
