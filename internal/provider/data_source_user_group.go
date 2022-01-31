package provider

import (
	"context"
	"encoding/json"
	"net/http"

	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about a JumpCloud System User.",
		ReadContext: dataSourceGroupRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": {
				Description: "The name of the group. Example: `My Group`.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceGroupRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*jcapiv2.Configuration)

	groupIdData, groupIdOk := d.GetOk("id")
	groupNameData, groupNameOk := d.GetOk("name")

	if groupIdOk {
		groupId := groupIdData.(string)

		group, ok, err := userGroupReadHelper(config, groupId)

		if err != nil {
			return diag.FromErr(err)
		}

		if !ok {
			// not found
			d.SetId("")
			return nil
		}

		d.SetId(group.ID)
		if err := d.Set("name", group.Name); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("type", group.Type); err != nil {
			return diag.FromErr(err)
		}

		return nil
	} else if groupNameOk {

		groupName := groupNameData.(string)

		group, ok, err := GroupReadHelperName(config, groupName)
		if err != nil {
			return diag.FromErr(err)
		}
		if len(group) <= 0 {
			return diag.Errorf(groupName + " group name not found!")
		}

		if !ok {
			// not found
			d.SetId("")
			return nil
		}

		gr := group[0]

		d.SetId(gr.ID)
		if err := d.Set("name", gr.Name); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("type", gr.Type); err != nil {
			return diag.FromErr(err)
		}

		return nil
	} else {
		return diag.Errorf("one of the following must be set: id, name")
	}

}

func GroupReadHelperName(config *jcapiv2.Configuration, name string) (ug []*UserGroup,
	ok bool, err error) {

	req, err := http.NewRequest(http.MethodGet,
		config.BasePath+"/usergroups/?filter=name:eq:"+name, nil)
	if err != nil {
		return
	}

	req.Header.Add("x-api-key", config.DefaultHeader["x-api-key"])
	if config.DefaultHeader["x-org-id"] != "" {
		req.Header.Add("x-org-id", config.DefaultHeader["x-org-id"])
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return
	}

	ok = true
	err = json.NewDecoder(res.Body).Decode(&ug)
	return
}
