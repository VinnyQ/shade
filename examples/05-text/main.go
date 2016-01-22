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
// Package app manages the main game loop.

package main

import (
	_ "image/png"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/hurricanerix/shade/display"
	"github.com/hurricanerix/shade/events"
	"github.com/hurricanerix/shade/fonts"
)

const windowWidth = 640
const windowHeight = 480

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	screen, err := display.SetMode("05-font", windowWidth, windowHeight)
	if err != nil {
		log.Fatalln("failed to set display mode:", err)
	}

	font, err := fonts.New()
	if err != nil {
		panic(err)
	}
	font.Bind(screen.Program)

	var msg string
	var w, h float32

	for running := true; running; {
		// TODO move this somewhere else (maybe a Clear method of display
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// TODO refector events to be cleaner
		if screen.Window.ShouldClose() {
			running = !screen.Window.ShouldClose()
		}

		for _, event := range events.Get() {
			if event.Action == glfw.Press && event.Key == glfw.KeyEscape {
				running = false
				event.Window.SetShouldClose(true)
			}
		}

		screen.Fill(200.0/256.0, 200/256.0, 200/256.0)

		msg = "Bottom Left"
		font.DrawText(0, 0, 3.0, 3.0, nil, msg)

		msg = "Bottom Right"
		w, _ = font.SizeText(3.0, 3.0, msg)
		font.DrawText(screen.Width-w, 0, 3.0, 3.0, nil, msg)

		msg = "Top Left"
		_, h = font.SizeText(3.0, 3.0, msg)
		font.DrawText(0, screen.Height-h, 3.0, 3.0, nil, msg)

		msg = "Top Right"
		w, h = font.SizeText(3.0, 3.0, msg)
		font.DrawText(screen.Width-w, screen.Height-h, 3.0, 3.0, nil, msg)

		msg = "Center\nMulti-Line\nText\nWith\nColor"
		color := mgl32.Vec4{1.0, 0.0, 0.0, 1.0}
		w, h = font.SizeText(3.0, 3.0, msg)
		font.DrawText(screen.Width/2-w/2, screen.Height/2+h/2, 3.0, 3.0, &color, msg)

		screen.Flip()

		// TODO refector events to be cleaner
		glfw.PollEvents()
	}

}
