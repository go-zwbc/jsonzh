package sure

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/sure_pkg_gen"
)

func TestGen(t *testing.T) {
	pkgPath := "github.com/go-zwbc/jsonzh"
	t.Log(pkgPath)
	srcRoot := runpath.PARENT.Up(1)
	t.Log(srcRoot)
	outRoot := runpath.PARENT.Path()
	t.Log(outRoot)
	sure_pkg_gen.GenerateSurePackageFiles(t, sure_pkg_gen.NewSurePackageConfig(srcRoot, sure.SOFT, pkgPath).WithOutputRoot(outRoot).WithNewPkgName("jsonszh"))
	sure_pkg_gen.GenerateSurePackageFiles(t, sure_pkg_gen.NewSurePackageConfig(srcRoot, sure.MUST, pkgPath).WithOutputRoot(outRoot).WithNewPkgName("jsonmzh"))
	sure_pkg_gen.GenerateSurePackageFiles(t, sure_pkg_gen.NewSurePackageConfig(srcRoot, sure.OMIT, pkgPath).WithOutputRoot(outRoot).WithNewPkgName("jsonozh"))
}
