// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/qanpb/json/client/filters"
	"github.com/percona/pmm/api/qanpb/json/client/metrics_names"
	"github.com/percona/pmm/api/qanpb/json/client/object_details"
	"github.com/percona/pmm/api/qanpb/json/client/profile"
)

// Default PMM QAN HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new PMM QAN HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PMMQAN {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new PMM QAN HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PMMQAN {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new PMM QAN client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PMMQAN {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PMMQAN)
	cli.Transport = transport
	cli.Filters = filters.New(transport, formats)
	cli.MetricsNames = metrics_names.New(transport, formats)
	cli.ObjectDetails = object_details.New(transport, formats)
	cli.Profile = profile.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PMMQAN is a client for PMM QAN
type PMMQAN struct {
	Filters filters.ClientService

	MetricsNames metrics_names.ClientService

	ObjectDetails object_details.ClientService

	Profile profile.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PMMQAN) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Filters.SetTransport(transport)
	c.MetricsNames.SetTransport(transport)
	c.ObjectDetails.SetTransport(transport)
	c.Profile.SetTransport(transport)
}
