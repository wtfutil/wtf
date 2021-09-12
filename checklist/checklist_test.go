package checklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCheckist(t *testing.T) {
	cl := NewChecklist("o", "-")

	assert.IsType(t, Checklist{}, cl)
	assert.Equal(t, "o", cl.checkedIcon)
	assert.Equal(t, -1, cl.selected)
	assert.Equal(t, "-", cl.uncheckedIcon)
	assert.Equal(t, 0, len(cl.Items))
}

func Test_Add(t *testing.T) {
	cl := NewChecklist("o", "-")
	cl.Add(true, nil, make([]string, 0), "test item")

	assert.Equal(t, 1, len(cl.Items))
}

func Test_CheckedItems(t *testing.T) {
	tests := []struct {
		name        string
		expectedLen int
		checkedLen  int
		before      func(cl *Checklist)
	}{
		{
			name:        "with no items",
			expectedLen: 0,
			checkedLen:  0,
			before:      func(cl *Checklist) {},
		},
		{
			name:        "with no checked items",
			expectedLen: 1,
			checkedLen:  0,
			before: func(cl *Checklist) {
				cl.Add(false, nil, make([]string, 0), "unchecked item")
			},
		},
		{
			name:        "with one checked item",
			expectedLen: 2,
			checkedLen:  1,
			before: func(cl *Checklist) {
				cl.Add(false, nil, make([]string, 0), "unchecked item")
				cl.Add(true, nil, make([]string, 0), "checked item")
			},
		},
		{
			name:        "with multiple checked items",
			expectedLen: 3,
			checkedLen:  2,
			before: func(cl *Checklist) {
				cl.Add(false, nil, make([]string, 0), "unchecked item")
				cl.Add(true, nil, make([]string, 0), "checked item 11")
				cl.Add(true, nil, make([]string, 0), "checked item 2")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewChecklist("o", "-")
			tt.before(&cl)

			assert.Equal(t, tt.expectedLen, len(cl.Items))
			assert.Equal(t, tt.checkedLen, len(cl.CheckedItems()))
		})
	}
}

func Test_Delete(t *testing.T) {
	tests := []struct {
		name        string
		idx         int
		expectedLen int
	}{
		{
			name:        "with valid index",
			idx:         0,
			expectedLen: 0,
		},
		{
			name:        "with invalid index",
			idx:         2,
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewChecklist("o", "-")

			cl.Add(true, nil, make([]string, 0), "test item")
			cl.Delete(tt.idx)

			assert.Equal(t, tt.expectedLen, len(cl.Items))
		})
	}
}

func Test_IsSelectable(t *testing.T) {
	tests := []struct {
		name     string
		selected int
		expected bool
	}{
		{
			name:     "nothing selected",
			selected: -1,
			expected: false,
		},
		{
			name:     "valid selection",
			selected: 1,
			expected: true,
		},
		{
			name:     "invalid selection",
			selected: 3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewChecklist("o", "-")
			cl.Add(true, nil, make([]string, 0), "test item 1")
			cl.Add(false, nil, make([]string, 0), "test item 2")

			cl.selected = tt.selected

			assert.Equal(t, tt.expected, cl.IsSelectable())
		})
	}
}

func Test_IsUnselectable(t *testing.T) {
	tests := []struct {
		name     string
		selected int
		expected bool
	}{
		{
			name:     "nothing selected",
			selected: -1,
			expected: true,
		},
		{
			name:     "valid selection",
			selected: 1,
			expected: false,
		},
		{
			name:     "invalid selection",
			selected: 3,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewChecklist("o", "-")
			cl.Add(true, nil, make([]string, 0), "test item 1")
			cl.Add(false, nil, make([]string, 0), "test item 2")

			cl.selected = tt.selected

			assert.Equal(t, tt.expected, cl.IsUnselectable())
		})
	}
}

func Test_LongestLine(t *testing.T) {
	tests := []struct {
		name        string
		expectedLen int
		before      func(cl *Checklist)
	}{
		{
			name:        "with no items",
			expectedLen: 0,
			before:      func(cl *Checklist) {},
		},
		{
			name:        "with different-length items",
			expectedLen: 12,
			before: func(cl *Checklist) {
				cl.Add(true, nil, make([]string, 0), "test item 1")
				cl.Add(false, nil, make([]string, 0), "test item 22")
			},
		},
		{
			name:        "with same-length items",
			expectedLen: 11,
			before: func(cl *Checklist) {
				cl.Add(true, nil, make([]string, 0), "test item 1")
				cl.Add(false, nil, make([]string, 0), "test item 2")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewChecklist("o", "-")
			tt.before(&cl)

			assert.Equal(t, tt.expectedLen, cl.LongestLine())
		})
	}
}

func Test_IndexByItem(t *testing.T) {
	cl := NewChecklist("o", "-")
	cl.Add(false, nil, make([]string, 0), "unchecked item")
	cl.Add(true, nil, make([]string, 0), "checked item")

	tests := []struct {
		name        string
		item        *ChecklistItem
		expectedIdx int
		expectedOk  bool
	}{
		{
			name:        "with nil",
			item:        nil,
			expectedIdx: 0,
			expectedOk:  false,
		},
		{
			name:        "with valid item",
			item:        cl.Items[1],
			expectedIdx: 1,
			expectedOk:  true,
		},
		{
			name:        "with valid item",
			item:        NewChecklistItem(false, nil, make([]string, 0), "invalid", "x", " "),
			expectedIdx: 0,
			expectedOk:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			idx, ok := cl.IndexByItem(tt.item)

			assert.Equal(t, tt.expectedIdx, idx)
			assert.Equal(t, tt.expectedOk, ok)
		})
	}
}

func Test_UncheckedItems(t *testing.T) {
	tests := []struct {
		name        string
		expectedLen int
		checkedLen  int
		before      func(cl *Checklist)
	}{
		{
			name:        "with no items",
			expectedLen: 0,
			checkedLen:  0,
			before:      func(cl *Checklist) {},
		},
		{
			name:        "with no unchecked items",
			expectedLen: 1,
			checkedLen:  0,
			before: func(cl *Checklist) {
				cl.Add(true, nil, make([]string, 0), "unchecked item")
			},
		},
		{
			name:        "with one unchecked item",
			expectedLen: 2,
			checkedLen:  1,
			before: func(cl *Checklist) {
				cl.Add(false, nil, make([]string, 0), "unchecked item")
				cl.Add(true, nil, make([]string, 0), "checked item")
			},
		},
		{
			name:        "with multiple unchecked items",
			expectedLen: 3,
			checkedLen:  2,
			before: func(cl *Checklist) {
				cl.Add(false, nil, make([]string, 0), "unchecked item")
				cl.Add(true, nil, make([]string, 0), "checked item 11")
				cl.Add(false, nil, make([]string, 0), "checked item 2")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewChecklist("o", "-")
			tt.before(&cl)

			assert.Equal(t, tt.expectedLen, len(cl.Items))
			assert.Equal(t, tt.checkedLen, len(cl.UncheckedItems()))
		})
	}
}

func Test_Unselect(t *testing.T) {
	cl := NewChecklist("o", "-")
	cl.Add(false, nil, make([]string, 0), "unchecked item")

	cl.selected = 0
	assert.Equal(t, 0, cl.selected)

	cl.Unselect()
	assert.Equal(t, -1, cl.selected)
}

/* -------------------- Sort Interface -------------------- */

func Test_Len(t *testing.T) {
	tests := []struct {
		name        string
		expectedLen int
		before      func(cl *Checklist)
	}{
		{
			name:        "with no items",
			expectedLen: 0,
			before:      func(cl *Checklist) {},
		},
		{
			name:        "with one item",
			expectedLen: 1,
			before: func(cl *Checklist) {
				cl.Add(false, nil, make([]string, 0), "unchecked item")
			},
		},
		{
			name:        "with multiple items",
			expectedLen: 3,
			before: func(cl *Checklist) {
				cl.Add(false, nil, make([]string, 0), "unchecked item")
				cl.Add(true, nil, make([]string, 0), "checked item 1")
				cl.Add(false, nil, make([]string, 0), "checked item 2")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewChecklist("o", "-")
			tt.before(&cl)

			assert.Equal(t, tt.expectedLen, cl.Len())
		})
	}
}

func Test_Less(t *testing.T) {
	tests := []struct {
		name     string
		first    string
		second   string
		expected bool
	}{
		{
			name:     "same",
			first:    "",
			second:   "",
			expected: false,
		},
		{
			name:     "last less",
			first:    "beta",
			second:   "alpha",
			expected: true,
		},
		{
			name:     "first less",
			first:    "alpha",
			second:   "beta",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewChecklist("o", "-")
			cl.Add(false, nil, make([]string, 0), tt.first)
			cl.Add(false, nil, make([]string, 0), tt.second)

			assert.Equal(t, tt.expected, cl.Less(0, 1))
		})
	}
}

func Test_Swap(t *testing.T) {
	tests := []struct {
		name     string
		first    string
		second   string
		expected bool
	}{
		{
			name:   "same",
			first:  "",
			second: "",
		},
		{
			name:   "last less",
			first:  "alpha",
			second: "beta",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewChecklist("o", "-")
			cl.Add(false, nil, make([]string, 0), tt.first)
			cl.Add(false, nil, make([]string, 0), tt.second)

			cl.Swap(0, 1)

			assert.Equal(t, tt.expected, cl.Items[0].Text == "beta")
			assert.Equal(t, tt.expected, cl.Items[1].Text == "alpha")
		})
	}
}
