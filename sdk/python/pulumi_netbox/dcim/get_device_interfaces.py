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
    'GetDeviceInterfacesResult',
    'AwaitableGetDeviceInterfacesResult',
    'get_device_interfaces',
    'get_device_interfaces_output',
]

@pulumi.output_type
class GetDeviceInterfacesResult:
    """
    A collection of values returned by getDeviceInterfaces.
    """
    def __init__(__self__, filters=None, id=None, interfaces=None, name_regex=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetDeviceInterfacesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interfaces(self) -> Sequence['outputs.GetDeviceInterfacesInterfaceResult']:
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")


class AwaitableGetDeviceInterfacesResult(GetDeviceInterfacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceInterfacesResult(
            filters=self.filters,
            id=self.id,
            interfaces=self.interfaces,
            name_regex=self.name_regex)


def get_device_interfaces(filters: Optional[Sequence[Union['GetDeviceInterfacesFilterArgs', 'GetDeviceInterfacesFilterArgsDict']]] = None,
                          name_regex: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceInterfacesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['nameRegex'] = name_regex
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('netbox:dcim/getDeviceInterfaces:getDeviceInterfaces', __args__, opts=opts, typ=GetDeviceInterfacesResult).value

    return AwaitableGetDeviceInterfacesResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        interfaces=pulumi.get(__ret__, 'interfaces'),
        name_regex=pulumi.get(__ret__, 'name_regex'))
def get_device_interfaces_output(filters: Optional[pulumi.Input[Optional[Sequence[Union['GetDeviceInterfacesFilterArgs', 'GetDeviceInterfacesFilterArgsDict']]]]] = None,
                                 name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDeviceInterfacesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['nameRegex'] = name_regex
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('netbox:dcim/getDeviceInterfaces:getDeviceInterfaces', __args__, opts=opts, typ=GetDeviceInterfacesResult)
    return __ret__.apply(lambda __response__: GetDeviceInterfacesResult(
        filters=pulumi.get(__response__, 'filters'),
        id=pulumi.get(__response__, 'id'),
        interfaces=pulumi.get(__response__, 'interfaces'),
        name_regex=pulumi.get(__response__, 'name_regex')))
