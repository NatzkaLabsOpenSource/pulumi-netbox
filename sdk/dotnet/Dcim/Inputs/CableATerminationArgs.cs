// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox.Dcim.Inputs
{

    public sealed class CableATerminationArgs : global::Pulumi.ResourceArgs
    {
        [Input("objectId", required: true)]
        public Input<int> ObjectId { get; set; } = null!;

        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        public CableATerminationArgs()
        {
        }
        public static new CableATerminationArgs Empty => new CableATerminationArgs();
    }
}
