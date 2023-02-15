package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

type Client struct {
	*analytics.Service

	spec *Spec

	// Fields set by multiplexers
	Account     *analytics.Account
	WebProperty *analytics.Webproperty

	current  *accountInfo
	accounts []*accountInfo

	clientOptions []option.ClientOption
	callOptions   []googleapi.CallOption
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) ID() string {
	res := "google-analytics"
	if c.Account == nil {
		return res
	}

	res += ":account:{" + c.Account.Id + "}"
	if c.WebProperty == nil {
		return res
	}

	return res + ":web-property:{" + c.WebProperty.Id + "}"
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, options source.Options) (schema.ClientMeta, error) {
	panic("todo")
}
