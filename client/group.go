package client

import (
	"context"
	"encoding/json"
	"fmt"
)

type Group struct {
	Name    string `json:"name"`
	OwnerID int    `json:"owner_id"`
}

type CreateGroupResponse struct {
	ID int `json:"id"`
}

func (c Client) CreateGroup(ctx context.Context, req Group) (int, error) {
	resp, err := c.post(ctx, "/groups", req)
	if err != nil {
		return 0, err
	}

	var response CreateGroupResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return 0, err
	}

	return response.ID, err
}

func (c Client) ReadGroup(ctx context.Context, id int) (Group, error) {
	resp, err := c.get(ctx, fmt.Sprintf("/groups/%v", id))
	if err != nil {
		return Group{}, err
	}

	var group Group
	err = json.Unmarshal(resp, &group)
	if err != nil {
		return Group{}, err
	}

	return group, nil
}

func (c Client) UpdateGroup(ctx context.Context, id int, req Group) error {
	return c.put(ctx, fmt.Sprintf("/groups/%v", id), req)
}

func (c Client) DeleteGroup(ctx context.Context, id int) error {
	return c.delete(ctx, fmt.Sprintf("/groups/%v", id))
}
