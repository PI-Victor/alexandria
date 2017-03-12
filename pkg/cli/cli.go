package cli

import (
	"fmt"
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	chkSum   string
	imageExt string
)

// PullImages pulls a remote image locally and stores in the library.
var PullImages = &cobra.Command{
	Use:   "pull",
	Short: "pulls one or more images from a remote location and stores it locally",
	Example: `alexctl pull https://remotewebsite.com/remote.iso
https://remotewebsite2.com/remote2.iso https://remotewebsite3.com/remote3.iso
`,
	Run: func(cmd *cobra.Command, args []string) {
		var validURLs []*url.URL

		for _, uri := range args {
			u, err := url.ParseRequestURI(uri)
			if err != nil {
				logrus.Warnf("Could not parse URL: %s! Skipping... \n", u)
				continue
			}
			validURLs = append(validURLs, u)
		}
		pullImages(validURLs)
	},
}

// ListImages lists all available images in the library.
var ListImages = &cobra.Command{
	Use:     "list",
	Short:   "lists all available images from the library",
	Example: "alexct list",
	Run: func(cmd *cobra.Command, args []string) {
		listImages()
	},
}

func listImages() error {
	fmt.Println("listing images not implemented yet")
	return nil
}

func pullImages(urls []*url.URL) error {
	fmt.Println("pulling images not implemented yet")
	return nil
}

func init() {
	PullImages.PersistentFlags().StringVar(&chkSum, "chksum", "", "Verify the checksum of the image after download.")
	ListImages.PersistentFlags().StringVar(&imageExt, "imgext", "iso", "Filter images by image extension.")
}
