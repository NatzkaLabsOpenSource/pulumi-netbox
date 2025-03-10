// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Dcim
{
    public static class GetRegion
    {
        public static Task<GetRegionResult> InvokeAsync(GetRegionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegionResult>("netbox:dcim/getRegion:getRegion", args ?? new GetRegionArgs(), options.WithDefaults());

        public static Output<GetRegionResult> Invoke(GetRegionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegionResult>("netbox:dcim/getRegion:getRegion", args ?? new GetRegionInvokeArgs(), options.WithDefaults());

        public static Output<GetRegionResult> Invoke(GetRegionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegionResult>("netbox:dcim/getRegion:getRegion", args ?? new GetRegionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetRegionFilterArgs>? _filters;
        public List<Inputs.GetRegionFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetRegionFilterArgs>());
            set => _filters = value;
        }

        public GetRegionArgs()
        {
        }
        public static new GetRegionArgs Empty => new GetRegionArgs();
    }

    public sealed class GetRegionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetRegionFilterInputArgs>? _filters;
        public InputList<Inputs.GetRegionFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetRegionFilterInputArgs>());
            set => _filters = value;
        }

        public GetRegionInvokeArgs()
        {
        }
        public static new GetRegionInvokeArgs Empty => new GetRegionInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegionResult
    {
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetRegionFilterResult> Filters;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly int Id;
        public readonly string Name;
        public readonly int ParentRegionId;
        public readonly string Slug;

        [OutputConstructor]
        private GetRegionResult(
            string description,

            ImmutableArray<Outputs.GetRegionFilterResult> filters,

            int id,

            string name,

            int parentRegionId,

            string slug)
        {
            Description = description;
            Filters = filters;
            Id = id;
            Name = name;
            ParentRegionId = parentRegionId;
            Slug = slug;
        }
    }
}
