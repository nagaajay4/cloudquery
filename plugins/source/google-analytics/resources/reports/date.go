package reports

import (
	"context"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
)

const dateLayout = "2006-01-02"

func currDate() string {
	return time.Now().UTC().Format(dateLayout)
}

func getStartDate(ctx context.Context, c *client.Client, table string) (string, error) {
	res, err := c.Backend.Get(ctx, table, c.ID())
	if err != nil {
		return "", err
	}

	if len(res) > 0 {
		return res, nil
	}

	return c.StartDate, nil
}
