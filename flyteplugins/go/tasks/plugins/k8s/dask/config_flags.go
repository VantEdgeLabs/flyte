// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package dask

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(raw)
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.cloudwatch-enabled"), defaultConfig.Logs.IsCloudwatchEnabled, "Enable Cloudwatch Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.cloudwatch-region"), defaultConfig.Logs.CloudwatchRegion, "AWS region in which Cloudwatch logs are stored.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.cloudwatch-log-group"), defaultConfig.Logs.CloudwatchLogGroup, "Log group to which streams are associated.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.cloudwatch-template-uri"), defaultConfig.Logs.CloudwatchTemplateURI, "Template Uri to use when building cloudwatch log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.kubernetes-enabled"), defaultConfig.Logs.IsKubernetesEnabled, "Enable Kubernetes Logging")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.kubernetes-url"), defaultConfig.Logs.KubernetesURL, "Console URL for Kubernetes logs")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.kubernetes-template-uri"), defaultConfig.Logs.KubernetesTemplateURI, "Template Uri to use when building kubernetes log links")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "logs.stackdriver-enabled"), defaultConfig.Logs.IsStackDriverEnabled, "Enable Log-links to stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.gcp-project"), defaultConfig.Logs.GCPProjectName, "Name of the project in GCP")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.stackdriver-logresourcename"), defaultConfig.Logs.StackdriverLogResourceName, "Name of the logresource in stackdriver")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "logs.stackdriver-template-uri"), defaultConfig.Logs.StackDriverTemplateURI, "Template Uri to use when building stackdriver log links")
	return cmdFlags
}
