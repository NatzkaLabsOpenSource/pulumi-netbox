// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tenancy

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
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/tenancy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tenancy.LookupTenant(ctx, &tenancy.LookupTenantArgs{
//				Name: pulumi.StringRef("Customer A"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTenant(ctx *pulumi.Context, args *LookupTenantArgs, opts ...pulumi.InvokeOption) (*LookupTenantResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTenantResult
	err := ctx.Invoke("netbox:tenancy/getTenant:getTenant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTenant.
type LookupTenantArgs struct {
	Description *string `pulumi:"description"`
	// At least one of `name` or `slug` must be given.
	Name *string `pulumi:"name"`
	// At least one of `name` or `slug` must be given.
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getTenant.
type LookupTenantResult struct {
	Description *string `pulumi:"description"`
	GroupId     int     `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// At least one of `name` or `slug` must be given.
	Name string `pulumi:"name"`
	// At least one of `name` or `slug` must be given.
	Slug string `pulumi:"slug"`
}

func LookupTenantOutput(ctx *pulumi.Context, args LookupTenantOutputArgs, opts ...pulumi.InvokeOption) LookupTenantResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTenantResultOutput, error) {
			args := v.(LookupTenantArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:tenancy/getTenant:getTenant", args, LookupTenantResultOutput{}, options).(LookupTenantResultOutput), nil
		}).(LookupTenantResultOutput)
}

// A collection of arguments for invoking getTenant.
type LookupTenantOutputArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	// At least one of `name` or `slug` must be given.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// At least one of `name` or `slug` must be given.
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupTenantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantArgs)(nil)).Elem()
}

// A collection of values returned by getTenant.
type LookupTenantResultOutput struct{ *pulumi.OutputState }

func (LookupTenantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantResult)(nil)).Elem()
}

func (o LookupTenantResultOutput) ToLookupTenantResultOutput() LookupTenantResultOutput {
	return o
}

func (o LookupTenantResultOutput) ToLookupTenantResultOutputWithContext(ctx context.Context) LookupTenantResultOutput {
	return o
}

func (o LookupTenantResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTenantResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTenantResultOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTenantResult) int { return v.GroupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTenantResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantResult) string { return v.Id }).(pulumi.StringOutput)
}

// At least one of `name` or `slug` must be given.
func (o LookupTenantResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantResult) string { return v.Name }).(pulumi.StringOutput)
}

// At least one of `name` or `slug` must be given.
func (o LookupTenantResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTenantResultOutput{})
}
