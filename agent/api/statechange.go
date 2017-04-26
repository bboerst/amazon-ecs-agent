// Copyright 2014-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package api

import (
	"fmt"
	"strconv"
)

// ContainerStateChange represents a state change that needs to be sent to the
// SubmitContainerStateChange API
type ContainerStateChange struct {
	// TaskArn is the unique identifier for the task
	TaskArn string
	// ContainerName is the name of the container
	ContainerName string
	// Status is the status to send
	Status ContainerStatus

	// Reason may contain details of why the container stopped
	Reason string
	// ExitCode is the exit code of the container, if available
	ExitCode *int
	// PortBindings are the details of the host ports picked for the specified
	// container ports
	PortBindings []PortBinding

	// Container is a pointer to the container involved in the state change that gives the event handler a hook into
	// storing what status was sent.  This is used to ensure the same event is handled only once.
	Container *Container
}

// TaskStateChange represents a state change that needs to be sent to the
// SubmitTaskStateChange API
type TaskStateChange struct {
	// TaskArn is the unique identifier for the task
	TaskArn string
	// Status is the status to send
	Status TaskStatus
	// Reason may contain details of why the task stopped
	Reason string

	// Task is a pointer to the task involved in the state change that gives the event handler a hook into storing
	// what status was sent.  This is used to ensure the same event is handled only once.
	Task *Task
}

// ENIAttachmentStateChange represents a state change that needs to be sent to
// SubmitTaskStateChange API to report the ENI attachment/detachment
type ENIAttachmentStateChange struct {
	TaskArn       string
	AttachmentArn string
	Status        string
	Reason        string
}

// String returns a human readable string representation of this object
func (c *ContainerStateChange) String() string {
	res := fmt.Sprintf("%s %s -> %s", c.TaskArn, c.ContainerName, c.Status.String())
	if c.ExitCode != nil {
		res += ", Exit " + strconv.Itoa(*c.ExitCode) + ", "
	}
	if c.Reason != "" {
		res += ", Reason " + c.Reason
	}
	if len(c.PortBindings) != 0 {
		res += fmt.Sprintf(", Ports %v", c.PortBindings)
	}
	if c.Container != nil {
		res += ", Known Sent: " + c.Container.GetSentStatus().String()
	}
	return res
}

// String returns a human readable string representation of this object
func (t *TaskStateChange) String() string {
	res := fmt.Sprintf("%s -> %s", t.TaskArn, t.Status.String())
	if t.Task != nil {
		res += ", Known Sent: " + t.Task.GetSentStatus().String()
	}
	return res
}
