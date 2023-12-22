// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#providers):
 *
 * > A circuit provider is any entity which provides some form of connectivity of among sites or organizations within a site. While this obviously includes carriers which offer Internet and private transit service, it might also include Internet exchange (IX) points and even organizations with whom you peer directly. Each circuit within NetBox must be assigned a provider and a circuit ID which is unique to that provider.
 * > 
 * > Each provider may be assigned an autonomous system number (ASN), an account number, and contact information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzkalabsopensource/pulumi-netbox";
 *
 * const test = new netbox.circuit.CircuitProvider("test", {});
 * ```
 */
export class CircuitProvider extends pulumi.CustomResource {
    /**
     * Get an existing CircuitProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CircuitProviderState, opts?: pulumi.CustomResourceOptions): CircuitProvider {
        return new CircuitProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:circuit/circuitProvider:CircuitProvider';

    /**
     * Returns true if the given object is an instance of CircuitProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CircuitProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CircuitProvider.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;
    public readonly slug!: pulumi.Output<string>;

    /**
     * Create a CircuitProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CircuitProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CircuitProviderArgs | CircuitProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CircuitProviderState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as CircuitProviderArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CircuitProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CircuitProvider resources.
 */
export interface CircuitProviderState {
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CircuitProvider resource.
 */
export interface CircuitProviderArgs {
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
}
