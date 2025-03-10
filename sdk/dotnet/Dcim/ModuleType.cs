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
    /// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/moduletype/):
    /// 
    /// &gt; A module type represents a specific make and model of hardware component which is installable within a device's module bay and has its own child components. For example, consider a chassis-based switch or router with a number of field-replaceable line cards. Each line card has its own model number and includes a certain set of components such as interfaces. Each module type may have a manufacturer, model number, and part number assigned to it.
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
    ///     var test = new Netbox.Dcim.Manufacturer("test", new()
    ///     {
    ///         Name = "Dell",
    ///     });
    /// 
    ///     var testModuleType = new Netbox.Dcim.ModuleType("test", new()
    ///     {
    ///         ManufacturerId = test.Id,
    ///         Model = "Networking",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:dcim/moduleType:ModuleType")]
    public partial class ModuleType : global::Pulumi.CustomResource
    {
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        [Output("customFields")]
        public Output<ImmutableDictionary<string, string>?> CustomFields { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("manufacturerId")]
        public Output<int> ManufacturerId { get; private set; } = null!;

        [Output("model")]
        public Output<string> Model { get; private set; } = null!;

        [Output("partNumber")]
        public Output<string?> PartNumber { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("weight")]
        public Output<double?> Weight { get; private set; } = null!;

        /// <summary>
        /// One of [kg, g, lb, oz]. Required when `weight` is set.
        /// </summary>
        [Output("weightUnit")]
        public Output<string?> WeightUnit { get; private set; } = null!;


        /// <summary>
        /// Create a ModuleType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModuleType(string name, ModuleTypeArgs args, CustomResourceOptions? options = null)
            : base("netbox:dcim/moduleType:ModuleType", name, args ?? new ModuleTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModuleType(string name, Input<string> id, ModuleTypeState? state = null, CustomResourceOptions? options = null)
            : base("netbox:dcim/moduleType:ModuleType", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ModuleType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModuleType Get(string name, Input<string> id, ModuleTypeState? state = null, CustomResourceOptions? options = null)
        {
            return new ModuleType(name, id, state, options);
        }
    }

    public sealed class ModuleTypeArgs : global::Pulumi.ResourceArgs
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

        [Input("manufacturerId", required: true)]
        public Input<int> ManufacturerId { get; set; } = null!;

        [Input("model", required: true)]
        public Input<string> Model { get; set; } = null!;

        [Input("partNumber")]
        public Input<string>? PartNumber { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("weight")]
        public Input<double>? Weight { get; set; }

        /// <summary>
        /// One of [kg, g, lb, oz]. Required when `weight` is set.
        /// </summary>
        [Input("weightUnit")]
        public Input<string>? WeightUnit { get; set; }

        public ModuleTypeArgs()
        {
        }
        public static new ModuleTypeArgs Empty => new ModuleTypeArgs();
    }

    public sealed class ModuleTypeState : global::Pulumi.ResourceArgs
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

        [Input("manufacturerId")]
        public Input<int>? ManufacturerId { get; set; }

        [Input("model")]
        public Input<string>? Model { get; set; }

        [Input("partNumber")]
        public Input<string>? PartNumber { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("weight")]
        public Input<double>? Weight { get; set; }

        /// <summary>
        /// One of [kg, g, lb, oz]. Required when `weight` is set.
        /// </summary>
        [Input("weightUnit")]
        public Input<string>? WeightUnit { get; set; }

        public ModuleTypeState()
        {
        }
        public static new ModuleTypeState Empty => new ModuleTypeState();
    }
}
