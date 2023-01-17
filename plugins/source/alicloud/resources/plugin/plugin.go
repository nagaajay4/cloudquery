package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/bss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/resources/services/oss"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"alicloud",
		Version,
		// Note:  this list should only include top-level tables
		[]*schema.Table{
			bss.BillOverview(),
			bss.Bill(),
			oss.Buckets(),
		},
		client.New,
	)
}