// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"

func Armrelay() []*Table {
	tables := []*Table{
		{
			NewFunc:        armrelay.NewNamespacesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Relay/namespaces",
			Namespace:      "Microsoft.Relay",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Relay)`,
			Pager:          `NewListPager`,
			ResponseStruct: "NamespacesClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armrelay())
}