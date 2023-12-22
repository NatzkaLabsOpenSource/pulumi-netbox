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
 * const example1 = netbox.ipam.getVlanGroup({
 *     name: "example1",
 * });
 * const example2 = netbox.ipam.getVlanGroup({
 *     slug: "example2",
 * });
 * const example3 = netbox.ipam.getVlanGroup({
 *     name: "example",
 *     scopeType: "dcim.site",
 *     scopeId: netbox_site.example.id,
 * });
 * ```
 */
export function getVlanGroup(args?: GetVlanGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetVlanGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:ipam/getVlanGroup:getVlanGroup", {
        "name": args.name,
        "scopeId": args.scopeId,
        "scopeType": args.scopeType,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getVlanGroup.
 */
export interface GetVlanGroupArgs {
    /**
     * At least one of `name` or `slug` must be given.
     */
    name?: string;
    /**
     * Required when `scopeType` is set.
     */
    scopeId?: number;
    /**
     * Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
     */
    scopeType?: string;
    /**
     * At least one of `name` or `slug` must be given.
     */
    slug?: string;
}

/**
 * A collection of values returned by getVlanGroup.
 */
export interface GetVlanGroupResult {
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly maxVid: number;
    readonly minVid: number;
    /**
     * At least one of `name` or `slug` must be given.
     */
    readonly name: string;
    /**
     * Required when `scopeType` is set.
     */
    readonly scopeId?: number;
    /**
     * Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
     */
    readonly scopeType?: string;
    /**
     * At least one of `name` or `slug` must be given.
     */
    readonly slug: string;
    readonly vlanCount: number;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@pulumi/netbox";
 *
 * const example1 = netbox.ipam.getVlanGroup({
 *     name: "example1",
 * });
 * const example2 = netbox.ipam.getVlanGroup({
 *     slug: "example2",
 * });
 * const example3 = netbox.ipam.getVlanGroup({
 *     name: "example",
 *     scopeType: "dcim.site",
 *     scopeId: netbox_site.example.id,
 * });
 * ```
 */
export function getVlanGroupOutput(args?: GetVlanGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVlanGroupResult> {
    return pulumi.output(args).apply((a: any) => getVlanGroup(a, opts))
}

/**
 * A collection of arguments for invoking getVlanGroup.
 */
export interface GetVlanGroupOutputArgs {
    /**
     * At least one of `name` or `slug` must be given.
     */
    name?: pulumi.Input<string>;
    /**
     * Required when `scopeType` is set.
     */
    scopeId?: pulumi.Input<number>;
    /**
     * Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
     */
    scopeType?: pulumi.Input<string>;
    /**
     * At least one of `name` or `slug` must be given.
     */
    slug?: pulumi.Input<string>;
}