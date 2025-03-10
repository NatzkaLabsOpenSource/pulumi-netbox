// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRole(ctx *pulumi.Context, args *LookupRoleArgs, opts ...pulumi.InvokeOption) (*LookupRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRoleResult
	err := ctx.Invoke("netbox:ipam/getRole:getRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRole.
type LookupRoleArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getRole.
type LookupRoleResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Name   string `pulumi:"name"`
	Slug   string `pulumi:"slug"`
	Weight int    `pulumi:"weight"`
}

func LookupRoleOutput(ctx *pulumi.Context, args LookupRoleOutputArgs, opts ...pulumi.InvokeOption) LookupRoleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRoleResultOutput, error) {
			args := v.(LookupRoleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:ipam/getRole:getRole", args, LookupRoleResultOutput{}, options).(LookupRoleResultOutput), nil
		}).(LookupRoleResultOutput)
}

// A collection of arguments for invoking getRole.
type LookupRoleOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleArgs)(nil)).Elem()
}

// A collection of values returned by getRole.
type LookupRoleResultOutput struct{ *pulumi.OutputState }

func (LookupRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleResult)(nil)).Elem()
}

func (o LookupRoleResultOutput) ToLookupRoleResultOutput() LookupRoleResultOutput {
	return o
}

func (o LookupRoleResultOutput) ToLookupRoleResultOutputWithContext(ctx context.Context) LookupRoleResultOutput {
	return o
}

func (o LookupRoleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Slug }).(pulumi.StringOutput)
}

func (o LookupRoleResultOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRoleResult) int { return v.Weight }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleResultOutput{})
}
