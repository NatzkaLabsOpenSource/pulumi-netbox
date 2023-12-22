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

// From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#interfaces):
//
// > Virtual machine interfaces behave similarly to device interfaces, and can be assigned to VRFs, and may have IP addresses, VLANs, and services attached to them. However, given their virtual nature, they lack properties pertaining to physical attributes. For example, VM interfaces do not have a physical type and cannot have cables attached to them.
type Interface struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Defaults to `true`.
	Enabled    pulumi.BoolPtrOutput   `pulumi:"enabled"`
	MacAddress pulumi.StringPtrOutput `pulumi:"macAddress"`
	// Valid values are `access`, `tagged` and `tagged-all`.
	Mode        pulumi.StringPtrOutput   `pulumi:"mode"`
	Mtu         pulumi.IntPtrOutput      `pulumi:"mtu"`
	Name        pulumi.StringOutput      `pulumi:"name"`
	TaggedVlans pulumi.IntArrayOutput    `pulumi:"taggedVlans"`
	Tags        pulumi.StringArrayOutput `pulumi:"tags"`
	// Deprecated: This attribute is not supported by netbox any longer. It will be removed in future versions of this provider.
	Type             pulumi.StringPtrOutput `pulumi:"type"`
	UntaggedVlan     pulumi.IntPtrOutput    `pulumi:"untaggedVlan"`
	VirtualMachineId pulumi.IntOutput       `pulumi:"virtualMachineId"`
}

// NewInterface registers a new resource with the given unique name, arguments, and options.
func NewInterface(ctx *pulumi.Context,
	name string, args *InterfaceArgs, opts ...pulumi.ResourceOption) (*Interface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VirtualMachineId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Interface
	err := ctx.RegisterResource("netbox:virt/interface:Interface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterface gets an existing Interface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterfaceState, opts ...pulumi.ResourceOption) (*Interface, error) {
	var resource Interface
	err := ctx.ReadResource("netbox:virt/interface:Interface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Interface resources.
type interfaceState struct {
	Description *string `pulumi:"description"`
	// Defaults to `true`.
	Enabled    *bool   `pulumi:"enabled"`
	MacAddress *string `pulumi:"macAddress"`
	// Valid values are `access`, `tagged` and `tagged-all`.
	Mode        *string  `pulumi:"mode"`
	Mtu         *int     `pulumi:"mtu"`
	Name        *string  `pulumi:"name"`
	TaggedVlans []int    `pulumi:"taggedVlans"`
	Tags        []string `pulumi:"tags"`
	// Deprecated: This attribute is not supported by netbox any longer. It will be removed in future versions of this provider.
	Type             *string `pulumi:"type"`
	UntaggedVlan     *int    `pulumi:"untaggedVlan"`
	VirtualMachineId *int    `pulumi:"virtualMachineId"`
}

type InterfaceState struct {
	Description pulumi.StringPtrInput
	// Defaults to `true`.
	Enabled    pulumi.BoolPtrInput
	MacAddress pulumi.StringPtrInput
	// Valid values are `access`, `tagged` and `tagged-all`.
	Mode        pulumi.StringPtrInput
	Mtu         pulumi.IntPtrInput
	Name        pulumi.StringPtrInput
	TaggedVlans pulumi.IntArrayInput
	Tags        pulumi.StringArrayInput
	// Deprecated: This attribute is not supported by netbox any longer. It will be removed in future versions of this provider.
	Type             pulumi.StringPtrInput
	UntaggedVlan     pulumi.IntPtrInput
	VirtualMachineId pulumi.IntPtrInput
}

func (InterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceState)(nil)).Elem()
}

type interfaceArgs struct {
	Description *string `pulumi:"description"`
	// Defaults to `true`.
	Enabled    *bool   `pulumi:"enabled"`
	MacAddress *string `pulumi:"macAddress"`
	// Valid values are `access`, `tagged` and `tagged-all`.
	Mode        *string  `pulumi:"mode"`
	Mtu         *int     `pulumi:"mtu"`
	Name        *string  `pulumi:"name"`
	TaggedVlans []int    `pulumi:"taggedVlans"`
	Tags        []string `pulumi:"tags"`
	// Deprecated: This attribute is not supported by netbox any longer. It will be removed in future versions of this provider.
	Type             *string `pulumi:"type"`
	UntaggedVlan     *int    `pulumi:"untaggedVlan"`
	VirtualMachineId int     `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a Interface resource.
type InterfaceArgs struct {
	Description pulumi.StringPtrInput
	// Defaults to `true`.
	Enabled    pulumi.BoolPtrInput
	MacAddress pulumi.StringPtrInput
	// Valid values are `access`, `tagged` and `tagged-all`.
	Mode        pulumi.StringPtrInput
	Mtu         pulumi.IntPtrInput
	Name        pulumi.StringPtrInput
	TaggedVlans pulumi.IntArrayInput
	Tags        pulumi.StringArrayInput
	// Deprecated: This attribute is not supported by netbox any longer. It will be removed in future versions of this provider.
	Type             pulumi.StringPtrInput
	UntaggedVlan     pulumi.IntPtrInput
	VirtualMachineId pulumi.IntInput
}

func (InterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceArgs)(nil)).Elem()
}

type InterfaceInput interface {
	pulumi.Input

	ToInterfaceOutput() InterfaceOutput
	ToInterfaceOutputWithContext(ctx context.Context) InterfaceOutput
}

func (*Interface) ElementType() reflect.Type {
	return reflect.TypeOf((**Interface)(nil)).Elem()
}

func (i *Interface) ToInterfaceOutput() InterfaceOutput {
	return i.ToInterfaceOutputWithContext(context.Background())
}

func (i *Interface) ToInterfaceOutputWithContext(ctx context.Context) InterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceOutput)
}

// InterfaceArrayInput is an input type that accepts InterfaceArray and InterfaceArrayOutput values.
// You can construct a concrete instance of `InterfaceArrayInput` via:
//
//	InterfaceArray{ InterfaceArgs{...} }
type InterfaceArrayInput interface {
	pulumi.Input

	ToInterfaceArrayOutput() InterfaceArrayOutput
	ToInterfaceArrayOutputWithContext(context.Context) InterfaceArrayOutput
}

type InterfaceArray []InterfaceInput

func (InterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Interface)(nil)).Elem()
}

func (i InterfaceArray) ToInterfaceArrayOutput() InterfaceArrayOutput {
	return i.ToInterfaceArrayOutputWithContext(context.Background())
}

func (i InterfaceArray) ToInterfaceArrayOutputWithContext(ctx context.Context) InterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceArrayOutput)
}

// InterfaceMapInput is an input type that accepts InterfaceMap and InterfaceMapOutput values.
// You can construct a concrete instance of `InterfaceMapInput` via:
//
//	InterfaceMap{ "key": InterfaceArgs{...} }
type InterfaceMapInput interface {
	pulumi.Input

	ToInterfaceMapOutput() InterfaceMapOutput
	ToInterfaceMapOutputWithContext(context.Context) InterfaceMapOutput
}

type InterfaceMap map[string]InterfaceInput

func (InterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Interface)(nil)).Elem()
}

func (i InterfaceMap) ToInterfaceMapOutput() InterfaceMapOutput {
	return i.ToInterfaceMapOutputWithContext(context.Background())
}

func (i InterfaceMap) ToInterfaceMapOutputWithContext(ctx context.Context) InterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceMapOutput)
}

type InterfaceOutput struct{ *pulumi.OutputState }

func (InterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Interface)(nil)).Elem()
}

func (o InterfaceOutput) ToInterfaceOutput() InterfaceOutput {
	return o
}

func (o InterfaceOutput) ToInterfaceOutputWithContext(ctx context.Context) InterfaceOutput {
	return o
}

func (o InterfaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Interface) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Defaults to `true`.
func (o InterfaceOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Interface) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o InterfaceOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Interface) pulumi.StringPtrOutput { return v.MacAddress }).(pulumi.StringPtrOutput)
}

// Valid values are `access`, `tagged` and `tagged-all`.
func (o InterfaceOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Interface) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o InterfaceOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Interface) pulumi.IntPtrOutput { return v.Mtu }).(pulumi.IntPtrOutput)
}

func (o InterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Interface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o InterfaceOutput) TaggedVlans() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Interface) pulumi.IntArrayOutput { return v.TaggedVlans }).(pulumi.IntArrayOutput)
}

func (o InterfaceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Interface) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Deprecated: This attribute is not supported by netbox any longer. It will be removed in future versions of this provider.
func (o InterfaceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Interface) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o InterfaceOutput) UntaggedVlan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Interface) pulumi.IntPtrOutput { return v.UntaggedVlan }).(pulumi.IntPtrOutput)
}

func (o InterfaceOutput) VirtualMachineId() pulumi.IntOutput {
	return o.ApplyT(func(v *Interface) pulumi.IntOutput { return v.VirtualMachineId }).(pulumi.IntOutput)
}

type InterfaceArrayOutput struct{ *pulumi.OutputState }

func (InterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Interface)(nil)).Elem()
}

func (o InterfaceArrayOutput) ToInterfaceArrayOutput() InterfaceArrayOutput {
	return o
}

func (o InterfaceArrayOutput) ToInterfaceArrayOutputWithContext(ctx context.Context) InterfaceArrayOutput {
	return o
}

func (o InterfaceArrayOutput) Index(i pulumi.IntInput) InterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Interface {
		return vs[0].([]*Interface)[vs[1].(int)]
	}).(InterfaceOutput)
}

type InterfaceMapOutput struct{ *pulumi.OutputState }

func (InterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Interface)(nil)).Elem()
}

func (o InterfaceMapOutput) ToInterfaceMapOutput() InterfaceMapOutput {
	return o
}

func (o InterfaceMapOutput) ToInterfaceMapOutputWithContext(ctx context.Context) InterfaceMapOutput {
	return o
}

func (o InterfaceMapOutput) MapIndex(k pulumi.StringInput) InterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Interface {
		return vs[0].(map[string]*Interface)[vs[1].(string)]
	}).(InterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InterfaceInput)(nil)).Elem(), &Interface{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfaceArrayInput)(nil)).Elem(), InterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfaceMapInput)(nil)).Elem(), InterfaceMap{})
	pulumi.RegisterOutputType(InterfaceOutput{})
	pulumi.RegisterOutputType(InterfaceArrayOutput{})
	pulumi.RegisterOutputType(InterfaceMapOutput{})
}
