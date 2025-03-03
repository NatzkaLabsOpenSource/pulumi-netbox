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
    'GetInterfacesFilterArgs',
    'GetInterfacesFilterArgsDict',
    'GetVirtualMachinesFilterArgs',
    'GetVirtualMachinesFilterArgsDict',
]

MYPY = False

if not MYPY:
    class GetInterfacesFilterArgsDict(TypedDict):
        name: str
        value: str
elif False:
    GetInterfacesFilterArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class GetInterfacesFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 value: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


if not MYPY:
    class GetVirtualMachinesFilterArgsDict(TypedDict):
        name: str
        value: str
elif False:
    GetVirtualMachinesFilterArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class GetVirtualMachinesFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 value: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


