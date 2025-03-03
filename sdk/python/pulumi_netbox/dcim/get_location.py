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
    'GetLocationResult',
    'AwaitableGetLocationResult',
    'get_location',
    'get_location_output',
]

@pulumi.output_type
class GetLocationResult:
    """
    A collection of values returned by getLocation.
    """
    def __init__(__self__, description=None, id=None, name=None, parent_id=None, site_id=None, slug=None, status=None, tenant_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_id and not isinstance(parent_id, int):
            raise TypeError("Expected argument 'parent_id' to be a int")
        pulumi.set(__self__, "parent_id", parent_id)
        if site_id and not isinstance(site_id, int):
            raise TypeError("Expected argument 'site_id' to be a int")
        pulumi.set(__self__, "site_id", site_id)
        if slug and not isinstance(slug, str):
            raise TypeError("Expected argument 'slug' to be a str")
        pulumi.set(__self__, "slug", slug)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tenant_id and not isinstance(tenant_id, int):
            raise TypeError("Expected argument 'tenant_id' to be a int")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> int:
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> int:
        return pulumi.get(self, "site_id")

    @property
    @pulumi.getter
    def slug(self) -> Optional[str]:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> int:
        return pulumi.get(self, "tenant_id")


class AwaitableGetLocationResult(GetLocationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocationResult(
            description=self.description,
            id=self.id,
            name=self.name,
            parent_id=self.parent_id,
            site_id=self.site_id,
            slug=self.slug,
            status=self.status,
            tenant_id=self.tenant_id)


def get_location(id: Optional[str] = None,
                 name: Optional[str] = None,
                 parent_id: Optional[int] = None,
                 site_id: Optional[int] = None,
                 slug: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLocationResult:
    """
    Use this data source to access information about an existing resource.

    :param str id: The ID of this resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['parentId'] = parent_id
    __args__['siteId'] = site_id
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('netbox:dcim/getLocation:getLocation', __args__, opts=opts, typ=GetLocationResult).value

    return AwaitableGetLocationResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        parent_id=pulumi.get(__ret__, 'parent_id'),
        site_id=pulumi.get(__ret__, 'site_id'),
        slug=pulumi.get(__ret__, 'slug'),
        status=pulumi.get(__ret__, 'status'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_location_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        parent_id: Optional[pulumi.Input[Optional[int]]] = None,
                        site_id: Optional[pulumi.Input[Optional[int]]] = None,
                        slug: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLocationResult]:
    """
    Use this data source to access information about an existing resource.

    :param str id: The ID of this resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['parentId'] = parent_id
    __args__['siteId'] = site_id
    __args__['slug'] = slug
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('netbox:dcim/getLocation:getLocation', __args__, opts=opts, typ=GetLocationResult)
    return __ret__.apply(lambda __response__: GetLocationResult(
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        parent_id=pulumi.get(__response__, 'parent_id'),
        site_id=pulumi.get(__response__, 'site_id'),
        slug=pulumi.get(__response__, 'slug'),
        status=pulumi.get(__response__, 'status'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
