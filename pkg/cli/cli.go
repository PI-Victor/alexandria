package cli

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verify  string
	filter  string
	encrypt string
	// TODO: see how an image signing might work in this case.
	sign string
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
	logrus.Info("Not implemented yet")
	return nil
}

// pullImages copies a remote image file locally.
// if this fails, it will only log the error to let the user know.
func pullImages(urls []*url.URL) {
	err := make(chan error)
	msg := make(chan string)

	for _, u := range urls {
		go func(u *url.URL) {
			if e := downloadFile(msg, u.String()); e != nil {
				err <- e
			}
		}(u)
		logrus.Warn(<-msg)
		logrus.Warn(<-err)
	}

}

func downloadFile(msg chan string, dlURL string) error {
	var err error
	tokens := strings.Split(dlURL, "/")
	fileName := tokens[len(tokens)-1]
	configDir := os.Getenv("HOME")

	if configDir == "" {
		configDir, err = os.Getwd()
		if err != nil {
			return err
		}
	}
	filePath := path.Join(configDir, ".alexandria", "images")
	file := path.Join(filePath, fileName)
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		if err = os.MkdirAll(filePath, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	fh, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fh.Close()

	response, err := http.Get(dlURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	msg <- fmt.Sprintf("Downloading %s...", dlURL)
	n, err := io.Copy(fh, response.Body)
	if err != nil {
		return err
	}
	logrus.Info(n)
	return nil
}

func init() {
	PullImages.PersistentFlags().StringVar(&verify, "verify", "", "Verify the checksum of the image after download.")
	PullImages.PersistentFlags().StringVar(&encrypt, "encrypt", "", "Encrypt image locally with personal GPG Key.")

	PullImages.PersistentFlags().StringVar(&sign, "sign", "", "Sign an image that you push to the library")

	ListImages.PersistentFlags().StringVar(&filter, "filter", "iso", "Filter images by image extension.")
}
