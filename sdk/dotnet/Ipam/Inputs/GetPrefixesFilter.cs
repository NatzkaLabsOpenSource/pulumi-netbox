// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Ipam.Inputs
{

    public sealed class GetPrefixesFilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the field to filter on. Supported fields are: `prefix`, `contains`, `vlan_vid`, `vrf_id`, `vlan_id`, `status`, `tenant_id`, `site_id`, &amp; `tag`.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The value to pass to the specified filter.
        /// </summary>
        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetPrefixesFilterArgs()
        {
        }
        public static new GetPrefixesFilterArgs Empty => new GetPrefixesFilterArgs();
    }
}
