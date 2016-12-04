package utils

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	WIDTH, HEIGHT, NEAR, FAR float32
	position, target mgl32.Vec3
	projection               mgl32.Mat4
	perspective              mgl32.Mat4
}

func (cam *Camera) Update() {
	cam.projection = mgl32.LookAtV(cam.position, cam.target, mgl32.Vec3{0, 1, 0})
	cam.perspective = mgl32.Perspective(mgl32.DegToRad(37.0), cam.WIDTH/cam.HEIGHT, cam.NEAR, cam.FAR)
}

func (cam Camera) Combined() mgl32.Mat4 {
	return cam.perspective.Mul4(cam.projection)
}

func NewCamera(width, height, near, far float32) Camera {
	var cam Camera = Camera{width, height, near, far, mgl32.Vec3{3,3,3}, mgl32.Vec3{0,0,1}, mgl32.Ident4(), mgl32.Ident4()}
	cam.Update();
	
	return cam
}
