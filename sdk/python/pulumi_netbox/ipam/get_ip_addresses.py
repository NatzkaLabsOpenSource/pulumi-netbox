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
from ._inputs import *

__all__ = [
    'GetIpAddressesResult',
    'AwaitableGetIpAddressesResult',
    'get_ip_addresses',
    'get_ip_addresses_output',
]

@pulumi.output_type
class GetIpAddressesResult:
    """
    A collection of values returned by getIpAddresses.
    """
    def __init__(__self__, filters=None, id=None, ip_addresses=None, limit=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if limit and not isinstance(limit, int):
            raise TypeError("Expected argument 'limit' to be a int")
        pulumi.set(__self__, "limit", limit)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetIpAddressesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Sequence['outputs.GetIpAddressesIpAddressResult']:
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter
    def limit(self) -> Optional[int]:
        """
        Defaults to `1000`.
        """
        return pulumi.get(self, "limit")


class AwaitableGetIpAddressesResult(GetIpAddressesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpAddressesResult(
            filters=self.filters,
            id=self.id,
            ip_addresses=self.ip_addresses,
            limit=self.limit)


def get_ip_addresses(filters: Optional[Sequence[Union['GetIpAddressesFilterArgs', 'GetIpAddressesFilterArgsDict']]] = None,
                     limit: Optional[int] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpAddressesResult:
    """
    Use this data source to access information about an existing resource.

    :param int limit: Defaults to `1000`.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['limit'] = limit
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('netbox:ipam/getIpAddresses:getIpAddresses', __args__, opts=opts, typ=GetIpAddressesResult).value

    return AwaitableGetIpAddressesResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        ip_addresses=pulumi.get(__ret__, 'ip_addresses'),
        limit=pulumi.get(__ret__, 'limit'))
def get_ip_addresses_output(filters: Optional[pulumi.Input[Optional[Sequence[Union['GetIpAddressesFilterArgs', 'GetIpAddressesFilterArgsDict']]]]] = None,
                            limit: Optional[pulumi.Input[Optional[int]]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIpAddressesResult]:
    """
    Use this data source to access information about an existing resource.

    :param int limit: Defaults to `1000`.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['limit'] = limit
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('netbox:ipam/getIpAddresses:getIpAddresses', __args__, opts=opts, typ=GetIpAddressesResult)
    return __ret__.apply(lambda __response__: GetIpAddressesResult(
        filters=pulumi.get(__response__, 'filters'),
        id=pulumi.get(__response__, 'id'),
        ip_addresses=pulumi.get(__response__, 'ip_addresses'),
        limit=pulumi.get(__response__, 'limit')))
