package traytor

import "math/rand"

// Random is a random generator with convenient methods
type Random struct {
	generator *rand.Rand
}

// NewRandom returns a new Random object initialized with the given seed
func NewRandom(seed int64) *Random {
	r := &Random{}
	source := rand.NewSource(seed)
	r.generator = rand.New(source)
	return r
}

// Float01 returns a random float between 0 and 1
func (r *Random) Float01() float64 {
	return r.generator.Float64()
}

// Float0A returns a random float between 0 and a
func (r *Random) Float0A(a float64) float64 {
	return r.Float01() * a
}

// FloatAB returns a random float between 0 and a
func (r *Random) FloatAB(a, b float64) float64 {
	return r.Float0A(b-a) + a
}

// Int640N returns a random int64 within [0..n]
func (r *Random) Int640N(n int64) int64 {
	return r.generator.Int63n(n + 1)
}

// Int320N returns a random int32 within [0..n]
func (r *Random) Int320N(n int32) int32 {
	return r.generator.Int31n(n + 1)
}

// Int0N returns a random int within [0..n]
func (r *Random) Int0N(n int) int {
	return r.generator.Intn(n + 1)
}

// Int64AB returns a random int64 within [a..b]
func (r *Random) Int64AB(a, b int64) int64 {
	return r.Int640N(b-a) + a
}

// Int32AB returns a random int32 within [a..b]
func (r *Random) Int32AB(a, b int32) int32 {
	return r.Int320N(b-a) + a
}

// IntAB returns a random int within [a..b]
func (r *Random) IntAB(a, b int) int {
	return r.Int0N(b-a) + a
}