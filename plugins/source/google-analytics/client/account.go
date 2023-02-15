package client

import (
	"context"

	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/analytics/v3"
)

type accountInfo struct {
	account       *analytics.Account
	webProperties []*analytics.Webproperty
}

func (c *Client) initAccounts(ctx context.Context, ids []string) error {
	eg, ctx := errgroup.WithContext(ctx)
	var accs []*accountInfo

	req := c.Service.Management.Accounts.List().Context(ctx)
	for {
		resp, err := req.Do(c.callOptions...)
		if err != nil {
			return err
		}

		for _, a := range resp.Items {
			if len(ids) == 0 || slices.Contains(ids, a.Id) {
				info := &accountInfo{account: a}
				accs = append(accs, info)
				eg.Go(func() error { return c.initAccountInfo(ctx, info) })
			}
		}

		if resp.NextLink == "" {
			break
		}

		// analytics pkg isn't compatible with iterator pkg, so we use offset instead
		req.StartIndex(resp.StartIndex + int64(len(resp.Items)))
	}

	return eg.Wait()
}

func (c *Client) initAccountInfo(ctx context.Context, ai *accountInfo) error {
	ai.webProperties = nil

	req := c.Service.Management.Webproperties.List(ai.account.Id).Context(ctx)
	for {
		resp, err := req.Do(c.callOptions...)
		if err != nil {
			return err
		}

		ai.webProperties = append(ai.webProperties, resp.Items...)

		if resp.NextLink == "" {
			break
		}

		// analytics pkg isn't compatible with iterator pkg, so we use offset instead
		req.StartIndex(resp.StartIndex + int64(len(resp.Items)))
	}

	return nil
}
