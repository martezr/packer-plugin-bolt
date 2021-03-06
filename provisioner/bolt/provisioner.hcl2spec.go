// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package bolt

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName      *string                     `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType    *string                     `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion    *string                     `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug          *bool                       `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce          *bool                       `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError        *string                     `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars       map[string]string           `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars  []string                    `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Command              *string                     `cty:"command" hcl:"command"`
	Backend              *string                     `mapstructure:"backend" cty:"backend" hcl:"backend"`
	Host                 *string                     `mapstructure:"host" cty:"host" hcl:"host"`
	Password             *string                     `mapstructure:"password" cty:"password" hcl:"password"`
	ExtraArguments       []string                    `mapstructure:"extra_arguments" cty:"extra_arguments" hcl:"extra_arguments"`
	BoltEnvVars          []string                    `mapstructure:"bolt_env_vars" cty:"bolt_env_vars" hcl:"bolt_env_vars"`
	BoltParams           map[interface{}]interface{} `mapstructure:"bolt_params" cty:"bolt_params" hcl:"bolt_params"`
	BoltTask             *string                     `mapstructure:"bolt_task" cty:"bolt_task" hcl:"bolt_task"`
	BoltPlan             *string                     `mapstructure:"bolt_plan" cty:"bolt_plan" hcl:"bolt_plan"`
	BoltModulePath       *string                     `mapstructure:"bolt_module_path" cty:"bolt_module_path" hcl:"bolt_module_path"`
	ProjectPath          *string                     `mapstructure:"project_path" cty:"project_path" hcl:"project_path"`
	InventoryFile        *string                     `mapstructure:"inventory_file" cty:"inventory_file" hcl:"inventory_file"`
	ConnectTimeout       *int                        `mapstructure:"connect_timeout" cty:"connect_timeout" hcl:"connect_timeout"`
	InventoryDirectory   *string                     `mapstructure:"inventory_directory" cty:"inventory_directory" hcl:"inventory_directory"`
	LocalPort            *int                        `mapstructure:"local_port" cty:"local_port" hcl:"local_port"`
	SkipVersionCheck     *bool                       `mapstructure:"skip_version_check" cty:"skip_version_check" hcl:"skip_version_check"`
	RunAs                *string                     `mapstructure:"run_as" cty:"run_as" hcl:"run_as"`
	User                 *string                     `mapstructure:"user" cty:"user" hcl:"user"`
	SSHHostKeyFile       *string                     `mapstructure:"ssh_host_key_file" cty:"ssh_host_key_file" hcl:"ssh_host_key_file"`
	SSHAuthorizedKeyFile *string                     `mapstructure:"ssh_authorized_key_file" cty:"ssh_authorized_key_file" hcl:"ssh_authorized_key_file"`
	NoWinRMSSLVerify     *bool                       `mapstructure:"no_winrm_ssl_verify" cty:"no_winrm_ssl_verify" hcl:"no_winrm_ssl_verify"`
	NoWinRMSSL           *bool                       `mapstructure:"no_winrm_ssl" cty:"no_winrm_ssl" hcl:"no_winrm_ssl"`
	LogLevel             *string                     `mapstructure:"log_level" cty:"log_level" hcl:"log_level"`
	InstallModules       *bool                       `mapstructure:"install_modules" cty:"install_modules" hcl:"install_modules"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"command":                    &hcldec.AttrSpec{Name: "command", Type: cty.String, Required: false},
		"backend":                    &hcldec.AttrSpec{Name: "backend", Type: cty.String, Required: false},
		"host":                       &hcldec.AttrSpec{Name: "host", Type: cty.String, Required: false},
		"password":                   &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"extra_arguments":            &hcldec.AttrSpec{Name: "extra_arguments", Type: cty.List(cty.String), Required: false},
		"bolt_env_vars":              &hcldec.AttrSpec{Name: "bolt_env_vars", Type: cty.List(cty.String), Required: false},
		"bolt_params":                &hcldec.AttrSpec{Name: "bolt_params", Type: cty.Map(cty.String), Required: false},
		"bolt_task":                  &hcldec.AttrSpec{Name: "bolt_task", Type: cty.String, Required: false},
		"bolt_plan":                  &hcldec.AttrSpec{Name: "bolt_plan", Type: cty.String, Required: false},
		"bolt_module_path":           &hcldec.AttrSpec{Name: "bolt_module_path", Type: cty.String, Required: false},
		"project_path":               &hcldec.AttrSpec{Name: "project_path", Type: cty.String, Required: false},
		"inventory_file":             &hcldec.AttrSpec{Name: "inventory_file", Type: cty.String, Required: false},
		"connect_timeout":            &hcldec.AttrSpec{Name: "connect_timeout", Type: cty.Number, Required: false},
		"inventory_directory":        &hcldec.AttrSpec{Name: "inventory_directory", Type: cty.String, Required: false},
		"local_port":                 &hcldec.AttrSpec{Name: "local_port", Type: cty.Number, Required: false},
		"skip_version_check":         &hcldec.AttrSpec{Name: "skip_version_check", Type: cty.Bool, Required: false},
		"run_as":                     &hcldec.AttrSpec{Name: "run_as", Type: cty.String, Required: false},
		"user":                       &hcldec.AttrSpec{Name: "user", Type: cty.String, Required: false},
		"ssh_host_key_file":          &hcldec.AttrSpec{Name: "ssh_host_key_file", Type: cty.String, Required: false},
		"ssh_authorized_key_file":    &hcldec.AttrSpec{Name: "ssh_authorized_key_file", Type: cty.String, Required: false},
		"no_winrm_ssl_verify":        &hcldec.AttrSpec{Name: "no_winrm_ssl_verify", Type: cty.Bool, Required: false},
		"no_winrm_ssl":               &hcldec.AttrSpec{Name: "no_winrm_ssl", Type: cty.Bool, Required: false},
		"log_level":                  &hcldec.AttrSpec{Name: "log_level", Type: cty.String, Required: false},
		"install_modules":            &hcldec.AttrSpec{Name: "install_modules", Type: cty.Bool, Required: false},
	}
	return s
}
