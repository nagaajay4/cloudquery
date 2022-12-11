// Code generated by codegen2; DO NOT EDIT.
package appservice

import (
	"encoding/json"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createResourceHealthMetadata(router *mux.Router) error {
	var item armappservice.ResourceHealthMetadataClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	return nil
}

func TestResourceHealthMetadata(t *testing.T) {
	client.MockTestHelper(t, ResourceHealthMetadata(), createResourceHealthMetadata)
}