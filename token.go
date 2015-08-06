package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func getAuthCodeURL(cs *clientSecret) string {
	conf := &oauth2.Config{
		ClientID:     cs.ClientID,
		ClientSecret: cs.ClientSecret,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Scopes:       []string{"http://www.google.com/m8/feeds/contacts/"},
		Endpoint:     google.Endpoint,
	}
	return conf.AuthCodeURL("state")
}

func getToken(cs *clientSecret, approval string) (*oauth2.Token, error) {
	conf := &oauth2.Config{
		ClientID:     cs.ClientID,
		ClientSecret: cs.ClientSecret,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Scopes:       []string{"http://www.google.com/m8/feeds/contacts/"},
		Endpoint:     google.Endpoint,
	}
	return conf.Exchange(oauth2.NoContext, approval)
}
