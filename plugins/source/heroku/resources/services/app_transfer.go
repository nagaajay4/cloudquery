package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func AppTransfers() *schema.Table {
	return &schema.Table{
		Name:        "heroku_app_transfers",
		Description: `https://devcenter.heroku.com/articles/platform-api-reference#app-transfer`,
		Resolver:    fetchAppTransfers,
		Transform:   transformers.TransformWithStruct(&heroku.AppTransfer{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchAppTransfers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange) // nolint:revive,staticcheck
		v, err := c.Heroku.AppTransferList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}