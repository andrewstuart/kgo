// Copyright © 2018 Andrew Stuart
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal // in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"text/template"

	"github.com/Masterminds/sprig"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:generate go-bindata -pkg cmd templates/
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kgo",
	Short: "Scaffolding for kubernetes go applications",
	Run: func(cmd *cobra.Command, args []string) {
		t := template.New("templates")
		fs, err := AssetDir("templates")
		if err != nil {
			log.Fatal(err)
		}

		data := map[string]interface{}{
			"DockerRoot": viper.GetString("docker.registry"),
			"Image":      viper.GetString("docker.image"),
			"Name":       viper.GetString("k8s.app"),
			"Namespace":  viper.GetString("k8s.namespace"),
			"Domain":     viper.GetString("k8s.domain"),
			"Ingress":    viper.GetBool("k8s.ingress"),
			"CORS":       viper.GetBool("go.cors"),
		}

		for _, file := range fs {
			t2, err := t.New(file).Funcs(sprig.TxtFuncMap()).Parse(string(MustAsset("templates/" + file)))
			if err != nil {
				log.Fatal(err)
			}

			force, _ := cmd.Flags().GetBool("force")
			if _, err := os.Stat(file); err == nil && !force {
				fmt.Println(file + " already exists; skipping")
				continue
			}

			f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0640)
			if err != nil {
				fmt.Printf("error opening file %q: %s\n", file, err)
				continue
			}

			err = t2.Execute(f, data)
			if err != nil {
				fmt.Println("error templating: ", err)
			}
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/kgo.yaml)")

	rootCmd.Flags().BoolP("force", "f", false, "force overwrite files")

	rootCmd.Flags().StringP("registry", "r", "", "The default docker registry")
	viper.BindPFlag("docker.registry", rootCmd.Flags().Lookup("registry"))
	viper.SetDefault("docker.registry", "")

	rootCmd.Flags().StringP("namespace", "n", "", "The namespace to override the default")
	viper.BindPFlag("k8s.namespace", rootCmd.Flags().Lookup("namespace"))
	viper.SetDefault("k8s.namespace", "default")

	rootCmd.Flags().StringP("app", "a", "", "The name of the app")
	viper.BindPFlag("k8s.app", rootCmd.Flags().Lookup("app"))

	rootCmd.Flags().StringP("image", "i", "", "The name of the docker image/path")
	viper.BindPFlag("docker.image", rootCmd.Flags().Lookup("image"))

	if dir, err := os.Getwd(); err == nil {
		viper.SetDefault("k8s.app", path.Base(dir))
		viper.SetDefault("docker.image", path.Base(dir))
	}

	rootCmd.Flags().Bool("ingress", false, "Whether to add ingress")
	viper.BindPFlag("k8s.ingress", rootCmd.Flags().Lookup("ingress"))

	rootCmd.Flags().String("domain", "astuart.co", "The dns domain to use for ingress")
	viper.BindPFlag("k8s.domain", rootCmd.Flags().Lookup("domain"))

	rootCmd.Flags().Bool("cors", true, "Use gorilla cors")
	viper.BindPFlag("go.cors", rootCmd.Flags().Lookup("cors"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

		// Search config in home directory with name ".kgo" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(path.Join(home, ".config"))
		viper.SetConfigName("kgo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
