// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/contacts/#contact-groups):
 *
 * > Contacts can be grouped arbitrarily into a recursive hierarchy, and a contact can be assigned to a group at any level within the hierarchy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * const test = new netbox.tenancy.ContactGroup("test", {name: "test"});
 * ```
 */
export class ContactGroup extends pulumi.CustomResource {
    /**
     * Get an existing ContactGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContactGroupState, opts?: pulumi.CustomResourceOptions): ContactGroup {
        return new ContactGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:tenancy/contactGroup:ContactGroup';

    /**
     * Returns true if the given object is an instance of ContactGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContactGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContactGroup.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly parentId!: pulumi.Output<number | undefined>;
    public readonly slug!: pulumi.Output<string>;

    /**
     * Create a ContactGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContactGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContactGroupArgs | ContactGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContactGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentId"] = state ? state.parentId : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as ContactGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentId"] = args ? args.parentId : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContactGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContactGroup resources.
 */
export interface ContactGroupState {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    parentId?: pulumi.Input<number>;
    slug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContactGroup resource.
 */
export interface ContactGroupArgs {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    parentId?: pulumi.Input<number>;
    slug?: pulumi.Input<string>;
}
