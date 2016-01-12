package traytor

import "math"
import "fmt"

// Vec3 is a 3 dimensional vector
type Vec3 struct {
	X, Y, Z float64
}

// NewVec3 returns new 3 dimensional vector
func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{X: x, Y: y, Z: z}
}

// ToZero makes all the dimentsions of the vector zero
func (v *Vec3) ToZero() {
	v.X, v.Y, v.Z = 0, 0, 0
}

//Length return the lenght of a vector
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared returns the square of the length of a vector
func (v *Vec3) LengthSquared() float64 {
	return (v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Scale multiplies all the dimension of the vector by the given multiplier
func (v *Vec3) Scale(multiplier float64) {
	v.X *= multiplier
	v.Y *= multiplier
	v.Z *= multiplier
}

// Add takes another vector and adds its dimensions to those of the given vector
func (v *Vec3) Add(other *Vec3) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

// Scaled returns a new Vec3 which is the product of the multiplication of the given vector and the multiplier
func (v *Vec3) Scaled(multiplier float64) *Vec3 {
	return NewVec3(
		v.X*multiplier,
		v.Y*multiplier,
		v.Z*multiplier,
	)
}

// Normalise sets the length to the given vector to 1
func (v *Vec3) Normalise() {
	v.SetLength(1.0)
}

// Normalised returns a new vector which length is 1
func (v *Vec3) Normalised() *Vec3 {
	normalisedVector := NewVec3(v.X, v.Y, v.Z)
	normalisedVector.SetLength(1.0)
	return normalisedVector
}

//Reflected makes the given vector equal to its reflected vector by the normal
func (v *Vec3) Reflect(normal *Vec3) {
	v.Normalise()
	v.Add(normal.Scaled(2 * DotProduct(normal, v.Negative())))
	v.Normalise()
}

//Reflected returns a new Vec3 witch is the reflected vector of the given vector by the normal
func (v *Vec3) Reflected(normal *Vec3) *Vec3 {
	reflectedVector := NewVec3(v.X, v.Y, v.Z)
	reflectedVector.Normalise()
	reflectedVector.Add(normal.Scaled(2 * DotProduct(normal, reflectedVector.Negative())))
	return reflectedVector.Normalised()
}

//Negative returns the opposite of the given vector
func (v *Vec3) Negative() *Vec3 {
	return v.Scaled(-1)
}

//Negate makes the given vector equal to its opposite vector
func (v *Vec3) Negate() {
	v.Scale(-1)
}

//SetLength makes the lenght of the vector equal to the given newLength
func (v *Vec3) SetLength(newLength float64) {
	v.Scale(newLength / v.Length())
}

//String returns the string representation of the vector in the form of (x, y, z)
func (v *Vec3) String() string {
	return fmt.Sprintf("(%.3g, %.3g, %.3g)", v.X, v.Y, v.Z)
}

//AddVector returns a new vector which is the sum of the two given vectors
func AddVectors(first, second *Vec3) *Vec3 {
	return NewVec3(first.X+second.X, first.Y+second.Y, first.Z+second.Z)
}

//MinusVectors returns a new vector which is the difference between the two given vectors
func MinusVectors(first, second *Vec3) *Vec3 {
	return AddVectors(first, second.Negative())
}

// DotProduct returns a float64 number which is the product of the two given vectors
func DotProduct(first, second *Vec3) float64 {
	return (first.X*second.X + first.Y*second.Y + first.Z*second.Z)
}

// CrossProduct returns a new Vec3 which is the cross product of the two given vectors
func CrossProduct(first, second *Vec3) *Vec3 {
	return NewVec3(
		first.Y*second.Z-first.Z*second.Y,
		first.Z*second.X-first.X*second.Z,
		first.X*second.Y-first.Y*second.X,
	)
}

//FaceForward return a new Vec3 which is the normal vector directed so that the ray is facing forward
func (normal *Vec3) FaceForward(ray *Vec3) *Vec3 {
	if DotProduct(ray, normal) < 0 {
		return normal
	}
	return normal.Negative()
}

// Ray is defined by its start, direction and depth which indicates how many materials it has passed through
type Ray struct {
	start     Vec3
	direction Vec3
	depth     int
}

// NewRay returns new ray
func NewRay(start Vec3, direction Vec3, depth int) *Ray {
	return &Ray{start: start, direction: direction, depth: depth}
}