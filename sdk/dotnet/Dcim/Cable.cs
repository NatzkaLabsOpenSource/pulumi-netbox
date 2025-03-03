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
    /// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/cable/):
    /// 
    /// &gt; All connections between device components in NetBox are represented using cables. A cable represents a direct physical connection between two sets of endpoints (A and B), such as a console port and a patch panel port, or between two network interfaces.
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
    ///     // assumes that the referenced console port resources exist
    ///     var test = new Netbox.Dcim.Cable("test", new()
    ///     {
    ///         ATerminations = new[]
    ///         {
    ///             new Netbox.Dcim.Inputs.CableATerminationArgs
    ///             {
    ///                 ObjectType = "dcim.consoleserverport",
    ///                 ObjectId = kvm1.Id,
    ///             },
    ///             new Netbox.Dcim.Inputs.CableATerminationArgs
    ///             {
    ///                 ObjectType = "dcim.consoleserverport",
    ///                 ObjectId = kvm2.Id,
    ///             },
    ///         },
    ///         BTerminations = new[]
    ///         {
    ///             new Netbox.Dcim.Inputs.CableBTerminationArgs
    ///             {
    ///                 ObjectType = "dcim.consoleport",
    ///                 ObjectId = server1.Id,
    ///             },
    ///             new Netbox.Dcim.Inputs.CableBTerminationArgs
    ///             {
    ///                 ObjectType = "dcim.consoleport",
    ///                 ObjectId = server2.Id,
    ///             },
    ///         },
    ///         Status = "connected",
    ///         Label = "KVM cable",
    ///         Type = "cat8",
    ///         ColorHex = "123456",
    ///         Length = 10,
    ///         LengthUnit = "m",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:dcim/cable:Cable")]
    public partial class Cable : global::Pulumi.CustomResource
    {
        [Output("aTerminations")]
        public Output<ImmutableArray<Outputs.CableATermination>> ATerminations { get; private set; } = null!;

        [Output("bTerminations")]
        public Output<ImmutableArray<Outputs.CableBTermination>> BTerminations { get; private set; } = null!;

        [Output("colorHex")]
        public Output<string?> ColorHex { get; private set; } = null!;

        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        [Output("customFields")]
        public Output<ImmutableDictionary<string, string>?> CustomFields { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("label")]
        public Output<string?> Label { get; private set; } = null!;

        [Output("length")]
        public Output<double?> Length { get; private set; } = null!;

        /// <summary>
        /// One of [km, m, cm, mi, ft, in]. Required when `length` is set.
        /// </summary>
        [Output("lengthUnit")]
        public Output<string?> LengthUnit { get; private set; } = null!;

        /// <summary>
        /// One of [connected, planned, decommissioning].
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("tenantId")]
        public Output<int?> TenantId { get; private set; } = null!;

        /// <summary>
        /// One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power].
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Cable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cable(string name, CableArgs args, CustomResourceOptions? options = null)
            : base("netbox:dcim/cable:Cable", name, args ?? new CableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cable(string name, Input<string> id, CableState? state = null, CustomResourceOptions? options = null)
            : base("netbox:dcim/cable:Cable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cable Get(string name, Input<string> id, CableState? state = null, CustomResourceOptions? options = null)
        {
            return new Cable(name, id, state, options);
        }
    }

    public sealed class CableArgs : global::Pulumi.ResourceArgs
    {
        [Input("aTerminations", required: true)]
        private InputList<Inputs.CableATerminationArgs>? _aTerminations;
        public InputList<Inputs.CableATerminationArgs> ATerminations
        {
            get => _aTerminations ?? (_aTerminations = new InputList<Inputs.CableATerminationArgs>());
            set => _aTerminations = value;
        }

        [Input("bTerminations", required: true)]
        private InputList<Inputs.CableBTerminationArgs>? _bTerminations;
        public InputList<Inputs.CableBTerminationArgs> BTerminations
        {
            get => _bTerminations ?? (_bTerminations = new InputList<Inputs.CableBTerminationArgs>());
            set => _bTerminations = value;
        }

        [Input("colorHex")]
        public Input<string>? ColorHex { get; set; }

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

        [Input("label")]
        public Input<string>? Label { get; set; }

        [Input("length")]
        public Input<double>? Length { get; set; }

        /// <summary>
        /// One of [km, m, cm, mi, ft, in]. Required when `length` is set.
        /// </summary>
        [Input("lengthUnit")]
        public Input<string>? LengthUnit { get; set; }

        /// <summary>
        /// One of [connected, planned, decommissioning].
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

        /// <summary>
        /// One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power].
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CableArgs()
        {
        }
        public static new CableArgs Empty => new CableArgs();
    }

    public sealed class CableState : global::Pulumi.ResourceArgs
    {
        [Input("aTerminations")]
        private InputList<Inputs.CableATerminationGetArgs>? _aTerminations;
        public InputList<Inputs.CableATerminationGetArgs> ATerminations
        {
            get => _aTerminations ?? (_aTerminations = new InputList<Inputs.CableATerminationGetArgs>());
            set => _aTerminations = value;
        }

        [Input("bTerminations")]
        private InputList<Inputs.CableBTerminationGetArgs>? _bTerminations;
        public InputList<Inputs.CableBTerminationGetArgs> BTerminations
        {
            get => _bTerminations ?? (_bTerminations = new InputList<Inputs.CableBTerminationGetArgs>());
            set => _bTerminations = value;
        }

        [Input("colorHex")]
        public Input<string>? ColorHex { get; set; }

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

        [Input("label")]
        public Input<string>? Label { get; set; }

        [Input("length")]
        public Input<double>? Length { get; set; }

        /// <summary>
        /// One of [km, m, cm, mi, ft, in]. Required when `length` is set.
        /// </summary>
        [Input("lengthUnit")]
        public Input<string>? LengthUnit { get; set; }

        /// <summary>
        /// One of [connected, planned, decommissioning].
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

        /// <summary>
        /// One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power].
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CableState()
        {
        }
        public static new CableState Empty => new CableState();
    }
}
