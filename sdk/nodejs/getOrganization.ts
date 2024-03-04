// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * GetOrganization gets the Organization information
 */
export function getOrganization(args?: GetOrganizationArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gitpod:index:getOrganization", {
    }, opts);
}

/**
 * Id of the organization
 */
export interface GetOrganizationArgs {
}

export interface GetOrganizationResult {
    /**
     * Name of the organization created
     */
    readonly name: string;
    /**
     * Id of the organization
     */
    readonly org_id: string;
    /**
     * Slug of the organization
     */
    readonly slug: string;
}
/**
 * GetOrganization gets the Organization information
 */
export function getOrganizationOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationResult> {
    return pulumi.output(getOrganization(opts))
}
