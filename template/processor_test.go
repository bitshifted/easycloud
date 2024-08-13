package template

import (
	"os"
	"path"
	"testing"

	"github.com/bitshifted/liftoff/common"
	"github.com/bitshifted/liftoff/config"
	"github.com/bitshifted/liftoff/log"
	"github.com/stretchr/testify/assert"
)

func TestProcessTerraformFiles(t *testing.T) {
	log.Init(true)
	tmpDir, err := os.MkdirTemp("", "template-test")
	assert.NoError(t, err)
	processor := TemplateProcessor{
		BaseDir:   "test_files",
		OutputDir: tmpDir,
	}
	t.Setenv("HCLOUD_TOKEN", "foo")
	config, err := config.LoadConfig("test_files/sample-config.yaml")
	assert.NoError(t, err)
	err = processor.ProcessTerraformTemplate(config)
	assert.NoError(t, err)
	_, err = os.Stat(path.Join(tmpDir, common.DefaultTerraformDir, "terraform.tf"))
	assert.NoError(t, err)
}
