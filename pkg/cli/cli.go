package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// PullImage pulls a remote image locally and stores in the library.
var PullImage = &cobra.Command{
	Use:     "pull",
	Short:   "pulls the image from a remote location and stores it locally",
	Example: `alexctl pull https://remotewebsite.com/remote.iso`,
	Run: func(cmd *cobra.Command, args []string) {
		// Placeholder for image pull function
		fmt.Println("Not Implemented!")
	},
}
