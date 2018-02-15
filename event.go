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
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebitenui/layout"
)

type Event interface {
	Name() string
	StopPropagation()
}

type event struct {
	name               string
	propagationStopped bool
}

func (e *event) Name() string {
	return e.name
}

func (e *event) StopPropagation() {
	e.propagationStopped = true
}

func propagateEvent(node *layout.Node, eventName string) {
	for ; node != nil; node = node.Parent() {
		if node.Object == nil {
			continue
		}
		e := &event{
			name: eventName,
		}
		node.Object.(WidgetContent).HandleEvent(e)
		if e.propagationStopped {
			break
		}
	}
}

type input struct {
	mouseNode *layout.Node
}

func (i *input) handleInputEvents(node *layout.Node) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		n := node.At(ebiten.CursorPosition())
		propagateEvent(n, "mousedown")
		i.mouseNode = n
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		n := node.At(ebiten.CursorPosition())
		propagateEvent(n, "mouseup")
		if i.mouseNode == n {
			propagateEvent(n, "click")
		}
		i.mouseNode = nil
	}
}
