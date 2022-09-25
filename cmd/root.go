package cmd


import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
  )

  func init() {

	cfgFile := ""
	projectBase := viper.GetString("project_base")
	userLicense := viper.GetString("user_license")

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "m", "mikej", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "MIKE JOHNSON <mike@mikej.dev>")
	viper.SetDefault("license", "apache")
  }

  func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	cfgFile := viper.GetString("config")
	if cfgFile != "" {
	  // Use config file from the flag.
	  viper.SetConfigFile(cfgFile)
	} else {
	  // Find home directory.
	  home, err := homedir.Dir()
	  if err != nil {
		fmt.Println(err)
		os.Exit(1)
	  }

	  // Search config in home directory with name ".cobra" (without extension).
	  viper.AddConfigPath(home)
	  viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
	  fmt.Println("Can't read config:", err)
	  os.Exit(1)
	}
  }

  var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	},
  }

  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }