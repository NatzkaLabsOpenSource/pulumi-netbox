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
    /// From the [official documentation](https://docs.netbox.dev/en/stable/features/devices/#device-roles):
    /// 
    /// &gt; Devices can be organized by functional roles, which are fully customizable by the user. For example, you might create roles for core switches, distribution switches, and access switches within your network.
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
    ///     var coreSw = new Netbox.Dcim.DeviceRole("core_sw", new()
    ///     {
    ///         ColorHex = "FF00FF",
    ///         Name = "core-sw",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:dcim/deviceRole:DeviceRole")]
    public partial class DeviceRole : global::Pulumi.CustomResource
    {
        [Output("colorHex")]
        public Output<string> ColorHex { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("vmRole")]
        public Output<bool?> VmRole { get; private set; } = null!;


        /// <summary>
        /// Create a DeviceRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeviceRole(string name, DeviceRoleArgs args, CustomResourceOptions? options = null)
            : base("netbox:dcim/deviceRole:DeviceRole", name, args ?? new DeviceRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeviceRole(string name, Input<string> id, DeviceRoleState? state = null, CustomResourceOptions? options = null)
            : base("netbox:dcim/deviceRole:DeviceRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeviceRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeviceRole Get(string name, Input<string> id, DeviceRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new DeviceRole(name, id, state, options);
        }
    }

    public sealed class DeviceRoleArgs : global::Pulumi.ResourceArgs
    {
        [Input("colorHex", required: true)]
        public Input<string> ColorHex { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("vmRole")]
        public Input<bool>? VmRole { get; set; }

        public DeviceRoleArgs()
        {
        }
        public static new DeviceRoleArgs Empty => new DeviceRoleArgs();
    }

    public sealed class DeviceRoleState : global::Pulumi.ResourceArgs
    {
        [Input("colorHex")]
        public Input<string>? ColorHex { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("vmRole")]
        public Input<bool>? VmRole { get; set; }

        public DeviceRoleState()
        {
        }
        public static new DeviceRoleState Empty => new DeviceRoleState();
    }
}
