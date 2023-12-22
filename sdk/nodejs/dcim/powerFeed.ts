// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/powerfeed/):
 *
 * > A power feed represents the distribution of power from a power panel to a particular device, typically a power distribution unit (PDU). The power port (inlet) on a device can be connected via a cable to a power feed. A power feed may optionally be assigned to a rack to allow more easily tracking the distribution of power among racks.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzkalabsopensource/pulumi-netbox";
 *
 * const testSite = new netbox.dcim.Site("testSite", {status: "active"});
 * const testLocation = new netbox.dcim.Location("testLocation", {siteId: testSite.id});
 * const testPowerPanel = new netbox.dcim.PowerPanel("testPowerPanel", {
 *     siteId: testSite.id,
 *     locationId: testLocation.id,
 * });
 * const testPowerFeed = new netbox.dcim.PowerFeed("testPowerFeed", {
 *     powerPanelId: testPowerPanel.id,
 *     status: "active",
 *     type: "primary",
 *     supply: "ac",
 *     phase: "single-phase",
 *     voltage: 250,
 *     amperage: 100,
 *     maxPercentUtilization: 80,
 * });
 * ```
 */
export class PowerFeed extends pulumi.CustomResource {
    /**
     * Get an existing PowerFeed resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PowerFeedState, opts?: pulumi.CustomResourceOptions): PowerFeed {
        return new PowerFeed(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:dcim/powerFeed:PowerFeed';

    /**
     * Returns true if the given object is an instance of PowerFeed.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PowerFeed {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PowerFeed.__pulumiType;
    }

    public readonly amperage!: pulumi.Output<number>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly customFields!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `false`.
     */
    public readonly markConnected!: pulumi.Output<boolean | undefined>;
    public readonly maxPercentUtilization!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    /**
     * One of [single-phase, three-phase].
     */
    public readonly phase!: pulumi.Output<string>;
    public readonly powerPanelId!: pulumi.Output<number>;
    public readonly rackId!: pulumi.Output<number | undefined>;
    /**
     * One of [offline, active, planned, failed].
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * One of [ac, dc].
     */
    public readonly supply!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * One of [primary, redundant].
     */
    public readonly type!: pulumi.Output<string>;
    public readonly voltage!: pulumi.Output<number>;

    /**
     * Create a PowerFeed resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PowerFeedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PowerFeedArgs | PowerFeedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PowerFeedState | undefined;
            resourceInputs["amperage"] = state ? state.amperage : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["customFields"] = state ? state.customFields : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["markConnected"] = state ? state.markConnected : undefined;
            resourceInputs["maxPercentUtilization"] = state ? state.maxPercentUtilization : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["phase"] = state ? state.phase : undefined;
            resourceInputs["powerPanelId"] = state ? state.powerPanelId : undefined;
            resourceInputs["rackId"] = state ? state.rackId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["supply"] = state ? state.supply : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["voltage"] = state ? state.voltage : undefined;
        } else {
            const args = argsOrState as PowerFeedArgs | undefined;
            if ((!args || args.amperage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'amperage'");
            }
            if ((!args || args.maxPercentUtilization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxPercentUtilization'");
            }
            if ((!args || args.phase === undefined) && !opts.urn) {
                throw new Error("Missing required property 'phase'");
            }
            if ((!args || args.powerPanelId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'powerPanelId'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.supply === undefined) && !opts.urn) {
                throw new Error("Missing required property 'supply'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.voltage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'voltage'");
            }
            resourceInputs["amperage"] = args ? args.amperage : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["customFields"] = args ? args.customFields : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["markConnected"] = args ? args.markConnected : undefined;
            resourceInputs["maxPercentUtilization"] = args ? args.maxPercentUtilization : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["phase"] = args ? args.phase : undefined;
            resourceInputs["powerPanelId"] = args ? args.powerPanelId : undefined;
            resourceInputs["rackId"] = args ? args.rackId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["supply"] = args ? args.supply : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["voltage"] = args ? args.voltage : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PowerFeed.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PowerFeed resources.
 */
export interface PowerFeedState {
    amperage?: pulumi.Input<number>;
    comments?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Defaults to `false`.
     */
    markConnected?: pulumi.Input<boolean>;
    maxPercentUtilization?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    /**
     * One of [single-phase, three-phase].
     */
    phase?: pulumi.Input<string>;
    powerPanelId?: pulumi.Input<number>;
    rackId?: pulumi.Input<number>;
    /**
     * One of [offline, active, planned, failed].
     */
    status?: pulumi.Input<string>;
    /**
     * One of [ac, dc].
     */
    supply?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One of [primary, redundant].
     */
    type?: pulumi.Input<string>;
    voltage?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PowerFeed resource.
 */
export interface PowerFeedArgs {
    amperage: pulumi.Input<number>;
    comments?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Defaults to `false`.
     */
    markConnected?: pulumi.Input<boolean>;
    maxPercentUtilization: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    /**
     * One of [single-phase, three-phase].
     */
    phase: pulumi.Input<string>;
    powerPanelId: pulumi.Input<number>;
    rackId?: pulumi.Input<number>;
    /**
     * One of [offline, active, planned, failed].
     */
    status: pulumi.Input<string>;
    /**
     * One of [ac, dc].
     */
    supply: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One of [primary, redundant].
     */
    type: pulumi.Input<string>;
    voltage: pulumi.Input<number>;
}