// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/customization/custom-fields/#custom-fields):
 *
 * > Each model in NetBox is represented in the database as a discrete table, and each attribute of a model exists as a column within its table. For example, sites are stored in the dcimSite table, which has columns named name, facility, physical_address, and so on. As new attributes are added to objects throughout the development of NetBox, tables are expanded to include new rows.
 * > 
 * > However, some users might want to store additional object attributes that are somewhat esoteric in nature, and that would not make sense to include in the core NetBox database schema. For instance, suppose your organization needs to associate each device with a ticket number correlating it with an internal support system record. This is certainly a legitimate use for NetBox, but it's not a common enough need to warrant including a field for every NetBox installation. Instead, you can create a custom field to hold this data.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzkalabsopensource/pulumi-netbox";
 *
 * const test = new netbox.CustomField("test", {
 *     contentTypes: ["virtualization.vminterface"],
 *     type: "text",
 *     validationRegex: "^.*$",
 *     weight: 100,
 * });
 * ```
 */
export class CustomField extends pulumi.CustomResource {
    /**
     * Get an existing CustomField resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomFieldState, opts?: pulumi.CustomResourceOptions): CustomField {
        return new CustomField(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:index/customField:CustomField';

    /**
     * Returns true if the given object is an instance of CustomField.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomField {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomField.__pulumiType;
    }

    public readonly choiceSetId!: pulumi.Output<number | undefined>;
    public readonly contentTypes!: pulumi.Output<string[]>;
    public readonly default!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly groupName!: pulumi.Output<string | undefined>;
    public readonly label!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly required!: pulumi.Output<boolean | undefined>;
    public readonly type!: pulumi.Output<string>;
    public readonly validationMaximum!: pulumi.Output<number | undefined>;
    public readonly validationMinimum!: pulumi.Output<number | undefined>;
    public readonly validationRegex!: pulumi.Output<string | undefined>;
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a CustomField resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomFieldArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomFieldArgs | CustomFieldState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomFieldState | undefined;
            resourceInputs["choiceSetId"] = state ? state.choiceSetId : undefined;
            resourceInputs["contentTypes"] = state ? state.contentTypes : undefined;
            resourceInputs["default"] = state ? state.default : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["required"] = state ? state.required : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["validationMaximum"] = state ? state.validationMaximum : undefined;
            resourceInputs["validationMinimum"] = state ? state.validationMinimum : undefined;
            resourceInputs["validationRegex"] = state ? state.validationRegex : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as CustomFieldArgs | undefined;
            if ((!args || args.contentTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contentTypes'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.weight === undefined) && !opts.urn) {
                throw new Error("Missing required property 'weight'");
            }
            resourceInputs["choiceSetId"] = args ? args.choiceSetId : undefined;
            resourceInputs["contentTypes"] = args ? args.contentTypes : undefined;
            resourceInputs["default"] = args ? args.default : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["required"] = args ? args.required : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["validationMaximum"] = args ? args.validationMaximum : undefined;
            resourceInputs["validationMinimum"] = args ? args.validationMinimum : undefined;
            resourceInputs["validationRegex"] = args ? args.validationRegex : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomField.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomField resources.
 */
export interface CustomFieldState {
    choiceSetId?: pulumi.Input<number>;
    contentTypes?: pulumi.Input<pulumi.Input<string>[]>;
    default?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    groupName?: pulumi.Input<string>;
    label?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    required?: pulumi.Input<boolean>;
    type?: pulumi.Input<string>;
    validationMaximum?: pulumi.Input<number>;
    validationMinimum?: pulumi.Input<number>;
    validationRegex?: pulumi.Input<string>;
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CustomField resource.
 */
export interface CustomFieldArgs {
    choiceSetId?: pulumi.Input<number>;
    contentTypes: pulumi.Input<pulumi.Input<string>[]>;
    default?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    groupName?: pulumi.Input<string>;
    label?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    required?: pulumi.Input<boolean>;
    type: pulumi.Input<string>;
    validationMaximum?: pulumi.Input<number>;
    validationMinimum?: pulumi.Input<number>;
    validationRegex?: pulumi.Input<string>;
    weight: pulumi.Input<number>;
}