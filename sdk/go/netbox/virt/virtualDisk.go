// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package virt

import (
	"context"
	"reflect"

	"errors"
	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/models/virtualization/virtualdisk/):
//
//	> A virtual disk is used to model discrete virtual hard disks assigned to virtual machines.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/virt"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Assumes vmw-cluster-01 exists in Netbox
//			vmwCluster01, err := virt.LookupCluster(ctx, &virt.LookupClusterArgs{
//				Name: pulumi.StringRef("vmw-cluster-01"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			baseVm, err := virt.NewVirtualMachine(ctx, "base_vm", &virt.VirtualMachineArgs{
//				ClusterId: pulumi.String(vmwCluster01.Id),
//				Name:      pulumi.String("myvm-1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = virt.NewVirtualDisk(ctx, "example", &virt.VirtualDiskArgs{
//				Name:             pulumi.String("disk-01"),
//				Description:      pulumi.String("Main disk"),
//				SizeMb:           pulumi.Int(50),
//				VirtualMachineId: baseVm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VirtualDisk struct {
	pulumi.CustomResourceState

	CustomFields     pulumi.StringMapOutput   `pulumi:"customFields"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	SizeMb           pulumi.IntOutput         `pulumi:"sizeMb"`
	Tags             pulumi.StringArrayOutput `pulumi:"tags"`
	VirtualMachineId pulumi.IntOutput         `pulumi:"virtualMachineId"`
}

// NewVirtualDisk registers a new resource with the given unique name, arguments, and options.
func NewVirtualDisk(ctx *pulumi.Context,
	name string, args *VirtualDiskArgs, opts ...pulumi.ResourceOption) (*VirtualDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SizeMb == nil {
		return nil, errors.New("invalid value for required argument 'SizeMb'")
	}
	if args.VirtualMachineId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualDisk
	err := ctx.RegisterResource("netbox:virt/virtualDisk:VirtualDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualDisk gets an existing VirtualDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualDiskState, opts ...pulumi.ResourceOption) (*VirtualDisk, error) {
	var resource VirtualDisk
	err := ctx.ReadResource("netbox:virt/virtualDisk:VirtualDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualDisk resources.
type virtualDiskState struct {
	CustomFields     map[string]string `pulumi:"customFields"`
	Description      *string           `pulumi:"description"`
	Name             *string           `pulumi:"name"`
	SizeMb           *int              `pulumi:"sizeMb"`
	Tags             []string          `pulumi:"tags"`
	VirtualMachineId *int              `pulumi:"virtualMachineId"`
}

type VirtualDiskState struct {
	CustomFields     pulumi.StringMapInput
	Description      pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	SizeMb           pulumi.IntPtrInput
	Tags             pulumi.StringArrayInput
	VirtualMachineId pulumi.IntPtrInput
}

func (VirtualDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDiskState)(nil)).Elem()
}

type virtualDiskArgs struct {
	CustomFields     map[string]string `pulumi:"customFields"`
	Description      *string           `pulumi:"description"`
	Name             *string           `pulumi:"name"`
	SizeMb           int               `pulumi:"sizeMb"`
	Tags             []string          `pulumi:"tags"`
	VirtualMachineId int               `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a VirtualDisk resource.
type VirtualDiskArgs struct {
	CustomFields     pulumi.StringMapInput
	Description      pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	SizeMb           pulumi.IntInput
	Tags             pulumi.StringArrayInput
	VirtualMachineId pulumi.IntInput
}

func (VirtualDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDiskArgs)(nil)).Elem()
}

type VirtualDiskInput interface {
	pulumi.Input

	ToVirtualDiskOutput() VirtualDiskOutput
	ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput
}

func (*VirtualDisk) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDisk)(nil)).Elem()
}

func (i *VirtualDisk) ToVirtualDiskOutput() VirtualDiskOutput {
	return i.ToVirtualDiskOutputWithContext(context.Background())
}

func (i *VirtualDisk) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskOutput)
}

// VirtualDiskArrayInput is an input type that accepts VirtualDiskArray and VirtualDiskArrayOutput values.
// You can construct a concrete instance of `VirtualDiskArrayInput` via:
//
//	VirtualDiskArray{ VirtualDiskArgs{...} }
type VirtualDiskArrayInput interface {
	pulumi.Input

	ToVirtualDiskArrayOutput() VirtualDiskArrayOutput
	ToVirtualDiskArrayOutputWithContext(context.Context) VirtualDiskArrayOutput
}

type VirtualDiskArray []VirtualDiskInput

func (VirtualDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return i.ToVirtualDiskArrayOutputWithContext(context.Background())
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskArrayOutput)
}

// VirtualDiskMapInput is an input type that accepts VirtualDiskMap and VirtualDiskMapOutput values.
// You can construct a concrete instance of `VirtualDiskMapInput` via:
//
//	VirtualDiskMap{ "key": VirtualDiskArgs{...} }
type VirtualDiskMapInput interface {
	pulumi.Input

	ToVirtualDiskMapOutput() VirtualDiskMapOutput
	ToVirtualDiskMapOutputWithContext(context.Context) VirtualDiskMapOutput
}

type VirtualDiskMap map[string]VirtualDiskInput

func (VirtualDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskMap) ToVirtualDiskMapOutput() VirtualDiskMapOutput {
	return i.ToVirtualDiskMapOutputWithContext(context.Background())
}

func (i VirtualDiskMap) ToVirtualDiskMapOutputWithContext(ctx context.Context) VirtualDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskMapOutput)
}

type VirtualDiskOutput struct{ *pulumi.OutputState }

func (VirtualDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskOutput) ToVirtualDiskOutput() VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

func (o VirtualDiskOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualDiskOutput) SizeMb() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.IntOutput { return v.SizeMb }).(pulumi.IntOutput)
}

func (o VirtualDiskOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o VirtualDiskOutput) VirtualMachineId() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.IntOutput { return v.VirtualMachineId }).(pulumi.IntOutput)
}

type VirtualDiskArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) Index(i pulumi.IntInput) VirtualDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualDisk {
		return vs[0].([]*VirtualDisk)[vs[1].(int)]
	}).(VirtualDiskOutput)
}

type VirtualDiskMapOutput struct{ *pulumi.OutputState }

func (VirtualDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskMapOutput) ToVirtualDiskMapOutput() VirtualDiskMapOutput {
	return o
}

func (o VirtualDiskMapOutput) ToVirtualDiskMapOutputWithContext(ctx context.Context) VirtualDiskMapOutput {
	return o
}

func (o VirtualDiskMapOutput) MapIndex(k pulumi.StringInput) VirtualDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualDisk {
		return vs[0].(map[string]*VirtualDisk)[vs[1].(string)]
	}).(VirtualDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskInput)(nil)).Elem(), &VirtualDisk{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskArrayInput)(nil)).Elem(), VirtualDiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskMapInput)(nil)).Elem(), VirtualDiskMap{})
	pulumi.RegisterOutputType(VirtualDiskOutput{})
	pulumi.RegisterOutputType(VirtualDiskArrayOutput{})
	pulumi.RegisterOutputType(VirtualDiskMapOutput{})
}
