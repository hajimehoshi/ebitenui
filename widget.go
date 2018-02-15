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

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenui/layout"
)

type WidgetContent interface {
	Update()
	Draw(screen *ebiten.Image, rect image.Rectangle)
	HandleEvent(event Event)
}

type Widget struct {
	node  layout.Node
	input input
}

func NewWidget(content WidgetContent) *Widget {
	w := &Widget{}
	w.node.Object = content
	return w
}

func forEachNode(node *layout.Node, f func(node *layout.Node)) {
	f(node)
	for _, c := range node.ChildNodes() {
		forEachNode(c, f)
	}
}

func (w *Widget) Update() {
	forEachNode(&w.node, func(node *layout.Node) {
		if node.Object != nil {
			node.Object.(WidgetContent).Update()
		}
	})
	w.input.handleInputEvents(&w.node)
}

func (w *Widget) Draw(screen *ebiten.Image) {
	forEachNode(&w.node, func(node *layout.Node) {
		if node.Object != nil {
			node.Object.(WidgetContent).Draw(screen, node.Abs())
		}
	})
}

func (w *Widget) AddChild(widget *Widget, anchor image.Rectangle) {
	w.node.AddChild(&widget.node, anchor)
}

func (w *Widget) Resize(rect image.Rectangle) {
	w.node.Resize(rect)
}
