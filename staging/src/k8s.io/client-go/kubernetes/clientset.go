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

// Code generated by client-gen. DO NOT EDIT.

package kubernetes

import (
	fmt "fmt"
	http "net/http"

	discovery "k8s.io/client-go/discovery"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	admissionregistrationv1alpha1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	admissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	internalv1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authenticationv1alpha1 "k8s.io/client-go/kubernetes/typed/authentication/v1alpha1"
	authenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	authorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	autoscalingv2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	autoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	certificatesv1 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	certificatesv1alpha1 "k8s.io/client-go/kubernetes/typed/certificates/v1alpha1"
	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	coordinationv1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	coordinationv1alpha2 "k8s.io/client-go/kubernetes/typed/coordination/v1alpha2"
	coordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	discoveryv1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	discoveryv1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	eventsv1 "k8s.io/client-go/kubernetes/typed/events/v1"
	eventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	flowcontrolv1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1"
	flowcontrolv1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	flowcontrolv1beta2 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	flowcontrolv1beta3 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta3"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	networkingv1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	nodev1 "k8s.io/client-go/kubernetes/typed/node/v1"
	nodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	nodev1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	policyv1 "k8s.io/client-go/kubernetes/typed/policy/v1"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	resourcev1 "k8s.io/client-go/kubernetes/typed/resource/v1"
	resourcev1alpha3 "k8s.io/client-go/kubernetes/typed/resource/v1alpha3"
	resourcev1beta1 "k8s.io/client-go/kubernetes/typed/resource/v1beta1"
	resourcev1beta2 "k8s.io/client-go/kubernetes/typed/resource/v1beta2"
	schedulingv1 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	storagemigrationv1alpha1 "k8s.io/client-go/kubernetes/typed/storagemigration/v1alpha1"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AdmissionregistrationV1() admissionregistrationv1.AdmissionregistrationV1Interface
	AdmissionregistrationV1alpha1() admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Interface
	AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface
	InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface
	AppsV1() appsv1.AppsV1Interface
	AppsV1beta1() appsv1beta1.AppsV1beta1Interface
	AppsV1beta2() appsv1beta2.AppsV1beta2Interface
	AuthenticationV1() authenticationv1.AuthenticationV1Interface
	AuthenticationV1alpha1() authenticationv1alpha1.AuthenticationV1alpha1Interface
	AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface
	AuthorizationV1() authorizationv1.AuthorizationV1Interface
	AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface
	AutoscalingV1() autoscalingv1.AutoscalingV1Interface
	AutoscalingV2() autoscalingv2.AutoscalingV2Interface
	AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface
	AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface
	BatchV1() batchv1.BatchV1Interface
	BatchV1beta1() batchv1beta1.BatchV1beta1Interface
	CertificatesV1() certificatesv1.CertificatesV1Interface
	CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface
	CertificatesV1alpha1() certificatesv1alpha1.CertificatesV1alpha1Interface
	CoordinationV1alpha2() coordinationv1alpha2.CoordinationV1alpha2Interface
	CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface
	CoordinationV1() coordinationv1.CoordinationV1Interface
	CoreV1() corev1.CoreV1Interface
	DiscoveryV1() discoveryv1.DiscoveryV1Interface
	DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface
	EventsV1() eventsv1.EventsV1Interface
	EventsV1beta1() eventsv1beta1.EventsV1beta1Interface
	ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface
	FlowcontrolV1() flowcontrolv1.FlowcontrolV1Interface
	FlowcontrolV1beta1() flowcontrolv1beta1.FlowcontrolV1beta1Interface
	FlowcontrolV1beta2() flowcontrolv1beta2.FlowcontrolV1beta2Interface
	FlowcontrolV1beta3() flowcontrolv1beta3.FlowcontrolV1beta3Interface
	NetworkingV1() networkingv1.NetworkingV1Interface
	NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface
	NodeV1() nodev1.NodeV1Interface
	NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface
	NodeV1beta1() nodev1beta1.NodeV1beta1Interface
	PolicyV1() policyv1.PolicyV1Interface
	PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface
	RbacV1() rbacv1.RbacV1Interface
	RbacV1beta1() rbacv1beta1.RbacV1beta1Interface
	RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface
	ResourceV1() resourcev1.ResourceV1Interface
	ResourceV1beta2() resourcev1beta2.ResourceV1beta2Interface
	ResourceV1beta1() resourcev1beta1.ResourceV1beta1Interface
	ResourceV1alpha3() resourcev1alpha3.ResourceV1alpha3Interface
	SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface
	SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface
	SchedulingV1() schedulingv1.SchedulingV1Interface
	StorageV1beta1() storagev1beta1.StorageV1beta1Interface
	StorageV1() storagev1.StorageV1Interface
	StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface
	StoragemigrationV1alpha1() storagemigrationv1alpha1.StoragemigrationV1alpha1Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	admissionregistrationV1       *admissionregistrationv1.AdmissionregistrationV1Client
	admissionregistrationV1alpha1 *admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Client
	admissionregistrationV1beta1  *admissionregistrationv1beta1.AdmissionregistrationV1beta1Client
	internalV1alpha1              *internalv1alpha1.InternalV1alpha1Client
	appsV1                        *appsv1.AppsV1Client
	appsV1beta1                   *appsv1beta1.AppsV1beta1Client
	appsV1beta2                   *appsv1beta2.AppsV1beta2Client
	authenticationV1              *authenticationv1.AuthenticationV1Client
	authenticationV1alpha1        *authenticationv1alpha1.AuthenticationV1alpha1Client
	authenticationV1beta1         *authenticationv1beta1.AuthenticationV1beta1Client
	authorizationV1               *authorizationv1.AuthorizationV1Client
	authorizationV1beta1          *authorizationv1beta1.AuthorizationV1beta1Client
	autoscalingV1                 *autoscalingv1.AutoscalingV1Client
	autoscalingV2                 *autoscalingv2.AutoscalingV2Client
	autoscalingV2beta1            *autoscalingv2beta1.AutoscalingV2beta1Client
	autoscalingV2beta2            *autoscalingv2beta2.AutoscalingV2beta2Client
	batchV1                       *batchv1.BatchV1Client
	batchV1beta1                  *batchv1beta1.BatchV1beta1Client
	certificatesV1                *certificatesv1.CertificatesV1Client
	certificatesV1beta1           *certificatesv1beta1.CertificatesV1beta1Client
	certificatesV1alpha1          *certificatesv1alpha1.CertificatesV1alpha1Client
	coordinationV1alpha2          *coordinationv1alpha2.CoordinationV1alpha2Client
	coordinationV1beta1           *coordinationv1beta1.CoordinationV1beta1Client
	coordinationV1                *coordinationv1.CoordinationV1Client
	coreV1                        *corev1.CoreV1Client
	discoveryV1                   *discoveryv1.DiscoveryV1Client
	discoveryV1beta1              *discoveryv1beta1.DiscoveryV1beta1Client
	eventsV1                      *eventsv1.EventsV1Client
	eventsV1beta1                 *eventsv1beta1.EventsV1beta1Client
	extensionsV1beta1             *extensionsv1beta1.ExtensionsV1beta1Client
	flowcontrolV1                 *flowcontrolv1.FlowcontrolV1Client
	flowcontrolV1beta1            *flowcontrolv1beta1.FlowcontrolV1beta1Client
	flowcontrolV1beta2            *flowcontrolv1beta2.FlowcontrolV1beta2Client
	flowcontrolV1beta3            *flowcontrolv1beta3.FlowcontrolV1beta3Client
	networkingV1                  *networkingv1.NetworkingV1Client
	networkingV1beta1             *networkingv1beta1.NetworkingV1beta1Client
	nodeV1                        *nodev1.NodeV1Client
	nodeV1alpha1                  *nodev1alpha1.NodeV1alpha1Client
	nodeV1beta1                   *nodev1beta1.NodeV1beta1Client
	policyV1                      *policyv1.PolicyV1Client
	policyV1beta1                 *policyv1beta1.PolicyV1beta1Client
	rbacV1                        *rbacv1.RbacV1Client
	rbacV1beta1                   *rbacv1beta1.RbacV1beta1Client
	rbacV1alpha1                  *rbacv1alpha1.RbacV1alpha1Client
	resourceV1                    *resourcev1.ResourceV1Client
	resourceV1beta2               *resourcev1beta2.ResourceV1beta2Client
	resourceV1beta1               *resourcev1beta1.ResourceV1beta1Client
	resourceV1alpha3              *resourcev1alpha3.ResourceV1alpha3Client
	schedulingV1alpha1            *schedulingv1alpha1.SchedulingV1alpha1Client
	schedulingV1beta1             *schedulingv1beta1.SchedulingV1beta1Client
	schedulingV1                  *schedulingv1.SchedulingV1Client
	storageV1beta1                *storagev1beta1.StorageV1beta1Client
	storageV1                     *storagev1.StorageV1Client
	storageV1alpha1               *storagev1alpha1.StorageV1alpha1Client
	storagemigrationV1alpha1      *storagemigrationv1alpha1.StoragemigrationV1alpha1Client
}

// AdmissionregistrationV1 retrieves the AdmissionregistrationV1Client
func (c *Clientset) AdmissionregistrationV1() admissionregistrationv1.AdmissionregistrationV1Interface {
	return c.admissionregistrationV1
}

// AdmissionregistrationV1alpha1 retrieves the AdmissionregistrationV1alpha1Client
func (c *Clientset) AdmissionregistrationV1alpha1() admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Interface {
	return c.admissionregistrationV1alpha1
}

// AdmissionregistrationV1beta1 retrieves the AdmissionregistrationV1beta1Client
func (c *Clientset) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	return c.admissionregistrationV1beta1
}

// InternalV1alpha1 retrieves the InternalV1alpha1Client
func (c *Clientset) InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface {
	return c.internalV1alpha1
}

// AppsV1 retrieves the AppsV1Client
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return c.appsV1
}

// AppsV1beta1 retrieves the AppsV1beta1Client
func (c *Clientset) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	return c.appsV1beta1
}

// AppsV1beta2 retrieves the AppsV1beta2Client
func (c *Clientset) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	return c.appsV1beta2
}

// AuthenticationV1 retrieves the AuthenticationV1Client
func (c *Clientset) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	return c.authenticationV1
}

// AuthenticationV1alpha1 retrieves the AuthenticationV1alpha1Client
func (c *Clientset) AuthenticationV1alpha1() authenticationv1alpha1.AuthenticationV1alpha1Interface {
	return c.authenticationV1alpha1
}

// AuthenticationV1beta1 retrieves the AuthenticationV1beta1Client
func (c *Clientset) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface {
	return c.authenticationV1beta1
}

// AuthorizationV1 retrieves the AuthorizationV1Client
func (c *Clientset) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	return c.authorizationV1
}

// AuthorizationV1beta1 retrieves the AuthorizationV1beta1Client
func (c *Clientset) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface {
	return c.authorizationV1beta1
}

// AutoscalingV1 retrieves the AutoscalingV1Client
func (c *Clientset) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	return c.autoscalingV1
}

// AutoscalingV2 retrieves the AutoscalingV2Client
func (c *Clientset) AutoscalingV2() autoscalingv2.AutoscalingV2Interface {
	return c.autoscalingV2
}

// AutoscalingV2beta1 retrieves the AutoscalingV2beta1Client
func (c *Clientset) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	return c.autoscalingV2beta1
}

// AutoscalingV2beta2 retrieves the AutoscalingV2beta2Client
func (c *Clientset) AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface {
	return c.autoscalingV2beta2
}

// BatchV1 retrieves the BatchV1Client
func (c *Clientset) BatchV1() batchv1.BatchV1Interface {
	return c.batchV1
}

// BatchV1beta1 retrieves the BatchV1beta1Client
func (c *Clientset) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	return c.batchV1beta1
}

// CertificatesV1 retrieves the CertificatesV1Client
func (c *Clientset) CertificatesV1() certificatesv1.CertificatesV1Interface {
	return c.certificatesV1
}

// CertificatesV1beta1 retrieves the CertificatesV1beta1Client
func (c *Clientset) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface {
	return c.certificatesV1beta1
}

// CertificatesV1alpha1 retrieves the CertificatesV1alpha1Client
func (c *Clientset) CertificatesV1alpha1() certificatesv1alpha1.CertificatesV1alpha1Interface {
	return c.certificatesV1alpha1
}

// CoordinationV1alpha2 retrieves the CoordinationV1alpha2Client
func (c *Clientset) CoordinationV1alpha2() coordinationv1alpha2.CoordinationV1alpha2Interface {
	return c.coordinationV1alpha2
}

// CoordinationV1beta1 retrieves the CoordinationV1beta1Client
func (c *Clientset) CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface {
	return c.coordinationV1beta1
}

// CoordinationV1 retrieves the CoordinationV1Client
func (c *Clientset) CoordinationV1() coordinationv1.CoordinationV1Interface {
	return c.coordinationV1
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return c.coreV1
}

// DiscoveryV1 retrieves the DiscoveryV1Client
func (c *Clientset) DiscoveryV1() discoveryv1.DiscoveryV1Interface {
	return c.discoveryV1
}

// DiscoveryV1beta1 retrieves the DiscoveryV1beta1Client
func (c *Clientset) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface {
	return c.discoveryV1beta1
}

// EventsV1 retrieves the EventsV1Client
func (c *Clientset) EventsV1() eventsv1.EventsV1Interface {
	return c.eventsV1
}

// EventsV1beta1 retrieves the EventsV1beta1Client
func (c *Clientset) EventsV1beta1() eventsv1beta1.EventsV1beta1Interface {
	return c.eventsV1beta1
}

// ExtensionsV1beta1 retrieves the ExtensionsV1beta1Client
func (c *Clientset) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	return c.extensionsV1beta1
}

// FlowcontrolV1 retrieves the FlowcontrolV1Client
func (c *Clientset) FlowcontrolV1() flowcontrolv1.FlowcontrolV1Interface {
	return c.flowcontrolV1
}

// FlowcontrolV1beta1 retrieves the FlowcontrolV1beta1Client
func (c *Clientset) FlowcontrolV1beta1() flowcontrolv1beta1.FlowcontrolV1beta1Interface {
	return c.flowcontrolV1beta1
}

// FlowcontrolV1beta2 retrieves the FlowcontrolV1beta2Client
func (c *Clientset) FlowcontrolV1beta2() flowcontrolv1beta2.FlowcontrolV1beta2Interface {
	return c.flowcontrolV1beta2
}

// FlowcontrolV1beta3 retrieves the FlowcontrolV1beta3Client
func (c *Clientset) FlowcontrolV1beta3() flowcontrolv1beta3.FlowcontrolV1beta3Interface {
	return c.flowcontrolV1beta3
}

// NetworkingV1 retrieves the NetworkingV1Client
func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	return c.networkingV1
}

// NetworkingV1beta1 retrieves the NetworkingV1beta1Client
func (c *Clientset) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface {
	return c.networkingV1beta1
}

// NodeV1 retrieves the NodeV1Client
func (c *Clientset) NodeV1() nodev1.NodeV1Interface {
	return c.nodeV1
}

// NodeV1alpha1 retrieves the NodeV1alpha1Client
func (c *Clientset) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	return c.nodeV1alpha1
}

// NodeV1beta1 retrieves the NodeV1beta1Client
func (c *Clientset) NodeV1beta1() nodev1beta1.NodeV1beta1Interface {
	return c.nodeV1beta1
}

// PolicyV1 retrieves the PolicyV1Client
func (c *Clientset) PolicyV1() policyv1.PolicyV1Interface {
	return c.policyV1
}

// PolicyV1beta1 retrieves the PolicyV1beta1Client
func (c *Clientset) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	return c.policyV1beta1
}

// RbacV1 retrieves the RbacV1Client
func (c *Clientset) RbacV1() rbacv1.RbacV1Interface {
	return c.rbacV1
}

// RbacV1beta1 retrieves the RbacV1beta1Client
func (c *Clientset) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	return c.rbacV1beta1
}

// RbacV1alpha1 retrieves the RbacV1alpha1Client
func (c *Clientset) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	return c.rbacV1alpha1
}

// ResourceV1 retrieves the ResourceV1Client
func (c *Clientset) ResourceV1() resourcev1.ResourceV1Interface {
	return c.resourceV1
}

// ResourceV1beta2 retrieves the ResourceV1beta2Client
func (c *Clientset) ResourceV1beta2() resourcev1beta2.ResourceV1beta2Interface {
	return c.resourceV1beta2
}

// ResourceV1beta1 retrieves the ResourceV1beta1Client
func (c *Clientset) ResourceV1beta1() resourcev1beta1.ResourceV1beta1Interface {
	return c.resourceV1beta1
}

// ResourceV1alpha3 retrieves the ResourceV1alpha3Client
func (c *Clientset) ResourceV1alpha3() resourcev1alpha3.ResourceV1alpha3Interface {
	return c.resourceV1alpha3
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1Client
func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
}

// SchedulingV1beta1 retrieves the SchedulingV1beta1Client
func (c *Clientset) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	return c.schedulingV1beta1
}

// SchedulingV1 retrieves the SchedulingV1Client
func (c *Clientset) SchedulingV1() schedulingv1.SchedulingV1Interface {
	return c.schedulingV1
}

// StorageV1beta1 retrieves the StorageV1beta1Client
func (c *Clientset) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	return c.storageV1beta1
}

// StorageV1 retrieves the StorageV1Client
func (c *Clientset) StorageV1() storagev1.StorageV1Interface {
	return c.storageV1
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return c.storageV1alpha1
}

// StoragemigrationV1alpha1 retrieves the StoragemigrationV1alpha1Client
func (c *Clientset) StoragemigrationV1alpha1() storagemigrationv1alpha1.StoragemigrationV1alpha1Interface {
	return c.storagemigrationV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.admissionregistrationV1, err = admissionregistrationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.admissionregistrationV1alpha1, err = admissionregistrationv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.admissionregistrationV1beta1, err = admissionregistrationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.internalV1alpha1, err = internalv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.appsV1, err = appsv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta1, err = appsv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta2, err = appsv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1, err = authenticationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1alpha1, err = authenticationv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1beta1, err = authenticationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.authorizationV1, err = authorizationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.authorizationV1beta1, err = authorizationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV1, err = autoscalingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2, err = autoscalingv2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2beta1, err = autoscalingv2beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2beta2, err = autoscalingv2beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.batchV1, err = batchv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.batchV1beta1, err = batchv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.certificatesV1, err = certificatesv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.certificatesV1beta1, err = certificatesv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.certificatesV1alpha1, err = certificatesv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.coordinationV1alpha2, err = coordinationv1alpha2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.coordinationV1beta1, err = coordinationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.coordinationV1, err = coordinationv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.coreV1, err = corev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.discoveryV1, err = discoveryv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.discoveryV1beta1, err = discoveryv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.eventsV1, err = eventsv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.eventsV1beta1, err = eventsv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.extensionsV1beta1, err = extensionsv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.flowcontrolV1, err = flowcontrolv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.flowcontrolV1beta1, err = flowcontrolv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.flowcontrolV1beta2, err = flowcontrolv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.flowcontrolV1beta3, err = flowcontrolv1beta3.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.networkingV1, err = networkingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.networkingV1beta1, err = networkingv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.nodeV1, err = nodev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.nodeV1alpha1, err = nodev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.nodeV1beta1, err = nodev1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.policyV1, err = policyv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.policyV1beta1, err = policyv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.rbacV1, err = rbacv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.rbacV1beta1, err = rbacv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.rbacV1alpha1, err = rbacv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.resourceV1, err = resourcev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.resourceV1beta2, err = resourcev1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.resourceV1beta1, err = resourcev1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.resourceV1alpha3, err = resourcev1alpha3.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1alpha1, err = schedulingv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1beta1, err = schedulingv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1, err = schedulingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.storageV1beta1, err = storagev1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.storageV1, err = storagev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.storageV1alpha1, err = storagev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.storagemigrationV1alpha1, err = storagemigrationv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.admissionregistrationV1 = admissionregistrationv1.New(c)
	cs.admissionregistrationV1alpha1 = admissionregistrationv1alpha1.New(c)
	cs.admissionregistrationV1beta1 = admissionregistrationv1beta1.New(c)
	cs.internalV1alpha1 = internalv1alpha1.New(c)
	cs.appsV1 = appsv1.New(c)
	cs.appsV1beta1 = appsv1beta1.New(c)
	cs.appsV1beta2 = appsv1beta2.New(c)
	cs.authenticationV1 = authenticationv1.New(c)
	cs.authenticationV1alpha1 = authenticationv1alpha1.New(c)
	cs.authenticationV1beta1 = authenticationv1beta1.New(c)
	cs.authorizationV1 = authorizationv1.New(c)
	cs.authorizationV1beta1 = authorizationv1beta1.New(c)
	cs.autoscalingV1 = autoscalingv1.New(c)
	cs.autoscalingV2 = autoscalingv2.New(c)
	cs.autoscalingV2beta1 = autoscalingv2beta1.New(c)
	cs.autoscalingV2beta2 = autoscalingv2beta2.New(c)
	cs.batchV1 = batchv1.New(c)
	cs.batchV1beta1 = batchv1beta1.New(c)
	cs.certificatesV1 = certificatesv1.New(c)
	cs.certificatesV1beta1 = certificatesv1beta1.New(c)
	cs.certificatesV1alpha1 = certificatesv1alpha1.New(c)
	cs.coordinationV1alpha2 = coordinationv1alpha2.New(c)
	cs.coordinationV1beta1 = coordinationv1beta1.New(c)
	cs.coordinationV1 = coordinationv1.New(c)
	cs.coreV1 = corev1.New(c)
	cs.discoveryV1 = discoveryv1.New(c)
	cs.discoveryV1beta1 = discoveryv1beta1.New(c)
	cs.eventsV1 = eventsv1.New(c)
	cs.eventsV1beta1 = eventsv1beta1.New(c)
	cs.extensionsV1beta1 = extensionsv1beta1.New(c)
	cs.flowcontrolV1 = flowcontrolv1.New(c)
	cs.flowcontrolV1beta1 = flowcontrolv1beta1.New(c)
	cs.flowcontrolV1beta2 = flowcontrolv1beta2.New(c)
	cs.flowcontrolV1beta3 = flowcontrolv1beta3.New(c)
	cs.networkingV1 = networkingv1.New(c)
	cs.networkingV1beta1 = networkingv1beta1.New(c)
	cs.nodeV1 = nodev1.New(c)
	cs.nodeV1alpha1 = nodev1alpha1.New(c)
	cs.nodeV1beta1 = nodev1beta1.New(c)
	cs.policyV1 = policyv1.New(c)
	cs.policyV1beta1 = policyv1beta1.New(c)
	cs.rbacV1 = rbacv1.New(c)
	cs.rbacV1beta1 = rbacv1beta1.New(c)
	cs.rbacV1alpha1 = rbacv1alpha1.New(c)
	cs.resourceV1 = resourcev1.New(c)
	cs.resourceV1beta2 = resourcev1beta2.New(c)
	cs.resourceV1beta1 = resourcev1beta1.New(c)
	cs.resourceV1alpha3 = resourcev1alpha3.New(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.New(c)
	cs.schedulingV1beta1 = schedulingv1beta1.New(c)
	cs.schedulingV1 = schedulingv1.New(c)
	cs.storageV1beta1 = storagev1beta1.New(c)
	cs.storageV1 = storagev1.New(c)
	cs.storageV1alpha1 = storagev1alpha1.New(c)
	cs.storagemigrationV1alpha1 = storagemigrationv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
