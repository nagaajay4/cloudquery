package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_MONGODB_TEST_CONN")
	if testConn == "" {
		return "mongodb://localhost:27017"
	}
	return testConn
}

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("mongodb", "development", New, destination.WithManagedWriter())
		},
		Spec{
			ConnectionString: getTestConnection(),
			Database:         "destination_mongodb_test",
		},
		destination.PluginTestSuiteTests{})
}
