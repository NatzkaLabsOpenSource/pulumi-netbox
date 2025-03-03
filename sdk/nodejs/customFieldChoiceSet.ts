// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/models/extras/customfieldchoiceset/):
 *
 * Single- and multi-selection custom fields must define a set of valid choices from which the user may choose when defining the field value. These choices are defined as sets that may be reused among multiple custom fields.
 *
 * A choice set must define a base choice set and/or a set of arbitrary extra choices.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * const test = new netbox.CustomFieldChoiceSet("test", {
 *     name: "my-custom-field-set",
 *     description: "Description",
 *     extraChoices: [
 *         [
 *             "choice1",
 *             "label1",
 *         ],
 *         [
 *             "choice2",
 *             "choice2",
 *         ],
 *     ],
 * });
 * ```
 */
export class CustomFieldChoiceSet extends pulumi.CustomResource {
    /**
     * Get an existing CustomFieldChoiceSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomFieldChoiceSetState, opts?: pulumi.CustomResourceOptions): CustomFieldChoiceSet {
        return new CustomFieldChoiceSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:index/customFieldChoiceSet:CustomFieldChoiceSet';

    /**
     * Returns true if the given object is an instance of CustomFieldChoiceSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomFieldChoiceSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomFieldChoiceSet.__pulumiType;
    }

    /**
     * Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `baseChoices` or `extraChoices` must be given.
     */
    public readonly baseChoices!: pulumi.Output<string | undefined>;
    public readonly customFields!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `baseChoices` or `extraChoices` must be given.
     */
    public readonly extraChoices!: pulumi.Output<string[][] | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * experimental. Defaults to `false`.
     */
    public readonly orderAlphabetically!: pulumi.Output<boolean | undefined>;

    /**
     * Create a CustomFieldChoiceSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CustomFieldChoiceSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomFieldChoiceSetArgs | CustomFieldChoiceSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomFieldChoiceSetState | undefined;
            resourceInputs["baseChoices"] = state ? state.baseChoices : undefined;
            resourceInputs["customFields"] = state ? state.customFields : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["extraChoices"] = state ? state.extraChoices : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orderAlphabetically"] = state ? state.orderAlphabetically : undefined;
        } else {
            const args = argsOrState as CustomFieldChoiceSetArgs | undefined;
            resourceInputs["baseChoices"] = args ? args.baseChoices : undefined;
            resourceInputs["customFields"] = args ? args.customFields : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["extraChoices"] = args ? args.extraChoices : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orderAlphabetically"] = args ? args.orderAlphabetically : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomFieldChoiceSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomFieldChoiceSet resources.
 */
export interface CustomFieldChoiceSetState {
    /**
     * Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `baseChoices` or `extraChoices` must be given.
     */
    baseChoices?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `baseChoices` or `extraChoices` must be given.
     */
    extraChoices?: pulumi.Input<pulumi.Input<pulumi.Input<string>[]>[]>;
    name?: pulumi.Input<string>;
    /**
     * experimental. Defaults to `false`.
     */
    orderAlphabetically?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a CustomFieldChoiceSet resource.
 */
export interface CustomFieldChoiceSetArgs {
    /**
     * Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `baseChoices` or `extraChoices` must be given.
     */
    baseChoices?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    description?: pulumi.Input<string>;
    /**
     * This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `baseChoices` or `extraChoices` must be given.
     */
    extraChoices?: pulumi.Input<pulumi.Input<pulumi.Input<string>[]>[]>;
    name?: pulumi.Input<string>;
    /**
     * experimental. Defaults to `false`.
     */
    orderAlphabetically?: pulumi.Input<boolean>;
}
