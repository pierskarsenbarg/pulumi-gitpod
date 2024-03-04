import * as pulumi from "@pulumi/pulumi";
/**
 * GetOrganization gets the Organization information
 */
export declare function getOrganization(args?: GetOrganizationArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult>;
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
export declare function getOrganizationOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationResult>;
