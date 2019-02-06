package manifest

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

/*
Application yaml represents the complete yaml structure
*/
type Application struct {
	Name                    string                 `yaml:"name"`
	Instances               int                    `yaml:"instances,omitempty"`
	Memory                  string                 `yaml:"memory,omitempty"`
	DiskQuota               string                 `yaml:"disk_quota,omitempty"`
	Routes                  []map[string]string    `yaml:"routes,omitempty"`
	NoRoute                 bool                   `yaml:"no-route,omitempty"`
	Buildpack               string                 `yaml:"buildpack,omitempty"`
	Command                 string                 `yaml:"command,omitempty"`
	Env                     map[string]interface{} `yaml:"env,omitempty"`
	Services                []string               `yaml:"services,omitempty"`
	Stack                   string                 `yaml:"stack,omitempty"`
	Timeout                 int                    `yaml:"timeout,omitempty"`
	HealthCheckType         string                 `yaml:"health-check-type,omitempty"`
	HealthCheckHTTPEndpoint string                 `yaml:"health-check-http-endpoint,omitempty"`
}

/*
Represents the manifest.
*/
type Manifest struct {
	Manifest []Application `yaml:"applications"`
}

func loadYmlFile(manifestFilePath string) (manifest Manifest, err error) {
	fileBytes, err := ioutil.ReadFile(manifestFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while reading manifest:", err)
		os.Exit(1)
	}

	var document Manifest
	err = yaml.Unmarshal(fileBytes, &document)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while parsing the manifest", err)
		os.Exit(1)
	}

	return document, nil
}
