// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox
{
    /// <summary>
    /// From the [official documentation](https://docs.netbox.dev/en/stable/models/extras/configcontext/):
    /// 
    /// &gt; Context data is made available to devices and/or virtual machines based on their relationships to other objects in NetBox. For example, context data can be associated only with devices assigned to a particular site, or only to virtual machines in a certain cluster.
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
    ///     var test = new Netbox.ConfigContext("test", new()
    ///     {
    ///         Name = "%s",
    ///         Data = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["testkey"] = "testval",
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:index/configContext:ConfigContext")]
    public partial class ConfigContext : global::Pulumi.CustomResource
    {
        [Output("clusterGroups")]
        public Output<ImmutableArray<int>> ClusterGroups { get; private set; } = null!;

        [Output("clusterTypes")]
        public Output<ImmutableArray<int>> ClusterTypes { get; private set; } = null!;

        [Output("clusters")]
        public Output<ImmutableArray<int>> Clusters { get; private set; } = null!;

        [Output("data")]
        public Output<string> Data { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("deviceTypes")]
        public Output<ImmutableArray<int>> DeviceTypes { get; private set; } = null!;

        [Output("locations")]
        public Output<ImmutableArray<int>> Locations { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("platforms")]
        public Output<ImmutableArray<int>> Platforms { get; private set; } = null!;

        [Output("regions")]
        public Output<ImmutableArray<int>> Regions { get; private set; } = null!;

        [Output("roles")]
        public Output<ImmutableArray<int>> Roles { get; private set; } = null!;

        [Output("siteGroups")]
        public Output<ImmutableArray<int>> SiteGroups { get; private set; } = null!;

        [Output("sites")]
        public Output<ImmutableArray<int>> Sites { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("tenantGroups")]
        public Output<ImmutableArray<int>> TenantGroups { get; private set; } = null!;

        [Output("tenants")]
        public Output<ImmutableArray<int>> Tenants { get; private set; } = null!;

        /// <summary>
        /// Defaults to `1000`.
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigContext resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigContext(string name, ConfigContextArgs args, CustomResourceOptions? options = null)
            : base("netbox:index/configContext:ConfigContext", name, args ?? new ConfigContextArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigContext(string name, Input<string> id, ConfigContextState? state = null, CustomResourceOptions? options = null)
            : base("netbox:index/configContext:ConfigContext", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigContext resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigContext Get(string name, Input<string> id, ConfigContextState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigContext(name, id, state, options);
        }
    }

    public sealed class ConfigContextArgs : global::Pulumi.ResourceArgs
    {
        [Input("clusterGroups")]
        private InputList<int>? _clusterGroups;
        public InputList<int> ClusterGroups
        {
            get => _clusterGroups ?? (_clusterGroups = new InputList<int>());
            set => _clusterGroups = value;
        }

        [Input("clusterTypes")]
        private InputList<int>? _clusterTypes;
        public InputList<int> ClusterTypes
        {
            get => _clusterTypes ?? (_clusterTypes = new InputList<int>());
            set => _clusterTypes = value;
        }

        [Input("clusters")]
        private InputList<int>? _clusters;
        public InputList<int> Clusters
        {
            get => _clusters ?? (_clusters = new InputList<int>());
            set => _clusters = value;
        }

        [Input("data", required: true)]
        public Input<string> Data { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("deviceTypes")]
        private InputList<int>? _deviceTypes;
        public InputList<int> DeviceTypes
        {
            get => _deviceTypes ?? (_deviceTypes = new InputList<int>());
            set => _deviceTypes = value;
        }

        [Input("locations")]
        private InputList<int>? _locations;
        public InputList<int> Locations
        {
            get => _locations ?? (_locations = new InputList<int>());
            set => _locations = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platforms")]
        private InputList<int>? _platforms;
        public InputList<int> Platforms
        {
            get => _platforms ?? (_platforms = new InputList<int>());
            set => _platforms = value;
        }

        [Input("regions")]
        private InputList<int>? _regions;
        public InputList<int> Regions
        {
            get => _regions ?? (_regions = new InputList<int>());
            set => _regions = value;
        }

        [Input("roles")]
        private InputList<int>? _roles;
        public InputList<int> Roles
        {
            get => _roles ?? (_roles = new InputList<int>());
            set => _roles = value;
        }

        [Input("siteGroups")]
        private InputList<int>? _siteGroups;
        public InputList<int> SiteGroups
        {
            get => _siteGroups ?? (_siteGroups = new InputList<int>());
            set => _siteGroups = value;
        }

        [Input("sites")]
        private InputList<int>? _sites;
        public InputList<int> Sites
        {
            get => _sites ?? (_sites = new InputList<int>());
            set => _sites = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tenantGroups")]
        private InputList<int>? _tenantGroups;
        public InputList<int> TenantGroups
        {
            get => _tenantGroups ?? (_tenantGroups = new InputList<int>());
            set => _tenantGroups = value;
        }

        [Input("tenants")]
        private InputList<int>? _tenants;
        public InputList<int> Tenants
        {
            get => _tenants ?? (_tenants = new InputList<int>());
            set => _tenants = value;
        }

        /// <summary>
        /// Defaults to `1000`.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ConfigContextArgs()
        {
        }
        public static new ConfigContextArgs Empty => new ConfigContextArgs();
    }

    public sealed class ConfigContextState : global::Pulumi.ResourceArgs
    {
        [Input("clusterGroups")]
        private InputList<int>? _clusterGroups;
        public InputList<int> ClusterGroups
        {
            get => _clusterGroups ?? (_clusterGroups = new InputList<int>());
            set => _clusterGroups = value;
        }

        [Input("clusterTypes")]
        private InputList<int>? _clusterTypes;
        public InputList<int> ClusterTypes
        {
            get => _clusterTypes ?? (_clusterTypes = new InputList<int>());
            set => _clusterTypes = value;
        }

        [Input("clusters")]
        private InputList<int>? _clusters;
        public InputList<int> Clusters
        {
            get => _clusters ?? (_clusters = new InputList<int>());
            set => _clusters = value;
        }

        [Input("data")]
        public Input<string>? Data { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("deviceTypes")]
        private InputList<int>? _deviceTypes;
        public InputList<int> DeviceTypes
        {
            get => _deviceTypes ?? (_deviceTypes = new InputList<int>());
            set => _deviceTypes = value;
        }

        [Input("locations")]
        private InputList<int>? _locations;
        public InputList<int> Locations
        {
            get => _locations ?? (_locations = new InputList<int>());
            set => _locations = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platforms")]
        private InputList<int>? _platforms;
        public InputList<int> Platforms
        {
            get => _platforms ?? (_platforms = new InputList<int>());
            set => _platforms = value;
        }

        [Input("regions")]
        private InputList<int>? _regions;
        public InputList<int> Regions
        {
            get => _regions ?? (_regions = new InputList<int>());
            set => _regions = value;
        }

        [Input("roles")]
        private InputList<int>? _roles;
        public InputList<int> Roles
        {
            get => _roles ?? (_roles = new InputList<int>());
            set => _roles = value;
        }

        [Input("siteGroups")]
        private InputList<int>? _siteGroups;
        public InputList<int> SiteGroups
        {
            get => _siteGroups ?? (_siteGroups = new InputList<int>());
            set => _siteGroups = value;
        }

        [Input("sites")]
        private InputList<int>? _sites;
        public InputList<int> Sites
        {
            get => _sites ?? (_sites = new InputList<int>());
            set => _sites = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tenantGroups")]
        private InputList<int>? _tenantGroups;
        public InputList<int> TenantGroups
        {
            get => _tenantGroups ?? (_tenantGroups = new InputList<int>());
            set => _tenantGroups = value;
        }

        [Input("tenants")]
        private InputList<int>? _tenants;
        public InputList<int> Tenants
        {
            get => _tenants ?? (_tenants = new InputList<int>());
            set => _tenants = value;
        }

        /// <summary>
        /// Defaults to `1000`.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ConfigContextState()
        {
        }
        public static new ConfigContextState Empty => new ConfigContextState();
    }
}
