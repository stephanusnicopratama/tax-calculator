package controllers_test

import (
	"os"
	"tax-calculator/boot"
	"testing"
)

func TestMain(m *testing.M)  {
	os.Chdir("../")

	boot.Bootstrap()

	os.Exit(m.Run())
}
