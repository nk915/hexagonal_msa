package terraform

import (
	"context"
	"os"

	"msa/internal/port"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

type tfmodule struct {
	//accessKey string
	//secretKey string
	CTX       context.Context
	Terraform *tfexec.Terraform
}

func NewTerraformExec(workDir string) (port.TerraformExec, error) {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		fmt.Fatalf("error installing Terraform: %s", err)
		fmt.Panic(err)
		return nil, err
	}

	tf, err := tfexec.NewTerraform(workDir, execPath)
	if err != nil {
		fmt.Fatalf("error running NewTerraform: %s", err)
		fmt.Panic(err)
		return nil, err
	}

	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	ctx := context.Background()
	err = tf.Init(ctx, tfexec.Upgrade(true))
	if err != nil {
		// log.Fatalf("error running Init: %s", err)
		return nil, err
	}

	return &tfmodule{Terraform: tf, CTX: ctx}, nil
}

func (tfm *tfmodule) Plan() (bool, error) {
	return tfm.Terraform.Plan(tfm.CTX)
}

func (tfm *tfmodule) PlanWithOpt(opts ...tfexec.PlanOption) (bool, error) {
	return tfm.Terraform.Plan(tfm.CTX, opts)
}

func (tfm *tfmodule) Apply() error {
	return tfm.Terraform.Apply(tfm.CTX)
}

func (tfm *tfmodule) ApplyWithOpt(opts ...tfexec.ApplyOption) error {
	return tfm.Terraform.Apply(tfm.CTX, opts)
}

func (tfm *tfmodule) Destroy() error {
	return tfm.Terraform.Destroy(tfm.CTX)
}

func (tfm *tfmodule) DestroyWithOpt(opts ...tfexec.DestroyOption) error {
	return tfm.Terraform.Destroy(tfm.CTX, opts)
}

func (tfm *tfmodule) WorkspaceList() ([]string, string, error) {
	return tfm.Terraform.WorkspaceList(tfm.CTX)
}

// workspace list
// workspace new
// workspace delete
