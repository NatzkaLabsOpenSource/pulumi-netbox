// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Dcim.Inputs
{

    public sealed class GetDeviceInterfacesFilterArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetDeviceInterfacesFilterArgs()
        {
        }
        public static new GetDeviceInterfacesFilterArgs Empty => new GetDeviceInterfacesFilterArgs();
    }
}
