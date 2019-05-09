package spec

import (
	"encoding/xml"
	"time"
)

// URLSitemap entry in sitemap.
type URLSitemap struct {
	Loc        string     `xml:"loc"`
	LastMod    *time.Time `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreq `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
}

// Sitemap represents a complete sitemap which can be marshaled to XML.
type Sitemap struct {
	XMLName xml.Name      `xml:"urlset"`
	XMLNS   string        `xml:"xmlns,attr"`
	URL     []*URLSitemap `xml:"url"`
}

// NewSitemap returns a new Sitemap.
func NewSitemap() *Sitemap {
	return &Sitemap{
		XMLNS: XMLNS,
	}
}
