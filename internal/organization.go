package gitpodapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type OrganizationRequest struct {
	Name string `json:"name"`
}

type OrganizationIdRequest struct {
	Id string `json:"organizationId"`
}

type OrganizationNameRequest struct {
	Id   string `json:"organizationId"`
	Name string `json:"name"`
}

type OrganizationResponse struct {
	Organization Organization `json:"organization"`
}

type Organization struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	CreationTime string `json:"creationTime"`
	Slug         string `json:"slug"`
}

func (c *Client) CreateOrganization(ctx context.Context, orgRequest OrganizationRequest) (*OrganizationResponse, error) {
	if len(orgRequest.Name) == 0 {
		return nil, errors.New("organization name must not be empty")
	}

	requestPath := "gitpod.v1.OrganizationService/CreateOrganization"
	var orgResponse OrganizationResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, orgRequest, &orgResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to create organization: %w", err)
	}
	return &orgResponse, nil
}

func (c *Client) DeleteOrganization(ctx context.Context, orgId OrganizationIdRequest) error {
	requestPath := "gitpod.v1.OrganizationService/DeleteOrganization"
	var orgResponse OrganizationResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, orgId, &orgResponse)
	if err != nil {
		return fmt.Errorf("couldn't get organisation: %w", err)
	}
	return nil
}

func (c *Client) GetOrganization(ctx context.Context, orgId OrganizationIdRequest) (*OrganizationResponse, error) {
	requestPath := "gitpod.v1.OrganizationService/GetOrganization"
	var orgResponse OrganizationResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, orgId, &orgResponse)
	if err != nil {
		return nil, fmt.Errorf("couldn't get organisation: %w", err)
	}
	return &orgResponse, nil
}

func (c *Client) UpdateOrganization(ctx context.Context, orgNameRequest OrganizationNameRequest) error {
	requestPath := "gitpod.v1.OrganizationService/UpdateOrganization"
	var orgRespone OrganizationResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, orgNameRequest, &orgRespone)
	if err != nil {
		return fmt.Errorf("couldn't update organization: %w", err)
	}
	return nil
}
