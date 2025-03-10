// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Outputs
{

    [OutputType]
    public sealed class GetTagsTagResult
    {
        public readonly string? Color;
        public readonly string? Description;
        public readonly string Name;
        public readonly string Slug;
        public readonly int TagId;

        [OutputConstructor]
        private GetTagsTagResult(
            string? color,

            string? description,

            string name,

            string slug,

            int tagId)
        {
            Color = color;
            Description = description;
            Name = name;
            Slug = slug;
            TagId = tagId;
        }
    }
}
