package store

import "context"

func (cl *Client) Fetch(key string) (string, error) {

	ctx := context.Background()

	return cl.conn.Get(ctx, key).Result()

}