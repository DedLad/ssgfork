package main

import (
	"log"

	"github.com/acmpesuecc/anna/cmd/anna"
	"github.com/spf13/cobra"
)

func main() {
	var serve bool
	var addr string
	var renderDrafts bool
	var validateHTML bool
	var prof bool

	rootCmd := &cobra.Command{
		Use:   "anna",
		Short: "Static Site Generator",
		Run: func(cmd *cobra.Command, args []string) {
			annaCmd := anna.Cmd{
				RenderDrafts: renderDrafts,
				Addr:         addr,
			}

			if serve {
				annaCmd.StartLiveReload()
			}

			if prof {
				// TODO: To be filled later
			}

			if validateHTML {
				// anna.ValidateHTMLContent()
			}
			annaCmd.VanillaRender()
		},
	}

	rootCmd.Flags().BoolVarP(&serve, "serve", "s", false, "serve the rendered content")
	rootCmd.Flags().StringVarP(&addr, "addr", "a", "8000", "ip address to serve rendered content to")
	rootCmd.Flags().BoolVarP(&renderDrafts, "draft", "d", false, "renders draft posts")
	rootCmd.Flags().BoolVarP(&validateHTML, "validate-html", "v", false, "validate semantic HTML")
	rootCmd.Flags().BoolVarP(&prof, "prof", "p", false, "enable profiling")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
