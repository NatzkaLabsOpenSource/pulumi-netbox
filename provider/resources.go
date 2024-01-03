// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netbox

import (
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/provider/pkg/version"
	"github.com/e-breuninger/terraform-provider-netbox/netbox"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
)

//go:embed cmd/pulumi-resource-netbox/bridge-metadata.json
var metadata []byte

// all of the token components used below.
const (
	// packages:
	mainPkg = "netbox"

	// modules:
	mainMod    = "index" // the netbox module
	authMod    = "auth"
	circuitMod = "circuit"
	dcimMod    = "dcim"
	ipamMod    = "ipam"
	tenancyMod = "tenancy"
	virtMod    = "virt"
)

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(netbox.Provider(), shimv2.WithDiffStrategy(shimv2.PlanState))

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "netbox",
		DisplayName:       "Netbox",
		Publisher:         "Natzka",
		Description:       "A Pulumi package for creating and managing Netbox cloud resources.",
		Keywords:          []string{"pulumi", "netbox", "category/infrastructure"},
		License:           "Apache-2.0",
		Homepage:          "https://www.pulumi.com",
		Repository:        "https://github.com/NatzkaLabsOpenSource/pulumi-netbox",
		GitHubOrg:         "e-breuninger",
		PluginDownloadURL: "github://api.github.com/NatzkaLabsOpenSource/pulumi-netbox",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Version:           version.Version,
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			// Auth
			"netbox_permission": {
				Tok: tfbridge.MakeResource(mainPkg, authMod, "Permission"),
			},
			"netbox_token": {
				Tok: tfbridge.MakeResource(mainPkg, authMod, "Token"),
			},
			"netbox_user": {
				Tok: tfbridge.MakeResource(mainPkg, authMod, "User"),
			},

			// Circuits
			"netbox_circuit": {
				Tok: tfbridge.MakeResource(mainPkg, circuitMod, "Circuit"),
			},
			"netbox_circuit_provider": {
				Tok: tfbridge.MakeResource(mainPkg, circuitMod, "CircuitProvider"),
			},
			"netbox_circuit_termination": {
				Tok: tfbridge.MakeResource(mainPkg, circuitMod, "CircuitTermination"),
			},
			"netbox_circuit_type": {
				Tok: tfbridge.MakeResource(mainPkg, circuitMod, "CircuitType"),
			},

			// DCIM
			"netbox_cable": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "Cable"),
			},
			"netbox_device": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "Device"),
			},
			"netbox_device_console_port": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DeviceConsolePort"),
			},
			"netbox_device_console_server_port": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DeviceConsoleServerPort"),
			},
			"netbox_device_front_port": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DeviceFrontPort"),
			},
			"netbox_device_interface": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DeviceInterface"),
			},
			"netbox_device_module_bay": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DeviceModuleBay"),
			},
			"netbox_device_power_outlet": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DevicePowerOutlet"),
			},
			"netbox_device_power_port": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DevicePowerPort"),
			},
			"netbox_device_primary_ip": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DevicePrimaryIp"),
			},
			"netbox_device_rear_port": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DeviceRearPort"),
			},
			"netbox_device_role": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DeviceRole"),
			},
			"netbox_device_type": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "DeviceType"),
			},
			"netbox_inventory_item": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "InventoryItem"),
			},
			"netbox_inventory_item_role": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "InventoryItemRole"),
			},
			"netbox_location": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "Location"),
			},
			"netbox_manufacturer": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "Manufacturer"),
			},
			"netbox_module": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "Module"),
			},
			"netbox_module_type": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "ModuleType"),
			},
			"netbox_platform": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "Platform"),
			},
			"netbox_power_feed": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "PowerFeed"),
			},
			"netbox_power_panel": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "PowerPanel"),
			},
			"netbox_rack": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "Rack"),
			},
			"netbox_rack_reservation": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "RackReservation"),
			},
			"netbox_rack_role": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "RackRole"),
			},
			"netbox_region": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "Region"),
			},
			"netbox_site": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "Site"),
			},
			"netbox_site_group": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "SiteGroup"),
			},
			"netbox_virtual_chassis": {
				Tok: tfbridge.MakeResource(mainPkg, dcimMod, "VirtualChassis"),
			},

			// IPAM
			"netbox_aggregate": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "Aggregate"),
			},
			"netbox_asn": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "Asn"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"asn": {
						CSharpName: "AsNumber",
					},
				},
			},
			"netbox_available_ip_address": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "AvailableIpAddress"),
			},
			"netbox_available_prefix": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "AvailablePrefix"),
			},
			"netbox_ip_address": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "IpAddress"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ip_address": {
						CSharpName: "Address",
					},
				},
			},
			"netbox_ip_range": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "IpRange"),
			},
			"netbox_ipam_role": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "Role"),
			},
			"netbox_prefix": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "Prefix"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"prefix": {
						CSharpName: "Address",
					},
				},
			},
			"netbox_rir": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "Rir"),
			},
			"netbox_route_target": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "RouteTarget"),
			},
			"netbox_service": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "Service"),
			},
			"netbox_vlan": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "Vlan"),
			},
			"netbox_vlan_group": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "VlanGroup"),
			},
			"netbox_vrf": {
				Tok: tfbridge.MakeResource(mainPkg, ipamMod, "Vrf"),
			},

			// Tenancy
			"netbox_contact": {
				Tok: tfbridge.MakeResource(mainPkg, tenancyMod, "Contact"),
			},
			"netbox_contact_assignment": {
				Tok: tfbridge.MakeResource(mainPkg, tenancyMod, "ContactAssignment"),
			},
			"netbox_contact_group": {
				Tok: tfbridge.MakeResource(mainPkg, tenancyMod, "ContactGroup"),
			},
			"netbox_contact_role": {
				Tok: tfbridge.MakeResource(mainPkg, tenancyMod, "ContactRole"),
			},
			"netbox_tenant": {
				Tok: tfbridge.MakeResource(mainPkg, tenancyMod, "Tenant"),
			},
			"netbox_tenant_group": {
				Tok: tfbridge.MakeResource(mainPkg, tenancyMod, "TenantGroup"),
			},

			// Virtualization
			"netbox_cluster": {
				Tok: tfbridge.MakeResource(mainPkg, virtMod, "Cluster"),
			},
			"netbox_cluster_group": {
				Tok: tfbridge.MakeResource(mainPkg, virtMod, "ClusterGroup"),
			},
			"netbox_cluster_type": {
				Tok: tfbridge.MakeResource(mainPkg, virtMod, "ClusterType"),
			},
			"netbox_interface": {
				Tok: tfbridge.MakeResource(mainPkg, virtMod, "Interface"),
			},
			"netbox_primary_ip": {
				Tok: tfbridge.MakeResource(mainPkg, virtMod, "PrimaryIp"),
			},
			"netbox_virtual_machine": {
				Tok: tfbridge.MakeResource(mainPkg, virtMod, "VirtualMachine"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// DCIM
			"netbox_device_interfaces": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getDeviceInterfaces"),
			},
			"netbox_device_role": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getDeviceRole"),
			},
			"netbox_device_type": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getDeviceType"),
			},
			"netbox_devices": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getDevices"),
			},
			"netbox_location": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getLocation"),
			},
			"netbox_locations": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getLocations"),
			},
			"netbox_platform": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getPlatform"),
			},
			"netbox_racks": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getRacks"),
			},
			"netbox_rack_role": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getRackRole"),
			},
			"netbox_region": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getRegion"),
			},
			"netbox_site": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getSite"),
			},
			"netbox_site_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, dcimMod, "getSiteGroup"),
			},

			// IPAM
			"netbox_asn": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getAsn"),
			},
			"netbox_asns": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getAsns"),
			},
			"netbox_available_prefix": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getAvailablePrefix"),
			},
			"netbox_ip_addresses": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getIpAddresses"),
			},
			"netbox_ip_range": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getIpRange"),
			},
			"netbox_ipam_role": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getRole"),
			},
			"netbox_prefix": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getPrefix"),
			},
			"netbox_prefixes": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getPrefixes"),
			},
			"netbox_route_target": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getRouteTarget"),
			},
			"netbox_vlan": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getVlan"),
			},
			"netbox_vlan_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getVlanGroup"),
			},
			"netbox_vlans": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getVlans"),
			},
			"netbox_vrf": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getVrf"),
			},
			"netbox_vrfs": {
				Tok: tfbridge.MakeDataSource(mainPkg, ipamMod, "getVrfs"),
			},

			// Tenancy
			"netbox_contact": {
				Tok: tfbridge.MakeDataSource(mainPkg, tenancyMod, "getContact"),
			},
			"netbox_contact_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, tenancyMod, "getContactGroup"),
			},
			"netbox_contact_role": {
				Tok: tfbridge.MakeDataSource(mainPkg, tenancyMod, "getContactRole"),
			},
			"netbox_tenant": {
				Tok: tfbridge.MakeDataSource(mainPkg, tenancyMod, "getTenant"),
			},
			"netbox_tenant_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, tenancyMod, "getTenantGroup"),
			},
			"netbox_tenants": {
				Tok: tfbridge.MakeDataSource(mainPkg, tenancyMod, "getTenants"),
			},

			// Virtualization
			"netbox_cluster": {
				Tok: tfbridge.MakeDataSource(mainPkg, virtMod, "getCluster"),
			},
			"netbox_cluster_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, virtMod, "getClusterGroup"),
			},
			"netbox_cluster_type": {
				Tok: tfbridge.MakeDataSource(mainPkg, virtMod, "getClusterType"),
			},
			"netbox_interfaces": {
				Tok: tfbridge.MakeDataSource(mainPkg, virtMod, "getInterfaces"),
			},
			"netbox_virtual_machines": {
				Tok: tfbridge.MakeDataSource(mainPkg, virtMod, "getVirtualMachines"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@natzka-oss/pulumi-netbox",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/NatzkaLabsOpenSource/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource tokens,
	// and apply auto aliasing for full backwards compatibility.
	// For more information, please reference: https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("netbox_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
