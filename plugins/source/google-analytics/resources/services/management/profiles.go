package management

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"google.golang.org/api/analytics/v3"
)

func profiles() *schema.Table {
	return &schema.Table{
		Name: "google_analytics_profiles",
		Transform: transformers.TransformWithStruct(new(analytics.Profile),
			transformers.WithPrimaryKeys("AccountId", "WebPropertyId", "Id")),
		Resolver: resolveProfile,
	}
}

func resolveProfile(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	for _, wp := range meta.(*client.Client).Account.WebProperties {
		res <- wp.Profiles
	}
	return nil
}
