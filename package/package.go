package xsocket

import (
	"github.com/jurgen-kluft/xbase/package"
	"github.com/jurgen-kluft/xhash/package"
	"github.com/jurgen-kluft/xtime/package"
	"github.com/jurgen-kluft/xuuid/package"
	"github.com/jurgen-kluft/xcode/denv"
	"github.com/jurgen-kluft/xentry/package"
	"github.com/jurgen-kluft/xunittest/package"
)

// GetPackage returns the package object of 'xsocket'
func GetPackage() *denv.Package {
	// Dependencies
	xunittestpkg := xunittest.GetPackage()
	xentrypkg := xentry.GetPackage()
	xbasepkg := xbase.GetPackage()
	xhashpkg := xhash.GetPackage()
	xtimepkg := xtime.GetPackage()
	xuuidpkg := xuuid.GetPackage()

	// The main (xsocket) package
	mainpkg := denv.NewPackage("xsocket")
	mainpkg.AddPackage(xunittestpkg)
	mainpkg.AddPackage(xentrypkg)
	mainpkg.AddPackage(xbasepkg)
	mainpkg.AddPackage(xhashpkg)
	mainpkg.AddPackage(xtimepkg)
	mainpkg.AddPackage(xuuidpkg)

	// 'xsocket' library
	mainlib := denv.SetupDefaultCppLibProject("xsocket", "github.com\\jurgen-kluft\\xsocket")
	mainlib.Dependencies = append(mainlib.Dependencies, xbasepkg.GetMainLib())
	mainlib.Dependencies = append(mainlib.Dependencies, xhashpkg.GetMainLib())
	mainlib.Dependencies = append(mainlib.Dependencies, xtimepkg.GetMainLib())
	mainlib.Dependencies = append(mainlib.Dependencies, xuuidpkg.GetMainLib())

	// 'xsocket' unittest project
	maintest := denv.SetupDefaultCppTestProject("xsocket_test", "github.com\\jurgen-kluft\\xsocket")
	maintest.Dependencies = append(maintest.Dependencies, xunittestpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, xentrypkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, xbasepkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, xhashpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, xtimepkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, xuuidpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
