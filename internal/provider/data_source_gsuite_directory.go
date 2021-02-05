package provider

import (
	"context"
	"fmt"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJumpCloudGSuiteDirectory() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about a JumpCloud G Suite directory.",
		Read:        dataSourceJumpCloudGSuiteDirectoryRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The user defined name, e.g. `My G Suite directory`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The directory type. This will always be `g_suite`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceJumpCloudGSuiteDirectoryRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*jcapiv2.Configuration)
	client := jcapiv2.NewAPIClient(config)

	directories, _, err := client.DirectoriesApi.DirectoriesList(
		context.TODO(), "", "", nil)
	if err != nil {
		return err
	}

	// there can only be a single GSuite directory per JumpCloud account
	// TODO this is no longer true, there can be multiple G Suite directories per account!
	for _, dir := range directories {
		if dir.Type_ == "g_suite" {
			d.SetId(dir.Id)
			d.Set("name", dir.Name)
			d.Set("type", dir.Type_)
		}
		return nil
	}

	return fmt.Errorf("couldn't find a directory with type 'g_suite'")
}
