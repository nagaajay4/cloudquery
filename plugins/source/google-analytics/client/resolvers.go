package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func ResolveAccountID(_ context.Context, m schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, m.(*Client).Account.Id)
}

func ResolveWebPropertyID(_ context.Context, m schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, m.(*Client).WebProperty.Id)
}
