// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netbox

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetTagsFilter struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

// GetTagsFilterInput is an input type that accepts GetTagsFilterArgs and GetTagsFilterOutput values.
// You can construct a concrete instance of `GetTagsFilterInput` via:
//
//	GetTagsFilterArgs{...}
type GetTagsFilterInput interface {
	pulumi.Input

	ToGetTagsFilterOutput() GetTagsFilterOutput
	ToGetTagsFilterOutputWithContext(context.Context) GetTagsFilterOutput
}

type GetTagsFilterArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (GetTagsFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsFilter)(nil)).Elem()
}

func (i GetTagsFilterArgs) ToGetTagsFilterOutput() GetTagsFilterOutput {
	return i.ToGetTagsFilterOutputWithContext(context.Background())
}

func (i GetTagsFilterArgs) ToGetTagsFilterOutputWithContext(ctx context.Context) GetTagsFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTagsFilterOutput)
}

// GetTagsFilterArrayInput is an input type that accepts GetTagsFilterArray and GetTagsFilterArrayOutput values.
// You can construct a concrete instance of `GetTagsFilterArrayInput` via:
//
//	GetTagsFilterArray{ GetTagsFilterArgs{...} }
type GetTagsFilterArrayInput interface {
	pulumi.Input

	ToGetTagsFilterArrayOutput() GetTagsFilterArrayOutput
	ToGetTagsFilterArrayOutputWithContext(context.Context) GetTagsFilterArrayOutput
}

type GetTagsFilterArray []GetTagsFilterInput

func (GetTagsFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTagsFilter)(nil)).Elem()
}

func (i GetTagsFilterArray) ToGetTagsFilterArrayOutput() GetTagsFilterArrayOutput {
	return i.ToGetTagsFilterArrayOutputWithContext(context.Background())
}

func (i GetTagsFilterArray) ToGetTagsFilterArrayOutputWithContext(ctx context.Context) GetTagsFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTagsFilterArrayOutput)
}

type GetTagsFilterOutput struct{ *pulumi.OutputState }

func (GetTagsFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsFilter)(nil)).Elem()
}

func (o GetTagsFilterOutput) ToGetTagsFilterOutput() GetTagsFilterOutput {
	return o
}

func (o GetTagsFilterOutput) ToGetTagsFilterOutputWithContext(ctx context.Context) GetTagsFilterOutput {
	return o
}

func (o GetTagsFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsFilter) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetTagsFilterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsFilter) string { return v.Value }).(pulumi.StringOutput)
}

type GetTagsFilterArrayOutput struct{ *pulumi.OutputState }

func (GetTagsFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTagsFilter)(nil)).Elem()
}

func (o GetTagsFilterArrayOutput) ToGetTagsFilterArrayOutput() GetTagsFilterArrayOutput {
	return o
}

func (o GetTagsFilterArrayOutput) ToGetTagsFilterArrayOutputWithContext(ctx context.Context) GetTagsFilterArrayOutput {
	return o
}

func (o GetTagsFilterArrayOutput) Index(i pulumi.IntInput) GetTagsFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetTagsFilter {
		return vs[0].([]GetTagsFilter)[vs[1].(int)]
	}).(GetTagsFilterOutput)
}

type GetTagsTag struct {
	Color       *string `pulumi:"color"`
	Description *string `pulumi:"description"`
	Name        string  `pulumi:"name"`
	Slug        string  `pulumi:"slug"`
	TagId       int     `pulumi:"tagId"`
}

// GetTagsTagInput is an input type that accepts GetTagsTagArgs and GetTagsTagOutput values.
// You can construct a concrete instance of `GetTagsTagInput` via:
//
//	GetTagsTagArgs{...}
type GetTagsTagInput interface {
	pulumi.Input

	ToGetTagsTagOutput() GetTagsTagOutput
	ToGetTagsTagOutputWithContext(context.Context) GetTagsTagOutput
}

type GetTagsTagArgs struct {
	Color       pulumi.StringPtrInput `pulumi:"color"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Slug        pulumi.StringInput    `pulumi:"slug"`
	TagId       pulumi.IntInput       `pulumi:"tagId"`
}

func (GetTagsTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsTag)(nil)).Elem()
}

func (i GetTagsTagArgs) ToGetTagsTagOutput() GetTagsTagOutput {
	return i.ToGetTagsTagOutputWithContext(context.Background())
}

func (i GetTagsTagArgs) ToGetTagsTagOutputWithContext(ctx context.Context) GetTagsTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTagsTagOutput)
}

// GetTagsTagArrayInput is an input type that accepts GetTagsTagArray and GetTagsTagArrayOutput values.
// You can construct a concrete instance of `GetTagsTagArrayInput` via:
//
//	GetTagsTagArray{ GetTagsTagArgs{...} }
type GetTagsTagArrayInput interface {
	pulumi.Input

	ToGetTagsTagArrayOutput() GetTagsTagArrayOutput
	ToGetTagsTagArrayOutputWithContext(context.Context) GetTagsTagArrayOutput
}

type GetTagsTagArray []GetTagsTagInput

func (GetTagsTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTagsTag)(nil)).Elem()
}

func (i GetTagsTagArray) ToGetTagsTagArrayOutput() GetTagsTagArrayOutput {
	return i.ToGetTagsTagArrayOutputWithContext(context.Background())
}

func (i GetTagsTagArray) ToGetTagsTagArrayOutputWithContext(ctx context.Context) GetTagsTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTagsTagArrayOutput)
}

type GetTagsTagOutput struct{ *pulumi.OutputState }

func (GetTagsTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsTag)(nil)).Elem()
}

func (o GetTagsTagOutput) ToGetTagsTagOutput() GetTagsTagOutput {
	return o
}

func (o GetTagsTagOutput) ToGetTagsTagOutputWithContext(ctx context.Context) GetTagsTagOutput {
	return o
}

func (o GetTagsTagOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTagsTag) *string { return v.Color }).(pulumi.StringPtrOutput)
}

func (o GetTagsTagOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTagsTag) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetTagsTagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsTag) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetTagsTagOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsTag) string { return v.Slug }).(pulumi.StringOutput)
}

func (o GetTagsTagOutput) TagId() pulumi.IntOutput {
	return o.ApplyT(func(v GetTagsTag) int { return v.TagId }).(pulumi.IntOutput)
}

type GetTagsTagArrayOutput struct{ *pulumi.OutputState }

func (GetTagsTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTagsTag)(nil)).Elem()
}

func (o GetTagsTagArrayOutput) ToGetTagsTagArrayOutput() GetTagsTagArrayOutput {
	return o
}

func (o GetTagsTagArrayOutput) ToGetTagsTagArrayOutputWithContext(ctx context.Context) GetTagsTagArrayOutput {
	return o
}

func (o GetTagsTagArrayOutput) Index(i pulumi.IntInput) GetTagsTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetTagsTag {
		return vs[0].([]GetTagsTag)[vs[1].(int)]
	}).(GetTagsTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetTagsFilterInput)(nil)).Elem(), GetTagsFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTagsFilterArrayInput)(nil)).Elem(), GetTagsFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTagsTagInput)(nil)).Elem(), GetTagsTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTagsTagArrayInput)(nil)).Elem(), GetTagsTagArray{})
	pulumi.RegisterOutputType(GetTagsFilterOutput{})
	pulumi.RegisterOutputType(GetTagsFilterArrayOutput{})
	pulumi.RegisterOutputType(GetTagsTagOutput{})
	pulumi.RegisterOutputType(GetTagsTagArrayOutput{})
}