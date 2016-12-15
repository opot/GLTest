package graphics

import (
	"time"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Program struct {
	Window *glfw.Window
	draw func()
	update func(delta float32)
	last time.Time
}

func NewGlfwProgram(title string, width, height int, draw func(), update func(delta float32)) Program {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	Window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	Window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	return Program{Window, draw, update, time.Now()}
}

func (program *Program) Update(){
	program.update(float32(time.Now().Sub(program.last).Nanoseconds())/100000000.0)
	program.draw()
	program.Window.SwapBuffers()
	glfw.PollEvents()
	program.last = time.Now()
}