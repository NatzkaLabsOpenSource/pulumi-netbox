// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcim

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDevices(ctx *pulumi.Context, args *GetDevicesArgs, opts ...pulumi.InvokeOption) (*GetDevicesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDevicesResult
	err := ctx.Invoke("netbox:dcim/getDevices:getDevices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDevices.
type GetDevicesArgs struct {
	Filters   []GetDevicesFilter `pulumi:"filters"`
	Limit     *int               `pulumi:"limit"`
	NameRegex *string            `pulumi:"nameRegex"`
}

// A collection of values returned by getDevices.
type GetDevicesResult struct {
	Devices []GetDevicesDevice `pulumi:"devices"`
	Filters []GetDevicesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Limit     *int    `pulumi:"limit"`
	NameRegex *string `pulumi:"nameRegex"`
}

func GetDevicesOutput(ctx *pulumi.Context, args GetDevicesOutputArgs, opts ...pulumi.InvokeOption) GetDevicesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDevicesResultOutput, error) {
			args := v.(GetDevicesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:dcim/getDevices:getDevices", args, GetDevicesResultOutput{}, options).(GetDevicesResultOutput), nil
		}).(GetDevicesResultOutput)
}

// A collection of arguments for invoking getDevices.
type GetDevicesOutputArgs struct {
	Filters   GetDevicesFilterArrayInput `pulumi:"filters"`
	Limit     pulumi.IntPtrInput         `pulumi:"limit"`
	NameRegex pulumi.StringPtrInput      `pulumi:"nameRegex"`
}

func (GetDevicesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDevicesArgs)(nil)).Elem()
}

// A collection of values returned by getDevices.
type GetDevicesResultOutput struct{ *pulumi.OutputState }

func (GetDevicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDevicesResult)(nil)).Elem()
}

func (o GetDevicesResultOutput) ToGetDevicesResultOutput() GetDevicesResultOutput {
	return o
}

func (o GetDevicesResultOutput) ToGetDevicesResultOutputWithContext(ctx context.Context) GetDevicesResultOutput {
	return o
}

func (o GetDevicesResultOutput) Devices() GetDevicesDeviceArrayOutput {
	return o.ApplyT(func(v GetDevicesResult) []GetDevicesDevice { return v.Devices }).(GetDevicesDeviceArrayOutput)
}

func (o GetDevicesResultOutput) Filters() GetDevicesFilterArrayOutput {
	return o.ApplyT(func(v GetDevicesResult) []GetDevicesFilter { return v.Filters }).(GetDevicesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDevicesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDevicesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDevicesResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDevicesResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetDevicesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDevicesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDevicesResultOutput{})
}
