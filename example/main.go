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

package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenui"
	"github.com/hajimehoshi/ebitenui/layout"
)

type App struct {
	root *ebitenui.Widget
}

func (a *App) Update(screen *ebiten.Image) error {

	a.root.Update()

	if ebiten.IsRunningSlowly() {
		return nil
	}

	a.root.Draw(screen)

	return nil
}

func main() {
	app := &App{
		root: ebitenui.NewWidget(nil),
	}
	app.root.Resize(image.Rect(0, 0, 640, 480))

	panel := ebitenui.NewWidget(&ebitenui.Panel{})
	panel.Resize(image.Rect(30, 30, 300, 200))

	button := ebitenui.NewWidget(&ebitenui.Button{
		OnClick: func() {
			fmt.Println("Button Click")
		},
		OnMouseDown: func() {
			fmt.Println("Button MouseDown")
		},
		OnMouseUp: func() {
			fmt.Println("Button MouseUp")
		},
	})
	button.Resize(image.Rect(10, 10, 100, 100))
	panel.AddChild(button, image.Rect(0, 0, layout.MaxAnchor, layout.MaxAnchor))

	app.root.AddChild(panel, image.Rect(0, 0, layout.MaxAnchor, layout.MaxAnchor))

	if err := ebiten.Run(app.Update, 640, 480, 1, "UI Example"); err != nil {
		panic(err)
	}
}
