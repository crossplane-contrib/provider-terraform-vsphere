/*
Copyright 2020 The Crossplane Authors.

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

package main

import (
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/crossplane-contrib/provider-terraform-vsphere/generated"
	"github.com/crossplane-contrib/terraform-runtime/pkg/client"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	"github.com/crossplane/crossplane-runtime/pkg/logging"

	"github.com/crossplane-contrib/terraform-runtime/pkg/controller"
)

func main() {
	var (
		app        = kingpin.New(filepath.Base(os.Args[0]), "Template support for Crossplane.").DefaultEnvars()
		debug      = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		// TODO: pluginPath should receive a default pointing to a canonical path in a container generated by the build process
		pluginPath = app.Flag("pluginPath", "Full path to directory where terraform plugin is located.").String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("terraform-runtime"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	providerInit := generated.ProviderInit()
	idxr := plugin.NewIndexer()
	err := generated.Index(idxr)
	if err != nil {
		kingpin.FatalIfError(err, "plugin.Indexer failed (likely issue with a resource Implementation overlay)")
	}
	idx, err := idxr.BuildIndex()
	kingpin.FatalIfError(err, "Failed to index provider plugin")

	opts := ctrl.Options{SyncPeriod: syncPeriod}
	ropts := client.NewRuntimeOptions().
		WithPluginPath(*pluginPath).
		WithPoolSize(5)
	log.Debug("Starting", "sync-period", syncPeriod.String())

	err = controller.StartTerraformManager(idx, providerInit, opts, ropts, log)
	kingpin.FatalIfError(err, "Cannot start the generated terraform controllers")
}
