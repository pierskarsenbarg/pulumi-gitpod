package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	gitpodapi "github.com/pierskarsenbarg/pulumi-gitpod/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var Version string

func main() {
	err := p.RunProvider("gitpod", "0.1.0", provider())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func provider() p.Provider {
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{infer.Resource[*Organization, OrganizationArgs, OrganizationState]()},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"gitpod": "index",
		},
		Functions: []infer.InferredFunction{
			infer.Function[*GetOrganization, GetOrganizationArgs, OrganizationState](),
		},
		Config: infer.Config[*Config](),
	})
}

type Config struct {
	AccessToken string `pulumi:"accessToken,optional" provider:"secret"`
	client      gitpodapi.Client
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.AccessToken, "Your Gitpod access token")
	// a.SetDefault(&c.AccessToken, "", "GITPOD_ACCESSTOKEN")
}

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx p.Context) error {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	accessToken := c.AccessToken

	client, err := gitpodapi.NewClient(&httpClient, accessToken, "")
	if err != nil {
		return err
	}

	c.client = *client

	return nil
}
