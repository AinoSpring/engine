package io

import (
	"log"

	"github.com/go-gl/glfw/v3.1/glfw"
)

type Window struct {
  Width int
  Height int
  Window *glfw.Window

  InputManager *InputManager
  DeltaTime float64
  LastFrameTime float64
}

func NewWindow(width, height int, title string) *Window {
  window, err := glfw.CreateWindow(width, height, title, nil, nil)
  if err != nil {
    log.Fatal(err)
  }

  window.MakeContextCurrent()
  
  manager := NewInputManager()

  window.SetKeyCallback(manager.keyCallback)
  window.SetCursorPosCallback(manager.mouseCallback)

  return &Window{
    Width: width,
    Height: height,
    Window: window,
    InputManager: manager,
  }
}

func (window *Window) PrepareFrame() {
  window.Window.SwapBuffers()

  glfw.PollEvents()

  frameTime := glfw.GetTime()

  window.DeltaTime = frameTime - window.LastFrameTime
  window.LastFrameTime = frameTime

  window.InputManager.UpdateCursor()
}

