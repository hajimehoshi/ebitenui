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

package layout

import (
	"image"
)

func in(r *image.Rectangle, x, y int) bool {
	return r.Min.X <= x && x < r.Max.X && r.Min.Y <= y && y < r.Max.Y
}

type Node struct {
	Object interface{}

	rect     image.Rectangle
	parent   *Node
	children []*Node

	anchor       image.Rectangle
	selfToAnchor image.Rectangle
}

const MaxAnchor = 65536

func (n *Node) AddChild(child *Node, anchor image.Rectangle) {
	s := n.rect.Size()
	child.anchor = anchor
	child.selfToAnchor = image.Rect(
		anchor.Min.X*s.X/MaxAnchor-child.rect.Min.X,
		anchor.Min.Y*s.Y/MaxAnchor-child.rect.Min.Y,
		anchor.Max.X*s.X/MaxAnchor-child.rect.Max.X,
		anchor.Max.Y*s.Y/MaxAnchor-child.rect.Max.Y,
	)
	n.children = append(n.children, child)
	child.parent = n
}

func (n *Node) ChildNodes() []*Node {
	nodes := []*Node{}
	for _, c := range n.children {
		nodes = append(nodes, c)
	}
	return nodes
}

func (n *Node) Parent() *Node {
	return n.parent
}

func subRects(a, b *image.Rectangle) image.Rectangle {
	return image.Rectangle{
		Min: a.Min.Sub(b.Min),
		Max: a.Max.Sub(b.Max),
	}
}

func (n *Node) Resize(rect image.Rectangle) {
	prev := n.rect
	n.rect = rect
	diff := subRects(&n.rect, &prev)
	n.selfToAnchor = subRects(&n.selfToAnchor, &diff)

	for _, c := range n.children {
		c.resize()
	}
}

func (n *Node) resize() {
	s := n.parent.Size()
	a := n.anchor
	n.rect = image.Rect(
		a.Min.X*s.X/MaxAnchor-n.selfToAnchor.Min.X,
		a.Min.Y*s.Y/MaxAnchor-n.selfToAnchor.Min.Y,
		a.Max.X*s.X/MaxAnchor-n.selfToAnchor.Max.X,
		a.Max.Y*s.Y/MaxAnchor-n.selfToAnchor.Max.Y,
	)

	for _, c := range n.children {
		c.resize()
	}
}

func (n *Node) Size() image.Point {
	return n.rect.Size()
}

func (n *Node) Abs() image.Rectangle {
	r := n.rect
	if n.parent != nil {
		return r.Add(n.parent.Abs().Min)
	}
	return r
}

func (n *Node) At(x, y int) *Node {
	abs := n.Abs()

	for _, c := range n.children {
		if n := c.At(x, y); n != nil {
			return n
		}
	}

	if in(&abs, x, y) {
		return n
	}

	return nil
}
