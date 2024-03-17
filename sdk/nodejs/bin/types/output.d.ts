import * as outputs from "../types/output";
export interface WorkspaceStateGitInfo {
    /**
     * Git branch used to create the workspace
     */
    branch: string;
    /**
     * The git clone url used to create the workspace
     */
    cloneUrl: string;
    /**
     * The commit hash used to create the workspace
     */
    latestCommit: string;
}
export interface WorkspaceStateMetaData {
    instanceId: string;
    /**
     * The name of the workspace
     */
    name: string;
    /**
     * The ID of the organization that the workspace belongs to
     */
    organizationId: string;
    /**
     * The original context URL of the workspace
     */
    originalContextUrl: string;
    /**
     * The ID of the user that owns the workspace
     */
    ownerId: string;
}
export interface WorkspaceStateSpec {
    class: string;
    /**
     * The IDE that has been chosen
     */
    editor: outputs.WorkspaceStateSpecEditor;
}
export interface WorkspaceStateSpecEditor {
    /**
     * Name of IDE that is being used in the workspace
     */
    name: string;
    /**
     * Version of the IDE that is being used in the workspace
     */
    version: string;
}
