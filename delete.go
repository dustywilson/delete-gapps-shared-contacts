package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type listingFeed struct {
	ID      string         `xml:"id,omitempty"`
	Updated time.Time      `xml:"updated,omitempty"`
	Entries []listingEntry `xml:"entry,omitempty"`
}

type listingEntry struct {
	ID      string    `xml:"id,omitempty"`
	Updated time.Time `xml:"updated,omitempty"`
	Title   string    `xml:"title,omitempty"`
	Links   []struct {
		Relationship string `xml:"rel,attr,omitempty"`
		Type         string `xml:"type,attr,omitempty"`
		URL          string `xml:"href,attr,omitempty"`
	} `xml:"link,omitempty"`
}

func enterTheIncinerator(cs *clientSecret, token string, domain string) error {
	res, err := http.Get(fmt.Sprintf("https://www.google.com/m8/feeds/contacts/%s/full?max-results=%d&access_token=%s", domain, *deleteCount, token))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	dec := xml.NewDecoder(res.Body)
	var feed listingFeed
	err = dec.Decode(&feed)
	if err != nil {
		return err
	}

	for _, entry := range feed.Entries {
		fmt.Printf("%s... ", entry.Title)
		hasEdit := false
		for _, link := range entry.Links {
			if link.Relationship == "edit" {
				hasEdit = true
				editURL, err := url.Parse(link.URL)
				if err != nil {
					fmt.Println("ERROR (bad edit URL).")
					return err
				}
				query := editURL.Query()
				query.Add("access_token", token)
				editURL.RawQuery = query.Encode()
				if doIt != nil && *doIt {
					req, err := http.NewRequest("DELETE", editURL.String(), nil)
					if err != nil {
						fmt.Println("ERROR.")
						return err
					}
					res, err := http.DefaultClient.Do(req)
					if err != nil {
						fmt.Println("ERROR.")
						return err
					}
					fmt.Printf("%s.\n", res.Status)
				} else {
					fmt.Println("skipped; lacking -DOIT flag.")
				}
			}
		}
		if !hasEdit {
			fmt.Println("ERROR: entry missing edit URL.")
		}
	}

	return nil
}
