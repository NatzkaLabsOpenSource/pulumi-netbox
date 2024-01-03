// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 * import * as netbox from "@pulumi/netbox";
 *
 * const testPrefix = netbox.ipam.getPrefix({
 *     cidr: "10.0.0.0/24",
 * });
 * const testAvailablePrefix = new netbox.ipam.AvailablePrefix("testAvailablePrefix", {
 *     parentPrefixId: testPrefix.then(testPrefix => testPrefix.id),
 *     prefixLength: 25,
 *     status: "active",
 * });
 * ```
 */
export class AvailablePrefix extends pulumi.CustomResource {
    /**
     * Get an existing AvailablePrefix resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AvailablePrefixState, opts?: pulumi.CustomResourceOptions): AvailablePrefix {
        return new AvailablePrefix(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:ipam/availablePrefix:AvailablePrefix';

    /**
     * Returns true if the given object is an instance of AvailablePrefix.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AvailablePrefix {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AvailablePrefix.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly isPool!: pulumi.Output<boolean | undefined>;
    public readonly markUtilized!: pulumi.Output<boolean | undefined>;
    public readonly parentPrefixId!: pulumi.Output<number>;
    public /*out*/ readonly prefix!: pulumi.Output<string>;
    public readonly prefixLength!: pulumi.Output<number>;
    public readonly roleId!: pulumi.Output<number | undefined>;
    public readonly siteId!: pulumi.Output<number | undefined>;
    /**
     * Valid values are `active`, `container`, `reserved` and `deprecated`.
     */
    public readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly tenantId!: pulumi.Output<number | undefined>;
    public readonly vlanId!: pulumi.Output<number | undefined>;
    public readonly vrfId!: pulumi.Output<number | undefined>;

    /**
     * Create a AvailablePrefix resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AvailablePrefixArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AvailablePrefixArgs | AvailablePrefixState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AvailablePrefixState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isPool"] = state ? state.isPool : undefined;
            resourceInputs["markUtilized"] = state ? state.markUtilized : undefined;
            resourceInputs["parentPrefixId"] = state ? state.parentPrefixId : undefined;
            resourceInputs["prefix"] = state ? state.prefix : undefined;
            resourceInputs["prefixLength"] = state ? state.prefixLength : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
            resourceInputs["siteId"] = state ? state.siteId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["vlanId"] = state ? state.vlanId : undefined;
            resourceInputs["vrfId"] = state ? state.vrfId : undefined;
        } else {
            const args = argsOrState as AvailablePrefixArgs | undefined;
            if ((!args || args.parentPrefixId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parentPrefixId'");
            }
            if ((!args || args.prefixLength === undefined) && !opts.urn) {
                throw new Error("Missing required property 'prefixLength'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isPool"] = args ? args.isPool : undefined;
            resourceInputs["markUtilized"] = args ? args.markUtilized : undefined;
            resourceInputs["parentPrefixId"] = args ? args.parentPrefixId : undefined;
            resourceInputs["prefixLength"] = args ? args.prefixLength : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["vlanId"] = args ? args.vlanId : undefined;
            resourceInputs["vrfId"] = args ? args.vrfId : undefined;
            resourceInputs["prefix"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AvailablePrefix.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AvailablePrefix resources.
 */
export interface AvailablePrefixState {
    description?: pulumi.Input<string>;
    isPool?: pulumi.Input<boolean>;
    markUtilized?: pulumi.Input<boolean>;
    parentPrefixId?: pulumi.Input<number>;
    prefix?: pulumi.Input<string>;
    prefixLength?: pulumi.Input<number>;
    roleId?: pulumi.Input<number>;
    siteId?: pulumi.Input<number>;
    /**
     * Valid values are `active`, `container`, `reserved` and `deprecated`.
     */
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    vlanId?: pulumi.Input<number>;
    vrfId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AvailablePrefix resource.
 */
export interface AvailablePrefixArgs {
    description?: pulumi.Input<string>;
    isPool?: pulumi.Input<boolean>;
    markUtilized?: pulumi.Input<boolean>;
    parentPrefixId: pulumi.Input<number>;
    prefixLength: pulumi.Input<number>;
    roleId?: pulumi.Input<number>;
    siteId?: pulumi.Input<number>;
    /**
     * Valid values are `active`, `container`, `reserved` and `deprecated`.
     */
    status: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    vlanId?: pulumi.Input<number>;
    vrfId?: pulumi.Input<number>;
}
