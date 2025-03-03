// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"errors"
	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#ip-ranges):
//
// > This model represents an arbitrary range of individual IPv4 or IPv6 addresses, inclusive of its starting and ending addresses. For instance, the range 192.0.2.10 to 192.0.2.20 has eleven members. (The total member count is available as the size property on an IPRange instance.) Like prefixes and IP addresses, each IP range may optionally be assigned to a VRF and/or tenant.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/ipam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ipam.NewIpRange(ctx, "cust_a_prod", &ipam.IpRangeArgs{
//				StartAddress: pulumi.String("10.0.0.1/24"),
//				EndAddress:   pulumi.String("10.0.0.50/24"),
//				Tags: pulumi.StringArray{
//					pulumi.String("customer-a"),
//					pulumi.String("prod"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IpRange struct {
	pulumi.CustomResourceState

	Description  pulumi.StringPtrOutput `pulumi:"description"`
	EndAddress   pulumi.StringOutput    `pulumi:"endAddress"`
	RoleId       pulumi.IntPtrOutput    `pulumi:"roleId"`
	StartAddress pulumi.StringOutput    `pulumi:"startAddress"`
	// Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
	Status   pulumi.StringPtrOutput   `pulumi:"status"`
	Tags     pulumi.StringArrayOutput `pulumi:"tags"`
	TenantId pulumi.IntPtrOutput      `pulumi:"tenantId"`
	VrfId    pulumi.IntPtrOutput      `pulumi:"vrfId"`
}

// NewIpRange registers a new resource with the given unique name, arguments, and options.
func NewIpRange(ctx *pulumi.Context,
	name string, args *IpRangeArgs, opts ...pulumi.ResourceOption) (*IpRange, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndAddress'")
	}
	if args.StartAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartAddress'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpRange
	err := ctx.RegisterResource("netbox:ipam/ipRange:IpRange", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpRange gets an existing IpRange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpRange(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpRangeState, opts ...pulumi.ResourceOption) (*IpRange, error) {
	var resource IpRange
	err := ctx.ReadResource("netbox:ipam/ipRange:IpRange", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpRange resources.
type ipRangeState struct {
	Description  *string `pulumi:"description"`
	EndAddress   *string `pulumi:"endAddress"`
	RoleId       *int    `pulumi:"roleId"`
	StartAddress *string `pulumi:"startAddress"`
	// Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
	Status   *string  `pulumi:"status"`
	Tags     []string `pulumi:"tags"`
	TenantId *int     `pulumi:"tenantId"`
	VrfId    *int     `pulumi:"vrfId"`
}

type IpRangeState struct {
	Description  pulumi.StringPtrInput
	EndAddress   pulumi.StringPtrInput
	RoleId       pulumi.IntPtrInput
	StartAddress pulumi.StringPtrInput
	// Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
	Status   pulumi.StringPtrInput
	Tags     pulumi.StringArrayInput
	TenantId pulumi.IntPtrInput
	VrfId    pulumi.IntPtrInput
}

func (IpRangeState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipRangeState)(nil)).Elem()
}

type ipRangeArgs struct {
	Description  *string `pulumi:"description"`
	EndAddress   string  `pulumi:"endAddress"`
	RoleId       *int    `pulumi:"roleId"`
	StartAddress string  `pulumi:"startAddress"`
	// Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
	Status   *string  `pulumi:"status"`
	Tags     []string `pulumi:"tags"`
	TenantId *int     `pulumi:"tenantId"`
	VrfId    *int     `pulumi:"vrfId"`
}

// The set of arguments for constructing a IpRange resource.
type IpRangeArgs struct {
	Description  pulumi.StringPtrInput
	EndAddress   pulumi.StringInput
	RoleId       pulumi.IntPtrInput
	StartAddress pulumi.StringInput
	// Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
	Status   pulumi.StringPtrInput
	Tags     pulumi.StringArrayInput
	TenantId pulumi.IntPtrInput
	VrfId    pulumi.IntPtrInput
}

func (IpRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipRangeArgs)(nil)).Elem()
}

type IpRangeInput interface {
	pulumi.Input

	ToIpRangeOutput() IpRangeOutput
	ToIpRangeOutputWithContext(ctx context.Context) IpRangeOutput
}

func (*IpRange) ElementType() reflect.Type {
	return reflect.TypeOf((**IpRange)(nil)).Elem()
}

func (i *IpRange) ToIpRangeOutput() IpRangeOutput {
	return i.ToIpRangeOutputWithContext(context.Background())
}

func (i *IpRange) ToIpRangeOutputWithContext(ctx context.Context) IpRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRangeOutput)
}

// IpRangeArrayInput is an input type that accepts IpRangeArray and IpRangeArrayOutput values.
// You can construct a concrete instance of `IpRangeArrayInput` via:
//
//	IpRangeArray{ IpRangeArgs{...} }
type IpRangeArrayInput interface {
	pulumi.Input

	ToIpRangeArrayOutput() IpRangeArrayOutput
	ToIpRangeArrayOutputWithContext(context.Context) IpRangeArrayOutput
}

type IpRangeArray []IpRangeInput

func (IpRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpRange)(nil)).Elem()
}

func (i IpRangeArray) ToIpRangeArrayOutput() IpRangeArrayOutput {
	return i.ToIpRangeArrayOutputWithContext(context.Background())
}

func (i IpRangeArray) ToIpRangeArrayOutputWithContext(ctx context.Context) IpRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRangeArrayOutput)
}

// IpRangeMapInput is an input type that accepts IpRangeMap and IpRangeMapOutput values.
// You can construct a concrete instance of `IpRangeMapInput` via:
//
//	IpRangeMap{ "key": IpRangeArgs{...} }
type IpRangeMapInput interface {
	pulumi.Input

	ToIpRangeMapOutput() IpRangeMapOutput
	ToIpRangeMapOutputWithContext(context.Context) IpRangeMapOutput
}

type IpRangeMap map[string]IpRangeInput

func (IpRangeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpRange)(nil)).Elem()
}

func (i IpRangeMap) ToIpRangeMapOutput() IpRangeMapOutput {
	return i.ToIpRangeMapOutputWithContext(context.Background())
}

func (i IpRangeMap) ToIpRangeMapOutputWithContext(ctx context.Context) IpRangeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRangeMapOutput)
}

type IpRangeOutput struct{ *pulumi.OutputState }

func (IpRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpRange)(nil)).Elem()
}

func (o IpRangeOutput) ToIpRangeOutput() IpRangeOutput {
	return o
}

func (o IpRangeOutput) ToIpRangeOutputWithContext(ctx context.Context) IpRangeOutput {
	return o
}

func (o IpRangeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IpRangeOutput) EndAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.EndAddress }).(pulumi.StringOutput)
}

func (o IpRangeOutput) RoleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpRange) pulumi.IntPtrOutput { return v.RoleId }).(pulumi.IntPtrOutput)
}

func (o IpRangeOutput) StartAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringOutput { return v.StartAddress }).(pulumi.StringOutput)
}

// Valid values are `active`, `reserved` and `deprecated`. Defaults to `active`.
func (o IpRangeOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o IpRangeOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpRange) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o IpRangeOutput) TenantId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpRange) pulumi.IntPtrOutput { return v.TenantId }).(pulumi.IntPtrOutput)
}

func (o IpRangeOutput) VrfId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpRange) pulumi.IntPtrOutput { return v.VrfId }).(pulumi.IntPtrOutput)
}

type IpRangeArrayOutput struct{ *pulumi.OutputState }

func (IpRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpRange)(nil)).Elem()
}

func (o IpRangeArrayOutput) ToIpRangeArrayOutput() IpRangeArrayOutput {
	return o
}

func (o IpRangeArrayOutput) ToIpRangeArrayOutputWithContext(ctx context.Context) IpRangeArrayOutput {
	return o
}

func (o IpRangeArrayOutput) Index(i pulumi.IntInput) IpRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpRange {
		return vs[0].([]*IpRange)[vs[1].(int)]
	}).(IpRangeOutput)
}

type IpRangeMapOutput struct{ *pulumi.OutputState }

func (IpRangeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpRange)(nil)).Elem()
}

func (o IpRangeMapOutput) ToIpRangeMapOutput() IpRangeMapOutput {
	return o
}

func (o IpRangeMapOutput) ToIpRangeMapOutputWithContext(ctx context.Context) IpRangeMapOutput {
	return o
}

func (o IpRangeMapOutput) MapIndex(k pulumi.StringInput) IpRangeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpRange {
		return vs[0].(map[string]*IpRange)[vs[1].(string)]
	}).(IpRangeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpRangeInput)(nil)).Elem(), &IpRange{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRangeArrayInput)(nil)).Elem(), IpRangeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRangeMapInput)(nil)).Elem(), IpRangeMap{})
	pulumi.RegisterOutputType(IpRangeOutput{})
	pulumi.RegisterOutputType(IpRangeArrayOutput{})
	pulumi.RegisterOutputType(IpRangeMapOutput{})
}
