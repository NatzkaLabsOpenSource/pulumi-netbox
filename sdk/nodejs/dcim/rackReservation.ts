// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rackreservation/):
 *
 * > Users can reserve specific units within a rack for future use. An arbitrary set of units within a rack can be associated with a single reservation, but reservations cannot span multiple racks. A description is required for each reservation, reservations may optionally be associated with a specific tenant.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * const test = new netbox.dcim.Site("test", {
 *     name: "test",
 *     status: "active",
 * });
 * const testRack = new netbox.dcim.Rack("test", {
 *     name: "test",
 *     siteId: test.id,
 *     status: "active",
 *     width: 10,
 *     uHeight: 40,
 * });
 * const testRackReservation = new netbox.dcim.RackReservation("test", {
 *     rackId: testRack.id,
 *     units: [
 *         1,
 *         2,
 *         3,
 *         4,
 *         5,
 *     ],
 *     userId: 1,
 *     description: "my description",
 * });
 * ```
 */
export class RackReservation extends pulumi.CustomResource {
    /**
     * Get an existing RackReservation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RackReservationState, opts?: pulumi.CustomResourceOptions): RackReservation {
        return new RackReservation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:dcim/rackReservation:RackReservation';

    /**
     * Returns true if the given object is an instance of RackReservation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RackReservation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RackReservation.__pulumiType;
    }

    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string>;
    public readonly rackId!: pulumi.Output<number>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly tenantId!: pulumi.Output<number | undefined>;
    public readonly units!: pulumi.Output<number[]>;
    public readonly userId!: pulumi.Output<number>;

    /**
     * Create a RackReservation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RackReservationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RackReservationArgs | RackReservationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RackReservationState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["rackId"] = state ? state.rackId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["units"] = state ? state.units : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as RackReservationArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.rackId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rackId'");
            }
            if ((!args || args.units === undefined) && !opts.urn) {
                throw new Error("Missing required property 'units'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["rackId"] = args ? args.rackId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["units"] = args ? args.units : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RackReservation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RackReservation resources.
 */
export interface RackReservationState {
    comments?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    rackId?: pulumi.Input<number>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    units?: pulumi.Input<pulumi.Input<number>[]>;
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RackReservation resource.
 */
export interface RackReservationArgs {
    comments?: pulumi.Input<string>;
    description: pulumi.Input<string>;
    rackId: pulumi.Input<number>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    units: pulumi.Input<pulumi.Input<number>[]>;
    userId: pulumi.Input<number>;
}
