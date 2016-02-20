package traytor

import (
	"encoding/json"
	"fmt"
	"math"
)

// AnyMaterial implements the Material interface and is deserialiseable from json
type AnyMaterial struct {
	Material
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (m *AnyMaterial) UnmarshalJSON(data []byte) error {
	materialType, err := jsonObjectType(data)
	if err != nil {
		return err
	}

	switch materialType {
	case "emissive":
		material := &EmissiveMaterial{}
		err = json.Unmarshal(data, &material)
		if err != nil {
			return err
		}
		*m = AnyMaterial{material}
	case "lambert":
		material := &LambertMaterial{}
		err = json.Unmarshal(data, &material)
		if err != nil {
			return err
		}
		*m = AnyMaterial{material}
	case "reflective":
		material := &ReflectiveMaterial{}
		err = json.Unmarshal(data, &material)
		if err != nil {
			return err
		}
		*m = AnyMaterial{material}
	case "refractive":
		material := &RefractiveMaterial{}
		err = json.Unmarshal(data, &material)
		if err != nil {
			return err
		}
		*m = AnyMaterial{material}
	default:
		return fmt.Errorf("Unknown material type: '%s'", materialType)
	}

	return nil
}

// Material objects are used to shade surfaces
type Material interface {
	Shade(intersection *Intersection, raytracer *Raytracer) *Colour
}

// EmissiveMaterial acts as a lamp
type EmissiveMaterial struct {
	Colour   *AnySampler
	Strength *AnySampler
}

// Shade returns the emitted colour after intersecting the material
func (m *EmissiveMaterial) Shade(intersection *Intersection, raytracer *Raytracer) *Colour {
	return m.Colour.GetColour(intersection).Scaled(float32(m.Strength.GetFac(intersection)))
}

// LambertMaterial is a simple diffuse material
type LambertMaterial struct {
	Colour *AnySampler
}

// Shade returns the emitted colour after intersecting the material
func (m *LambertMaterial) Shade(intersection *Intersection, raytracer *Raytracer) *Colour {
	randomRayDir := *raytracer.Random.Vec3HemiCos(intersection.Normal)
	randomRayStart := *AddVectors(intersection.Point, intersection.Normal.Scaled(Epsilon))
	ray := &Ray{Start: randomRayStart, Direction: randomRayDir, Depth: intersection.Incoming.Depth + 1}
	colour := raytracer.Raytrace(ray)
	colour.MultiplyBy(m.Colour.GetColour(intersection))
	return colour
}

// ReflectiveMaterial is a reflective material
type ReflectiveMaterial struct {
	Colour    *AnySampler
	Roughness *AnySampler
}

// Shade returns the emitted colour after intersecting the material
func (m *ReflectiveMaterial) Shade(intersection *Intersection, raytracer *Raytracer) *Colour {
	ray := intersection.Incoming
	reflectedRay := &Ray{Depth: ray.Depth + 1}
	reflectedRay.Direction = *ray.Direction.Reflected(intersection.Normal)
	reflectedRay.Start = *AddVectors(intersection.Point, intersection.Normal.Scaled(Epsilon))
	return raytracer.Raytrace(reflectedRay)
}

// RefractiveMaterial is a material for modeling glass, etc
type RefractiveMaterial struct {
	Colour    *AnySampler
	Roughness *AnySampler
	IOR       *AnySampler
}

// Shade returns the emitted colour after intersecting the material
func (m *RefractiveMaterial) Shade(intersection *Intersection, raytracer *Raytracer) *Colour {
	ray := intersection.Incoming
	normal := intersection.Normal
	eta := m.IOR.GetFac(intersection)
	if DotProduct(normal, &ray.Direction) < 0 {
		eta = 1 / eta
	}
	forwardNormal := FaceForward(normal, &ray.Direction)
	refracted := Refract(&ray.Direction, forwardNormal, eta)
	if math.Abs(refracted.Length()) < Epsilon {
		return NewColour(0, 0, 0)
	}
	newRay := &Ray{}
	newRay.Start = *AddVectors(intersection.Point, ray.Direction.Scaled(Epsilon))
	newRay.Direction = *refracted
	newRay.Depth = ray.Depth + 1
	return raytracer.Raytrace(newRay)
}
