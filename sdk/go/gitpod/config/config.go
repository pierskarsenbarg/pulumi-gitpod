// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pierskarsenbarg/pulumi-gitpod/sdk/go/gitpod/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// Your Gitpod access token
func GetAccessToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "gitpod:accessToken")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault("", nil, "GITPOD_ACCESSTOKEN"); d != nil {
		value = d.(string)
	}
	return value
}
func GetOrganizationId(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "gitpod:organizationId")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault("", nil, "GITPOD_ORGANISATIONID"); d != nil {
		value = d.(string)
	}
	return value
}

// Id of owner account
func GetOwnerId(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "gitpod:ownerId")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault("", nil, "GITPOD_OWNERID"); d != nil {
		value = d.(string)
	}
	return value
}
