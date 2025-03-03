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
    'GetIpRangeResult',
    'AwaitableGetIpRangeResult',
    'get_ip_range',
    'get_ip_range_output',
]

@pulumi.output_type
class GetIpRangeResult:
    """
    A collection of values returned by getIpRange.
    """
    def __init__(__self__, contains=None, id=None):
        if contains and not isinstance(contains, str):
            raise TypeError("Expected argument 'contains' to be a str")
        pulumi.set(__self__, "contains", contains)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def contains(self) -> str:
        return pulumi.get(self, "contains")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetIpRangeResult(GetIpRangeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpRangeResult(
            contains=self.contains,
            id=self.id)


def get_ip_range(contains: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpRangeResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_netbox as netbox

    cust_a_prod = netbox.ipam.get_ip_range(contains="10.0.0.1/24")
    ```
    """
    __args__ = dict()
    __args__['contains'] = contains
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('netbox:ipam/getIpRange:getIpRange', __args__, opts=opts, typ=GetIpRangeResult).value

    return AwaitableGetIpRangeResult(
        contains=pulumi.get(__ret__, 'contains'),
        id=pulumi.get(__ret__, 'id'))
def get_ip_range_output(contains: Optional[pulumi.Input[str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIpRangeResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_netbox as netbox

    cust_a_prod = netbox.ipam.get_ip_range(contains="10.0.0.1/24")
    ```
    """
    __args__ = dict()
    __args__['contains'] = contains
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('netbox:ipam/getIpRange:getIpRange', __args__, opts=opts, typ=GetIpRangeResult)
    return __ret__.apply(lambda __response__: GetIpRangeResult(
        contains=pulumi.get(__response__, 'contains'),
        id=pulumi.get(__response__, 'id')))
