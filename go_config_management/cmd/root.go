/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "add",                                // The command's name, used for invocation.
	Short: "Adds an increment to a base number", // A brief description of the command.
	Run: func(cmd *cobra.Command, args []string) {
		// This function runs when the 'add' command is executed.
		base := viper.GetInt("base")           // Retrieves the 'base' value from Viper.
		increment := viper.GetInt("increment") // Retrieves the 'increment' value from Viper.
		result := base + increment             // Calculates the result.
		fmt.Printf("Result: %d\n", result)     // Displays the result.

		// Debugging output to indicate where configuration values are sourced from.
		checkConfigSource(cmd, "base")
		checkConfigSource(cmd, "increment")
	},
}

func init() {
	// Initialize command flags and configuration bindings.
	rootCmd.Flags().Int("base", 0, "Base number to add to")
	rootCmd.Flags().Int("increment", 1, "Increment to add to the base number")
	viper.BindPFlag("base", rootCmd.Flags().Lookup("base"))
	viper.BindPFlag("increment", rootCmd.Flags().Lookup("increment"))

	cobra.OnInitialize(initConfig) // Setup configuration when Cobra initializes.
}

func initConfig() {
	// Configures Viper to manage the application settings.
	viper.SetDefault("base", 10)     // Sets default for 'base'.
	viper.SetDefault("increment", 5) // Sets default for 'increment'.

	viper.SetEnvPrefix("myapp") // Prefix for environment variables.
	viper.BindEnv("base")
	viper.BindEnv("increment")

	viper.SetConfigName("config") // Name of the configuration file.
	viper.SetConfigType("yaml")   // File format for configuration.
	viper.AddConfigPath(".")      // Directory to search for the configuration file.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// checkConfigSource prints the source of a configuration setting.
func checkConfigSource(cmd *cobra.Command, key string) {
	// Construct the environment variable name based on the prefix and key.
	envVarName := "MYAPP_" + strings.ToUpper(key) // Construct the environment variable name.

	// Check if the CLI flag was explicitly set.
	if cmd.Flags().Changed(key) {
		fmt.Printf("%s is set by CLI flag.\n", key)
	} else {
		// Check if the environment variable was set.
		if os.Getenv(envVarName) != "" {
			fmt.Printf("%s is set by environment variable.\n", key)
		} else {
			// Check if the value is set in the config file.
			if viper.InConfig(key) {
				fmt.Printf("%s is set by config file.\n", key)
			} else {
				// If none of the above, then it's using the default value.
				fmt.Printf("%s is using default value.\n", key)
			}
		}
	}
}

// Execute starts the command execution.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
