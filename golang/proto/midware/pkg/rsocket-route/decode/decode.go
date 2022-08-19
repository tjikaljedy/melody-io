// Copyright by Tjikal
// Ref: https://github.com/rsocket/rsocket-js/blob/1.0.x-alpha/packages/rsocket-composite-metadata/src/AuthMetadata.ts
package decode

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/fxamacker/cbor/v2"
	"github.com/rsocket/rsocket-go/extension"
)

var simpleAuthLength = uint16(2)

const (
	_simpleAuth = "simple"
	_bearerAuth = "bearer"
)

type AuthorizeData struct {
	authType string
	username string
	password string
	bearer   string
}

func Routes(metadataRaw []byte) ([]string, bool) {
	headers := MimeType(metadataRaw)
	if routingTags, exists := headers[extension.MessageRouting.String()]; exists {
		tags, err := extension.ParseRoutingTags(routingTags)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if err == nil && tags != nil && len(tags) > 0 {
			return tags, true
		}
	}
	return nil, false
}

func Authorize(metadataRaw []byte) (auth *extension.Authentication, err error) {
	headers := MimeType(metadataRaw)
	if authTags, exists := headers[extension.MessageAuthentication.String()]; exists {
		auth, err := extension.ParseAuthentication(authTags)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if err == nil && auth != nil {
			return auth, err
		}
	}
	return nil, err
}

func CompositeMetadata(metadataRaw []byte) (headers map[string][]byte) {
	headers = make(map[string][]byte)
	metadata := extension.NewCompositeMetadataBytes(metadataRaw).Scanner()
	for metadata.Scan() {
		mimeType, payloadMetadata, err := metadata.Metadata()
		if err != nil {
			_ = fmt.Errorf("%v", err)
			continue
		}
		headers[mimeType] = payloadMetadata
	}
	return
}

func MimeType(raw []byte) (headers map[string][]byte) {
	var mimeType string
	m := raw[0]
	idOrLen := (m << 1) >> 1
	if m&0x80 == 0x80 {
		mimeType = extension.MIME(idOrLen).String()
	} else {
		mimeTypeLen := int(idOrLen) + 1
		if cap(raw) < 1+mimeTypeLen {
			mimeType = string(raw)
		} else {
			mimeType = string(raw[1 : 1+mimeTypeLen])
		}
	}

	switch mimeType {
	case extension.MessageCompositeMetadata.String():
	case extension.MessageRouting.String():
		headers = CompositeMetadata(raw)
	default:
		headers = make(map[string][]byte)
		headers[mimeType] = []byte{}
	}
	return
}

func Metadata(metadataMimeType string, metadata []byte) (response map[string]string) {
	response = make(map[string]string)
	mime, ok := extension.ParseMIME(metadataMimeType)
	if !ok {
		return
	}

	switch mime {
	case extension.ApplicationJSON:
		_ = json.Unmarshal(metadata, &response)
	case extension.ApplicationCBOR:
		_ = cbor.Unmarshal(metadata, &response)
	}
	return
}

func DecodeSimpleAuthPayload(authType string, auth []byte) *AuthorizeData {
	usernameLength := binary.BigEndian.Uint16(auth)
	if authType == _simpleAuth {
		return &AuthorizeData{
			authType: _simpleAuth,
			username: string(auth[1 : usernameLength+simpleAuthLength]),
			password: string(auth[usernameLength+simpleAuthLength:]),
		}
	}
	return &AuthorizeData{
		authType: _bearerAuth,
		bearer:   string(auth),
	}
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
