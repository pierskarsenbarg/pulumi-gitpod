package main

import (
	"fmt"
	"os"

	"github.com/pierskarsenbarg/pulumi-gitpod/pkg"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
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
		Metadata: schema.Metadata{
			DisplayName: "gitpod",
			Description: "Gitpod provider",
			LanguageMap: map[string]any{
				"go": gen.GoPackageInfo{
					Generics:       gen.GenericsSettingGenericsOnly,
					ImportBasePath: "github.com/pierskarsenbarg/pulumi-gitpod/sdk/go/gitpod",
				},
			},
		},
		Resources: []infer.InferredResource{
			infer.Resource[*pkg.Organization, pkg.OrganizationArgs, pkg.OrganizationState](),
			infer.Resource[*pkg.Workspace, pkg.WorkspaceArgs, pkg.WorkspaceState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"pkg": "index",
		},
		Functions: []infer.InferredFunction{
			infer.Function[*pkg.GetOrganization, pkg.GetOrganizationArgs, pkg.OrganizationState](),
		},
		Config: infer.Config[*pkg.Config](),
	})
}
