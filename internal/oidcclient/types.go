// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package oidcclient

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AccessToken is an OAuth2 access token.
type AccessToken struct {
	// Token is the token that authorizes and authenticates the requests.
	Token string `json:"token"`

	// Type is the type of token.
	Type string `json:"type,omitempty"`

	// Expiry is the optional expiration time of the access token.
	Expiry metav1.Time `json:"expiryTimestamp,omitempty"`
}

// RefreshToken is an OAuth2 refresh token.
type RefreshToken struct {
	// Token is a token that's used by the application (as opposed to the user) to refresh the access token if it expires.
	Token string `json:"token"`
}

// IDToken is an OpenID Connect ID token.
type IDToken struct {
	// Token is an OpenID Connect ID token.
	Token string `json:"token"`

	// Expiry is the optional expiration time of the ID token.
	Expiry metav1.Time `json:"expiryTimestamp,omitempty"`
}

// Token contains the elements of an OIDC session.
type Token struct {
	// AccessToken is the token that authorizes and authenticates the requests.
	AccessToken *AccessToken `json:"access,omitempty"`

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken *RefreshToken `json:"refresh,omitempty"`

	// IDToken is an OpenID Connect ID token.
	IDToken *IDToken `json:"id,omitempty"`
}

// SessionCacheKey contains the data used to select a valid session cache entry.
type SessionCacheKey struct {
	Issuer      string   `json:"issuer"`
	ClientID    string   `json:"clientID"`
	Scopes      []string `json:"scopes"`
	RedirectURI string   `json:"redirect_uri"`
}

type SessionCache interface {
	GetToken(SessionCacheKey) *Token
	PutToken(SessionCacheKey, *Token)
}
