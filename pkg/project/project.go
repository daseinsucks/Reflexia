package project

import (
	"errors"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	util "reflexia/internal"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type ProjectConfig struct {
	FileFilter        []string `toml:"file_filter"`
	ProjectRootFilter []string `toml:"project_root_filter"`
	ModuleMatch       string   `toml:"module_match"`
	StopWords         []string `toml:"stop_words"`
	CodePrompt        string   `toml:"code_prompt"`
	SummaryPrompt     string   `toml:"summary_prompt"`
	PackagePrompt     string   `toml:"package_prompt"`
	ReadmePrompt      string   `toml:"readme_prompt"`
	RootPath          string
}

func GetProjectConfig(currentDirectory string, lightCheck bool) (*ProjectConfig, error) {
	var projectConfigs = []ProjectConfig{}

	if err := util.WalkDirIgnored(
		"project_config", filepath.Join(currentDirectory, ".gitignore"),
		func(path string, d fs.DirEntry) error {
			if d.IsDir() {
				return nil
			}
			if strings.HasSuffix(path, ".toml") {
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				var projectConfig ProjectConfig
				if err = toml.Unmarshal(content, &projectConfig); err != nil {
					return err
				}
				projectConfig.RootPath = currentDirectory
				projectConfigs = append(projectConfigs, projectConfig)
			}
			return nil
		}); err != nil {
		return nil, err
	}

	for _, config := range projectConfigs {
		if lightCheck && hasFilterFiles(currentDirectory, config.FileFilter) {
			return &config, nil
		}
		if hasFilterFiles(currentDirectory, config.FileFilter) &&
			hasRootFilterFile(currentDirectory, config.ProjectRootFilter) {
			return &config, nil
		}
	}

	return nil, errors.New(
		"failed to detect project language, available languages: go, python, typescript",
	)
}

func (pc *ProjectConfig) BuildPackageFiles() (map[string][]string, error) {
	packageFileMap := map[string][]string{}
	switch pc.ModuleMatch {
	case "directory":
		if err := util.WalkDirIgnored(
			pc.RootPath,
			filepath.Join(pc.RootPath, ".gitignore"),
			func(path string, d fs.DirEntry) error {
				for _, filter := range pc.FileFilter {
					if strings.HasSuffix(d.Name(), filter) {
						key := filepath.Base(filepath.Dir(path))
						relPath, err := filepath.Rel(pc.RootPath, path)
						if err != nil {
							return err
						}

						if _, exists := packageFileMap[key]; !exists {
							packageFileMap[key] = []string{}
						}
						packageFileMap[key] = append(packageFileMap[key], relPath)
					}
				}

				return nil
			}); err != nil {
			return nil, err
		}

	case "package_name":
		if err := util.WalkDirIgnored(
			pc.RootPath,
			filepath.Join(pc.RootPath, ".gitignore"),
			func(path string, d fs.DirEntry) error {
				for _, filter := range pc.FileFilter {
					if strings.HasSuffix(d.Name(), filter) {
						fset := token.NewFileSet()
						ast, err := parser.ParseFile(fset, path, nil, 0)
						if err != nil {
							return err
						}
						key := ast.Name.Name
						relPath, err := filepath.Rel(pc.RootPath, path)
						if err != nil {
							return err
						}

						if _, exists := packageFileMap[key]; !exists {
							packageFileMap[key] = []string{}
						}
						packageFileMap[key] = append(packageFileMap[key], relPath)
					}
				}

				return nil
			}); err != nil {
			return nil, err
		}

	default:
		return nil, errors.New(pc.ModuleMatch + " module match mode unimplemented")
	}

	return packageFileMap, nil
}

func hasFilterFiles(workdir string, filters []string) bool {
	found := false

	err := util.WalkDirIgnored(
		workdir, filepath.Join(workdir, ".gitignore"),
		func(path string, d fs.DirEntry) error {
			for _, filter := range filters {
				if strings.HasSuffix(d.Name(), filter) {
					found = true
					return io.EOF
				}
			}
			return nil
		})
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		log.Fatal(err)
	}
	return found
}

func hasRootFilterFile(workdir string, filters []string) bool {
	for _, filter := range filters {
		if _, err := os.Stat(filepath.Join(workdir, filter)); err != nil {
			if !os.IsNotExist(err) {
				panic(err)
			}
			return false
		}
	}

	return true
}
