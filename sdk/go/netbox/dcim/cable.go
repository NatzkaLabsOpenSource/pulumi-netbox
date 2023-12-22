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

// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/cable/):
//
// > All connections between device components in NetBox are represented using cables. A cable represents a direct physical connection between two sets of endpoints (A and B), such as a console port and a patch panel port, or between two network interfaces.
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
//			_, err := dcim.NewCable(ctx, "test", &dcim.CableArgs{
//				ATerminations: dcim.CableATerminationArray{
//					&dcim.CableATerminationArgs{
//						ObjectType: pulumi.String("dcim.consoleserverport"),
//						ObjectId:   pulumi.Any(netbox_device_console_server_port.Kvm1.Id),
//					},
//					&dcim.CableATerminationArgs{
//						ObjectType: pulumi.String("dcim.consoleserverport"),
//						ObjectId:   pulumi.Any(netbox_device_console_server_port.Kvm2.Id),
//					},
//				},
//				BTerminations: dcim.CableBTerminationArray{
//					&dcim.CableBTerminationArgs{
//						ObjectType: pulumi.String("dcim.consoleport"),
//						ObjectId:   pulumi.Any(netbox_device_console_port.Server1.Id),
//					},
//					&dcim.CableBTerminationArgs{
//						ObjectType: pulumi.String("dcim.consoleport"),
//						ObjectId:   pulumi.Any(netbox_device_console_port.Server2.Id),
//					},
//				},
//				Status:     pulumi.String("connected"),
//				Label:      pulumi.String("KVM cable"),
//				Type:       pulumi.String("cat8"),
//				ColorHex:   pulumi.String("123456"),
//				Length:     pulumi.Float64(10),
//				LengthUnit: pulumi.String("m"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Cable struct {
	pulumi.CustomResourceState

	ATerminations CableATerminationArrayOutput `pulumi:"aTerminations"`
	BTerminations CableBTerminationArrayOutput `pulumi:"bTerminations"`
	ColorHex      pulumi.StringPtrOutput       `pulumi:"colorHex"`
	Comments      pulumi.StringPtrOutput       `pulumi:"comments"`
	CustomFields  pulumi.StringMapOutput       `pulumi:"customFields"`
	Description   pulumi.StringPtrOutput       `pulumi:"description"`
	Label         pulumi.StringPtrOutput       `pulumi:"label"`
	Length        pulumi.Float64PtrOutput      `pulumi:"length"`
	// One of [km, m, cm, mi, ft, in]. Required when `length` is set.
	LengthUnit pulumi.StringPtrOutput `pulumi:"lengthUnit"`
	// One of [connected, planned, decommissioning].
	Status   pulumi.StringOutput      `pulumi:"status"`
	Tags     pulumi.StringArrayOutput `pulumi:"tags"`
	TenantId pulumi.IntPtrOutput      `pulumi:"tenantId"`
	// One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power].
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewCable registers a new resource with the given unique name, arguments, and options.
func NewCable(ctx *pulumi.Context,
	name string, args *CableArgs, opts ...pulumi.ResourceOption) (*Cable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ATerminations == nil {
		return nil, errors.New("invalid value for required argument 'ATerminations'")
	}
	if args.BTerminations == nil {
		return nil, errors.New("invalid value for required argument 'BTerminations'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cable
	err := ctx.RegisterResource("netbox:dcim/cable:Cable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCable gets an existing Cable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CableState, opts ...pulumi.ResourceOption) (*Cable, error) {
	var resource Cable
	err := ctx.ReadResource("netbox:dcim/cable:Cable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cable resources.
type cableState struct {
	ATerminations []CableATermination `pulumi:"aTerminations"`
	BTerminations []CableBTermination `pulumi:"bTerminations"`
	ColorHex      *string             `pulumi:"colorHex"`
	Comments      *string             `pulumi:"comments"`
	CustomFields  map[string]string   `pulumi:"customFields"`
	Description   *string             `pulumi:"description"`
	Label         *string             `pulumi:"label"`
	Length        *float64            `pulumi:"length"`
	// One of [km, m, cm, mi, ft, in]. Required when `length` is set.
	LengthUnit *string `pulumi:"lengthUnit"`
	// One of [connected, planned, decommissioning].
	Status   *string  `pulumi:"status"`
	Tags     []string `pulumi:"tags"`
	TenantId *int     `pulumi:"tenantId"`
	// One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power].
	Type *string `pulumi:"type"`
}

type CableState struct {
	ATerminations CableATerminationArrayInput
	BTerminations CableBTerminationArrayInput
	ColorHex      pulumi.StringPtrInput
	Comments      pulumi.StringPtrInput
	CustomFields  pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	Label         pulumi.StringPtrInput
	Length        pulumi.Float64PtrInput
	// One of [km, m, cm, mi, ft, in]. Required when `length` is set.
	LengthUnit pulumi.StringPtrInput
	// One of [connected, planned, decommissioning].
	Status   pulumi.StringPtrInput
	Tags     pulumi.StringArrayInput
	TenantId pulumi.IntPtrInput
	// One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power].
	Type pulumi.StringPtrInput
}

func (CableState) ElementType() reflect.Type {
	return reflect.TypeOf((*cableState)(nil)).Elem()
}

type cableArgs struct {
	ATerminations []CableATermination `pulumi:"aTerminations"`
	BTerminations []CableBTermination `pulumi:"bTerminations"`
	ColorHex      *string             `pulumi:"colorHex"`
	Comments      *string             `pulumi:"comments"`
	CustomFields  map[string]string   `pulumi:"customFields"`
	Description   *string             `pulumi:"description"`
	Label         *string             `pulumi:"label"`
	Length        *float64            `pulumi:"length"`
	// One of [km, m, cm, mi, ft, in]. Required when `length` is set.
	LengthUnit *string `pulumi:"lengthUnit"`
	// One of [connected, planned, decommissioning].
	Status   string   `pulumi:"status"`
	Tags     []string `pulumi:"tags"`
	TenantId *int     `pulumi:"tenantId"`
	// One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power].
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Cable resource.
type CableArgs struct {
	ATerminations CableATerminationArrayInput
	BTerminations CableBTerminationArrayInput
	ColorHex      pulumi.StringPtrInput
	Comments      pulumi.StringPtrInput
	CustomFields  pulumi.StringMapInput
	Description   pulumi.StringPtrInput
	Label         pulumi.StringPtrInput
	Length        pulumi.Float64PtrInput
	// One of [km, m, cm, mi, ft, in]. Required when `length` is set.
	LengthUnit pulumi.StringPtrInput
	// One of [connected, planned, decommissioning].
	Status   pulumi.StringInput
	Tags     pulumi.StringArrayInput
	TenantId pulumi.IntPtrInput
	// One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power].
	Type pulumi.StringPtrInput
}

func (CableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cableArgs)(nil)).Elem()
}

type CableInput interface {
	pulumi.Input

	ToCableOutput() CableOutput
	ToCableOutputWithContext(ctx context.Context) CableOutput
}

func (*Cable) ElementType() reflect.Type {
	return reflect.TypeOf((**Cable)(nil)).Elem()
}

func (i *Cable) ToCableOutput() CableOutput {
	return i.ToCableOutputWithContext(context.Background())
}

func (i *Cable) ToCableOutputWithContext(ctx context.Context) CableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CableOutput)
}

// CableArrayInput is an input type that accepts CableArray and CableArrayOutput values.
// You can construct a concrete instance of `CableArrayInput` via:
//
//	CableArray{ CableArgs{...} }
type CableArrayInput interface {
	pulumi.Input

	ToCableArrayOutput() CableArrayOutput
	ToCableArrayOutputWithContext(context.Context) CableArrayOutput
}

type CableArray []CableInput

func (CableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cable)(nil)).Elem()
}

func (i CableArray) ToCableArrayOutput() CableArrayOutput {
	return i.ToCableArrayOutputWithContext(context.Background())
}

func (i CableArray) ToCableArrayOutputWithContext(ctx context.Context) CableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CableArrayOutput)
}

// CableMapInput is an input type that accepts CableMap and CableMapOutput values.
// You can construct a concrete instance of `CableMapInput` via:
//
//	CableMap{ "key": CableArgs{...} }
type CableMapInput interface {
	pulumi.Input

	ToCableMapOutput() CableMapOutput
	ToCableMapOutputWithContext(context.Context) CableMapOutput
}

type CableMap map[string]CableInput

func (CableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cable)(nil)).Elem()
}

func (i CableMap) ToCableMapOutput() CableMapOutput {
	return i.ToCableMapOutputWithContext(context.Background())
}

func (i CableMap) ToCableMapOutputWithContext(ctx context.Context) CableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CableMapOutput)
}

type CableOutput struct{ *pulumi.OutputState }

func (CableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cable)(nil)).Elem()
}

func (o CableOutput) ToCableOutput() CableOutput {
	return o
}

func (o CableOutput) ToCableOutputWithContext(ctx context.Context) CableOutput {
	return o
}

func (o CableOutput) ATerminations() CableATerminationArrayOutput {
	return o.ApplyT(func(v *Cable) CableATerminationArrayOutput { return v.ATerminations }).(CableATerminationArrayOutput)
}

func (o CableOutput) BTerminations() CableBTerminationArrayOutput {
	return o.ApplyT(func(v *Cable) CableBTerminationArrayOutput { return v.BTerminations }).(CableBTerminationArrayOutput)
}

func (o CableOutput) ColorHex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cable) pulumi.StringPtrOutput { return v.ColorHex }).(pulumi.StringPtrOutput)
}

func (o CableOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cable) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o CableOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cable) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

func (o CableOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cable) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CableOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cable) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

func (o CableOutput) Length() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Cable) pulumi.Float64PtrOutput { return v.Length }).(pulumi.Float64PtrOutput)
}

// One of [km, m, cm, mi, ft, in]. Required when `length` is set.
func (o CableOutput) LengthUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cable) pulumi.StringPtrOutput { return v.LengthUnit }).(pulumi.StringPtrOutput)
}

// One of [connected, planned, decommissioning].
func (o CableOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Cable) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o CableOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cable) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o CableOutput) TenantId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cable) pulumi.IntPtrOutput { return v.TenantId }).(pulumi.IntPtrOutput)
}

// One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power].
func (o CableOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cable) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type CableArrayOutput struct{ *pulumi.OutputState }

func (CableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cable)(nil)).Elem()
}

func (o CableArrayOutput) ToCableArrayOutput() CableArrayOutput {
	return o
}

func (o CableArrayOutput) ToCableArrayOutputWithContext(ctx context.Context) CableArrayOutput {
	return o
}

func (o CableArrayOutput) Index(i pulumi.IntInput) CableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cable {
		return vs[0].([]*Cable)[vs[1].(int)]
	}).(CableOutput)
}

type CableMapOutput struct{ *pulumi.OutputState }

func (CableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cable)(nil)).Elem()
}

func (o CableMapOutput) ToCableMapOutput() CableMapOutput {
	return o
}

func (o CableMapOutput) ToCableMapOutputWithContext(ctx context.Context) CableMapOutput {
	return o
}

func (o CableMapOutput) MapIndex(k pulumi.StringInput) CableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cable {
		return vs[0].(map[string]*Cable)[vs[1].(string)]
	}).(CableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CableInput)(nil)).Elem(), &Cable{})
	pulumi.RegisterInputType(reflect.TypeOf((*CableArrayInput)(nil)).Elem(), CableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CableMapInput)(nil)).Elem(), CableMap{})
	pulumi.RegisterOutputType(CableOutput{})
	pulumi.RegisterOutputType(CableArrayOutput{})
	pulumi.RegisterOutputType(CableMapOutput{})
}
