/*
Copyright 2024 Nokia.

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

import (
	"reflect"

	conditionv1alpha1 "github.com/kuidio/kuid/apis/condition/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkDeviceDeviceSpec defines the desired state of NetworkDevice
type NetworkDeviceSpec struct {
	Topology string `json:"topology" yaml:"topology" protobuf:"bytes,1,opt,name=topology"`
	// Interfaces defines the interfaces of the device config
	// +optional
	Interfaces []*NetworkDeviceInterface `json:"interfaces,omitempty" yaml:"interfaces,omitempty" protobuf:"bytes,2,rep,name=topology"`
	// NetworkInstances defines the network instances of the device config
	// +optional
	NetworkInstances []*NetworkDeviceNetworkInstance `json:"networkInstances,omitempty" yaml:"networkInstances,omitempty" protobuf:"bytes,2,rep,name=topology"`
}

// TODO LAG, etc
type NetworkDeviceInterface struct {
	Name string `json:"name" yaml:"name" protobuf:"bytes,1,opt,name=name"`
	// tunnel, regular, etc
	InterfaceType string                               `json:"type" yaml:"type" protobuf:"bytes,2,opt,name=type"`
	SubInterfaces []*NetworkDeviceInterfaceSubInterface `json:"subInterfaces,omitempty" yaml:"subInterfaces,omitempty" protobuf:"bytes,3,rep,name=subInterfaces"`
	VLANTagging   bool                                 `json:"vlanTagging" yaml:"vlanTagging" protobuf:"bytes,4,opt,name=vlanTagging"`
	Speed         string                               `json:"speed" yaml:"speed" protobuf:"bytes,5,opt,name=speed"`
	LAGMember     bool                                 `json:"lagMember" yaml:"lagMember" protobuf:"bytes,6,opt,name=lagMember"`

}

type NetworkDeviceInterfaceSubInterface struct {
	ID uint32 `json:"id" yaml:"id" protobuf:"bytes,1,opt,name=id"`
	// routed or bridged
	SubInterfaceType string                                    `json:"type" yaml:"type" protobuf:"bytes,2,opt,name=type"`
	VLAN             *uint32                                   `json:"vlan,omitempty" yaml:"vlan,omitempty" protobuf:"bytes,3,opt,name=vlan"`
	IPv4             *NetworkDeviceInterfaceSubInterfaceIPv4 `json:"ipv4,omitempty" yaml:"ipv4,omitempty" protobuf:"bytes,4,rep,name=ipv4"`
	IPv6             *NetworkDeviceInterfaceSubInterfaceIPv6 `json:"ipv6,omitempty" yaml:"ipv6,omitempty" protobuf:"bytes,5,rep,name=ipv6"`
}

type NetworkDeviceInterfaceSubInterfaceIPv4 struct {
	Addresses []string `json:"addresses" yaml:"addresses" protobuf:"bytes,1,opt,name=addresses"`
}

type NetworkDeviceInterfaceSubInterfaceIPv6 struct {
	Addresses []string `json:"addresses" yaml:"addresses" protobuf:"bytes,1,opt,name=addresses"`
}

type NetworkDeviceNetworkInstance struct {
	Name string `json:"name" yaml:"name" protobuf:"bytes,1,opt,name=name"`
	// mac-vrf, ip-vrf
	NetworkInstanceType NetworkInstanceType                    `json:"type" yaml:"type" protobuf:"bytes,2,opt,name=type"`
	Protocols           *NetworkDeviceNetworkInstanceProtocols `json:"protocols,omitempty" yaml:"protocols,omitempty" protobuf:"bytes,3,opt,name=protocols"`
	Interfaces          []string                               `json:"interfaces,omitempty" yaml:"interfaces,omitempty" protobuf:"bytes,4,opt,name=interfaces"`
	VXLANInterface      string                                 `json:"vxlanInterface,omitempty" yaml:"vxlanInterface,omitempty" protobuf:"bytes,5,opt,name=vxlanInterface"`
}

type NetworkDeviceNetworkInstanceProtocols struct {
	BGP  *NetworkDeviceNetworkInstanceProtocolBGP  `json:"bgp,omitempty" yaml:"bgp,omitempty" protobuf:"bytes,1,opt,name=bgp"`
	EVPN *NetworkDeviceNetworkInstanceProtocolEVPN `json:"evpn,omitempty" yaml:"evpn,omitempty" protobuf:"bytes,2,opt,name=evpn"`
}

type NetworkDeviceNetworkInstanceProtocolBGP struct {
	AS         uint32                                              `json:"as" yaml:"as" protobuf:"bytes,1,opt,name=as"`
	RouterID   string                                              `json:"routerID" yaml:"routerID" protobuf:"bytes,2,opt,name=routerID"`
	PeerGroups []*NetworkDeviceNetworkInstanceProtocolBGPPeerGroup `json:"peerGroups,omitempty" yaml:"peerGroups,omitempty" protobuf:"bytes,3,opt,name=peerGroups"`
	Neighbors  []*NetworkDeviceNetworkInstanceProtocolBGPNeighbor  `json:"neighbors,omitempty" yaml:"neighbors,omitempty" protobuf:"bytes,4,opt,name=neighbors"`
}

type NetworkDeviceNetworkInstanceProtocolEVPN struct {
}

type NetworkDeviceNetworkInstanceProtocolBGPPeerGroup struct {
	Name            string   `json:"name" yaml:"name" protobuf:"bytes,1,opt,name=name"`
	AddressFamilies []string `json:"addressFamilies,omitempty" yaml:"addressFamilies,omitempty" protobuf:"bytes,2,rep,name=addressFamilies"`
}

type NetworkDeviceNetworkInstanceProtocolBGPNeighbor struct {
	PeerAddress  string `json:"peerAddress" yaml:"peerAddress" protobuf:"bytes,1,opt,name=peerAddress"`
	PeerAS       uint32 `json:"peerAS" yaml:"peerAS" protobuf:"bytes,2,opt,name=peerAS"`
	PeerGroup    string `json:"peerGroup" yaml:"peerGroup" protobuf:"bytes,2,opt,name=peerGroup"`
	LocalAS      uint32 `json:"localAS" yaml:"localAS" protobuf:"bytes,2,opt,name=localAS"`
	LocalAddress string `json:"localAddress" yaml:"localAddress" protobuf:"bytes,2,opt,name=localAddress"`
}

// NetworkDeviceStatus defines the observed state of NetworkDevice
type NetworkDeviceStatus struct {
	// ConditionedStatus provides the status of the NetworkDevice using conditions
	// - a ready condition indicates the overall status of the resource
	conditionv1alpha1.ConditionedStatus `json:",inline" yaml:",inline" protobuf:"bytes,1,opt,name=conditionedStatus"`
}

// +kubebuilder:object:root=true
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:resource:categories={kuid, net}
// NetworkDevice is the NetworkDevice for the NetworkDevice API
// +k8s:openapi-gen=true
type NetworkDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   NetworkDeviceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status NetworkDeviceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +kubebuilder:object:root=true
// NetworkDeviceClabList contains a list of NetworkDeviceClabs
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NetworkDeviceList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []NetworkDevice `json:"items" yaml:"items" protobuf:"bytes,2,rep,name=items"`
}

var (
	NetworkDeviceKind = reflect.TypeOf(NetworkDevice{}).Name()
)
