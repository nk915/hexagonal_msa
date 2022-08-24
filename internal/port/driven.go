package port

import "github.com/hashicorp/terraform-exec/tfexec"

type TerraformExec interface {
	Plan() (bool, error)
	PlanWithOpt(...tfexec.PlanOption) (bool, error)

	Apply() error
	ApplyWithOpt(...tfexec.ApplyOption) error

	Destroy() error
	DestroyWithOpt(...tfexec.DestroyOption) error
}


