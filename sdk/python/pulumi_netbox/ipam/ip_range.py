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

__all__ = ['IpRangeArgs', 'IpRange']

@pulumi.input_type
class IpRangeArgs:
    def __init__(__self__, *,
                 end_address: pulumi.Input[str],
                 start_address: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[int]] = None,
                 vrf_id: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a IpRange resource.
        :param pulumi.Input[str] status: Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
        """
        pulumi.set(__self__, "end_address", end_address)
        pulumi.set(__self__, "start_address", start_address)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if vrf_id is not None:
            pulumi.set(__self__, "vrf_id", vrf_id)

    @property
    @pulumi.getter(name="endAddress")
    def end_address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "end_address")

    @end_address.setter
    def end_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "end_address", value)

    @property
    @pulumi.getter(name="startAddress")
    def start_address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "start_address")

    @start_address.setter
    def start_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "start_address", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="vrfId")
    def vrf_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vrf_id")

    @vrf_id.setter
    def vrf_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vrf_id", value)


@pulumi.input_type
class _IpRangeState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 end_address: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[int]] = None,
                 start_address: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[int]] = None,
                 vrf_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering IpRange resources.
        :param pulumi.Input[str] status: Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if end_address is not None:
            pulumi.set(__self__, "end_address", end_address)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)
        if start_address is not None:
            pulumi.set(__self__, "start_address", start_address)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if vrf_id is not None:
            pulumi.set(__self__, "vrf_id", vrf_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endAddress")
    def end_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "end_address")

    @end_address.setter
    def end_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_address", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter(name="startAddress")
    def start_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "start_address")

    @start_address.setter
    def start_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_address", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="vrfId")
    def vrf_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vrf_id")

    @vrf_id.setter
    def vrf_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vrf_id", value)


class IpRange(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_address: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[int]] = None,
                 start_address: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[int]] = None,
                 vrf_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#ip-ranges):

        > This model represents an arbitrary range of individual IPv4 or IPv6 addresses, inclusive of its starting and ending addresses. For instance, the range 192.0.2.10 to 192.0.2.20 has eleven members. (The total member count is available as the size property on an IPRange instance.) Like prefixes and IP addresses, each IP range may optionally be assigned to a VRF and/or tenant.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        cust_a_prod = netbox.ipam.IpRange("cust_a_prod",
            start_address="10.0.0.1/24",
            end_address="10.0.0.50/24",
            tags=[
                "customer-a",
                "prod",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] status: Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpRangeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#ip-ranges):

        > This model represents an arbitrary range of individual IPv4 or IPv6 addresses, inclusive of its starting and ending addresses. For instance, the range 192.0.2.10 to 192.0.2.20 has eleven members. (The total member count is available as the size property on an IPRange instance.) Like prefixes and IP addresses, each IP range may optionally be assigned to a VRF and/or tenant.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        cust_a_prod = netbox.ipam.IpRange("cust_a_prod",
            start_address="10.0.0.1/24",
            end_address="10.0.0.50/24",
            tags=[
                "customer-a",
                "prod",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param IpRangeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpRangeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_address: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[int]] = None,
                 start_address: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[int]] = None,
                 vrf_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpRangeArgs.__new__(IpRangeArgs)

            __props__.__dict__["description"] = description
            if end_address is None and not opts.urn:
                raise TypeError("Missing required property 'end_address'")
            __props__.__dict__["end_address"] = end_address
            __props__.__dict__["role_id"] = role_id
            if start_address is None and not opts.urn:
                raise TypeError("Missing required property 'start_address'")
            __props__.__dict__["start_address"] = start_address
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["vrf_id"] = vrf_id
        super(IpRange, __self__).__init__(
            'netbox:ipam/ipRange:IpRange',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            end_address: Optional[pulumi.Input[str]] = None,
            role_id: Optional[pulumi.Input[int]] = None,
            start_address: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tenant_id: Optional[pulumi.Input[int]] = None,
            vrf_id: Optional[pulumi.Input[int]] = None) -> 'IpRange':
        """
        Get an existing IpRange resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] status: Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpRangeState.__new__(_IpRangeState)

        __props__.__dict__["description"] = description
        __props__.__dict__["end_address"] = end_address
        __props__.__dict__["role_id"] = role_id
        __props__.__dict__["start_address"] = start_address
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["vrf_id"] = vrf_id
        return IpRange(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endAddress")
    def end_address(self) -> pulumi.Output[str]:
        return pulumi.get(self, "end_address")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter(name="startAddress")
    def start_address(self) -> pulumi.Output[str]:
        return pulumi.get(self, "start_address")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="vrfId")
    def vrf_id(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "vrf_id")

