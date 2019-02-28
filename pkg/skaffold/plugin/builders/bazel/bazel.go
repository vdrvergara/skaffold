/*
Copyright 2019 The Skaffold Authors

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

package bazel

import (
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/plugin/shared"
	plugin "github.com/hashicorp/go-plugin"
	"github.com/pkg/errors"
)

// Execute an image build with docker
func Execute() error {
	builder, err := NewBuilder()
	if err != nil {
		return errors.Wrap(err, "getting bazel builder")
	}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"bazel": &shared.BuilderPlugin{Impl: builder},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         pluginMap,
	})
	return nil
}
