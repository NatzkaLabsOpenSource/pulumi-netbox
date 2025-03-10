// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcim

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLocation(ctx *pulumi.Context, args *LookupLocationArgs, opts ...pulumi.InvokeOption) (*LookupLocationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocationResult
	err := ctx.Invoke("netbox:dcim/getLocation:getLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocation.
type LookupLocationArgs struct {
	// The ID of this resource.
	Id       *string `pulumi:"id"`
	Name     *string `pulumi:"name"`
	ParentId *int    `pulumi:"parentId"`
	SiteId   *int    `pulumi:"siteId"`
	Slug     *string `pulumi:"slug"`
}

// A collection of values returned by getLocation.
type LookupLocationResult struct {
	Description string `pulumi:"description"`
	// The ID of this resource.
	Id       string  `pulumi:"id"`
	Name     *string `pulumi:"name"`
	ParentId int     `pulumi:"parentId"`
	SiteId   int     `pulumi:"siteId"`
	Slug     *string `pulumi:"slug"`
	Status   string  `pulumi:"status"`
	TenantId int     `pulumi:"tenantId"`
}

func LookupLocationOutput(ctx *pulumi.Context, args LookupLocationOutputArgs, opts ...pulumi.InvokeOption) LookupLocationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLocationResultOutput, error) {
			args := v.(LookupLocationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:dcim/getLocation:getLocation", args, LookupLocationResultOutput{}, options).(LookupLocationResultOutput), nil
		}).(LookupLocationResultOutput)
}

// A collection of arguments for invoking getLocation.
type LookupLocationOutputArgs struct {
	// The ID of this resource.
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	ParentId pulumi.IntPtrInput    `pulumi:"parentId"`
	SiteId   pulumi.IntPtrInput    `pulumi:"siteId"`
	Slug     pulumi.StringPtrInput `pulumi:"slug"`
}

func (LookupLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationArgs)(nil)).Elem()
}

// A collection of values returned by getLocation.
type LookupLocationResultOutput struct{ *pulumi.OutputState }

func (LookupLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationResult)(nil)).Elem()
}

func (o LookupLocationResultOutput) ToLookupLocationResultOutput() LookupLocationResultOutput {
	return o
}

func (o LookupLocationResultOutput) ToLookupLocationResultOutputWithContext(ctx context.Context) LookupLocationResultOutput {
	return o
}

func (o LookupLocationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationResult) string { return v.Description }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupLocationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLocationResultOutput) ParentId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLocationResult) int { return v.ParentId }).(pulumi.IntOutput)
}

func (o LookupLocationResultOutput) SiteId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLocationResult) int { return v.SiteId }).(pulumi.IntOutput)
}

func (o LookupLocationResultOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationResult) *string { return v.Slug }).(pulumi.StringPtrOutput)
}

func (o LookupLocationResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocationResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupLocationResultOutput) TenantId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLocationResult) int { return v.TenantId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationResultOutput{})
}
