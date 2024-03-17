import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
/**
 * Gitpod workspace
 */
export declare class Workspace extends pulumi.CustomResource {
    /**
     * Get an existing Workspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workspace;
    /**
     * Returns true if the given object is an instance of Workspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Workspace;
    /**
     * The information from git used to build the workspace
     */
    readonly gitInfo: pulumi.Output<outputs.WorkspaceStateGitInfo>;
    /**
     * The metadata of the workspace
     */
    readonly metadata: pulumi.Output<outputs.WorkspaceStateMetaData>;
    readonly spec: pulumi.Output<outputs.WorkspaceStateSpec>;
    /**
     * The ID of the workspace
     */
    readonly workspaceId: pulumi.Output<string>;
    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * Define the compute resources that you want your workspace to use. Defaults to `g1-standard`.
     */
    class?: pulumi.Input<enums.WorkspaceClass>;
    contextUrl: pulumi.Input<string>;
    /**
     * Choose which editor you want to be able to use.
     */
    editor?: pulumi.Input<inputs.WorkspaceEditorArgs>;
    /**
     * Set the Organization Id of the workspace. If not set then the Organization ID from the stack config will be used.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Set the Owner Id of the workspace. If not set then the Owner ID from the stack config will be used.
     */
    ownerId?: pulumi.Input<string>;
}
