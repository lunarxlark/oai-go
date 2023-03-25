package model

import (
	"strings"
	"time"
)

type ListAPIResponse struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

type Model struct {
	ID          string       `json:"id"`
	Object      string       `json:"object"`
	CreatedAt   int64        `json:"created"`
	OwnedBy     string       `json:"owned_by"`
	Permissoins []Permission `json:"permission"`
	Root        string       `json:"root"`
	Parent      *string      `json:"parent"`
}

func Header() string {
	return strings.Join([]string{"ID", "Created", "OwnedBy"}, "\t")
}

func (m Model) String() string {
	return m.ID + "\t" +
		time.Unix(m.CreatedAt, 0).Format("2006/01/02 15:04:05") + "\t" +
		m.OwnedBy
}

type Permission struct {
	ID                 string  `json:"id"`
	Object             string  `json:"object"`
	CreatedAt          int     `json:"created"`
	AllowCreateEngine  bool    `json:"allow_create_engine"`
	AllowSampling      bool    `json:"allow_sampling"`
	AllowLogprobs      bool    `json:"allow_logprobs"`
	AllowSearchIndices bool    `json:"allow_search_indices"`
	AllowView          bool    `json:"allow_view"`
	AllowFineTuning    bool    `json:"allow_fine_tuning"`
	Organizaton        string  `json:"organization"`
	Group              *string `json:"group"`
	IsBlocking         bool    `json:"is_blocking"`
}
