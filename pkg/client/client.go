package client

import (
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/venturemark/apigengo/pkg/pbf/update"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
)

type Config struct {
	Address string
}

type Client struct {
	connection *grpc.ClientConn

	audience audience.APIClient
	texupd   texupd.APIClient
	timeline timeline.APIClient
	update   update.APIClient
}

func New(c Config) (*Client, error) {
	if c.Address == "" {
		c.Address = "127.0.0.1:7777"
	}

	var err error

	var con *grpc.ClientConn
	{
		con, err = grpc.Dial(c.Address, grpc.WithInsecure())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var aud audience.APIClient
	{
		aud = audience.NewAPIClient(con)
	}

	var tex texupd.APIClient
	{
		tex = texupd.NewAPIClient(con)
	}

	var tim timeline.APIClient
	{
		tim = timeline.NewAPIClient(con)
	}

	var upd update.APIClient
	{
		upd = update.NewAPIClient(con)
	}

	cli := &Client{
		connection: con,

		audience: aud,
		texupd:   tex,
		timeline: tim,
		update:   upd,
	}

	return cli, nil
}

func (c *Client) Connection() *grpc.ClientConn {
	return c.connection
}

func (c *Client) Audience() audience.APIClient {
	return c.audience
}

func (c *Client) TexUpd() texupd.APIClient {
	return c.texupd
}

func (c *Client) Timeline() timeline.APIClient {
	return c.timeline
}

func (c *Client) Update() update.APIClient {
	return c.update
}
