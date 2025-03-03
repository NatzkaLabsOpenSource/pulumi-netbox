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
    public sealed class GetVlansVlanResult
    {
        public readonly string Description;
        public readonly int GroupId;
        public readonly string Name;
        public readonly int Role;
        public readonly int Site;
        public readonly string Status;
        public readonly ImmutableArray<int> TagIds;
        public readonly int Tenant;
        public readonly int Vid;

        [OutputConstructor]
        private GetVlansVlanResult(
            string description,

            int groupId,

            string name,

            int role,

            int site,

            string status,

            ImmutableArray<int> tagIds,

            int tenant,

            int vid)
        {
            Description = description;
            GroupId = groupId;
            Name = name;
            Role = role;
            Site = site;
            Status = status;
            TagIds = tagIds;
            Tenant = tenant;
            Vid = vid;
        }
    }
}
