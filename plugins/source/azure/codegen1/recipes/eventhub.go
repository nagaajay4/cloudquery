// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"

func Armeventhub() []*Table {
	tables := []*Table{
		{
			NewFunc:        armeventhub.NewNamespacesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/namespaces",
			Namespace:      "Microsoft.EventHub",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_EventHub)`,
			Pager:          `NewListPager`,
			ResponseStruct: "NamespacesClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armeventhub())
}