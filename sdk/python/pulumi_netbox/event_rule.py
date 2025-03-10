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

__all__ = ['EventRuleArgs', 'EventRule']

@pulumi.input_type
class EventRuleArgs:
    def __init__(__self__, *,
                 action_object_id: pulumi.Input[int],
                 action_type: pulumi.Input[str],
                 content_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 event_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 conditions: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a EventRule resource.
        :param pulumi.Input[str] action_type: Valid values are `webhook`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: The types of event which will trigger this rule. By default, valid values are `object_created`, `oject_updated`, `object_deleted`, `job_started`, `job_completed`, `job_failed` and `job_errored`.
        :param pulumi.Input[bool] enabled: Defaults to `true`.
        """
        pulumi.set(__self__, "action_object_id", action_object_id)
        pulumi.set(__self__, "action_type", action_type)
        pulumi.set(__self__, "content_types", content_types)
        pulumi.set(__self__, "event_types", event_types)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="actionObjectId")
    def action_object_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "action_object_id")

    @action_object_id.setter
    def action_object_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "action_object_id", value)

    @property
    @pulumi.getter(name="actionType")
    def action_type(self) -> pulumi.Input[str]:
        """
        Valid values are `webhook`.
        """
        return pulumi.get(self, "action_type")

    @action_type.setter
    def action_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "action_type", value)

    @property
    @pulumi.getter(name="contentTypes")
    def content_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "content_types")

    @content_types.setter
    def content_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "content_types", value)

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The types of event which will trigger this rule. By default, valid values are `object_created`, `oject_updated`, `object_deleted`, `job_started`, `job_completed`, `job_failed` and `job_errored`.
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "event_types", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EventRuleState:
    def __init__(__self__, *,
                 action_object_id: Optional[pulumi.Input[int]] = None,
                 action_type: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[str]] = None,
                 content_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering EventRule resources.
        :param pulumi.Input[str] action_type: Valid values are `webhook`.
        :param pulumi.Input[bool] enabled: Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: The types of event which will trigger this rule. By default, valid values are `object_created`, `oject_updated`, `object_deleted`, `job_started`, `job_completed`, `job_failed` and `job_errored`.
        """
        if action_object_id is not None:
            pulumi.set(__self__, "action_object_id", action_object_id)
        if action_type is not None:
            pulumi.set(__self__, "action_type", action_type)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if content_types is not None:
            pulumi.set(__self__, "content_types", content_types)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if event_types is not None:
            pulumi.set(__self__, "event_types", event_types)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="actionObjectId")
    def action_object_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "action_object_id")

    @action_object_id.setter
    def action_object_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "action_object_id", value)

    @property
    @pulumi.getter(name="actionType")
    def action_type(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `webhook`.
        """
        return pulumi.get(self, "action_type")

    @action_type.setter
    def action_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_type", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="contentTypes")
    def content_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "content_types")

    @content_types.setter
    def content_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "content_types", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The types of event which will trigger this rule. By default, valid values are `object_created`, `oject_updated`, `object_deleted`, `job_started`, `job_completed`, `job_failed` and `job_errored`.
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_types", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class EventRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_object_id: Optional[pulumi.Input[int]] = None,
                 action_type: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[str]] = None,
                 content_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/features/event-rules/):

        > NetBox can be configured via Event Rules to transmit outgoing webhooks to remote systems in response to internal object changes. The receiver can act on the data in these webhook messages to perform related tasks.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        test = netbox.Webhook("test",
            name="my-webhook",
            payload_url="https://example.com/webhook")
        test_event_rule = netbox.EventRule("test",
            name="my-event-rule",
            content_types=[
                "dcim.site",
                "virtualization.cluster",
            ],
            action_type="webhook",
            action_object_id=test.id,
            event_types=[
                "object_created",
                "object_updated",
                "object_deleted",
                "job_started",
                "job_completed",
                "job_failed",
                "job_errored",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action_type: Valid values are `webhook`.
        :param pulumi.Input[bool] enabled: Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: The types of event which will trigger this rule. By default, valid values are `object_created`, `oject_updated`, `object_deleted`, `job_started`, `job_completed`, `job_failed` and `job_errored`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        From the [official documentation](https://docs.netbox.dev/en/stable/features/event-rules/):

        > NetBox can be configured via Event Rules to transmit outgoing webhooks to remote systems in response to internal object changes. The receiver can act on the data in these webhook messages to perform related tasks.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_netbox as netbox

        test = netbox.Webhook("test",
            name="my-webhook",
            payload_url="https://example.com/webhook")
        test_event_rule = netbox.EventRule("test",
            name="my-event-rule",
            content_types=[
                "dcim.site",
                "virtualization.cluster",
            ],
            action_type="webhook",
            action_object_id=test.id,
            event_types=[
                "object_created",
                "object_updated",
                "object_deleted",
                "job_started",
                "job_completed",
                "job_failed",
                "job_errored",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param EventRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_object_id: Optional[pulumi.Input[int]] = None,
                 action_type: Optional[pulumi.Input[str]] = None,
                 conditions: Optional[pulumi.Input[str]] = None,
                 content_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventRuleArgs.__new__(EventRuleArgs)

            if action_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'action_object_id'")
            __props__.__dict__["action_object_id"] = action_object_id
            if action_type is None and not opts.urn:
                raise TypeError("Missing required property 'action_type'")
            __props__.__dict__["action_type"] = action_type
            __props__.__dict__["conditions"] = conditions
            if content_types is None and not opts.urn:
                raise TypeError("Missing required property 'content_types'")
            __props__.__dict__["content_types"] = content_types
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            if event_types is None and not opts.urn:
                raise TypeError("Missing required property 'event_types'")
            __props__.__dict__["event_types"] = event_types
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
        super(EventRule, __self__).__init__(
            'netbox:index/eventRule:EventRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action_object_id: Optional[pulumi.Input[int]] = None,
            action_type: Optional[pulumi.Input[str]] = None,
            conditions: Optional[pulumi.Input[str]] = None,
            content_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'EventRule':
        """
        Get an existing EventRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action_type: Valid values are `webhook`.
        :param pulumi.Input[bool] enabled: Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: The types of event which will trigger this rule. By default, valid values are `object_created`, `oject_updated`, `object_deleted`, `job_started`, `job_completed`, `job_failed` and `job_errored`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventRuleState.__new__(_EventRuleState)

        __props__.__dict__["action_object_id"] = action_object_id
        __props__.__dict__["action_type"] = action_type
        __props__.__dict__["conditions"] = conditions
        __props__.__dict__["content_types"] = content_types
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["event_types"] = event_types
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        return EventRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actionObjectId")
    def action_object_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "action_object_id")

    @property
    @pulumi.getter(name="actionType")
    def action_type(self) -> pulumi.Output[str]:
        """
        Valid values are `webhook`.
        """
        return pulumi.get(self, "action_type")

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="contentTypes")
    def content_types(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "content_types")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> pulumi.Output[Sequence[str]]:
        """
        The types of event which will trigger this rule. By default, valid values are `object_created`, `oject_updated`, `object_deleted`, `job_started`, `job_completed`, `job_failed` and `job_errored`.
        """
        return pulumi.get(self, "event_types")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

