// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#regional-internet-registries-rirs):
//
// > Regional Internet registries are responsible for the allocation of globally-routable address space. The five RIRs are ARIN, RIPE, APNIC, LACNIC, and AFRINIC. However, some address space has been set aside for internal use, such as defined in RFCs 1918 and 6598. NetBox considers these RFCs as a sort of RIR as well; that is, an authority which "owns" certain address space. There also exist lower-tier registries which serve particular geographic areas.
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
//			_, err := ipam.NewRir(ctx, "test", &ipam.RirArgs{
//				Name:        pulumi.String("test"),
//				Description: pulumi.String("my description"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Rir struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Defaults to `false`.
	IsPrivate pulumi.BoolPtrOutput `pulumi:"isPrivate"`
	Name      pulumi.StringOutput  `pulumi:"name"`
	Slug      pulumi.StringOutput  `pulumi:"slug"`
}

// NewRir registers a new resource with the given unique name, arguments, and options.
func NewRir(ctx *pulumi.Context,
	name string, args *RirArgs, opts ...pulumi.ResourceOption) (*Rir, error) {
	if args == nil {
		args = &RirArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rir
	err := ctx.RegisterResource("netbox:ipam/rir:Rir", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRir gets an existing Rir resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRir(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RirState, opts ...pulumi.ResourceOption) (*Rir, error) {
	var resource Rir
	err := ctx.ReadResource("netbox:ipam/rir:Rir", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rir resources.
type rirState struct {
	Description *string `pulumi:"description"`
	// Defaults to `false`.
	IsPrivate *bool   `pulumi:"isPrivate"`
	Name      *string `pulumi:"name"`
	Slug      *string `pulumi:"slug"`
}

type RirState struct {
	Description pulumi.StringPtrInput
	// Defaults to `false`.
	IsPrivate pulumi.BoolPtrInput
	Name      pulumi.StringPtrInput
	Slug      pulumi.StringPtrInput
}

func (RirState) ElementType() reflect.Type {
	return reflect.TypeOf((*rirState)(nil)).Elem()
}

type rirArgs struct {
	Description *string `pulumi:"description"`
	// Defaults to `false`.
	IsPrivate *bool   `pulumi:"isPrivate"`
	Name      *string `pulumi:"name"`
	Slug      *string `pulumi:"slug"`
}

// The set of arguments for constructing a Rir resource.
type RirArgs struct {
	Description pulumi.StringPtrInput
	// Defaults to `false`.
	IsPrivate pulumi.BoolPtrInput
	Name      pulumi.StringPtrInput
	Slug      pulumi.StringPtrInput
}

func (RirArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rirArgs)(nil)).Elem()
}

type RirInput interface {
	pulumi.Input

	ToRirOutput() RirOutput
	ToRirOutputWithContext(ctx context.Context) RirOutput
}

func (*Rir) ElementType() reflect.Type {
	return reflect.TypeOf((**Rir)(nil)).Elem()
}

func (i *Rir) ToRirOutput() RirOutput {
	return i.ToRirOutputWithContext(context.Background())
}

func (i *Rir) ToRirOutputWithContext(ctx context.Context) RirOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RirOutput)
}

// RirArrayInput is an input type that accepts RirArray and RirArrayOutput values.
// You can construct a concrete instance of `RirArrayInput` via:
//
//	RirArray{ RirArgs{...} }
type RirArrayInput interface {
	pulumi.Input

	ToRirArrayOutput() RirArrayOutput
	ToRirArrayOutputWithContext(context.Context) RirArrayOutput
}

type RirArray []RirInput

func (RirArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rir)(nil)).Elem()
}

func (i RirArray) ToRirArrayOutput() RirArrayOutput {
	return i.ToRirArrayOutputWithContext(context.Background())
}

func (i RirArray) ToRirArrayOutputWithContext(ctx context.Context) RirArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RirArrayOutput)
}

// RirMapInput is an input type that accepts RirMap and RirMapOutput values.
// You can construct a concrete instance of `RirMapInput` via:
//
//	RirMap{ "key": RirArgs{...} }
type RirMapInput interface {
	pulumi.Input

	ToRirMapOutput() RirMapOutput
	ToRirMapOutputWithContext(context.Context) RirMapOutput
}

type RirMap map[string]RirInput

func (RirMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rir)(nil)).Elem()
}

func (i RirMap) ToRirMapOutput() RirMapOutput {
	return i.ToRirMapOutputWithContext(context.Background())
}

func (i RirMap) ToRirMapOutputWithContext(ctx context.Context) RirMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RirMapOutput)
}

type RirOutput struct{ *pulumi.OutputState }

func (RirOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rir)(nil)).Elem()
}

func (o RirOutput) ToRirOutput() RirOutput {
	return o
}

func (o RirOutput) ToRirOutputWithContext(ctx context.Context) RirOutput {
	return o
}

func (o RirOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rir) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Defaults to `false`.
func (o RirOutput) IsPrivate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Rir) pulumi.BoolPtrOutput { return v.IsPrivate }).(pulumi.BoolPtrOutput)
}

func (o RirOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Rir) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RirOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Rir) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

type RirArrayOutput struct{ *pulumi.OutputState }

func (RirArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rir)(nil)).Elem()
}

func (o RirArrayOutput) ToRirArrayOutput() RirArrayOutput {
	return o
}

func (o RirArrayOutput) ToRirArrayOutputWithContext(ctx context.Context) RirArrayOutput {
	return o
}

func (o RirArrayOutput) Index(i pulumi.IntInput) RirOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rir {
		return vs[0].([]*Rir)[vs[1].(int)]
	}).(RirOutput)
}

type RirMapOutput struct{ *pulumi.OutputState }

func (RirMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rir)(nil)).Elem()
}

func (o RirMapOutput) ToRirMapOutput() RirMapOutput {
	return o
}

func (o RirMapOutput) ToRirMapOutputWithContext(ctx context.Context) RirMapOutput {
	return o
}

func (o RirMapOutput) MapIndex(k pulumi.StringInput) RirOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rir {
		return vs[0].(map[string]*Rir)[vs[1].(string)]
	}).(RirOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RirInput)(nil)).Elem(), &Rir{})
	pulumi.RegisterInputType(reflect.TypeOf((*RirArrayInput)(nil)).Elem(), RirArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RirMapInput)(nil)).Elem(), RirMap{})
	pulumi.RegisterOutputType(RirOutput{})
	pulumi.RegisterOutputType(RirArrayOutput{})
	pulumi.RegisterOutputType(RirMapOutput{})
}
