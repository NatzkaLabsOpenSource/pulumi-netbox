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

// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rack/):
//
// > The rack model represents a physical two- or four-post equipment rack in which devices can be installed. Each rack must be assigned to a site, and may optionally be assigned to a location within that site. Racks can also be organized by user-defined functional roles. The name and facility ID of each rack within a location must be unique.
//
// Rack height is measured in rack units (U); racks are commonly between 42U and 48U tall, but NetBox allows you to define racks of arbitrary height. A toggle is provided to indicate whether rack units are in ascending (from the ground up) or descending order.
//
// Each rack is assigned a name and (optionally) a separate facility ID. This is helpful when leasing space in a data center your organization does not own: The facility will often assign a seemingly arbitrary ID to a rack (for example, "M204.313") whereas internally you refer to is simply as "R113." A unique serial number and asset tag may also be associated with each rack.
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
//			testSite, err := dcim.NewSite(ctx, "testSite", &dcim.SiteArgs{
//				Status: pulumi.String("active"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dcim.NewRack(ctx, "testRack", &dcim.RackArgs{
//				SiteId:  testSite.ID(),
//				Status:  pulumi.String("reserved"),
//				Width:   pulumi.Int(19),
//				UHeight: pulumi.Int(48),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Rack struct {
	pulumi.CustomResourceState

	AssetTag     pulumi.StringPtrOutput `pulumi:"assetTag"`
	Comments     pulumi.StringPtrOutput `pulumi:"comments"`
	CustomFields pulumi.StringMapOutput `pulumi:"customFields"`
	// If rack units are descending. Defaults to `false`.
	DescUnits     pulumi.BoolPtrOutput   `pulumi:"descUnits"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	FacilityId    pulumi.StringPtrOutput `pulumi:"facilityId"`
	LocationId    pulumi.IntPtrOutput    `pulumi:"locationId"`
	MaxWeight     pulumi.IntPtrOutput    `pulumi:"maxWeight"`
	MountingDepth pulumi.IntPtrOutput    `pulumi:"mountingDepth"`
	Name          pulumi.StringOutput    `pulumi:"name"`
	OuterDepth    pulumi.IntPtrOutput    `pulumi:"outerDepth"`
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit  pulumi.StringPtrOutput `pulumi:"outerUnit"`
	OuterWidth pulumi.IntPtrOutput    `pulumi:"outerWidth"`
	RoleId     pulumi.IntPtrOutput    `pulumi:"roleId"`
	Serial     pulumi.StringPtrOutput `pulumi:"serial"`
	SiteId     pulumi.IntOutput       `pulumi:"siteId"`
	// Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
	Status   pulumi.StringOutput      `pulumi:"status"`
	Tags     pulumi.StringArrayOutput `pulumi:"tags"`
	TenantId pulumi.IntPtrOutput      `pulumi:"tenantId"`
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	Type    pulumi.StringPtrOutput  `pulumi:"type"`
	UHeight pulumi.IntOutput        `pulumi:"uHeight"`
	Weight  pulumi.Float64PtrOutput `pulumi:"weight"`
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit pulumi.StringPtrOutput `pulumi:"weightUnit"`
	// Valid values are `10`, `19`, `21` and `23`.
	Width pulumi.IntOutput `pulumi:"width"`
}

// NewRack registers a new resource with the given unique name, arguments, and options.
func NewRack(ctx *pulumi.Context,
	name string, args *RackArgs, opts ...pulumi.ResourceOption) (*Rack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.UHeight == nil {
		return nil, errors.New("invalid value for required argument 'UHeight'")
	}
	if args.Width == nil {
		return nil, errors.New("invalid value for required argument 'Width'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rack
	err := ctx.RegisterResource("netbox:dcim/rack:Rack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRack gets an existing Rack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RackState, opts ...pulumi.ResourceOption) (*Rack, error) {
	var resource Rack
	err := ctx.ReadResource("netbox:dcim/rack:Rack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rack resources.
type rackState struct {
	AssetTag     *string           `pulumi:"assetTag"`
	Comments     *string           `pulumi:"comments"`
	CustomFields map[string]string `pulumi:"customFields"`
	// If rack units are descending. Defaults to `false`.
	DescUnits     *bool   `pulumi:"descUnits"`
	Description   *string `pulumi:"description"`
	FacilityId    *string `pulumi:"facilityId"`
	LocationId    *int    `pulumi:"locationId"`
	MaxWeight     *int    `pulumi:"maxWeight"`
	MountingDepth *int    `pulumi:"mountingDepth"`
	Name          *string `pulumi:"name"`
	OuterDepth    *int    `pulumi:"outerDepth"`
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit  *string `pulumi:"outerUnit"`
	OuterWidth *int    `pulumi:"outerWidth"`
	RoleId     *int    `pulumi:"roleId"`
	Serial     *string `pulumi:"serial"`
	SiteId     *int    `pulumi:"siteId"`
	// Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
	Status   *string  `pulumi:"status"`
	Tags     []string `pulumi:"tags"`
	TenantId *int     `pulumi:"tenantId"`
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	Type    *string  `pulumi:"type"`
	UHeight *int     `pulumi:"uHeight"`
	Weight  *float64 `pulumi:"weight"`
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit *string `pulumi:"weightUnit"`
	// Valid values are `10`, `19`, `21` and `23`.
	Width *int `pulumi:"width"`
}

type RackState struct {
	AssetTag     pulumi.StringPtrInput
	Comments     pulumi.StringPtrInput
	CustomFields pulumi.StringMapInput
	// If rack units are descending. Defaults to `false`.
	DescUnits     pulumi.BoolPtrInput
	Description   pulumi.StringPtrInput
	FacilityId    pulumi.StringPtrInput
	LocationId    pulumi.IntPtrInput
	MaxWeight     pulumi.IntPtrInput
	MountingDepth pulumi.IntPtrInput
	Name          pulumi.StringPtrInput
	OuterDepth    pulumi.IntPtrInput
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit  pulumi.StringPtrInput
	OuterWidth pulumi.IntPtrInput
	RoleId     pulumi.IntPtrInput
	Serial     pulumi.StringPtrInput
	SiteId     pulumi.IntPtrInput
	// Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
	Status   pulumi.StringPtrInput
	Tags     pulumi.StringArrayInput
	TenantId pulumi.IntPtrInput
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	Type    pulumi.StringPtrInput
	UHeight pulumi.IntPtrInput
	Weight  pulumi.Float64PtrInput
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit pulumi.StringPtrInput
	// Valid values are `10`, `19`, `21` and `23`.
	Width pulumi.IntPtrInput
}

func (RackState) ElementType() reflect.Type {
	return reflect.TypeOf((*rackState)(nil)).Elem()
}

type rackArgs struct {
	AssetTag     *string           `pulumi:"assetTag"`
	Comments     *string           `pulumi:"comments"`
	CustomFields map[string]string `pulumi:"customFields"`
	// If rack units are descending. Defaults to `false`.
	DescUnits     *bool   `pulumi:"descUnits"`
	Description   *string `pulumi:"description"`
	FacilityId    *string `pulumi:"facilityId"`
	LocationId    *int    `pulumi:"locationId"`
	MaxWeight     *int    `pulumi:"maxWeight"`
	MountingDepth *int    `pulumi:"mountingDepth"`
	Name          *string `pulumi:"name"`
	OuterDepth    *int    `pulumi:"outerDepth"`
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit  *string `pulumi:"outerUnit"`
	OuterWidth *int    `pulumi:"outerWidth"`
	RoleId     *int    `pulumi:"roleId"`
	Serial     *string `pulumi:"serial"`
	SiteId     int     `pulumi:"siteId"`
	// Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
	Status   string   `pulumi:"status"`
	Tags     []string `pulumi:"tags"`
	TenantId *int     `pulumi:"tenantId"`
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	Type    *string  `pulumi:"type"`
	UHeight int      `pulumi:"uHeight"`
	Weight  *float64 `pulumi:"weight"`
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit *string `pulumi:"weightUnit"`
	// Valid values are `10`, `19`, `21` and `23`.
	Width int `pulumi:"width"`
}

// The set of arguments for constructing a Rack resource.
type RackArgs struct {
	AssetTag     pulumi.StringPtrInput
	Comments     pulumi.StringPtrInput
	CustomFields pulumi.StringMapInput
	// If rack units are descending. Defaults to `false`.
	DescUnits     pulumi.BoolPtrInput
	Description   pulumi.StringPtrInput
	FacilityId    pulumi.StringPtrInput
	LocationId    pulumi.IntPtrInput
	MaxWeight     pulumi.IntPtrInput
	MountingDepth pulumi.IntPtrInput
	Name          pulumi.StringPtrInput
	OuterDepth    pulumi.IntPtrInput
	// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
	OuterUnit  pulumi.StringPtrInput
	OuterWidth pulumi.IntPtrInput
	RoleId     pulumi.IntPtrInput
	Serial     pulumi.StringPtrInput
	SiteId     pulumi.IntInput
	// Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
	Status   pulumi.StringInput
	Tags     pulumi.StringArrayInput
	TenantId pulumi.IntPtrInput
	// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
	Type    pulumi.StringPtrInput
	UHeight pulumi.IntInput
	Weight  pulumi.Float64PtrInput
	// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
	WeightUnit pulumi.StringPtrInput
	// Valid values are `10`, `19`, `21` and `23`.
	Width pulumi.IntInput
}

func (RackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rackArgs)(nil)).Elem()
}

type RackInput interface {
	pulumi.Input

	ToRackOutput() RackOutput
	ToRackOutputWithContext(ctx context.Context) RackOutput
}

func (*Rack) ElementType() reflect.Type {
	return reflect.TypeOf((**Rack)(nil)).Elem()
}

func (i *Rack) ToRackOutput() RackOutput {
	return i.ToRackOutputWithContext(context.Background())
}

func (i *Rack) ToRackOutputWithContext(ctx context.Context) RackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RackOutput)
}

// RackArrayInput is an input type that accepts RackArray and RackArrayOutput values.
// You can construct a concrete instance of `RackArrayInput` via:
//
//	RackArray{ RackArgs{...} }
type RackArrayInput interface {
	pulumi.Input

	ToRackArrayOutput() RackArrayOutput
	ToRackArrayOutputWithContext(context.Context) RackArrayOutput
}

type RackArray []RackInput

func (RackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rack)(nil)).Elem()
}

func (i RackArray) ToRackArrayOutput() RackArrayOutput {
	return i.ToRackArrayOutputWithContext(context.Background())
}

func (i RackArray) ToRackArrayOutputWithContext(ctx context.Context) RackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RackArrayOutput)
}

// RackMapInput is an input type that accepts RackMap and RackMapOutput values.
// You can construct a concrete instance of `RackMapInput` via:
//
//	RackMap{ "key": RackArgs{...} }
type RackMapInput interface {
	pulumi.Input

	ToRackMapOutput() RackMapOutput
	ToRackMapOutputWithContext(context.Context) RackMapOutput
}

type RackMap map[string]RackInput

func (RackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rack)(nil)).Elem()
}

func (i RackMap) ToRackMapOutput() RackMapOutput {
	return i.ToRackMapOutputWithContext(context.Background())
}

func (i RackMap) ToRackMapOutputWithContext(ctx context.Context) RackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RackMapOutput)
}

type RackOutput struct{ *pulumi.OutputState }

func (RackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rack)(nil)).Elem()
}

func (o RackOutput) ToRackOutput() RackOutput {
	return o
}

func (o RackOutput) ToRackOutputWithContext(ctx context.Context) RackOutput {
	return o
}

func (o RackOutput) AssetTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringPtrOutput { return v.AssetTag }).(pulumi.StringPtrOutput)
}

func (o RackOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o RackOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

// If rack units are descending. Defaults to `false`.
func (o RackOutput) DescUnits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.BoolPtrOutput { return v.DescUnits }).(pulumi.BoolPtrOutput)
}

func (o RackOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RackOutput) FacilityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringPtrOutput { return v.FacilityId }).(pulumi.StringPtrOutput)
}

func (o RackOutput) LocationId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntPtrOutput { return v.LocationId }).(pulumi.IntPtrOutput)
}

func (o RackOutput) MaxWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntPtrOutput { return v.MaxWeight }).(pulumi.IntPtrOutput)
}

func (o RackOutput) MountingDepth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntPtrOutput { return v.MountingDepth }).(pulumi.IntPtrOutput)
}

func (o RackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RackOutput) OuterDepth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntPtrOutput { return v.OuterDepth }).(pulumi.IntPtrOutput)
}

// Valid values are `mm` and `in`. Required when `outerWidth` and `outerDepth` is set.
func (o RackOutput) OuterUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringPtrOutput { return v.OuterUnit }).(pulumi.StringPtrOutput)
}

func (o RackOutput) OuterWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntPtrOutput { return v.OuterWidth }).(pulumi.IntPtrOutput)
}

func (o RackOutput) RoleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntPtrOutput { return v.RoleId }).(pulumi.IntPtrOutput)
}

func (o RackOutput) Serial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringPtrOutput { return v.Serial }).(pulumi.StringPtrOutput)
}

func (o RackOutput) SiteId() pulumi.IntOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntOutput { return v.SiteId }).(pulumi.IntOutput)
}

// Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.
func (o RackOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o RackOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o RackOutput) TenantId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntPtrOutput { return v.TenantId }).(pulumi.IntPtrOutput)
}

// Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
func (o RackOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o RackOutput) UHeight() pulumi.IntOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntOutput { return v.UHeight }).(pulumi.IntOutput)
}

func (o RackOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.Float64PtrOutput { return v.Weight }).(pulumi.Float64PtrOutput)
}

// Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `maxWeight` is set.
func (o RackOutput) WeightUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rack) pulumi.StringPtrOutput { return v.WeightUnit }).(pulumi.StringPtrOutput)
}

// Valid values are `10`, `19`, `21` and `23`.
func (o RackOutput) Width() pulumi.IntOutput {
	return o.ApplyT(func(v *Rack) pulumi.IntOutput { return v.Width }).(pulumi.IntOutput)
}

type RackArrayOutput struct{ *pulumi.OutputState }

func (RackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rack)(nil)).Elem()
}

func (o RackArrayOutput) ToRackArrayOutput() RackArrayOutput {
	return o
}

func (o RackArrayOutput) ToRackArrayOutputWithContext(ctx context.Context) RackArrayOutput {
	return o
}

func (o RackArrayOutput) Index(i pulumi.IntInput) RackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rack {
		return vs[0].([]*Rack)[vs[1].(int)]
	}).(RackOutput)
}

type RackMapOutput struct{ *pulumi.OutputState }

func (RackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rack)(nil)).Elem()
}

func (o RackMapOutput) ToRackMapOutput() RackMapOutput {
	return o
}

func (o RackMapOutput) ToRackMapOutputWithContext(ctx context.Context) RackMapOutput {
	return o
}

func (o RackMapOutput) MapIndex(k pulumi.StringInput) RackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rack {
		return vs[0].(map[string]*Rack)[vs[1].(string)]
	}).(RackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RackInput)(nil)).Elem(), &Rack{})
	pulumi.RegisterInputType(reflect.TypeOf((*RackArrayInput)(nil)).Elem(), RackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RackMapInput)(nil)).Elem(), RackMap{})
	pulumi.RegisterOutputType(RackOutput{})
	pulumi.RegisterOutputType(RackArrayOutput{})
	pulumi.RegisterOutputType(RackMapOutput{})
}
