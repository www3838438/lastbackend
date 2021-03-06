//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package v1

import (
	"github.com/lastbackend/lastbackend/pkg/api/container/views/v1"
	"time"
)

type Pod struct {
	// Pod Meta
	Meta PodMeta `json:"meta"`
	// Pod state
	State PodState `json:"state"`
	// Container spec
	Spec PodSpec `json:"spec"`
}

type PodInfo struct {
	// Pod Meta
	Meta PodMeta `json:"meta"`
	// Pod state
	State PodState `json:"state"`
	// Pod containers
	Containers []v1.Container `json:"containers"`
}

type PodMeta struct {
	// Meta id
	Name string `json:"name"`
	// Meta endpoint
	Endpoint string `json:"endpoint"`
	// Meta labels
	Labels map[string]string `json:"labels"`
	// Meta created time
	Created time.Time `json:"created"`
	// Meta updated time
	Updated time.Time `json:"updated"`
}

type PodState struct {
	// Pod current state
	State string `json:"state"`
	// Pod current status
	Status string `json:"status"`
	// Pod provision flag
	Provision bool `json:"provision"`
	// Pod ready flag
	Ready bool `json:"ready"`
}

type PodContainersState struct {
	// Total containers
	Total int `json:"total"`
	// Total running containers
	Running int `json:"running"`
	// Total created containers
	Created int `json:"created"`
	// Total stopped containers
	Stopped int `json:"stopped"`
	// Total errored containers
	Errored int `json:"errored"`
}

type PodSpec struct {
	// Provision ID
	ID string `json:"id"`
	// Provision state
	State string `json:"state"`
	// Provision status
	Status string `json:"status"`

	// Containers spec for pod
	Containers map[string]v1.ContainerSpec `json:"containers"`

	// Provision create time
	Created time.Time `json:"created"`
	// Provision update time
	Updated time.Time `json:"updated"`
}
