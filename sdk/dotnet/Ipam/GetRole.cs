// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Ipam
{
    public static class GetRole
    {
        public static Task<GetRoleResult> InvokeAsync(GetRoleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRoleResult>("netbox:ipam/getRole:getRole", args ?? new GetRoleArgs(), options.WithDefaults());

        public static Output<GetRoleResult> Invoke(GetRoleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoleResult>("netbox:ipam/getRole:getRole", args ?? new GetRoleInvokeArgs(), options.WithDefaults());

        public static Output<GetRoleResult> Invoke(GetRoleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoleResult>("netbox:ipam/getRole:getRole", args ?? new GetRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRoleArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRoleArgs()
        {
        }
        public static new GetRoleArgs Empty => new GetRoleArgs();
    }

    public sealed class GetRoleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetRoleInvokeArgs()
        {
        }
        public static new GetRoleInvokeArgs Empty => new GetRoleInvokeArgs();
    }


    [OutputType]
    public sealed class GetRoleResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Slug;
        public readonly int Weight;

        [OutputConstructor]
        private GetRoleResult(
            string description,

            string id,

            string name,

            string slug,

            int weight)
        {
            Description = description;
            Id = id;
            Name = name;
            Slug = slug;
            Weight = weight;
        }
    }
}
