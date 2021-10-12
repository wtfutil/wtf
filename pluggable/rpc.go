package pluggable

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type RPCPluggable interface {
	Module() (ExternalModule, error)
}

type RPCPluggablePlugin struct {
	Impl RPCPluggable
}

func (p *RPCPluggablePlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (p *RPCPluggablePlugin) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

type RPCServer struct {
	Impl RPCPluggable
}

func (s *RPCServer) Module(_ map[string]interface{}, resp *ExternalModule) error {
	res, err := s.Impl.Module()
	if err != nil {
		return err
	}

	resp = &res

	return nil
}

type RPCClient struct {
	client *rpc.Client
}

func (c *RPCClient) Module() (p RPCPluggablePlugin, err error) {
	err = c.client.Call("Plugin.Module", nil, &p)

	return
}
