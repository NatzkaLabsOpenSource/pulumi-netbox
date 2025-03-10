// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Ipam
{
    public static class GetAsn
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
        ///     var asn1 = Netbox.Ipam.GetAsn.Invoke(new()
        ///     {
        ///         Asn = "1111",
        ///         Tag = "tag-1",
        ///     });
        /// 
        ///     var asn2 = Netbox.Ipam.GetAsn.Invoke(new()
        ///     {
        ///         Tag = "tag-1",
        ///         TagN = "tag-2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAsnResult> InvokeAsync(GetAsnArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAsnResult>("netbox:ipam/getAsn:getAsn", args ?? new GetAsnArgs(), options.WithDefaults());

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
        ///     var asn1 = Netbox.Ipam.GetAsn.Invoke(new()
        ///     {
        ///         Asn = "1111",
        ///         Tag = "tag-1",
        ///     });
        /// 
        ///     var asn2 = Netbox.Ipam.GetAsn.Invoke(new()
        ///     {
        ///         Tag = "tag-1",
        ///         TagN = "tag-2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAsnResult> Invoke(GetAsnInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAsnResult>("netbox:ipam/getAsn:getAsn", args ?? new GetAsnInvokeArgs(), options.WithDefaults());

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
        ///     var asn1 = Netbox.Ipam.GetAsn.Invoke(new()
        ///     {
        ///         Asn = "1111",
        ///         Tag = "tag-1",
        ///     });
        /// 
        ///     var asn2 = Netbox.Ipam.GetAsn.Invoke(new()
        ///     {
        ///         Tag = "tag-1",
        ///         TagN = "tag-2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAsnResult> Invoke(GetAsnInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAsnResult>("netbox:ipam/getAsn:getAsn", args ?? new GetAsnInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAsnArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// At least one of `asn` or `tag` must be given.
        /// </summary>
        [Input("asn")]
        public string? Asn { get; set; }

        /// <summary>
        /// Tag to include in the data source filter (must match the tag's slug). At least one of `asn` or `tag` must be given.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        /// <summary>
        /// Tag to exclude from the data source filter (must match the tag's slug).
        /// Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
        /// for more information on available lookup expressions.
        /// </summary>
        [Input("tagN")]
        public string? TagN { get; set; }

        public GetAsnArgs()
        {
        }
        public static new GetAsnArgs Empty => new GetAsnArgs();
    }

    public sealed class GetAsnInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// At least one of `asn` or `tag` must be given.
        /// </summary>
        [Input("asn")]
        public Input<string>? Asn { get; set; }

        /// <summary>
        /// Tag to include in the data source filter (must match the tag's slug). At least one of `asn` or `tag` must be given.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// Tag to exclude from the data source filter (must match the tag's slug).
        /// Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
        /// for more information on available lookup expressions.
        /// </summary>
        [Input("tagN")]
        public Input<string>? TagN { get; set; }

        public GetAsnInvokeArgs()
        {
        }
        public static new GetAsnInvokeArgs Empty => new GetAsnInvokeArgs();
    }


    [OutputType]
    public sealed class GetAsnResult
    {
        /// <summary>
        /// At least one of `asn` or `tag` must be given.
        /// </summary>
        public readonly string? Asn;
        public readonly string Description;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Tag to include in the data source filter (must match the tag's slug). At least one of `asn` or `tag` must be given.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// Tag to exclude from the data source filter (must match the tag's slug).
        /// Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
        /// for more information on available lookup expressions.
        /// </summary>
        public readonly string? TagN;
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetAsnResult(
            string? asn,

            string description,

            int id,

            string? tag,

            string? tagN,

            ImmutableArray<string> tags)
        {
            Asn = asn;
            Description = description;
            Id = id;
            Tag = tag;
            TagN = tagN;
            Tags = tags;
        }
    }
}
