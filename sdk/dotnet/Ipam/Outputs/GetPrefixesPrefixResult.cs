// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Ipam.Outputs
{

    [OutputType]
    public sealed class GetPrefixesPrefixResult
    {
        public readonly string Description;
        public readonly int Id;
        public readonly string Prefix;
        public readonly int SiteId;
        public readonly string Status;
        public readonly ImmutableArray<string> Tags;
        public readonly int TenantId;
        public readonly int VlanId;
        public readonly double VlanVid;
        public readonly int VrfId;

        [OutputConstructor]
        private GetPrefixesPrefixResult(
            string description,

            int id,

            string prefix,

            int siteId,

            string status,

            ImmutableArray<string> tags,

            int tenantId,

            int vlanId,

            double vlanVid,

            int vrfId)
        {
            Description = description;
            Id = id;
            Prefix = prefix;
            SiteId = siteId;
            Status = status;
            Tags = tags;
            TenantId = tenantId;
            VlanId = vlanId;
            VlanVid = vlanVid;
            VrfId = vrfId;
        }
    }
}
