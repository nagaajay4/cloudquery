package client

import (
	"context"

	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/analytics/v3"
)

type (
	AccountInfo struct {
		*analytics.Account
		WebProperties []*WebProperty
	}
	WebProperty struct {
		*analytics.Webproperty
		Profiles []*analytics.Profile
	}
)

func (c *Client) initAccounts(ctx context.Context, ids []string) error {
	eg, ctx := errgroup.WithContext(ctx)
	var accs []*AccountInfo

	req := c.Service.Management.Accounts.List().Context(ctx)
	for {
		resp, err := req.Do(c.callOptions...)
		if err != nil {
			return err
		}

		for _, a := range resp.Items {
			if len(ids) == 0 || slices.Contains(ids, a.Id) {
				info := &AccountInfo{Account: a}
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

func (c *Client) initAccountInfo(ctx context.Context, ai *AccountInfo) error {
	eg, ctx := errgroup.WithContext(ctx)
	ai.WebProperties = nil

	req := c.Service.Management.Webproperties.List(ai.Account.Id).Context(ctx)
	for {
		resp, err := req.Do(c.callOptions...)
		if err != nil {
			return err
		}

		for _, p := range resp.Items {
			property := &WebProperty{Webproperty: p}
			eg.Go(func() error {
				return c.initProfiles(ctx, property)
			})
			ai.WebProperties = append(ai.WebProperties, property)
		}

		if resp.NextLink == "" {
			break
		}

		// analytics pkg isn't compatible with iterator pkg, so we use offset instead
		req.StartIndex(resp.StartIndex + int64(len(resp.Items)))
	}

	return eg.Wait()
}

func (c *Client) initProfiles(ctx context.Context, wp *WebProperty) error {
	wp.Profiles = nil

	req := c.Service.Management.Profiles.List(wp.AccountId, wp.Id).Context(ctx)
	for {
		resp, err := req.Do(c.callOptions...)
		if err != nil {
			return err
		}

		wp.Profiles = append(wp.Profiles, resp.Items...)

		if resp.NextLink == "" {
			break
		}

		// analytics pkg isn't compatible with iterator pkg, so we use offset instead
		req.StartIndex(resp.StartIndex + int64(len(resp.Items)))
	}

	return nil
}
