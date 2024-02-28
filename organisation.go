package main

import "github.com/pulumi/pulumi-go-provider/infer"

type Organization struct{}

type OrganizationArgs struct {
	Name string `pulumi:name`
}

type OrganizationState struct {
	Id   string `pulumi:id`
	Name string `pulumi:name`
	Slug string `pulumi:slug`
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
