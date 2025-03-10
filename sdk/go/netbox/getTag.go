// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netbox

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
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := netbox.LookupTag(ctx, &netbox.LookupTagArgs{
//				Name: "DMZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTag(ctx *pulumi.Context, args *LookupTagArgs, opts ...pulumi.InvokeOption) (*LookupTagResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTagResult
	err := ctx.Invoke("netbox:index/getTag:getTag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTag.
type LookupTagArgs struct {
	Description *string `pulumi:"description"`
	Name        string  `pulumi:"name"`
}

// A collection of values returned by getTag.
type LookupTagResult struct {
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	Slug string `pulumi:"slug"`
}

func LookupTagOutput(ctx *pulumi.Context, args LookupTagOutputArgs, opts ...pulumi.InvokeOption) LookupTagResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTagResultOutput, error) {
			args := v.(LookupTagArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:index/getTag:getTag", args, LookupTagResultOutput{}, options).(LookupTagResultOutput), nil
		}).(LookupTagResultOutput)
}

// A collection of arguments for invoking getTag.
type LookupTagOutputArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringInput    `pulumi:"name"`
}

func (LookupTagOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagArgs)(nil)).Elem()
}

// A collection of values returned by getTag.
type LookupTagResultOutput struct{ *pulumi.OutputState }

func (LookupTagResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagResult)(nil)).Elem()
}

func (o LookupTagResultOutput) ToLookupTagResultOutput() LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) ToLookupTagResultOutputWithContext(ctx context.Context) LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTagResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagResultOutput{})
}
