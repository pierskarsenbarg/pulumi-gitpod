package pkg

import (
	"context"
	"fmt"

	gitpodapi "github.com/pierskarsenbarg/pulumi-gitpod/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Workspace struct {
}

func (w *Workspace) Annotate(a infer.Annotator) {
	a.Describe(&w, "Gitpod workspace")
}

type WorkspaceArgs struct {
	Class          WorkspaceClass  `pulumi:"class,optional"`
	Editor         WorkspaceEditor `pulumi:"editor,optional"`
	OrganizationId string          `pulumi:"organizationId,optional"`
	OwnerId        string          `pulumi:"ownerId,optional"`
	ContextUrl     string          `pulumi:"contextUrl"`
}

func (wa *WorkspaceArgs) Annotate(a infer.Annotator) {
	a.Describe(&wa.Class, "Define the compute resources that you want your workspace to use. Defaults to `g1-standard`.")
	a.Describe(&wa.Editor, "Choose which editor you want to be able to use.")
	a.Describe(&wa.OrganizationId, "Set the Organization Id of the workspace. If not set then the Organization ID from the stack config will be used.")
	a.Describe(&wa.OwnerId, "Set the Owner Id of the workspace. If not set then the Owner ID from the stack config will be used.")
}

type WorkspaceEditor struct {
	Name    EditorName    `pulumi:"name,optional"`
	Version EditorVersion `pulumi:"version,optional"`
}

func (we *WorkspaceEditor) Annotate(a infer.Annotator) {
	a.Describe(&we.Name, "Name of the editor that you'd like to use in your Gitpod workspace. Defaults to VS Code desktop.")
	a.Describe(&we.Version, "Version of editor to use. Options are `latest` or `stable`. Defaults to `stable`")
	a.SetDefault(&we.Name, "code")
	a.SetDefault(&we.Version, "stable")
}

type WorkspaceState struct {
	Id       string                 `pulumi:"workspaceId"`
	Metadata WorkspaceStateMetaData `pulumi:"metadata"`
	GitInfo  WorkspaceStateGitInfo  `pulumi:"gitInfo"`
	Spec     WorkspaceStateSpec     `pulumi:"spec"`
}

func (ws *WorkspaceState) Annotate(a infer.Annotator) {
	a.Describe(&ws.Id, "The ID of the workspace")
	a.Describe(&ws.Metadata, "The metadata of the workspace")
	a.Describe(&ws.GitInfo, "The information from git used to build the workspace")
}

type WorkspaceStateMetaData struct {
	OwnerId            string `pulumi:"ownerId"`
	OrganizationId     string `pulumi:"organizationId"`
	Name               string `pulumi:"name"`
	OriginalContextUrl string `pulumi:"originalContextUrl"`
	InstanceId         string `pulumi:"instanceId"`
}

func (wsm *WorkspaceStateMetaData) Annotate(a infer.Annotator) {
	a.Describe(&wsm.OwnerId, "The ID of the user that owns the workspace")
	a.Describe(&wsm.OrganizationId, "The ID of the organization that the workspace belongs to")
	a.Describe(&wsm.Name, "The name of the workspace")
	a.Describe(&wsm.OriginalContextUrl, "The original context URL of the workspace")

}

type WorkspaceStateGitInfo struct {
	LatestCommit string `pulumi:"latestCommit"`
	CloneUrl     string `pulumi:"cloneUrl"`
	Branch       string `pulumi:"branch"`
}

func (wgi *WorkspaceStateGitInfo) Annotate(a infer.Annotator) {
	a.Describe(&wgi.Branch, "Git branch used to create the workspace")
	a.Describe(&wgi.CloneUrl, "The git clone url used to create the workspace")
	a.Describe(&wgi.LatestCommit, "The commit hash used to create the workspace")
}

type WorkspaceStateSpec struct {
	Editor WorkspaceStateSpecEditor `pulumi:"editor"`
	Class  string                   `pulumi:"class"`
}

func (ws *WorkspaceStateSpec) Annotate(a infer.Annotator) {
	a.Describe(&ws.Editor, "The IDE that has been chosen")
}

type WorkspaceStateSpecEditor struct {
	Name    string `pulumi:"name"`
	Version string `pulumi:"version"`
}

func (wse *WorkspaceStateSpecEditor) Annotate(a infer.Annotator) {
	a.Describe(&wse.Name, "Name of IDE that is being used in the workspace")
	a.Describe(&wse.Version, "Version of the IDE that is being used in the workspace")
}

type GetWorkspace struct{}

type GetWorkspaceArgs struct {
	WorkspaceId string `json:"workspaceId"`
}

func (w *Workspace) Create(ctx p.Context, name string, input WorkspaceArgs, preview bool) (
	id string, output WorkspaceState, err error) {
	config := infer.GetConfig[Config](ctx)
	if len(config.OrganizationId) == 0 && len(input.OrganizationId) == 0 {
		err := fmt.Errorf("you must set either the organization id in the stack config or the organization id in the resource")
		return "", WorkspaceState{}, err
	}
	if len(config.OwnerId) == 0 && len(input.OwnerId) == 0 {
		err := fmt.Errorf("you must set either the owner id in the stack config or the owner id in the resource")
		return "", WorkspaceState{}, err
	}
	if preview {
		return "", WorkspaceState{}, nil
	}
	res, err := w.createWorkspace(input, config)
	if err != nil {
		return "", WorkspaceState{}, err
	}

	return name, WorkspaceState{
		Id: res.Workspace.Id,
		Metadata: WorkspaceStateMetaData{
			OwnerId:            res.Workspace.MetaData.OwnerId,
			OrganizationId:     res.Workspace.MetaData.OrganizationId,
			OriginalContextUrl: res.Workspace.MetaData.OriginalContextUrl,
			Name:               res.Workspace.MetaData.Name,
			InstanceId:         res.Workspace.Status.InstanceId,
		},
		GitInfo: WorkspaceStateGitInfo{
			LatestCommit: res.Workspace.Status.GitStatus.LatestCommit,
			CloneUrl:     res.Workspace.Status.GitStatus.CloneUrl,
			Branch:       res.Workspace.Status.GitStatus.Branch,
		},
		Spec: WorkspaceStateSpec{
			Editor: WorkspaceStateSpecEditor{
				Name:    res.Workspace.Spec.Editor.Name,
				Version: res.Workspace.Spec.Editor.Version,
			},
			Class: res.Workspace.Spec.Class,
		},
	}, nil
}

func (w *Workspace) Delete(ctx p.Context, id string, props WorkspaceState) error {
	config := infer.GetConfig[Config](ctx)
	err := w.deleteWorkspace(props.Id, config)
	if err != nil {
		return err
	}
	return nil
}

func (*Workspace) Diff(ctx p.Context, id string, olds WorkspaceState, news WorkspaceArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if (olds.Metadata.OrganizationId != news.OrganizationId) && len(news.OrganizationId) > 0 {
		diff["organizationId"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if olds.Metadata.OwnerId != news.OwnerId && len(news.OwnerId) > 0 {
		diff["ownerId"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if news.Editor.Name != EditorName(olds.Spec.Editor.Name) {
		diff["name"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if news.Editor.Version != EditorVersion(olds.Spec.Editor.Version) {
		diff["version"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if news.Class != WorkspaceClass(olds.Spec.Class) {
		diff["class"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if (news.ContextUrl) != olds.Metadata.OriginalContextUrl {
		diff["contextUrl"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (w *Workspace) Read(ctx p.Context, id string, inputs WorkspaceArgs, state WorkspaceState) (
	string, WorkspaceArgs, WorkspaceState, error) {
	config := infer.GetConfig[Config](ctx)
	res, err := w.getWorkspace(id, config)
	if err != nil {
		return "", WorkspaceArgs{}, WorkspaceState{}, err
	}
	return res.Workspace.Id, inputs, WorkspaceState{
		Id: res.Workspace.Id,
		Metadata: WorkspaceStateMetaData{
			OwnerId:            res.Workspace.MetaData.OwnerId,
			OrganizationId:     res.Workspace.MetaData.OrganizationId,
			Name:               res.Workspace.Id,
			OriginalContextUrl: res.Workspace.MetaData.OriginalContextUrl,
		},
		GitInfo: WorkspaceStateGitInfo{
			LatestCommit: res.Workspace.Status.GitStatus.LatestCommit,
			CloneUrl:     res.Workspace.Status.GitStatus.CloneUrl,
			Branch:       res.Workspace.Status.GitStatus.Branch,
		},
		Spec: WorkspaceStateSpec{
			Editor: WorkspaceStateSpecEditor{
				Name:    res.Workspace.Spec.Editor.Name,
				Version: res.Workspace.Spec.Editor.Version,
			},
			Class: res.Workspace.Spec.Class,
		},
	}, nil
}

func (GetWorkspace) Call(ctx p.Context, args GetWorkspaceArgs) (WorkspaceState, error) {
	config := infer.GetConfig[Config](ctx)
	req := gitpodapi.WorkspaceRequest{
		WorkspaceId: args.WorkspaceId,
	}
	ws, err := config.Client.GetWorkspace(ctx, req)
	orgResponse := WorkspaceState{
		Id: args.WorkspaceId,
		Metadata: WorkspaceStateMetaData{
			OwnerId: ws.Workspace.MetaData.OwnerId,
		},
	}
	if err != nil {
		return WorkspaceState{}, err
	}
	return orgResponse, nil
}

func (w *Workspace) createWorkspace(args WorkspaceArgs, config Config) (*gitpodapi.WorkspaceResponse, error) {
	ctx := context.Background()
	if len(args.OrganizationId) == 0 {
		args.OrganizationId = config.OrganizationId
	}
	if len(args.OwnerId) == 0 {
		args.OwnerId = config.OwnerId
	}
	req := gitpodapi.CreateWorkspaceRequest{
		Context: gitpodapi.WorkspaceRequestContext{
			Editor: gitpodapi.WorkspaceRequestContextUrlEditor{
				Name:    string(args.Editor.Name),
				Version: string(args.Editor.Version),
			},
			ContextUrl:     args.ContextUrl,
			WorkspaceClass: string(args.Class),
		},
		MetaData: gitpodapi.WorkspaceRequestMetadata{
			OwnerId:        config.OwnerId,
			OrganizationId: args.OrganizationId,
		},
	}
	workspace, err := config.Client.CreateWorkspace(ctx, req)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (w *Workspace) deleteWorkspace(id string, config Config) error {
	ctx := context.Background()
	req := gitpodapi.WorkspaceRequest{
		WorkspaceId: id,
	}
	err := config.Client.DeleteWorkspace(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (w *Workspace) getWorkspace(id string, config Config) (*gitpodapi.WorkspaceResponse, error) {
	ctx := context.Background()
	req := gitpodapi.WorkspaceRequest{
		WorkspaceId: id,
	}
	res, err := config.Client.GetWorkspace(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
