// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/powerpanel/):
 *
 * > A power panel represents the origin point in NetBox for electrical power being disseminated by one or more power feeds. In a data center environment, one power panel often serves a group of racks, with an individual power feed extending to each rack, though this is not always the case. It is common to have two sets of panels and feeds arranged in parallel to provide redundant power to each rack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * const testSite = new netbox.dcim.Site("testSite", {status: "active"});
 * const testLocation = new netbox.dcim.Location("testLocation", {siteId: testSite.id});
 * const testPowerPanel = new netbox.dcim.PowerPanel("testPowerPanel", {
 *     siteId: testSite.id,
 *     locationId: testLocation.id,
 * });
 * ```
 */
export class PowerPanel extends pulumi.CustomResource {
    /**
     * Get an existing PowerPanel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PowerPanelState, opts?: pulumi.CustomResourceOptions): PowerPanel {
        return new PowerPanel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:dcim/powerPanel:PowerPanel';

    /**
     * Returns true if the given object is an instance of PowerPanel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PowerPanel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PowerPanel.__pulumiType;
    }

    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly customFields!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly locationId!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly siteId!: pulumi.Output<number>;
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a PowerPanel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PowerPanelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PowerPanelArgs | PowerPanelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PowerPanelState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["customFields"] = state ? state.customFields : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["locationId"] = state ? state.locationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["siteId"] = state ? state.siteId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as PowerPanelArgs | undefined;
            if ((!args || args.siteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'siteId'");
            }
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["customFields"] = args ? args.customFields : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["locationId"] = args ? args.locationId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PowerPanel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PowerPanel resources.
 */
export interface PowerPanelState {
    comments?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    locationId?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    siteId?: pulumi.Input<number>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a PowerPanel resource.
 */
export interface PowerPanelArgs {
    comments?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    locationId?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    siteId: pulumi.Input<number>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
