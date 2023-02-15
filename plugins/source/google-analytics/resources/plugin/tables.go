package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/resources/services/management"
	"github.com/cloudquery/plugin-sdk/schema"
)

func tables() schema.Tables {
	return schema.Tables{
		management.Accounts(),
	}
}
