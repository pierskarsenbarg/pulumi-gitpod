// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'gitpod';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * Your Gitpod access token. You can create these in your user settings: https://gitpod.io/user/tokens
     */
    public readonly accessToken!: pulumi.Output<string | undefined>;
    /**
     * Id of the organisation who is going to own the workspaces. You can find this on the organisation settings page of the currently selected organisation: https://gitpod.io/settings
     */
    public readonly organizationId!: pulumi.Output<string | undefined>;
    /**
     * Id of owner account. This can be found on your user account page: https://gitpod.io/user/account
     */
    public readonly ownerId!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["accessToken"] = (args?.accessToken ? pulumi.secret(args.accessToken) : undefined) ?? (utilities.getEnv("GITPOD_ACCESSTOKEN") || "");
            resourceInputs["organizationId"] = (args ? args.organizationId : undefined) ?? (utilities.getEnv("GITPOD_ORGANISATIONID") || "");
            resourceInputs["ownerId"] = (args ? args.ownerId : undefined) ?? (utilities.getEnv("GITPOD_OWNERID") || "");
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
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
