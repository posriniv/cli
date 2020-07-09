/*
Package actions includes actions for create a new project via the Create Go App CLI.

Copyright © 2020 Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

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
package actions

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/create-go-app/cli/pkg/registry"
	"github.com/create-go-app/cli/pkg/utils"
)

// CreateProjectFromRegistry function for create a new project from repository.
func CreateProjectFromRegistry(p *registry.Project, r map[string]*registry.Repository) {
	// Define vars.
	var pattern string

	// Checking for nil.
	if p == nil || r == nil {
		utils.SendMsg(true, "[ERROR]", "Project template or registry not found!", "red", true)
		os.Exit(1)
	}

	// Create path in project root folder.
	folder := filepath.Join(p.RootFolder, p.Type)

	// Switch project type.
	switch p.Type {
	case "roles":
		pattern = registry.RegexpAnsiblePattern
		folder = filepath.Join(p.RootFolder, p.Type, p.Name) // re-define folder
		break
	case "backend":
		pattern = registry.RegexpBackendPattern
		break
	case "webserver":
		pattern = registry.RegexpWebServerPattern
		break
	case "database":
		pattern = registry.RegexpDatabasePattern
		break
	}

	// Create match expration.
	match, err := regexp.MatchString(pattern, p.Name)
	if err != nil {
		utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Check for regexp.
	if match {
		// Re-define vars.
		template := r[p.Type].List[p.Name]

		// If match, create from default template.
		if err := utils.GitClone(folder, template); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// Show success report.
		utils.SendMsg(false, "[OK]", strings.Title(p.Type)+": created with default template `"+template+"`!", "cyan", false)
	} else {
		// Else create from user template (from GitHub, etc).
		if err := utils.GitClone(folder, p.Name); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// Show success report.
		utils.SendMsg(false, "[OK]", strings.Title(p.Type)+": created with user template `"+p.Name+"`!", "cyan", false)
	}

	// Cleanup project.
	foldersToRemove := []string{".git", ".github"}
	utils.RemoveFolders(folder, foldersToRemove)
}