package backend

import (
	"fmt"
	"log"

	"github.com/adlio/trello"
	"github.com/olebedev/config"
)

type Trello struct {
	username  string
	boardName string
	client    *trello.Client
	board     string
	projects  []interface{}
}

func (todo *Trello) Title() string {
	return "Trello"
}

func (todo *Trello) Setup(config *config.Config) {
	todo.username = config.UString("username")
	todo.boardName = config.UString("board")
	todo.client = trello.NewClient(
		config.UString("apiKey"),
		config.UString("accessToken"),
	)
	board, err := getBoardID(todo.client, todo.username, todo.boardName)
	if err != nil {
		log.Fatal(err)
	}
	todo.board = board
	todo.projects = config.UList("lists")
}

func getBoardID(client *trello.Client, username, boardName string) (string, error) {
	member, err := client.GetMember(username, trello.Defaults())
	if err != nil {
		return "", err
	}

	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		return "", err
	}

	for _, board := range boards {
		if board.Name == boardName {
			return board.ID, nil
		}
	}

	return "", fmt.Errorf("could not find board with name %s", boardName)
}

func getListId(client *trello.Client, boardID string, listName string) (string, error) {
	board, err := client.GetBoard(boardID, trello.Defaults())
	if err != nil {
		return "", err
	}

	boardLists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return "", err
	}

	for _, list := range boardLists {
		if list.Name == listName {
			return list.ID, nil
		}
	}

	return "", nil
}

func getCardsOnList(client *trello.Client, listID string) ([]*trello.Card, error) {
	list, err := client.GetList(listID, trello.Defaults())
	if err != nil {
		return nil, err
	}

	cards, err := list.GetCards(trello.Defaults())
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (todo *Trello) BuildProjects() []*Project {
	projects := []*Project{}

	for _, id := range todo.projects {
		proj := todo.GetProject(id.(string))
		projects = append(projects, proj)
	}
	return projects
}

func (todo *Trello) GetProject(id string) *Project {
	proj := &Project{
		Index:   -1,
		backend: todo,
	}

	listId, err := getListId(todo.client, todo.board, id)
	if err != nil {
		proj.Err = err
		return proj
	}
	proj.ID = listId
	proj.Name = id

	tasks, err := todo.LoadTasks(listId)
	proj.Err = err
	proj.Tasks = tasks

	return proj
}

func fromTrello(task *trello.Card) Task {
	return Task{
		ID:        task.ID,
		Completed: task.Closed,
		Name:      task.Name,
	}
}

func (todo *Trello) LoadTasks(id string) ([]Task, error) {
	tasks, err := getCardsOnList(todo.client, id)

	if err != nil {
		return nil, err
	}
	var finalTasks []Task
	for _, item := range tasks {
		finalTasks = append(finalTasks, fromTrello(item))
	}
	return finalTasks, nil
}

func (todo *Trello) CloseTask(task *Task) error {
	args := trello.Arguments{
		"closed": "true",
	}
	if task != nil {
		// Card has an internal client rep which we can't access
		// Just force a lookup
		internal, err := todo.client.GetCard(task.ID, trello.Arguments{})
		if err != nil {
			return err
		}
		return internal.Update(args)
	}
	return nil
}

func (todo *Trello) DeleteTask(_ *Task) error {
	return nil
}

func (todo *Trello) Sources() []string {
	var result []string
	for _, id := range todo.projects {
		result = append(result, id.(string))
	}
	return result
}
