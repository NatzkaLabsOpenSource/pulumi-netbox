// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Dcim.Inputs
{

    public sealed class GetRegionFilterInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public GetRegionFilterInputArgs()
        {
        }
        public static new GetRegionFilterInputArgs Empty => new GetRegionFilterInputArgs();
    }
}
