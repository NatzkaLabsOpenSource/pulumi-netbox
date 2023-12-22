// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@pulumi/netbox";
 *
 * const getByName = netbox.dcim.getSiteGroup({
 *     name: "example-sitegroup-1",
 * });
 * const getBySlug = netbox.dcim.getSiteGroup({
 *     slug: "sitegrp",
 * });
 * ```
 */
export function getSiteGroup(args?: GetSiteGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSiteGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:dcim/getSiteGroup:getSiteGroup", {
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getSiteGroup.
 */
export interface GetSiteGroupArgs {
    /**
     * At least one of `name` or `slug` must be given.
     */
    name?: string;
    /**
     * At least one of `name` or `slug` must be given.
     */
    slug?: string;
}

/**
 * A collection of values returned by getSiteGroup.
 */
export interface GetSiteGroupResult {
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * At least one of `name` or `slug` must be given.
     */
    readonly name: string;
    /**
     * At least one of `name` or `slug` must be given.
     */
    readonly slug: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@pulumi/netbox";
 *
 * const getByName = netbox.dcim.getSiteGroup({
 *     name: "example-sitegroup-1",
 * });
 * const getBySlug = netbox.dcim.getSiteGroup({
 *     slug: "sitegrp",
 * });
 * ```
 */
export function getSiteGroupOutput(args?: GetSiteGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSiteGroupResult> {
    return pulumi.output(args).apply((a: any) => getSiteGroup(a, opts))
}

/**
 * A collection of arguments for invoking getSiteGroup.
 */
export interface GetSiteGroupOutputArgs {
    /**
     * At least one of `name` or `slug` must be given.
     */
    name?: pulumi.Input<string>;
    /**
     * At least one of `name` or `slug` must be given.
     */
    slug?: pulumi.Input<string>;
}