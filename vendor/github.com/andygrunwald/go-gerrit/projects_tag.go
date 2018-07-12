package gerrit

import (
	"fmt"
	"net/url"
)

// TagInfo entity contains information about a tag.
type TagInfo struct {
	Ref      string        `json:"ref"`
	Revision string        `json:"revision"`
	Object   string        `json:"object"`
	Message  string        `json:"message"`
	Tagger   GitPersonInfo `json:"tagger"`
	Created  *Timestamp    `json:"created,omitempty"`
}

// ListTags list the tags of a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#list-tags
func (s *ProjectsService) ListTags(projectName string, opt *ProjectBaseOptions) (*[]TagInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/tags/", url.QueryEscape(projectName))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]TagInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetTag retrieves a tag of a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-tag
func (s *ProjectsService) GetTag(projectName, tagName string) (*TagInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/tags/%s", url.QueryEscape(projectName), url.QueryEscape(tagName))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(TagInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}
