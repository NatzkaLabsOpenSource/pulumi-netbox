// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Ipam
{
    public static class GetPrefix
    {
        public static Task<GetPrefixResult> InvokeAsync(GetPrefixArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrefixResult>("netbox:ipam/getPrefix:getPrefix", args ?? new GetPrefixArgs(), options.WithDefaults());

        public static Output<GetPrefixResult> Invoke(GetPrefixInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrefixResult>("netbox:ipam/getPrefix:getPrefix", args ?? new GetPrefixInvokeArgs(), options.WithDefaults());

        public static Output<GetPrefixResult> Invoke(GetPrefixInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrefixResult>("netbox:ipam/getPrefix:getPrefix", args ?? new GetPrefixInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrefixArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given. Conflicts with `prefix`.
        /// </summary>
        [Input("cidr")]
        public string? Cidr { get; set; }

        [Input("customFields")]
        private Dictionary<string, string>? _customFields;
        public Dictionary<string, string> CustomFields
        {
            get => _customFields ?? (_customFields = new Dictionary<string, string>());
            set => _customFields = value;
        }

        /// <summary>
        /// Description to include in the data source filter. At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The IP family of the prefix. One of 4 or 6. At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("family")]
        public int? Family { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given. Conflicts with `cidr`.
        /// </summary>
        [Input("prefix")]
        public string? Prefix { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("roleId")]
        public int? RoleId { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("siteId")]
        public int? SiteId { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// Tag to include in the data source filter (must match the tag's slug). At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        /// <summary>
        /// Tag to exclude from the data source filter (must match the tag's slug).
        /// Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
        /// for more information on available lookup expressions.
        /// </summary>
        [Input("tagN")]
        public string? TagN { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("tenantId")]
        public int? TenantId { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("vlanId")]
        public int? VlanId { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("vlanVid")]
        public double? VlanVid { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("vrfId")]
        public int? VrfId { get; set; }

        public GetPrefixArgs()
        {
        }
        public static new GetPrefixArgs Empty => new GetPrefixArgs();
    }

    public sealed class GetPrefixInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given. Conflicts with `prefix`.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        /// <summary>
        /// Description to include in the data source filter. At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP family of the prefix. One of 4 or 6. At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("family")]
        public Input<int>? Family { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given. Conflicts with `cidr`.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("roleId")]
        public Input<int>? RoleId { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("siteId")]
        public Input<int>? SiteId { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Tag to include in the data source filter (must match the tag's slug). At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// Tag to exclude from the data source filter (must match the tag's slug).
        /// Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
        /// for more information on available lookup expressions.
        /// </summary>
        [Input("tagN")]
        public Input<string>? TagN { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("tenantId")]
        public Input<int>? TenantId { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("vlanVid")]
        public Input<double>? VlanVid { get; set; }

        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        [Input("vrfId")]
        public Input<int>? VrfId { get; set; }

        public GetPrefixInvokeArgs()
        {
        }
        public static new GetPrefixInvokeArgs Empty => new GetPrefixInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrefixResult
    {
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given. Conflicts with `prefix`.
        /// </summary>
        public readonly string? Cidr;
        public readonly ImmutableDictionary<string, string>? CustomFields;
        /// <summary>
        /// Description to include in the data source filter. At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The IP family of the prefix. One of 4 or 6. At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly int Family;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given. Conflicts with `cidr`.
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly int RoleId;
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly int? SiteId;
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Tag to include in the data source filter (must match the tag's slug). At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// Tag to exclude from the data source filter (must match the tag's slug).
        /// Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
        /// for more information on available lookup expressions.
        /// </summary>
        public readonly string? TagN;
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly int? TenantId;
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly int? VlanId;
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly double? VlanVid;
        /// <summary>
        /// At least one of `description`, `family`, `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `tenant_id`, `site_id`, `role_id`, `cidr`, `tag` or `status` must be given.
        /// </summary>
        public readonly int? VrfId;

        [OutputConstructor]
        private GetPrefixResult(
            string? cidr,

            ImmutableDictionary<string, string>? customFields,

            string description,

            int family,

            int id,

            string? prefix,

            int roleId,

            int? siteId,

            string? status,

            string? tag,

            string? tagN,

            ImmutableArray<string> tags,

            int? tenantId,

            int? vlanId,

            double? vlanVid,

            int? vrfId)
        {
            Cidr = cidr;
            CustomFields = customFields;
            Description = description;
            Family = family;
            Id = id;
            Prefix = prefix;
            RoleId = roleId;
            SiteId = siteId;
            Status = status;
            Tag = tag;
            TagN = tagN;
            Tags = tags;
            TenantId = tenantId;
            VlanId = vlanId;
            VlanVid = vlanVid;
            VrfId = vrfId;
        }
    }
}
