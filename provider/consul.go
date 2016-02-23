package provider

import (
	"fmt"

	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"
	"github.com/emilevauge/traefik/types"
)

// Consul holds configurations of the Consul provider.
type Consul struct {
	Kv  `mapstructure:",squash"`
	TLS *KvTLS
}

// Provide allows the provider to provide configurations to traefik
// using the given configuration channel.
func (provider *Consul) Provide(configurationChan chan<- types.ConfigMessage) error {

	fmt.Printf("1: %q\n", provider.TLS)

	provider.storeType = store.CONSUL
	consul.Register()
	return provider.provide(configurationChan)
}
