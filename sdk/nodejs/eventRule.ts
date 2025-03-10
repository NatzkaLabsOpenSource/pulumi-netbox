// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/event-rules/):
 *
 * > NetBox can be configured via Event Rules to transmit outgoing webhooks to remote systems in response to internal object changes. The receiver can act on the data in these webhook messages to perform related tasks.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * const test = new netbox.Webhook("test", {
 *     name: "my-webhook",
 *     payloadUrl: "https://example.com/webhook",
 * });
 * const testEventRule = new netbox.EventRule("test", {
 *     name: "my-event-rule",
 *     contentTypes: [
 *         "dcim.site",
 *         "virtualization.cluster",
 *     ],
 *     actionType: "webhook",
 *     actionObjectId: test.id,
 *     eventTypes: [
 *         "object_created",
 *         "object_updated",
 *         "object_deleted",
 *         "job_started",
 *         "job_completed",
 *         "job_failed",
 *         "job_errored",
 *     ],
 * });
 * ```
 */
export class EventRule extends pulumi.CustomResource {
    /**
     * Get an existing EventRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventRuleState, opts?: pulumi.CustomResourceOptions): EventRule {
        return new EventRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:index/eventRule:EventRule';

    /**
     * Returns true if the given object is an instance of EventRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventRule.__pulumiType;
    }

    public readonly actionObjectId!: pulumi.Output<number>;
    /**
     * Valid values are `webhook`.
     */
    public readonly actionType!: pulumi.Output<string>;
    public readonly conditions!: pulumi.Output<string | undefined>;
    public readonly contentTypes!: pulumi.Output<string[]>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The types of event which will trigger this rule. By default, valid values are `objectCreated`, `ojectUpdated`, `objectDeleted`, `jobStarted`, `jobCompleted`, `jobFailed` and `jobErrored`.
     */
    public readonly eventTypes!: pulumi.Output<string[]>;
    public readonly name!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a EventRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventRuleArgs | EventRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventRuleState | undefined;
            resourceInputs["actionObjectId"] = state ? state.actionObjectId : undefined;
            resourceInputs["actionType"] = state ? state.actionType : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["contentTypes"] = state ? state.contentTypes : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["eventTypes"] = state ? state.eventTypes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EventRuleArgs | undefined;
            if ((!args || args.actionObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actionObjectId'");
            }
            if ((!args || args.actionType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actionType'");
            }
            if ((!args || args.contentTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contentTypes'");
            }
            if ((!args || args.eventTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventTypes'");
            }
            resourceInputs["actionObjectId"] = args ? args.actionObjectId : undefined;
            resourceInputs["actionType"] = args ? args.actionType : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["contentTypes"] = args ? args.contentTypes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["eventTypes"] = args ? args.eventTypes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventRule resources.
 */
export interface EventRuleState {
    actionObjectId?: pulumi.Input<number>;
    /**
     * Valid values are `webhook`.
     */
    actionType?: pulumi.Input<string>;
    conditions?: pulumi.Input<string>;
    contentTypes?: pulumi.Input<pulumi.Input<string>[]>;
    description?: pulumi.Input<string>;
    /**
     * Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The types of event which will trigger this rule. By default, valid values are `objectCreated`, `ojectUpdated`, `objectDeleted`, `jobStarted`, `jobCompleted`, `jobFailed` and `jobErrored`.
     */
    eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a EventRule resource.
 */
export interface EventRuleArgs {
    actionObjectId: pulumi.Input<number>;
    /**
     * Valid values are `webhook`.
     */
    actionType: pulumi.Input<string>;
    conditions?: pulumi.Input<string>;
    contentTypes: pulumi.Input<pulumi.Input<string>[]>;
    description?: pulumi.Input<string>;
    /**
     * Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The types of event which will trigger this rule. By default, valid values are `objectCreated`, `ojectUpdated`, `objectDeleted`, `jobStarted`, `jobCompleted`, `jobFailed` and `jobErrored`.
     */
    eventTypes: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
