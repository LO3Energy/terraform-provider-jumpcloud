package provider

import (
	"context"

	jcapiv1 "github.com/TheJumpCloud/jcapi-go/v1"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about a JumpCloud System User.",
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"username": {
				Description: "The technical user name. See JumpCloud's [user naming conventions](https://support.jumpcloud.com/support/s/article/naming-convention-for-users1) for naming restrictions. Example: `john.doe`.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"email": {
				Description: "The users e-mail address, which is also used for log ins. E-mail addresses have to be unique across all JumpCloud accounts, there cannot be two users with the same e-mail address. Example: `john.doe@acme.org`.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"firstname": {
				Description: "The user's first name. Example: `john`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"lastname": {
				Description: "The user's last name. Example: `doe`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceUserRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	configv1 := convertV2toV1Config(meta.(*jcapiv2.Configuration))
	client := jcapiv1.NewAPIClient(configv1)

	userIDData, userIDOk := d.GetOk("id")
	usernameData, usernameOk := d.GetOk("username")
	emailData, emailOk := d.GetOk("email")

	var result jcapiv1.Systemuserreturn
	var err error

	if userIDOk {
		userID := userIDData.(string)

		res, _, errr := client.SystemusersApi.SystemusersGet(context.TODO(),
			userID, "", "", nil)

		result = res
		err = errr
	} else if emailOk {
		email := emailData.(string)

		res, _, errr := client.SystemusersApi.SystemusersList(context.TODO(),
			"", "", map[string]interface{}{
				"filter": "email:$eq:" + email,
			})

		if len(res.Results) <= 0 {
			return diag.Errorf(email + " Email not found!")
		}

		result = res.Results[0]
		err = errr
	} else if usernameOk {
		username := usernameData.(string)

		res, _, errr := client.SystemusersApi.SystemusersList(context.TODO(),
			"", "", map[string]interface{}{
				"filter": "username:$eq:" + username,
			})

		if len(res.Results) <= 0 {
			return diag.Errorf(username + " Username not found!")
		}

		result = res.Results[0]
		err = errr
	} else {
		return diag.Errorf("one of the following must be set: id, username, email")
	}

	// If the object does not exist in our infrastructure, we unset the ID
	// Unfortunately, the http request returns 200 even if the resource does not exist
	if err != nil {
		if err.Error() == "EOF" {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(result.Id)

	if err := d.Set("username", result.Username); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("email", result.Email); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("firstname", result.Firstname); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("lastname", result.Lastname); err != nil {
		return diag.FromErr(err)
	}

	// indicates that everything went well
	return nil

}
