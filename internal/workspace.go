package gitpodapi

import (
	"context"
	"fmt"
	"net/http"
)

type CreateWorkspaceRequest struct {
	MetaData WorkspaceRequestMetadata `json:"metadata"`
	Context  WorkspaceRequestContext  `json:"contextUrl"`
}

type WorkspaceRequestContext struct {
	Editor         WorkspaceRequestContextUrlEditor `json:"editor"`
	ContextUrl     string                           `json:"url"`
	WorkspaceClass string                           `json:"workspaceClass"`
}

type WorkspaceRequestContextUrlEditor struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type WorkspaceRequestMetadata struct {
	OwnerId        string `json:"ownerId"`
	OrganizationId string `json:"organizationId"`
}

type WorkspaceResponse struct {
	Workspace WorkspaceResponseWorkspace `json:"workspace"`
}

type WorkspaceResponseWorkspace struct {
	Id       string                    `json:"id"`
	MetaData WorkspaceResponseMetadata `json:"metadata"`
	Status   WorkspaceResponseStatus   `json:"status"`
	Spec     WorkspaceResponseSpec     `json:"spec"`
}

type WorkspaceResponseMetadata struct {
	OwnerId            string `json:"ownerId"`
	OrganizationId     string `json:"organizationId"`
	Name               string `json:"name"`
	OriginalContextUrl string `json:"originalContextUrl"`
}

type WorkspaceResponseSpec struct {
	Editor WorkspaceResponseSpecEditor `json:"editor"`
	Class  string                      `json:"class"`
}

type WorkspaceResponseSpecEditor struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type WorkspaceResponseStatus struct {
	InstanceId string                           `json:"instanceId"`
	GitStatus  WorkspaceResponseStatusGitStatus `json:"gitStatus"`
}

type WorkspaceResponseStatusGitStatus struct {
	CloneUrl     string `json:"cloneUrl"`
	Branch       string `json:"branch"`
	LatestCommit string `json:"latestCommit"`
}

type WorkspaceRequest struct {
	WorkspaceId string `json:"workspaceId"`
}

type DeleteWorkspaceResponse struct{}

func (c *Client) CreateWorkspace(ctx context.Context, workspace CreateWorkspaceRequest) (*WorkspaceResponse, error) {
	requestPath := "gitpod.v1.WorkspaceService/CreateAndStartWorkspace"
	var workspaceResponse WorkspaceResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, workspace, &workspaceResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to create workspace: %w", err)
	}
	return &workspaceResponse, nil
}

func (c *Client) DeleteWorkspace(ctx context.Context, workspace WorkspaceRequest) error {
	requestPath := "gitpod.v1.WorkspaceService/DeleteWorkspace"
	var res DeleteWorkspaceResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, workspace, &res)
	if err != nil {
		return fmt.Errorf("failed to delete workspace: %w", err)
	}
	return nil
}

func (c *Client) GetWorkspace(ctx context.Context, workspace WorkspaceRequest) (*WorkspaceResponse, error) {
	requestPath := "gitpod.v1.WorkspaceService/GetWorkspace"
	var res WorkspaceResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, workspace, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to get workspace: %w", err)
	}
	return &res, nil
}
