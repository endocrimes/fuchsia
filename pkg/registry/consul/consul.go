package consul

import (
	"fmt"

	"github.com/endocrimes/fuchsia/pkg/registry"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
)

var _ registry.Registry = (*Registry)(nil)

type Config struct {
	ConsulAddr        string
	ConsulServiceName string
}

type Registry struct {
	cfg    *Config
	client *api.Client
	csvc   *connect.Service
}

func (r *Registry) Name() string {
	return "consul"
}

func (r *Registry) Init() error {
	c, err := r.createClient()
	if err != nil {
		return err
	}

	csvc, err := r.createConnectService(c)
	if err != nil {
		return err
	}

	r.client = c
	r.csvc = csvc

	return nil
}

func (r *Registry) createClient() (*api.Client, error) {
	return nil, fmt.Errorf("Unimplemented")
}

func (r *Registry) createConnectService(client *api.Client) (*connect.Service, error) {
	return connect.NewService(r.cfg.ConsulServiceName, client)
}

func (r *Registry) WatchServiceUpdates() chan registry.ServiceUpdate {
	return nil
}

func (r *Registry) Shutdown() error {
	return nil
}
