package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
  position mgl32.Vec3
  normal mgl32.Vec3
  uv mgl32.Vec2
}

type Mesh struct {
  vertices []Vertex
  vaoHandle uint32
}

func NewMesh(glVertices []float32) *Mesh {
  vertices := make([]Vertex, 0)
  for i := 0; i < len(glVertices); i += 8 {
    vertex := Vertex{}
    vertex.position = mgl32.Vec3{
      glVertices[i],
      glVertices[i + 1],
      glVertices[i + 2],
    }
    vertex.normal = mgl32.Vec3{
      glVertices[i + 3],
      glVertices[i + 4],
      glVertices[i + 5],
    }
    vertex.uv = mgl32.Vec2{
      glVertices[i + 6],
      glVertices[i + 7],
    }
    vertices = append(vertices, vertex)
  }

  return &Mesh{vertices: vertices}
}

func (mesh *Mesh) CreateBuffers() {
  vertices := mesh.VertexSlice()

  gl.GenVertexArrays(1, &mesh.vaoHandle)

  var vboHandle uint32
  gl.GenBuffers(1, &vboHandle)
 
  mesh.Bind()

  gl.BindBuffer(gl.ARRAY_BUFFER, vboHandle)
  gl.BufferData(gl.ARRAY_BUFFER, len(vertices) * 4, gl.Ptr(vertices), gl.STATIC_DRAW)

  stride := int32((3 * 4) + (3 * 4) + (2 * 4))  // 4 = size of float in bytes
  offset := 0

  gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
  gl.EnableVertexAttribArray(0);
  offset += 3 * 4

  gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
  gl.EnableVertexAttribArray(1);
  offset += 3 * 4

  gl.VertexAttribPointer(2, 2, gl.FLOAT, false, stride, gl.PtrOffset(offset))
  gl.EnableVertexAttribArray(2);
  offset += 2 * 4

  mesh.Unbind()
}

func (mesh *Mesh) Bind() {
  gl.BindVertexArray(mesh.vaoHandle)
}

func (mesh *Mesh) Unbind() {
  gl.BindVertexArray(0)
}

func (mesh *Mesh) Draw() {
  gl.DrawArrays(gl.TRIANGLES, 0, int32(len(mesh.vertices)))
}

func (mesh *Mesh) VertexSlice() (vertices []float32) {
  vertices = make([]float32, 0)
  for _, vertex := range mesh.vertices {
    vertices = append(vertices, vertex.position[:]...)
    vertices = append(vertices, vertex.normal[:]...)
    vertices = append(vertices, vertex.uv[:]...)
  }
  return
}

