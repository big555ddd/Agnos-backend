package console

import (
	"os/exec"

	"github.com/spf13/cobra"

	"app/internal/cmd"
	"app/internal/logger"
)

func helloCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "hello",
		Args: cmd.NotReqArgs,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof("Hello, world")
		},
	}
	return cmd
}

func testCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "test",
		Args: cmd.NotReqArgs,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof("Available test commands:")
			logger.Infof("test patient - Run patient controller tests")
			logger.Infof("test staff   - Run staff controller tests")
		},
	}

	// Add sub-commands for specific module testing
	cmd.AddCommand(testPatientCmd())
	cmd.AddCommand(testStaffCmd())

	return cmd
}

func testPatientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "patient",
		Args: cmd.NotReqArgs,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof("🧪 Running Patient Controller Tests...")
			logger.Infof("📁 File: app/modules/patient/controller_test.go")
			logger.Infof("🎯 Focus: Success/Fail scenarios only")

			// Execute the actual tests
			testCmd := exec.Command("go", "test", "-v", "./app/modules/patient/", "-run", "TestPatientController")
			output, err := testCmd.CombinedOutput()

			if err != nil {
				logger.Err(err)
				logger.Infof("Output: %s", string(output))
			} else {
				logger.Infof("✅ Patient tests completed!")
				logger.Infof("Output: %s", string(output))
			}
		},
	}
	return cmd
}

func testStaffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "staff",
		Args: cmd.NotReqArgs,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof("🧪 Running Staff Controller Tests...")
			logger.Infof("📁 File: app/modules/staff/controller_test.go")
			logger.Infof("🎯 Focus: Success/Fail scenarios only")

			// Execute the actual tests
			testCmd := exec.Command("go", "test", "-v", "./app/modules/staff/", "-run", "TestStaffController")
			output, err := testCmd.CombinedOutput()

			if err != nil {
				logger.Err(err)
				logger.Infof("Output: %s", string(output))
			} else {
				logger.Infof("✅ Staff tests completed!")
				logger.Infof("Output: %s", string(output))
			}
		},
	}
	return cmd
}
