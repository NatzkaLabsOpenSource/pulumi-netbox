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
    'GetTenantResult',
    'AwaitableGetTenantResult',
    'get_tenant',
    'get_tenant_output',
]

@pulumi.output_type
class GetTenantResult:
    """
    A collection of values returned by getTenant.
    """
    def __init__(__self__, description=None, group_id=None, id=None, name=None, slug=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if group_id and not isinstance(group_id, int):
            raise TypeError("Expected argument 'group_id' to be a int")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if slug and not isinstance(slug, str):
            raise TypeError("Expected argument 'slug' to be a str")
        pulumi.set(__self__, "slug", slug)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> int:
        return pulumi.get(self, "group_id")

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
        """
        At least one of `name` or `slug` must be given.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def slug(self) -> str:
        """
        At least one of `name` or `slug` must be given.
        """
        return pulumi.get(self, "slug")


class AwaitableGetTenantResult(GetTenantResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTenantResult(
            description=self.description,
            group_id=self.group_id,
            id=self.id,
            name=self.name,
            slug=self.slug)


def get_tenant(description: Optional[str] = None,
               name: Optional[str] = None,
               slug: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTenantResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_netbox as netbox

    customer_a = netbox.tenancy.get_tenant(name="Customer A")
    ```


    :param str name: At least one of `name` or `slug` must be given.
    :param str slug: At least one of `name` or `slug` must be given.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['name'] = name
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('netbox:tenancy/getTenant:getTenant', __args__, opts=opts, typ=GetTenantResult).value

    return AwaitableGetTenantResult(
        description=pulumi.get(__ret__, 'description'),
        group_id=pulumi.get(__ret__, 'group_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        slug=pulumi.get(__ret__, 'slug'))
def get_tenant_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                      name: Optional[pulumi.Input[Optional[str]]] = None,
                      slug: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTenantResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_netbox as netbox

    customer_a = netbox.tenancy.get_tenant(name="Customer A")
    ```


    :param str name: At least one of `name` or `slug` must be given.
    :param str slug: At least one of `name` or `slug` must be given.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['name'] = name
    __args__['slug'] = slug
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('netbox:tenancy/getTenant:getTenant', __args__, opts=opts, typ=GetTenantResult)
    return __ret__.apply(lambda __response__: GetTenantResult(
        description=pulumi.get(__response__, 'description'),
        group_id=pulumi.get(__response__, 'group_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        slug=pulumi.get(__response__, 'slug')))
