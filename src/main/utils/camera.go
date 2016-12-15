package utils

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	ASPECT, NEAR, FAR float32
	position, target, top mgl32.Vec3
	projection               mgl32.Mat4
	perspective              mgl32.Mat4
}

func (cam *Camera) Update() {
	cam.projection = mgl32.LookAtV(cam.position, cam.target,cam.top)
	cam.perspective = mgl32.Perspective(mgl32.DegToRad(37.0), cam.ASPECT, cam.NEAR, cam.FAR)
}

func (cam *Camera) SetPosition(pos mgl32.Vec3) {
	cam.position = pos;
}

func (cam *Camera) SetTarget(target mgl32.Vec3) {
	cam.target = target;
}

func (cam *Camera) SetTop(top mgl32.Vec3) {
	cam.top = top;
}

func (cam Camera) Combined() mgl32.Mat4 {
	return cam.perspective.Mul4(cam.projection)
}

func NewCamera(aspect, near, far float32) Camera {
	var cam Camera = Camera{aspect, near, far, mgl32.Vec3{0,0,0}, mgl32.Vec3{0,0,0}, mgl32.Vec3{0,0,0}, mgl32.Ident4(), mgl32.Ident4()}
	cam.Update();
	
	return cam
}
