// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetTagsFilter {
    name: string;
    value: string;
}

export interface GetTagsFilterArgs {
    name: pulumi.Input<string>;
    value: pulumi.Input<string>;
}

export namespace dcim {
    export interface CableATermination {
        objectId: pulumi.Input<number>;
        objectType: pulumi.Input<string>;
    }

    export interface CableBTermination {
        objectId: pulumi.Input<number>;
        objectType: pulumi.Input<string>;
    }

    export interface GetDeviceInterfacesFilter {
        name: string;
        value: string;
    }

    export interface GetDeviceInterfacesFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetDevicesFilter {
        name: string;
        value: string;
    }

    export interface GetDevicesFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
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

    export interface GetLocationsFilterArgs {
        /**
         * The name of the field to filter on. Supported fields are: .
         */
        name: pulumi.Input<string>;
        /**
         * The value to pass to the specified filter.
         */
        value: pulumi.Input<string>;
    }

    export interface GetRacksFilter {
        name: string;
        value: string;
    }

    export interface GetRacksFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetRegionFilter {
        /**
         * The ID of this resource.
         */
        id?: number;
        name?: string;
        slug?: string;
    }

    export interface GetRegionFilterArgs {
        /**
         * The ID of this resource.
         */
        id?: pulumi.Input<number>;
        name?: pulumi.Input<string>;
        slug?: pulumi.Input<string>;
    }
}

export namespace ipam {
    export interface GetAsnsFilter {
        name: string;
        value: string;
    }

    export interface GetAsnsFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetIpAddressesFilter {
        name: string;
        value: string;
    }

    export interface GetIpAddressesFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
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

    export interface GetPrefixesFilterArgs {
        /**
         * The name of the field to filter on. Supported fields are: `prefix`, `vlanVid`, `vrfId`, `vlanId`, `status`, `siteId`, & `tag`.
         */
        name: pulumi.Input<string>;
        /**
         * The value to pass to the specified filter.
         */
        value: pulumi.Input<string>;
    }

    export interface GetVlansFilter {
        name: string;
        value: string;
    }

    export interface GetVlansFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetVrfsFilter {
        name: string;
        value: string;
    }

    export interface GetVrfsFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface IpAddressNatOutsideAddress {
        addressFamily?: pulumi.Input<number>;
        id?: pulumi.Input<number>;
        ipAddress?: pulumi.Input<string>;
    }
}

export namespace tenancy {
    export interface GetTenantsFilter {
        name: string;
        value: string;
    }

    export interface GetTenantsFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

}

export namespace virt {
    export interface GetInterfacesFilter {
        name: string;
        value: string;
    }

    export interface GetInterfacesFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    export interface GetVirtualMachinesFilter {
        name: string;
        value: string;
    }

    export interface GetVirtualMachinesFilterArgs {
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

}