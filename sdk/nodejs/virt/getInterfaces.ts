// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInterfaces(args?: GetInterfacesArgs, opts?: pulumi.InvokeOptions): Promise<GetInterfacesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:virt/getInterfaces:getInterfaces", {
        "filters": args.filters,
        "nameRegex": args.nameRegex,
    }, opts);
}

/**
 * A collection of arguments for invoking getInterfaces.
 */
export interface GetInterfacesArgs {
    filters?: inputs.virt.GetInterfacesFilter[];
    nameRegex?: string;
}

/**
 * A collection of values returned by getInterfaces.
 */
export interface GetInterfacesResult {
    readonly filters?: outputs.virt.GetInterfacesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly interfaces: outputs.virt.GetInterfacesInterface[];
    readonly nameRegex?: string;
}
export function getInterfacesOutput(args?: GetInterfacesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInterfacesResult> {
    return pulumi.output(args).apply((a: any) => getInterfaces(a, opts))
}

/**
 * A collection of arguments for invoking getInterfaces.
 */
export interface GetInterfacesOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.virt.GetInterfacesFilterArgs>[]>;
    nameRegex?: pulumi.Input<string>;
}
