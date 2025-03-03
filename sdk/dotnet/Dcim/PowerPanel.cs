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
    /// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/powerpanel/):
    /// 
    /// &gt; A power panel represents the origin point in NetBox for electrical power being disseminated by one or more power feeds. In a data center environment, one power panel often serves a group of racks, with an individual power feed extending to each rack, though this is not always the case. It is common to have two sets of panels and feeds arranged in parallel to provide redundant power to each rack.
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
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:dcim/powerPanel:PowerPanel")]
    public partial class PowerPanel : global::Pulumi.CustomResource
    {
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        [Output("customFields")]
        public Output<ImmutableDictionary<string, string>?> CustomFields { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("locationId")]
        public Output<int?> LocationId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("siteId")]
        public Output<int> SiteId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a PowerPanel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PowerPanel(string name, PowerPanelArgs args, CustomResourceOptions? options = null)
            : base("netbox:dcim/powerPanel:PowerPanel", name, args ?? new PowerPanelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PowerPanel(string name, Input<string> id, PowerPanelState? state = null, CustomResourceOptions? options = null)
            : base("netbox:dcim/powerPanel:PowerPanel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PowerPanel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PowerPanel Get(string name, Input<string> id, PowerPanelState? state = null, CustomResourceOptions? options = null)
        {
            return new PowerPanel(name, id, state, options);
        }
    }

    public sealed class PowerPanelArgs : global::Pulumi.ResourceArgs
    {
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

        [Input("locationId")]
        public Input<int>? LocationId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("siteId", required: true)]
        public Input<int> SiteId { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public PowerPanelArgs()
        {
        }
        public static new PowerPanelArgs Empty => new PowerPanelArgs();
    }

    public sealed class PowerPanelState : global::Pulumi.ResourceArgs
    {
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

        [Input("locationId")]
        public Input<int>? LocationId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("siteId")]
        public Input<int>? SiteId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public PowerPanelState()
        {
        }
        public static new PowerPanelState Empty => new PowerPanelState();
    }
}
