// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitpod

import (
	"context"
	"reflect"

	"github.com/pierskarsenbarg/pulumi-gitpod/sdk/go/gitpod/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type Provider struct {
	pulumi.ProviderResourceState

	// Your Gitpod access token. You can create these in your user settings: https://gitpod.io/user/tokens
	AccessToken pulumix.Output[*string] `pulumi:"accessToken"`
	// Id of the organisation who is going to own the workspaces. You can find this on the organisation settings page of the currently selected organisation: https://gitpod.io/settings
	OrganizationId pulumix.Output[*string] `pulumi:"organizationId"`
	// Id of owner account. This can be found on your user account page: https://gitpod.io/user/account
	OwnerId pulumix.Output[*string] `pulumi:"ownerId"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.AccessToken == nil {
		if d := internal.GetEnvOrDefault("", nil, "GITPOD_ACCESSTOKEN"); d != nil {
			args.AccessToken = pulumix.Ptr(d.(string))
		}
	}
	if args.OrganizationId == nil {
		if d := internal.GetEnvOrDefault("", nil, "GITPOD_ORGANISATIONID"); d != nil {
			args.OrganizationId = pulumix.Ptr(d.(string))
		}
	}
	if args.OwnerId == nil {
		if d := internal.GetEnvOrDefault("", nil, "GITPOD_OWNERID"); d != nil {
			args.OwnerId = pulumix.Ptr(d.(string))
		}
	}
	if args.AccessToken != nil {
		untypedSecretValue := pulumi.ToSecret(args.AccessToken.ToOutput(ctx.Context()).Untyped())
		args.AccessToken = pulumix.MustConvertTyped[*string](untypedSecretValue)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:gitpod", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Your Gitpod access token. You can create these in your user settings: https://gitpod.io/user/tokens
	AccessToken *string `pulumi:"accessToken"`
	// Id of the organisation who is going to own the workspaces. You can find this on the organisation settings page of the currently selected organisation: https://gitpod.io/settings
	OrganizationId *string `pulumi:"organizationId"`
	// Id of owner account. This can be found on your user account page: https://gitpod.io/user/account
	OwnerId *string `pulumi:"ownerId"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Your Gitpod access token. You can create these in your user settings: https://gitpod.io/user/tokens
	AccessToken pulumix.Input[*string]
	// Id of the organisation who is going to own the workspaces. You can find this on the organisation settings page of the currently selected organisation: https://gitpod.io/settings
	OrganizationId pulumix.Input[*string]
	// Id of owner account. This can be found on your user account page: https://gitpod.io/user/account
	OwnerId pulumix.Input[*string]
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToOutput(ctx context.Context) pulumix.Output[Provider] {
	return pulumix.Output[Provider]{
		OutputState: o.OutputState,
	}
}

// Your Gitpod access token. You can create these in your user settings: https://gitpod.io/user/tokens
func (o ProviderOutput) AccessToken() pulumix.Output[*string] {
	value := pulumix.Apply[Provider](o, func(v Provider) pulumix.Output[*string] { return v.AccessToken })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// Id of the organisation who is going to own the workspaces. You can find this on the organisation settings page of the currently selected organisation: https://gitpod.io/settings
func (o ProviderOutput) OrganizationId() pulumix.Output[*string] {
	value := pulumix.Apply[Provider](o, func(v Provider) pulumix.Output[*string] { return v.OrganizationId })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

// Id of owner account. This can be found on your user account page: https://gitpod.io/user/account
func (o ProviderOutput) OwnerId() pulumix.Output[*string] {
	value := pulumix.Apply[Provider](o, func(v Provider) pulumix.Output[*string] { return v.OwnerId })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
}
