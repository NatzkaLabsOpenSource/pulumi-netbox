// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"fmt"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "netbox:vpn/vpnTunnel:VpnTunnel":
		r = &VpnTunnel{}
	case "netbox:vpn/vpnTunnelGroup:VpnTunnelGroup":
		r = &VpnTunnelGroup{}
	case "netbox:vpn/vpnTunnelTermination:VpnTunnelTermination":
		r = &VpnTunnelTermination{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"netbox",
		"vpn/vpnTunnel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"vpn/vpnTunnelGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"vpn/vpnTunnelTermination",
		&module{version},
	)
}
