package io

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl64"
)

type InputManager struct {
  KeysPressed map[glfw.Key]bool
  Cursor mgl64.Vec2
  LastCursor mgl64.Vec2
  CursorDelta mgl64.Vec2
}

func NewInputManager() *InputManager {
  return &InputManager{
    KeysPressed: make(map[glfw.Key]bool),
  }
}

func (manager *InputManager) UpdateCursor() {
  if manager.LastCursor == manager.Cursor {
    manager.CursorDelta = mgl64.Vec2{0, 0}
  }
  manager.LastCursor = manager.Cursor
}

func (manager *InputManager) keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
  switch action {
  case glfw.Press:
    manager.KeysPressed[key] = true
  case glfw.Release:
    manager.KeysPressed[key] = false
  }
}

func (manager *InputManager) mouseCallback(window *glfw.Window, xpos, ypos float64) {
  manager.CursorDelta = mgl64.Vec2{xpos, ypos}.Sub(manager.Cursor)
  manager.Cursor = mgl64.Vec2{xpos, ypos}
}


