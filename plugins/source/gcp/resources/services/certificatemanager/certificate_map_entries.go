package certificatemanager

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	certificatemanager "cloud.google.com/go/certificatemanager/apiv1"
)

func CertificateMapEntries() *schema.Table {
	return &schema.Table{
		Name:        "gcp_certificatemanager_certificate_map_entries",
		Description: `https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificateMaps.certificateMapEntries#CertificateMapEntry`,
		Resolver:    fetchCertificateMapEntries,
		Multiplex:   client.ProjectMultiplexEnabledServices("certificatemanager.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.CertificateMapEntry{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchCertificateMapEntries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListCertificateMapEntriesRequest{
		Parent: parent.Item.(*pb.CertificateMap).Name,
	}
	gcpClient, err := certificatemanager.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListCertificateMapEntries(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}