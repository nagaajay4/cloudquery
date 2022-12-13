// Code generated by codegen2; DO NOT EDIT.
package keyvault

import (
	"encoding/json"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createKeyvault(router *mux.Router) error {
	var item armkeyvault.VaultsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resources", func(w http.ResponseWriter, r *http.Request) {
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
	createKeyvaultKeys(router)
	createKeyvaultSecrets(router)
	return nil
}

func TestKeyvault(t *testing.T) {
	client.MockTestHelper(t, Keyvault(), createKeyvault)
}