// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouteTarget(ctx *pulumi.Context, args *LookupRouteTargetArgs, opts ...pulumi.InvokeOption) (*LookupRouteTargetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteTargetResult
	err := ctx.Invoke("netbox:ipam/getRouteTarget:getRouteTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTarget.
type LookupRouteTargetArgs struct {
	Name string   `pulumi:"name"`
	Tags []string `pulumi:"tags"`
}

// A collection of values returned by getRouteTarget.
type LookupRouteTargetResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id       string   `pulumi:"id"`
	Name     string   `pulumi:"name"`
	Tags     []string `pulumi:"tags"`
	TenantId int      `pulumi:"tenantId"`
}

func LookupRouteTargetOutput(ctx *pulumi.Context, args LookupRouteTargetOutputArgs, opts ...pulumi.InvokeOption) LookupRouteTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteTargetResult, error) {
			args := v.(LookupRouteTargetArgs)
			r, err := LookupRouteTarget(ctx, &args, opts...)
			var s LookupRouteTargetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteTargetResultOutput)
}

// A collection of arguments for invoking getRouteTarget.
type LookupRouteTargetOutputArgs struct {
	Name pulumi.StringInput      `pulumi:"name"`
	Tags pulumi.StringArrayInput `pulumi:"tags"`
}

func (LookupRouteTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTargetArgs)(nil)).Elem()
}

// A collection of values returned by getRouteTarget.
type LookupRouteTargetResultOutput struct{ *pulumi.OutputState }

func (LookupRouteTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTargetResult)(nil)).Elem()
}

func (o LookupRouteTargetResultOutput) ToLookupRouteTargetResultOutput() LookupRouteTargetResultOutput {
	return o
}

func (o LookupRouteTargetResultOutput) ToLookupRouteTargetResultOutputWithContext(ctx context.Context) LookupRouteTargetResultOutput {
	return o
}

func (o LookupRouteTargetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTargetResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouteTargetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTargetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouteTargetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTargetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRouteTargetResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteTargetResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupRouteTargetResultOutput) TenantId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouteTargetResult) int { return v.TenantId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteTargetResultOutput{})
}
