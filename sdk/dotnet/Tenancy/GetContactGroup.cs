// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Tenancy
{
    public static class GetContactGroup
    {
        public static Task<GetContactGroupResult> InvokeAsync(GetContactGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContactGroupResult>("netbox:tenancy/getContactGroup:getContactGroup", args ?? new GetContactGroupArgs(), options.WithDefaults());

        public static Output<GetContactGroupResult> Invoke(GetContactGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactGroupResult>("netbox:tenancy/getContactGroup:getContactGroup", args ?? new GetContactGroupInvokeArgs(), options.WithDefaults());

        public static Output<GetContactGroupResult> Invoke(GetContactGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactGroupResult>("netbox:tenancy/getContactGroup:getContactGroup", args ?? new GetContactGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContactGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetContactGroupArgs()
        {
        }
        public static new GetContactGroupArgs Empty => new GetContactGroupArgs();
    }

    public sealed class GetContactGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetContactGroupInvokeArgs()
        {
        }
        public static new GetContactGroupInvokeArgs Empty => new GetContactGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetContactGroupResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly int ParentId;
        public readonly string Slug;

        [OutputConstructor]
        private GetContactGroupResult(
            string description,

            string id,

            string name,

            int parentId,

            string slug)
        {
            Description = description;
            Id = id;
            Name = name;
            ParentId = parentId;
            Slug = slug;
        }
    }
}
