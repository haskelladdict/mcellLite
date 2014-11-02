// Package geometry contains geometry related functionality
package geometry

// Face describes a triangular face on the model geometry. The orientation of
// the face is defined by the ordering of the vertices and a right hand rule.
type Face [3]Vec3

// Object describes a single geometrical object consisting of a collection
// of triangular faces
type Object struct {
	Faces []*Face
}

// Vec3 is a 3D vector
type Vec3 [3]float64
