// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package virt

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/virt"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := virt.LookupClusterGroup(ctx, &virt.LookupClusterGroupArgs{
//				Name: "dc-west",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupClusterGroup(ctx *pulumi.Context, args *LookupClusterGroupArgs, opts ...pulumi.InvokeOption) (*LookupClusterGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterGroupResult
	err := ctx.Invoke("netbox:virt/getClusterGroup:getClusterGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterGroup.
type LookupClusterGroupArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getClusterGroup.
type LookupClusterGroupResult struct {
	ClusterGroupId int `pulumi:"clusterGroupId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupClusterGroupOutput(ctx *pulumi.Context, args LookupClusterGroupOutputArgs, opts ...pulumi.InvokeOption) LookupClusterGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClusterGroupResultOutput, error) {
			args := v.(LookupClusterGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:virt/getClusterGroup:getClusterGroup", args, LookupClusterGroupResultOutput{}, options).(LookupClusterGroupResultOutput), nil
		}).(LookupClusterGroupResultOutput)
}

// A collection of arguments for invoking getClusterGroup.
type LookupClusterGroupOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupClusterGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterGroupArgs)(nil)).Elem()
}

// A collection of values returned by getClusterGroup.
type LookupClusterGroupResultOutput struct{ *pulumi.OutputState }

func (LookupClusterGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterGroupResult)(nil)).Elem()
}

func (o LookupClusterGroupResultOutput) ToLookupClusterGroupResultOutput() LookupClusterGroupResultOutput {
	return o
}

func (o LookupClusterGroupResultOutput) ToLookupClusterGroupResultOutputWithContext(ctx context.Context) LookupClusterGroupResultOutput {
	return o
}

func (o LookupClusterGroupResultOutput) ClusterGroupId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterGroupResult) int { return v.ClusterGroupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClusterGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterGroupResultOutput{})
}
