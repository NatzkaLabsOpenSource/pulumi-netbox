// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/facilities/#site-groups):
 *
 * > Like regions, site groups can be arranged in a recursive hierarchy for grouping sites. However, whereas regions are intended for geographic organization, site groups may be used for functional grouping. For example, you might classify sites as corporate, branch, or customer sites in addition to where they are physically located.
 * > 
 * > The use of both regions and site groups affords to independent but complementary dimensions across which sites can be organized.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * const parent = new netbox.dcim.SiteGroup("parent", {
 *     name: "parent",
 *     description: "sample description",
 * });
 * const child = new netbox.dcim.SiteGroup("child", {
 *     name: "child",
 *     description: "sample description",
 *     parentId: parent.id,
 * });
 * ```
 */
export class SiteGroup extends pulumi.CustomResource {
    /**
     * Get an existing SiteGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SiteGroupState, opts?: pulumi.CustomResourceOptions): SiteGroup {
        return new SiteGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:dcim/siteGroup:SiteGroup';

    /**
     * Returns true if the given object is an instance of SiteGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SiteGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SiteGroup.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly parentId!: pulumi.Output<number | undefined>;
    public readonly slug!: pulumi.Output<string>;

    /**
     * Create a SiteGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SiteGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SiteGroupArgs | SiteGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SiteGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentId"] = state ? state.parentId : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as SiteGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentId"] = args ? args.parentId : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SiteGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SiteGroup resources.
 */
export interface SiteGroupState {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    parentId?: pulumi.Input<number>;
    slug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SiteGroup resource.
 */
export interface SiteGroupArgs {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    parentId?: pulumi.Input<number>;
    slug?: pulumi.Input<string>;
}
