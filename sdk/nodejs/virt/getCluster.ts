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
 * const vmwCluster01 = netbox.virt.getCluster({
 *     name: "vmw-cluster-01",
 * });
 * ```
 */
export function getCluster(args?: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:virt/getCluster:getCluster", {
        "clusterGroupId": args.clusterGroupId,
        "id": args.id,
        "name": args.name,
        "siteId": args.siteId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    clusterGroupId?: number;
    /**
     * At least one of `name`, `siteId` or `id` must be given.
     */
    id?: string;
    /**
     * At least one of `name`, `siteId` or `id` must be given.
     */
    name?: string;
    /**
     * At least one of `name`, `siteId` or `id` must be given.
     */
    siteId?: number;
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    readonly clusterGroupId: number;
    readonly clusterId: number;
    readonly clusterTypeId: number;
    readonly comments: string;
    readonly customFields: {[key: string]: string};
    readonly description: string;
    /**
     * At least one of `name`, `siteId` or `id` must be given.
     */
    readonly id: string;
    /**
     * At least one of `name`, `siteId` or `id` must be given.
     */
    readonly name?: string;
    /**
     * At least one of `name`, `siteId` or `id` must be given.
     */
    readonly siteId: number;
    readonly tags: string[];
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@pulumi/netbox";
 *
 * const vmwCluster01 = netbox.virt.getCluster({
 *     name: "vmw-cluster-01",
 * });
 * ```
 */
export function getClusterOutput(args?: GetClusterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetClusterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("netbox:virt/getCluster:getCluster", {
        "clusterGroupId": args.clusterGroupId,
        "id": args.id,
        "name": args.name,
        "siteId": args.siteId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterOutputArgs {
    clusterGroupId?: pulumi.Input<number>;
    /**
     * At least one of `name`, `siteId` or `id` must be given.
     */
    id?: pulumi.Input<string>;
    /**
     * At least one of `name`, `siteId` or `id` must be given.
     */
    name?: pulumi.Input<string>;
    /**
     * At least one of `name`, `siteId` or `id` must be given.
     */
    siteId?: pulumi.Input<number>;
}
