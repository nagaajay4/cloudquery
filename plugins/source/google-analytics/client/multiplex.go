package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccountMX(meta schema.ClientMeta) []schema.ClientMeta {
	c := meta.(*Client)
	res := make([]schema.ClientMeta, len(c.accounts))
	for _, acc := range c.accounts {
		res = append(res, &Client{
			Service:  c.Service,
			Account:  acc.account,
			current:  acc,
			accounts: c.accounts,
		})
	}
	return res
}

func AccountWebPropertyMX(meta schema.ClientMeta) []schema.ClientMeta {
	var res []schema.ClientMeta
	for _, m := range AccountMX(meta) {
		c := m.(*Client)
		for _, webProperty := range c.current.webProperties {
			n := *c
			n.WebProperty = webProperty
			res = append(res, &n)
		}
	}

	return res
}
