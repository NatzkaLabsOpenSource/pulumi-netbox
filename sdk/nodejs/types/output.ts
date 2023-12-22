// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetTagsFilter {
    name: string;
    value: string;
}

export interface GetTagsTag {
    color?: string;
    description?: string;
    name: string;
    slug: string;
    tagId: number;
}

export namespace dcim {
    export interface CableATermination {
        objectId: number;
        objectType: string;
    }

    export interface CableBTermination {
        objectId: number;
        objectType: string;
    }

    export interface GetDeviceInterfacesFilter {
        name: string;
        value: string;
    }

    export interface GetDeviceInterfacesInterface {
        description: string;
        deviceId: number;
        enabled: boolean;
        id: number;
        macAddress: string;
        mode: {[key: string]: string};
        mtu: number;
        name: string;
        tagIds: number[];
        taggedVlans: outputs.dcim.GetDeviceInterfacesInterfaceTaggedVlan[];
        untaggedVlans: outputs.dcim.GetDeviceInterfacesInterfaceUntaggedVlan[];
    }

    export interface GetDeviceInterfacesInterfaceTaggedVlan {
        id: number;
        name: string;
        vid: number;
    }

    export interface GetDeviceInterfacesInterfaceUntaggedVlan {
        id: number;
        name: string;
        vid: number;
    }

    export interface GetDevicesDevice {
        assetTag: string;
        clusterId: number;
        comments: string;
        customFields: {[key: string]: any};
        description: string;
        deviceId: number;
        deviceTypeId: number;
        locationId: number;
        manufacturerId: number;
        model: string;
        name: string;
        platformId: number;
        primaryIpv4: string;
        primaryIpv6: string;
        rackFace: string;
        rackId: number;
        rackPosition: number;
        roleId: number;
        serial: string;
        siteId: number;
        status: string;
        tags: string[];
        tenantId: number;
    }

    export interface GetDevicesFilter {
        name: string;
        value: string;
    }

    export interface GetLocationsFilter {
        /**
         * The name of the field to filter on. Supported fields are: .
         */
        name: string;
        /**
         * The value to pass to the specified filter.
         */
        value: string;
    }

    export interface GetLocationsLocation {
        description: string;
        id: string;
        name?: string;
        siteId: number;
        slug?: string;
        status: string;
        tenantId: number;
    }

    export interface GetRacksFilter {
        name: string;
        value: string;
    }

    export interface GetRacksRack {
        assetTag: string;
        comments: string;
        customFields: {[key: string]: any};
        descUnits: boolean;
        description: string;
        facilityId: string;
        id: number;
        locationId: number;
        maxWeight: number;
        mountingDepth: number;
        name: string;
        outerDepth: number;
        outerUnit: string;
        outerWidth: number;
        roleId: number;
        serial: string;
        siteId: number;
        status: string;
        tags: string[];
        tenantId: number;
        type: string;
        uHeight: number;
        weight: number;
        weightUnit: string;
        width: number;
    }

    export interface GetRegionFilter {
        /**
         * The ID of this resource.
         */
        id?: number;
        name?: string;
        slug?: string;
    }

}

export namespace ipam {
    export interface GetAsnsAsn {
        asn: number;
        id: number;
        rirId: number;
        tags: string[];
    }

    export interface GetAsnsFilter {
        name: string;
        value: string;
    }

    export interface GetAvailablePrefixPrefixesAvailable {
        family: number;
        prefix: string;
        vrfId: number;
    }

    export interface GetIpAddressesFilter {
        name: string;
        value: string;
    }

    export interface GetIpAddressesIpAddress {
        addressFamily: string;
        created: string;
        customFields: {[key: string]: any};
        description: string;
        dnsName: string;
        id: number;
        ipAddress: string;
        lastUpdated: string;
        role: string;
        status: string;
        tags: outputs.ipam.GetIpAddressesIpAddressTag[];
        tenants: outputs.ipam.GetIpAddressesIpAddressTenant[];
    }

    export interface GetIpAddressesIpAddressTag {
        display: string;
        id: number;
        name: string;
        slug: string;
    }

    export interface GetIpAddressesIpAddressTenant {
        id: number;
        name: string;
        slug: string;
    }

    export interface GetPrefixesFilter {
        /**
         * The name of the field to filter on. Supported fields are: `prefix`, `vlanVid`, `vrfId`, `vlanId`, `status`, `siteId`, & `tag`.
         */
        name: string;
        /**
         * The value to pass to the specified filter.
         */
        value: string;
    }

    export interface GetPrefixesPrefix {
        description: string;
        id: number;
        prefix: string;
        status: string;
        tags: string[];
        vlanId: number;
        vlanVid: number;
        vrfId: number;
    }

    export interface GetVlansFilter {
        name: string;
        value: string;
    }

    export interface GetVlansVlan {
        description: string;
        groupId: number;
        name: string;
        role: number;
        site: number;
        status: string;
        tenant: number;
        vid: number;
    }

    export interface GetVrfsFilter {
        name: string;
        value: string;
    }

    export interface GetVrfsVrf {
        description: string;
        id: number;
        name: string;
        rd: string;
        tenant: number;
    }

    export interface IpAddressNatOutsideAddress {
        addressFamily: number;
        id: number;
        ipAddress: string;
    }

}

export namespace tenancy {
    export interface GetTenantsFilter {
        name: string;
        value: string;
    }

    export interface GetTenantsTenant {
        circuitCount: number;
        clusterCount: number;
        comments: string;
        created: string;
        customFields: {[key: string]: any};
        description: string;
        deviceCount: number;
        id: number;
        ipAddressCount: number;
        lastUpdated: string;
        name: string;
        prefixCount: number;
        rackCount: number;
        siteCount: number;
        slug: string;
        tenantGroups: outputs.tenancy.GetTenantsTenantTenantGroup[];
        vlanCount: number;
        vmCount: number;
        vrfCount: number;
    }

    export interface GetTenantsTenantTenantGroup {
        id: number;
        name: string;
        slug: string;
        tenantCount: number;
    }

}

export namespace virt {
    export interface GetInterfacesFilter {
        name: string;
        value: string;
    }

    export interface GetInterfacesInterface {
        description: string;
        enabled: boolean;
        id: number;
        macAddress: string;
        mode: {[key: string]: string};
        mtu: number;
        name: string;
        tagIds: number[];
        taggedVlans: outputs.virt.GetInterfacesInterfaceTaggedVlan[];
        untaggedVlans: outputs.virt.GetInterfacesInterfaceUntaggedVlan[];
        vmId: number;
    }

    export interface GetInterfacesInterfaceTaggedVlan {
        id: number;
        name: string;
        vid: number;
    }

    export interface GetInterfacesInterfaceUntaggedVlan {
        id: number;
        name: string;
        vid: number;
    }

    export interface GetVirtualMachinesFilter {
        name: string;
        value: string;
    }

    export interface GetVirtualMachinesVm {
        clusterId: number;
        comments: string;
        configContext: string;
        customFields: {[key: string]: any};
        description: string;
        deviceId: number;
        deviceName: string;
        diskSizeGb: number;
        localContextData: string;
        memoryMb: number;
        name: string;
        platformId: number;
        primaryIp: string;
        primaryIp4: string;
        primaryIp6: string;
        roleId: number;
        siteId: number;
        status: string;
        tagIds: number[];
        tenantId: number;
        vcpus: number;
        vmId: number;
    }

}