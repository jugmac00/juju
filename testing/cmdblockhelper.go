// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package testing

import (
	"strings"

	"github.com/juju/cmd"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/api/base"
	"github.com/juju/juju/api/block"
)

// CmdBlockHelper is a helper struct used to block commands.
type CmdBlockHelper struct {
	blockClient *block.Client
}

// NewCmdBlockHelper creates a block switch used in testing
// to manage desired juju blocks.
func NewCmdBlockHelper(api base.APICallCloser) CmdBlockHelper {
	return CmdBlockHelper{
		blockClient: block.NewClient(api),
	}
}

// on switches on desired block and
// asserts that no errors were encountered.
func (s *CmdBlockHelper) on(c *gc.C, blockType, msg string) {
	c.Assert(s.blockClient.SwitchBlockOn(blockType, msg), gc.IsNil)
}

// BlockAllChanges switches changes block on.
// This prevents all changes to juju environment.
func (s *CmdBlockHelper) BlockAllChanges(c *gc.C, msg string) {
	s.on(c, "BlockChange", msg)
}

// BlockRemoveObject switches remove block on.
// This prevents any object/entity removal on juju environment
func (s *CmdBlockHelper) BlockRemoveObject(c *gc.C, msg string) {
	s.on(c, "BlockRemove", msg)
}

// BlockDestroyModel switches destroy block on.
// This prevents juju environment destruction.
func (s *CmdBlockHelper) BlockDestroyModel(c *gc.C, msg string) {
	s.on(c, "BlockDestroy", msg)
}

func (s *CmdBlockHelper) Close() {
	s.blockClient.Close()
}

func (s *CmdBlockHelper) AssertBlocked(c *gc.C, err error, msg string) {
	c.Assert(err, gc.ErrorMatches, cmd.ErrSilent.Error())
	// msg is logged
	stripped := strings.Replace(c.GetTestLog(), "\n", "", -1)
	c.Check(stripped, gc.Matches, msg)
}

func AssertOperationWasBlocked(c *gc.C, err error, msg string) {
	c.Assert(err, gc.ErrorMatches, cmd.ErrSilent.Error())
	// msg is logged
	stripped := strings.Replace(c.GetTestLog(), "\n", "", -1)
	c.Check(stripped, gc.Matches, msg)
}
