// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { VpnTunnelArgs, VpnTunnelState } from "./vpnTunnel";
export type VpnTunnel = import("./vpnTunnel").VpnTunnel;
export const VpnTunnel: typeof import("./vpnTunnel").VpnTunnel = null as any;
utilities.lazyLoad(exports, ["VpnTunnel"], () => require("./vpnTunnel"));

export { VpnTunnelGroupArgs, VpnTunnelGroupState } from "./vpnTunnelGroup";
export type VpnTunnelGroup = import("./vpnTunnelGroup").VpnTunnelGroup;
export const VpnTunnelGroup: typeof import("./vpnTunnelGroup").VpnTunnelGroup = null as any;
utilities.lazyLoad(exports, ["VpnTunnelGroup"], () => require("./vpnTunnelGroup"));

export { VpnTunnelTerminationArgs, VpnTunnelTerminationState } from "./vpnTunnelTermination";
export type VpnTunnelTermination = import("./vpnTunnelTermination").VpnTunnelTermination;
export const VpnTunnelTermination: typeof import("./vpnTunnelTermination").VpnTunnelTermination = null as any;
utilities.lazyLoad(exports, ["VpnTunnelTermination"], () => require("./vpnTunnelTermination"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "netbox:vpn/vpnTunnel:VpnTunnel":
                return new VpnTunnel(name, <any>undefined, { urn })
            case "netbox:vpn/vpnTunnelGroup:VpnTunnelGroup":
                return new VpnTunnelGroup(name, <any>undefined, { urn })
            case "netbox:vpn/vpnTunnelTermination:VpnTunnelTermination":
                return new VpnTunnelTermination(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("netbox", "vpn/vpnTunnel", _module)
pulumi.runtime.registerResourceModule("netbox", "vpn/vpnTunnelGroup", _module)
pulumi.runtime.registerResourceModule("netbox", "vpn/vpnTunnelTermination", _module)
