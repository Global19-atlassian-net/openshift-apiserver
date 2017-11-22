// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	v1 "github.com/openshift/api/project/v1"
	project "github.com/openshift/origin/pkg/project/apis/project"
	core_v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_Project_To_project_Project,
		Convert_project_Project_To_v1_Project,
		Convert_v1_ProjectList_To_project_ProjectList,
		Convert_project_ProjectList_To_v1_ProjectList,
		Convert_v1_ProjectRequest_To_project_ProjectRequest,
		Convert_project_ProjectRequest_To_v1_ProjectRequest,
		Convert_v1_ProjectSpec_To_project_ProjectSpec,
		Convert_project_ProjectSpec_To_v1_ProjectSpec,
		Convert_v1_ProjectStatus_To_project_ProjectStatus,
		Convert_project_ProjectStatus_To_v1_ProjectStatus,
	)
}

func autoConvert_v1_Project_To_project_Project(in *v1.Project, out *project.Project, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ProjectSpec_To_project_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ProjectStatus_To_project_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Project_To_project_Project is an autogenerated conversion function.
func Convert_v1_Project_To_project_Project(in *v1.Project, out *project.Project, s conversion.Scope) error {
	return autoConvert_v1_Project_To_project_Project(in, out, s)
}

func autoConvert_project_Project_To_v1_Project(in *project.Project, out *v1.Project, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_project_ProjectSpec_To_v1_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_project_ProjectStatus_To_v1_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_project_Project_To_v1_Project is an autogenerated conversion function.
func Convert_project_Project_To_v1_Project(in *project.Project, out *v1.Project, s conversion.Scope) error {
	return autoConvert_project_Project_To_v1_Project(in, out, s)
}

func autoConvert_v1_ProjectList_To_project_ProjectList(in *v1.ProjectList, out *project.ProjectList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]project.Project)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ProjectList_To_project_ProjectList is an autogenerated conversion function.
func Convert_v1_ProjectList_To_project_ProjectList(in *v1.ProjectList, out *project.ProjectList, s conversion.Scope) error {
	return autoConvert_v1_ProjectList_To_project_ProjectList(in, out, s)
}

func autoConvert_project_ProjectList_To_v1_ProjectList(in *project.ProjectList, out *v1.ProjectList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.Project)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_project_ProjectList_To_v1_ProjectList is an autogenerated conversion function.
func Convert_project_ProjectList_To_v1_ProjectList(in *project.ProjectList, out *v1.ProjectList, s conversion.Scope) error {
	return autoConvert_project_ProjectList_To_v1_ProjectList(in, out, s)
}

func autoConvert_v1_ProjectRequest_To_project_ProjectRequest(in *v1.ProjectRequest, out *project.ProjectRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}

// Convert_v1_ProjectRequest_To_project_ProjectRequest is an autogenerated conversion function.
func Convert_v1_ProjectRequest_To_project_ProjectRequest(in *v1.ProjectRequest, out *project.ProjectRequest, s conversion.Scope) error {
	return autoConvert_v1_ProjectRequest_To_project_ProjectRequest(in, out, s)
}

func autoConvert_project_ProjectRequest_To_v1_ProjectRequest(in *project.ProjectRequest, out *v1.ProjectRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}

// Convert_project_ProjectRequest_To_v1_ProjectRequest is an autogenerated conversion function.
func Convert_project_ProjectRequest_To_v1_ProjectRequest(in *project.ProjectRequest, out *v1.ProjectRequest, s conversion.Scope) error {
	return autoConvert_project_ProjectRequest_To_v1_ProjectRequest(in, out, s)
}

func autoConvert_v1_ProjectSpec_To_project_ProjectSpec(in *v1.ProjectSpec, out *project.ProjectSpec, s conversion.Scope) error {
	out.Finalizers = *(*[]api.FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}

// Convert_v1_ProjectSpec_To_project_ProjectSpec is an autogenerated conversion function.
func Convert_v1_ProjectSpec_To_project_ProjectSpec(in *v1.ProjectSpec, out *project.ProjectSpec, s conversion.Scope) error {
	return autoConvert_v1_ProjectSpec_To_project_ProjectSpec(in, out, s)
}

func autoConvert_project_ProjectSpec_To_v1_ProjectSpec(in *project.ProjectSpec, out *v1.ProjectSpec, s conversion.Scope) error {
	out.Finalizers = *(*[]core_v1.FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}

// Convert_project_ProjectSpec_To_v1_ProjectSpec is an autogenerated conversion function.
func Convert_project_ProjectSpec_To_v1_ProjectSpec(in *project.ProjectSpec, out *v1.ProjectSpec, s conversion.Scope) error {
	return autoConvert_project_ProjectSpec_To_v1_ProjectSpec(in, out, s)
}

func autoConvert_v1_ProjectStatus_To_project_ProjectStatus(in *v1.ProjectStatus, out *project.ProjectStatus, s conversion.Scope) error {
	out.Phase = api.NamespacePhase(in.Phase)
	return nil
}

// Convert_v1_ProjectStatus_To_project_ProjectStatus is an autogenerated conversion function.
func Convert_v1_ProjectStatus_To_project_ProjectStatus(in *v1.ProjectStatus, out *project.ProjectStatus, s conversion.Scope) error {
	return autoConvert_v1_ProjectStatus_To_project_ProjectStatus(in, out, s)
}

func autoConvert_project_ProjectStatus_To_v1_ProjectStatus(in *project.ProjectStatus, out *v1.ProjectStatus, s conversion.Scope) error {
	out.Phase = core_v1.NamespacePhase(in.Phase)
	return nil
}

// Convert_project_ProjectStatus_To_v1_ProjectStatus is an autogenerated conversion function.
func Convert_project_ProjectStatus_To_v1_ProjectStatus(in *project.ProjectStatus, out *v1.ProjectStatus, s conversion.Scope) error {
	return autoConvert_project_ProjectStatus_To_v1_ProjectStatus(in, out, s)
}
