package zeptomail_test

import (
	_ "embed"
	"os"
)

var (
	//go:embed file/favicon.ico
	fileAttachment     []byte
	zeptoMailAgent     = os.Getenv("ZEPTO_MAIL_AGENT")
	zeptoMailToken     = os.Getenv("ZEPTO_MAIL_TOKEN")
	zeptoMailMgmtToken = os.Getenv("ZEPTO_MAIL_MGMT_TOKEN")
)
