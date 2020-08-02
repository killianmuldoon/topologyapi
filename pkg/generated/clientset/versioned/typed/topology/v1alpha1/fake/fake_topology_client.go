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

// This package imports things required by build scripts, to force `go mod` to see them as dependencies

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/swatisehgal/topologyapi/pkg/generated/clientset/versioned/typed/topology/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTopologyV1alpha1 struct {
	*testing.Fake
}

func (c *FakeTopologyV1alpha1) NodeResourceTopologies(namespace string) v1alpha1.NodeResourceTopologyInterface {
	return &FakeNodeResourceTopologies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTopologyV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
