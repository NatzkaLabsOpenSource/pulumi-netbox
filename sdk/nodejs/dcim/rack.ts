// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rack/):
 *
 * > The rack model represents a physical two- or four-post equipment rack in which devices can be installed. Each rack must be assigned to a site, and may optionally be assigned to a location within that site. Racks can also be organized by user-defined functional roles. The name and facility ID of each rack within a location must be unique.
 *
 * Rack height is measured in rack units (U); racks are commonly between 42U and 48U tall, but NetBox allows you to define racks of arbitrary height. A toggle is provided to indicate whether rack units are in ascending (from the ground up) or descending order.
 *
 * Each rack is assigned a name and (optionally) a separate facility ID. This is helpful when leasing space in a data center your organization does not own: The facility will often assign a seemingly arbitrary ID to a rack (for example, "M204.313") whereas internally you refer to is simply as "R113." A unique serial number and asset tag may also be associated with each rack.
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
 *     status: "reserved",
 *     width: 19,
 *     uHeight: 48,
 * });
 * ```
 */
export class Rack extends pulumi.CustomResource {
    /**
     * Get an existing Rack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RackState, opts?: pulumi.CustomResourceOptions): Rack {
        return new Rack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:dcim/rack:Rack';

    /**
     * Returns true if the given object is an instance of Rack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rack.__pulumiType;
    }

    public readonly assetTag!: pulumi.Output<string | undefined>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly customFields!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * If rack units are descending. Defaults to `false`.
     */
    public readonly descUnits!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly facilityId!: pulumi.Output<string | undefined>;
    /**
     * Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
     */
    public readonly formFactor!: pulumi.Output<string | undefined>;
    public readonly locationId!: pulumi.Output<number | undefined>;
    public readonly maxWeight!: pulumi.Output<number | undefined>;
    public readonly mountingDepth!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly outerDepth!: pulumi.Output<number | undefined>;
    /**
     * Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
     */
    public readonly outerUnit!: pulumi.Output<string | undefined>;
    public readonly outerWidth!: pulumi.Output<number | undefined>;
    public readonly roleId!: pulumi.Output<number | undefined>;
    public readonly serial!: pulumi.Output<string | undefined>;
    public readonly siteId!: pulumi.Output<number>;
    /**
     * Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
     */
    public readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly tenantId!: pulumi.Output<number | undefined>;
    public readonly uHeight!: pulumi.Output<number | undefined>;
    public readonly weight!: pulumi.Output<number | undefined>;
    /**
     * Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
     */
    public readonly weightUnit!: pulumi.Output<string | undefined>;
    /**
     * Valid values are `10`, `19`, `21` and `23`.
     */
    public readonly width!: pulumi.Output<number | undefined>;

    /**
     * Create a Rack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RackArgs | RackState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RackState | undefined;
            resourceInputs["assetTag"] = state ? state.assetTag : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["customFields"] = state ? state.customFields : undefined;
            resourceInputs["descUnits"] = state ? state.descUnits : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["facilityId"] = state ? state.facilityId : undefined;
            resourceInputs["formFactor"] = state ? state.formFactor : undefined;
            resourceInputs["locationId"] = state ? state.locationId : undefined;
            resourceInputs["maxWeight"] = state ? state.maxWeight : undefined;
            resourceInputs["mountingDepth"] = state ? state.mountingDepth : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outerDepth"] = state ? state.outerDepth : undefined;
            resourceInputs["outerUnit"] = state ? state.outerUnit : undefined;
            resourceInputs["outerWidth"] = state ? state.outerWidth : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
            resourceInputs["serial"] = state ? state.serial : undefined;
            resourceInputs["siteId"] = state ? state.siteId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["uHeight"] = state ? state.uHeight : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
            resourceInputs["weightUnit"] = state ? state.weightUnit : undefined;
            resourceInputs["width"] = state ? state.width : undefined;
        } else {
            const args = argsOrState as RackArgs | undefined;
            if ((!args || args.siteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'siteId'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["assetTag"] = args ? args.assetTag : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["customFields"] = args ? args.customFields : undefined;
            resourceInputs["descUnits"] = args ? args.descUnits : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["facilityId"] = args ? args.facilityId : undefined;
            resourceInputs["formFactor"] = args ? args.formFactor : undefined;
            resourceInputs["locationId"] = args ? args.locationId : undefined;
            resourceInputs["maxWeight"] = args ? args.maxWeight : undefined;
            resourceInputs["mountingDepth"] = args ? args.mountingDepth : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outerDepth"] = args ? args.outerDepth : undefined;
            resourceInputs["outerUnit"] = args ? args.outerUnit : undefined;
            resourceInputs["outerWidth"] = args ? args.outerWidth : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
            resourceInputs["serial"] = args ? args.serial : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["uHeight"] = args ? args.uHeight : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
            resourceInputs["weightUnit"] = args ? args.weightUnit : undefined;
            resourceInputs["width"] = args ? args.width : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Rack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rack resources.
 */
export interface RackState {
    assetTag?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If rack units are descending. Defaults to `false`.
     */
    descUnits?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    facilityId?: pulumi.Input<string>;
    /**
     * Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
     */
    formFactor?: pulumi.Input<string>;
    locationId?: pulumi.Input<number>;
    maxWeight?: pulumi.Input<number>;
    mountingDepth?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    outerDepth?: pulumi.Input<number>;
    /**
     * Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
     */
    outerUnit?: pulumi.Input<string>;
    outerWidth?: pulumi.Input<number>;
    roleId?: pulumi.Input<number>;
    serial?: pulumi.Input<string>;
    siteId?: pulumi.Input<number>;
    /**
     * Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
     */
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    uHeight?: pulumi.Input<number>;
    weight?: pulumi.Input<number>;
    /**
     * Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
     */
    weightUnit?: pulumi.Input<string>;
    /**
     * Valid values are `10`, `19`, `21` and `23`.
     */
    width?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Rack resource.
 */
export interface RackArgs {
    assetTag?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If rack units are descending. Defaults to `false`.
     */
    descUnits?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    facilityId?: pulumi.Input<string>;
    /**
     * Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
     */
    formFactor?: pulumi.Input<string>;
    locationId?: pulumi.Input<number>;
    maxWeight?: pulumi.Input<number>;
    mountingDepth?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    outerDepth?: pulumi.Input<number>;
    /**
     * Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
     */
    outerUnit?: pulumi.Input<string>;
    outerWidth?: pulumi.Input<number>;
    roleId?: pulumi.Input<number>;
    serial?: pulumi.Input<string>;
    siteId: pulumi.Input<number>;
    /**
     * Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
     */
    status: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    uHeight?: pulumi.Input<number>;
    weight?: pulumi.Input<number>;
    /**
     * Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
     */
    weightUnit?: pulumi.Input<string>;
    /**
     * Valid values are `10`, `19`, `21` and `23`.
     */
    width?: pulumi.Input<number>;
}
