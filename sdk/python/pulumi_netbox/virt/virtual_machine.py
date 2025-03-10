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

__all__ = ['VirtualMachineArgs', 'VirtualMachine']

@pulumi.input_type
class VirtualMachineArgs:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[int]] = None,
                 disk_size_mb: Optional[pulumi.Input[int]] = None,
                 local_context_data: Optional[pulumi.Input[str]] = None,
                 memory_mb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_id: Optional[pulumi.Input[int]] = None,
                 role_id: Optional[pulumi.Input[int]] = None,
                 site_id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[int]] = None,
                 vcpus: Optional[pulumi.Input[float]] = None):
        """
        The set of arguments for constructing a VirtualMachine resource.
        :param pulumi.Input[int] cluster_id: At least one of `site_id` or `cluster_id` must be given.
        :param pulumi.Input[str] local_context_data: This is best managed through the use of `jsonencode` and a map of settings.
        :param pulumi.Input[int] site_id: At least one of `site_id` or `cluster_id` must be given.
        :param pulumi.Input[str] status: Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if custom_fields is not None:
            pulumi.set(__self__, "custom_fields", custom_fields)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if disk_size_mb is not None:
            pulumi.set(__self__, "disk_size_mb", disk_size_mb)
        if local_context_data is not None:
            pulumi.set(__self__, "local_context_data", local_context_data)
        if memory_mb is not None:
            pulumi.set(__self__, "memory_mb", memory_mb)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_id is not None:
            pulumi.set(__self__, "platform_id", platform_id)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)
        if site_id is not None:
            pulumi.set(__self__, "site_id", site_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if vcpus is not None:
            pulumi.set(__self__, "vcpus", vcpus)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[int]]:
        """
        At least one of `site_id` or `cluster_id` must be given.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="customFields")
    def custom_fields(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_fields")

    @custom_fields.setter
    def custom_fields(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_fields", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="diskSizeMb")
    def disk_size_mb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "disk_size_mb")

    @disk_size_mb.setter
    def disk_size_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_size_mb", value)

    @property
    @pulumi.getter(name="localContextData")
    def local_context_data(self) -> Optional[pulumi.Input[str]]:
        """
        This is best managed through the use of `jsonencode` and a map of settings.
        """
        return pulumi.get(self, "local_context_data")

    @local_context_data.setter
    def local_context_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_context_data", value)

    @property
    @pulumi.getter(name="memoryMb")
    def memory_mb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "memory_mb")

    @memory_mb.setter
    def memory_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_mb", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="platformId")
    def platform_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "platform_id")

    @platform_id.setter
    def platform_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "platform_id", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> Optional[pulumi.Input[int]]:
        """
        At least one of `site_id` or `cluster_id` must be given.
        """
        return pulumi.get(self, "site_id")

    @site_id.setter
    def site_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "site_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
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
    @pulumi.getter
    def vcpus(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "vcpus")

    @vcpus.setter
    def vcpus(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "vcpus", value)


@pulumi.input_type
class _VirtualMachineState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[int]] = None,
                 disk_size_mb: Optional[pulumi.Input[int]] = None,
                 local_context_data: Optional[pulumi.Input[str]] = None,
                 memory_mb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_id: Optional[pulumi.Input[int]] = None,
                 primary_ipv4: Optional[pulumi.Input[int]] = None,
                 primary_ipv6: Optional[pulumi.Input[int]] = None,
                 role_id: Optional[pulumi.Input[int]] = None,
                 site_id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[int]] = None,
                 vcpus: Optional[pulumi.Input[float]] = None):
        """
        Input properties used for looking up and filtering VirtualMachine resources.
        :param pulumi.Input[int] cluster_id: At least one of `site_id` or `cluster_id` must be given.
        :param pulumi.Input[str] local_context_data: This is best managed through the use of `jsonencode` and a map of settings.
        :param pulumi.Input[int] site_id: At least one of `site_id` or `cluster_id` must be given.
        :param pulumi.Input[str] status: Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if custom_fields is not None:
            pulumi.set(__self__, "custom_fields", custom_fields)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if disk_size_mb is not None:
            pulumi.set(__self__, "disk_size_mb", disk_size_mb)
        if local_context_data is not None:
            pulumi.set(__self__, "local_context_data", local_context_data)
        if memory_mb is not None:
            pulumi.set(__self__, "memory_mb", memory_mb)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_id is not None:
            pulumi.set(__self__, "platform_id", platform_id)
        if primary_ipv4 is not None:
            pulumi.set(__self__, "primary_ipv4", primary_ipv4)
        if primary_ipv6 is not None:
            pulumi.set(__self__, "primary_ipv6", primary_ipv6)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)
        if site_id is not None:
            pulumi.set(__self__, "site_id", site_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if vcpus is not None:
            pulumi.set(__self__, "vcpus", vcpus)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[int]]:
        """
        At least one of `site_id` or `cluster_id` must be given.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="customFields")
    def custom_fields(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_fields")

    @custom_fields.setter
    def custom_fields(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_fields", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="diskSizeMb")
    def disk_size_mb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "disk_size_mb")

    @disk_size_mb.setter
    def disk_size_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_size_mb", value)

    @property
    @pulumi.getter(name="localContextData")
    def local_context_data(self) -> Optional[pulumi.Input[str]]:
        """
        This is best managed through the use of `jsonencode` and a map of settings.
        """
        return pulumi.get(self, "local_context_data")

    @local_context_data.setter
    def local_context_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_context_data", value)

    @property
    @pulumi.getter(name="memoryMb")
    def memory_mb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "memory_mb")

    @memory_mb.setter
    def memory_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_mb", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="platformId")
    def platform_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "platform_id")

    @platform_id.setter
    def platform_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "platform_id", value)

    @property
    @pulumi.getter(name="primaryIpv4")
    def primary_ipv4(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "primary_ipv4")

    @primary_ipv4.setter
    def primary_ipv4(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "primary_ipv4", value)

    @property
    @pulumi.getter(name="primaryIpv6")
    def primary_ipv6(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "primary_ipv6")

    @primary_ipv6.setter
    def primary_ipv6(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "primary_ipv6", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> Optional[pulumi.Input[int]]:
        """
        At least one of `site_id` or `cluster_id` must be given.
        """
        return pulumi.get(self, "site_id")

    @site_id.setter
    def site_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "site_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
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
    @pulumi.getter
    def vcpus(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "vcpus")

    @vcpus.setter
    def vcpus(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "vcpus", value)


class VirtualMachine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[int]] = None,
                 disk_size_mb: Optional[pulumi.Input[int]] = None,
                 local_context_data: Optional[pulumi.Input[str]] = None,
                 memory_mb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_id: Optional[pulumi.Input[int]] = None,
                 role_id: Optional[pulumi.Input[int]] = None,
                 site_id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[int]] = None,
                 vcpus: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#virtual-machines):

        > A virtual machine is a virtualized compute instance. These behave in NetBox very similarly to device objects, but without any physical attributes. For example, a VM may have interfaces assigned to it with IP addresses and VLANs, however its interfaces cannot be connected via cables (because they are virtual). Each VM may also define its compute, memory, and storage resources as well.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cluster_id: At least one of `site_id` or `cluster_id` must be given.
        :param pulumi.Input[str] local_context_data: This is best managed through the use of `jsonencode` and a map of settings.
        :param pulumi.Input[int] site_id: At least one of `site_id` or `cluster_id` must be given.
        :param pulumi.Input[str] status: Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VirtualMachineArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#virtual-machines):

        > A virtual machine is a virtualized compute instance. These behave in NetBox very similarly to device objects, but without any physical attributes. For example, a VM may have interfaces assigned to it with IP addresses and VLANs, however its interfaces cannot be connected via cables (because they are virtual). Each VM may also define its compute, memory, and storage resources as well.

        :param str resource_name: The name of the resource.
        :param VirtualMachineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualMachineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[int]] = None,
                 disk_size_mb: Optional[pulumi.Input[int]] = None,
                 local_context_data: Optional[pulumi.Input[str]] = None,
                 memory_mb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_id: Optional[pulumi.Input[int]] = None,
                 role_id: Optional[pulumi.Input[int]] = None,
                 site_id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[int]] = None,
                 vcpus: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualMachineArgs.__new__(VirtualMachineArgs)

            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["comments"] = comments
            __props__.__dict__["custom_fields"] = custom_fields
            __props__.__dict__["description"] = description
            __props__.__dict__["device_id"] = device_id
            __props__.__dict__["disk_size_mb"] = disk_size_mb
            __props__.__dict__["local_context_data"] = local_context_data
            __props__.__dict__["memory_mb"] = memory_mb
            __props__.__dict__["name"] = name
            __props__.__dict__["platform_id"] = platform_id
            __props__.__dict__["role_id"] = role_id
            __props__.__dict__["site_id"] = site_id
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["vcpus"] = vcpus
            __props__.__dict__["primary_ipv4"] = None
            __props__.__dict__["primary_ipv6"] = None
        super(VirtualMachine, __self__).__init__(
            'netbox:virt/virtualMachine:VirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[int]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            device_id: Optional[pulumi.Input[int]] = None,
            disk_size_mb: Optional[pulumi.Input[int]] = None,
            local_context_data: Optional[pulumi.Input[str]] = None,
            memory_mb: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            platform_id: Optional[pulumi.Input[int]] = None,
            primary_ipv4: Optional[pulumi.Input[int]] = None,
            primary_ipv6: Optional[pulumi.Input[int]] = None,
            role_id: Optional[pulumi.Input[int]] = None,
            site_id: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tenant_id: Optional[pulumi.Input[int]] = None,
            vcpus: Optional[pulumi.Input[float]] = None) -> 'VirtualMachine':
        """
        Get an existing VirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cluster_id: At least one of `site_id` or `cluster_id` must be given.
        :param pulumi.Input[str] local_context_data: This is best managed through the use of `jsonencode` and a map of settings.
        :param pulumi.Input[int] site_id: At least one of `site_id` or `cluster_id` must be given.
        :param pulumi.Input[str] status: Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualMachineState.__new__(_VirtualMachineState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["comments"] = comments
        __props__.__dict__["custom_fields"] = custom_fields
        __props__.__dict__["description"] = description
        __props__.__dict__["device_id"] = device_id
        __props__.__dict__["disk_size_mb"] = disk_size_mb
        __props__.__dict__["local_context_data"] = local_context_data
        __props__.__dict__["memory_mb"] = memory_mb
        __props__.__dict__["name"] = name
        __props__.__dict__["platform_id"] = platform_id
        __props__.__dict__["primary_ipv4"] = primary_ipv4
        __props__.__dict__["primary_ipv6"] = primary_ipv6
        __props__.__dict__["role_id"] = role_id
        __props__.__dict__["site_id"] = site_id
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["vcpus"] = vcpus
        return VirtualMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[Optional[int]]:
        """
        At least one of `site_id` or `cluster_id` must be given.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="customFields")
    def custom_fields(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "custom_fields")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="diskSizeMb")
    def disk_size_mb(self) -> pulumi.Output[int]:
        return pulumi.get(self, "disk_size_mb")

    @property
    @pulumi.getter(name="localContextData")
    def local_context_data(self) -> pulumi.Output[Optional[str]]:
        """
        This is best managed through the use of `jsonencode` and a map of settings.
        """
        return pulumi.get(self, "local_context_data")

    @property
    @pulumi.getter(name="memoryMb")
    def memory_mb(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "memory_mb")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformId")
    def platform_id(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "platform_id")

    @property
    @pulumi.getter(name="primaryIpv4")
    def primary_ipv4(self) -> pulumi.Output[int]:
        return pulumi.get(self, "primary_ipv4")

    @property
    @pulumi.getter(name="primaryIpv6")
    def primary_ipv6(self) -> pulumi.Output[int]:
        return pulumi.get(self, "primary_ipv6")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> pulumi.Output[Optional[int]]:
        """
        At least one of `site_id` or `cluster_id` must be given.
        """
        return pulumi.get(self, "site_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
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
    @pulumi.getter
    def vcpus(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "vcpus")

