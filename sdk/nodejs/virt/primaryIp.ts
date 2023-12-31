// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource is used to define the primary IP for a given virtual machine. The primary IP is reflected in the Virtual machine Netbox UI, which identifies the Primary IPv4 and IPv6 addresses.
 */
export class PrimaryIp extends pulumi.CustomResource {
    /**
     * Get an existing PrimaryIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrimaryIpState, opts?: pulumi.CustomResourceOptions): PrimaryIp {
        return new PrimaryIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:virt/primaryIp:PrimaryIp';

    /**
     * Returns true if the given object is an instance of PrimaryIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrimaryIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrimaryIp.__pulumiType;
    }

    public readonly ipAddressId!: pulumi.Output<number>;
    /**
     * Defaults to `4`.
     */
    public readonly ipAddressVersion!: pulumi.Output<number | undefined>;
    public readonly virtualMachineId!: pulumi.Output<number>;

    /**
     * Create a PrimaryIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrimaryIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrimaryIpArgs | PrimaryIpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrimaryIpState | undefined;
            resourceInputs["ipAddressId"] = state ? state.ipAddressId : undefined;
            resourceInputs["ipAddressVersion"] = state ? state.ipAddressVersion : undefined;
            resourceInputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as PrimaryIpArgs | undefined;
            if ((!args || args.ipAddressId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddressId'");
            }
            if ((!args || args.virtualMachineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualMachineId'");
            }
            resourceInputs["ipAddressId"] = args ? args.ipAddressId : undefined;
            resourceInputs["ipAddressVersion"] = args ? args.ipAddressVersion : undefined;
            resourceInputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrimaryIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrimaryIp resources.
 */
export interface PrimaryIpState {
    ipAddressId?: pulumi.Input<number>;
    /**
     * Defaults to `4`.
     */
    ipAddressVersion?: pulumi.Input<number>;
    virtualMachineId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PrimaryIp resource.
 */
export interface PrimaryIpArgs {
    ipAddressId: pulumi.Input<number>;
    /**
     * Defaults to `4`.
     */
    ipAddressVersion?: pulumi.Input<number>;
    virtualMachineId: pulumi.Input<number>;
}
