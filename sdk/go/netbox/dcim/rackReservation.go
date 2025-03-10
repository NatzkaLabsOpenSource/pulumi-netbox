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

// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rackreservation/):
//
// > Users can reserve specific units within a rack for future use. An arbitrary set of units within a rack can be associated with a single reservation, but reservations cannot span multiple racks. A description is required for each reservation, reservations may optionally be associated with a specific tenant.
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
//			test, err := dcim.NewSite(ctx, "test", &dcim.SiteArgs{
//				Name:   pulumi.String("test"),
//				Status: pulumi.String("active"),
//			})
//			if err != nil {
//				return err
//			}
//			testRack, err := dcim.NewRack(ctx, "test", &dcim.RackArgs{
//				Name:    pulumi.String("test"),
//				SiteId:  test.ID(),
//				Status:  pulumi.String("active"),
//				Width:   pulumi.Int(10),
//				UHeight: pulumi.Int(40),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dcim.NewRackReservation(ctx, "test", &dcim.RackReservationArgs{
//				RackId: testRack.ID(),
//				Units: pulumi.IntArray{
//					pulumi.Int(1),
//					pulumi.Int(2),
//					pulumi.Int(3),
//					pulumi.Int(4),
//					pulumi.Int(5),
//				},
//				UserId:      pulumi.Int(1),
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
type RackReservation struct {
	pulumi.CustomResourceState

	Comments    pulumi.StringPtrOutput   `pulumi:"comments"`
	Description pulumi.StringOutput      `pulumi:"description"`
	RackId      pulumi.IntOutput         `pulumi:"rackId"`
	Tags        pulumi.StringArrayOutput `pulumi:"tags"`
	TenantId    pulumi.IntPtrOutput      `pulumi:"tenantId"`
	Units       pulumi.IntArrayOutput    `pulumi:"units"`
	UserId      pulumi.IntOutput         `pulumi:"userId"`
}

// NewRackReservation registers a new resource with the given unique name, arguments, and options.
func NewRackReservation(ctx *pulumi.Context,
	name string, args *RackReservationArgs, opts ...pulumi.ResourceOption) (*RackReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.RackId == nil {
		return nil, errors.New("invalid value for required argument 'RackId'")
	}
	if args.Units == nil {
		return nil, errors.New("invalid value for required argument 'Units'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RackReservation
	err := ctx.RegisterResource("netbox:dcim/rackReservation:RackReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRackReservation gets an existing RackReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRackReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RackReservationState, opts ...pulumi.ResourceOption) (*RackReservation, error) {
	var resource RackReservation
	err := ctx.ReadResource("netbox:dcim/rackReservation:RackReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RackReservation resources.
type rackReservationState struct {
	Comments    *string  `pulumi:"comments"`
	Description *string  `pulumi:"description"`
	RackId      *int     `pulumi:"rackId"`
	Tags        []string `pulumi:"tags"`
	TenantId    *int     `pulumi:"tenantId"`
	Units       []int    `pulumi:"units"`
	UserId      *int     `pulumi:"userId"`
}

type RackReservationState struct {
	Comments    pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	RackId      pulumi.IntPtrInput
	Tags        pulumi.StringArrayInput
	TenantId    pulumi.IntPtrInput
	Units       pulumi.IntArrayInput
	UserId      pulumi.IntPtrInput
}

func (RackReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*rackReservationState)(nil)).Elem()
}

type rackReservationArgs struct {
	Comments    *string  `pulumi:"comments"`
	Description string   `pulumi:"description"`
	RackId      int      `pulumi:"rackId"`
	Tags        []string `pulumi:"tags"`
	TenantId    *int     `pulumi:"tenantId"`
	Units       []int    `pulumi:"units"`
	UserId      int      `pulumi:"userId"`
}

// The set of arguments for constructing a RackReservation resource.
type RackReservationArgs struct {
	Comments    pulumi.StringPtrInput
	Description pulumi.StringInput
	RackId      pulumi.IntInput
	Tags        pulumi.StringArrayInput
	TenantId    pulumi.IntPtrInput
	Units       pulumi.IntArrayInput
	UserId      pulumi.IntInput
}

func (RackReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rackReservationArgs)(nil)).Elem()
}

type RackReservationInput interface {
	pulumi.Input

	ToRackReservationOutput() RackReservationOutput
	ToRackReservationOutputWithContext(ctx context.Context) RackReservationOutput
}

func (*RackReservation) ElementType() reflect.Type {
	return reflect.TypeOf((**RackReservation)(nil)).Elem()
}

func (i *RackReservation) ToRackReservationOutput() RackReservationOutput {
	return i.ToRackReservationOutputWithContext(context.Background())
}

func (i *RackReservation) ToRackReservationOutputWithContext(ctx context.Context) RackReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RackReservationOutput)
}

// RackReservationArrayInput is an input type that accepts RackReservationArray and RackReservationArrayOutput values.
// You can construct a concrete instance of `RackReservationArrayInput` via:
//
//	RackReservationArray{ RackReservationArgs{...} }
type RackReservationArrayInput interface {
	pulumi.Input

	ToRackReservationArrayOutput() RackReservationArrayOutput
	ToRackReservationArrayOutputWithContext(context.Context) RackReservationArrayOutput
}

type RackReservationArray []RackReservationInput

func (RackReservationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RackReservation)(nil)).Elem()
}

func (i RackReservationArray) ToRackReservationArrayOutput() RackReservationArrayOutput {
	return i.ToRackReservationArrayOutputWithContext(context.Background())
}

func (i RackReservationArray) ToRackReservationArrayOutputWithContext(ctx context.Context) RackReservationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RackReservationArrayOutput)
}

// RackReservationMapInput is an input type that accepts RackReservationMap and RackReservationMapOutput values.
// You can construct a concrete instance of `RackReservationMapInput` via:
//
//	RackReservationMap{ "key": RackReservationArgs{...} }
type RackReservationMapInput interface {
	pulumi.Input

	ToRackReservationMapOutput() RackReservationMapOutput
	ToRackReservationMapOutputWithContext(context.Context) RackReservationMapOutput
}

type RackReservationMap map[string]RackReservationInput

func (RackReservationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RackReservation)(nil)).Elem()
}

func (i RackReservationMap) ToRackReservationMapOutput() RackReservationMapOutput {
	return i.ToRackReservationMapOutputWithContext(context.Background())
}

func (i RackReservationMap) ToRackReservationMapOutputWithContext(ctx context.Context) RackReservationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RackReservationMapOutput)
}

type RackReservationOutput struct{ *pulumi.OutputState }

func (RackReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RackReservation)(nil)).Elem()
}

func (o RackReservationOutput) ToRackReservationOutput() RackReservationOutput {
	return o
}

func (o RackReservationOutput) ToRackReservationOutputWithContext(ctx context.Context) RackReservationOutput {
	return o
}

func (o RackReservationOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RackReservation) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o RackReservationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RackReservation) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o RackReservationOutput) RackId() pulumi.IntOutput {
	return o.ApplyT(func(v *RackReservation) pulumi.IntOutput { return v.RackId }).(pulumi.IntOutput)
}

func (o RackReservationOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RackReservation) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o RackReservationOutput) TenantId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RackReservation) pulumi.IntPtrOutput { return v.TenantId }).(pulumi.IntPtrOutput)
}

func (o RackReservationOutput) Units() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RackReservation) pulumi.IntArrayOutput { return v.Units }).(pulumi.IntArrayOutput)
}

func (o RackReservationOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *RackReservation) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type RackReservationArrayOutput struct{ *pulumi.OutputState }

func (RackReservationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RackReservation)(nil)).Elem()
}

func (o RackReservationArrayOutput) ToRackReservationArrayOutput() RackReservationArrayOutput {
	return o
}

func (o RackReservationArrayOutput) ToRackReservationArrayOutputWithContext(ctx context.Context) RackReservationArrayOutput {
	return o
}

func (o RackReservationArrayOutput) Index(i pulumi.IntInput) RackReservationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RackReservation {
		return vs[0].([]*RackReservation)[vs[1].(int)]
	}).(RackReservationOutput)
}

type RackReservationMapOutput struct{ *pulumi.OutputState }

func (RackReservationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RackReservation)(nil)).Elem()
}

func (o RackReservationMapOutput) ToRackReservationMapOutput() RackReservationMapOutput {
	return o
}

func (o RackReservationMapOutput) ToRackReservationMapOutputWithContext(ctx context.Context) RackReservationMapOutput {
	return o
}

func (o RackReservationMapOutput) MapIndex(k pulumi.StringInput) RackReservationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RackReservation {
		return vs[0].(map[string]*RackReservation)[vs[1].(string)]
	}).(RackReservationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RackReservationInput)(nil)).Elem(), &RackReservation{})
	pulumi.RegisterInputType(reflect.TypeOf((*RackReservationArrayInput)(nil)).Elem(), RackReservationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RackReservationMapInput)(nil)).Elem(), RackReservationMap{})
	pulumi.RegisterOutputType(RackReservationOutput{})
	pulumi.RegisterOutputType(RackReservationArrayOutput{})
	pulumi.RegisterOutputType(RackReservationMapOutput{})
}
