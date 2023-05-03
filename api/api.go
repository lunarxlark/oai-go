package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/lunarxlark/openai-cli/config"
)

func Request(method, url string, req, res any) error {
	var payload []byte
	if req != nil {
		var err error
		payload, err = json.Marshal(req)
		if err != nil {
			return err
		}
	}

	rq, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	if req != nil {
		rq.Header.Set("Content-Type", "application/json")
	}
	rq.Header.Set("Authorization", "Bearer "+config.OAIConfig.APIKey)

	client := new(http.Client)
	rs, err := client.Do(rq)
	if err != nil {
		return err
	}

	return json.NewDecoder(rs.Body).Decode(res)
}
