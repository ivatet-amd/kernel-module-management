package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-logr/logr"
	kmmcmd "github.com/rh-ecosystem-edge/kernel-module-management/internal/cmd"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/worker"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2/textlogger"
)

var (
	GitCommit = "undefined"
	Version   = "undefined"

	configHelper = worker.NewConfigHelper()
	logger       logr.Logger
	w            worker.Worker
)

var rootCmd = &cobra.Command{
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	SilenceUsage:      true,
	SilenceErrors:     true,
	Use:               "worker",
	Version:           Version,
}

var kmodCmd = &cobra.Command{
	Use:   "kmod",
	Short: "Manage kernel modules",
}

var kmodLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a kernel module",
	Args:  cobra.ExactArgs(1),
	RunE:  kmodLoadFunc,
}

var kmodUnloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "Unload a kernel module",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgPath := args[0]

		logger.V(1).Info("Reading config", "path", cfgPath)

		cfg, err := configHelper.ReadConfigFile(cfgPath)
		if err != nil {
			return fmt.Errorf("could not read config file %s: %v", cfgPath, err)
		}

		mountPathFlag := cmd.Flags().Lookup(worker.FlagFirmwareMountPath)

		return w.UnloadKmod(cmd.Context(), cfg, mountPathFlag.Value.String())
	},
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()

	rootCmd.AddCommand(kmodCmd)

	kmodCmd.AddCommand(kmodLoadCmd, kmodUnloadCmd)

	klogFlagSet := flag.NewFlagSet("klog", flag.ContinueOnError)

	logConfig := textlogger.NewConfig()
	logConfig.AddFlags(klogFlagSet)

	rootCmd.PersistentFlags().AddGoFlagSet(klogFlagSet)

	kmodLoadCmd.Flags().String(
		worker.FlagFirmwareClassPath,
		"",
		"if set, this value will be written to "+worker.FirmwareClassPathLocation,
	)

	kmodLoadCmd.Flags().String(
		worker.FlagFirmwareMountPath,
		"",
		"if set, this the value that firmware host path is mounted to")

	kmodUnloadCmd.Flags().String(
		worker.FlagFirmwareMountPath,
		"",
		"if set, this the value that firmware host path is mounted to")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		logger = textlogger.NewLogger(logConfig).WithName("kmm-worker")

		logger.Info("Starting worker", "version", rootCmd.Version, "git commit", GitCommit)
		logger.Info("Reading pull secrets", "base dir", worker.PullSecretsDir)

		keyChain, err := worker.ReadKubernetesSecrets(cmd.Context(), worker.PullSecretsDir, logger)
		if err != nil {
			return fmt.Errorf("could not read pull secrets: %v", err)
		}

		ip := worker.NewImagePuller(worker.ImagesDir, keyChain, logger)
		mr := worker.NewModprobeRunner(logger)
		res := worker.NewMirrorResolver(logger)
		w = worker.NewWorker(ip, mr, res, logger)

		return nil
	}

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		kmmcmd.FatalError(logger, err, "Fatal error")
		os.Exit(1)
	}
}
