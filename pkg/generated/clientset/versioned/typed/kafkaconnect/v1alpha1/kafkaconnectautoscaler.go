// Copyright 2020 The Amadeus Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/amadeusitgroup/kubernetes-kafka-connect-operator/pkg/apis/kafkaconnect/v1alpha1"
	scheme "github.com/amadeusitgroup/kubernetes-kafka-connect-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KafkaConnectAutoScalersGetter has a method to return a KafkaConnectAutoScalerInterface.
// A group's client should implement this interface.
type KafkaConnectAutoScalersGetter interface {
	KafkaConnectAutoScalers(namespace string) KafkaConnectAutoScalerInterface
}

// KafkaConnectAutoScalerInterface has methods to work with KafkaConnectAutoScaler resources.
type KafkaConnectAutoScalerInterface interface {
	Create(*v1alpha1.KafkaConnectAutoScaler) (*v1alpha1.KafkaConnectAutoScaler, error)
	Update(*v1alpha1.KafkaConnectAutoScaler) (*v1alpha1.KafkaConnectAutoScaler, error)
	UpdateStatus(*v1alpha1.KafkaConnectAutoScaler) (*v1alpha1.KafkaConnectAutoScaler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KafkaConnectAutoScaler, error)
	List(opts v1.ListOptions) (*v1alpha1.KafkaConnectAutoScalerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KafkaConnectAutoScaler, err error)
	KafkaConnectAutoScalerExpansion
}

// kafkaConnectAutoScalers implements KafkaConnectAutoScalerInterface
type kafkaConnectAutoScalers struct {
	client rest.Interface
	ns     string
}

// newKafkaConnectAutoScalers returns a KafkaConnectAutoScalers
func newKafkaConnectAutoScalers(c *KafkaconnectV1alpha1Client, namespace string) *kafkaConnectAutoScalers {
	return &kafkaConnectAutoScalers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kafkaConnectAutoScaler, and returns the corresponding kafkaConnectAutoScaler object, and an error if there is any.
func (c *kafkaConnectAutoScalers) Get(name string, options v1.GetOptions) (result *v1alpha1.KafkaConnectAutoScaler, err error) {
	result = &v1alpha1.KafkaConnectAutoScaler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkaconnectautoscalers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KafkaConnectAutoScalers that match those selectors.
func (c *kafkaConnectAutoScalers) List(opts v1.ListOptions) (result *v1alpha1.KafkaConnectAutoScalerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KafkaConnectAutoScalerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkaconnectautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kafkaConnectAutoScalers.
func (c *kafkaConnectAutoScalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kafkaconnectautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kafkaConnectAutoScaler and creates it.  Returns the server's representation of the kafkaConnectAutoScaler, and an error, if there is any.
func (c *kafkaConnectAutoScalers) Create(kafkaConnectAutoScaler *v1alpha1.KafkaConnectAutoScaler) (result *v1alpha1.KafkaConnectAutoScaler, err error) {
	result = &v1alpha1.KafkaConnectAutoScaler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kafkaconnectautoscalers").
		Body(kafkaConnectAutoScaler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kafkaConnectAutoScaler and updates it. Returns the server's representation of the kafkaConnectAutoScaler, and an error, if there is any.
func (c *kafkaConnectAutoScalers) Update(kafkaConnectAutoScaler *v1alpha1.KafkaConnectAutoScaler) (result *v1alpha1.KafkaConnectAutoScaler, err error) {
	result = &v1alpha1.KafkaConnectAutoScaler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkaconnectautoscalers").
		Name(kafkaConnectAutoScaler.Name).
		Body(kafkaConnectAutoScaler).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kafkaConnectAutoScalers) UpdateStatus(kafkaConnectAutoScaler *v1alpha1.KafkaConnectAutoScaler) (result *v1alpha1.KafkaConnectAutoScaler, err error) {
	result = &v1alpha1.KafkaConnectAutoScaler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkaconnectautoscalers").
		Name(kafkaConnectAutoScaler.Name).
		SubResource("status").
		Body(kafkaConnectAutoScaler).
		Do().
		Into(result)
	return
}

// Delete takes name of the kafkaConnectAutoScaler and deletes it. Returns an error if one occurs.
func (c *kafkaConnectAutoScalers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkaconnectautoscalers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kafkaConnectAutoScalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkaconnectautoscalers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kafkaConnectAutoScaler.
func (c *kafkaConnectAutoScalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KafkaConnectAutoScaler, err error) {
	result = &v1alpha1.KafkaConnectAutoScaler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kafkaconnectautoscalers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
