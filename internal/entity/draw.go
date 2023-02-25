package entity

import (
	"encoding/xml"
	"time"
)

type DrawImage struct {
	Alt       string    `json:"alt"`
	Data      string    `json:"data"`
	Xml       string    `json:"xml"`
	CreatedAt time.Time `json:"created_at"`
	Size      int64     `json:"size"`
}

type DrawFile struct {
	XMLName  xml.Name `xml:"mxfile"`
	Host     string   `xml:"host,attr"`
	Modified string   `xml:"modified,attr"`
	Agent    string   `xml:"agent,attr"`
	Etag     string   `xml:"etag,attr"`
	Version  string   `xml:"version,attr"`
	Type     string   `xml:"type,attr"`
	Pages    int      `xml:"pages,attr"`
	Diagrams []string `xml:"diagram"`
}

// type Diagram struct {
// 	XMLName xml.Name `xml:"diagram"`
// 	Id      string   `xml:"id,attr"`
// 	Name    string   `xml:"name,attr"`
// }

func UnmarshalDiagram(data []byte) (*DrawFile, error) {
	var diagram DrawFile
	if err := xml.Unmarshal(data, &diagram); err != nil {
		return nil, err
	}
	return &diagram, nil
}
