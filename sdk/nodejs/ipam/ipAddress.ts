// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#ip-addresses):
 *
 * > An IP address comprises a single host address (either IPv4 or IPv6) and its subnet mask. Its mask should match exactly how the IP address is configured on an interface in the real world.
 * > 
 * > Like a prefix, an IP address can optionally be assigned to a VRF (otherwise, it will appear in the "global" table). IP addresses are automatically arranged under parent prefixes within their respective VRFs according to the IP hierarchy.
 *
 * ## Example Usage
 *
 * ### Creating an IP address that is assigned to a virtual machine interface
 *
 * Starting with provider version 3.5.0, you can use the `virtualMachineInterfaceId` attribute to assign an IP address to a virtual machine interface.
 * You can also use the `interfaceId` and `objectType` attributes instead.
 *
 * With `virtualMachineInterfaceId`:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * // Assuming a virtual machine with the id `123` exists
 * const _this = new netbox.virt.Interface("this", {
 *     name: "eth0",
 *     virtualMachineId: 123,
 * });
 * const thisIpAddress = new netbox.ipam.IpAddress("this", {
 *     ipAddress: "10.0.0.60/24",
 *     status: "active",
 *     virtualMachineInterfaceId: _this.id,
 * });
 * ```
 *
 * With `objectType` and `interfaceId`:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * // Assuming a virtual machine with the id `123` exists
 * const _this = new netbox.virt.Interface("this", {
 *     name: "eth0",
 *     virtualMachineId: 123,
 * });
 * const thisIpAddress = new netbox.ipam.IpAddress("this", {
 *     ipAddress: "10.0.0.60/24",
 *     status: "active",
 *     interfaceId: _this.id,
 *     objectType: "virtualization.vminterface",
 * });
 * ```
 *
 * ### Creating an IP address that is assigned to a device interface
 *
 * Starting with provider version 3.5.0, you can use the `deviceInterfaceId` attribute to assign an IP address to a device interface.
 * You can also use the `interfaceId` and `objectType` attributes instead.
 *
 * With `deviceInterfaceId`:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * // Assuming a device with the id `123` exists
 * const _this = new netbox.dcim.DeviceInterface("this", {
 *     name: "eth0",
 *     deviceId: 123,
 *     type: "1000base-t",
 * });
 * const thisIpAddress = new netbox.ipam.IpAddress("this", {
 *     ipAddress: "10.0.0.60/24",
 *     status: "active",
 *     deviceInterfaceId: _this.id,
 * });
 * ```
 *
 * With `objectType` and `interfaceId`:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * // Assuming a device with the id `123` exists
 * const _this = new netbox.dcim.DeviceInterface("this", {
 *     name: "eth0",
 *     deviceId: 123,
 *     type: "1000base-t",
 * });
 * const thisIpAddress = new netbox.ipam.IpAddress("this", {
 *     ipAddress: "10.0.0.60/24",
 *     status: "active",
 *     interfaceId: _this.id,
 *     objectType: "dcim.interface",
 * });
 * ```
 *
 * ### Creating an IP address that is not assigned to anything
 *
 * You can create an IP address that is not assigend to anything by omitting the attributes mentioned above.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * const _this = new netbox.ipam.IpAddress("this", {
 *     ipAddress: "10.0.0.50/24",
 *     status: "reserved",
 * });
 * ```
 */
export class IpAddress extends pulumi.CustomResource {
    /**
     * Get an existing IpAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpAddressState, opts?: pulumi.CustomResourceOptions): IpAddress {
        return new IpAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:ipam/ipAddress:IpAddress';

    /**
     * Returns true if the given object is an instance of IpAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpAddress.__pulumiType;
    }

    public readonly customFields!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Conflicts with `interfaceId` and `virtualMachineInterfaceId`.
     */
    public readonly deviceInterfaceId!: pulumi.Output<number | undefined>;
    public readonly dnsName!: pulumi.Output<string | undefined>;
    /**
     * Required when `objectType` is set.
     */
    public readonly interfaceId!: pulumi.Output<number | undefined>;
    public readonly ipAddress!: pulumi.Output<string>;
    public readonly natInsideAddressId!: pulumi.Output<number | undefined>;
    public /*out*/ readonly natOutsideAddresses!: pulumi.Output<outputs.ipam.IpAddressNatOutsideAddress[]>;
    /**
     * Valid values are `virtualization.vminterface` and `dcim.interface`. Required when `interfaceId` is set.
     */
    public readonly objectType!: pulumi.Output<string | undefined>;
    /**
     * Valid values are `loopback`, `secondary`, `anycast`, `vip`, `vrrp`, `hsrp`, `glbp` and `carp`.
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * Valid values are `active`, `reserved`, `deprecated`, `dhcp` and `slaac`.
     */
    public readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly tenantId!: pulumi.Output<number | undefined>;
    /**
     * Conflicts with `interfaceId` and `deviceInterfaceId`.
     */
    public readonly virtualMachineInterfaceId!: pulumi.Output<number | undefined>;
    public readonly vrfId!: pulumi.Output<number | undefined>;

    /**
     * Create a IpAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpAddressArgs | IpAddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpAddressState | undefined;
            resourceInputs["customFields"] = state ? state.customFields : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceInterfaceId"] = state ? state.deviceInterfaceId : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["natInsideAddressId"] = state ? state.natInsideAddressId : undefined;
            resourceInputs["natOutsideAddresses"] = state ? state.natOutsideAddresses : undefined;
            resourceInputs["objectType"] = state ? state.objectType : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["virtualMachineInterfaceId"] = state ? state.virtualMachineInterfaceId : undefined;
            resourceInputs["vrfId"] = state ? state.vrfId : undefined;
        } else {
            const args = argsOrState as IpAddressArgs | undefined;
            if ((!args || args.ipAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddress'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["customFields"] = args ? args.customFields : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceInterfaceId"] = args ? args.deviceInterfaceId : undefined;
            resourceInputs["dnsName"] = args ? args.dnsName : undefined;
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["natInsideAddressId"] = args ? args.natInsideAddressId : undefined;
            resourceInputs["objectType"] = args ? args.objectType : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["virtualMachineInterfaceId"] = args ? args.virtualMachineInterfaceId : undefined;
            resourceInputs["vrfId"] = args ? args.vrfId : undefined;
            resourceInputs["natOutsideAddresses"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpAddress resources.
 */
export interface IpAddressState {
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Conflicts with `interfaceId` and `virtualMachineInterfaceId`.
     */
    deviceInterfaceId?: pulumi.Input<number>;
    dnsName?: pulumi.Input<string>;
    /**
     * Required when `objectType` is set.
     */
    interfaceId?: pulumi.Input<number>;
    ipAddress?: pulumi.Input<string>;
    natInsideAddressId?: pulumi.Input<number>;
    natOutsideAddresses?: pulumi.Input<pulumi.Input<inputs.ipam.IpAddressNatOutsideAddress>[]>;
    /**
     * Valid values are `virtualization.vminterface` and `dcim.interface`. Required when `interfaceId` is set.
     */
    objectType?: pulumi.Input<string>;
    /**
     * Valid values are `loopback`, `secondary`, `anycast`, `vip`, `vrrp`, `hsrp`, `glbp` and `carp`.
     */
    role?: pulumi.Input<string>;
    /**
     * Valid values are `active`, `reserved`, `deprecated`, `dhcp` and `slaac`.
     */
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    /**
     * Conflicts with `interfaceId` and `deviceInterfaceId`.
     */
    virtualMachineInterfaceId?: pulumi.Input<number>;
    vrfId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a IpAddress resource.
 */
export interface IpAddressArgs {
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * Conflicts with `interfaceId` and `virtualMachineInterfaceId`.
     */
    deviceInterfaceId?: pulumi.Input<number>;
    dnsName?: pulumi.Input<string>;
    /**
     * Required when `objectType` is set.
     */
    interfaceId?: pulumi.Input<number>;
    ipAddress: pulumi.Input<string>;
    natInsideAddressId?: pulumi.Input<number>;
    /**
     * Valid values are `virtualization.vminterface` and `dcim.interface`. Required when `interfaceId` is set.
     */
    objectType?: pulumi.Input<string>;
    /**
     * Valid values are `loopback`, `secondary`, `anycast`, `vip`, `vrrp`, `hsrp`, `glbp` and `carp`.
     */
    role?: pulumi.Input<string>;
    /**
     * Valid values are `active`, `reserved`, `deprecated`, `dhcp` and `slaac`.
     */
    status: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<number>;
    /**
     * Conflicts with `interfaceId` and `deviceInterfaceId`.
     */
    virtualMachineInterfaceId?: pulumi.Input<number>;
    vrfId?: pulumi.Input<number>;
}
