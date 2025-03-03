// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netbox

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigContext(ctx *pulumi.Context, args *LookupConfigContextArgs, opts ...pulumi.InvokeOption) (*LookupConfigContextResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigContextResult
	err := ctx.Invoke("netbox:index/getConfigContext:getConfigContext", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConfigContext.
type LookupConfigContextArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getConfigContext.
type LookupConfigContextResult struct {
	ClusterGroups []int  `pulumi:"clusterGroups"`
	ClusterTypes  []int  `pulumi:"clusterTypes"`
	Clusters      []int  `pulumi:"clusters"`
	Data          string `pulumi:"data"`
	Description   string `pulumi:"description"`
	DeviceTypes   []int  `pulumi:"deviceTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id           string   `pulumi:"id"`
	Locations    []int    `pulumi:"locations"`
	Name         string   `pulumi:"name"`
	Platforms    []int    `pulumi:"platforms"`
	Regions      []int    `pulumi:"regions"`
	Roles        []int    `pulumi:"roles"`
	SiteGroups   []int    `pulumi:"siteGroups"`
	Sites        []int    `pulumi:"sites"`
	Tags         []string `pulumi:"tags"`
	TenantGroups []int    `pulumi:"tenantGroups"`
	Tenants      []int    `pulumi:"tenants"`
	Weight       int      `pulumi:"weight"`
}

func LookupConfigContextOutput(ctx *pulumi.Context, args LookupConfigContextOutputArgs, opts ...pulumi.InvokeOption) LookupConfigContextResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupConfigContextResultOutput, error) {
			args := v.(LookupConfigContextArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:index/getConfigContext:getConfigContext", args, LookupConfigContextResultOutput{}, options).(LookupConfigContextResultOutput), nil
		}).(LookupConfigContextResultOutput)
}

// A collection of arguments for invoking getConfigContext.
type LookupConfigContextOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupConfigContextOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigContextArgs)(nil)).Elem()
}

// A collection of values returned by getConfigContext.
type LookupConfigContextResultOutput struct{ *pulumi.OutputState }

func (LookupConfigContextResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigContextResult)(nil)).Elem()
}

func (o LookupConfigContextResultOutput) ToLookupConfigContextResultOutput() LookupConfigContextResultOutput {
	return o
}

func (o LookupConfigContextResultOutput) ToLookupConfigContextResultOutputWithContext(ctx context.Context) LookupConfigContextResultOutput {
	return o
}

func (o LookupConfigContextResultOutput) ClusterGroups() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.ClusterGroups }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) ClusterTypes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.ClusterTypes }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) Clusters() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.Clusters }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigContextResult) string { return v.Data }).(pulumi.StringOutput)
}

func (o LookupConfigContextResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigContextResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupConfigContextResultOutput) DeviceTypes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.DeviceTypes }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupConfigContextResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigContextResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigContextResultOutput) Locations() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.Locations }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigContextResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigContextResultOutput) Platforms() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.Platforms }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) Regions() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.Regions }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) Roles() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.Roles }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) SiteGroups() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.SiteGroups }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) Sites() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.Sites }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupConfigContextResultOutput) TenantGroups() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.TenantGroups }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) Tenants() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupConfigContextResult) []int { return v.Tenants }).(pulumi.IntArrayOutput)
}

func (o LookupConfigContextResultOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConfigContextResult) int { return v.Weight }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigContextResultOutput{})
}
