# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs

__all__ = [
    'GetTenantsFilterResult',
    'GetTenantsTenantResult',
    'GetTenantsTenantTenantGroupResult',
]

@pulumi.output_type
class GetTenantsFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class GetTenantsTenantResult(dict):
    def __init__(__self__, *,
                 circuit_count: int,
                 cluster_count: int,
                 comments: str,
                 created: str,
                 custom_fields: Mapping[str, str],
                 description: str,
                 device_count: int,
                 id: int,
                 ip_address_count: int,
                 last_updated: str,
                 name: str,
                 prefix_count: int,
                 rack_count: int,
                 site_count: int,
                 slug: str,
                 tenant_groups: Sequence['outputs.GetTenantsTenantTenantGroupResult'],
                 vlan_count: int,
                 vm_count: int,
                 vrf_count: int):
        pulumi.set(__self__, "circuit_count", circuit_count)
        pulumi.set(__self__, "cluster_count", cluster_count)
        pulumi.set(__self__, "comments", comments)
        pulumi.set(__self__, "created", created)
        pulumi.set(__self__, "custom_fields", custom_fields)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "device_count", device_count)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip_address_count", ip_address_count)
        pulumi.set(__self__, "last_updated", last_updated)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "prefix_count", prefix_count)
        pulumi.set(__self__, "rack_count", rack_count)
        pulumi.set(__self__, "site_count", site_count)
        pulumi.set(__self__, "slug", slug)
        pulumi.set(__self__, "tenant_groups", tenant_groups)
        pulumi.set(__self__, "vlan_count", vlan_count)
        pulumi.set(__self__, "vm_count", vm_count)
        pulumi.set(__self__, "vrf_count", vrf_count)

    @property
    @pulumi.getter(name="circuitCount")
    def circuit_count(self) -> int:
        return pulumi.get(self, "circuit_count")

    @property
    @pulumi.getter(name="clusterCount")
    def cluster_count(self) -> int:
        return pulumi.get(self, "cluster_count")

    @property
    @pulumi.getter
    def comments(self) -> str:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def created(self) -> str:
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="customFields")
    def custom_fields(self) -> Mapping[str, str]:
        return pulumi.get(self, "custom_fields")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceCount")
    def device_count(self) -> int:
        return pulumi.get(self, "device_count")

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddressCount")
    def ip_address_count(self) -> int:
        return pulumi.get(self, "ip_address_count")

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> str:
        return pulumi.get(self, "last_updated")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="prefixCount")
    def prefix_count(self) -> int:
        return pulumi.get(self, "prefix_count")

    @property
    @pulumi.getter(name="rackCount")
    def rack_count(self) -> int:
        return pulumi.get(self, "rack_count")

    @property
    @pulumi.getter(name="siteCount")
    def site_count(self) -> int:
        return pulumi.get(self, "site_count")

    @property
    @pulumi.getter
    def slug(self) -> str:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter(name="tenantGroups")
    def tenant_groups(self) -> Sequence['outputs.GetTenantsTenantTenantGroupResult']:
        return pulumi.get(self, "tenant_groups")

    @property
    @pulumi.getter(name="vlanCount")
    def vlan_count(self) -> int:
        return pulumi.get(self, "vlan_count")

    @property
    @pulumi.getter(name="vmCount")
    def vm_count(self) -> int:
        return pulumi.get(self, "vm_count")

    @property
    @pulumi.getter(name="vrfCount")
    def vrf_count(self) -> int:
        return pulumi.get(self, "vrf_count")


@pulumi.output_type
class GetTenantsTenantTenantGroupResult(dict):
    def __init__(__self__, *,
                 id: int,
                 name: str,
                 slug: str,
                 tenant_count: int):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "slug", slug)
        pulumi.set(__self__, "tenant_count", tenant_count)

    @property
    @pulumi.getter
    def id(self) -> int:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def slug(self) -> str:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter(name="tenantCount")
    def tenant_count(self) -> int:
        return pulumi.get(self, "tenant_count")


