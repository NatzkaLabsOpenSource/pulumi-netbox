// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/modulebay/):
 *
 * > Module bays represent a space or slot within a device in which a field-replaceable module may be installed. A common example is that of a chassis-based switch such as the Cisco Nexus 9000 or Juniper EX9200. Modules in turn hold additional components that become available to the parent device.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * // Note that some terraform code is not included in the example for brevity
 * const testDevice = new netbox.dcim.Device("testDevice", {
 *     deviceTypeId: netbox_device_type.test.id,
 *     roleId: netbox_device_role.test.id,
 *     siteId: netbox_site.test.id,
 * });
 * const testDeviceModuleBay = new netbox.dcim.DeviceModuleBay("testDeviceModuleBay", {deviceId: testDevice.id});
 * ```
 */
export class DeviceModuleBay extends pulumi.CustomResource {
    /**
     * Get an existing DeviceModuleBay resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceModuleBayState, opts?: pulumi.CustomResourceOptions): DeviceModuleBay {
        return new DeviceModuleBay(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:dcim/deviceModuleBay:DeviceModuleBay';

    /**
     * Returns true if the given object is an instance of DeviceModuleBay.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceModuleBay {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceModuleBay.__pulumiType;
    }

    public readonly customFields!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly deviceId!: pulumi.Output<number>;
    public readonly label!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly position!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a DeviceModuleBay resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceModuleBayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceModuleBayArgs | DeviceModuleBayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceModuleBayState | undefined;
            resourceInputs["customFields"] = state ? state.customFields : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DeviceModuleBayArgs | undefined;
            if ((!args || args.deviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceId'");
            }
            resourceInputs["customFields"] = args ? args.customFields : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeviceModuleBay.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeviceModuleBay resources.
 */
export interface DeviceModuleBayState {
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    deviceId?: pulumi.Input<number>;
    label?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    position?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DeviceModuleBay resource.
 */
export interface DeviceModuleBayArgs {
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    deviceId: pulumi.Input<number>;
    label?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    position?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
