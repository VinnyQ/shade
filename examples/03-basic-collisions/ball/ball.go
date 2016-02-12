// Copyright 2016 Richard Hawkins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package ball TODO doc

package ball

import (
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/hurricanerix/shade/entity"
	"github.com/hurricanerix/shade/shapes"
	"github.com/hurricanerix/shade/sprite"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

// Ball TODO doc
type Ball struct {
	Pos    mgl32.Vec3
	Sprite *sprite.Context
	Shape  *shapes.Shape
}

// New TODO doc
func New(x, y float32, s *sprite.Context, group *[]entity.Entity) (*Ball, error) {
	// TODO should take a group in as a argument
	b := Ball{
		Pos:    mgl32.Vec3{x, y, 1.0},
		Sprite: s,
		Shape:  shapes.NewCircle(mgl32.Vec2{float32(s.Width) / 2, float32(s.Height) / 2}, float32(s.Width)/2),
	}

	// TODO: this should probably be added outside of ball
	if group != nil {
		*group = append(*group, &b)
	}
	return &b, nil
}

func (b Ball) Type() string {
	return "ball"
}

func (b Ball) Label() string {
	return ""
}

// Bind TODO doc
func (b *Ball) Bind(program uint32) error {
	return b.Sprite.Bind(program)
}

func (b Ball) Bounds() *shapes.Shape {
	return b.Shape
}

func (b Ball) Pos2() *mgl32.Vec3 {
	return &b.Pos
}

// Update TODO doc
func (b *Ball) Update(dt float32, g []entity.Entity) {
}

// Draw TODO doc
func (b Ball) Draw() {
	b.Sprite.Draw(b.Pos, nil)
}
