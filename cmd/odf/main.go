package main

import (
	rook "github.com/red-hat-storage/odf-cli/cmd/odf/ceph"
	"github.com/red-hat-storage/odf-cli/cmd/odf/get"
	"github.com/red-hat-storage/odf-cli/cmd/odf/root"
	"github.com/red-hat-storage/odf-cli/cmd/odf/set"
	"github.com/red-hat-storage/odf-cli/cmd/odf/subvolume"
	"github.com/rook/kubectl-rook-ceph/pkg/logging"
)

func main() {
	addcommands()
	err := root.RootCmd.Execute()
	if err != nil {
		logging.Fatal(err)
	}
}

func addcommands() {
	root.RootCmd.AddCommand(
		set.SetCmd,
		get.GetCmd,
		rook.CephCmd,
		subvolume.SubvolumeCmd,
	)
}
