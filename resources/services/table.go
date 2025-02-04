package services

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/mnorbury/cq-source-apod/client"
)

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:     "apod_sample_table",
		Resolver: fetchSampleTable,
		Columns: []schema.Column{
			{
				Name: "column",
				Type: arrow.BinaryTypes.String,
			},
		},
	}
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	return fmt.Errorf("not implemented. client id: " + cl.ID())
}
