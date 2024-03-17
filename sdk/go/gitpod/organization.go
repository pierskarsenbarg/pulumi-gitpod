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

// A Gitpod Organization
type Organization struct {
	pulumi.CustomResourceState

	// Name of the organization created
	Name pulumix.Output[string] `pulumi:"name"`
	// Id of the organization
	Org_id pulumix.Output[string] `pulumi:"org_id"`
	// Slug of the organization
	Slug pulumix.Output[string] `pulumi:"slug"`
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		args = &OrganizationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Organization
	err := ctx.RegisterResource("gitpod:index:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("gitpod:index:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type organizationState struct {
}

type OrganizationState struct {
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
	// Name of the organization to create
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// Name of the organization to create
	Name pulumix.Input[*string]
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationOutput struct{ *pulumi.OutputState }

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Organization)(nil)).Elem()
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOutput(ctx context.Context) pulumix.Output[Organization] {
	return pulumix.Output[Organization]{
		OutputState: o.OutputState,
	}
}

// Name of the organization created
func (o OrganizationOutput) Name() pulumix.Output[string] {
	value := pulumix.Apply[Organization](o, func(v Organization) pulumix.Output[string] { return v.Name })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

// Id of the organization
func (o OrganizationOutput) Org_id() pulumix.Output[string] {
	value := pulumix.Apply[Organization](o, func(v Organization) pulumix.Output[string] { return v.Org_id })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

// Slug of the organization
func (o OrganizationOutput) Slug() pulumix.Output[string] {
	value := pulumix.Apply[Organization](o, func(v Organization) pulumix.Output[string] { return v.Slug })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func init() {
	pulumi.RegisterOutputType(OrganizationOutput{})
}
