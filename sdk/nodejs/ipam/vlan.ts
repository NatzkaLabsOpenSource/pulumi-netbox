// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/vlans/#vlans):
 *
 * > A VLAN represents an isolated layer two domain, identified by a name and a numeric ID (1-4094) as defined in IEEE 802.1Q. VLANs are arranged into VLAN groups to define scope and to enforce uniqueness.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzkalabsopensource/pulumi-netbox";
 *
 * const example1 = new netbox.ipam.Vlan("example1", {
 *     vid: 1777,
 *     tags: [],
 * });
 * // Assume netbox_tenant, netbox_site, and netbox_tag resources exist
 * const example2 = new netbox.ipam.Vlan("example2", {
 *     vid: 1778,
 *     status: "reserved",
 *     description: "Reserved example VLAN",
 *     tenantId: netbox_tenant.ex.id,
 *     siteId: netbox_site.ex.id,
 *     groupId: netbox_vlan_group.ex.id,
 *     tags: [netbox_tag.ex.name],
 * });
 * ```
 */
export class Vlan extends pulumi.CustomResource {
    /**
     * Get an existing Vlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VlanState, opts?: pulumi.CustomResourceOptions): Vlan {
        return new Vlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:ipam/vlan:Vlan';

    /**
     * Returns true if the given object is an instance of Vlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vlan.__pulumiType;
    }

    /**
     * Defaults to `""`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly groupId!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly roleId!: pulumi.Output<number | undefined>;
    public readonly siteId!: pulumi.Output<number | undefined>;
    /**
     * Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly tenantId!: pulumi.Output<number | undefined>;
    public readonly vid!: pulumi.Output<number>;

    /**
     * Create a Vlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VlanArgs | VlanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VlanState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
            resourceInputs["siteId"] = state ? state.siteId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["vid"] = state ? state.vid : undefined;
        } else {
            const args = argsOrState as VlanArgs | undefined;
            if ((!args || args.vid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vid'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["vid"] = args ? args.vid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vlan resources.
 */
export interface VlanState {
    /**
     * Defaults to `""`.
     */
    description?: pulumi.Input<string>;
    groupId?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    roleId?: pulumi.Input<number>;
    siteId?: pulumi.Input<number>;
    /**
     * Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
     */
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    vid?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Vlan resource.
 */
export interface VlanArgs {
    /**
     * Defaults to `""`.
     */
    description?: pulumi.Input<string>;
    groupId?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    roleId?: pulumi.Input<number>;
    siteId?: pulumi.Input<number>;
    /**
     * Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
     */
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    vid: pulumi.Input<number>;
}
