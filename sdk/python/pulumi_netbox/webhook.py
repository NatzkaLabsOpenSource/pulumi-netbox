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

__all__ = ['WebhookArgs', 'Webhook']

@pulumi.input_type
class WebhookArgs:
    def __init__(__self__, *,
                 payload_url: pulumi.Input[str],
                 additional_headers: Optional[pulumi.Input[str]] = None,
                 body_template: Optional[pulumi.Input[str]] = None,
                 http_content_type: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Webhook resource.
        :param pulumi.Input[str] http_content_type: The complete list of official content types is available [here](https://www.iana.org/assignments/media-types/media-types.xhtml). Defaults to `application/json`.
        :param pulumi.Input[str] http_method: Valid values are `GET`, `POST`, `PUT`, `PATCH` and `DELETE`. Defaults to `POST`.
        """
        pulumi.set(__self__, "payload_url", payload_url)
        if additional_headers is not None:
            pulumi.set(__self__, "additional_headers", additional_headers)
        if body_template is not None:
            pulumi.set(__self__, "body_template", body_template)
        if http_content_type is not None:
            pulumi.set(__self__, "http_content_type", http_content_type)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="payloadUrl")
    def payload_url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "payload_url")

    @payload_url.setter
    def payload_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "payload_url", value)

    @property
    @pulumi.getter(name="additionalHeaders")
    def additional_headers(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "additional_headers")

    @additional_headers.setter
    def additional_headers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "additional_headers", value)

    @property
    @pulumi.getter(name="bodyTemplate")
    def body_template(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "body_template")

    @body_template.setter
    def body_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body_template", value)

    @property
    @pulumi.getter(name="httpContentType")
    def http_content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The complete list of official content types is available [here](https://www.iana.org/assignments/media-types/media-types.xhtml). Defaults to `application/json`.
        """
        return pulumi.get(self, "http_content_type")

    @http_content_type.setter
    def http_content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_content_type", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `GET`, `POST`, `PUT`, `PATCH` and `DELETE`. Defaults to `POST`.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _WebhookState:
    def __init__(__self__, *,
                 additional_headers: Optional[pulumi.Input[str]] = None,
                 body_template: Optional[pulumi.Input[str]] = None,
                 http_content_type: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Webhook resources.
        :param pulumi.Input[str] http_content_type: The complete list of official content types is available [here](https://www.iana.org/assignments/media-types/media-types.xhtml). Defaults to `application/json`.
        :param pulumi.Input[str] http_method: Valid values are `GET`, `POST`, `PUT`, `PATCH` and `DELETE`. Defaults to `POST`.
        """
        if additional_headers is not None:
            pulumi.set(__self__, "additional_headers", additional_headers)
        if body_template is not None:
            pulumi.set(__self__, "body_template", body_template)
        if http_content_type is not None:
            pulumi.set(__self__, "http_content_type", http_content_type)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if payload_url is not None:
            pulumi.set(__self__, "payload_url", payload_url)

    @property
    @pulumi.getter(name="additionalHeaders")
    def additional_headers(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "additional_headers")

    @additional_headers.setter
    def additional_headers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "additional_headers", value)

    @property
    @pulumi.getter(name="bodyTemplate")
    def body_template(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "body_template")

    @body_template.setter
    def body_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body_template", value)

    @property
    @pulumi.getter(name="httpContentType")
    def http_content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The complete list of official content types is available [here](https://www.iana.org/assignments/media-types/media-types.xhtml). Defaults to `application/json`.
        """
        return pulumi.get(self, "http_content_type")

    @http_content_type.setter
    def http_content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_content_type", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `GET`, `POST`, `PUT`, `PATCH` and `DELETE`. Defaults to `POST`.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="payloadUrl")
    def payload_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "payload_url")

    @payload_url.setter
    def payload_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_url", value)


class Webhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_headers: Optional[pulumi.Input[str]] = None,
                 body_template: Optional[pulumi.Input[str]] = None,
                 http_content_type: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/integrations/webhooks/):

        > A webhook is a mechanism for conveying to some external system a change that took place in NetBox. For example, you may want to notify a monitoring system whenever the status of a device is updated in NetBox. This can be done by creating a webhook for the device model in NetBox and identifying the webhook receiver. When NetBox detects a change to a device, an HTTP request containing the details of the change and who made it be sent to the specified receiver.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        test = netbox.Webhook("test",
            name="test",
            payload_url="https://example.com/webhook",
            bodytemplate="Sample body")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_content_type: The complete list of official content types is available [here](https://www.iana.org/assignments/media-types/media-types.xhtml). Defaults to `application/json`.
        :param pulumi.Input[str] http_method: Valid values are `GET`, `POST`, `PUT`, `PATCH` and `DELETE`. Defaults to `POST`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/integrations/webhooks/):

        > A webhook is a mechanism for conveying to some external system a change that took place in NetBox. For example, you may want to notify a monitoring system whenever the status of a device is updated in NetBox. This can be done by creating a webhook for the device model in NetBox and identifying the webhook receiver. When NetBox detects a change to a device, an HTTP request containing the details of the change and who made it be sent to the specified receiver.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        test = netbox.Webhook("test",
            name="test",
            payload_url="https://example.com/webhook",
            bodytemplate="Sample body")
        ```

        :param str resource_name: The name of the resource.
        :param WebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_headers: Optional[pulumi.Input[str]] = None,
                 body_template: Optional[pulumi.Input[str]] = None,
                 http_content_type: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebhookArgs.__new__(WebhookArgs)

            __props__.__dict__["additional_headers"] = additional_headers
            __props__.__dict__["body_template"] = body_template
            __props__.__dict__["http_content_type"] = http_content_type
            __props__.__dict__["http_method"] = http_method
            __props__.__dict__["name"] = name
            if payload_url is None and not opts.urn:
                raise TypeError("Missing required property 'payload_url'")
            __props__.__dict__["payload_url"] = payload_url
        super(Webhook, __self__).__init__(
            'netbox:index/webhook:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_headers: Optional[pulumi.Input[str]] = None,
            body_template: Optional[pulumi.Input[str]] = None,
            http_content_type: Optional[pulumi.Input[str]] = None,
            http_method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            payload_url: Optional[pulumi.Input[str]] = None) -> 'Webhook':
        """
        Get an existing Webhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_content_type: The complete list of official content types is available [here](https://www.iana.org/assignments/media-types/media-types.xhtml). Defaults to `application/json`.
        :param pulumi.Input[str] http_method: Valid values are `GET`, `POST`, `PUT`, `PATCH` and `DELETE`. Defaults to `POST`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebhookState.__new__(_WebhookState)

        __props__.__dict__["additional_headers"] = additional_headers
        __props__.__dict__["body_template"] = body_template
        __props__.__dict__["http_content_type"] = http_content_type
        __props__.__dict__["http_method"] = http_method
        __props__.__dict__["name"] = name
        __props__.__dict__["payload_url"] = payload_url
        return Webhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalHeaders")
    def additional_headers(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "additional_headers")

    @property
    @pulumi.getter(name="bodyTemplate")
    def body_template(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "body_template")

    @property
    @pulumi.getter(name="httpContentType")
    def http_content_type(self) -> pulumi.Output[Optional[str]]:
        """
        The complete list of official content types is available [here](https://www.iana.org/assignments/media-types/media-types.xhtml). Defaults to `application/json`.
        """
        return pulumi.get(self, "http_content_type")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> pulumi.Output[Optional[str]]:
        """
        Valid values are `GET`, `POST`, `PUT`, `PATCH` and `DELETE`. Defaults to `POST`.
        """
        return pulumi.get(self, "http_method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="payloadUrl")
    def payload_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "payload_url")

