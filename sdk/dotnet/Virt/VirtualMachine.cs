// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Virt
{
    /// <summary>
    /// From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#virtual-machines):
    /// 
    /// &gt; A virtual machine is a virtualized compute instance. These behave in NetBox very similarly to device objects, but without any physical attributes. For example, a VM may have interfaces assigned to it with IP addresses and VLANs, however its interfaces cannot be connected via cables (because they are virtual). Each VM may also define its compute, memory, and storage resources as well.
    /// </summary>
    [NetboxResourceType("netbox:virt/virtualMachine:VirtualMachine")]
    public partial class VirtualMachine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// At least one of `site_id` or `cluster_id` must be given.
        /// </summary>
        [Output("clusterId")]
        public Output<int?> ClusterId { get; private set; } = null!;

        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        [Output("customFields")]
        public Output<ImmutableDictionary<string, string>?> CustomFields { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("deviceId")]
        public Output<int?> DeviceId { get; private set; } = null!;

        [Output("diskSizeMb")]
        public Output<int> DiskSizeMb { get; private set; } = null!;

        /// <summary>
        /// This is best managed through the use of `jsonencode` and a map of settings.
        /// </summary>
        [Output("localContextData")]
        public Output<string?> LocalContextData { get; private set; } = null!;

        [Output("memoryMb")]
        public Output<int?> MemoryMb { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("platformId")]
        public Output<int?> PlatformId { get; private set; } = null!;

        [Output("primaryIpv4")]
        public Output<int> PrimaryIpv4 { get; private set; } = null!;

        [Output("primaryIpv6")]
        public Output<int> PrimaryIpv6 { get; private set; } = null!;

        [Output("roleId")]
        public Output<int?> RoleId { get; private set; } = null!;

        /// <summary>
        /// At least one of `site_id` or `cluster_id` must be given.
        /// </summary>
        [Output("siteId")]
        public Output<int?> SiteId { get; private set; } = null!;

        /// <summary>
        /// Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("tenantId")]
        public Output<int?> TenantId { get; private set; } = null!;

        [Output("vcpus")]
        public Output<double?> Vcpus { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachine(string name, VirtualMachineArgs? args = null, CustomResourceOptions? options = null)
            : base("netbox:virt/virtualMachine:VirtualMachine", name, args ?? new VirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachine(string name, Input<string> id, VirtualMachineState? state = null, CustomResourceOptions? options = null)
            : base("netbox:virt/virtualMachine:VirtualMachine", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachine Get(string name, Input<string> id, VirtualMachineState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualMachine(name, id, state, options);
        }
    }

    public sealed class VirtualMachineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// At least one of `site_id` or `cluster_id` must be given.
        /// </summary>
        [Input("clusterId")]
        public Input<int>? ClusterId { get; set; }

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

        [Input("deviceId")]
        public Input<int>? DeviceId { get; set; }

        [Input("diskSizeMb")]
        public Input<int>? DiskSizeMb { get; set; }

        /// <summary>
        /// This is best managed through the use of `jsonencode` and a map of settings.
        /// </summary>
        [Input("localContextData")]
        public Input<string>? LocalContextData { get; set; }

        [Input("memoryMb")]
        public Input<int>? MemoryMb { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platformId")]
        public Input<int>? PlatformId { get; set; }

        [Input("roleId")]
        public Input<int>? RoleId { get; set; }

        /// <summary>
        /// At least one of `site_id` or `cluster_id` must be given.
        /// </summary>
        [Input("siteId")]
        public Input<int>? SiteId { get; set; }

        /// <summary>
        /// Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
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

        [Input("vcpus")]
        public Input<double>? Vcpus { get; set; }

        public VirtualMachineArgs()
        {
        }
        public static new VirtualMachineArgs Empty => new VirtualMachineArgs();
    }

    public sealed class VirtualMachineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// At least one of `site_id` or `cluster_id` must be given.
        /// </summary>
        [Input("clusterId")]
        public Input<int>? ClusterId { get; set; }

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

        [Input("deviceId")]
        public Input<int>? DeviceId { get; set; }

        [Input("diskSizeMb")]
        public Input<int>? DiskSizeMb { get; set; }

        /// <summary>
        /// This is best managed through the use of `jsonencode` and a map of settings.
        /// </summary>
        [Input("localContextData")]
        public Input<string>? LocalContextData { get; set; }

        [Input("memoryMb")]
        public Input<int>? MemoryMb { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platformId")]
        public Input<int>? PlatformId { get; set; }

        [Input("primaryIpv4")]
        public Input<int>? PrimaryIpv4 { get; set; }

        [Input("primaryIpv6")]
        public Input<int>? PrimaryIpv6 { get; set; }

        [Input("roleId")]
        public Input<int>? RoleId { get; set; }

        /// <summary>
        /// At least one of `site_id` or `cluster_id` must be given.
        /// </summary>
        [Input("siteId")]
        public Input<int>? SiteId { get; set; }

        /// <summary>
        /// Valid values are `offline`, `active`, `planned`, `staged`, `failed` and `decommissioning`. Defaults to `active`.
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

        [Input("vcpus")]
        public Input<double>? Vcpus { get; set; }

        public VirtualMachineState()
        {
        }
        public static new VirtualMachineState Empty => new VirtualMachineState();
    }
}
