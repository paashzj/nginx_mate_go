package path

import (
	"os"
	"path/filepath"
)

// nginx
var (
	OpenRestyHome                   = os.Getenv("OPENRESTY_HOME")
	NginxHome                       = filepath.FromSlash(OpenRestyHome + "/nginx")
	NginxConfDir                    = filepath.FromSlash(NginxHome + "/conf")
	NginxMainConf                   = filepath.FromSlash(NginxConfDir + "/nginx.conf")
)

// mate
var (
	OpenRestyMatePath  = filepath.FromSlash(OpenRestyHome + "/mate")
	OpenRestyScripts   = filepath.FromSlash(OpenRestyMatePath + "/scripts")
	NginxStartScript   = filepath.FromSlash(OpenRestyScripts + "/start-nginx.sh")
	NginxRestartScript = filepath.FromSlash(OpenRestyScripts + "/restart-nginx.sh")
)
