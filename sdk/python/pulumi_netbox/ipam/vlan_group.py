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

__all__ = ['VlanGroupArgs', 'VlanGroup']

@pulumi.input_type
class VlanGroupArgs:
    def __init__(__self__, *,
                 slug: pulumi.Input[str],
                 vid_ranges: pulumi.Input[Sequence[pulumi.Input[Sequence[pulumi.Input[int]]]]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[int]] = None,
                 scope_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VlanGroup resource.
        :param pulumi.Input[str] description: Defaults to `""`.
        :param pulumi.Input[int] scope_id: Required when `scope_type` is set.
        :param pulumi.Input[str] scope_type: Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
        """
        pulumi.set(__self__, "slug", slug)
        pulumi.set(__self__, "vid_ranges", vid_ranges)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)
        if scope_type is not None:
            pulumi.set(__self__, "scope_type", scope_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Input[str]:
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: pulumi.Input[str]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter(name="vidRanges")
    def vid_ranges(self) -> pulumi.Input[Sequence[pulumi.Input[Sequence[pulumi.Input[int]]]]]:
        return pulumi.get(self, "vid_ranges")

    @vid_ranges.setter
    def vid_ranges(self, value: pulumi.Input[Sequence[pulumi.Input[Sequence[pulumi.Input[int]]]]]):
        pulumi.set(self, "vid_ranges", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `""`.
        """
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
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[int]]:
        """
        Required when `scope_type` is set.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter(name="scopeType")
    def scope_type(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
        """
        return pulumi.get(self, "scope_type")

    @scope_type.setter
    def scope_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VlanGroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[int]] = None,
                 scope_type: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vid_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[Sequence[pulumi.Input[int]]]]]] = None):
        """
        Input properties used for looking up and filtering VlanGroup resources.
        :param pulumi.Input[str] description: Defaults to `""`.
        :param pulumi.Input[int] scope_id: Required when `scope_type` is set.
        :param pulumi.Input[str] scope_type: Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)
        if scope_type is not None:
            pulumi.set(__self__, "scope_type", scope_type)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vid_ranges is not None:
            pulumi.set(__self__, "vid_ranges", vid_ranges)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to `""`.
        """
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
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[int]]:
        """
        Required when `scope_type` is set.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter(name="scopeType")
    def scope_type(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
        """
        return pulumi.get(self, "scope_type")

    @scope_type.setter
    def scope_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_type", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vidRanges")
    def vid_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Sequence[pulumi.Input[int]]]]]]:
        return pulumi.get(self, "vid_ranges")

    @vid_ranges.setter
    def vid_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Sequence[pulumi.Input[int]]]]]]):
        pulumi.set(self, "vid_ranges", value)


class VlanGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[int]] = None,
                 scope_type: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vid_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[Sequence[pulumi.Input[int]]]]]] = None,
                 __props__=None):
        """
        > A VLAN Group represents a collection of VLANs. Generally, these are limited by one of a number of scopes such as "Site" or "Virtualization Cluster".

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        #Basic VLAN Group example
        example1 = netbox.ipam.VlanGroup("example1",
            name="example1",
            slug="example1")
        #Full VLAN Group example
        example2 = netbox.ipam.VlanGroup("example2",
            name="Second Example",
            slug="example2",
            scope_type="dcim.site",
            scope_id=example["id"],
            description="Second Example VLAN Group",
            tags=[example_netbox_tag["id"]],
            vid_ranges=[
                [
                    1,
                    2,
                ],
                [
                    3,
                    4,
                ],
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Defaults to `""`.
        :param pulumi.Input[int] scope_id: Required when `scope_type` is set.
        :param pulumi.Input[str] scope_type: Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VlanGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        > A VLAN Group represents a collection of VLANs. Generally, these are limited by one of a number of scopes such as "Site" or "Virtualization Cluster".

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        #Basic VLAN Group example
        example1 = netbox.ipam.VlanGroup("example1",
            name="example1",
            slug="example1")
        #Full VLAN Group example
        example2 = netbox.ipam.VlanGroup("example2",
            name="Second Example",
            slug="example2",
            scope_type="dcim.site",
            scope_id=example["id"],
            description="Second Example VLAN Group",
            tags=[example_netbox_tag["id"]],
            vid_ranges=[
                [
                    1,
                    2,
                ],
                [
                    3,
                    4,
                ],
            ])
        ```

        :param str resource_name: The name of the resource.
        :param VlanGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VlanGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[int]] = None,
                 scope_type: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vid_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[Sequence[pulumi.Input[int]]]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VlanGroupArgs.__new__(VlanGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["scope_id"] = scope_id
            __props__.__dict__["scope_type"] = scope_type
            if slug is None and not opts.urn:
                raise TypeError("Missing required property 'slug'")
            __props__.__dict__["slug"] = slug
            __props__.__dict__["tags"] = tags
            if vid_ranges is None and not opts.urn:
                raise TypeError("Missing required property 'vid_ranges'")
            __props__.__dict__["vid_ranges"] = vid_ranges
        super(VlanGroup, __self__).__init__(
            'netbox:ipam/vlanGroup:VlanGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            scope_id: Optional[pulumi.Input[int]] = None,
            scope_type: Optional[pulumi.Input[str]] = None,
            slug: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            vid_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[Sequence[pulumi.Input[int]]]]]] = None) -> 'VlanGroup':
        """
        Get an existing VlanGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Defaults to `""`.
        :param pulumi.Input[int] scope_id: Required when `scope_type` is set.
        :param pulumi.Input[str] scope_type: Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VlanGroupState.__new__(_VlanGroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["scope_id"] = scope_id
        __props__.__dict__["scope_type"] = scope_type
        __props__.__dict__["slug"] = slug
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vid_ranges"] = vid_ranges
        return VlanGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to `""`.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[Optional[int]]:
        """
        Required when `scope_type` is set.
        """
        return pulumi.get(self, "scope_id")

    @property
    @pulumi.getter(name="scopeType")
    def scope_type(self) -> pulumi.Output[Optional[str]]:
        """
        Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
        """
        return pulumi.get(self, "scope_type")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vidRanges")
    def vid_ranges(self) -> pulumi.Output[Sequence[Sequence[int]]]:
        return pulumi.get(self, "vid_ranges")

