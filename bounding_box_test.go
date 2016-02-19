package traytor

import (
	"fmt"
	"testing"
)

func ExampleBoundingBox_AddPoint() {
	box := NewBoundingBox()

	box.AddPoint(NewVec3(-1, -1, -1))
	box.AddPoint(NewVec3(0, 5, -0.5))
	box.AddPoint(NewVec3(1, 0, 2))

	fmt.Printf("min: %s, max: %s\n", &box.MinVolume, &box.MaxVolume)

	// Output:
	// min: (-1, -1, -1), max: (1, 5, 2)
}

func ExampleBoundingBox_Inside() {
	box := &BoundingBox{
		MinVolume: *NewVec3(-1, -1, -1),
		MaxVolume: *NewVec3(1, 1, 1),
	}

	fmt.Printf("0, 0, 0: %v\n", box.Inside(NewVec3(0, 0, 0)))
	fmt.Printf("0, 0, 2: %v\n", box.Inside(NewVec3(0, 0, 2)))

	// Output:
	// 0, 0, 0: true
	// 0, 0, 2: false
}

func ExampleBoundingBox_Intersect() {
	box := &BoundingBox{
		MinVolume: *NewVec3(-1, -1, -1),
		MaxVolume: *NewVec3(1, 1, 1),
	}

	ray1 := &Ray{
		Start:     *NewVec3(0, 0, 0),
		Direction: *NewVec3(1, 0, 0),
	}

	ray2 := &Ray{
		Start:     *NewVec3(-5, 0, 0.5),
		Direction: *NewVec3(1, 0, 0),
	}

	ray3 := &Ray{
		Start:     *NewVec3(-5, 0, 0.5),
		Direction: *NewVec3(-1, 0, 0),
	}

	ray4 := &Ray{
		Start:     *NewVec3(-5, 0, 0),
		Direction: *NewVec3(1, 0, 5),
	}

	fmt.Printf("ray 1: %v\n", box.Intersect(ray1))
	fmt.Printf("ray 2: %v\n", box.Intersect(ray2))
	fmt.Printf("ray 3: %v\n", box.Intersect(ray3))
	fmt.Printf("ray 4: %v\n", box.Intersect(ray4))

	// Output:
	// ray 1: true
	// ray 2: true
	// ray 3: false
	// ray 4: false
}

func TestIntersectBoundingBoxZeroVolume(t *testing.T) {
	box := NewBoundingBox()
	box.AddPoint(NewVec3(0, 0, 0))
	box.AddPoint(NewVec3(1, 0, 0))
	box.AddPoint(NewVec3(0, 1, 0))
	box.AddPoint(NewVec3(1, 1, 0))

	ray := &Ray{
		Start:     *NewVec3(0.5, 0.5, 1),
		Direction: *NewVec3(0, 0, -1),
	}

	if !box.Intersect(ray) {
		t.Error("Ray must intersect box")
	}
}

func TestBoundingBoxSplit(t *testing.T) {
	box := &BoundingBox{
		MinVolume: *NewVec3(-3.18, -3.96, -1.77),
		MaxVolume: *NewVec3(4.58, 1.74, 4.56),
	}

	ray := &Ray{
		Start:     *NewVec3(4.49, -7.3, 5.53),
		Direction: *NewVec3(-0.60, 0.5, -0.62),
	}

	if !box.Intersect(ray) {
		t.Error("ray should intersect big box")
	}

	left, right := box.Split(2, 1.39)

	asserEqualVectors(t, NewVec3(-3.18, -3.96, -1.77), &left.MinVolume)
	asserEqualVectors(t, NewVec3(-3.18, -3.96, 1.39), &right.MinVolume)
	asserEqualVectors(t, NewVec3(4.58, 1.74, 1.39), &left.MaxVolume)
	asserEqualVectors(t, NewVec3(4.58, 1.74, 4.56), &right.MaxVolume)

	if box.IntersectWall(2, 1.39, ray) {
		t.Error("ray shouldn't intersect the wall")
	}

	if right.Intersect(ray) {
		t.Error("ray shouldn't intersect right")
	}

	if !left.Intersect(ray) {
		t.Error("ray should intersect left")
	}
}
