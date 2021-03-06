// Copyright 2017 Diego Fernández Barrera
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pipeline

import (
	"testing"

	"github.com/Bigomby/gopiper/component"
)

type TestComponent struct{}

func (c *TestComponent) Handle(m component.Message) *component.Report {
	return &component.Report{}
}

type TestFactory struct{}

func (f TestFactory) Create(postponed chan component.Message) component.Component {
	return &TestComponent{}
}

func (f TestFactory) SetAttribute(key string, value interface{}) error {
	return nil
}

func (f TestFactory) ChannelSize() int { return 1 }
func (f TestFactory) PoolSize() int    { return 1 }
func (f TestFactory) Destroy()         {}

func TestPipeline(t *testing.T) {
	factory := TestFactory{}

	factories := []component.Factory{factory}

	NewPipeline(factories)
}
