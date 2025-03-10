// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/services/#services):
 *
 * > A service represents a layer four TCP or UDP service available on a device or virtual machine. For example, you might want to document that an HTTP service is running on a device. Each service includes a name, protocol, and port number; for example, "SSH (TCP/22)" or "DNS (UDP/53)."
 * > 
 * > A service may optionally be bound to one or more specific IP addresses belonging to its parent device or VM. (If no IP addresses are bound, the service is assumed to be reachable via any assigned IP address.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:ipam/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    public readonly customFields!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Exactly one of `virtualMachineId` or `deviceId` must be given.
     */
    public readonly deviceId!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Exactly one of `port` or `ports` must be given.
     *
     * @deprecated This field is deprecated. Please use the new "ports" attribute instead.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * Exactly one of `port` or `ports` must be given.
     */
    public readonly ports!: pulumi.Output<number[] | undefined>;
    /**
     * Valid values are `tcp`, `udp` and `sctp`.
     */
    public readonly protocol!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Exactly one of `virtualMachineId` or `deviceId` must be given.
     */
    public readonly virtualMachineId!: pulumi.Output<number | undefined>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["customFields"] = state ? state.customFields : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["customFields"] = args ? args.customFields : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Exactly one of `virtualMachineId` or `deviceId` must be given.
     */
    deviceId?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    /**
     * Exactly one of `port` or `ports` must be given.
     *
     * @deprecated This field is deprecated. Please use the new "ports" attribute instead.
     */
    port?: pulumi.Input<number>;
    /**
     * Exactly one of `port` or `ports` must be given.
     */
    ports?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Valid values are `tcp`, `udp` and `sctp`.
     */
    protocol?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Exactly one of `virtualMachineId` or `deviceId` must be given.
     */
    virtualMachineId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Exactly one of `virtualMachineId` or `deviceId` must be given.
     */
    deviceId?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    /**
     * Exactly one of `port` or `ports` must be given.
     *
     * @deprecated This field is deprecated. Please use the new "ports" attribute instead.
     */
    port?: pulumi.Input<number>;
    /**
     * Exactly one of `port` or `ports` must be given.
     */
    ports?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Valid values are `tcp`, `udp` and `sctp`.
     */
    protocol: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Exactly one of `virtualMachineId` or `deviceId` must be given.
     */
    virtualMachineId?: pulumi.Input<number>;
}
