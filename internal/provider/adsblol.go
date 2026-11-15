// Package provider retrieves observed aircraft from adsb.lol.
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const BaseURL = "https://api.adsb.lol"

// Pointers preserve the difference between an absent reading and a real zero.
// Altitude is raw JSON because the provider returns a number or "ground".
type Aircraft struct {
	Hex          string          `json:"hex"`
	Callsign     string          `json:"flight,omitempty"`
	Registration string          `json:"r,omitempty"`
	Type         string          `json:"t,omitempty"`
	Altitude     json.RawMessage `json:"alt_baro,omitempty"`
	Speed        *float64        `json:"gs,omitempty"`
	Track        *float64        `json:"track,omitempty"`
	Latitude     *float64        `json:"lat,omitempty"`
	Longitude    *float64        `json:"lon,omitempty"`
	Seen         *float64        `json:"seen,omitempty"`
	PositionSeen *float64        `json:"seen_pos,omitempty"`
}

type Result struct {
	Source      string     `json:"source"`
	RetrievedAt time.Time  `json:"retrieved_at"`
	SnapshotAt  *time.Time `json:"snapshot_at,omitempty"`
	Aircraft    []Aircraft `json:"aircraft"`
}

type Client struct {
	HTTP    *http.Client
	BaseURL string
	APIKey  string
}

func (c *Client) Fetch(ctx context.Context, path string) (Result, error) {
	result := Result{Source: "adsb.lol", Aircraft: []Aircraft{}}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.TrimRight(c.BaseURL, "/")+path, nil)
	if err != nil {
		return result, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "sky-cli/0.1.0")
	if c.APIKey != "" {
		req.Header.Set("x-api-key", c.APIKey)
	}
	client := c.HTTP
	if client == nil {
		client = &http.Client{Timeout: 15 * time.Second}
	}
	resp, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("request aircraft data: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusTooManyRequests {
		hint := ""
		if retry := resp.Header.Get("Retry-After"); retry != "" {
			hint = " (Retry-After: " + retry + ")"
		}
		return result, fmt.Errorf("adsb.lol rate limit reached%s; wait before retrying or configure SKY_API_KEY", hint)
	}
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("adsb.lol returned HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 16*1024*1024+1))
	if err != nil {
		return result, fmt.Errorf("read response: %w", err)
	}
	if len(body) > 16*1024*1024 {
		return result, fmt.Errorf("aircraft response exceeds 16 MiB")
	}
	var payload struct {
		Aircraft []Aircraft `json:"ac"`
		Now      int64      `json:"now"`
		Message  string     `json:"msg"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return result, fmt.Errorf("decode aircraft response: %w", err)
	}
	if payload.Message != "" && payload.Message != "No error" {
		return result, fmt.Errorf("adsb.lol: %s", payload.Message)
	}
	// A missing ac field is a schema error; an empty array is a successful lookup.
	if payload.Aircraft == nil {
		return result, fmt.Errorf("aircraft response is missing the ac array")
	}
	result.Aircraft = payload.Aircraft
	for i := range result.Aircraft {
		result.Aircraft[i].Callsign = strings.TrimSpace(result.Aircraft[i].Callsign)
	}
	result.RetrievedAt = time.Now().UTC()
	if payload.Now > 0 {
		t := time.UnixMilli(payload.Now).UTC()
		result.SnapshotAt = &t
	}
	return result, nil
}
