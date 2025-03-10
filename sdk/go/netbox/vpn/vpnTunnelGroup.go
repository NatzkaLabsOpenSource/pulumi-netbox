// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/vpn-tunnels/):
//
// > NetBox can model private tunnels formed among virtual termination points across your network. Typical tunnel implementations include GRE, IP-in-IP, and IPSec. A tunnel may be terminated to two or more device or virtual machine interfaces. For convenient organization, tunnels may be assigned to user-defined groups.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/vpn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpn.NewVpnTunnelGroup(ctx, "test", &vpn.VpnTunnelGroupArgs{
//				Name:        pulumi.String("my-tunnel-group"),
//				Description: pulumi.String("My description"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VpnTunnelGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Slug        pulumi.StringOutput    `pulumi:"slug"`
}

// NewVpnTunnelGroup registers a new resource with the given unique name, arguments, and options.
func NewVpnTunnelGroup(ctx *pulumi.Context,
	name string, args *VpnTunnelGroupArgs, opts ...pulumi.ResourceOption) (*VpnTunnelGroup, error) {
	if args == nil {
		args = &VpnTunnelGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpnTunnelGroup
	err := ctx.RegisterResource("netbox:vpn/vpnTunnelGroup:VpnTunnelGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnTunnelGroup gets an existing VpnTunnelGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnTunnelGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnTunnelGroupState, opts ...pulumi.ResourceOption) (*VpnTunnelGroup, error) {
	var resource VpnTunnelGroup
	err := ctx.ReadResource("netbox:vpn/vpnTunnelGroup:VpnTunnelGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnTunnelGroup resources.
type vpnTunnelGroupState struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Slug        *string `pulumi:"slug"`
}

type VpnTunnelGroupState struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Slug        pulumi.StringPtrInput
}

func (VpnTunnelGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnTunnelGroupState)(nil)).Elem()
}

type vpnTunnelGroupArgs struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Slug        *string `pulumi:"slug"`
}

// The set of arguments for constructing a VpnTunnelGroup resource.
type VpnTunnelGroupArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Slug        pulumi.StringPtrInput
}

func (VpnTunnelGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnTunnelGroupArgs)(nil)).Elem()
}

type VpnTunnelGroupInput interface {
	pulumi.Input

	ToVpnTunnelGroupOutput() VpnTunnelGroupOutput
	ToVpnTunnelGroupOutputWithContext(ctx context.Context) VpnTunnelGroupOutput
}

func (*VpnTunnelGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnTunnelGroup)(nil)).Elem()
}

func (i *VpnTunnelGroup) ToVpnTunnelGroupOutput() VpnTunnelGroupOutput {
	return i.ToVpnTunnelGroupOutputWithContext(context.Background())
}

func (i *VpnTunnelGroup) ToVpnTunnelGroupOutputWithContext(ctx context.Context) VpnTunnelGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnTunnelGroupOutput)
}

// VpnTunnelGroupArrayInput is an input type that accepts VpnTunnelGroupArray and VpnTunnelGroupArrayOutput values.
// You can construct a concrete instance of `VpnTunnelGroupArrayInput` via:
//
//	VpnTunnelGroupArray{ VpnTunnelGroupArgs{...} }
type VpnTunnelGroupArrayInput interface {
	pulumi.Input

	ToVpnTunnelGroupArrayOutput() VpnTunnelGroupArrayOutput
	ToVpnTunnelGroupArrayOutputWithContext(context.Context) VpnTunnelGroupArrayOutput
}

type VpnTunnelGroupArray []VpnTunnelGroupInput

func (VpnTunnelGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnTunnelGroup)(nil)).Elem()
}

func (i VpnTunnelGroupArray) ToVpnTunnelGroupArrayOutput() VpnTunnelGroupArrayOutput {
	return i.ToVpnTunnelGroupArrayOutputWithContext(context.Background())
}

func (i VpnTunnelGroupArray) ToVpnTunnelGroupArrayOutputWithContext(ctx context.Context) VpnTunnelGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnTunnelGroupArrayOutput)
}

// VpnTunnelGroupMapInput is an input type that accepts VpnTunnelGroupMap and VpnTunnelGroupMapOutput values.
// You can construct a concrete instance of `VpnTunnelGroupMapInput` via:
//
//	VpnTunnelGroupMap{ "key": VpnTunnelGroupArgs{...} }
type VpnTunnelGroupMapInput interface {
	pulumi.Input

	ToVpnTunnelGroupMapOutput() VpnTunnelGroupMapOutput
	ToVpnTunnelGroupMapOutputWithContext(context.Context) VpnTunnelGroupMapOutput
}

type VpnTunnelGroupMap map[string]VpnTunnelGroupInput

func (VpnTunnelGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnTunnelGroup)(nil)).Elem()
}

func (i VpnTunnelGroupMap) ToVpnTunnelGroupMapOutput() VpnTunnelGroupMapOutput {
	return i.ToVpnTunnelGroupMapOutputWithContext(context.Background())
}

func (i VpnTunnelGroupMap) ToVpnTunnelGroupMapOutputWithContext(ctx context.Context) VpnTunnelGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnTunnelGroupMapOutput)
}

type VpnTunnelGroupOutput struct{ *pulumi.OutputState }

func (VpnTunnelGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnTunnelGroup)(nil)).Elem()
}

func (o VpnTunnelGroupOutput) ToVpnTunnelGroupOutput() VpnTunnelGroupOutput {
	return o
}

func (o VpnTunnelGroupOutput) ToVpnTunnelGroupOutputWithContext(ctx context.Context) VpnTunnelGroupOutput {
	return o
}

func (o VpnTunnelGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnTunnelGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VpnTunnelGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnelGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnTunnelGroupOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnelGroup) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

type VpnTunnelGroupArrayOutput struct{ *pulumi.OutputState }

func (VpnTunnelGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnTunnelGroup)(nil)).Elem()
}

func (o VpnTunnelGroupArrayOutput) ToVpnTunnelGroupArrayOutput() VpnTunnelGroupArrayOutput {
	return o
}

func (o VpnTunnelGroupArrayOutput) ToVpnTunnelGroupArrayOutputWithContext(ctx context.Context) VpnTunnelGroupArrayOutput {
	return o
}

func (o VpnTunnelGroupArrayOutput) Index(i pulumi.IntInput) VpnTunnelGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnTunnelGroup {
		return vs[0].([]*VpnTunnelGroup)[vs[1].(int)]
	}).(VpnTunnelGroupOutput)
}

type VpnTunnelGroupMapOutput struct{ *pulumi.OutputState }

func (VpnTunnelGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnTunnelGroup)(nil)).Elem()
}

func (o VpnTunnelGroupMapOutput) ToVpnTunnelGroupMapOutput() VpnTunnelGroupMapOutput {
	return o
}

func (o VpnTunnelGroupMapOutput) ToVpnTunnelGroupMapOutputWithContext(ctx context.Context) VpnTunnelGroupMapOutput {
	return o
}

func (o VpnTunnelGroupMapOutput) MapIndex(k pulumi.StringInput) VpnTunnelGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnTunnelGroup {
		return vs[0].(map[string]*VpnTunnelGroup)[vs[1].(string)]
	}).(VpnTunnelGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnTunnelGroupInput)(nil)).Elem(), &VpnTunnelGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnTunnelGroupArrayInput)(nil)).Elem(), VpnTunnelGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnTunnelGroupMapInput)(nil)).Elem(), VpnTunnelGroupMap{})
	pulumi.RegisterOutputType(VpnTunnelGroupOutput{})
	pulumi.RegisterOutputType(VpnTunnelGroupArrayOutput{})
	pulumi.RegisterOutputType(VpnTunnelGroupMapOutput{})
}
