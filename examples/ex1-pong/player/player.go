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

// Package player manages a player's state

package player

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/hurricanerix/shade/entity"
	"github.com/hurricanerix/shade/events"
	"github.com/hurricanerix/shade/sprite"
)

const NumToWin = 5
const TopY = 455
const BottomY = 75

// Player state
type Player struct {
	pos    mgl32.Vec3
	Score  int
	Sprite *sprite.Context
	PlayerNum	int  // player 1 or player 2
	paddleSize int
	upKey  bool
	downKey bool
	keyUp	glfw.Key
	keyDown	glfw.Key
}

func New(playerNum int, x, y float32, s *sprite.Context) *Player {
	// create initial paddle
	p := Player{
		pos:    mgl32.Vec3{x, y, 0.0},
		Sprite: s,
		paddleSize: 8,
		PlayerNum: playerNum,
	}

	// assign keys to player
	if p.PlayerNum == 1 { 
		p.keyUp = glfw.KeyQ
		p.keyDown = glfw.KeyA
	} else {
		p.keyUp = glfw.KeyP
		p.keyDown = glfw.KeyL
	}

	return &p
}

func (p Player) Pos() mgl32.Vec3 {
	return p.pos
}

//func (p *Player) HandleEvent(event events.Event, dt float32) {
func (p *Player) Handle(event events.Event) {
	// TODO: move this to SDK to handle things like holding Left & Right at the same time correctly

	if (event.Action == glfw.Press || event.Action == glfw.Repeat) && event.Key == p.keyUp {
		p.upKey = true
	}
	if (event.Action == glfw.Press || event.Action == glfw.Repeat) && event.Key == p.keyDown {
		p.downKey = true
	}

	if event.Action == glfw.Release && event.Key == p.keyUp {
		p.upKey = false
	}
	if event.Action == glfw.Release && event.Key == p.keyDown {
		p.downKey = false
	}	
}

// Update(dt?, group?)
func (p *Player) Update(dt float32, group *[]entity.Entity) {
	posY := p.pos[1]
	if p.upKey && posY <= TopY {
		p.pos[1] += dt
	}
	if p.downKey && posY >= BottomY {
		p.pos[1] -= dt
	}
}

func (p Player) Draw() {
	posX := p.pos[0]
	posY := p.pos[1]

	// DrawFrame(frame to render, position, effect); postion 0,0 is bottom left corner of screen
	// draw top of paddle
	p.Sprite.DrawFrame(mgl32.Vec2{0, 0}, mgl32.Vec3{posX, posY, 0}, nil)
	
	// draw middle part(s) of paddle; paddle shrink by 1 on every win
	for i:=0; i < NumToWin + 1 - p.Score; i++ {
		// position of paddle middle parts are offset by player posY minus (paddleSize * i + 1, i.e. index of loop + 1)
		midPosY := posY - float32(p.paddleSize * (i + 1))
		p.Sprite.DrawFrame(mgl32.Vec2{0, 1}, mgl32.Vec3{posX, midPosY, 0}, nil)
	}

	// draw bottom of paddle
	p.Sprite.DrawFrame(mgl32.Vec2{0, 2}, mgl32.Vec3{posX, posY - float32(p.paddleSize * (NumToWin + 2)) , 0}, nil)
}
