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

// > A VLAN Group represents a collection of VLANs. Generally, these are limited by one of a number of scopes such as "Site" or "Virtualization Cluster".
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
//			_, err := ipam.NewVlanGroup(ctx, "example1", &ipam.VlanGroupArgs{
//				Slug:   pulumi.String("example1"),
//				MinVid: pulumi.Int(1),
//				MaxVid: pulumi.Int(4094),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ipam.NewVlanGroup(ctx, "example2", &ipam.VlanGroupArgs{
//				Slug:        pulumi.String("example2"),
//				MinVid:      pulumi.Int(1),
//				MaxVid:      pulumi.Int(4094),
//				ScopeType:   pulumi.String("dcim.site"),
//				ScopeId:     pulumi.Any(netbox_site.Example.Id),
//				Description: pulumi.String("Second Example VLAN Group"),
//				Tags: pulumi.StringArray{
//					netbox_tag.Example.Id,
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
type VlanGroup struct {
	pulumi.CustomResourceState

	// Defaults to `""`.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	MaxVid      pulumi.IntOutput       `pulumi:"maxVid"`
	MinVid      pulumi.IntOutput       `pulumi:"minVid"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	// Required when `scopeType` is set.
	ScopeId pulumi.IntPtrOutput `pulumi:"scopeId"`
	// Valid values are `active`, `container`, `reserved` and `deprecated`.
	ScopeType pulumi.StringPtrOutput   `pulumi:"scopeType"`
	Slug      pulumi.StringOutput      `pulumi:"slug"`
	Tags      pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewVlanGroup registers a new resource with the given unique name, arguments, and options.
func NewVlanGroup(ctx *pulumi.Context,
	name string, args *VlanGroupArgs, opts ...pulumi.ResourceOption) (*VlanGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxVid == nil {
		return nil, errors.New("invalid value for required argument 'MaxVid'")
	}
	if args.MinVid == nil {
		return nil, errors.New("invalid value for required argument 'MinVid'")
	}
	if args.Slug == nil {
		return nil, errors.New("invalid value for required argument 'Slug'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VlanGroup
	err := ctx.RegisterResource("netbox:ipam/vlanGroup:VlanGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVlanGroup gets an existing VlanGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVlanGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VlanGroupState, opts ...pulumi.ResourceOption) (*VlanGroup, error) {
	var resource VlanGroup
	err := ctx.ReadResource("netbox:ipam/vlanGroup:VlanGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VlanGroup resources.
type vlanGroupState struct {
	// Defaults to `""`.
	Description *string `pulumi:"description"`
	MaxVid      *int    `pulumi:"maxVid"`
	MinVid      *int    `pulumi:"minVid"`
	Name        *string `pulumi:"name"`
	// Required when `scopeType` is set.
	ScopeId *int `pulumi:"scopeId"`
	// Valid values are `active`, `container`, `reserved` and `deprecated`.
	ScopeType *string  `pulumi:"scopeType"`
	Slug      *string  `pulumi:"slug"`
	Tags      []string `pulumi:"tags"`
}

type VlanGroupState struct {
	// Defaults to `""`.
	Description pulumi.StringPtrInput
	MaxVid      pulumi.IntPtrInput
	MinVid      pulumi.IntPtrInput
	Name        pulumi.StringPtrInput
	// Required when `scopeType` is set.
	ScopeId pulumi.IntPtrInput
	// Valid values are `active`, `container`, `reserved` and `deprecated`.
	ScopeType pulumi.StringPtrInput
	Slug      pulumi.StringPtrInput
	Tags      pulumi.StringArrayInput
}

func (VlanGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanGroupState)(nil)).Elem()
}

type vlanGroupArgs struct {
	// Defaults to `""`.
	Description *string `pulumi:"description"`
	MaxVid      int     `pulumi:"maxVid"`
	MinVid      int     `pulumi:"minVid"`
	Name        *string `pulumi:"name"`
	// Required when `scopeType` is set.
	ScopeId *int `pulumi:"scopeId"`
	// Valid values are `active`, `container`, `reserved` and `deprecated`.
	ScopeType *string  `pulumi:"scopeType"`
	Slug      string   `pulumi:"slug"`
	Tags      []string `pulumi:"tags"`
}

// The set of arguments for constructing a VlanGroup resource.
type VlanGroupArgs struct {
	// Defaults to `""`.
	Description pulumi.StringPtrInput
	MaxVid      pulumi.IntInput
	MinVid      pulumi.IntInput
	Name        pulumi.StringPtrInput
	// Required when `scopeType` is set.
	ScopeId pulumi.IntPtrInput
	// Valid values are `active`, `container`, `reserved` and `deprecated`.
	ScopeType pulumi.StringPtrInput
	Slug      pulumi.StringInput
	Tags      pulumi.StringArrayInput
}

func (VlanGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanGroupArgs)(nil)).Elem()
}

type VlanGroupInput interface {
	pulumi.Input

	ToVlanGroupOutput() VlanGroupOutput
	ToVlanGroupOutputWithContext(ctx context.Context) VlanGroupOutput
}

func (*VlanGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**VlanGroup)(nil)).Elem()
}

func (i *VlanGroup) ToVlanGroupOutput() VlanGroupOutput {
	return i.ToVlanGroupOutputWithContext(context.Background())
}

func (i *VlanGroup) ToVlanGroupOutputWithContext(ctx context.Context) VlanGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanGroupOutput)
}

// VlanGroupArrayInput is an input type that accepts VlanGroupArray and VlanGroupArrayOutput values.
// You can construct a concrete instance of `VlanGroupArrayInput` via:
//
//	VlanGroupArray{ VlanGroupArgs{...} }
type VlanGroupArrayInput interface {
	pulumi.Input

	ToVlanGroupArrayOutput() VlanGroupArrayOutput
	ToVlanGroupArrayOutputWithContext(context.Context) VlanGroupArrayOutput
}

type VlanGroupArray []VlanGroupInput

func (VlanGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VlanGroup)(nil)).Elem()
}

func (i VlanGroupArray) ToVlanGroupArrayOutput() VlanGroupArrayOutput {
	return i.ToVlanGroupArrayOutputWithContext(context.Background())
}

func (i VlanGroupArray) ToVlanGroupArrayOutputWithContext(ctx context.Context) VlanGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanGroupArrayOutput)
}

// VlanGroupMapInput is an input type that accepts VlanGroupMap and VlanGroupMapOutput values.
// You can construct a concrete instance of `VlanGroupMapInput` via:
//
//	VlanGroupMap{ "key": VlanGroupArgs{...} }
type VlanGroupMapInput interface {
	pulumi.Input

	ToVlanGroupMapOutput() VlanGroupMapOutput
	ToVlanGroupMapOutputWithContext(context.Context) VlanGroupMapOutput
}

type VlanGroupMap map[string]VlanGroupInput

func (VlanGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VlanGroup)(nil)).Elem()
}

func (i VlanGroupMap) ToVlanGroupMapOutput() VlanGroupMapOutput {
	return i.ToVlanGroupMapOutputWithContext(context.Background())
}

func (i VlanGroupMap) ToVlanGroupMapOutputWithContext(ctx context.Context) VlanGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanGroupMapOutput)
}

type VlanGroupOutput struct{ *pulumi.OutputState }

func (VlanGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VlanGroup)(nil)).Elem()
}

func (o VlanGroupOutput) ToVlanGroupOutput() VlanGroupOutput {
	return o
}

func (o VlanGroupOutput) ToVlanGroupOutputWithContext(ctx context.Context) VlanGroupOutput {
	return o
}

// Defaults to `""`.
func (o VlanGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VlanGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VlanGroupOutput) MaxVid() pulumi.IntOutput {
	return o.ApplyT(func(v *VlanGroup) pulumi.IntOutput { return v.MaxVid }).(pulumi.IntOutput)
}

func (o VlanGroupOutput) MinVid() pulumi.IntOutput {
	return o.ApplyT(func(v *VlanGroup) pulumi.IntOutput { return v.MinVid }).(pulumi.IntOutput)
}

func (o VlanGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VlanGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Required when `scopeType` is set.
func (o VlanGroupOutput) ScopeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VlanGroup) pulumi.IntPtrOutput { return v.ScopeId }).(pulumi.IntPtrOutput)
}

// Valid values are `active`, `container`, `reserved` and `deprecated`.
func (o VlanGroupOutput) ScopeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VlanGroup) pulumi.StringPtrOutput { return v.ScopeType }).(pulumi.StringPtrOutput)
}

func (o VlanGroupOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *VlanGroup) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

func (o VlanGroupOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VlanGroup) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type VlanGroupArrayOutput struct{ *pulumi.OutputState }

func (VlanGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VlanGroup)(nil)).Elem()
}

func (o VlanGroupArrayOutput) ToVlanGroupArrayOutput() VlanGroupArrayOutput {
	return o
}

func (o VlanGroupArrayOutput) ToVlanGroupArrayOutputWithContext(ctx context.Context) VlanGroupArrayOutput {
	return o
}

func (o VlanGroupArrayOutput) Index(i pulumi.IntInput) VlanGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VlanGroup {
		return vs[0].([]*VlanGroup)[vs[1].(int)]
	}).(VlanGroupOutput)
}

type VlanGroupMapOutput struct{ *pulumi.OutputState }

func (VlanGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VlanGroup)(nil)).Elem()
}

func (o VlanGroupMapOutput) ToVlanGroupMapOutput() VlanGroupMapOutput {
	return o
}

func (o VlanGroupMapOutput) ToVlanGroupMapOutputWithContext(ctx context.Context) VlanGroupMapOutput {
	return o
}

func (o VlanGroupMapOutput) MapIndex(k pulumi.StringInput) VlanGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VlanGroup {
		return vs[0].(map[string]*VlanGroup)[vs[1].(string)]
	}).(VlanGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VlanGroupInput)(nil)).Elem(), &VlanGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanGroupArrayInput)(nil)).Elem(), VlanGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanGroupMapInput)(nil)).Elem(), VlanGroupMap{})
	pulumi.RegisterOutputType(VlanGroupOutput{})
	pulumi.RegisterOutputType(VlanGroupArrayOutput{})
	pulumi.RegisterOutputType(VlanGroupMapOutput{})
}