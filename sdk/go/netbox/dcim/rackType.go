// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcim

import (
	"context"
	"reflect"

	"errors"
	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://netboxlabs.com/docs/netbox/en/stable/models/dcim/racktype/):
//
// > A rack type defines the physical characteristics of a particular model of rack.
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
//			test, err := dcim.NewManufacturer(ctx, "test", &dcim.ManufacturerArgs{
//				Name: pulumi.String("my-manufacturer"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dcim.NewRackType(ctx, "test", &dcim.RackTypeArgs{
//				Model:           pulumi.String("mymodel"),
//				ManufacturerId:  test.ID(),
//				Width:           pulumi.Int(19),
//				UHeight:         pulumi.Int(48),
//				StartingUnit:    pulumi.Int(1),
//				FormFactor:      pulumi.String("2-post-frame"),
//				Description:     pulumi.String("My description"),
//				OuterWidth:      pulumi.Int(10),
//				OuterDepth:      pulumi.Int(15),
//				OuterUnit:       pulumi.String("mm"),
//				Weight:          pulumi.Float64(15),
//				MaxWeight:       pulumi.Int(20),
//				WeightUnit:      pulumi.String("kg"),
//				MountingDepthMm: pulumi.Int(21),
//				Comments:        pulumi.String("My comments"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type RackType struct {
	pulumi.CustomResourceState

	Comments    pulumi.StringPtrOutput `pulumi:"comments"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	FormFactor      pulumi.StringOutput `pulumi:"formFactor"`
	ManufacturerId  pulumi.IntPtrOutput `pulumi:"manufacturerId"`
	MaxWeight       pulumi.IntPtrOutput `pulumi:"maxWeight"`
	Model           pulumi.StringOutput `pulumi:"model"`
	MountingDepthMm pulumi.IntPtrOutput `pulumi:"mountingDepthMm"`
	OuterDepth      pulumi.IntPtrOutput `pulumi:"outerDepth"`
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit    pulumi.StringPtrOutput   `pulumi:"outerUnit"`
	OuterWidth   pulumi.IntPtrOutput      `pulumi:"outerWidth"`
	Slug         pulumi.StringOutput      `pulumi:"slug"`
	StartingUnit pulumi.IntOutput         `pulumi:"startingUnit"`
	Tags         pulumi.StringArrayOutput `pulumi:"tags"`
	UHeight      pulumi.IntOutput         `pulumi:"uHeight"`
	Weight       pulumi.Float64PtrOutput  `pulumi:"weight"`
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit pulumi.StringPtrOutput `pulumi:"weightUnit"`
	// Valid values are `10`, `19`, `21` and `23`.
	Width pulumi.IntOutput `pulumi:"width"`
}

// NewRackType registers a new resource with the given unique name, arguments, and options.
func NewRackType(ctx *pulumi.Context,
	name string, args *RackTypeArgs, opts ...pulumi.ResourceOption) (*RackType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FormFactor == nil {
		return nil, errors.New("invalid value for required argument 'FormFactor'")
	}
	if args.Model == nil {
		return nil, errors.New("invalid value for required argument 'Model'")
	}
	if args.StartingUnit == nil {
		return nil, errors.New("invalid value for required argument 'StartingUnit'")
	}
	if args.UHeight == nil {
		return nil, errors.New("invalid value for required argument 'UHeight'")
	}
	if args.Width == nil {
		return nil, errors.New("invalid value for required argument 'Width'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RackType
	err := ctx.RegisterResource("netbox:dcim/rackType:RackType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRackType gets an existing RackType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRackType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RackTypeState, opts ...pulumi.ResourceOption) (*RackType, error) {
	var resource RackType
	err := ctx.ReadResource("netbox:dcim/rackType:RackType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RackType resources.
type rackTypeState struct {
	Comments    *string `pulumi:"comments"`
	Description *string `pulumi:"description"`
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	FormFactor      *string `pulumi:"formFactor"`
	ManufacturerId  *int    `pulumi:"manufacturerId"`
	MaxWeight       *int    `pulumi:"maxWeight"`
	Model           *string `pulumi:"model"`
	MountingDepthMm *int    `pulumi:"mountingDepthMm"`
	OuterDepth      *int    `pulumi:"outerDepth"`
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit    *string  `pulumi:"outerUnit"`
	OuterWidth   *int     `pulumi:"outerWidth"`
	Slug         *string  `pulumi:"slug"`
	StartingUnit *int     `pulumi:"startingUnit"`
	Tags         []string `pulumi:"tags"`
	UHeight      *int     `pulumi:"uHeight"`
	Weight       *float64 `pulumi:"weight"`
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit *string `pulumi:"weightUnit"`
	// Valid values are `10`, `19`, `21` and `23`.
	Width *int `pulumi:"width"`
}

type RackTypeState struct {
	Comments    pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	FormFactor      pulumi.StringPtrInput
	ManufacturerId  pulumi.IntPtrInput
	MaxWeight       pulumi.IntPtrInput
	Model           pulumi.StringPtrInput
	MountingDepthMm pulumi.IntPtrInput
	OuterDepth      pulumi.IntPtrInput
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit    pulumi.StringPtrInput
	OuterWidth   pulumi.IntPtrInput
	Slug         pulumi.StringPtrInput
	StartingUnit pulumi.IntPtrInput
	Tags         pulumi.StringArrayInput
	UHeight      pulumi.IntPtrInput
	Weight       pulumi.Float64PtrInput
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit pulumi.StringPtrInput
	// Valid values are `10`, `19`, `21` and `23`.
	Width pulumi.IntPtrInput
}

func (RackTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*rackTypeState)(nil)).Elem()
}

type rackTypeArgs struct {
	Comments    *string `pulumi:"comments"`
	Description *string `pulumi:"description"`
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	FormFactor      string `pulumi:"formFactor"`
	ManufacturerId  *int   `pulumi:"manufacturerId"`
	MaxWeight       *int   `pulumi:"maxWeight"`
	Model           string `pulumi:"model"`
	MountingDepthMm *int   `pulumi:"mountingDepthMm"`
	OuterDepth      *int   `pulumi:"outerDepth"`
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit    *string  `pulumi:"outerUnit"`
	OuterWidth   *int     `pulumi:"outerWidth"`
	Slug         *string  `pulumi:"slug"`
	StartingUnit int      `pulumi:"startingUnit"`
	Tags         []string `pulumi:"tags"`
	UHeight      int      `pulumi:"uHeight"`
	Weight       *float64 `pulumi:"weight"`
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit *string `pulumi:"weightUnit"`
	// Valid values are `10`, `19`, `21` and `23`.
	Width int `pulumi:"width"`
}

// The set of arguments for constructing a RackType resource.
type RackTypeArgs struct {
	Comments    pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	FormFactor      pulumi.StringInput
	ManufacturerId  pulumi.IntPtrInput
	MaxWeight       pulumi.IntPtrInput
	Model           pulumi.StringInput
	MountingDepthMm pulumi.IntPtrInput
	OuterDepth      pulumi.IntPtrInput
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit    pulumi.StringPtrInput
	OuterWidth   pulumi.IntPtrInput
	Slug         pulumi.StringPtrInput
	StartingUnit pulumi.IntInput
	Tags         pulumi.StringArrayInput
	UHeight      pulumi.IntInput
	Weight       pulumi.Float64PtrInput
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit pulumi.StringPtrInput
	// Valid values are `10`, `19`, `21` and `23`.
	Width pulumi.IntInput
}

func (RackTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rackTypeArgs)(nil)).Elem()
}

type RackTypeInput interface {
	pulumi.Input

	ToRackTypeOutput() RackTypeOutput
	ToRackTypeOutputWithContext(ctx context.Context) RackTypeOutput
}

func (*RackType) ElementType() reflect.Type {
	return reflect.TypeOf((**RackType)(nil)).Elem()
}

func (i *RackType) ToRackTypeOutput() RackTypeOutput {
	return i.ToRackTypeOutputWithContext(context.Background())
}

func (i *RackType) ToRackTypeOutputWithContext(ctx context.Context) RackTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RackTypeOutput)
}

// RackTypeArrayInput is an input type that accepts RackTypeArray and RackTypeArrayOutput values.
// You can construct a concrete instance of `RackTypeArrayInput` via:
//
//	RackTypeArray{ RackTypeArgs{...} }
type RackTypeArrayInput interface {
	pulumi.Input

	ToRackTypeArrayOutput() RackTypeArrayOutput
	ToRackTypeArrayOutputWithContext(context.Context) RackTypeArrayOutput
}

type RackTypeArray []RackTypeInput

func (RackTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RackType)(nil)).Elem()
}

func (i RackTypeArray) ToRackTypeArrayOutput() RackTypeArrayOutput {
	return i.ToRackTypeArrayOutputWithContext(context.Background())
}

func (i RackTypeArray) ToRackTypeArrayOutputWithContext(ctx context.Context) RackTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RackTypeArrayOutput)
}

// RackTypeMapInput is an input type that accepts RackTypeMap and RackTypeMapOutput values.
// You can construct a concrete instance of `RackTypeMapInput` via:
//
//	RackTypeMap{ "key": RackTypeArgs{...} }
type RackTypeMapInput interface {
	pulumi.Input

	ToRackTypeMapOutput() RackTypeMapOutput
	ToRackTypeMapOutputWithContext(context.Context) RackTypeMapOutput
}

type RackTypeMap map[string]RackTypeInput

func (RackTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RackType)(nil)).Elem()
}

func (i RackTypeMap) ToRackTypeMapOutput() RackTypeMapOutput {
	return i.ToRackTypeMapOutputWithContext(context.Background())
}

func (i RackTypeMap) ToRackTypeMapOutputWithContext(ctx context.Context) RackTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RackTypeMapOutput)
}

type RackTypeOutput struct{ *pulumi.OutputState }

func (RackTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RackType)(nil)).Elem()
}

func (o RackTypeOutput) ToRackTypeOutput() RackTypeOutput {
	return o
}

func (o RackTypeOutput) ToRackTypeOutputWithContext(ctx context.Context) RackTypeOutput {
	return o
}

func (o RackTypeOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o RackTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
func (o RackTypeOutput) FormFactor() pulumi.StringOutput {
	return o.ApplyT(func(v *RackType) pulumi.StringOutput { return v.FormFactor }).(pulumi.StringOutput)
}

func (o RackTypeOutput) ManufacturerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.IntPtrOutput { return v.ManufacturerId }).(pulumi.IntPtrOutput)
}

func (o RackTypeOutput) MaxWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.IntPtrOutput { return v.MaxWeight }).(pulumi.IntPtrOutput)
}

func (o RackTypeOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v *RackType) pulumi.StringOutput { return v.Model }).(pulumi.StringOutput)
}

func (o RackTypeOutput) MountingDepthMm() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.IntPtrOutput { return v.MountingDepthMm }).(pulumi.IntPtrOutput)
}

func (o RackTypeOutput) OuterDepth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.IntPtrOutput { return v.OuterDepth }).(pulumi.IntPtrOutput)
}

// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
func (o RackTypeOutput) OuterUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.StringPtrOutput { return v.OuterUnit }).(pulumi.StringPtrOutput)
}

func (o RackTypeOutput) OuterWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.IntPtrOutput { return v.OuterWidth }).(pulumi.IntPtrOutput)
}

func (o RackTypeOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *RackType) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

func (o RackTypeOutput) StartingUnit() pulumi.IntOutput {
	return o.ApplyT(func(v *RackType) pulumi.IntOutput { return v.StartingUnit }).(pulumi.IntOutput)
}

func (o RackTypeOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RackType) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o RackTypeOutput) UHeight() pulumi.IntOutput {
	return o.ApplyT(func(v *RackType) pulumi.IntOutput { return v.UHeight }).(pulumi.IntOutput)
}

func (o RackTypeOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.Float64PtrOutput { return v.Weight }).(pulumi.Float64PtrOutput)
}

// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
func (o RackTypeOutput) WeightUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RackType) pulumi.StringPtrOutput { return v.WeightUnit }).(pulumi.StringPtrOutput)
}

// Valid values are `10`, `19`, `21` and `23`.
func (o RackTypeOutput) Width() pulumi.IntOutput {
	return o.ApplyT(func(v *RackType) pulumi.IntOutput { return v.Width }).(pulumi.IntOutput)
}

type RackTypeArrayOutput struct{ *pulumi.OutputState }

func (RackTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RackType)(nil)).Elem()
}

func (o RackTypeArrayOutput) ToRackTypeArrayOutput() RackTypeArrayOutput {
	return o
}

func (o RackTypeArrayOutput) ToRackTypeArrayOutputWithContext(ctx context.Context) RackTypeArrayOutput {
	return o
}

func (o RackTypeArrayOutput) Index(i pulumi.IntInput) RackTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RackType {
		return vs[0].([]*RackType)[vs[1].(int)]
	}).(RackTypeOutput)
}

type RackTypeMapOutput struct{ *pulumi.OutputState }

func (RackTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RackType)(nil)).Elem()
}

func (o RackTypeMapOutput) ToRackTypeMapOutput() RackTypeMapOutput {
	return o
}

func (o RackTypeMapOutput) ToRackTypeMapOutputWithContext(ctx context.Context) RackTypeMapOutput {
	return o
}

func (o RackTypeMapOutput) MapIndex(k pulumi.StringInput) RackTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RackType {
		return vs[0].(map[string]*RackType)[vs[1].(string)]
	}).(RackTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RackTypeInput)(nil)).Elem(), &RackType{})
	pulumi.RegisterInputType(reflect.TypeOf((*RackTypeArrayInput)(nil)).Elem(), RackTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RackTypeMapInput)(nil)).Elem(), RackTypeMap{})
	pulumi.RegisterOutputType(RackTypeOutput{})
	pulumi.RegisterOutputType(RackTypeArrayOutput{})
	pulumi.RegisterOutputType(RackTypeMapOutput{})
}
