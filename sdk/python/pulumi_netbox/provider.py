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
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 api_token: pulumi.Input[str],
                 server_url: pulumi.Input[str],
                 allow_insecure_https: Optional[pulumi.Input[bool]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 request_timeout: Optional[pulumi.Input[int]] = None,
                 skip_version_check: Optional[pulumi.Input[bool]] = None,
                 strip_trailing_slashes_from_url: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] api_token: Netbox API authentication token. Can be set via the `NETBOX_API_TOKEN` environment variable.
        :param pulumi.Input[str] server_url: Location of Netbox server including scheme (http or https) and optional port. Can be set via the `NETBOX_SERVER_URL`
               environment variable.
        :param pulumi.Input[bool] allow_insecure_https: Flag to set whether to allow https with invalid certificates. Can be set via the `NETBOX_ALLOW_INSECURE_HTTPS`
               environment variable. Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: Set these header on all requests to Netbox. Can be set via the `NETBOX_HEADERS` environment variable.
        :param pulumi.Input[int] request_timeout: Netbox API HTTP request timeout in seconds. Can be set via the `NETBOX_REQUEST_TIMEOUT` environment variable.
        :param pulumi.Input[bool] strip_trailing_slashes_from_url: If true, strip trailing slashes from the `server_url` parameter and print a warning when doing so. Note that using
               trailing slashes in the `server_url` parameter will usually lead to errors. Can be set via the
               `NETBOX_STRIP_TRAILING_SLASHES_FROM_URL` environment variable. Defaults to `true`.
        """
        pulumi.set(__self__, "api_token", api_token)
        pulumi.set(__self__, "server_url", server_url)
        if allow_insecure_https is not None:
            pulumi.set(__self__, "allow_insecure_https", allow_insecure_https)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if request_timeout is not None:
            pulumi.set(__self__, "request_timeout", request_timeout)
        if skip_version_check is not None:
            pulumi.set(__self__, "skip_version_check", skip_version_check)
        if strip_trailing_slashes_from_url is not None:
            pulumi.set(__self__, "strip_trailing_slashes_from_url", strip_trailing_slashes_from_url)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Input[str]:
        """
        Netbox API authentication token. Can be set via the `NETBOX_API_TOKEN` environment variable.
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_token", value)

    @property
    @pulumi.getter(name="serverUrl")
    def server_url(self) -> pulumi.Input[str]:
        """
        Location of Netbox server including scheme (http or https) and optional port. Can be set via the `NETBOX_SERVER_URL`
        environment variable.
        """
        return pulumi.get(self, "server_url")

    @server_url.setter
    def server_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_url", value)

    @property
    @pulumi.getter(name="allowInsecureHttps")
    def allow_insecure_https(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to set whether to allow https with invalid certificates. Can be set via the `NETBOX_ALLOW_INSECURE_HTTPS`
        environment variable. Defaults to `false`.
        """
        return pulumi.get(self, "allow_insecure_https")

    @allow_insecure_https.setter
    def allow_insecure_https(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_insecure_https", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Set these header on all requests to Netbox. Can be set via the `NETBOX_HEADERS` environment variable.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="requestTimeout")
    def request_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Netbox API HTTP request timeout in seconds. Can be set via the `NETBOX_REQUEST_TIMEOUT` environment variable.
        """
        return pulumi.get(self, "request_timeout")

    @request_timeout.setter
    def request_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "request_timeout", value)

    @property
    @pulumi.getter(name="skipVersionCheck")
    def skip_version_check(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "skip_version_check")

    @skip_version_check.setter
    def skip_version_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_version_check", value)

    @property
    @pulumi.getter(name="stripTrailingSlashesFromUrl")
    def strip_trailing_slashes_from_url(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, strip trailing slashes from the `server_url` parameter and print a warning when doing so. Note that using
        trailing slashes in the `server_url` parameter will usually lead to errors. Can be set via the
        `NETBOX_STRIP_TRAILING_SLASHES_FROM_URL` environment variable. Defaults to `true`.
        """
        return pulumi.get(self, "strip_trailing_slashes_from_url")

    @strip_trailing_slashes_from_url.setter
    def strip_trailing_slashes_from_url(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "strip_trailing_slashes_from_url", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_insecure_https: Optional[pulumi.Input[bool]] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 request_timeout: Optional[pulumi.Input[int]] = None,
                 server_url: Optional[pulumi.Input[str]] = None,
                 skip_version_check: Optional[pulumi.Input[bool]] = None,
                 strip_trailing_slashes_from_url: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The provider type for the netbox package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_insecure_https: Flag to set whether to allow https with invalid certificates. Can be set via the `NETBOX_ALLOW_INSECURE_HTTPS`
               environment variable. Defaults to `false`.
        :param pulumi.Input[str] api_token: Netbox API authentication token. Can be set via the `NETBOX_API_TOKEN` environment variable.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: Set these header on all requests to Netbox. Can be set via the `NETBOX_HEADERS` environment variable.
        :param pulumi.Input[int] request_timeout: Netbox API HTTP request timeout in seconds. Can be set via the `NETBOX_REQUEST_TIMEOUT` environment variable.
        :param pulumi.Input[str] server_url: Location of Netbox server including scheme (http or https) and optional port. Can be set via the `NETBOX_SERVER_URL`
               environment variable.
        :param pulumi.Input[bool] strip_trailing_slashes_from_url: If true, strip trailing slashes from the `server_url` parameter and print a warning when doing so. Note that using
               trailing slashes in the `server_url` parameter will usually lead to errors. Can be set via the
               `NETBOX_STRIP_TRAILING_SLASHES_FROM_URL` environment variable. Defaults to `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the netbox package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_insecure_https: Optional[pulumi.Input[bool]] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 request_timeout: Optional[pulumi.Input[int]] = None,
                 server_url: Optional[pulumi.Input[str]] = None,
                 skip_version_check: Optional[pulumi.Input[bool]] = None,
                 strip_trailing_slashes_from_url: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["allow_insecure_https"] = pulumi.Output.from_input(allow_insecure_https).apply(pulumi.runtime.to_json) if allow_insecure_https is not None else None
            if api_token is None and not opts.urn:
                raise TypeError("Missing required property 'api_token'")
            __props__.__dict__["api_token"] = api_token
            __props__.__dict__["headers"] = pulumi.Output.from_input(headers).apply(pulumi.runtime.to_json) if headers is not None else None
            __props__.__dict__["request_timeout"] = pulumi.Output.from_input(request_timeout).apply(pulumi.runtime.to_json) if request_timeout is not None else None
            if server_url is None and not opts.urn:
                raise TypeError("Missing required property 'server_url'")
            __props__.__dict__["server_url"] = server_url
            __props__.__dict__["skip_version_check"] = pulumi.Output.from_input(skip_version_check).apply(pulumi.runtime.to_json) if skip_version_check is not None else None
            __props__.__dict__["strip_trailing_slashes_from_url"] = pulumi.Output.from_input(strip_trailing_slashes_from_url).apply(pulumi.runtime.to_json) if strip_trailing_slashes_from_url is not None else None
        super(Provider, __self__).__init__(
            'netbox',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Output[str]:
        """
        Netbox API authentication token. Can be set via the `NETBOX_API_TOKEN` environment variable.
        """
        return pulumi.get(self, "api_token")

    @property
    @pulumi.getter(name="serverUrl")
    def server_url(self) -> pulumi.Output[str]:
        """
        Location of Netbox server including scheme (http or https) and optional port. Can be set via the `NETBOX_SERVER_URL`
        environment variable.
        """
        return pulumi.get(self, "server_url")

