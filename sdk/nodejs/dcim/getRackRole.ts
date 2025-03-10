// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getRackRole(args: GetRackRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRackRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:dcim/getRackRole:getRackRole", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRackRole.
 */
export interface GetRackRoleArgs {
    name: string;
}

/**
 * A collection of values returned by getRackRole.
 */
export interface GetRackRoleResult {
    readonly colorHex: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly slug: string;
    readonly tags: string[];
}
export function getRackRoleOutput(args: GetRackRoleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRackRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("netbox:dcim/getRackRole:getRackRole", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRackRole.
 */
export interface GetRackRoleOutputArgs {
    name: pulumi.Input<string>;
}
