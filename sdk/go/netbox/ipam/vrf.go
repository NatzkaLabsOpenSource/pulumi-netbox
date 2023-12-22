// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#virtual-routing-and-forwarding-vrf):
//
// > A VRF object in NetBox represents a virtual routing and forwarding (VRF) domain. Each VRF is essentially a separate routing table. VRFs are commonly used to isolate customers or organizations from one another within a network, or to route overlapping address space (e.g. multiple instances of the 10.0.0.0/8 space). Each VRF may be assigned to a specific tenant to aid in organizing the available IP space by customer or internal user.
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
//			_, err := ipam.NewVrf(ctx, "custAProd", &ipam.VrfArgs{
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
type Vrf struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput   `pulumi:"description"`
	Name        pulumi.StringOutput      `pulumi:"name"`
	Tags        pulumi.StringArrayOutput `pulumi:"tags"`
	TenantId    pulumi.IntPtrOutput      `pulumi:"tenantId"`
}

// NewVrf registers a new resource with the given unique name, arguments, and options.
func NewVrf(ctx *pulumi.Context,
	name string, args *VrfArgs, opts ...pulumi.ResourceOption) (*Vrf, error) {
	if args == nil {
		args = &VrfArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vrf
	err := ctx.RegisterResource("netbox:ipam/vrf:Vrf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVrf gets an existing Vrf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVrf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VrfState, opts ...pulumi.ResourceOption) (*Vrf, error) {
	var resource Vrf
	err := ctx.ReadResource("netbox:ipam/vrf:Vrf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vrf resources.
type vrfState struct {
	Description *string  `pulumi:"description"`
	Name        *string  `pulumi:"name"`
	Tags        []string `pulumi:"tags"`
	TenantId    *int     `pulumi:"tenantId"`
}

type VrfState struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Tags        pulumi.StringArrayInput
	TenantId    pulumi.IntPtrInput
}

func (VrfState) ElementType() reflect.Type {
	return reflect.TypeOf((*vrfState)(nil)).Elem()
}

type vrfArgs struct {
	Description *string  `pulumi:"description"`
	Name        *string  `pulumi:"name"`
	Tags        []string `pulumi:"tags"`
	TenantId    *int     `pulumi:"tenantId"`
}

// The set of arguments for constructing a Vrf resource.
type VrfArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Tags        pulumi.StringArrayInput
	TenantId    pulumi.IntPtrInput
}

func (VrfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vrfArgs)(nil)).Elem()
}

type VrfInput interface {
	pulumi.Input

	ToVrfOutput() VrfOutput
	ToVrfOutputWithContext(ctx context.Context) VrfOutput
}

func (*Vrf) ElementType() reflect.Type {
	return reflect.TypeOf((**Vrf)(nil)).Elem()
}

func (i *Vrf) ToVrfOutput() VrfOutput {
	return i.ToVrfOutputWithContext(context.Background())
}

func (i *Vrf) ToVrfOutputWithContext(ctx context.Context) VrfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrfOutput)
}

// VrfArrayInput is an input type that accepts VrfArray and VrfArrayOutput values.
// You can construct a concrete instance of `VrfArrayInput` via:
//
//	VrfArray{ VrfArgs{...} }
type VrfArrayInput interface {
	pulumi.Input

	ToVrfArrayOutput() VrfArrayOutput
	ToVrfArrayOutputWithContext(context.Context) VrfArrayOutput
}

type VrfArray []VrfInput

func (VrfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vrf)(nil)).Elem()
}

func (i VrfArray) ToVrfArrayOutput() VrfArrayOutput {
	return i.ToVrfArrayOutputWithContext(context.Background())
}

func (i VrfArray) ToVrfArrayOutputWithContext(ctx context.Context) VrfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrfArrayOutput)
}

// VrfMapInput is an input type that accepts VrfMap and VrfMapOutput values.
// You can construct a concrete instance of `VrfMapInput` via:
//
//	VrfMap{ "key": VrfArgs{...} }
type VrfMapInput interface {
	pulumi.Input

	ToVrfMapOutput() VrfMapOutput
	ToVrfMapOutputWithContext(context.Context) VrfMapOutput
}

type VrfMap map[string]VrfInput

func (VrfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vrf)(nil)).Elem()
}

func (i VrfMap) ToVrfMapOutput() VrfMapOutput {
	return i.ToVrfMapOutputWithContext(context.Background())
}

func (i VrfMap) ToVrfMapOutputWithContext(ctx context.Context) VrfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrfMapOutput)
}

type VrfOutput struct{ *pulumi.OutputState }

func (VrfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vrf)(nil)).Elem()
}

func (o VrfOutput) ToVrfOutput() VrfOutput {
	return o
}

func (o VrfOutput) ToVrfOutputWithContext(ctx context.Context) VrfOutput {
	return o
}

func (o VrfOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vrf) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VrfOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vrf) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VrfOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vrf) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o VrfOutput) TenantId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Vrf) pulumi.IntPtrOutput { return v.TenantId }).(pulumi.IntPtrOutput)
}

type VrfArrayOutput struct{ *pulumi.OutputState }

func (VrfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vrf)(nil)).Elem()
}

func (o VrfArrayOutput) ToVrfArrayOutput() VrfArrayOutput {
	return o
}

func (o VrfArrayOutput) ToVrfArrayOutputWithContext(ctx context.Context) VrfArrayOutput {
	return o
}

func (o VrfArrayOutput) Index(i pulumi.IntInput) VrfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vrf {
		return vs[0].([]*Vrf)[vs[1].(int)]
	}).(VrfOutput)
}

type VrfMapOutput struct{ *pulumi.OutputState }

func (VrfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vrf)(nil)).Elem()
}

func (o VrfMapOutput) ToVrfMapOutput() VrfMapOutput {
	return o
}

func (o VrfMapOutput) ToVrfMapOutputWithContext(ctx context.Context) VrfMapOutput {
	return o
}

func (o VrfMapOutput) MapIndex(k pulumi.StringInput) VrfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vrf {
		return vs[0].(map[string]*Vrf)[vs[1].(string)]
	}).(VrfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VrfInput)(nil)).Elem(), &Vrf{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrfArrayInput)(nil)).Elem(), VrfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrfMapInput)(nil)).Elem(), VrfMap{})
	pulumi.RegisterOutputType(VrfOutput{})
	pulumi.RegisterOutputType(VrfArrayOutput{})
	pulumi.RegisterOutputType(VrfMapOutput{})
}
