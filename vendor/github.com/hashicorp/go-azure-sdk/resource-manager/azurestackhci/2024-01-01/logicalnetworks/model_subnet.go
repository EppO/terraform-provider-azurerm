package logicalnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Subnet struct {
	Name       *string                 `json:"name,omitempty"`
	Properties *SubnetPropertiesFormat `json:"properties,omitempty"`
}
