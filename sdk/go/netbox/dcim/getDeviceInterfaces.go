// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcim

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDeviceInterfaces(ctx *pulumi.Context, args *GetDeviceInterfacesArgs, opts ...pulumi.InvokeOption) (*GetDeviceInterfacesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDeviceInterfacesResult
	err := ctx.Invoke("netbox:dcim/getDeviceInterfaces:getDeviceInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDeviceInterfaces.
type GetDeviceInterfacesArgs struct {
	Filters   []GetDeviceInterfacesFilter `pulumi:"filters"`
	NameRegex *string                     `pulumi:"nameRegex"`
}

// A collection of values returned by getDeviceInterfaces.
type GetDeviceInterfacesResult struct {
	Filters []GetDeviceInterfacesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                         `pulumi:"id"`
	Interfaces []GetDeviceInterfacesInterface `pulumi:"interfaces"`
	NameRegex  *string                        `pulumi:"nameRegex"`
}

func GetDeviceInterfacesOutput(ctx *pulumi.Context, args GetDeviceInterfacesOutputArgs, opts ...pulumi.InvokeOption) GetDeviceInterfacesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDeviceInterfacesResultOutput, error) {
			args := v.(GetDeviceInterfacesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("netbox:dcim/getDeviceInterfaces:getDeviceInterfaces", args, GetDeviceInterfacesResultOutput{}, options).(GetDeviceInterfacesResultOutput), nil
		}).(GetDeviceInterfacesResultOutput)
}

// A collection of arguments for invoking getDeviceInterfaces.
type GetDeviceInterfacesOutputArgs struct {
	Filters   GetDeviceInterfacesFilterArrayInput `pulumi:"filters"`
	NameRegex pulumi.StringPtrInput               `pulumi:"nameRegex"`
}

func (GetDeviceInterfacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceInterfacesArgs)(nil)).Elem()
}

// A collection of values returned by getDeviceInterfaces.
type GetDeviceInterfacesResultOutput struct{ *pulumi.OutputState }

func (GetDeviceInterfacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceInterfacesResult)(nil)).Elem()
}

func (o GetDeviceInterfacesResultOutput) ToGetDeviceInterfacesResultOutput() GetDeviceInterfacesResultOutput {
	return o
}

func (o GetDeviceInterfacesResultOutput) ToGetDeviceInterfacesResultOutputWithContext(ctx context.Context) GetDeviceInterfacesResultOutput {
	return o
}

func (o GetDeviceInterfacesResultOutput) Filters() GetDeviceInterfacesFilterArrayOutput {
	return o.ApplyT(func(v GetDeviceInterfacesResult) []GetDeviceInterfacesFilter { return v.Filters }).(GetDeviceInterfacesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDeviceInterfacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceInterfacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDeviceInterfacesResultOutput) Interfaces() GetDeviceInterfacesInterfaceArrayOutput {
	return o.ApplyT(func(v GetDeviceInterfacesResult) []GetDeviceInterfacesInterface { return v.Interfaces }).(GetDeviceInterfacesInterfaceArrayOutput)
}

func (o GetDeviceInterfacesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceInterfacesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeviceInterfacesResultOutput{})
}
