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
// Package rect TODO doc

package rect

type Rect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

// New TODO doc
func New(x, y, width, height float32) (*Rect, error) {
	r := Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
	return &r, nil
}