// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuits_1):
 *
 * > A communications circuit represents a single physical link connecting exactly two endpoints, commonly referred to as its A and Z terminations. A circuit in NetBox may have zero, one, or two terminations defined. It is common to have only one termination defined when you don't necessarily care about the details of the provider side of the circuit, e.g. for Internet access circuits. Both terminations would likely be modeled for circuits which connect one customer site to another.
 * > 
 * > Each circuit is associated with a provider and a user-defined type. For example, you might have Internet access circuits delivered to each site by one provider, and private MPLS circuits delivered by another. Each circuit must be assigned a circuit ID, each of which must be unique per provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzkalabsopensource/pulumi-netbox";
 *
 * const testTenant = new netbox.tenancy.Tenant("testTenant", {});
 * const testCircuitProvider = new netbox.circuit.CircuitProvider("testCircuitProvider", {});
 * const testCircuitType = new netbox.circuit.CircuitType("testCircuitType", {});
 * const testCircuit = new netbox.circuit.Circuit("testCircuit", {
 *     cid: "test",
 *     status: "active",
 *     providerId: testCircuitProvider.id,
 *     typeId: testCircuitType.id,
 * });
 * ```
 */
export class Circuit extends pulumi.CustomResource {
    /**
     * Get an existing Circuit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CircuitState, opts?: pulumi.CustomResourceOptions): Circuit {
        return new Circuit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:circuit/circuit:Circuit';

    /**
     * Returns true if the given object is an instance of Circuit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Circuit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Circuit.__pulumiType;
    }

    public readonly cid!: pulumi.Output<string>;
    public readonly providerId!: pulumi.Output<number>;
    /**
     * Valid values are `planned`, `provisioning`, `active`, `offline`, `deprovisioning` and `decommissioning`.
     */
    public readonly status!: pulumi.Output<string>;
    public readonly tenantId!: pulumi.Output<number | undefined>;
    public readonly typeId!: pulumi.Output<number>;

    /**
     * Create a Circuit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CircuitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CircuitArgs | CircuitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CircuitState | undefined;
            resourceInputs["cid"] = state ? state.cid : undefined;
            resourceInputs["providerId"] = state ? state.providerId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["typeId"] = state ? state.typeId : undefined;
        } else {
            const args = argsOrState as CircuitArgs | undefined;
            if ((!args || args.cid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cid'");
            }
            if ((!args || args.providerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerId'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.typeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'typeId'");
            }
            resourceInputs["cid"] = args ? args.cid : undefined;
            resourceInputs["providerId"] = args ? args.providerId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["typeId"] = args ? args.typeId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Circuit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Circuit resources.
 */
export interface CircuitState {
    cid?: pulumi.Input<string>;
    providerId?: pulumi.Input<number>;
    /**
     * Valid values are `planned`, `provisioning`, `active`, `offline`, `deprovisioning` and `decommissioning`.
     */
    status?: pulumi.Input<string>;
    tenantId?: pulumi.Input<number>;
    typeId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Circuit resource.
 */
export interface CircuitArgs {
    cid: pulumi.Input<string>;
    providerId: pulumi.Input<number>;
    /**
     * Valid values are `planned`, `provisioning`, `active`, `offline`, `deprovisioning` and `decommissioning`.
     */
    status: pulumi.Input<string>;
    tenantId?: pulumi.Input<number>;
    typeId: pulumi.Input<number>;
}
