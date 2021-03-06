package spec

import (
	"encoding/xml"
)

// URLSitemapIndex entry in sitemap index.
type URLSitemapIndex struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod,omitempty"`
}

// SitemapIndex is like Sitemap except the elements are named differently.
type SitemapIndex struct {
	XMLName xml.Name           `xml:"sitemapindex"`
	XMLNS   string             `xml:"xmlns,attr"`
	Sitemap []*URLSitemapIndex `xml:"sitemap"`
}

// NewSitemapIndex returns new SitemapIndex.
func NewSitemapIndex() *SitemapIndex {
	return &SitemapIndex{
		XMLNS:   XMLNS,
		Sitemap: []*URLSitemapIndex{},
	}
}
