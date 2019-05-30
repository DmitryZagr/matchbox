package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_MissingEndpoints(t *testing.T) {
	cfg := &Config{
		Endpoints: []string{},
	}
	client, err := New(cfg)
	assert.Nil(t, client)
	assert.Equal(t, errNoEndpoints, err)
}

// gRPC expects host:port with no scheme (e.g. 192.168.2.2:8081)
func TestNew_InvalidEndpoints(t *testing.T) {
	invalid := []string{
		"192.168.2.2",
		"http://192.168.2.2:8081",
		"https://192.168.2.2:8081",
	}

	for _, endpoint := range invalid {
		client, err := New(&Config{
			Endpoints: []string{endpoint},
		})
		assert.Nil(t, client)
		assert.Error(t, err)
	}
}
