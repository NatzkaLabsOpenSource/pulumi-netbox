// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Ipam
{
    public static class GetAsns
    {
        /// <summary>
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
        ///     var asns = Netbox.Ipam.GetAsns.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Netbox.Ipam.Inputs.GetAsnsFilterInputArgs
        ///             {
        ///                 Name = "asn__gte",
        ///                 Value = "1000",
        ///             },
        ///             new Netbox.Ipam.Inputs.GetAsnsFilterInputArgs
        ///             {
        ///                 Name = "asn__lte",
        ///                 Value = "2000",
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAsnsResult> InvokeAsync(GetAsnsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAsnsResult>("netbox:ipam/getAsns:getAsns", args ?? new GetAsnsArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var asns = Netbox.Ipam.GetAsns.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Netbox.Ipam.Inputs.GetAsnsFilterInputArgs
        ///             {
        ///                 Name = "asn__gte",
        ///                 Value = "1000",
        ///             },
        ///             new Netbox.Ipam.Inputs.GetAsnsFilterInputArgs
        ///             {
        ///                 Name = "asn__lte",
        ///                 Value = "2000",
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAsnsResult> Invoke(GetAsnsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAsnsResult>("netbox:ipam/getAsns:getAsns", args ?? new GetAsnsInvokeArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var asns = Netbox.Ipam.GetAsns.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Netbox.Ipam.Inputs.GetAsnsFilterInputArgs
        ///             {
        ///                 Name = "asn__gte",
        ///                 Value = "1000",
        ///             },
        ///             new Netbox.Ipam.Inputs.GetAsnsFilterInputArgs
        ///             {
        ///                 Name = "asn__lte",
        ///                 Value = "2000",
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAsnsResult> Invoke(GetAsnsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAsnsResult>("netbox:ipam/getAsns:getAsns", args ?? new GetAsnsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAsnsArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetAsnsFilterArgs>? _filters;
        public List<Inputs.GetAsnsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetAsnsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Defaults to `0`.
        /// </summary>
        [Input("limit")]
        public int? Limit { get; set; }

        public GetAsnsArgs()
        {
        }
        public static new GetAsnsArgs Empty => new GetAsnsArgs();
    }

    public sealed class GetAsnsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetAsnsFilterInputArgs>? _filters;
        public InputList<Inputs.GetAsnsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetAsnsFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Defaults to `0`.
        /// </summary>
        [Input("limit")]
        public Input<int>? Limit { get; set; }

        public GetAsnsInvokeArgs()
        {
        }
        public static new GetAsnsInvokeArgs Empty => new GetAsnsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAsnsResult
    {
        public readonly ImmutableArray<Outputs.GetAsnsAsnResult> Asns;
        public readonly ImmutableArray<Outputs.GetAsnsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Defaults to `0`.
        /// </summary>
        public readonly int? Limit;

        [OutputConstructor]
        private GetAsnsResult(
            ImmutableArray<Outputs.GetAsnsAsnResult> asns,

            ImmutableArray<Outputs.GetAsnsFilterResult> filters,

            string id,

            int? limit)
        {
            Asns = asns;
            Filters = filters;
            Id = id;
            Limit = limit;
        }
    }
}
