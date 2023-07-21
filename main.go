package main

import (
	"log"
	"math"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"

	"aino-spring.com/engine/engine"
	"aino-spring.com/engine/io"
)

var cubeVertices = []float32{
	// position        // normal vector  // texture uvs
	-0.5, -0.5, -0.5,  0.0,  0.0, -1.0,  0.0, 0.0,
	 0.5, -0.5, -0.5,  0.0,  0.0, -1.0,  1.0, 0.0,
	 0.5,  0.5, -0.5,  0.0,  0.0, -1.0,  1.0, 1.0,
	 0.5,  0.5, -0.5,  0.0,  0.0, -1.0,  1.0, 1.0,
	-0.5,  0.5, -0.5,  0.0,  0.0, -1.0,  0.0, 1.0,
	-0.5, -0.5, -0.5,  0.0,  0.0, -1.0,  0.0, 0.0,

	-0.5, -0.5,  0.5,  0.0,  0.0, 1.0,  0.0, 0.0,
	 0.5, -0.5,  0.5,  0.0,  0.0, 1.0,  1.0, 0.0,
	 0.5,  0.5,  0.5,  0.0,  0.0, 1.0,  1.0, 1.0,
	 0.5,  0.5,  0.5,  0.0,  0.0, 1.0,  1.0, 1.0,
	-0.5,  0.5,  0.5,  0.0,  0.0, 1.0,  0.0, 1.0,
	-0.5, -0.5,  0.5,  0.0,  0.0, 1.0,  0.0, 0.0,

	-0.5,  0.5,  0.5, -1.0,  0.0,  0.0,  1.0, 0.0,
	-0.5,  0.5, -0.5, -1.0,  0.0,  0.0,  1.0, 1.0,
	-0.5, -0.5, -0.5, -1.0,  0.0,  0.0,  0.0, 1.0,
	-0.5, -0.5, -0.5, -1.0,  0.0,  0.0,  0.0, 1.0,
	-0.5, -0.5,  0.5, -1.0,  0.0,  0.0,  0.0, 0.0,
	-0.5,  0.5,  0.5, -1.0,  0.0,  0.0,  1.0, 0.0,

	 0.5,  0.5,  0.5,  1.0,  0.0,  0.0,  1.0, 0.0,
	 0.5,  0.5, -0.5,  1.0,  0.0,  0.0,  1.0, 1.0,
	 0.5, -0.5, -0.5,  1.0,  0.0,  0.0,  0.0, 1.0,
	 0.5, -0.5, -0.5,  1.0,  0.0,  0.0,  0.0, 1.0,
	 0.5, -0.5,  0.5,  1.0,  0.0,  0.0,  0.0, 0.0,
	 0.5,  0.5,  0.5,  1.0,  0.0,  0.0,  1.0, 0.0,

	-0.5, -0.5, -0.5,  0.0, -1.0,  0.0,  0.0, 1.0,
	 0.5, -0.5, -0.5,  0.0, -1.0,  0.0,  1.0, 1.0,
	 0.5, -0.5,  0.5,  0.0, -1.0,  0.0,  1.0, 0.0,
	 0.5, -0.5,  0.5,  0.0, -1.0,  0.0,  1.0, 0.0,
	-0.5, -0.5,  0.5,  0.0, -1.0,  0.0,  0.0, 0.0,
	-0.5, -0.5, -0.5,  0.0, -1.0,  0.0,  0.0, 1.0,

	-0.5,  0.5, -0.5,  0.0,  1.0,  0.0,  0.0, 1.0,
	 0.5,  0.5, -0.5,  0.0,  1.0,  0.0,  1.0, 1.0,
	 0.5,  0.5,  0.5,  0.0,  1.0,  0.0,  1.0, 0.0,
	 0.5,  0.5,  0.5,  0.0,  1.0,  0.0,  1.0, 0.0,
	-0.5,  0.5,  0.5,  0.0,  1.0,  0.0,  0.0, 0.0,
	-0.5,  0.5, -0.5,  0.0,  1.0,  0.0,  0.0, 1.0,
}

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to inifitialize glfw:", err)
	}
	defer glfw.Terminate()

	log.Println(glfw.GetVersionString())

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
  glfw.WindowHint(glfw.Samples, 4)

	window := io.NewWindow(1280, 720, "engine")

  window.Window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	if err := gl.Init(); err != nil {
    log.Fatal(err)
	}

  gl.Enable(gl.MULTISAMPLE)
	gl.Enable(gl.DEPTH_TEST)

  if err := loop(window); err != nil {
		log.Fatal(err)
	}
}

func loop(window *io.Window) error {

	vertShader := engine.NewShaderFromFile("shaders/phong.vert", gl.VERTEX_SHADER)
	fragShader := engine.NewShaderFromFile("shaders/phong.frag", gl.FRAGMENT_SHADER)

	program := engine.NewProgram(vertShader, fragShader)
	defer program.Delete()

	lightFragShader := engine.NewShaderFromFile("shaders/light.frag", gl.FRAGMENT_SHADER)

	lightProgram := engine.NewProgram(vertShader, lightFragShader)
  defer lightProgram.Delete()

  cube := engine.NewMesh(cubeVertices)
  cube.CreateBuffers()
  light := engine.NewMesh(cubeVertices)
  light.CreateBuffers()

  texture := engine.NewTextureFromFile("images/RTS_Crate.png", gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)

  var pitch float64
  var yaw float64
  var position mgl32.Vec3
  var front mgl32.Vec3

	for !window.Window.ShouldClose() {
		window.PrepareFrame()

    speed := float32(window.DeltaTime * 2)
    right := mgl32.Vec3{0, 1, 0}.Cross(front).Normalize()
    up := right.Cross(front).Normalize()

    {
      if window.InputManager.KeysPressed[glfw.KeyW] {
        position = position.Add(front.Mul(speed))
      }
      if window.InputManager.KeysPressed[glfw.KeyS] {
        position = position.Sub(front.Mul(speed))
      }
      if window.InputManager.KeysPressed[glfw.KeyA] {
        position = position.Sub(front.Cross(up).Normalize().Mul(speed))
      }
      if window.InputManager.KeysPressed[glfw.KeyD] {
        position = position.Add(front.Cross(up).Normalize().Mul(speed))
      }

      dCursor := window.InputManager.CursorDelta

      dx := -0.5 * dCursor[0]
      dy := 0.5 * dCursor[1]

      pitch += dy
      if pitch > 89.0 {
        pitch = 89.0
      } else if pitch < -89.0 {
        pitch = -89.0
      }

      yaw = math.Mod(yaw + dx, 360)

      front[0] = float32(math.Cos(mgl64.DegToRad(pitch)) * math.Cos(mgl64.DegToRad(yaw)))
      front[1] = float32(math.Sin(mgl64.DegToRad(pitch)))
      front[2] = float32(math.Cos(mgl64.DegToRad(pitch)) * math.Sin(mgl64.DegToRad(yaw)))
    }

		gl.ClearColor(0, 0, 0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)  // depth buffer needed for DEPTH_TEST

		fov := float32(60.0)
		projectTransform := mgl32.Perspective(mgl32.DegToRad(fov), float32(window.Width) / float32(window.Height), 0.1, 100.0)

    camTransform := mgl32.LookAtV(position, position.Add(front), up)

		lightPos := mgl32.Vec3{0.6, 1, 0.1}
		lightTransform := mgl32.Translate3D(lightPos.X(), lightPos.Y(), lightPos.Z()).Mul4(mgl32.Scale3D(0.2, 0.2, 0.2))

    {
      program.Use()
      texture.Bind(gl.TEXTURE0)
      texture.SetUniform(program.GetUniformLocation("bg_texture"))

      gl.UniformMatrix4fv(program.GetUniformLocation("view"), 1, false, &camTransform[0])
      gl.UniformMatrix4fv(program.GetUniformLocation("project"), 1, false, &projectTransform[0])

      cube.Bind()

      gl.Uniform3f(program.GetUniformLocation("lightColor"), 1.0, 1.0, 1.0)
      gl.Uniform3f(program.GetUniformLocation("lightPos"), lightPos.X(), lightPos.Y(), lightPos.Z())

      worldTransform := mgl32.Ident4()

      gl.UniformMatrix4fv(program.GetUniformLocation("model"), 1, false, &worldTransform[0])

      cube.Draw()

      cube.Unbind()
    }
    {
      lightProgram.Use()
      light.Bind()

      gl.UniformMatrix4fv(lightProgram.GetUniformLocation("model"), 1, false, &lightTransform[0])
      gl.UniformMatrix4fv(lightProgram.GetUniformLocation("view"), 1, false, &camTransform[0])
      gl.UniformMatrix4fv(lightProgram.GetUniformLocation("project"), 1, false, &projectTransform[0])

      light.Draw()
      light.Unbind()
    }
	}

	return nil
}
