// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@pulumi/netbox";
 *
 * const dmz = netbox.getTag({
 *     name: "DMZ",
 * });
 * ```
 */
export function getTag(args: GetTagArgs, opts?: pulumi.InvokeOptions): Promise<GetTagResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:index/getTag:getTag", {
        "description": args.description,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTag.
 */
export interface GetTagArgs {
    description?: string;
    name: string;
}

/**
 * A collection of values returned by getTag.
 */
export interface GetTagResult {
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly slug: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@pulumi/netbox";
 *
 * const dmz = netbox.getTag({
 *     name: "DMZ",
 * });
 * ```
 */
export function getTagOutput(args: GetTagOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTagResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("netbox:index/getTag:getTag", {
        "description": args.description,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTag.
 */
export interface GetTagOutputArgs {
    description?: pulumi.Input<string>;
    name: pulumi.Input<string>;
}
