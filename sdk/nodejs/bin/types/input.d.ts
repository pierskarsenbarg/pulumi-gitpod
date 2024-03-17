import * as pulumi from "@pulumi/pulumi";
import * as enums from "../types/enums";
export interface WorkspaceEditorArgs {
    /**
     * Name of the editor that you'd like to use in your Gitpod workspace. Defaults to VS Code desktop.
     */
    name?: pulumi.Input<enums.EditorName>;
    /**
     * Version of editor to use. Options are `latest` or `stable`. Defaults to `stable`
     */
    version?: pulumi.Input<enums.EditorVersion>;
}
/**
 * workspaceEditorArgsProvideDefaults sets the appropriate defaults for WorkspaceEditorArgs
 */
export declare function workspaceEditorArgsProvideDefaults(val: WorkspaceEditorArgs): WorkspaceEditorArgs;
