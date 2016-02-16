package traytor

// Scene contains all the information for a scene
type Scene struct {
	camera    *AnyCamera
	materials []Material
	mesh      Mesh
}

// Intersection represents a point on a surface struck by a ray
type Intersection struct {
	Point     *Vec3
	Incoming  *Ray
	distance  float64
	u, v      float64
	Normal    *Vec3
	SurfaceOx *Vec3
	SurfaceOy *Vec3
}
