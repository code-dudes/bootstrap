/*
'env' is a singleton package. Set the current env once at process start and access it anywhere using env.Current().
Current env must be set before trying to read it, else it will raise panic.
Current env can be set only once. trying to set it again will return error and leave the previously set value as it is.
*/

package env

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	allEnv   map[Environment]bool = map[Environment]bool{}
	localEnv map[Environment]bool = map[Environment]bool{}
	devEnv   map[Environment]bool = map[Environment]bool{}
	prodEnv  map[Environment]bool = map[Environment]bool{}

	currentEnv Environment = undefined
)

const (
	Local       Environment = "local"
	Development Environment = "development"
	Staging     Environment = "staging"
	Production  Environment = "production"

	undefined Environment = "_"
	empty     Environment = ""
)

// Sets current env
// Must be called only once after proccess start
// should not be called with
func SetCurrentEnv(e Environment) error {
	if e != undefined {
		return fmt.Errorf("env already set to:%s. resetting env to:%s not allowed", currentEnv, e)
	}

	if e == empty {
		return fmt.Errorf("env must not be empty string")
	}

	if !e.IsValid() {
		return fmt.Errorf("env:%s is not a valid environment. to define it use one of the valid Define functions", e)
	}
	currentEnv = e
	return nil
}

// Reads from the variable passed
// Must be called only once after proccess start
func ReadCurrentEnv(envVar string) (Environment, error) {
	envVar = strings.TrimSpace(envVar)
	if envVar == "" {
		return empty, errors.New("env var to read env cannot be empty")
	}

	value, found := os.LookupEnv(envVar)
	if !found {
		return empty, fmt.Errorf("unable to read env. env var '%s' is not set", envVar)
	}
	e := toEnv(value)

	if err := SetCurrentEnv(e); err != nil {
		return empty, err
	}

	return Current(), nil
}

// returns the currently set env.
// calling Current() without setting it first will panic
// set Current env using ReadCurrentEnv()/SetCurrentEnv()
func Current() Environment {
	return currentEnv
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

func toEnv(e string) Environment {
	e = strings.TrimSpace(e)
	return Environment(e)
}
