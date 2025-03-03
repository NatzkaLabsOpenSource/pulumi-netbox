// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Ipam
{
    /// <summary>
    /// From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#prefixes):
    /// 
    /// &gt; A prefix is an IPv4 or IPv6 network and mask expressed in CIDR notation (e.g. 192.0.2.0/24). A prefix entails only the "network portion" of an IP address: All bits in the address not covered by the mask must be zero. (In other words, a prefix cannot be a specific IP address.)
    /// &gt; 
    /// &gt; Prefixes are automatically organized by their parent aggregates. Additionally, each prefix can be assigned to a particular site and virtual routing and forwarding instance (VRF). Each VRF represents a separate IP space or routing table. All prefixes not assigned to a VRF are considered to be in the "global" table.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Netbox = Pulumi.Netbox;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myPrefix = new Netbox.Ipam.Prefix("my_prefix", new()
    ///     {
    ///         Address = "10.0.0.0/24",
    ///         Status = "active",
    ///         Description = "test prefix",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:ipam/prefix:Prefix")]
    public partial class Prefix : global::Pulumi.CustomResource
    {
        [Output("customFields")]
        public Output<ImmutableDictionary<string, string>?> CustomFields { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("isPool")]
        public Output<bool?> IsPool { get; private set; } = null!;

        [Output("markUtilized")]
        public Output<bool?> MarkUtilized { get; private set; } = null!;

        [Output("prefix")]
        public Output<string> Address { get; private set; } = null!;

        [Output("roleId")]
        public Output<int?> RoleId { get; private set; } = null!;

        [Output("siteId")]
        public Output<int?> SiteId { get; private set; } = null!;

        /// <summary>
        /// Valid values are `active`, `container`, `reserved` and `deprecated`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("tenantId")]
        public Output<int?> TenantId { get; private set; } = null!;

        [Output("vlanId")]
        public Output<int?> VlanId { get; private set; } = null!;

        [Output("vrfId")]
        public Output<int?> VrfId { get; private set; } = null!;


        /// <summary>
        /// Create a Prefix resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Prefix(string name, PrefixArgs args, CustomResourceOptions? options = null)
            : base("netbox:ipam/prefix:Prefix", name, args ?? new PrefixArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Prefix(string name, Input<string> id, PrefixState? state = null, CustomResourceOptions? options = null)
            : base("netbox:ipam/prefix:Prefix", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/NatzkaLabsOpenSource/pulumi-netbox",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Prefix resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Prefix Get(string name, Input<string> id, PrefixState? state = null, CustomResourceOptions? options = null)
        {
            return new Prefix(name, id, state, options);
        }
    }

    public sealed class PrefixArgs : global::Pulumi.ResourceArgs
    {
        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("isPool")]
        public Input<bool>? IsPool { get; set; }

        [Input("markUtilized")]
        public Input<bool>? MarkUtilized { get; set; }

        [Input("prefix", required: true)]
        public Input<string> Address { get; set; } = null!;

        [Input("roleId")]
        public Input<int>? RoleId { get; set; }

        [Input("siteId")]
        public Input<int>? SiteId { get; set; }

        /// <summary>
        /// Valid values are `active`, `container`, `reserved` and `deprecated`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tenantId")]
        public Input<int>? TenantId { get; set; }

        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        [Input("vrfId")]
        public Input<int>? VrfId { get; set; }

        public PrefixArgs()
        {
        }
        public static new PrefixArgs Empty => new PrefixArgs();
    }

    public sealed class PrefixState : global::Pulumi.ResourceArgs
    {
        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("isPool")]
        public Input<bool>? IsPool { get; set; }

        [Input("markUtilized")]
        public Input<bool>? MarkUtilized { get; set; }

        [Input("prefix")]
        public Input<string>? Address { get; set; }

        [Input("roleId")]
        public Input<int>? RoleId { get; set; }

        [Input("siteId")]
        public Input<int>? SiteId { get; set; }

        /// <summary>
        /// Valid values are `active`, `container`, `reserved` and `deprecated`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tenantId")]
        public Input<int>? TenantId { get; set; }

        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        [Input("vrfId")]
        public Input<int>? VrfId { get; set; }

        public PrefixState()
        {
        }
        public static new PrefixState Empty => new PrefixState();
    }
}
