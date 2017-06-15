// Copyright 2016-2017 Authors of Cilium
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

package nodeaddress

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

type NodeAddressSuite struct{}

var _ = Suite(&NodeAddressSuite{})

func (s *NodeAddressSuite) TestNewNodeAddress(c *C) {
	err := SetNodeAddress("b007::aaaa:bbbb:0:0", "10.1.0.1", "")
	c.Assert(err, Equals, nil)
	c.Assert(IPv6Address.String(), Equals, "b007::aaaa:bbbb:0:0")
	c.Assert(IPv6Address.EndpointID(), Equals, uint16(0))
	c.Assert(IPv6Address.State(), Equals, uint16(0))
	c.Assert(IPv6Address.NodeID(), Equals, uint32(0xaaaabbbb))
	c.Assert(IPv4Address.EndpointID(), Equals, uint16(1))
	c.Assert(IPv4Address.NodeID(), Equals, uint32(0xa010000))

	err = SetNodeAddress("b007::", "20.2.0.1", "")
	c.Assert(err, Equals, nil)
	c.Assert(IPv6Address.EndpointID(), Equals, uint16(0))
	c.Assert(IPv6Address.State(), Equals, uint16(0))
	c.Assert(IPv6Address.NodeID(), Not(Equals), uint32(0))
	c.Assert(IPv4Address.EndpointID(), Equals, uint16(1))
	c.Assert(IPv4Address.NodeID(), Equals, uint32(0x14020000))

	// container bits set, should fail
	err = SetNodeAddress("b007::aaaa:bbbb:0:1", "10.1.0.0", "")
	c.Assert(err, Equals, ErrNodeIPEndpointIDSet)

	err = SetNodeAddress("b007::aaaa:bbbb:0:0", "0.0.0.0", "")
	c.Assert(err, Equals, ErrIPv4Invalid)

	err = SetNodeAddress("b007::aaaa:bbbb:0:1", "10.0.1.0", "")
	c.Assert(err, Equals, ErrNodeIPEndpointIDSet)
}
