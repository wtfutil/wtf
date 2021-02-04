package asana

import (
	"fmt"
	"strings"
	"time"

	asana "bitbucket.org/mikehouston/asana-go"
)

func fetchTasksFromProject(token, projectId, mode string) ([]*TaskItem, error) {
	taskItems := []*TaskItem{}
	uidToName := make(map[string]string)

	client := asana.NewClientWithAccessToken(token)

	uid, err := getCurrentUserId(client, mode)
	if err != nil {
		return nil, err
	}

	q := &asana.TaskQuery{
		Project: projectId,
	}

	fetchedTasks, _, err := getTasksFromAsana(client, q)
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks: %s", err)
	}

	processFetchedTasks(client, &fetchedTasks, &taskItems, &uidToName, mode, projectId, uid)

	return taskItems, nil
}

func fetchTasksFromProjectSections(token, projectId string, sections []string, mode string) ([]*TaskItem, error) {
	taskItems := []*TaskItem{}
	uidToName := make(map[string]string)

	client := asana.NewClientWithAccessToken(token)

	uid, err := getCurrentUserId(client, mode)
	if err != nil {
		return nil, err
	}

	p := &asana.Project{
		ID: projectId,
	}

	for _, section := range sections {

		sectionId, err := findSection(client, p, section)
		if err != nil {
			return nil, fmt.Errorf("error fetching tasks: %s", err)
		}

		q := &asana.TaskQuery{
			Section: sectionId,
		}

		fetchedTasks, _, err := getTasksFromAsana(client, q)
		if err != nil {
			return nil, fmt.Errorf("error fetching tasks: %s", err)
		}

		if len(fetchedTasks) > 0 {
			taskItem := &TaskItem{
				name:     section,
				taskType: TASK_SECTION,
			}

			taskItems = append(taskItems, taskItem)
		}

		processFetchedTasks(client, &fetchedTasks, &taskItems, &uidToName, mode, projectId, uid)

	}

	return taskItems, nil
}

func fetchTasksFromWorkspace(token, workspaceId, mode string) ([]*TaskItem, error) {
	taskItems := []*TaskItem{}
	uidToName := make(map[string]string)

	client := asana.NewClientWithAccessToken(token)

	uid, err := getCurrentUserId(client, mode)
	if err != nil {
		return nil, err
	}

	q := &asana.TaskQuery{
		Workspace: workspaceId,
		Assignee:  "me",
	}

	fetchedTasks, _, err := getTasksFromAsana(client, q)
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks: %s", err)
	}

	processFetchedTasks(client, &fetchedTasks, &taskItems, &uidToName, mode, workspaceId, uid)

	return taskItems, nil

}

func toggleTaskCompletionById(token, taskId string) error {
	client := asana.NewClientWithAccessToken(token)

	t := &asana.Task{
		ID: taskId,
	}

	err := t.Fetch(client)
	if err != nil {
		return fmt.Errorf("error fetching task: %s", err)
	}

	updateReq := &asana.UpdateTaskRequest{}

	if *t.Completed {
		f := false
		updateReq.Completed = &f
	} else {
		t := true
		updateReq.Completed = &t
	}

	err = t.Update(client, updateReq)
	if err != nil {
		return fmt.Errorf("error updating task: %s", err)
	}

	return nil
}

func processFetchedTasks(client *asana.Client, fetchedTasks *[]*asana.Task, taskItems *[]*TaskItem, uidToName *map[string]string, mode, projectId, uid string) {

	for _, task := range *fetchedTasks {
		switch {
		case strings.HasSuffix(mode, "_all"):
			if task.Assignee != nil {
				// Check if we have already looked up this user
				if assigneeName, ok := (*uidToName)[task.Assignee.ID]; ok {
					task.Assignee.Name = assigneeName
				} else {
					// We haven't looked up this user before, perform the lookup now
					assigneeName, err := getOtherUserEmail(client, task.Assignee.ID)
					if err != nil {
						task.Assignee.Name = "Error"
					}
					(*uidToName)[task.Assignee.ID] = assigneeName
					task.Assignee.Name = assigneeName
				}
			} else {
				task.Assignee = &asana.User{
					Name: "Unassigned",
				}
			}
			taskItem := buildTaskItem(task, projectId)
			(*taskItems) = append((*taskItems), taskItem)
		case !strings.HasSuffix(mode, "_all") && task.Assignee != nil && task.Assignee.ID == uid:
			taskItem := buildTaskItem(task, projectId)
			(*taskItems) = append((*taskItems), taskItem)
		}
	}
}

func buildTaskItem(task *asana.Task, projectId string) *TaskItem {
	dueOnString := ""
	if task.DueOn != nil {
		dueOn := time.Time(*task.DueOn)
		currentYear, _, _ := time.Now().Date()
		if currentYear != dueOn.Year() {
			dueOnString = dueOn.Format("Jan 2 2006")
		} else {
			dueOnString = dueOn.Format("Jan 2")
		}
	}

	assignString := ""
	if task.Assignee != nil {
		assignString = task.Assignee.Name
	}

	taskItem := &TaskItem{
		name:        task.Name,
		id:          task.ID,
		numSubtasks: task.NumSubtasks,
		dueOn:       dueOnString,
		url:         fmt.Sprintf("https://app.asana.com/0/%s/%s/f", projectId, task.ID),
		taskType:    TASK_TYPE,
		completed:   *task.Completed,
		assignee:    assignString,
	}

	return taskItem

}

func getOtherUserEmail(client *asana.Client, uid string) (string, error) {
	if uid == "" {
		return "", fmt.Errorf("missing uid")
	}

	u := &asana.User{
		ID: uid,
	}

	err := u.Fetch(client, nil)
	if err != nil {
		return "", fmt.Errorf("error fetching user: %s", err)
	}

	return u.Email, nil
}

func getCurrentUserId(client *asana.Client, mode string) (string, error) {
	if strings.HasSuffix(mode, "_all") {
		return "", nil
	}
	u, err := client.CurrentUser()
	if err != nil {
		return "", fmt.Errorf("error getting current user: %s", err)
	}

	return u.ID, nil
}

func findSection(client *asana.Client, project *asana.Project, sectionName string) (string, error) {
	sectionId := ""

	sections, _, err := project.Sections(client, &asana.Options{
		Limit: 100,
	})
	if err != nil {
		return "", fmt.Errorf("error getting sections: %s", err)
	}

	for _, section := range sections {
		if section.Name == sectionName {
			sectionId = section.ID
			break
		}
	}

	if sectionId == "" {
		return "", fmt.Errorf("we didn't find the section %s", sectionName)
	}

	return sectionId, nil
}

func getTasksFromAsana(client *asana.Client, q *asana.TaskQuery) ([]*asana.Task, bool, error) {
	moreTasks := false

	tasks, np, err := client.QueryTasks(q, &asana.Options{
		Limit: 100,
		Fields: []string{
			"assignee",
			"name",
			"num_subtasks",
			"due_on",
			"completed",
		},
	})

	if err != nil {
		return nil, false, fmt.Errorf("error querying tasks: %s", err)
	}

	if np != nil {
		moreTasks = true
	}

	return tasks, moreTasks, nil
}
