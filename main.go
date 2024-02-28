package main

import (
	"fmt"
	"os"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

func main() {
	err := p.RunProvider("file", "0.1.0", provider())
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
		Config: infer.Config[Config](),
	})
}

type Config struct {
	AccessToken string `pulumi:"accessToken"`
}
