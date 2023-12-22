// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#asn):
 * > ASN is short for Autonomous System Number. This identifier is used in the BGP protocol to identify which "autonomous system" a particular prefix is originating and transiting through.
 * > 
 * > The AS number model within NetBox allows you to model some of this real-world relationship.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzkalabsopensource/pulumi-netbox";
 *
 * const testRir = new netbox.ipam.Rir("testRir", {});
 * const testAsn = new netbox.ipam.Asn("testAsn", {
 *     asn: 1337,
 *     rirId: testRir.id,
 * });
 * ```
 */
export class Asn extends pulumi.CustomResource {
    /**
     * Get an existing Asn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AsnState, opts?: pulumi.CustomResourceOptions): Asn {
        return new Asn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:ipam/asn:Asn';

    /**
     * Returns true if the given object is an instance of Asn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Asn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Asn.__pulumiType;
    }

    public readonly asn!: pulumi.Output<number>;
    public readonly rirId!: pulumi.Output<number>;
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Asn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AsnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AsnArgs | AsnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AsnState | undefined;
            resourceInputs["asn"] = state ? state.asn : undefined;
            resourceInputs["rirId"] = state ? state.rirId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AsnArgs | undefined;
            if ((!args || args.asn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'asn'");
            }
            if ((!args || args.rirId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rirId'");
            }
            resourceInputs["asn"] = args ? args.asn : undefined;
            resourceInputs["rirId"] = args ? args.rirId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Asn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Asn resources.
 */
export interface AsnState {
    asn?: pulumi.Input<number>;
    rirId?: pulumi.Input<number>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Asn resource.
 */
export interface AsnArgs {
    asn: pulumi.Input<number>;
    rirId: pulumi.Input<number>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}