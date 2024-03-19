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
	a.Describe(&c.AccessToken, "Your Gitpod access token")
	a.Describe(&c.OwnerId, "Id of owner account")
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
