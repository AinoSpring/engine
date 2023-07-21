package engine

import (
	"image"
	"image/draw"
  _ "image/png"
  _ "image/jpeg"
	"log"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Texture struct {
  handle uint32
  target uint32
  unit uint32
}

func NewTextureFromFile(path string, wrapR, wrapS int32) *Texture {
  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  img, _, err := image.Decode(file)
  if err != nil {
    log.Fatal(err)
  }
  return NewTexture(img, wrapR, wrapS)
}

func NewTexture(img image.Image, wrapR, wrapS int32) *Texture {
  rgba := image.NewRGBA(img.Bounds())
  draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
  
  var handle uint32
  gl.GenTextures(1, &handle)
  
  texture := Texture{
    handle: handle,
    target: gl.TEXTURE_2D,
  }

  texture.Bind(gl.TEXTURE0)
  defer texture.Unbind()

  gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_R, wrapR)
  gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_S, wrapS)
  gl.TexParameteri(texture.target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
  gl.TexParameteri(texture.target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

  gl.TexImage2D(gl.TEXTURE_2D, 0, gl.SRGB_ALPHA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

  gl.GenerateMipmap(texture.handle)

  return &texture
}

func (texture *Texture) Bind(unit uint32) {
  gl.ActiveTexture(unit)
  gl.BindTexture(texture.target, texture.handle)
  texture.unit = unit
}

func (texture *Texture) Unbind() {
  texture.unit = 0
  gl.BindTexture(texture.target, 0)
}

func (texture *Texture) SetUniform(location int32) {
  if texture.unit == 0 {
    log.Fatal("Texture has not been bound")
  }
  gl.Uniform1i(location, int32(texture.unit - gl.TEXTURE0))
}

