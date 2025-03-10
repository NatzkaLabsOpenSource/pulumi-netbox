// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Tenancy
{
    public static class GetContactRole
    {
        public static Task<GetContactRoleResult> InvokeAsync(GetContactRoleArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContactRoleResult>("netbox:tenancy/getContactRole:getContactRole", args ?? new GetContactRoleArgs(), options.WithDefaults());

        public static Output<GetContactRoleResult> Invoke(GetContactRoleInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactRoleResult>("netbox:tenancy/getContactRole:getContactRole", args ?? new GetContactRoleInvokeArgs(), options.WithDefaults());

        public static Output<GetContactRoleResult> Invoke(GetContactRoleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactRoleResult>("netbox:tenancy/getContactRole:getContactRole", args ?? new GetContactRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContactRoleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// At least one of `name` or `slug` must be given.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// At least one of `name` or `slug` must be given.
        /// </summary>
        [Input("slug")]
        public string? Slug { get; set; }

        public GetContactRoleArgs()
        {
        }
        public static new GetContactRoleArgs Empty => new GetContactRoleArgs();
    }

    public sealed class GetContactRoleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// At least one of `name` or `slug` must be given.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// At least one of `name` or `slug` must be given.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public GetContactRoleInvokeArgs()
        {
        }
        public static new GetContactRoleInvokeArgs Empty => new GetContactRoleInvokeArgs();
    }


    [OutputType]
    public sealed class GetContactRoleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// At least one of `name` or `slug` must be given.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// At least one of `name` or `slug` must be given.
        /// </summary>
        public readonly string Slug;

        [OutputConstructor]
        private GetContactRoleResult(
            string id,

            string name,

            string slug)
        {
            Id = id;
            Name = name;
            Slug = slug;
        }
    }
}
