// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package libkb

import (
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
)

// PassphraseGeneration represents which generation of the passphrase is
// currently in use.  It's used to guard against race conditions in which
// the passphrase is changed on one device which the other still has it cached.
type PassphraseGeneration int

// IsNil returns true if this PassphraseGeneration isn't initialized.
func (p PassphraseGeneration) IsNil() bool { return p == PassphraseGeneration(0) }

// LoginContext is passed to all loginHandler functions.  It
// allows them safe access to various parts of the LoginState during
// the login process.
type LoginContext interface {
	LoggedInLoad() (bool, error)
	Salt() []byte
	CreateStreamCache(tsec Triplesec, pps *PassphraseStream)
	SetStreamCache(c *PassphraseStreamCache)
	PassphraseStreamCache() *PassphraseStreamCache
	ClearStreamCache()
	SetStreamGeneration(gen PassphraseGeneration, nilPPStreamOK bool)
	GetStreamGeneration() PassphraseGeneration

	CreateLoginSessionWithSalt(emailOrUsername string, salt []byte) error
	LoadLoginSession(emailOrUsername string) error
	LoginSession() *LoginSession
	ClearLoginSession()
	SetLoginSession(l *LoginSession)

	LocalSession() *Session
	GetUID() keybase1.UID
	GetUsername() NormalizedUsername
	EnsureUsername(username NormalizedUsername)
	SaveState(sessionID, csrf string, username NormalizedUsername, uid keybase1.UID, deviceID keybase1.DeviceID) error
	SetUsernameUID(username NormalizedUsername, uid keybase1.UID) error

	Keyring(m MetaContext) (*SKBKeyringFile, error)
	ClearKeyring()
	SecretSyncer() *SecretSyncer
	RunSecretSyncer(m MetaContext, uid keybase1.UID) error
	Dump(m MetaContext, prefix string)
}

type loginAPIResult struct {
	sessionID string
	csrfToken string
	uid       keybase1.UID
	username  string
	ppGen     PassphraseGeneration
}
