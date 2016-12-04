package utils

import (
	"os"
	"image"
	"image/draw"
	_ "image/png"
	"github.com/go-gl/gl/v4.5-core/gl"
)

type Texture uint32

func (tex Texture) Bind(id int) {
	gl.BindTexture(uint32(id), uint32(tex))
}

func (tex Texture) Unbind(id int) {
	gl.BindTexture(uint32(id), 0);
}

func GetTexture(path string) Texture {
	file, err := os.Open(path)
	if err != nil{
		panic(err)
	}
	
	img, _, err := image.Decode(file)
	if err != nil{
		panic(err)
	}
	
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	
	return Texture(texture)
}