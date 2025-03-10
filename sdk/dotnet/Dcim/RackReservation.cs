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
    /// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rackreservation/):
    /// 
    /// &gt; Users can reserve specific units within a rack for future use. An arbitrary set of units within a rack can be associated with a single reservation, but reservations cannot span multiple racks. A description is required for each reservation, reservations may optionally be associated with a specific tenant.
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
    ///         Status = "active",
    ///         Width = 10,
    ///         UHeight = 40,
    ///     });
    /// 
    ///     var testRackReservation = new Netbox.Dcim.RackReservation("test", new()
    ///     {
    ///         RackId = testRack.Id,
    ///         Units = new[]
    ///         {
    ///             1,
    ///             2,
    ///             3,
    ///             4,
    ///             5,
    ///         },
    ///         UserId = 1,
    ///         Description = "my description",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:dcim/rackReservation:RackReservation")]
    public partial class RackReservation : global::Pulumi.CustomResource
    {
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("rackId")]
        public Output<int> RackId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("tenantId")]
        public Output<int?> TenantId { get; private set; } = null!;

        [Output("units")]
        public Output<ImmutableArray<int>> Units { get; private set; } = null!;

        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a RackReservation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RackReservation(string name, RackReservationArgs args, CustomResourceOptions? options = null)
            : base("netbox:dcim/rackReservation:RackReservation", name, args ?? new RackReservationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RackReservation(string name, Input<string> id, RackReservationState? state = null, CustomResourceOptions? options = null)
            : base("netbox:dcim/rackReservation:RackReservation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RackReservation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RackReservation Get(string name, Input<string> id, RackReservationState? state = null, CustomResourceOptions? options = null)
        {
            return new RackReservation(name, id, state, options);
        }
    }

    public sealed class RackReservationArgs : global::Pulumi.ResourceArgs
    {
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("rackId", required: true)]
        public Input<int> RackId { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tenantId")]
        public Input<int>? TenantId { get; set; }

        [Input("units", required: true)]
        private InputList<int>? _units;
        public InputList<int> Units
        {
            get => _units ?? (_units = new InputList<int>());
            set => _units = value;
        }

        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public RackReservationArgs()
        {
        }
        public static new RackReservationArgs Empty => new RackReservationArgs();
    }

    public sealed class RackReservationState : global::Pulumi.ResourceArgs
    {
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("rackId")]
        public Input<int>? RackId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tenantId")]
        public Input<int>? TenantId { get; set; }

        [Input("units")]
        private InputList<int>? _units;
        public InputList<int> Units
        {
            get => _units ?? (_units = new InputList<int>());
            set => _units = value;
        }

        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public RackReservationState()
        {
        }
        public static new RackReservationState Empty => new RackReservationState();
    }
}
