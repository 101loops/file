package fileutil

import (
	. "github.com/101loops/bdd"
)

var _ = Describe("File Utility", func() {

	It("checks if file exists", func() {
		Check(Exists("README.md"), IsTrue)
		Check(Exists("nonsense"), IsFalse)
	})
})
