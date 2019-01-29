package mcpacknpc

import (
	"bytes"
	"mcpack"
	"npc"
)

type Client struct {
	*npc.Client
}

func NewClient(server []string) *Client {
	c := npc.NewClient(server)
	return &Client{Client: c}
}

func (c *Client) Call(args interface{}, reply interface{}, logId uint32) error {
	content, err := mcpack.Marshal(args)
	if err != nil {
		return err
	}
	resp, err := c.Client.Do(npc.NewRequest(bytes.NewReader(content), logId))
	if err != nil {
		return err
	}
	return mcpack.Unmarshal(resp.Body, reply)
}
