package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/testprojects"
	"testing"
)

func TestProjectsFullTest(t *testing.T) {

	options := testprojects.TestProjectOptionsDefault(&testprojects.TestProjectsOptions{
		Testing:        t,
		Prefix:         "genai", // setting prefix here gets a random string appended to it
		ParallelDeploy: true,
	})

	options.StackInputs = map[string]interface{}{
		"resource_group_name":          options.ResourceGroup,
		"ibmcloud_api_key":             options.RequiredEnvironmentVars["TF_VAR_ibmcloud_api_key"],
		"prefix":                       options.Prefix,
		"secret_manager_service_plan":  "trial",
		"enable_platform_logs_metrics": false,
	}

	err := options.RunProjectsTest()
	if assert.NoError(t, err) {
		t.Log("TestProjectsFullTest Passed")
	} else {
		t.Error("TestProjectsFullTest Failed")
	}
}
