package delegated

import (
	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	projectapi "github.com/openshift/origin/pkg/project/api"
	templateapi "github.com/openshift/origin/pkg/template/api"
)

const (
	ProjectNameParam        = "PROJECT_NAME"
	ProjectDisplayNameParam = "PROJECT_DISPLAYNAME"
	ProjectDescriptionParam = "PROJECT_DESCRIPTION"
	ProjectAdminUserParam   = "PROJECT_ADMIN_USER"
)

var (
	parameters = []string{ProjectNameParam, ProjectDisplayNameParam, ProjectDescriptionParam, ProjectAdminUserParam}
)

func DefaultTemplate() *templateapi.Template {
	ret := &templateapi.Template{}
	ret.Name = "project-request"
	ret.Namespace = kapi.NamespaceDefault

	project := &projectapi.Project{}
	project.Name = "${" + ProjectNameParam + "}"
	project.Annotations = map[string]string{
		"description": "${" + ProjectDescriptionParam + "}",
		"displayName": "${" + ProjectDisplayNameParam + "}",
	}
	ret.Objects = append(ret.Objects, project)

	binding := &authorizationapi.RoleBinding{}
	binding.Name = "admins"
	binding.Namespace = "${" + ProjectNameParam + "}"
	binding.Users = util.NewStringSet("${" + ProjectAdminUserParam + "}")
	binding.RoleRef.Name = bootstrappolicy.AdminRoleName
	ret.Objects = append(ret.Objects, binding)

	for _, parameterName := range parameters {
		parameter := templateapi.Parameter{}
		parameter.Name = parameterName
		ret.Parameters = append(ret.Parameters, parameter)
	}

	return ret
}
