// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Virt.Outputs
{

    [OutputType]
    public sealed class GetVirtualMachinesVmResult
    {
        public readonly int ClusterId;
        public readonly string Comments;
        public readonly string ConfigContext;
        public readonly ImmutableDictionary<string, string> CustomFields;
        public readonly string Description;
        public readonly int DeviceId;
        public readonly string DeviceName;
        public readonly int DiskSizeMb;
        public readonly string LocalContextData;
        public readonly int MemoryMb;
        public readonly string Name;
        public readonly int PlatformId;
        public readonly string PlatformSlug;
        public readonly string PrimaryIp;
        public readonly string PrimaryIp4;
        public readonly string PrimaryIp6;
        public readonly int RoleId;
        public readonly int SiteId;
        public readonly string Status;
        public readonly ImmutableArray<int> TagIds;
        public readonly int TenantId;
        public readonly double Vcpus;
        public readonly int VmId;

        [OutputConstructor]
        private GetVirtualMachinesVmResult(
            int clusterId,

            string comments,

            string configContext,

            ImmutableDictionary<string, string> customFields,

            string description,

            int deviceId,

            string deviceName,

            int diskSizeMb,

            string localContextData,

            int memoryMb,

            string name,

            int platformId,

            string platformSlug,

            string primaryIp,

            string primaryIp4,

            string primaryIp6,

            int roleId,

            int siteId,

            string status,

            ImmutableArray<int> tagIds,

            int tenantId,

            double vcpus,

            int vmId)
        {
            ClusterId = clusterId;
            Comments = comments;
            ConfigContext = configContext;
            CustomFields = customFields;
            Description = description;
            DeviceId = deviceId;
            DeviceName = deviceName;
            DiskSizeMb = diskSizeMb;
            LocalContextData = localContextData;
            MemoryMb = memoryMb;
            Name = name;
            PlatformId = platformId;
            PlatformSlug = platformSlug;
            PrimaryIp = primaryIp;
            PrimaryIp4 = primaryIp4;
            PrimaryIp6 = primaryIp6;
            RoleId = roleId;
            SiteId = siteId;
            Status = status;
            TagIds = tagIds;
            TenantId = tenantId;
            Vcpus = vcpus;
            VmId = vmId;
        }
    }
}
