package engine

import (
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
  handle uint32
}

type Program struct {
  handle uint32
  shaders []*Shader
}

func NewProgram(shaders ...*Shader) *Program {
  program := &Program{handle: gl.CreateProgram()}
  program.Attach(shaders...)
  program.Link()

  return program
}

func NewShaderFromFile(path string, shaderType uint32) *Shader {
  data, err := os.ReadFile(path)
  if err != nil {
    log.Fatal(err)
  }
  return NewShader(string(data), shaderType)
}

func NewShader(source string, shaderType uint32) *Shader {
  handle := gl.CreateShader(shaderType)
  glSource, free := gl.Strs(source + "\x00")
  defer free()
  gl.ShaderSource(handle, 1, glSource, nil)
  gl.CompileShader(handle)
  glRaiseError(handle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog, "Failed to compile shader")

  return &Shader{
    handle: handle,
  }
}

func (shader *Shader) Delete() {
  gl.DeleteShader(shader.handle)
}

func (program *Program) Delete() {
  for _, shader := range program.shaders {
    shader.Delete()
  }
  gl.DeleteProgram(program.handle)
}

func (program *Program) Attach(shaders ...*Shader) {
  for _, shader := range shaders {
    gl.AttachShader(program.handle, shader.handle)
    program.shaders = append(program.shaders, shader)
  }
}

func (program *Program) Use() {
  gl.UseProgram(program.handle)
}

func (program *Program) Link() {
  gl.LinkProgram(program.handle)
  glRaiseError(program.handle, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog, "Failed to link program")
}

func (program *Program) GetUniformLocation(name string) int32 {
  return gl.GetUniformLocation(program.handle, gl.Str(name + "\x00"))
}

type getObjIv func(uint32, uint32, *int32)
type getObjInfoLog func(uint32, int32, *int32, *uint8)

func glRaiseError(glHandle uint32, checkTrueParam uint32, getObjIvFn getObjIv, getObjInfoLogFn getObjInfoLog, failMsg string) error {

	var success int32
	getObjIvFn(glHandle, checkTrueParam, &success)

	if success == gl.FALSE {
		var logLength int32
		getObjIvFn(glHandle, gl.INFO_LOG_LENGTH, &logLength)

		logValue := gl.Str(strings.Repeat("\x00", int(logLength)))
		getObjInfoLogFn(glHandle, logLength, nil, logValue)

		log.Fatalf("%s: %s", failMsg, gl.GoStr(logValue))
	}

	return nil
}
