package library

import (
	"time"
)

// Manifest holds the image manifest,
type Manifest struct {
	User  user  `json:"user"`
	Image image `json:"image"`
}

type user struct {
	Name  string `json:"user_name"`
	Email string `json:"user_email"`
}

type image struct {
	Type      string    `json:"image_type"`
	Size      string    `json:"image_size"`
	CheckSum  string    `json:"image_checksum"`
	EntryDate time.Time `json:"entry_date"`
}

func createManifest() error {
	return nil
}

func retrieveManifest() error {
	return nil
}
