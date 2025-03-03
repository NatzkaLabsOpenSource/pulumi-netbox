// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ClusterArgs, ClusterState } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { ClusterGroupArgs, ClusterGroupState } from "./clusterGroup";
export type ClusterGroup = import("./clusterGroup").ClusterGroup;
export const ClusterGroup: typeof import("./clusterGroup").ClusterGroup = null as any;
utilities.lazyLoad(exports, ["ClusterGroup"], () => require("./clusterGroup"));

export { ClusterTypeArgs, ClusterTypeState } from "./clusterType";
export type ClusterType = import("./clusterType").ClusterType;
export const ClusterType: typeof import("./clusterType").ClusterType = null as any;
utilities.lazyLoad(exports, ["ClusterType"], () => require("./clusterType"));

export { GetClusterArgs, GetClusterResult, GetClusterOutputArgs } from "./getCluster";
export const getCluster: typeof import("./getCluster").getCluster = null as any;
export const getClusterOutput: typeof import("./getCluster").getClusterOutput = null as any;
utilities.lazyLoad(exports, ["getCluster","getClusterOutput"], () => require("./getCluster"));

export { GetClusterGroupArgs, GetClusterGroupResult, GetClusterGroupOutputArgs } from "./getClusterGroup";
export const getClusterGroup: typeof import("./getClusterGroup").getClusterGroup = null as any;
export const getClusterGroupOutput: typeof import("./getClusterGroup").getClusterGroupOutput = null as any;
utilities.lazyLoad(exports, ["getClusterGroup","getClusterGroupOutput"], () => require("./getClusterGroup"));

export { GetClusterTypeArgs, GetClusterTypeResult, GetClusterTypeOutputArgs } from "./getClusterType";
export const getClusterType: typeof import("./getClusterType").getClusterType = null as any;
export const getClusterTypeOutput: typeof import("./getClusterType").getClusterTypeOutput = null as any;
utilities.lazyLoad(exports, ["getClusterType","getClusterTypeOutput"], () => require("./getClusterType"));

export { GetInterfacesArgs, GetInterfacesResult, GetInterfacesOutputArgs } from "./getInterfaces";
export const getInterfaces: typeof import("./getInterfaces").getInterfaces = null as any;
export const getInterfacesOutput: typeof import("./getInterfaces").getInterfacesOutput = null as any;
utilities.lazyLoad(exports, ["getInterfaces","getInterfacesOutput"], () => require("./getInterfaces"));

export { GetVirtualMachinesArgs, GetVirtualMachinesResult, GetVirtualMachinesOutputArgs } from "./getVirtualMachines";
export const getVirtualMachines: typeof import("./getVirtualMachines").getVirtualMachines = null as any;
export const getVirtualMachinesOutput: typeof import("./getVirtualMachines").getVirtualMachinesOutput = null as any;
utilities.lazyLoad(exports, ["getVirtualMachines","getVirtualMachinesOutput"], () => require("./getVirtualMachines"));

export { InterfaceArgs, InterfaceState } from "./interface";
export type Interface = import("./interface").Interface;
export const Interface: typeof import("./interface").Interface = null as any;
utilities.lazyLoad(exports, ["Interface"], () => require("./interface"));

export { PrimaryIpArgs, PrimaryIpState } from "./primaryIp";
export type PrimaryIp = import("./primaryIp").PrimaryIp;
export const PrimaryIp: typeof import("./primaryIp").PrimaryIp = null as any;
utilities.lazyLoad(exports, ["PrimaryIp"], () => require("./primaryIp"));

export { VirtualDiskArgs, VirtualDiskState } from "./virtualDisk";
export type VirtualDisk = import("./virtualDisk").VirtualDisk;
export const VirtualDisk: typeof import("./virtualDisk").VirtualDisk = null as any;
utilities.lazyLoad(exports, ["VirtualDisk"], () => require("./virtualDisk"));

export { VirtualMachineArgs, VirtualMachineState } from "./virtualMachine";
export type VirtualMachine = import("./virtualMachine").VirtualMachine;
export const VirtualMachine: typeof import("./virtualMachine").VirtualMachine = null as any;
utilities.lazyLoad(exports, ["VirtualMachine"], () => require("./virtualMachine"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "netbox:virt/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "netbox:virt/clusterGroup:ClusterGroup":
                return new ClusterGroup(name, <any>undefined, { urn })
            case "netbox:virt/clusterType:ClusterType":
                return new ClusterType(name, <any>undefined, { urn })
            case "netbox:virt/interface:Interface":
                return new Interface(name, <any>undefined, { urn })
            case "netbox:virt/primaryIp:PrimaryIp":
                return new PrimaryIp(name, <any>undefined, { urn })
            case "netbox:virt/virtualDisk:VirtualDisk":
                return new VirtualDisk(name, <any>undefined, { urn })
            case "netbox:virt/virtualMachine:VirtualMachine":
                return new VirtualMachine(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("netbox", "virt/cluster", _module)
pulumi.runtime.registerResourceModule("netbox", "virt/clusterGroup", _module)
pulumi.runtime.registerResourceModule("netbox", "virt/clusterType", _module)
pulumi.runtime.registerResourceModule("netbox", "virt/interface", _module)
pulumi.runtime.registerResourceModule("netbox", "virt/primaryIp", _module)
pulumi.runtime.registerResourceModule("netbox", "virt/virtualDisk", _module)
pulumi.runtime.registerResourceModule("netbox", "virt/virtualMachine", _module)
