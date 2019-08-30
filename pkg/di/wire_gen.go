// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/int128/kubelogin/pkg/adaptors/cmd"
	"github.com/int128/kubelogin/pkg/adaptors/credentialplugin"
	"github.com/int128/kubelogin/pkg/adaptors/env"
	"github.com/int128/kubelogin/pkg/adaptors/kubeconfig"
	"github.com/int128/kubelogin/pkg/adaptors/logger"
	"github.com/int128/kubelogin/pkg/adaptors/oidc"
	"github.com/int128/kubelogin/pkg/adaptors/tokencache"
	"github.com/int128/kubelogin/pkg/usecases/auth"
	credentialplugin2 "github.com/int128/kubelogin/pkg/usecases/credentialplugin"
	"github.com/int128/kubelogin/pkg/usecases/standalone"
)

// Injectors from di.go:

func NewCmd() cmd.Interface {
	loggerInterface := logger.New()
	factory := &oidc.Factory{
		Logger: loggerInterface,
	}
	decoder := &oidc.Decoder{}
	envEnv := &env.Env{}
	localServerReadyFunc := _wireLocalServerReadyFuncValue
	authentication := &auth.Authentication{
		OIDCFactory:          factory,
		OIDCDecoder:          decoder,
		Env:                  envEnv,
		Logger:               loggerInterface,
		LocalServerReadyFunc: localServerReadyFunc,
	}
	kubeconfigKubeconfig := &kubeconfig.Kubeconfig{}
	standaloneStandalone := &standalone.Standalone{
		Authentication: authentication,
		Kubeconfig:     kubeconfigKubeconfig,
		Logger:         loggerInterface,
	}
	root := &cmd.Root{
		Standalone: standaloneStandalone,
		Logger:     loggerInterface,
	}
	repository := &tokencache.Repository{}
	interaction := &credentialplugin.Interaction{}
	getToken := &credentialplugin2.GetToken{
		Authentication:       authentication,
		TokenCacheRepository: repository,
		Interaction:          interaction,
		Logger:               loggerInterface,
	}
	cmdGetToken := &cmd.GetToken{
		GetToken: getToken,
		Logger:   loggerInterface,
	}
	cmdCmd := &cmd.Cmd{
		Root:     root,
		GetToken: cmdGetToken,
		Logger:   loggerInterface,
	}
	return cmdCmd
}

var (
	_wireLocalServerReadyFuncValue = auth.DefaultLocalServerReadyFunc
)

func NewCmdForHeadless(loggerInterface logger.Interface, localServerReadyFunc auth.LocalServerReadyFunc, credentialpluginInterface credentialplugin.Interface) cmd.Interface {
	factory := &oidc.Factory{
		Logger: loggerInterface,
	}
	decoder := &oidc.Decoder{}
	envEnv := &env.Env{}
	authentication := &auth.Authentication{
		OIDCFactory:          factory,
		OIDCDecoder:          decoder,
		Env:                  envEnv,
		Logger:               loggerInterface,
		LocalServerReadyFunc: localServerReadyFunc,
	}
	kubeconfigKubeconfig := &kubeconfig.Kubeconfig{}
	standaloneStandalone := &standalone.Standalone{
		Authentication: authentication,
		Kubeconfig:     kubeconfigKubeconfig,
		Logger:         loggerInterface,
	}
	root := &cmd.Root{
		Standalone: standaloneStandalone,
		Logger:     loggerInterface,
	}
	repository := &tokencache.Repository{}
	getToken := &credentialplugin2.GetToken{
		Authentication:       authentication,
		TokenCacheRepository: repository,
		Interaction:          credentialpluginInterface,
		Logger:               loggerInterface,
	}
	cmdGetToken := &cmd.GetToken{
		GetToken: getToken,
		Logger:   loggerInterface,
	}
	cmdCmd := &cmd.Cmd{
		Root:     root,
		GetToken: cmdGetToken,
		Logger:   loggerInterface,
	}
	return cmdCmd
}
