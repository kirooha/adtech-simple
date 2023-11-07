package adserverclient

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

const (
	campaignCreationURL = "http://localhost:9091/create-campaign"
)

type campaignCreationRequest struct {
	Name        string
	Description string
}

type campaignCreationResponse struct {
	ID uuid.UUID
}

func (c *Client) CreateCampaign(ctx context.Context, name, description string) (*uuid.UUID, error) {
	var (
		req  campaignCreationRequest
		resp campaignCreationResponse
	)

	req.Name = name
	req.Description = description

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, campaignCreationURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp.ID, nil
}
