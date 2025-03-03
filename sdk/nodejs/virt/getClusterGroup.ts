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
 * const dcWest = netbox.virt.getClusterGroup({
 *     name: "dc-west",
 * });
 * ```
 */
export function getClusterGroup(args: GetClusterGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:virt/getClusterGroup:getClusterGroup", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterGroup.
 */
export interface GetClusterGroupArgs {
    name: string;
}

/**
 * A collection of values returned by getClusterGroup.
 */
export interface GetClusterGroupResult {
    readonly clusterGroupId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@pulumi/netbox";
 *
 * const dcWest = netbox.virt.getClusterGroup({
 *     name: "dc-west",
 * });
 * ```
 */
export function getClusterGroupOutput(args: GetClusterGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetClusterGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("netbox:virt/getClusterGroup:getClusterGroup", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterGroup.
 */
export interface GetClusterGroupOutputArgs {
    name: pulumi.Input<string>;
}
