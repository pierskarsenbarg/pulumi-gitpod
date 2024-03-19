package pkg

import (
	"net/http"
	"time"

	gitpodapi "github.com/pierskarsenbarg/pulumi-gitpod/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	AccessToken    string `pulumi:"accessToken,optional" provider:"secret"`
	OwnerId        string `pulumi:"ownerId,optional"`
	OrganizationId string `pulumi:"organizationId,optional"`
	Client         gitpodapi.Client
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.AccessToken, "Your Gitpod access token. You can create these in your user settings: https://gitpod.io/user/tokens")
	a.Describe(&c.OwnerId, "Id of owner account. This can be found on your user account page: https://gitpod.io/user/account")
	a.Describe(&c.OrganizationId, "Id of the organisation who is going to own the workspaces. You can find this on the organisation settings page of the currently selected organisation: https://gitpod.io/settings")
	a.SetDefault(&c.AccessToken, "", "GITPOD_ACCESSTOKEN")
	a.SetDefault(&c.OrganizationId, "", "GITPOD_ORGANISATIONID")
	a.SetDefault(&c.OwnerId, "", "GITPOD_OWNERID")
}

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx p.Context) error {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	client, err := gitpodapi.NewClient(&httpClient, c.AccessToken, "")
	if err != nil {
		return err
	}

	c.Client = *client

	return nil
}
