// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getPrefixes(args?: GetPrefixesArgs, opts?: pulumi.InvokeOptions): Promise<GetPrefixesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:ipam/getPrefixes:getPrefixes", {
        "filters": args.filters,
        "limit": args.limit,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrefixes.
 */
export interface GetPrefixesArgs {
    /**
     * A list of filters to apply to the API query when requesting prefixes.
     */
    filters?: inputs.ipam.GetPrefixesFilter[];
    /**
     * The limit of objects to return from the API lookup. Defaults to `0`.
     */
    limit?: number;
}

/**
 * A collection of values returned by getPrefixes.
 */
export interface GetPrefixesResult {
    /**
     * A list of filters to apply to the API query when requesting prefixes.
     */
    readonly filters?: outputs.ipam.GetPrefixesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The limit of objects to return from the API lookup. Defaults to `0`.
     */
    readonly limit?: number;
    readonly prefixes: outputs.ipam.GetPrefixesPrefix[];
}
export function getPrefixesOutput(args?: GetPrefixesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrefixesResult> {
    return pulumi.output(args).apply((a: any) => getPrefixes(a, opts))
}

/**
 * A collection of arguments for invoking getPrefixes.
 */
export interface GetPrefixesOutputArgs {
    /**
     * A list of filters to apply to the API query when requesting prefixes.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.ipam.GetPrefixesFilterArgs>[]>;
    /**
     * The limit of objects to return from the API lookup. Defaults to `0`.
     */
    limit?: pulumi.Input<number>;
}