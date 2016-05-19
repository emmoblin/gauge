// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package run_failed

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getgauge/common"
	"github.com/getgauge/gauge/config"
	"github.com/getgauge/gauge/execution/event"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestIfFailedFileIsCreated(c *C) {
	p, _ := filepath.Abs("_testdata")
	config.ProjectRoot = p
	event.InitRegistry()

	ListenFailedScenarios()
	event.Notify(event.NewExecutionEvent(event.SuiteEnd, nil, nil, 0))

	c.Assert(common.FileExists(filepath.Join(config.ProjectRoot, dotGauge, failedFile)), Equals, true)
	os.RemoveAll(filepath.Join(config.ProjectRoot, dotGauge))
}
