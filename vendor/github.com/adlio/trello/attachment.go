// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

type Attachment struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	Pos       float32             `json:"pos"`
	Bytes     int                 `json:"int"`
	Date      string              `json:"date"`
	EdgeColor string              `json:"edgeColor"`
	IDMember  string              `json:"idMember"`
	IsUpload  bool                `json:"isUpload"`
	MimeType  string              `json:"mimeType"`
	Previews  []AttachmentPreview `json:"previews"`
	URL       string              `json:"url"`
}

type AttachmentPreview struct {
	ID     string `json:"_id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Bytes  int    `json:"bytes"`
	Scaled bool   `json:"scaled"`
}
