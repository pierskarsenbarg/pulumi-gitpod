import * as pulumi from "@pulumi/pulumi";
export declare class Provider extends pulumi.ProviderResource {
    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Provider;
    /**
     * Your Gitpod access token. You can create these in your user settings: https://gitpod.io/user/tokens
     */
    readonly accessToken: pulumi.Output<string | undefined>;
    /**
     * Id of the organisation who is going to own the workspaces. You can find this on the organisation settings page of the currently selected organisation: https://gitpod.io/settings
     */
    readonly organizationId: pulumi.Output<string | undefined>;
    /**
     * Id of owner account. This can be found on your user account page: https://gitpod.io/user/account
     */
    readonly ownerId: pulumi.Output<string | undefined>;
    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions);
}
/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * Your Gitpod access token. You can create these in your user settings: https://gitpod.io/user/tokens
     */
    accessToken?: pulumi.Input<string>;
    /**
     * Id of the organisation who is going to own the workspaces. You can find this on the organisation settings page of the currently selected organisation: https://gitpod.io/settings
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Id of owner account. This can be found on your user account page: https://gitpod.io/user/account
     */
    ownerId?: pulumi.Input<string>;
}
