// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package kubeadm

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// Deprecated: register deep-copy functions.
func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// Deprecated: RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*API).DeepCopyInto(out.(*API))
			return nil
		}, InType: reflect.TypeOf(&API{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Etcd).DeepCopyInto(out.(*Etcd))
			return nil
		}, InType: reflect.TypeOf(&Etcd{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MasterConfiguration).DeepCopyInto(out.(*MasterConfiguration))
			return nil
		}, InType: reflect.TypeOf(&MasterConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Networking).DeepCopyInto(out.(*Networking))
			return nil
		}, InType: reflect.TypeOf(&Networking{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NodeConfiguration).DeepCopyInto(out.(*NodeConfiguration))
			return nil
		}, InType: reflect.TypeOf(&NodeConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TokenDiscovery).DeepCopyInto(out.(*TokenDiscovery))
			return nil
		}, InType: reflect.TypeOf(&TokenDiscovery{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (x *API) DeepCopy() *API {
	if x == nil {
		return nil
	}
	out := new(API)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (x *Etcd) DeepCopy() *Etcd {
	if x == nil {
		return nil
	}
	out := new(Etcd)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterConfiguration) DeepCopyInto(out *MasterConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.API = in.API
	in.Etcd.DeepCopyInto(&out.Etcd)
	out.Networking = in.Networking
	if in.AuthorizationModes != nil {
		in, out := &in.AuthorizationModes, &out.AuthorizationModes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIServerExtraArgs != nil {
		in, out := &in.APIServerExtraArgs, &out.APIServerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ControllerManagerExtraArgs != nil {
		in, out := &in.ControllerManagerExtraArgs, &out.ControllerManagerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SchedulerExtraArgs != nil {
		in, out := &in.SchedulerExtraArgs, &out.SchedulerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIServerCertSANs != nil {
		in, out := &in.APIServerCertSANs, &out.APIServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FeatureFlags != nil {
		in, out := &in.FeatureFlags, &out.FeatureFlags
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new MasterConfiguration.
func (x *MasterConfiguration) DeepCopy() *MasterConfiguration {
	if x == nil {
		return nil
	}
	out := new(MasterConfiguration)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *MasterConfiguration) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networking) DeepCopyInto(out *Networking) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Networking.
func (x *Networking) DeepCopy() *Networking {
	if x == nil {
		return nil
	}
	out := new(Networking)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfiguration) DeepCopyInto(out *NodeConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DiscoveryTokenAPIServers != nil {
		in, out := &in.DiscoveryTokenAPIServers, &out.DiscoveryTokenAPIServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfiguration.
func (x *NodeConfiguration) DeepCopy() *NodeConfiguration {
	if x == nil {
		return nil
	}
	out := new(NodeConfiguration)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *NodeConfiguration) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenDiscovery) DeepCopyInto(out *TokenDiscovery) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new TokenDiscovery.
func (x *TokenDiscovery) DeepCopy() *TokenDiscovery {
	if x == nil {
		return nil
	}
	out := new(TokenDiscovery)
	x.DeepCopyInto(out)
	return out
}
