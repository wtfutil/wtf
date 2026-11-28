package view

import (
	"fmt"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

func testScrollableWidget() ScrollableWidget {
	return NewScrollableWidget(
		tview.NewApplication(),
		make(chan bool, 1),
		tview.NewPages(),
		&cfg.Common{
			Module: cfg.Module{
				Name: "test scrollable widget",
			},
		},
	)
}

func Test_ScrollableWidget_SelectionNavigation(t *testing.T) {
	widget := testScrollableWidget()
	widget.SetItemCount(3)

	rendered := 0
	widget.SetRenderFunction(func() {
		rendered++
	})

	assert.Equal(t, -1, widget.GetSelected())

	// Next cycles 0 -> 1 -> 2 -> 0
	widget.Next()
	assert.Equal(t, 0, widget.GetSelected())
	assert.Equal(t, 1, rendered)

	widget.Next()
	assert.Equal(t, 1, widget.GetSelected())

	widget.Next()
	assert.Equal(t, 2, widget.GetSelected())

	widget.Next()
	assert.Equal(t, 0, widget.GetSelected())

	// Prev cycles backwards
	widget.Prev()
	assert.Equal(t, 2, widget.GetSelected())

	widget.Prev()
	assert.Equal(t, 1, widget.GetSelected())

	// Unselect resets to -1
	widget.Unselect()
	assert.Equal(t, -1, widget.GetSelected())
}

func Test_ScrollableWidget_EmptyItems(t *testing.T) {
	widget := testScrollableWidget()
	widget.SetItemCount(0)

	assert.Equal(t, -1, widget.GetSelected())

	rendered := false
	widget.SetRenderFunction(func() {
		rendered = true
	})

	widget.Next()
	assert.Equal(t, -1, widget.GetSelected())
	assert.True(t, rendered)

	widget.Prev()
	assert.Equal(t, -1, widget.GetSelected())
}

func Test_ScrollableWidget_RedrawHighlightAndScroll(t *testing.T) {
	widget := testScrollableWidget()
	widget.SetItemCount(5)

	rendered := false
	renderData := func() (string, string, bool) {
		content := ""
		for i := 0; i < 5; i++ {
			content += utils.HighlightableHelper(widget.View, fmt.Sprintf("Row %d", i), i, 5)
		}
		return "Title", content, false
	}
	widget.SetRenderFunction(func() {
		rendered = true
		widget.Redraw(renderData)
	})

	// When unselected (-1), highlights should be empty
	widget.Redraw(renderData)
	assert.Empty(t, widget.View.GetHighlights())

	// Select row 2
	widget.Selected = 2
	widget.Redraw(renderData)

	highlights := widget.View.GetHighlights()
	assert.Equal(t, []string{"2"}, highlights)

	// Unselect should clear highlights
	widget.Unselect()
	assert.True(t, rendered)
	assert.Empty(t, widget.View.GetHighlights())
}
