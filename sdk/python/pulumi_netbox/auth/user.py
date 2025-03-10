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

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 username: pulumi.Input[str],
                 active: Optional[pulumi.Input[bool]] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 staff: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[bool] active: Defaults to `true`.
        :param pulumi.Input[bool] staff: Defaults to `false`.
        """
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)
        if active is not None:
            pulumi.set(__self__, "active", active)
        if group_ids is not None:
            pulumi.set(__self__, "group_ids", group_ids)
        if staff is not None:
            pulumi.set(__self__, "staff", staff)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        return pulumi.get(self, "group_ids")

    @group_ids.setter
    def group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "group_ids", value)

    @property
    @pulumi.getter
    def staff(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "staff")

    @staff.setter
    def staff(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "staff", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 staff: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[bool] active: Defaults to `true`.
        :param pulumi.Input[bool] staff: Defaults to `false`.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if group_ids is not None:
            pulumi.set(__self__, "group_ids", group_ids)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if staff is not None:
            pulumi.set(__self__, "staff", staff)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        return pulumi.get(self, "group_ids")

    @group_ids.setter
    def group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "group_ids", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def staff(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "staff")

    @staff.setter
    def staff(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "staff", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 staff: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource is used to manage users.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        test = netbox.auth.User("test",
            username="johndoe",
            password="Abcdefghijkl1",
            active=True,
            staff=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Defaults to `true`.
        :param pulumi.Input[bool] staff: Defaults to `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource is used to manage users.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        test = netbox.auth.User("test",
            username="johndoe",
            password="Abcdefghijkl1",
            active=True,
            staff=True)
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 staff: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["active"] = active
            __props__.__dict__["group_ids"] = group_ids
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["staff"] = staff
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(User, __self__).__init__(
            'netbox:auth/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            password: Optional[pulumi.Input[str]] = None,
            staff: Optional[pulumi.Input[bool]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Defaults to `true`.
        :param pulumi.Input[bool] staff: Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["active"] = active
        __props__.__dict__["group_ids"] = group_ids
        __props__.__dict__["password"] = password
        __props__.__dict__["staff"] = staff
        __props__.__dict__["username"] = username
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> pulumi.Output[Optional[Sequence[int]]]:
        return pulumi.get(self, "group_ids")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def staff(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`.
        """
        return pulumi.get(self, "staff")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        return pulumi.get(self, "username")

