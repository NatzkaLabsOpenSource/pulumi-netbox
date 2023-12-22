// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPrefixes(ctx *pulumi.Context, args *GetPrefixesArgs, opts ...pulumi.InvokeOption) (*GetPrefixesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPrefixesResult
	err := ctx.Invoke("netbox:ipam/getPrefixes:getPrefixes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrefixes.
type GetPrefixesArgs struct {
	// A list of filters to apply to the API query when requesting prefixes.
	Filters []GetPrefixesFilter `pulumi:"filters"`
	// The limit of objects to return from the API lookup. Defaults to `0`.
	Limit *int `pulumi:"limit"`
}

// A collection of values returned by getPrefixes.
type GetPrefixesResult struct {
	// A list of filters to apply to the API query when requesting prefixes.
	Filters []GetPrefixesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The limit of objects to return from the API lookup. Defaults to `0`.
	Limit    *int                `pulumi:"limit"`
	Prefixes []GetPrefixesPrefix `pulumi:"prefixes"`
}

func GetPrefixesOutput(ctx *pulumi.Context, args GetPrefixesOutputArgs, opts ...pulumi.InvokeOption) GetPrefixesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPrefixesResult, error) {
			args := v.(GetPrefixesArgs)
			r, err := GetPrefixes(ctx, &args, opts...)
			var s GetPrefixesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPrefixesResultOutput)
}

// A collection of arguments for invoking getPrefixes.
type GetPrefixesOutputArgs struct {
	// A list of filters to apply to the API query when requesting prefixes.
	Filters GetPrefixesFilterArrayInput `pulumi:"filters"`
	// The limit of objects to return from the API lookup. Defaults to `0`.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
}

func (GetPrefixesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixesArgs)(nil)).Elem()
}

// A collection of values returned by getPrefixes.
type GetPrefixesResultOutput struct{ *pulumi.OutputState }

func (GetPrefixesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixesResult)(nil)).Elem()
}

func (o GetPrefixesResultOutput) ToGetPrefixesResultOutput() GetPrefixesResultOutput {
	return o
}

func (o GetPrefixesResultOutput) ToGetPrefixesResultOutputWithContext(ctx context.Context) GetPrefixesResultOutput {
	return o
}

// A list of filters to apply to the API query when requesting prefixes.
func (o GetPrefixesResultOutput) Filters() GetPrefixesFilterArrayOutput {
	return o.ApplyT(func(v GetPrefixesResult) []GetPrefixesFilter { return v.Filters }).(GetPrefixesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPrefixesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPrefixesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The limit of objects to return from the API lookup. Defaults to `0`.
func (o GetPrefixesResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetPrefixesResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetPrefixesResultOutput) Prefixes() GetPrefixesPrefixArrayOutput {
	return o.ApplyT(func(v GetPrefixesResult) []GetPrefixesPrefix { return v.Prefixes }).(GetPrefixesPrefixArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPrefixesResultOutput{})
}