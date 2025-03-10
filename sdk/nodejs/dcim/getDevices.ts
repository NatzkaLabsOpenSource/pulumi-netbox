// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDevices(args?: GetDevicesArgs, opts?: pulumi.InvokeOptions): Promise<GetDevicesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:dcim/getDevices:getDevices", {
        "filters": args.filters,
        "limit": args.limit,
        "nameRegex": args.nameRegex,
    }, opts);
}

/**
 * A collection of arguments for invoking getDevices.
 */
export interface GetDevicesArgs {
    filters?: inputs.dcim.GetDevicesFilter[];
    limit?: number;
    nameRegex?: string;
}

/**
 * A collection of values returned by getDevices.
 */
export interface GetDevicesResult {
    readonly devices: outputs.dcim.GetDevicesDevice[];
    readonly filters?: outputs.dcim.GetDevicesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly limit?: number;
    readonly nameRegex?: string;
}
export function getDevicesOutput(args?: GetDevicesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDevicesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("netbox:dcim/getDevices:getDevices", {
        "filters": args.filters,
        "limit": args.limit,
        "nameRegex": args.nameRegex,
    }, opts);
}

/**
 * A collection of arguments for invoking getDevices.
 */
export interface GetDevicesOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.dcim.GetDevicesFilterArgs>[]>;
    limit?: pulumi.Input<number>;
    nameRegex?: pulumi.Input<string>;
}
