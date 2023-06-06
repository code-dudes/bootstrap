package env

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	allEnv   map[Environment]bool = map[Environment]bool{}
	localEnv map[Environment]bool = map[Environment]bool{}
	devEnv   map[Environment]bool = map[Environment]bool{}
	prodEnv  map[Environment]bool = map[Environment]bool{}
)

const (
	Local       Environment = "local"
	Development Environment = "development"
	Staging     Environment = "staging"
	Production  Environment = "production"
)

func Env(env string) Environment {
	return Environment(env)
}

func DefaultEnvironments() {
	DefineLocal(Local)
	DefineDevelopment(Development)
	DefineProduction(Production)
	DefineEnvironments(Staging)
}

func DefineEnvironments(environments ...Environment) {
	for _, env := range environments {
		// ignore  empty values
		env = Environment(strings.TrimSpace(env.String()))
		if env != "" {
			allEnv[env] = true
		}
	}
}

// Define Local environnments
// IsLocal will return True
func DefineLocal(environments ...Environment) error {
	for _, env := range environments {
		// ignore  empty values
		env = Environment(strings.TrimSpace(env.String()))
		if env != "" {

			if isDev := devEnv[env]; isDev {
				return fmt.Errorf("env: %s cannot be 'local' since already defined as 'development'", env)
			}

			if isProd := prodEnv[env]; isProd {
				return fmt.Errorf("env: %s cannot be 'local' since already defined as 'production'", env)
			}

			allEnv[env] = true
			localEnv[env] = true
		}
	}
	return nil
}

// Define Development environnments
// IsDevelopment will return True
func DefineDevelopment(environments ...Environment) error {
	for _, env := range environments {
		// ignore  empty values
		env = Environment(strings.TrimSpace(env.String()))
		if env != "" {

			if isLocal := localEnv[env]; isLocal {
				return fmt.Errorf("env: %s cannot be 'development' since already defined as 'local'", env)
			}

			if isProd := prodEnv[env]; isProd {
				return fmt.Errorf("env: %s cannot be 'development' since already defined as 'production'", env)
			}
			allEnv[env] = true
			devEnv[env] = true
		}
	}
	return nil
}

// Define Production environnments
// IsProduction will return True
func DefineProduction(environments ...Environment) error {
	for _, env := range environments {
		// ignore  empty values
		env = Environment(strings.TrimSpace(env.String()))
		if env != "" {

			if isLocal := localEnv[env]; isLocal {
				return fmt.Errorf("env: %s cannot be 'production' since already defined as 'local'", env)
			}

			if isDev := devEnv[env]; isDev {
				return fmt.Errorf("env: %s cannot be 'production' since already defined as 'development'", env)
			}
			allEnv[env] = true
			prodEnv[env] = true
		}
	}
	return nil
}

type Environment string

func (e Environment) IsValid() bool {
	_, ok := allEnv[e]
	return ok
}

func (e Environment) String() string {
	return string(e)
}

func (e Environment) IsLocal() bool {
	_, ok := localEnv[e]
	return ok
}

func (e Environment) IsDevelopment() bool {
	_, ok := devEnv[e]
	return ok
}

func (e Environment) IsProduction() bool {
	_, ok := prodEnv[e]
	return ok
}

func (e Environment) Upper(lang ...language.Tag) string {
	return cases.Upper(langOrEn(lang)).String(e.String())
}

func (e Environment) Lower(lang ...language.Tag) string {
	return cases.Lower(langOrEn(lang)).String(e.String())
}

func (e Environment) Title(lang ...language.Tag) string {
	return cases.Title(langOrEn(lang)).String(e.String())
}

func langOrEn(langs []language.Tag) language.Tag {
	if len(langs) > 0 {
		return langs[0]
	}
	return language.English
}
