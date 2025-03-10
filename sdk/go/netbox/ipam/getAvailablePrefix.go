// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAvailablePrefix(ctx *pulumi.Context, args *LookupAvailablePrefixArgs, opts ...pulumi.InvokeOption) (*LookupAvailablePrefixResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAvailablePrefixResult
	err := ctx.Invoke("netbox:ipam/getAvailablePrefix:getAvailablePrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailablePrefix.
type LookupAvailablePrefixArgs struct {
	PrefixId int `pulumi:"prefixId"`
}

// A collection of values returned by getAvailablePrefix.
type LookupAvailablePrefixResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string                                `pulumi:"id"`
	PrefixId           int                                   `pulumi:"prefixId"`
	PrefixesAvailables []GetAvailablePrefixPrefixesAvailable `pulumi:"prefixesAvailables"`
}

func LookupAvailablePrefixOutput(ctx *pulumi.Context, args LookupAvailablePrefixOutputArgs, opts ...pulumi.InvokeOption) LookupAvailablePrefixResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAvailablePrefixResultOutput, error) {
			args := v.(LookupAvailablePrefixArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:ipam/getAvailablePrefix:getAvailablePrefix", args, LookupAvailablePrefixResultOutput{}, options).(LookupAvailablePrefixResultOutput), nil
		}).(LookupAvailablePrefixResultOutput)
}

// A collection of arguments for invoking getAvailablePrefix.
type LookupAvailablePrefixOutputArgs struct {
	PrefixId pulumi.IntInput `pulumi:"prefixId"`
}

func (LookupAvailablePrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailablePrefixArgs)(nil)).Elem()
}

// A collection of values returned by getAvailablePrefix.
type LookupAvailablePrefixResultOutput struct{ *pulumi.OutputState }

func (LookupAvailablePrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailablePrefixResult)(nil)).Elem()
}

func (o LookupAvailablePrefixResultOutput) ToLookupAvailablePrefixResultOutput() LookupAvailablePrefixResultOutput {
	return o
}

func (o LookupAvailablePrefixResultOutput) ToLookupAvailablePrefixResultOutputWithContext(ctx context.Context) LookupAvailablePrefixResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAvailablePrefixResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailablePrefixResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAvailablePrefixResultOutput) PrefixId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAvailablePrefixResult) int { return v.PrefixId }).(pulumi.IntOutput)
}

func (o LookupAvailablePrefixResultOutput) PrefixesAvailables() GetAvailablePrefixPrefixesAvailableArrayOutput {
	return o.ApplyT(func(v LookupAvailablePrefixResult) []GetAvailablePrefixPrefixesAvailable { return v.PrefixesAvailables }).(GetAvailablePrefixPrefixesAvailableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAvailablePrefixResultOutput{})
}
