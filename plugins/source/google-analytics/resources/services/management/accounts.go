package management

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"google.golang.org/api/analytics/v3"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "google_analytics_accounts",
		Transform: transformers.TransformWithStruct(new(analytics.Account), transformers.WithPrimaryKeys("Id")),
		Multiplex: client.AccountMX,
		Resolver:  resolveAccount,
		Relations: schema.Tables{webProperties()},
	}
}

func resolveAccount(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	res <- meta.(*client.Client).Account
	meta.(*client.Client).Management.AccountSummaries.List()
	return nil
}
