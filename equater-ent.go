package main

import (
	"image"

	"github.com/disintegration/imaging"
)

func init() {
	imageEquater = &AdvancedImageEquater{}
}

type AdvancedImageEquater struct {
	simple SimpleImageEquater
}

func (s *AdvancedImageEquater) Equal(left, right image.Image) bool {
	flippedRight := imaging.FlipH(right)

	switch {
	case s.simple.Equal(left, right):
	case s.simple.Equal(left, imaging.Rotate90(right)):
	case s.simple.Equal(left, imaging.Rotate180(right)):
	case s.simple.Equal(left, imaging.Rotate270(right)):
	case s.simple.Equal(left, flippedRight):
	case s.simple.Equal(left, imaging.Rotate90(flippedRight)):
	case s.simple.Equal(left, imaging.Rotate180(flippedRight)):
	case s.simple.Equal(left, imaging.Rotate270(flippedRight)):
	default:
		return false
	}

	return true
}
