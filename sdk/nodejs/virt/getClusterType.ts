// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getClusterType(args: GetClusterTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterTypeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:virt/getClusterType:getClusterType", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterType.
 */
export interface GetClusterTypeArgs {
    name: string;
}

/**
 * A collection of values returned by getClusterType.
 */
export interface GetClusterTypeResult {
    readonly clusterTypeId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
export function getClusterTypeOutput(args: GetClusterTypeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterTypeResult> {
    return pulumi.output(args).apply((a: any) => getClusterType(a, opts))
}

/**
 * A collection of arguments for invoking getClusterType.
 */
export interface GetClusterTypeOutputArgs {
    name: pulumi.Input<string>;
}