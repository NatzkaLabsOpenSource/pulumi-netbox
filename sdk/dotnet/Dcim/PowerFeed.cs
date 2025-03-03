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
    /// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/powerfeed/):
    /// 
    /// &gt; A power feed represents the distribution of power from a power panel to a particular device, typically a power distribution unit (PDU). The power port (inlet) on a device can be connected via a cable to a power feed. A power feed may optionally be assigned to a rack to allow more easily tracking the distribution of power among racks.
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
    ///         Name = "Site 1",
    ///         Status = "active",
    ///     });
    /// 
    ///     var testLocation = new Netbox.Dcim.Location("test", new()
    ///     {
    ///         Name = "Location 1",
    ///         SiteId = test.Id,
    ///     });
    /// 
    ///     var testPowerPanel = new Netbox.Dcim.PowerPanel("test", new()
    ///     {
    ///         Name = "Power Panel 1",
    ///         SiteId = test.Id,
    ///         LocationId = testLocation.Id,
    ///     });
    /// 
    ///     var testPowerFeed = new Netbox.Dcim.PowerFeed("test", new()
    ///     {
    ///         PowerPanelId = testPowerPanel.Id,
    ///         Name = "Power Feed 1",
    ///         Status = "active",
    ///         Type = "primary",
    ///         Supply = "ac",
    ///         Phase = "single-phase",
    ///         Voltage = 250,
    ///         Amperage = 100,
    ///         MaxPercentUtilization = 80,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:dcim/powerFeed:PowerFeed")]
    public partial class PowerFeed : global::Pulumi.CustomResource
    {
        [Output("amperage")]
        public Output<int> Amperage { get; private set; } = null!;

        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        [Output("customFields")]
        public Output<ImmutableDictionary<string, string>?> CustomFields { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Output("markConnected")]
        public Output<bool?> MarkConnected { get; private set; } = null!;

        [Output("maxPercentUtilization")]
        public Output<int> MaxPercentUtilization { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One of [single-phase, three-phase].
        /// </summary>
        [Output("phase")]
        public Output<string> Phase { get; private set; } = null!;

        [Output("powerPanelId")]
        public Output<int> PowerPanelId { get; private set; } = null!;

        [Output("rackId")]
        public Output<int?> RackId { get; private set; } = null!;

        /// <summary>
        /// One of [offline, active, planned, failed].
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// One of [ac, dc].
        /// </summary>
        [Output("supply")]
        public Output<string> Supply { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// One of [primary, redundant].
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("voltage")]
        public Output<int> Voltage { get; private set; } = null!;


        /// <summary>
        /// Create a PowerFeed resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PowerFeed(string name, PowerFeedArgs args, CustomResourceOptions? options = null)
            : base("netbox:dcim/powerFeed:PowerFeed", name, args ?? new PowerFeedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PowerFeed(string name, Input<string> id, PowerFeedState? state = null, CustomResourceOptions? options = null)
            : base("netbox:dcim/powerFeed:PowerFeed", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PowerFeed resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PowerFeed Get(string name, Input<string> id, PowerFeedState? state = null, CustomResourceOptions? options = null)
        {
            return new PowerFeed(name, id, state, options);
        }
    }

    public sealed class PowerFeedArgs : global::Pulumi.ResourceArgs
    {
        [Input("amperage", required: true)]
        public Input<int> Amperage { get; set; } = null!;

        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("markConnected")]
        public Input<bool>? MarkConnected { get; set; }

        [Input("maxPercentUtilization", required: true)]
        public Input<int> MaxPercentUtilization { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// One of [single-phase, three-phase].
        /// </summary>
        [Input("phase", required: true)]
        public Input<string> Phase { get; set; } = null!;

        [Input("powerPanelId", required: true)]
        public Input<int> PowerPanelId { get; set; } = null!;

        [Input("rackId")]
        public Input<int>? RackId { get; set; }

        /// <summary>
        /// One of [offline, active, planned, failed].
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// One of [ac, dc].
        /// </summary>
        [Input("supply", required: true)]
        public Input<string> Supply { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// One of [primary, redundant].
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("voltage", required: true)]
        public Input<int> Voltage { get; set; } = null!;

        public PowerFeedArgs()
        {
        }
        public static new PowerFeedArgs Empty => new PowerFeedArgs();
    }

    public sealed class PowerFeedState : global::Pulumi.ResourceArgs
    {
        [Input("amperage")]
        public Input<int>? Amperage { get; set; }

        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Defaults to `false`.
        /// </summary>
        [Input("markConnected")]
        public Input<bool>? MarkConnected { get; set; }

        [Input("maxPercentUtilization")]
        public Input<int>? MaxPercentUtilization { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// One of [single-phase, three-phase].
        /// </summary>
        [Input("phase")]
        public Input<string>? Phase { get; set; }

        [Input("powerPanelId")]
        public Input<int>? PowerPanelId { get; set; }

        [Input("rackId")]
        public Input<int>? RackId { get; set; }

        /// <summary>
        /// One of [offline, active, planned, failed].
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// One of [ac, dc].
        /// </summary>
        [Input("supply")]
        public Input<string>? Supply { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// One of [primary, redundant].
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("voltage")]
        public Input<int>? Voltage { get; set; }

        public PowerFeedState()
        {
        }
        public static new PowerFeedState Empty => new PowerFeedState();
    }
}
