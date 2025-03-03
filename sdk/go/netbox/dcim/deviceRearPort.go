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

// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rearport/):
//
// > Like front ports, rear ports are pass-through ports which represent the continuation of a path from one cable to the next. Each rear port is defined with its physical type and a number of positions: Rear ports with more than one position can be mapped to multiple front ports. This can be useful for modeling instances where multiple paths share a common cable (for example, six discrete two-strand fiber connections sharing a 12-strand MPO cable).
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
//			// Note that some terraform code is not included in the example for brevity
//			test, err := dcim.NewDevice(ctx, "test", &dcim.DeviceArgs{
//				Name:         pulumi.String("%[1]s"),
//				DeviceTypeId: pulumi.Any(testNetboxDeviceType.Id),
//				RoleId:       pulumi.Any(testNetboxDeviceRole.Id),
//				SiteId:       pulumi.Any(testNetboxSite.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dcim.NewDeviceRearPort(ctx, "test", &dcim.DeviceRearPortArgs{
//				DeviceId:      test.ID(),
//				Name:          pulumi.String("rear port 1"),
//				Type:          pulumi.String("8p8c"),
//				Positions:     pulumi.Int(2),
//				MarkConnected: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DeviceRearPort struct {
	pulumi.CustomResourceState

	ColorHex     pulumi.StringPtrOutput `pulumi:"colorHex"`
	CustomFields pulumi.StringMapOutput `pulumi:"customFields"`
	Description  pulumi.StringPtrOutput `pulumi:"description"`
	DeviceId     pulumi.IntOutput       `pulumi:"deviceId"`
	Label        pulumi.StringPtrOutput `pulumi:"label"`
	// Defaults to `false`.
	MarkConnected pulumi.BoolPtrOutput     `pulumi:"markConnected"`
	ModuleId      pulumi.IntPtrOutput      `pulumi:"moduleId"`
	Name          pulumi.StringOutput      `pulumi:"name"`
	Positions     pulumi.IntOutput         `pulumi:"positions"`
	Tags          pulumi.StringArrayOutput `pulumi:"tags"`
	// One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other].
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDeviceRearPort registers a new resource with the given unique name, arguments, and options.
func NewDeviceRearPort(ctx *pulumi.Context,
	name string, args *DeviceRearPortArgs, opts ...pulumi.ResourceOption) (*DeviceRearPort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	if args.Positions == nil {
		return nil, errors.New("invalid value for required argument 'Positions'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DeviceRearPort
	err := ctx.RegisterResource("netbox:dcim/deviceRearPort:DeviceRearPort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceRearPort gets an existing DeviceRearPort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceRearPort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceRearPortState, opts ...pulumi.ResourceOption) (*DeviceRearPort, error) {
	var resource DeviceRearPort
	err := ctx.ReadResource("netbox:dcim/deviceRearPort:DeviceRearPort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceRearPort resources.
type deviceRearPortState struct {
	ColorHex     *string           `pulumi:"colorHex"`
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	DeviceId     *int              `pulumi:"deviceId"`
	Label        *string           `pulumi:"label"`
	// Defaults to `false`.
	MarkConnected *bool    `pulumi:"markConnected"`
	ModuleId      *int     `pulumi:"moduleId"`
	Name          *string  `pulumi:"name"`
	Positions     *int     `pulumi:"positions"`
	Tags          []string `pulumi:"tags"`
	// One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other].
	Type *string `pulumi:"type"`
}

type DeviceRearPortState struct {
	ColorHex     pulumi.StringPtrInput
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	DeviceId     pulumi.IntPtrInput
	Label        pulumi.StringPtrInput
	// Defaults to `false`.
	MarkConnected pulumi.BoolPtrInput
	ModuleId      pulumi.IntPtrInput
	Name          pulumi.StringPtrInput
	Positions     pulumi.IntPtrInput
	Tags          pulumi.StringArrayInput
	// One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other].
	Type pulumi.StringPtrInput
}

func (DeviceRearPortState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceRearPortState)(nil)).Elem()
}

type deviceRearPortArgs struct {
	ColorHex     *string           `pulumi:"colorHex"`
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	DeviceId     int               `pulumi:"deviceId"`
	Label        *string           `pulumi:"label"`
	// Defaults to `false`.
	MarkConnected *bool    `pulumi:"markConnected"`
	ModuleId      *int     `pulumi:"moduleId"`
	Name          *string  `pulumi:"name"`
	Positions     int      `pulumi:"positions"`
	Tags          []string `pulumi:"tags"`
	// One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other].
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DeviceRearPort resource.
type DeviceRearPortArgs struct {
	ColorHex     pulumi.StringPtrInput
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	DeviceId     pulumi.IntInput
	Label        pulumi.StringPtrInput
	// Defaults to `false`.
	MarkConnected pulumi.BoolPtrInput
	ModuleId      pulumi.IntPtrInput
	Name          pulumi.StringPtrInput
	Positions     pulumi.IntInput
	Tags          pulumi.StringArrayInput
	// One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other].
	Type pulumi.StringInput
}

func (DeviceRearPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceRearPortArgs)(nil)).Elem()
}

type DeviceRearPortInput interface {
	pulumi.Input

	ToDeviceRearPortOutput() DeviceRearPortOutput
	ToDeviceRearPortOutputWithContext(ctx context.Context) DeviceRearPortOutput
}

func (*DeviceRearPort) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceRearPort)(nil)).Elem()
}

func (i *DeviceRearPort) ToDeviceRearPortOutput() DeviceRearPortOutput {
	return i.ToDeviceRearPortOutputWithContext(context.Background())
}

func (i *DeviceRearPort) ToDeviceRearPortOutputWithContext(ctx context.Context) DeviceRearPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceRearPortOutput)
}

// DeviceRearPortArrayInput is an input type that accepts DeviceRearPortArray and DeviceRearPortArrayOutput values.
// You can construct a concrete instance of `DeviceRearPortArrayInput` via:
//
//	DeviceRearPortArray{ DeviceRearPortArgs{...} }
type DeviceRearPortArrayInput interface {
	pulumi.Input

	ToDeviceRearPortArrayOutput() DeviceRearPortArrayOutput
	ToDeviceRearPortArrayOutputWithContext(context.Context) DeviceRearPortArrayOutput
}

type DeviceRearPortArray []DeviceRearPortInput

func (DeviceRearPortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceRearPort)(nil)).Elem()
}

func (i DeviceRearPortArray) ToDeviceRearPortArrayOutput() DeviceRearPortArrayOutput {
	return i.ToDeviceRearPortArrayOutputWithContext(context.Background())
}

func (i DeviceRearPortArray) ToDeviceRearPortArrayOutputWithContext(ctx context.Context) DeviceRearPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceRearPortArrayOutput)
}

// DeviceRearPortMapInput is an input type that accepts DeviceRearPortMap and DeviceRearPortMapOutput values.
// You can construct a concrete instance of `DeviceRearPortMapInput` via:
//
//	DeviceRearPortMap{ "key": DeviceRearPortArgs{...} }
type DeviceRearPortMapInput interface {
	pulumi.Input

	ToDeviceRearPortMapOutput() DeviceRearPortMapOutput
	ToDeviceRearPortMapOutputWithContext(context.Context) DeviceRearPortMapOutput
}

type DeviceRearPortMap map[string]DeviceRearPortInput

func (DeviceRearPortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceRearPort)(nil)).Elem()
}

func (i DeviceRearPortMap) ToDeviceRearPortMapOutput() DeviceRearPortMapOutput {
	return i.ToDeviceRearPortMapOutputWithContext(context.Background())
}

func (i DeviceRearPortMap) ToDeviceRearPortMapOutputWithContext(ctx context.Context) DeviceRearPortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceRearPortMapOutput)
}

type DeviceRearPortOutput struct{ *pulumi.OutputState }

func (DeviceRearPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceRearPort)(nil)).Elem()
}

func (o DeviceRearPortOutput) ToDeviceRearPortOutput() DeviceRearPortOutput {
	return o
}

func (o DeviceRearPortOutput) ToDeviceRearPortOutputWithContext(ctx context.Context) DeviceRearPortOutput {
	return o
}

func (o DeviceRearPortOutput) ColorHex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.StringPtrOutput { return v.ColorHex }).(pulumi.StringPtrOutput)
}

func (o DeviceRearPortOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

func (o DeviceRearPortOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DeviceRearPortOutput) DeviceId() pulumi.IntOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.IntOutput { return v.DeviceId }).(pulumi.IntOutput)
}

func (o DeviceRearPortOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// Defaults to `false`.
func (o DeviceRearPortOutput) MarkConnected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.BoolPtrOutput { return v.MarkConnected }).(pulumi.BoolPtrOutput)
}

func (o DeviceRearPortOutput) ModuleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.IntPtrOutput { return v.ModuleId }).(pulumi.IntPtrOutput)
}

func (o DeviceRearPortOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeviceRearPortOutput) Positions() pulumi.IntOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.IntOutput { return v.Positions }).(pulumi.IntOutput)
}

func (o DeviceRearPortOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other].
func (o DeviceRearPortOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceRearPort) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type DeviceRearPortArrayOutput struct{ *pulumi.OutputState }

func (DeviceRearPortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceRearPort)(nil)).Elem()
}

func (o DeviceRearPortArrayOutput) ToDeviceRearPortArrayOutput() DeviceRearPortArrayOutput {
	return o
}

func (o DeviceRearPortArrayOutput) ToDeviceRearPortArrayOutputWithContext(ctx context.Context) DeviceRearPortArrayOutput {
	return o
}

func (o DeviceRearPortArrayOutput) Index(i pulumi.IntInput) DeviceRearPortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeviceRearPort {
		return vs[0].([]*DeviceRearPort)[vs[1].(int)]
	}).(DeviceRearPortOutput)
}

type DeviceRearPortMapOutput struct{ *pulumi.OutputState }

func (DeviceRearPortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceRearPort)(nil)).Elem()
}

func (o DeviceRearPortMapOutput) ToDeviceRearPortMapOutput() DeviceRearPortMapOutput {
	return o
}

func (o DeviceRearPortMapOutput) ToDeviceRearPortMapOutputWithContext(ctx context.Context) DeviceRearPortMapOutput {
	return o
}

func (o DeviceRearPortMapOutput) MapIndex(k pulumi.StringInput) DeviceRearPortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeviceRearPort {
		return vs[0].(map[string]*DeviceRearPort)[vs[1].(string)]
	}).(DeviceRearPortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceRearPortInput)(nil)).Elem(), &DeviceRearPort{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceRearPortArrayInput)(nil)).Elem(), DeviceRearPortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceRearPortMapInput)(nil)).Elem(), DeviceRearPortMap{})
	pulumi.RegisterOutputType(DeviceRearPortOutput{})
	pulumi.RegisterOutputType(DeviceRearPortArrayOutput{})
	pulumi.RegisterOutputType(DeviceRearPortMapOutput{})
}
