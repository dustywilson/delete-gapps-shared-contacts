package main

import (
	"encoding/json"
	"io/ioutil"
)

type clientSecretRoot struct {
	Installed clientSecret `json:"installed,omitempty"`
}

type clientSecret struct {
	AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url,omitempty"`
	AuthURI                 string   `json:"auth_uri,omitempty"`
	ClientEmail             string   `json:"client_email,omitempty"`
	ClientID                string   `json:"client_id,omitempty"`
	ClientSecret            string   `json:"client_secret,omitempty"`
	ClientX509CertURL       string   `json:"client_x509_cert_url,omitempty"`
	RedirectURIs            []string `json:"redirect_uris,omitempty"`
	TokenURI                string   `json:"token_uri,omitempty"`
}

func getClientSecret(filename string) (*clientSecret, error) {
	csf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var csr clientSecretRoot
	err = json.Unmarshal(csf, &csr)
	if err != nil {
		return nil, err
	}

	return &csr.Installed, nil
}
