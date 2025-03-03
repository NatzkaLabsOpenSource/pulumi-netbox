// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox
{
    public static class GetConfigContext
    {
        public static Task<GetConfigContextResult> InvokeAsync(GetConfigContextArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigContextResult>("netbox:index/getConfigContext:getConfigContext", args ?? new GetConfigContextArgs(), options.WithDefaults());

        public static Output<GetConfigContextResult> Invoke(GetConfigContextInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigContextResult>("netbox:index/getConfigContext:getConfigContext", args ?? new GetConfigContextInvokeArgs(), options.WithDefaults());

        public static Output<GetConfigContextResult> Invoke(GetConfigContextInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigContextResult>("netbox:index/getConfigContext:getConfigContext", args ?? new GetConfigContextInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigContextArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetConfigContextArgs()
        {
        }
        public static new GetConfigContextArgs Empty => new GetConfigContextArgs();
    }

    public sealed class GetConfigContextInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetConfigContextInvokeArgs()
        {
        }
        public static new GetConfigContextInvokeArgs Empty => new GetConfigContextInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigContextResult
    {
        public readonly ImmutableArray<int> ClusterGroups;
        public readonly ImmutableArray<int> ClusterTypes;
        public readonly ImmutableArray<int> Clusters;
        public readonly string Data;
        public readonly string Description;
        public readonly ImmutableArray<int> DeviceTypes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<int> Locations;
        public readonly string Name;
        public readonly ImmutableArray<int> Platforms;
        public readonly ImmutableArray<int> Regions;
        public readonly ImmutableArray<int> Roles;
        public readonly ImmutableArray<int> SiteGroups;
        public readonly ImmutableArray<int> Sites;
        public readonly ImmutableArray<string> Tags;
        public readonly ImmutableArray<int> TenantGroups;
        public readonly ImmutableArray<int> Tenants;
        public readonly int Weight;

        [OutputConstructor]
        private GetConfigContextResult(
            ImmutableArray<int> clusterGroups,

            ImmutableArray<int> clusterTypes,

            ImmutableArray<int> clusters,

            string data,

            string description,

            ImmutableArray<int> deviceTypes,

            string id,

            ImmutableArray<int> locations,

            string name,

            ImmutableArray<int> platforms,

            ImmutableArray<int> regions,

            ImmutableArray<int> roles,

            ImmutableArray<int> siteGroups,

            ImmutableArray<int> sites,

            ImmutableArray<string> tags,

            ImmutableArray<int> tenantGroups,

            ImmutableArray<int> tenants,

            int weight)
        {
            ClusterGroups = clusterGroups;
            ClusterTypes = clusterTypes;
            Clusters = clusters;
            Data = data;
            Description = description;
            DeviceTypes = deviceTypes;
            Id = id;
            Locations = locations;
            Name = name;
            Platforms = platforms;
            Regions = regions;
            Roles = roles;
            SiteGroups = siteGroups;
            Sites = sites;
            Tags = tags;
            TenantGroups = tenantGroups;
            Tenants = tenants;
            Weight = weight;
        }
    }
}
