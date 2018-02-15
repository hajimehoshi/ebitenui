// Copyright 2018 Hajime Hoshi
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

package ebitenui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Button struct {
	Text        string
	OnClick     func()
	OnMouseDown func()
	OnMouseUp   func()
}

func (b *Button) Update() {
}

func (b *Button) Draw(screen *ebiten.Image, rect image.Rectangle) {
	ebitenutil.DrawRect(screen, float64(rect.Min.X), float64(rect.Min.Y), float64(rect.Dx()), float64(rect.Dy()), color.RGBA{0xff, 0x80, 0x80, 0xff})
}

func (b *Button) HandleEvent(event Event) {
	switch event.Name() {
	case "click":
		if b.OnClick != nil {
			b.OnClick()
		}
	case "mousedown":
		if b.OnMouseDown != nil {
			b.OnMouseDown()
		}
	case "mouseup":
		if b.OnMouseUp != nil {
			b.OnMouseUp()
		}
	}
	event.StopPropagation()
}
