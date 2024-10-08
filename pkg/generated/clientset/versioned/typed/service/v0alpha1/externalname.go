// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by client-gen. DO NOT EDIT.

package v0alpha1

import (
	"context"

	v0alpha1 "github.com/grafana/grafana/pkg/apis/service/v0alpha1"
	servicev0alpha1 "github.com/grafana/grafana/pkg/generated/applyconfiguration/service/v0alpha1"
	scheme "github.com/grafana/grafana/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ExternalNamesGetter has a method to return a ExternalNameInterface.
// A group's client should implement this interface.
type ExternalNamesGetter interface {
	ExternalNames(namespace string) ExternalNameInterface
}

// ExternalNameInterface has methods to work with ExternalName resources.
type ExternalNameInterface interface {
	Create(ctx context.Context, externalName *v0alpha1.ExternalName, opts v1.CreateOptions) (*v0alpha1.ExternalName, error)
	Update(ctx context.Context, externalName *v0alpha1.ExternalName, opts v1.UpdateOptions) (*v0alpha1.ExternalName, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v0alpha1.ExternalName, error)
	List(ctx context.Context, opts v1.ListOptions) (*v0alpha1.ExternalNameList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v0alpha1.ExternalName, err error)
	Apply(ctx context.Context, externalName *servicev0alpha1.ExternalNameApplyConfiguration, opts v1.ApplyOptions) (result *v0alpha1.ExternalName, err error)
	ExternalNameExpansion
}

// externalNames implements ExternalNameInterface
type externalNames struct {
	*gentype.ClientWithListAndApply[*v0alpha1.ExternalName, *v0alpha1.ExternalNameList, *servicev0alpha1.ExternalNameApplyConfiguration]
}

// newExternalNames returns a ExternalNames
func newExternalNames(c *ServiceV0alpha1Client, namespace string) *externalNames {
	return &externalNames{
		gentype.NewClientWithListAndApply[*v0alpha1.ExternalName, *v0alpha1.ExternalNameList, *servicev0alpha1.ExternalNameApplyConfiguration](
			"externalnames",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v0alpha1.ExternalName { return &v0alpha1.ExternalName{} },
			func() *v0alpha1.ExternalNameList { return &v0alpha1.ExternalNameList{} }),
	}
}
