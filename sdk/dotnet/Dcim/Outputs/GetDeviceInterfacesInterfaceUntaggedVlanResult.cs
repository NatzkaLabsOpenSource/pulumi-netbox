// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Dcim.Outputs
{

    [OutputType]
    public sealed class GetDeviceInterfacesInterfaceUntaggedVlanResult
    {
        public readonly int Id;
        public readonly string Name;
        public readonly int Vid;

        [OutputConstructor]
        private GetDeviceInterfacesInterfaceUntaggedVlanResult(
            int id,

            string name,

            int vid)
        {
            Id = id;
            Name = name;
            Vid = vid;
        }
    }
}
