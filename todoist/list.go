package todoist

import (
	"fmt"

	"github.com/darkSasori/todoist"
)

type List struct {
	todoist.Project
	items []todoist.Task
	index int
}

func NewList(id int) *List {
	project, err := todoist.GetProject(id)
	if err != nil {
		panic(err)
	}

	list := &List{
		Project: project,
		index:   -1,
	}
	list.loadItems()
	return list
}

func (l List) isFirst() bool {
	return l.index == 0
}

func (l List) isLast() bool {
	return l.index >= len(l.items)-1
}

func (l *List) up() {
	l.index = l.index - 1
	if l.index < 0 {
		l.index = len(l.items) - 1
	}
}

func (l *List) down() {
	if l.index == -1 {
		l.index = 0
		return
	}

	l.index = l.index + 1
	if l.index >= len(l.items) {
		l.index = 0
	}
}

func (l *List) loadItems() {
	tasks, err := todoist.ListTask(todoist.QueryParam{"project_id": fmt.Sprintf("%d", l.ID)})
	if err != nil {
		panic(err)
	}

	l.items = tasks
}

func (list *List) LongestLine() int {
	maxLen := 0

	for _, item := range list.items {
		if len(item.Content) > maxLen {
			maxLen = len(item.Content)
		}
	}

	return maxLen
}

func (l *List) close() {
	if err := l.items[l.index].Close(); err != nil {
		panic(err)
	}

	l.loadItems()
}

func (l *List) delete() {
	if err := l.items[l.index].Delete(); err != nil {
		panic(err)
	}

	l.loadItems()
}
