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

__all__ = [
    'GetRackRoleResult',
    'AwaitableGetRackRoleResult',
    'get_rack_role',
    'get_rack_role_output',
]

@pulumi.output_type
class GetRackRoleResult:
    """
    A collection of values returned by getRackRole.
    """
    def __init__(__self__, color_hex=None, description=None, id=None, name=None, slug=None, tags=None):
        if color_hex and not isinstance(color_hex, str):
            raise TypeError("Expected argument 'color_hex' to be a str")
        pulumi.set(__self__, "color_hex", color_hex)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if slug and not isinstance(slug, str):
            raise TypeError("Expected argument 'slug' to be a str")
        pulumi.set(__self__, "slug", slug)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="colorHex")
    def color_hex(self) -> str:
        return pulumi.get(self, "color_hex")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
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
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")


class AwaitableGetRackRoleResult(GetRackRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRackRoleResult(
            color_hex=self.color_hex,
            description=self.description,
            id=self.id,
            name=self.name,
            slug=self.slug,
            tags=self.tags)


def get_rack_role(name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRackRoleResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('netbox:dcim/getRackRole:getRackRole', __args__, opts=opts, typ=GetRackRoleResult).value

    return AwaitableGetRackRoleResult(
        color_hex=pulumi.get(__ret__, 'color_hex'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        slug=pulumi.get(__ret__, 'slug'),
        tags=pulumi.get(__ret__, 'tags'))
def get_rack_role_output(name: Optional[pulumi.Input[str]] = None,
                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRackRoleResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('netbox:dcim/getRackRole:getRackRole', __args__, opts=opts, typ=GetRackRoleResult)
    return __ret__.apply(lambda __response__: GetRackRoleResult(
        color_hex=pulumi.get(__response__, 'color_hex'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        slug=pulumi.get(__response__, 'slug'),
        tags=pulumi.get(__response__, 'tags')))
