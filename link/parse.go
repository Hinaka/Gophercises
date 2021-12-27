package link

import "io"

// Link represent a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take and HTML document and will return a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
