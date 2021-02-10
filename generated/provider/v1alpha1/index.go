/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

/*
import (
	"context"
	"reflect"
	"github.com/crossplane-contrib/terraform-runtime/pkg/client"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

 */

import (
	"context"
	"fmt"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	"reflect"

	"github.com/crossplane-contrib/terraform-runtime/pkg/client"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"github.com/zclconf/go-cty/cty"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	kubeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group                         = "vsphere.terraform-plugin.crossplane.io"
	Version                       = "v1alpha1"
	errProviderNotRetrieved       = "provider could not be retrieved"
	errProviderSecretNotRetrieved = "secret referred in provider could not be retrieved"
	errProviderSecretNil          = "cannot find Secret reference on Provider"
	ProviderName                  = "vsphere"
)

func GetProviderInit() *plugin.ProviderInit {
	schemeBuilder := &scheme.Builder{GroupVersion: SchemeGroupVersion}
	schemeBuilder.Register(&ProviderConfig{}, &ProviderConfigList{})
	schemeBuilder.Register(&ProviderConfigUsage{}, &ProviderConfigUsageList{})
	return &plugin.ProviderInit{
		SchemeBuilder: schemeBuilder,
		Initializer:   initializeProvider,
	}
}

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}
	// Provider type metadata.
	ProviderKind             = reflect.TypeOf(ProviderConfig{}).Name()
	ProviderGroupKind        = schema.GroupKind{Group: Group, Kind: ProviderKind}.String()
	ProviderKindAPIVersion   = ProviderKind + "." + SchemeGroupVersion.String()
	ProviderGroupVersionKind = SchemeGroupVersion.WithKind(ProviderKind)
)


func initializeProvider(ctx context.Context, mr resource.Managed, ropts *client.RuntimeOptions, kube kubeclient.Client) (*client.Provider, error) {
	pc := &ProviderConfig{}
	if err := kube.Get(ctx, types.NamespacedName{Name: mr.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, "cannot get referenced Provider")
	}

	t := resource.NewProviderConfigUsageTracker(kube, &ProviderConfigUsage{})
	if err := t.Track(ctx, mr); err != nil {
		return nil, errors.Wrap(err, "cannot track ProviderConfig usage")
	}

	//creds, err := providerConfigCredentials(ctx, kube, pc)
	//cfg := populateConfig(pc, creds)
	up, err := readCredentials(ctx, kube, mr)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read credentials from ProviderConfig")
	}
	cfg := populateConfig(pc, up)

	p, err := client.NewProvider(ProviderName, ropts.PluginPath)
	if err != nil {
		return p, err
	}
	err = p.Configure(cfg)
	return p, err
}

func populateConfig(p *ProviderConfig, up *UserPass) cty.Value {
	merged := make(map[string]cty.Value)
	merged["api_timeout"] = cty.NumberIntVal(p.Spec.ApiTimeout)
	merged["rest_session_path"] = cty.StringVal(p.Spec.RestSessionPath)
	merged["vcenter_server"] = cty.StringVal(p.Spec.VcenterServer)
	merged["vim_keep_alive"] = cty.NumberIntVal(p.Spec.VimKeepAlive)
	merged["allow_unverified_ssl"] = cty.BoolVal(p.Spec.AllowUnverifiedSsl)
	merged["client_debug"] = cty.BoolVal(p.Spec.ClientDebug)
	merged["client_debug_path"] = cty.StringVal(p.Spec.ClientDebugPath)
	merged["client_debug_path_run"] = cty.StringVal(p.Spec.ClientDebugPathRun)
	merged["persist_session"] = cty.BoolVal(p.Spec.PersistSession)
	merged["vim_session_path"] = cty.StringVal(p.Spec.VimSessionPath)
	merged["vsphere_server"] = cty.StringVal(p.Spec.VsphereServer)

	merged["user"] = cty.StringVal(up.User)
	merged["password"] = cty.StringVal(up.Pass)
	return cty.ObjectVal(merged)
}

type UserPass struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func readCredentials(ctx context.Context, kube kubeclient.Client, mr resource.Managed) (*UserPass, error) {
	pc := &ProviderConfig{}
	t := resource.NewProviderConfigUsageTracker(kube, &ProviderConfigUsage{})
	if err := t.Track(ctx, mr); err != nil {
		return nil, err
	}
	if err := kube.Get(ctx, types.NamespacedName{Name: mr.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, err
	}

	if s := pc.Spec.Credentials.Source; s != xpv1.CredentialsSourceSecret {
		return nil, errors.Errorf("unsupported credentials source %q", s)
	}

	ref := pc.Spec.Credentials.SecretRef
	if ref == nil {
		return nil, errors.New("no credentials secret reference was provided")
	}

	s := &corev1.Secret{}
	if err := kube.Get(ctx, types.NamespacedName{Name: ref.Name, Namespace: ref.Namespace}, s); err != nil {
		return nil, err
	}

	return userPassFromSecret(s)
}

func userPassFromSecret(s *corev1.Secret) (*UserPass, error) {
	if s.Type != corev1.SecretTypeBasicAuth {
		basicAuthUrl := "https://kubernetes.io/docs/concepts/configuration/secret/#basic-authentication-secret"
		return nil, fmt.Errorf("expected key to be of type %s -- see %s", corev1.SecretTypeBasicAuth, basicAuthUrl)
	}

	// we can assume these fields exist because k8s api guarantees them based on the .Type
	// but we need to check their length because k8s api only requires one of them to be non-zero
	username := s.Data[corev1.BasicAuthUsernameKey]
	if len(username) == 0 {
		return nil, errors.New("username field of basic-auth provider credential secret is empty")
	}
	password := s.Data[corev1.BasicAuthPasswordKey]
	if len(username) == 0 {
		return nil, errors.New("password field of basic-auth provider credential secret is empty")
	}
	return &UserPass{
		User: string(username),
		Pass: string(password),
	}, nil
}
