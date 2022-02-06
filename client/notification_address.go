package client

import (
	"context"
	"encoding/json"
	"fmt"
)

type NotificationAddress struct {
	Address               string `json:"address"`
	NotificationChannelID string `json:"notification_channel_id"`
	GroupID               int    `json:"group_id"`
	Description           string `json:"description,omitempty"`
	OwnerID               int    `json:"owner_id"`
}

type CreateNotificationAddressResponse struct {
	ID int `json:"id"`
}

func (c Client) CreateNotificationAddress(ctx context.Context, req NotificationAddress) (int, error) {
	resp, err := c.post(ctx, "/notification_addresses", req)
	if err != nil {
		return 0, err
	}

	var response CreateNotificationAddressResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return 0, fmt.Errorf("can't unamrshal %v: %w", string(resp), err)
	}

	return response.ID, err
}

func (c Client) ReadNotificationAddress(ctx context.Context, id int) (NotificationAddress, error) {
	resp, err := c.get(ctx, fmt.Sprintf("/notification_addresses/%v", id))
	if err != nil {
		return NotificationAddress{}, err
	}

	var notificationAddress NotificationAddress
	err = json.Unmarshal(resp, &notificationAddress)
	if err != nil {
		return NotificationAddress{}, err
	}

	return notificationAddress, nil
}

func (c Client) UpdateNotificationAddress(ctx context.Context, id int, req NotificationAddress) error {
	return c.put(ctx, fmt.Sprintf("/notification_addresses/%v", id), req)
}

func (c Client) DeleteNotificationAddress(ctx context.Context, id int) error {
	return c.delete(ctx, fmt.Sprintf("/notification_addresses/%v", id))
}
