import * as pulumi from "@pulumi/pulumi";
export declare class Provider extends pulumi.ProviderResource {
    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Provider;
    /**
     * Your Gitpod access token
     */
    readonly accessToken: pulumi.Output<string | undefined>;
    readonly organizationId: pulumi.Output<string | undefined>;
    /**
     * Id of owner account
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
     * Your Gitpod access token
     */
    accessToken?: pulumi.Input<string>;
    organizationId?: pulumi.Input<string>;
    /**
     * Id of owner account
     */
    ownerId?: pulumi.Input<string>;
}
