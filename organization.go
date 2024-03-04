package main

import (
	"context"
	"fmt"

	gitpodapi "github.com/pierskarsenbarg/pulumi-gitpod/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type Organization struct {
}

type OrganizationArgs struct {
	Name string `pulumi:"name,optional"`
}

type OrganizationState struct {
	Id   string `pulumi:"org_id"`
	Name string `pulumi:"name"`
	Slug string `pulumi:"slug"`
}

type GitpodOrganization struct {
	Id string
}

type GetOrganization struct{}

type GetOrganizationArgs struct {
	OrganizationId string `json:"organizationId"`
}

func (o *Organization) Annotate(a infer.Annotator) {
	a.Describe(&o, "A Gitpod Organization")
}

func (o *OrganizationArgs) Annotate(a infer.Annotator) {
	a.Describe(&o.Name, "Name of the organization to create")
}

func (o *OrganizationState) Annotate(a infer.Annotator) {
	a.Describe(&o.Id, "Id of the organization")
	a.Describe(&o.Name, "Name of the organization created")
	a.Describe(&o.Slug, "Slug of the organization")
}

func (goa *GetOrganization) Annotate(a infer.Annotator) {
	a.Describe(goa, "GetOrganization gets the Organization information")
}

func (goa *GetOrganizationArgs) Annotate(a infer.Annotator) {
	a.Describe(&goa.OrganizationId, "Id of the organization")
}

func (o *Organization) Create(ctx p.Context, name string, input OrganizationArgs, preview bool) (
	id string, output OrganizationState, err error) {
	config := infer.GetConfig[Config](ctx)
	if preview {
		return "", OrganizationState{}, nil
	}
	orgName, _ := resource.NewUniqueHex(fmt.Sprintf("%s-", name), 8, 0)
	if len(input.Name) > 0 {
		orgName = input.Name
	}
	resp, err := o.createOrganization(orgName, config)
	if err != nil {
		return "", OrganizationState{}, err
	}
	state := OrganizationState{
		Id:   resp.Organization.Id,
		Name: resp.Organization.Name,
		Slug: resp.Organization.Slug,
	}
	return name, state, nil
}

func (o *Organization) Delete(ctx p.Context, id string, props OrganizationState) error {
	config := infer.GetConfig[Config](ctx)
	err := o.deleteOrganization(props.Id, config)
	if err != nil {
		return err
	}
	return nil
}

func (o *Organization) Read(ctx p.Context, id string, inputs OrganizationArgs, state OrganizationState) (
	string, OrganizationArgs, OrganizationState, error) {
	config := infer.GetConfig[Config](ctx)
	resp, err := o.getOrganization(state.Id, config)
	if err != nil {
		return "", OrganizationArgs{}, OrganizationState{}, err
	}
	return resp.Organization.Id, OrganizationArgs{
			Name: resp.Organization.Name,
		}, OrganizationState{
			Id:   resp.Organization.Id,
			Name: resp.Organization.Name,
			Slug: resp.Organization.Slug,
		}, nil
}

func (*Organization) Diff(ctx p.Context, id string, olds OrganizationState, news OrganizationArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if len(olds.Name) > 0 && len(news.Name) == 0 {
		// name removed
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	} else if len(news.Name) > 0 && olds.Name != news.Name {
		// name updated
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	} else {
		diff = map[string]p.PropertyDiff{} // reset because there are no changes
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (o *Organization) Update(ctx p.Context, id string, olds OrganizationState, news OrganizationArgs, preview bool) (OrganizationState, error) {
	var updateName string
	if !preview {
		var err error
		if news.Name != olds.Name && len(news.Name) == 0 {
			updateName, err = resource.NewUniqueHex(fmt.Sprintf("%s-", id), 8, 0)
			if err != nil {
				return OrganizationState{}, err
			}
		} else if news.Name != olds.Name {
			updateName = news.Name
		}

		config := infer.GetConfig[Config](ctx)
		err = o.updateOrganization(olds.Id, updateName, config)
		if err != nil {
			return OrganizationState{}, err
		}
	}
	return OrganizationState{
		Id:   olds.Id,
		Name: updateName,
		Slug: olds.Slug,
	}, nil
}

func (GetOrganization) Call(ctx p.Context, args GetOrganizationArgs) (OrganizationState, error) {
	config := infer.GetConfig[Config](ctx)
	req := gitpodapi.OrganizationIdRequest{
		Id: args.OrganizationId,
	}
	org, err := config.client.GetOrganization(ctx, req)
	orgResponse := OrganizationState{
		Id:   args.OrganizationId,
		Name: org.Organization.Name,
		Slug: org.Organization.Slug,
	}
	if err != nil {
		return OrganizationState{}, err
	}
	return orgResponse, nil
}

func (o *Organization) getOrganization(id string, config Config) (*gitpodapi.OrganizationResponse, error) {
	ctx := context.Background()
	req := gitpodapi.OrganizationIdRequest{
		Id: id,
	}
	org, err := config.client.GetOrganization(ctx, req)
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (o *Organization) createOrganization(name string, config Config) (*gitpodapi.OrganizationResponse, error) {
	ctx := context.Background()
	req := gitpodapi.OrganizationRequest{
		Name: name,
	}
	org, err := config.client.CreateOrganization(ctx, req)
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (*Organization) deleteOrganization(id string, config Config) error {
	ctx := context.Background()
	req := gitpodapi.OrganizationIdRequest{
		Id: id,
	}
	err := config.client.DeleteOrganization(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (*Organization) updateOrganization(id string, name string, config Config) error {
	ctx := context.Background()
	req := gitpodapi.OrganizationNameRequest{
		Id:   id,
		Name: name,
	}
	err := config.client.UpdateOrganization(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
