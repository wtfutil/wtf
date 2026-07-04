package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/gopherlibs/todoist/api"
	"github.com/olebedev/config"
)

type todoistFilter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Query string `json:"query"`
}

type todoistProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type syncResponse struct {
	Filters []todoistFilter `json:"filters"`
}

type filterTaskItem struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Content   string `json:"content"`
	Checked   bool   `json:"checked"`
	Due       *struct {
		Date     string `json:"date"`
		Datetime string `json:"datetime"`
	} `json:"due"`
}

type Todoist struct {
	client         *api.Client
	apiKey         string
	projects       []interface{}
	filters        []interface{}
	filterMap      map[string]todoistFilter
	projectNameMap map[string]string
}

func (todo *Todoist) Title() string {
	return "Todoist"
}

func (todo *Todoist) Setup(config *config.Config) {
	todo.apiKey = config.UString("apiKey")
	todo.client = api.New(todo.apiKey)
	todo.projects = config.UList("projects")
	todo.filters = config.UList("filters")
	todo.fetchProjects()
	todo.fetchFilters()
}

func (todo *Todoist) doGet(urlStr string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Add("Authorization", "Bearer "+todo.apiKey)
	return http.DefaultClient.Do(req)
}

func (todo *Todoist) doPostForm(path string, body io.Reader) (*http.Response, error) {
	u, _ := url.Parse("https://api.todoist.com" + path)
	req, _ := http.NewRequest("POST", u.String(), body)
	req.Header.Add("Authorization", "Bearer "+todo.apiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return http.DefaultClient.Do(req)
}

func (todo *Todoist) fetchProjects() {
	todo.projectNameMap = make(map[string]string)

	projects, err := todo.client.Projects()
	if err != nil {
		return
	}

	for _, p := range projects.Results {
		todo.projectNameMap[p.ID] = p.Name
	}
}

func (todo *Todoist) fetchFilters() {
	todo.filterMap = make(map[string]todoistFilter)

	resp, err := todo.doPostForm("/api/v1/sync", strings.NewReader("sync_token=*&resource_types=%5B%22filters%22%5D"))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var syncResp syncResponse
	if err := json.Unmarshal(body, &syncResp); err != nil {
		return
	}

	for _, f := range syncResp.Filters {
		todo.filterMap[f.ID] = f
	}
}

func (todo *Todoist) BuildProjects() []*Project {
	projects := []*Project{}

	for _, id := range todo.projects {
		i := fmt.Sprintf("%v", id)
		proj := todo.GetProject(i)
		projects = append(projects, proj)
	}

	for _, f := range todo.filters {
		filterID := fmt.Sprintf("%v", f)
		proj := todo.GetFilterProject(filterID)
		projects = append(projects, proj)
	}

	return projects
}

func (todo *Todoist) GetProject(id string) *Project {
	proj := &Project{
		Index:   -1,
		backend: todo,
	}

	proj.ID = id
	proj.Name = "Error"

	p, err := todo.client.Project(id)
	if err != nil {
		return proj
	}

	proj.Name = p.Name

	tasks, err := todo.LoadTasks(proj.ID)
	proj.Err = err
	proj.Tasks = tasks

	return proj
}

func (todo *Todoist) GetFilterProject(filterID string) *Project {
	proj := &Project{
		Index:   -1,
		backend: todo,
	}
	proj.ID = "filter:" + filterID
	proj.Name = "Filter #" + filterID

	f, ok := todo.filterMap[filterID]
	if !ok {
		proj.Err = fmt.Errorf("filter %s not found (loaded %d filters)", filterID, len(todo.filterMap))
		return proj
	}

	proj.Name = f.Name
	tasks, err := todo.LoadTasksByFilter(f.Query)
	proj.Err = err
	proj.Tasks = tasks

	return proj
}

func toTask(task api.Task) Task {
	return Task{
		ID:        task.ID,
		Completed: task.Checked,
		Name:      task.Content,
	}
}

func (todo *Todoist) LoadTasks(id string) ([]Task, error) {
	if strings.HasPrefix(id, "filter:") {
		filterID := strings.TrimPrefix(id, "filter:")
		f, ok := todo.filterMap[filterID]
		if !ok {
			return nil, fmt.Errorf("filter %s not found", filterID)
		}
		return todo.LoadTasksByFilter(f.Query)
	}

	tasks, err := todo.client.Tasks(id)
	if err != nil {
		return nil, err
	}

	var finalTasks []Task
	for _, item := range tasks.Results {
		finalTasks = append(finalTasks, toTask(item))
	}
	return finalTasks, nil
}

func (todo *Todoist) LoadTasksByFilter(query string) ([]Task, error) {
	u, _ := url.Parse("https://api.todoist.com/api/v1/tasks/filter")
	params := u.Query()
	params.Set("query", query)
	params.Set("limit", "200")
	u.RawQuery = params.Encode()

	resp, err := todo.doGet(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("tasks/filter API returned %d: %s", resp.StatusCode, string(body))
	}

	var tasksResp struct {
		Results []filterTaskItem `json:"results"`
	}
	err = json.Unmarshal(body, &tasksResp)
	if err != nil {
		return nil, err
	}

	type taskWithDate struct {
		task     Task
		datetime *time.Time
	}

	var withDate []taskWithDate
	var withoutDate []Task

	for _, item := range tasksResp.Results {
		projName := todo.projectNameMap[item.ProjectID]

		prefix := ""
		if projName != "" {
			prefix = "#" + projName
		}

		suffix := ""
		var dt *time.Time

		if item.Due != nil && item.Due.Date != "" {
			var t time.Time
			var err error
			for _, layout := range []string{time.RFC3339, "2006-01-02T15:04:05", "2006-01-02"} {
				t, err = time.Parse(layout, item.Due.Date)
				if err == nil {
					break
				}
			}
			if err == nil {
				if t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0 {
					suffix = "  " + t.Format("Jan 2")
				} else {
					suffix = "  " + t.Format("Jan 2 15:04")
				}
				dt = &t
			}
		}

		overdue := false
		if dt != nil {
			dueUTC := dt.UTC()
			now := time.Now().UTC()
			if dueUTC.Before(now) {
				if dueUTC.Hour() == 0 && dueUTC.Minute() == 0 && dueUTC.Second() == 0 {
					if now.Year() > dueUTC.Year() || now.YearDay() > dueUTC.YearDay() {
						overdue = true
					}
				} else {
					overdue = true
				}
			}
		}

		t := Task{
			ID:         item.ID,
			Completed:  item.Checked,
			Name:       item.Content,
			Prefix:     prefix,
			DateSuffix: suffix,
			Overdue:    overdue,
		}

		if dt != nil {
			withDate = append(withDate, taskWithDate{t, dt})
		} else {
			withoutDate = append(withoutDate, t)
		}
	}

	sort.Slice(withDate, func(i, j int) bool {
		return withDate[i].datetime.Before(*withDate[j].datetime)
	})

	var finalTasks []Task
	for _, wd := range withDate {
		finalTasks = append(finalTasks, wd.task)
	}
	finalTasks = append(finalTasks, withoutDate...)

	return finalTasks, nil
}

func (todo *Todoist) CloseTask(task *Task) error {
	if task != nil {
		_, err := todo.client.TaskClose(task.ID)
		return err
	}
	return nil
}

func (todo *Todoist) DeleteTask(task *Task) error {
	if task != nil {
		_, err := todo.client.TaskDelete(task.ID)
		return err
	}
	return nil
}

func (todo *Todoist) Sources() []string {
	var result []string
	for _, id := range todo.projects {
		i := fmt.Sprintf("%v", id)
		result = append(result, i)
	}
	for _, f := range todo.filters {
		filterID := fmt.Sprintf("%v", f)
		result = append(result, "filter:"+filterID)
	}
	return result
}
