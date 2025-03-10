// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getPrefix(args?: GetPrefixArgs, opts?: pulumi.InvokeOptions): Promise<GetPrefixResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("netbox:ipam/getPrefix:getPrefix", {
        "cidr": args.cidr,
        "customFields": args.customFields,
        "description": args.description,
        "family": args.family,
        "prefix": args.prefix,
        "roleId": args.roleId,
        "siteId": args.siteId,
        "status": args.status,
        "tag": args.tag,
        "tagN": args.tagN,
        "tenantId": args.tenantId,
        "vlanId": args.vlanId,
        "vlanVid": args.vlanVid,
        "vrfId": args.vrfId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrefix.
 */
export interface GetPrefixArgs {
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given. Conflicts with `prefix`.
     *
     * @deprecated The `cidr` parameter is deprecated in favor of the canonical `prefix` attribute.
     */
    cidr?: string;
    customFields?: {[key: string]: string};
    /**
     * Description to include in the data source filter. At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    description?: string;
    /**
     * The IP family of the prefix. One of 4 or 6. At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    family?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given. Conflicts with `cidr`.
     */
    prefix?: string;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    roleId?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    siteId?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    status?: string;
    /**
     * Tag to include in the data source filter (must match the tag's slug). At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    tag?: string;
    /**
     * Tag to exclude from the data source filter (must match the tag's slug).
     * Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
     * for more information on available lookup expressions.
     */
    tagN?: string;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    tenantId?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    vlanId?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    vlanVid?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    vrfId?: number;
}

/**
 * A collection of values returned by getPrefix.
 */
export interface GetPrefixResult {
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given. Conflicts with `prefix`.
     *
     * @deprecated The `cidr` parameter is deprecated in favor of the canonical `prefix` attribute.
     */
    readonly cidr?: string;
    readonly customFields?: {[key: string]: string};
    /**
     * Description to include in the data source filter. At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly description: string;
    /**
     * The IP family of the prefix. One of 4 or 6. At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly family: number;
    /**
     * The ID of this resource.
     */
    readonly id: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given. Conflicts with `cidr`.
     */
    readonly prefix?: string;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly roleId: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly siteId?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly status?: string;
    /**
     * Tag to include in the data source filter (must match the tag's slug). At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly tag?: string;
    /**
     * Tag to exclude from the data source filter (must match the tag's slug).
     * Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
     * for more information on available lookup expressions.
     */
    readonly tagN?: string;
    readonly tags: string[];
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly tenantId?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly vlanId?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly vlanVid?: number;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    readonly vrfId?: number;
}
export function getPrefixOutput(args?: GetPrefixOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPrefixResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("netbox:ipam/getPrefix:getPrefix", {
        "cidr": args.cidr,
        "customFields": args.customFields,
        "description": args.description,
        "family": args.family,
        "prefix": args.prefix,
        "roleId": args.roleId,
        "siteId": args.siteId,
        "status": args.status,
        "tag": args.tag,
        "tagN": args.tagN,
        "tenantId": args.tenantId,
        "vlanId": args.vlanId,
        "vlanVid": args.vlanVid,
        "vrfId": args.vrfId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrefix.
 */
export interface GetPrefixOutputArgs {
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given. Conflicts with `prefix`.
     *
     * @deprecated The `cidr` parameter is deprecated in favor of the canonical `prefix` attribute.
     */
    cidr?: pulumi.Input<string>;
    customFields?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Description to include in the data source filter. At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP family of the prefix. One of 4 or 6. At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    family?: pulumi.Input<number>;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given. Conflicts with `cidr`.
     */
    prefix?: pulumi.Input<string>;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    roleId?: pulumi.Input<number>;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    siteId?: pulumi.Input<number>;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    status?: pulumi.Input<string>;
    /**
     * Tag to include in the data source filter (must match the tag's slug). At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    tag?: pulumi.Input<string>;
    /**
     * Tag to exclude from the data source filter (must match the tag's slug).
     * Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
     * for more information on available lookup expressions.
     */
    tagN?: pulumi.Input<string>;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    tenantId?: pulumi.Input<number>;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    vlanId?: pulumi.Input<number>;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    vlanVid?: pulumi.Input<number>;
    /**
     * At least one of `description`, `family`, `prefix`, `vlanVid`, `vrfId`, `vlanId`, `tenantId`, `siteId`, `roleId`, `cidr`, `tag` or `status` must be given.
     */
    vrfId?: pulumi.Input<number>;
}
