// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Dcim
{
    /// <summary>
    /// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rack/):
    /// 
    /// &gt; The rack model represents a physical two- or four-post equipment rack in which devices can be installed. Each rack must be assigned to a site, and may optionally be assigned to a location within that site. Racks can also be organized by user-defined functional roles. The name and facility ID of each rack within a location must be unique.
    /// 
    /// Rack height is measured in rack units (U); racks are commonly between 42U and 48U tall, but NetBox allows you to define racks of arbitrary height. A toggle is provided to indicate whether rack units are in ascending (from the ground up) or descending order.
    /// 
    /// Each rack is assigned a name and (optionally) a separate facility ID. This is helpful when leasing space in a data center your organization does not own: The facility will often assign a seemingly arbitrary ID to a rack (for example, "M204.313") whereas internally you refer to is simply as "R113." A unique serial number and asset tag may also be associated with each rack.
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
    ///     var test = new Netbox.Dcim.Site("test", new()
    ///     {
    ///         Name = "test",
    ///         Status = "active",
    ///     });
    /// 
    ///     var testRack = new Netbox.Dcim.Rack("test", new()
    ///     {
    ///         Name = "test",
    ///         SiteId = test.Id,
    ///         Status = "reserved",
    ///         Width = 19,
    ///         UHeight = 48,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:dcim/rack:Rack")]
    public partial class Rack : global::Pulumi.CustomResource
    {
        [Output("assetTag")]
        public Output<string?> AssetTag { get; private set; } = null!;

        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        [Output("customFields")]
        public Output<ImmutableDictionary<string, string>?> CustomFields { get; private set; } = null!;

        /// <summary>
        /// If rack units are descending. Defaults to `false`.
        /// </summary>
        [Output("descUnits")]
        public Output<bool?> DescUnits { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("facilityId")]
        public Output<string?> FacilityId { get; private set; } = null!;

        /// <summary>
        /// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
        /// </summary>
        [Output("formFactor")]
        public Output<string?> FormFactor { get; private set; } = null!;

        [Output("locationId")]
        public Output<int?> LocationId { get; private set; } = null!;

        [Output("maxWeight")]
        public Output<int?> MaxWeight { get; private set; } = null!;

        [Output("mountingDepth")]
        public Output<int?> MountingDepth { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("outerDepth")]
        public Output<int?> OuterDepth { get; private set; } = null!;

        /// <summary>
        /// Valid values are `mm` and `in`. Required when `outer_width` and `outer_depth` is set.
        /// </summary>
        [Output("outerUnit")]
        public Output<string?> OuterUnit { get; private set; } = null!;

        [Output("outerWidth")]
        public Output<int?> OuterWidth { get; private set; } = null!;

        [Output("roleId")]
        public Output<int?> RoleId { get; private set; } = null!;

        [Output("serial")]
        public Output<string?> Serial { get; private set; } = null!;

        [Output("siteId")]
        public Output<int> SiteId { get; private set; } = null!;

        /// <summary>
        /// Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("tenantId")]
        public Output<int?> TenantId { get; private set; } = null!;

        [Output("uHeight")]
        public Output<int?> UHeight { get; private set; } = null!;

        [Output("weight")]
        public Output<double?> Weight { get; private set; } = null!;

        /// <summary>
        /// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `max_weight` is set.
        /// </summary>
        [Output("weightUnit")]
        public Output<string?> WeightUnit { get; private set; } = null!;

        /// <summary>
        /// Valid values are `10`, `19`, `21` and `23`.
        /// </summary>
        [Output("width")]
        public Output<int?> Width { get; private set; } = null!;


        /// <summary>
        /// Create a Rack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rack(string name, RackArgs args, CustomResourceOptions? options = null)
            : base("netbox:dcim/rack:Rack", name, args ?? new RackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rack(string name, Input<string> id, RackState? state = null, CustomResourceOptions? options = null)
            : base("netbox:dcim/rack:Rack", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Rack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rack Get(string name, Input<string> id, RackState? state = null, CustomResourceOptions? options = null)
        {
            return new Rack(name, id, state, options);
        }
    }

    public sealed class RackArgs : global::Pulumi.ResourceArgs
    {
        [Input("assetTag")]
        public Input<string>? AssetTag { get; set; }

        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        /// <summary>
        /// If rack units are descending. Defaults to `false`.
        /// </summary>
        [Input("descUnits")]
        public Input<bool>? DescUnits { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("facilityId")]
        public Input<string>? FacilityId { get; set; }

        /// <summary>
        /// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
        /// </summary>
        [Input("formFactor")]
        public Input<string>? FormFactor { get; set; }

        [Input("locationId")]
        public Input<int>? LocationId { get; set; }

        [Input("maxWeight")]
        public Input<int>? MaxWeight { get; set; }

        [Input("mountingDepth")]
        public Input<int>? MountingDepth { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("outerDepth")]
        public Input<int>? OuterDepth { get; set; }

        /// <summary>
        /// Valid values are `mm` and `in`. Required when `outer_width` and `outer_depth` is set.
        /// </summary>
        [Input("outerUnit")]
        public Input<string>? OuterUnit { get; set; }

        [Input("outerWidth")]
        public Input<int>? OuterWidth { get; set; }

        [Input("roleId")]
        public Input<int>? RoleId { get; set; }

        [Input("serial")]
        public Input<string>? Serial { get; set; }

        [Input("siteId", required: true)]
        public Input<int> SiteId { get; set; } = null!;

        /// <summary>
        /// Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
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

        [Input("uHeight")]
        public Input<int>? UHeight { get; set; }

        [Input("weight")]
        public Input<double>? Weight { get; set; }

        /// <summary>
        /// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `max_weight` is set.
        /// </summary>
        [Input("weightUnit")]
        public Input<string>? WeightUnit { get; set; }

        /// <summary>
        /// Valid values are `10`, `19`, `21` and `23`.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public RackArgs()
        {
        }
        public static new RackArgs Empty => new RackArgs();
    }

    public sealed class RackState : global::Pulumi.ResourceArgs
    {
        [Input("assetTag")]
        public Input<string>? AssetTag { get; set; }

        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        /// <summary>
        /// If rack units are descending. Defaults to `false`.
        /// </summary>
        [Input("descUnits")]
        public Input<bool>? DescUnits { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("facilityId")]
        public Input<string>? FacilityId { get; set; }

        /// <summary>
        /// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
        /// </summary>
        [Input("formFactor")]
        public Input<string>? FormFactor { get; set; }

        [Input("locationId")]
        public Input<int>? LocationId { get; set; }

        [Input("maxWeight")]
        public Input<int>? MaxWeight { get; set; }

        [Input("mountingDepth")]
        public Input<int>? MountingDepth { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("outerDepth")]
        public Input<int>? OuterDepth { get; set; }

        /// <summary>
        /// Valid values are `mm` and `in`. Required when `outer_width` and `outer_depth` is set.
        /// </summary>
        [Input("outerUnit")]
        public Input<string>? OuterUnit { get; set; }

        [Input("outerWidth")]
        public Input<int>? OuterWidth { get; set; }

        [Input("roleId")]
        public Input<int>? RoleId { get; set; }

        [Input("serial")]
        public Input<string>? Serial { get; set; }

        [Input("siteId")]
        public Input<int>? SiteId { get; set; }

        /// <summary>
        /// Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
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

        [Input("uHeight")]
        public Input<int>? UHeight { get; set; }

        [Input("weight")]
        public Input<double>? Weight { get; set; }

        /// <summary>
        /// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `max_weight` is set.
        /// </summary>
        [Input("weightUnit")]
        public Input<string>? WeightUnit { get; set; }

        /// <summary>
        /// Valid values are `10`, `19`, `21` and `23`.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public RackState()
        {
        }
        public static new RackState Empty => new RackState();
    }
}
