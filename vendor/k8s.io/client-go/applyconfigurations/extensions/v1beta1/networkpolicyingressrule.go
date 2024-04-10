/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// NetworkPolicyIngressRuleApplyConfiguration represents an declarative configuration of the NetworkPolicyIngressRule type for use
// with apply.
type NetworkPolicyIngressRuleApplyConfiguration struct {
	Ports []NetworkPolicyPortApplyConfiguration `json:"ports,omitempty"`
	From  []NetworkPolicyPeerApplyConfiguration `json:"from,omitempty"`
}

// NetworkPolicyIngressRuleApplyConfiguration constructs an declarative configuration of the NetworkPolicyIngressRule type for use with
// apply.
func NetworkPolicyIngressRule() *NetworkPolicyIngressRuleApplyConfiguration {
	return &NetworkPolicyIngressRuleApplyConfiguration{}
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *NetworkPolicyIngressRuleApplyConfiguration) WithPorts(values ...*NetworkPolicyPortApplyConfiguration) *NetworkPolicyIngressRuleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}

// WithFrom adds the given value to the From field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the From field.
func (b *NetworkPolicyIngressRuleApplyConfiguration) WithFrom(values ...*NetworkPolicyPeerApplyConfiguration) *NetworkPolicyIngressRuleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFrom")
		}
		b.From = append(b.From, *values[i])
	}
	return b
}
