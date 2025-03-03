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
    /// From the [official documentation](https://docs.netbox.dev/en/stable/features/devices/#devices):
    /// 
    /// &gt; Every piece of hardware which is installed within a site or rack exists in NetBox as a device. Devices are measured in rack units (U) and can be half depth or full depth. A device may have a height of 0U: These devices do not consume vertical rack space and cannot be assigned to a particular rack unit. A common example of a 0U device is a vertically-mounted PDU.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Netbox = Pulumi.Netbox;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Netbox.Dcim.Site("test", new()
    ///     {
    ///         Name = "%[1]s",
    ///     });
    /// 
    ///     var testDeviceRole = new Netbox.Dcim.DeviceRole("test", new()
    ///     {
    ///         Name = "%[1]s",
    ///         ColorHex = "123456",
    ///     });
    /// 
    ///     var testManufacturer = new Netbox.Dcim.Manufacturer("test", new()
    ///     {
    ///         Name = "test",
    ///     });
    /// 
    ///     var testDeviceType = new Netbox.Dcim.DeviceType("test", new()
    ///     {
    ///         Model = "test",
    ///         ManufacturerId = testManufacturer.Id,
    ///     });
    /// 
    ///     var testDevice = new Netbox.Dcim.Device("test", new()
    ///     {
    ///         Name = "%[1]s",
    ///         DeviceTypeId = testDeviceType.Id,
    ///         RoleId = testDeviceRole.Id,
    ///         SiteId = test.Id,
    ///         LocalContextData = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["setting_a"] = "Some Setting",
    ///             ["setting_b"] = 42,
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:dcim/device:Device")]
    public partial class Device : global::Pulumi.CustomResource
    {
        [Output("assetTag")]
        public Output<string?> AssetTag { get; private set; } = null!;

        [Output("clusterId")]
        public Output<int?> ClusterId { get; private set; } = null!;

        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        [Output("configTemplateId")]
        public Output<int?> ConfigTemplateId { get; private set; } = null!;

        [Output("customFields")]
        public Output<ImmutableDictionary<string, string>?> CustomFields { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("deviceTypeId")]
        public Output<int> DeviceTypeId { get; private set; } = null!;

        /// <summary>
        /// This is best managed through the use of `jsonencode` and a map of settings.
        /// </summary>
        [Output("localContextData")]
        public Output<string?> LocalContextData { get; private set; } = null!;

        [Output("locationId")]
        public Output<int?> LocationId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("platformId")]
        public Output<int?> PlatformId { get; private set; } = null!;

        [Output("primaryIpv4")]
        public Output<int> PrimaryIpv4 { get; private set; } = null!;

        [Output("primaryIpv6")]
        public Output<int> PrimaryIpv6 { get; private set; } = null!;

        /// <summary>
        /// Valid values are `front` and `rear`. Required when `rack_position` is set.
        /// </summary>
        [Output("rackFace")]
        public Output<string?> RackFace { get; private set; } = null!;

        [Output("rackId")]
        public Output<int?> RackId { get; private set; } = null!;

        [Output("rackPosition")]
        public Output<double?> RackPosition { get; private set; } = null!;

        [Output("roleId")]
        public Output<int> RoleId { get; private set; } = null!;

        [Output("serial")]
        public Output<string?> Serial { get; private set; } = null!;

        [Output("siteId")]
        public Output<int> SiteId { get; private set; } = null!;

        /// <summary>
        /// Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `inventory`. Defaults to `active`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("tenantId")]
        public Output<int?> TenantId { get; private set; } = null!;

        /// <summary>
        /// Required when `virtual_chassis_master` and `virtual_chassis_id` is set.
        /// </summary>
        [Output("virtualChassisId")]
        public Output<int?> VirtualChassisId { get; private set; } = null!;

        /// <summary>
        /// Required when `virtual_chassis_master` and `virtual_chassis_id` is set.
        /// </summary>
        [Output("virtualChassisMaster")]
        public Output<bool?> VirtualChassisMaster { get; private set; } = null!;

        [Output("virtualChassisPosition")]
        public Output<int?> VirtualChassisPosition { get; private set; } = null!;

        [Output("virtualChassisPriority")]
        public Output<int?> VirtualChassisPriority { get; private set; } = null!;


        /// <summary>
        /// Create a Device resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Device(string name, DeviceArgs args, CustomResourceOptions? options = null)
            : base("netbox:dcim/device:Device", name, args ?? new DeviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Device(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
            : base("netbox:dcim/device:Device", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Device resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Device Get(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
        {
            return new Device(name, id, state, options);
        }
    }

    public sealed class DeviceArgs : global::Pulumi.ResourceArgs
    {
        [Input("assetTag")]
        public Input<string>? AssetTag { get; set; }

        [Input("clusterId")]
        public Input<int>? ClusterId { get; set; }

        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("configTemplateId")]
        public Input<int>? ConfigTemplateId { get; set; }

        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("deviceTypeId", required: true)]
        public Input<int> DeviceTypeId { get; set; } = null!;

        /// <summary>
        /// This is best managed through the use of `jsonencode` and a map of settings.
        /// </summary>
        [Input("localContextData")]
        public Input<string>? LocalContextData { get; set; }

        [Input("locationId")]
        public Input<int>? LocationId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platformId")]
        public Input<int>? PlatformId { get; set; }

        /// <summary>
        /// Valid values are `front` and `rear`. Required when `rack_position` is set.
        /// </summary>
        [Input("rackFace")]
        public Input<string>? RackFace { get; set; }

        [Input("rackId")]
        public Input<int>? RackId { get; set; }

        [Input("rackPosition")]
        public Input<double>? RackPosition { get; set; }

        [Input("roleId", required: true)]
        public Input<int> RoleId { get; set; } = null!;

        [Input("serial")]
        public Input<string>? Serial { get; set; }

        [Input("siteId", required: true)]
        public Input<int> SiteId { get; set; } = null!;

        /// <summary>
        /// Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `inventory`. Defaults to `active`.
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
        /// Required when `virtual_chassis_master` and `virtual_chassis_id` is set.
        /// </summary>
        [Input("virtualChassisId")]
        public Input<int>? VirtualChassisId { get; set; }

        /// <summary>
        /// Required when `virtual_chassis_master` and `virtual_chassis_id` is set.
        /// </summary>
        [Input("virtualChassisMaster")]
        public Input<bool>? VirtualChassisMaster { get; set; }

        [Input("virtualChassisPosition")]
        public Input<int>? VirtualChassisPosition { get; set; }

        [Input("virtualChassisPriority")]
        public Input<int>? VirtualChassisPriority { get; set; }

        public DeviceArgs()
        {
        }
        public static new DeviceArgs Empty => new DeviceArgs();
    }

    public sealed class DeviceState : global::Pulumi.ResourceArgs
    {
        [Input("assetTag")]
        public Input<string>? AssetTag { get; set; }

        [Input("clusterId")]
        public Input<int>? ClusterId { get; set; }

        [Input("comments")]
        public Input<string>? Comments { get; set; }

        [Input("configTemplateId")]
        public Input<int>? ConfigTemplateId { get; set; }

        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("deviceTypeId")]
        public Input<int>? DeviceTypeId { get; set; }

        /// <summary>
        /// This is best managed through the use of `jsonencode` and a map of settings.
        /// </summary>
        [Input("localContextData")]
        public Input<string>? LocalContextData { get; set; }

        [Input("locationId")]
        public Input<int>? LocationId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platformId")]
        public Input<int>? PlatformId { get; set; }

        [Input("primaryIpv4")]
        public Input<int>? PrimaryIpv4 { get; set; }

        [Input("primaryIpv6")]
        public Input<int>? PrimaryIpv6 { get; set; }

        /// <summary>
        /// Valid values are `front` and `rear`. Required when `rack_position` is set.
        /// </summary>
        [Input("rackFace")]
        public Input<string>? RackFace { get; set; }

        [Input("rackId")]
        public Input<int>? RackId { get; set; }

        [Input("rackPosition")]
        public Input<double>? RackPosition { get; set; }

        [Input("roleId")]
        public Input<int>? RoleId { get; set; }

        [Input("serial")]
        public Input<string>? Serial { get; set; }

        [Input("siteId")]
        public Input<int>? SiteId { get; set; }

        /// <summary>
        /// Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `inventory`. Defaults to `active`.
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
        /// Required when `virtual_chassis_master` and `virtual_chassis_id` is set.
        /// </summary>
        [Input("virtualChassisId")]
        public Input<int>? VirtualChassisId { get; set; }

        /// <summary>
        /// Required when `virtual_chassis_master` and `virtual_chassis_id` is set.
        /// </summary>
        [Input("virtualChassisMaster")]
        public Input<bool>? VirtualChassisMaster { get; set; }

        [Input("virtualChassisPosition")]
        public Input<int>? VirtualChassisPosition { get; set; }

        [Input("virtualChassisPriority")]
        public Input<int>? VirtualChassisPriority { get; set; }

        public DeviceState()
        {
        }
        public static new DeviceState Empty => new DeviceState();
    }
}
