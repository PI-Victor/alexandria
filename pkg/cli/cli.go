package cli

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verify    string
	filter    string
	encrypt   string
	overWrite bool
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
	e := make(chan error)
	done := make(chan bool)
	var err error

	for _, u := range urls {
		go func(u *url.URL) {
			downloadFile(u.String(), done, e)
		}(u)

		if err != nil {
			logrus.Warn(<-e)
			continue
		}
		<-done
	}
}

func downloadFile(dlURL string, done chan bool, err chan error) {
	// TODO: clean this up a bit.
	var e error
	var fh *os.File
	tokens := strings.Split(dlURL, "/")
	fileName := tokens[len(tokens)-1]
	configDir := os.Getenv("HOME")

	if configDir == "" {
		configDir, e = os.Getwd()
		if e != nil {
			err <- e
		}
	}
	filePath := path.Join(configDir, ".alexandria", "images")
	file := path.Join(filePath, fileName)
	if _, e = os.Stat(filePath); os.IsNotExist(e) {
		if e = os.MkdirAll(filePath, 0755); e != nil {
			err <- e
		}
	} else if e != nil {
		err <- e
	}
	_, e = os.Stat(filepath.Join(filePath, file))
	if e == nil && !overWrite {
		err <- fmt.Errorf("Image %s already exists, provide --overwrite flag to overwrite", file)
	}
	if !os.IsNotExist(e) {
		err <- e
	}

	fh, e = os.Create(file)
	if e != nil {
		err <- e
	}
	defer fh.Close()

	logrus.Infof("Downloading %s...", fileName)
	response, e := http.Get(dlURL)
	if e != nil {
		err <- e
	}
	defer response.Body.Close()

	n, e := io.Copy(fh, response.Body)
	if e != nil {
		err <- e
	}
	logrus.Infof("Copied %d to %s", n, filePath)
	done <- true
}

func init() {
	PullImages.PersistentFlags().StringVar(&verify, "verify", "", "Verify the checksum of the image after download.")
	PullImages.PersistentFlags().StringVar(&encrypt, "encrypt", "", "Encrypt image locally with personal GPG Key.")
	PullImages.PersistentFlags().StringVar(&sign, "sign", "", "Sign an image that you push to the library")
	PullImages.PersistentFlags().BoolVar(&overWrite, "overwrite", false, "Overwrite images already in the library")

	ListImages.PersistentFlags().StringVar(&filter, "filter", "iso", "Filter images by image extension.")
}
