// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rackrole/):
 *
 * > Each rack can optionally be assigned a user-defined functional role. For example, you might designate a rack for compute or storage resources, or to house colocated customer devices.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzkalabsopensource/pulumi-netbox";
 *
 * const test = new netbox.dcim.RackRole("test", {colorHex: "111111"});
 * ```
 */
export class RackRole extends pulumi.CustomResource {
    /**
     * Get an existing RackRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RackRoleState, opts?: pulumi.CustomResourceOptions): RackRole {
        return new RackRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:dcim/rackRole:RackRole';

    /**
     * Returns true if the given object is an instance of RackRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RackRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RackRole.__pulumiType;
    }

    public readonly colorHex!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly slug!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a RackRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RackRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RackRoleArgs | RackRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RackRoleState | undefined;
            resourceInputs["colorHex"] = state ? state.colorHex : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RackRoleArgs | undefined;
            if ((!args || args.colorHex === undefined) && !opts.urn) {
                throw new Error("Missing required property 'colorHex'");
            }
            resourceInputs["colorHex"] = args ? args.colorHex : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RackRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RackRole resources.
 */
export interface RackRoleState {
    colorHex?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a RackRole resource.
 */
export interface RackRoleArgs {
    colorHex: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}