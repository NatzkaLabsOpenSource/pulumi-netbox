// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > A VLAN Group represents a collection of VLANs. Generally, these are limited by one of a number of scopes such as "Site" or "Virtualization Cluster".
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * //Basic VLAN Group example
 * const example1 = new netbox.ipam.VlanGroup("example1", {
 *     name: "example1",
 *     slug: "example1",
 * });
 * //Full VLAN Group example
 * const example2 = new netbox.ipam.VlanGroup("example2", {
 *     name: "Second Example",
 *     slug: "example2",
 *     scopeType: "dcim.site",
 *     scopeId: example.id,
 *     description: "Second Example VLAN Group",
 *     tags: [exampleNetboxTag.id],
 *     vidRanges: [
 *         [
 *             1,
 *             2,
 *         ],
 *         [
 *             3,
 *             4,
 *         ],
 *     ],
 * });
 * ```
 */
export class VlanGroup extends pulumi.CustomResource {
    /**
     * Get an existing VlanGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VlanGroupState, opts?: pulumi.CustomResourceOptions): VlanGroup {
        return new VlanGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:ipam/vlanGroup:VlanGroup';

    /**
     * Returns true if the given object is an instance of VlanGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VlanGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VlanGroup.__pulumiType;
    }

    /**
     * Defaults to `""`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Required when `scopeType` is set.
     */
    public readonly scopeId!: pulumi.Output<number | undefined>;
    /**
     * Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
     */
    public readonly scopeType!: pulumi.Output<string | undefined>;
    public readonly slug!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly vidRanges!: pulumi.Output<number[][]>;

    /**
     * Create a VlanGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VlanGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VlanGroupArgs | VlanGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VlanGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
            resourceInputs["scopeType"] = state ? state.scopeType : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vidRanges"] = state ? state.vidRanges : undefined;
        } else {
            const args = argsOrState as VlanGroupArgs | undefined;
            if ((!args || args.slug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slug'");
            }
            if ((!args || args.vidRanges === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vidRanges'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
            resourceInputs["scopeType"] = args ? args.scopeType : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vidRanges"] = args ? args.vidRanges : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VlanGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VlanGroup resources.
 */
export interface VlanGroupState {
    /**
     * Defaults to `""`.
     */
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * Required when `scopeType` is set.
     */
    scopeId?: pulumi.Input<number>;
    /**
     * Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
     */
    scopeType?: pulumi.Input<string>;
    slug?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    vidRanges?: pulumi.Input<pulumi.Input<pulumi.Input<number>[]>[]>;
}

/**
 * The set of arguments for constructing a VlanGroup resource.
 */
export interface VlanGroupArgs {
    /**
     * Defaults to `""`.
     */
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * Required when `scopeType` is set.
     */
    scopeId?: pulumi.Input<number>;
    /**
     * Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
     */
    scopeType?: pulumi.Input<string>;
    slug: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    vidRanges: pulumi.Input<pulumi.Input<pulumi.Input<number>[]>[]>;
}
