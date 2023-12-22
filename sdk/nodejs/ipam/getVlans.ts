// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getVlans(args?: GetVlansArgs, opts?: pulumi.InvokeOptions): Promise<GetVlansResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:ipam/getVlans:getVlans", {
        "filters": args.filters,
        "limit": args.limit,
    }, opts);
}

/**
 * A collection of arguments for invoking getVlans.
 */
export interface GetVlansArgs {
    filters?: inputs.ipam.GetVlansFilter[];
    /**
     * Defaults to `0`.
     */
    limit?: number;
}

/**
 * A collection of values returned by getVlans.
 */
export interface GetVlansResult {
    readonly filters?: outputs.ipam.GetVlansFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Defaults to `0`.
     */
    readonly limit?: number;
    readonly vlans: outputs.ipam.GetVlansVlan[];
}
export function getVlansOutput(args?: GetVlansOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVlansResult> {
    return pulumi.output(args).apply((a: any) => getVlans(a, opts))
}

/**
 * A collection of arguments for invoking getVlans.
 */
export interface GetVlansOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.ipam.GetVlansFilterArgs>[]>;
    /**
     * Defaults to `0`.
     */
    limit?: pulumi.Input<number>;
}