// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

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
	case "netbox:ipam/aggregate:Aggregate":
		r = &Aggregate{}
	case "netbox:ipam/asn:Asn":
		r = &Asn{}
	case "netbox:ipam/availableIpAddress:AvailableIpAddress":
		r = &AvailableIpAddress{}
	case "netbox:ipam/availablePrefix:AvailablePrefix":
		r = &AvailablePrefix{}
	case "netbox:ipam/ipAddress:IpAddress":
		r = &IpAddress{}
	case "netbox:ipam/ipRange:IpRange":
		r = &IpRange{}
	case "netbox:ipam/prefix:Prefix":
		r = &Prefix{}
	case "netbox:ipam/rir:Rir":
		r = &Rir{}
	case "netbox:ipam/role:Role":
		r = &Role{}
	case "netbox:ipam/routeTarget:RouteTarget":
		r = &RouteTarget{}
	case "netbox:ipam/service:Service":
		r = &Service{}
	case "netbox:ipam/vlan:Vlan":
		r = &Vlan{}
	case "netbox:ipam/vlanGroup:VlanGroup":
		r = &VlanGroup{}
	case "netbox:ipam/vrf:Vrf":
		r = &Vrf{}
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
		"ipam/aggregate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/asn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/availableIpAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/availablePrefix",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/ipAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/ipRange",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/prefix",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/rir",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/routeTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/service",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/vlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/vlanGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"netbox",
		"ipam/vrf",
		&module{version},
	)
}
