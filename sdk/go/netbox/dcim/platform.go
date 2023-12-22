// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcim

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/devices/#platforms):
//
// > A platform defines the type of software running on a device or virtual machine. This can be helpful to model when it is necessary to distinguish between different versions or feature sets. Note that two devices of the same type may be assigned different platforms: For example, one Juniper MX240 might run Junos 14 while another runs Junos 15.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/dcim"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dcim.NewPlatform(ctx, "pANOS", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Platform struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
	Slug pulumi.StringOutput `pulumi:"slug"`
}

// NewPlatform registers a new resource with the given unique name, arguments, and options.
func NewPlatform(ctx *pulumi.Context,
	name string, args *PlatformArgs, opts ...pulumi.ResourceOption) (*Platform, error) {
	if args == nil {
		args = &PlatformArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Platform
	err := ctx.RegisterResource("netbox:dcim/platform:Platform", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlatform gets an existing Platform resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlatform(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlatformState, opts ...pulumi.ResourceOption) (*Platform, error) {
	var resource Platform
	err := ctx.ReadResource("netbox:dcim/platform:Platform", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Platform resources.
type platformState struct {
	Name *string `pulumi:"name"`
	Slug *string `pulumi:"slug"`
}

type PlatformState struct {
	Name pulumi.StringPtrInput
	Slug pulumi.StringPtrInput
}

func (PlatformState) ElementType() reflect.Type {
	return reflect.TypeOf((*platformState)(nil)).Elem()
}

type platformArgs struct {
	Name *string `pulumi:"name"`
	Slug *string `pulumi:"slug"`
}

// The set of arguments for constructing a Platform resource.
type PlatformArgs struct {
	Name pulumi.StringPtrInput
	Slug pulumi.StringPtrInput
}

func (PlatformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*platformArgs)(nil)).Elem()
}

type PlatformInput interface {
	pulumi.Input

	ToPlatformOutput() PlatformOutput
	ToPlatformOutputWithContext(ctx context.Context) PlatformOutput
}

func (*Platform) ElementType() reflect.Type {
	return reflect.TypeOf((**Platform)(nil)).Elem()
}

func (i *Platform) ToPlatformOutput() PlatformOutput {
	return i.ToPlatformOutputWithContext(context.Background())
}

func (i *Platform) ToPlatformOutputWithContext(ctx context.Context) PlatformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformOutput)
}

// PlatformArrayInput is an input type that accepts PlatformArray and PlatformArrayOutput values.
// You can construct a concrete instance of `PlatformArrayInput` via:
//
//	PlatformArray{ PlatformArgs{...} }
type PlatformArrayInput interface {
	pulumi.Input

	ToPlatformArrayOutput() PlatformArrayOutput
	ToPlatformArrayOutputWithContext(context.Context) PlatformArrayOutput
}

type PlatformArray []PlatformInput

func (PlatformArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Platform)(nil)).Elem()
}

func (i PlatformArray) ToPlatformArrayOutput() PlatformArrayOutput {
	return i.ToPlatformArrayOutputWithContext(context.Background())
}

func (i PlatformArray) ToPlatformArrayOutputWithContext(ctx context.Context) PlatformArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformArrayOutput)
}

// PlatformMapInput is an input type that accepts PlatformMap and PlatformMapOutput values.
// You can construct a concrete instance of `PlatformMapInput` via:
//
//	PlatformMap{ "key": PlatformArgs{...} }
type PlatformMapInput interface {
	pulumi.Input

	ToPlatformMapOutput() PlatformMapOutput
	ToPlatformMapOutputWithContext(context.Context) PlatformMapOutput
}

type PlatformMap map[string]PlatformInput

func (PlatformMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Platform)(nil)).Elem()
}

func (i PlatformMap) ToPlatformMapOutput() PlatformMapOutput {
	return i.ToPlatformMapOutputWithContext(context.Background())
}

func (i PlatformMap) ToPlatformMapOutputWithContext(ctx context.Context) PlatformMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformMapOutput)
}

type PlatformOutput struct{ *pulumi.OutputState }

func (PlatformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Platform)(nil)).Elem()
}

func (o PlatformOutput) ToPlatformOutput() PlatformOutput {
	return o
}

func (o PlatformOutput) ToPlatformOutputWithContext(ctx context.Context) PlatformOutput {
	return o
}

func (o PlatformOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Platform) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PlatformOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Platform) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

type PlatformArrayOutput struct{ *pulumi.OutputState }

func (PlatformArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Platform)(nil)).Elem()
}

func (o PlatformArrayOutput) ToPlatformArrayOutput() PlatformArrayOutput {
	return o
}

func (o PlatformArrayOutput) ToPlatformArrayOutputWithContext(ctx context.Context) PlatformArrayOutput {
	return o
}

func (o PlatformArrayOutput) Index(i pulumi.IntInput) PlatformOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Platform {
		return vs[0].([]*Platform)[vs[1].(int)]
	}).(PlatformOutput)
}

type PlatformMapOutput struct{ *pulumi.OutputState }

func (PlatformMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Platform)(nil)).Elem()
}

func (o PlatformMapOutput) ToPlatformMapOutput() PlatformMapOutput {
	return o
}

func (o PlatformMapOutput) ToPlatformMapOutputWithContext(ctx context.Context) PlatformMapOutput {
	return o
}

func (o PlatformMapOutput) MapIndex(k pulumi.StringInput) PlatformOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Platform {
		return vs[0].(map[string]*Platform)[vs[1].(string)]
	}).(PlatformOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlatformInput)(nil)).Elem(), &Platform{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlatformArrayInput)(nil)).Elem(), PlatformArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlatformMapInput)(nil)).Elem(), PlatformMap{})
	pulumi.RegisterOutputType(PlatformOutput{})
	pulumi.RegisterOutputType(PlatformArrayOutput{})
	pulumi.RegisterOutputType(PlatformMapOutput{})
}
