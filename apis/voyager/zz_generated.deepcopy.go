// +build !ignore_autogenerated

/*
Copyright 2017 The Stash Authors.

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

package voyager

import (
	reflect "reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/client-go/pkg/api/v1"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_ACMECertificateDetails, InType: reflect.TypeOf(&ACMECertificateDetails{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_Certificate, InType: reflect.TypeOf(&Certificate{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_CertificateList, InType: reflect.TypeOf(&CertificateList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_CertificateSpec, InType: reflect.TypeOf(&CertificateSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_CertificateStatus, InType: reflect.TypeOf(&CertificateStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_DNSResolver, InType: reflect.TypeOf(&DNSResolver{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_FrontendRule, InType: reflect.TypeOf(&FrontendRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_HTTPIngressBackend, InType: reflect.TypeOf(&HTTPIngressBackend{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_HTTPIngressPath, InType: reflect.TypeOf(&HTTPIngressPath{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_HTTPIngressRuleValue, InType: reflect.TypeOf(&HTTPIngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_Ingress, InType: reflect.TypeOf(&Ingress{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_IngressBackend, InType: reflect.TypeOf(&IngressBackend{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_IngressList, InType: reflect.TypeOf(&IngressList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_IngressRef, InType: reflect.TypeOf(&IngressRef{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_IngressRule, InType: reflect.TypeOf(&IngressRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_IngressRuleValue, InType: reflect.TypeOf(&IngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_IngressSpec, InType: reflect.TypeOf(&IngressSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_IngressStatus, InType: reflect.TypeOf(&IngressStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_IngressTLS, InType: reflect.TypeOf(&IngressTLS{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_MonitorSpec, InType: reflect.TypeOf(&MonitorSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_PrometheusSpec, InType: reflect.TypeOf(&PrometheusSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_TCPIngressRuleValue, InType: reflect.TypeOf(&TCPIngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_voyager_Target, InType: reflect.TypeOf(&Target{})},
	)
}

// DeepCopy_voyager_ACMECertificateDetails is an autogenerated deepcopy function.
func DeepCopy_voyager_ACMECertificateDetails(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ACMECertificateDetails)
		out := out.(*ACMECertificateDetails)
		*out = *in
		return nil
	}
}

// DeepCopy_voyager_Certificate is an autogenerated deepcopy function.
func DeepCopy_voyager_Certificate(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Certificate)
		out := out.(*Certificate)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_voyager_CertificateSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_voyager_CertificateStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_voyager_CertificateList is an autogenerated deepcopy function.
func DeepCopy_voyager_CertificateList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateList)
		out := out.(*CertificateList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Certificate, len(*in))
			for i := range *in {
				if err := DeepCopy_voyager_Certificate(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_voyager_CertificateSpec is an autogenerated deepcopy function.
func DeepCopy_voyager_CertificateSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSpec)
		out := out.(*CertificateSpec)
		*out = *in
		if in.Domains != nil {
			in, out := &in.Domains, &out.Domains
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_voyager_CertificateStatus is an autogenerated deepcopy function.
func DeepCopy_voyager_CertificateStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateStatus)
		out := out.(*CertificateStatus)
		*out = *in
		if in.CreationTime != nil {
			in, out := &in.CreationTime, &out.CreationTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		return nil
	}
}

// DeepCopy_voyager_DNSResolver is an autogenerated deepcopy function.
func DeepCopy_voyager_DNSResolver(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DNSResolver)
		out := out.(*DNSResolver)
		*out = *in
		if in.NameServer != nil {
			in, out := &in.NameServer, &out.NameServer
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Timeout != nil {
			in, out := &in.Timeout, &out.Timeout
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.Hold != nil {
			in, out := &in.Hold, &out.Hold
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_voyager_FrontendRule is an autogenerated deepcopy function.
func DeepCopy_voyager_FrontendRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*FrontendRule)
		out := out.(*FrontendRule)
		*out = *in
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_voyager_HTTPIngressBackend is an autogenerated deepcopy function.
func DeepCopy_voyager_HTTPIngressBackend(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HTTPIngressBackend)
		out := out.(*HTTPIngressBackend)
		*out = *in
		if err := DeepCopy_voyager_IngressBackend(&in.IngressBackend, &out.IngressBackend, c); err != nil {
			return err
		}
		if in.RewriteRule != nil {
			in, out := &in.RewriteRule, &out.RewriteRule
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.HeaderRule != nil {
			in, out := &in.HeaderRule, &out.HeaderRule
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_voyager_HTTPIngressPath is an autogenerated deepcopy function.
func DeepCopy_voyager_HTTPIngressPath(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HTTPIngressPath)
		out := out.(*HTTPIngressPath)
		*out = *in
		if err := DeepCopy_voyager_HTTPIngressBackend(&in.Backend, &out.Backend, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_voyager_HTTPIngressRuleValue is an autogenerated deepcopy function.
func DeepCopy_voyager_HTTPIngressRuleValue(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HTTPIngressRuleValue)
		out := out.(*HTTPIngressRuleValue)
		*out = *in
		if in.Paths != nil {
			in, out := &in.Paths, &out.Paths
			*out = make([]HTTPIngressPath, len(*in))
			for i := range *in {
				if err := DeepCopy_voyager_HTTPIngressPath(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_voyager_Ingress is an autogenerated deepcopy function.
func DeepCopy_voyager_Ingress(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Ingress)
		out := out.(*Ingress)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_voyager_IngressSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_voyager_IngressStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_voyager_IngressBackend is an autogenerated deepcopy function.
func DeepCopy_voyager_IngressBackend(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressBackend)
		out := out.(*IngressBackend)
		*out = *in
		if in.HostNames != nil {
			in, out := &in.HostNames, &out.HostNames
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.BackendRule != nil {
			in, out := &in.BackendRule, &out.BackendRule
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_voyager_IngressList is an autogenerated deepcopy function.
func DeepCopy_voyager_IngressList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressList)
		out := out.(*IngressList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Ingress, len(*in))
			for i := range *in {
				if err := DeepCopy_voyager_Ingress(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_voyager_IngressRef is an autogenerated deepcopy function.
func DeepCopy_voyager_IngressRef(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressRef)
		out := out.(*IngressRef)
		*out = *in
		return nil
	}
}

// DeepCopy_voyager_IngressRule is an autogenerated deepcopy function.
func DeepCopy_voyager_IngressRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressRule)
		out := out.(*IngressRule)
		*out = *in
		if err := DeepCopy_voyager_IngressRuleValue(&in.IngressRuleValue, &out.IngressRuleValue, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_voyager_IngressRuleValue is an autogenerated deepcopy function.
func DeepCopy_voyager_IngressRuleValue(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressRuleValue)
		out := out.(*IngressRuleValue)
		*out = *in
		if in.HTTP != nil {
			in, out := &in.HTTP, &out.HTTP
			*out = new(HTTPIngressRuleValue)
			if err := DeepCopy_voyager_HTTPIngressRuleValue(*in, *out, c); err != nil {
				return err
			}
		}
		if in.TCP != nil {
			in, out := &in.TCP, &out.TCP
			*out = new(TCPIngressRuleValue)
			if err := DeepCopy_voyager_TCPIngressRuleValue(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_voyager_IngressSpec is an autogenerated deepcopy function.
func DeepCopy_voyager_IngressSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressSpec)
		out := out.(*IngressSpec)
		*out = *in
		if in.Backend != nil {
			in, out := &in.Backend, &out.Backend
			*out = new(HTTPIngressBackend)
			if err := DeepCopy_voyager_HTTPIngressBackend(*in, *out, c); err != nil {
				return err
			}
		}
		if in.TLS != nil {
			in, out := &in.TLS, &out.TLS
			*out = make([]IngressTLS, len(*in))
			for i := range *in {
				if err := DeepCopy_voyager_IngressTLS(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.FrontendRules != nil {
			in, out := &in.FrontendRules, &out.FrontendRules
			*out = make([]FrontendRule, len(*in))
			for i := range *in {
				if err := DeepCopy_voyager_FrontendRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]IngressRule, len(*in))
			for i := range *in {
				if err := DeepCopy_voyager_IngressRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.LoadBalancerSourceRanges != nil {
			in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if newVal, err := c.DeepCopy(&in.Resources); err != nil {
			return err
		} else {
			out.Resources = *newVal.(*api_v1.ResourceRequirements)
		}
		return nil
	}
}

// DeepCopy_voyager_IngressStatus is an autogenerated deepcopy function.
func DeepCopy_voyager_IngressStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressStatus)
		out := out.(*IngressStatus)
		*out = *in
		if newVal, err := c.DeepCopy(&in.LoadBalancer); err != nil {
			return err
		} else {
			out.LoadBalancer = *newVal.(*api_v1.LoadBalancerStatus)
		}
		return nil
	}
}

// DeepCopy_voyager_IngressTLS is an autogenerated deepcopy function.
func DeepCopy_voyager_IngressTLS(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressTLS)
		out := out.(*IngressTLS)
		*out = *in
		if in.Hosts != nil {
			in, out := &in.Hosts, &out.Hosts
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_voyager_MonitorSpec is an autogenerated deepcopy function.
func DeepCopy_voyager_MonitorSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*MonitorSpec)
		out := out.(*MonitorSpec)
		*out = *in
		if in.Prometheus != nil {
			in, out := &in.Prometheus, &out.Prometheus
			*out = new(PrometheusSpec)
			if err := DeepCopy_voyager_PrometheusSpec(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_voyager_PrometheusSpec is an autogenerated deepcopy function.
func DeepCopy_voyager_PrometheusSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PrometheusSpec)
		out := out.(*PrometheusSpec)
		*out = *in
		if in.Labels != nil {
			in, out := &in.Labels, &out.Labels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_voyager_TCPIngressRuleValue is an autogenerated deepcopy function.
func DeepCopy_voyager_TCPIngressRuleValue(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TCPIngressRuleValue)
		out := out.(*TCPIngressRuleValue)
		*out = *in
		if err := DeepCopy_voyager_IngressBackend(&in.Backend, &out.Backend, c); err != nil {
			return err
		}
		if in.ALPN != nil {
			in, out := &in.ALPN, &out.ALPN
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_voyager_Target is an autogenerated deepcopy function.
func DeepCopy_voyager_Target(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Target)
		out := out.(*Target)
		*out = *in
		return nil
	}
}
