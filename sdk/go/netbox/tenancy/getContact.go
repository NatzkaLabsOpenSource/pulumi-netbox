// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tenancy

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContact(ctx *pulumi.Context, args *LookupContactArgs, opts ...pulumi.InvokeOption) (*LookupContactResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContactResult
	err := ctx.Invoke("netbox:tenancy/getContact:getContact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContact.
type LookupContactArgs struct {
	Description *string `pulumi:"description"`
	// At least one of `name` or `slug` must be given.
	Name *string `pulumi:"name"`
	// At least one of `name` or `slug` must be given.
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getContact.
type LookupContactResult struct {
	Description *string `pulumi:"description"`
	GroupId     int     `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// At least one of `name` or `slug` must be given.
	Name string `pulumi:"name"`
	// At least one of `name` or `slug` must be given.
	Slug string `pulumi:"slug"`
}

func LookupContactOutput(ctx *pulumi.Context, args LookupContactOutputArgs, opts ...pulumi.InvokeOption) LookupContactResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupContactResultOutput, error) {
			args := v.(LookupContactArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:tenancy/getContact:getContact", args, LookupContactResultOutput{}, options).(LookupContactResultOutput), nil
		}).(LookupContactResultOutput)
}

// A collection of arguments for invoking getContact.
type LookupContactOutputArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	// At least one of `name` or `slug` must be given.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// At least one of `name` or `slug` must be given.
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupContactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactArgs)(nil)).Elem()
}

// A collection of values returned by getContact.
type LookupContactResultOutput struct{ *pulumi.OutputState }

func (LookupContactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactResult)(nil)).Elem()
}

func (o LookupContactResultOutput) ToLookupContactResultOutput() LookupContactResultOutput {
	return o
}

func (o LookupContactResultOutput) ToLookupContactResultOutputWithContext(ctx context.Context) LookupContactResultOutput {
	return o
}

func (o LookupContactResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupContactResultOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContactResult) int { return v.GroupId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Id }).(pulumi.StringOutput)
}

// At least one of `name` or `slug` must be given.
func (o LookupContactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Name }).(pulumi.StringOutput)
}

// At least one of `name` or `slug` must be given.
func (o LookupContactResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContactResultOutput{})
}
