// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcim

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
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/dcim"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Assumes the corresponding site groups exist
//			_, err := dcim.LookupSiteGroup(ctx, &dcim.LookupSiteGroupArgs{
//				Name: pulumi.StringRef("example-sitegroup-1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = dcim.LookupSiteGroup(ctx, &dcim.LookupSiteGroupArgs{
//				Slug: pulumi.StringRef("sitegrp"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSiteGroup(ctx *pulumi.Context, args *LookupSiteGroupArgs, opts ...pulumi.InvokeOption) (*LookupSiteGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSiteGroupResult
	err := ctx.Invoke("netbox:dcim/getSiteGroup:getSiteGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSiteGroup.
type LookupSiteGroupArgs struct {
	// At least one of `name` or `slug` must be given.
	Name *string `pulumi:"name"`
	// At least one of `name` or `slug` must be given.
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getSiteGroup.
type LookupSiteGroupResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// At least one of `name` or `slug` must be given.
	Name string `pulumi:"name"`
	// At least one of `name` or `slug` must be given.
	Slug string `pulumi:"slug"`
}

func LookupSiteGroupOutput(ctx *pulumi.Context, args LookupSiteGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSiteGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSiteGroupResultOutput, error) {
			args := v.(LookupSiteGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:dcim/getSiteGroup:getSiteGroup", args, LookupSiteGroupResultOutput{}, options).(LookupSiteGroupResultOutput), nil
		}).(LookupSiteGroupResultOutput)
}

// A collection of arguments for invoking getSiteGroup.
type LookupSiteGroupOutputArgs struct {
	// At least one of `name` or `slug` must be given.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// At least one of `name` or `slug` must be given.
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupSiteGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteGroupArgs)(nil)).Elem()
}

// A collection of values returned by getSiteGroup.
type LookupSiteGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSiteGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteGroupResult)(nil)).Elem()
}

func (o LookupSiteGroupResultOutput) ToLookupSiteGroupResultOutput() LookupSiteGroupResultOutput {
	return o
}

func (o LookupSiteGroupResultOutput) ToLookupSiteGroupResultOutputWithContext(ctx context.Context) LookupSiteGroupResultOutput {
	return o
}

func (o LookupSiteGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSiteGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// At least one of `name` or `slug` must be given.
func (o LookupSiteGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// At least one of `name` or `slug` must be given.
func (o LookupSiteGroupResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteGroupResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteGroupResultOutput{})
}
