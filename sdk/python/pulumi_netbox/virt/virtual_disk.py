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

__all__ = ['VirtualDiskArgs', 'VirtualDisk']

@pulumi.input_type
class VirtualDiskArgs:
    def __init__(__self__, *,
                 size_mb: pulumi.Input[int],
                 virtual_machine_id: pulumi.Input[int],
                 custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VirtualDisk resource.
        """
        pulumi.set(__self__, "size_mb", size_mb)
        pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)
        if custom_fields is not None:
            pulumi.set(__self__, "custom_fields", custom_fields)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="sizeMb")
    def size_mb(self) -> pulumi.Input[int]:
        return pulumi.get(self, "size_mb")

    @size_mb.setter
    def size_mb(self, value: pulumi.Input[int]):
        pulumi.set(self, "size_mb", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "virtual_machine_id", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VirtualDiskState:
    def __init__(__self__, *,
                 custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size_mb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_machine_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering VirtualDisk resources.
        """
        if custom_fields is not None:
            pulumi.set(__self__, "custom_fields", custom_fields)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if size_mb is not None:
            pulumi.set(__self__, "size_mb", size_mb)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if virtual_machine_id is not None:
            pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sizeMb")
    def size_mb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "size_mb")

    @size_mb.setter
    def size_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_mb", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "virtual_machine_id", value)


class VirtualDisk(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size_mb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_machine_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/models/virtualization/virtualdisk/):

                > A virtual disk is used to model discrete virtual hard disks assigned to virtual machines.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        # Assumes vmw-cluster-01 exists in Netbox
        vmw_cluster01 = netbox.virt.get_cluster(name="vmw-cluster-01")
        base_vm = netbox.virt.VirtualMachine("base_vm",
            cluster_id=vmw_cluster01.id,
            name="myvm-1")
        example = netbox.virt.VirtualDisk("example",
            name="disk-01",
            description="Main disk",
            size_mb=50,
            virtual_machine_id=base_vm.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualDiskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/models/virtualization/virtualdisk/):

                > A virtual disk is used to model discrete virtual hard disks assigned to virtual machines.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        # Assumes vmw-cluster-01 exists in Netbox
        vmw_cluster01 = netbox.virt.get_cluster(name="vmw-cluster-01")
        base_vm = netbox.virt.VirtualMachine("base_vm",
            cluster_id=vmw_cluster01.id,
            name="myvm-1")
        example = netbox.virt.VirtualDisk("example",
            name="disk-01",
            description="Main disk",
            size_mb=50,
            virtual_machine_id=base_vm.id)
        ```

        :param str resource_name: The name of the resource.
        :param VirtualDiskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualDiskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size_mb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_machine_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualDiskArgs.__new__(VirtualDiskArgs)

            __props__.__dict__["custom_fields"] = custom_fields
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if size_mb is None and not opts.urn:
                raise TypeError("Missing required property 'size_mb'")
            __props__.__dict__["size_mb"] = size_mb
            __props__.__dict__["tags"] = tags
            if virtual_machine_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_machine_id'")
            __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        super(VirtualDisk, __self__).__init__(
            'netbox:virt/virtualDisk:VirtualDisk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            size_mb: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            virtual_machine_id: Optional[pulumi.Input[int]] = None) -> 'VirtualDisk':
        """
        Get an existing VirtualDisk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualDiskState.__new__(_VirtualDiskState)

        __props__.__dict__["custom_fields"] = custom_fields
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["size_mb"] = size_mb
        __props__.__dict__["tags"] = tags
        __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        return VirtualDisk(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customFields")
    def custom_fields(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "custom_fields")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sizeMb")
    def size_mb(self) -> pulumi.Output[int]:
        return pulumi.get(self, "size_mb")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "virtual_machine_id")

