package graphics

import (
	"fmt"
	"github.com/go-gl/gl/v4.5-core/gl"
	"io/ioutil"
	"strings"
)

type ShaderProgram uint32

func (s ShaderProgram) Begin() {
	gl.UseProgram(uint32(s))
}

func (s ShaderProgram) End() {
	gl.UseProgram(uint32(0))
}

func NewShaderProgram(vertexShaderSource, fragmentShaderSource string) ShaderProgram {
	vertexShader := compileShader(ReadShaderFile(vertexShaderSource) + "\x00", gl.VERTEX_SHADER)
	fragmentShader := compileShader(ReadShaderFile(fragmentShaderSource)+ "\x00", gl.FRAGMENT_SHADER)
	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		panic(fmt.Sprintf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return ShaderProgram(program)
}

func compileShader(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()

	gl.CompileShader(shader)
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		panic(fmt.Sprintf("failed to compile %v: %v", source, log))
	}

	return shader
}

func ReadShaderFile(name string) string {
	dat, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(dat)

}
