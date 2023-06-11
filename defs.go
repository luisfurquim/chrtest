package main

import (
	"hash"
	"net/url"
	"net/http"
	"github.com/luisfurquim/goose"
)

const (
//	WindowWidth int = 1280
//	WindowHeight int = 720
	WindowWidth int = 1920
	WindowHeight int = 1080
)


type GooseG struct {
	Init goose.Alert
	Search goose.Alert
	Collect goose.Alert
}

var Goose GooseG = GooseG{
	Init: goose.Alert(3),
	Search: goose.Alert(5),
	Collect: goose.Alert(5),
}

type Metadata struct {
	ReqMethod              string
	ReqURL                *url.URL
	ReqHeader              http.Header
	ReqBody              []byte
	ReqTransferEncoding  []string

	RespStatus             string // e.g. "200 OK"
	RespProto              string // e.g. "HTTP/1.0"
	RespHeader             http.Header
	RespTransferEncoding []string
	Name                   string
	Ext						  string
	Mime						  string

	hashMD5, hashSHA1, hashSHA256, hashSHA512 hash.Hash
}

