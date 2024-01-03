// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * From the [official documentation](https://docs.netbox.dev/en/stable/features/contacts#contactassignments_1):
 *
 * > Much like tenancy, contact assignment enables you to track ownership of resources modeled in NetBox.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as netbox from "@natzka-oss/pulumi-netbox";
 *
 * const testContact = new netbox.tenancy.Contact("testContact", {});
 * const testContactRole = new netbox.tenancy.ContactRole("testContactRole", {});
 * // Assumes that a device with id 123 exists
 * const testContactAssignment = new netbox.tenancy.ContactAssignment("testContactAssignment", {
 *     contentType: "dcim.device",
 *     objectId: 123,
 *     contactId: testContact.id,
 *     roleId: testContactRole.id,
 *     priority: "primary",
 * });
 * ```
 */
export class ContactAssignment extends pulumi.CustomResource {
    /**
     * Get an existing ContactAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContactAssignmentState, opts?: pulumi.CustomResourceOptions): ContactAssignment {
        return new ContactAssignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'netbox:tenancy/contactAssignment:ContactAssignment';

    /**
     * Returns true if the given object is an instance of ContactAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContactAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContactAssignment.__pulumiType;
    }

    public readonly contactId!: pulumi.Output<number>;
    public readonly contentType!: pulumi.Output<string>;
    public readonly objectId!: pulumi.Output<number>;
    /**
     * Valid values are `primary`, `secondary`, `tertiary` and `inactive`.
     */
    public readonly priority!: pulumi.Output<string | undefined>;
    public readonly roleId!: pulumi.Output<number>;

    /**
     * Create a ContactAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContactAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContactAssignmentArgs | ContactAssignmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContactAssignmentState | undefined;
            resourceInputs["contactId"] = state ? state.contactId : undefined;
            resourceInputs["contentType"] = state ? state.contentType : undefined;
            resourceInputs["objectId"] = state ? state.objectId : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
        } else {
            const args = argsOrState as ContactAssignmentArgs | undefined;
            if ((!args || args.contactId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contactId'");
            }
            if ((!args || args.contentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contentType'");
            }
            if ((!args || args.objectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectId'");
            }
            if ((!args || args.roleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleId'");
            }
            resourceInputs["contactId"] = args ? args.contactId : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["objectId"] = args ? args.objectId : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContactAssignment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContactAssignment resources.
 */
export interface ContactAssignmentState {
    contactId?: pulumi.Input<number>;
    contentType?: pulumi.Input<string>;
    objectId?: pulumi.Input<number>;
    /**
     * Valid values are `primary`, `secondary`, `tertiary` and `inactive`.
     */
    priority?: pulumi.Input<string>;
    roleId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ContactAssignment resource.
 */
export interface ContactAssignmentArgs {
    contactId: pulumi.Input<number>;
    contentType: pulumi.Input<string>;
    objectId: pulumi.Input<number>;
    /**
     * Valid values are `primary`, `secondary`, `tertiary` and `inactive`.
     */
    priority?: pulumi.Input<string>;
    roleId: pulumi.Input<number>;
}
