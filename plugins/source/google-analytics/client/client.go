package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/option"
)

type Client struct {
	*analytics.Service
	backend.Backend

	accounts []*AccountInfo

	// Fields set by multiplexers
	Account     *AccountInfo
	WebProperty *WebProperty
	Profile     *analytics.Profile
	logger      zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	res := "google-analytics"
	if c.Account == nil {
		return res
	}

	res += ":account:{" + c.Account.Id + "}"
	if c.WebProperty == nil {
		return res
	}

	res += ":web-property:{" + c.WebProperty.Id + "}"
	if c.Profile == nil {
		return res
	}

	return res + ":profile:{" + c.Profile.Id + "}"
}

func Configure(ctx context.Context, logger zerolog.Logger, srcSpec specs.Source, options source.Options) (schema.ClientMeta, error) {
	var spec Spec
	if err := srcSpec.UnmarshalSpec(&spec); err != nil {
		return nil, err
	}

	opts := []option.ClientOption{
		option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the same side with telemetry
		option.WithTelemetryDisabled(),
	}
	if len(spec.APIKey) > 0 {
		opts = append(opts, option.WithAPIKey(spec.APIKey))
	}

	svc, err := analytics.NewService(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	svc.UserAgent = "cloudquery:source-google-analytics/" + srcSpec.Version

	c := &Client{
		Service: svc,
		Backend: options.Backend,
		logger:  logger.With().Str("plugin", "google-analytics").Logger(),
	}

	return c, c.initAccounts(ctx, spec.AccountIDs)
}
