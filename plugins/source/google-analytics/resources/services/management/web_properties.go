package management

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"google.golang.org/api/analytics/v3"
)

func webProperties() *schema.Table {
	return &schema.Table{
		Name: "google_analytics_web_properties",
		Transform: transformers.TransformWithStruct(new(analytics.Webproperty),
			transformers.WithPrimaryKeys("AccountId", "Id")),
		Resolver:  resolveWebProperties,
		Relations: schema.Tables{profiles()},
	}
}

func resolveWebProperties(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	res <- meta.(*client.Client).Account.WebProperties
	return nil
}
