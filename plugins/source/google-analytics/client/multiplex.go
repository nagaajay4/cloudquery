package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

// AccountMX provides account-only mx
func AccountMX(meta schema.ClientMeta) []schema.ClientMeta {
	c := meta.(*Client)
	res := make([]schema.ClientMeta, len(c.accounts))
	for _, acc := range c.accounts {
		n := *c
		n.Account = acc
		n.WebProperty = nil
		n.Profile = nil
		res = append(res, &n)
	}
	return res
}

// WebPropertyMX provides account+web property mx
func WebPropertyMX(meta schema.ClientMeta) []schema.ClientMeta {
	var res []schema.ClientMeta
	for _, m := range AccountMX(meta) {
		c := m.(*Client)
		for _, webProperty := range c.Account.WebProperties {
			n := *c
			n.WebProperty = webProperty
			n.Profile = nil
			res = append(res, &n)
		}
	}

	return res
}

// ProfileMX provides account+web property+profile mx
func ProfileMX(meta schema.ClientMeta) []schema.ClientMeta {
	var res []schema.ClientMeta
	for _, m := range WebPropertyMX(meta) {
		c := m.(*Client)
		for _, profile := range c.WebProperty.Profiles {
			n := *c
			n.Profile = profile
			res = append(res, &n)
		}
	}

	return res
}
