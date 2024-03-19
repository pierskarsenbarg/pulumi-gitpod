import * as pulumi from "@pulumi/pulumi";
import * as outputs from "./types/output";
export declare function getWorkspace(args?: GetWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceResult>;
export interface GetWorkspaceArgs {
}
export interface GetWorkspaceResult {
    /**
     * The information from git used to build the workspace
     */
    readonly gitInfo: outputs.WorkspaceStateGitInfo;
    /**
     * The metadata of the workspace
     */
    readonly metadata: outputs.WorkspaceStateMetaData;
    readonly spec: outputs.WorkspaceStateSpec;
    /**
     * The ID of the workspace
     */
    readonly workspaceId: string;
}
export declare function getWorkspaceOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkspaceResult>;
