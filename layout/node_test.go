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

package layout_test

import (
	"image"
	"testing"

	. "github.com/hajimehoshi/ebitenui/layout"
)

func TestResize(t *testing.T) {
	cases := []struct {
		Name     string
		Parent   image.Rectangle
		Child    image.Rectangle
		Anchor   image.Rectangle
		Resize   image.Rectangle
		ChildAbs image.Rectangle
	}{
		{
			Name:     "no resize",
			Parent:   image.Rect(10, 10, 20, 20),
			Child:    image.Rect(1, 1, 9, 9),
			Anchor:   image.Rect(0, 0, MaxAnchor, MaxAnchor),
			Resize:   image.Rect(10, 10, 20, 20),
			ChildAbs: image.Rect(11, 11, 19, 19),
		},
		{
			Name:     "enlarge",
			Parent:   image.Rect(10, 10, 20, 20),
			Child:    image.Rect(1, 1, 9, 9),
			Anchor:   image.Rect(0, 0, MaxAnchor, MaxAnchor),
			Resize:   image.Rect(10, 10, 30, 30),
			ChildAbs: image.Rect(11, 11, 29, 29),
		},
		{
			Name:     "enlarge (anchor: left upper)",
			Parent:   image.Rect(10, 10, 20, 20),
			Child:    image.Rect(1, 1, 9, 9),
			Anchor:   image.Rect(0, 0, 0, 0),
			Resize:   image.Rect(10, 10, 30, 30),
			ChildAbs: image.Rect(11, 11, 19, 19),
		},
		{
			Name:     "enlarge (anchor: upper)",
			Parent:   image.Rect(10, 10, 20, 20),
			Child:    image.Rect(1, 1, 9, 9),
			Anchor:   image.Rect(0, 0, MaxAnchor, 0),
			Resize:   image.Rect(10, 10, 30, 30),
			ChildAbs: image.Rect(11, 11, 29, 19),
		},
		{
			Name:     "enlarge (anchor: lower)",
			Parent:   image.Rect(10, 10, 20, 20),
			Child:    image.Rect(1, 1, 9, 9),
			Anchor:   image.Rect(0, MaxAnchor, MaxAnchor, MaxAnchor),
			Resize:   image.Rect(10, 10, 30, 30),
			ChildAbs: image.Rect(11, 21, 29, 29),
		},
		{
			Name:     "enlarge (anchor: center)",
			Parent:   image.Rect(10, 10, 20, 20),
			Child:    image.Rect(1, 1, 9, 9),
			Anchor:   image.Rect(MaxAnchor/2, MaxAnchor/2, MaxAnchor/2, MaxAnchor/2),
			Resize:   image.Rect(10, 10, 30, 30),
			ChildAbs: image.Rect(16, 16, 24, 24),
		},
		{
			Name:     "move",
			Parent:   image.Rect(10, 10, 20, 20),
			Child:    image.Rect(1, 1, 9, 9),
			Anchor:   image.Rect(0, 0, MaxAnchor, MaxAnchor),
			Resize:   image.Rect(20, 20, 30, 30),
			ChildAbs: image.Rect(21, 21, 29, 29),
		},
	}
	for _, c := range cases {
		n1 := &Node{}
		n1.Resize(image.Rect(10, 10, 20, 20))
		n2 := &Node{}
		n2.Resize(image.Rect(1, 1, 9, 9))
		n1.AddChild(n2, c.Anchor)
		n1.Resize(c.Resize)

		got := n2.Abs()
		want := c.ChildAbs
		if got != want {
			t.Errorf("case (%s): n2.Abs() = %v, want: %v", c.Name, got, want)
		}
	}
}

func TestResizeChild(t *testing.T) {
	parent := &Node{}
	parent.Resize(image.Rect(10, 10, 20, 20))
	child := &Node{}
	child.Resize(image.Rect(1, 1, 9, 9))
	parent.AddChild(child, image.Rect(0, 0, MaxAnchor, MaxAnchor))

	// Resize the child first.
	child.Resize(image.Rect(1, 1, 19, 19))
	got := child.Abs()
	want := image.Rect(11, 11, 29, 29)
	if got != want {
		t.Errorf("child resized: got %v, want %v", got, want)
	}

	// Resize the parent with the same size.
	parent.Resize(image.Rect(10, 10, 20, 20))
	got = child.Abs()
	want = image.Rect(11, 11, 29, 29)
	if got != want {
		t.Errorf("parent resized (no effect): got %v, want %v", got, want)
	}

	parent.Resize(image.Rect(10, 10, 30, 30))
	got = child.Abs()
	want = image.Rect(11, 11, 39, 39)
	if got != want {
		t.Errorf("parent resized: got %v, want %v", got, want)
	}
}
