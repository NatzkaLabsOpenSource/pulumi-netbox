// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/moduletype/):
 *
 * > A module type represents a specific make and model of hardware component which is installable within a device's module bay and has its own child components. For example, consider a chassis-based switch or router with a number of field-replaceable line cards. Each line card has its own model number and includes a certain set of components such as interfaces. Each module type may have a manufacturer, model number, and part number assigned to it.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * const test = new netbox.dcim.Manufacturer("test", {name: "Dell"});
 * const testModuleType = new netbox.dcim.ModuleType("test", {
 *     manufacturerId: test.id,
 *     model: "Networking",
 * });
 * ```
 */
export class ModuleType extends pulumi.CustomResource {
    /**
     * Get an existing ModuleType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ModuleTypeState, opts?: pulumi.CustomResourceOptions): ModuleType {
        return new ModuleType(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:dcim/moduleType:ModuleType';

    /**
     * Returns true if the given object is an instance of ModuleType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ModuleType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ModuleType.__pulumiType;
    }

    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly customFields!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly manufacturerId!: pulumi.Output<number>;
    public readonly model!: pulumi.Output<string>;
    public readonly partNumber!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly weight!: pulumi.Output<number | undefined>;
    /**
     * One of [kg, g, lb, oz]. Required when `weight` is set.
     */
    public readonly weightUnit!: pulumi.Output<string | undefined>;

    /**
     * Create a ModuleType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModuleTypeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ModuleTypeArgs | ModuleTypeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ModuleTypeState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["customFields"] = state ? state.customFields : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["manufacturerId"] = state ? state.manufacturerId : undefined;
            resourceInputs["model"] = state ? state.model : undefined;
            resourceInputs["partNumber"] = state ? state.partNumber : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
            resourceInputs["weightUnit"] = state ? state.weightUnit : undefined;
        } else {
            const args = argsOrState as ModuleTypeArgs | undefined;
            if ((!args || args.manufacturerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'manufacturerId'");
            }
            if ((!args || args.model === undefined) && !opts.urn) {
                throw new Error("Missing required property 'model'");
            }
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["customFields"] = args ? args.customFields : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["manufacturerId"] = args ? args.manufacturerId : undefined;
            resourceInputs["model"] = args ? args.model : undefined;
            resourceInputs["partNumber"] = args ? args.partNumber : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
            resourceInputs["weightUnit"] = args ? args.weightUnit : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ModuleType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ModuleType resources.
 */
export interface ModuleTypeState {
    comments?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    manufacturerId?: pulumi.Input<number>;
    model?: pulumi.Input<string>;
    partNumber?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    weight?: pulumi.Input<number>;
    /**
     * One of [kg, g, lb, oz]. Required when `weight` is set.
     */
    weightUnit?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ModuleType resource.
 */
export interface ModuleTypeArgs {
    comments?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    manufacturerId: pulumi.Input<number>;
    model: pulumi.Input<string>;
    partNumber?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    weight?: pulumi.Input<number>;
    /**
     * One of [kg, g, lb, oz]. Required when `weight` is set.
     */
    weightUnit?: pulumi.Input<string>;
}
