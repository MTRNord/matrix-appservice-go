package registration

import (
	"github.com/MTRNord/matrix-appservice-go/utils"
	"log"
	"regexp"
	"gopkg.in/yaml.v2"
	"os"
	"bufio"
)

type RegexObject struct {
	Exclusive bool           `yaml:"exclusive"`
	Regex     *regexp.Regexp `yaml:"regex"`
}

type Namespaces struct {
	Users   []RegexObject `yaml:",users"`
	Aliases []RegexObject `yaml:",aliases"`
	Rooms   []RegexObject `yaml:",rooms"`
}

// Something is the structure we work with
type AppServiceRegistration struct {
	Url             string   `yaml:"url"`
	Id              string   `yaml:"id"`
	HsToken         string   `yaml:"hs_token"`
	AsToken         string   `yaml:"as_token"`
	SenderLocalpart string   `yaml:"sender_localpart"`
	RateLimited     bool     `yaml:"rate_limited"`
	Namespaces               `yaml:"namespaces"`
	Protocols       []string `yaml:"protocols"`
}

// NewSomething create new instance of Something
func NewAppServiceRegistration(appServiceUrl string) AppServiceRegistration {
	AppServiceRegistrationStruct := AppServiceRegistration {
		Url:         appServiceUrl,
		RateLimited: true,
		}
	return AppServiceRegistrationStruct
}

func GenerateToken() string {
	return utils.RandomString(32)
}

func (a *AppServiceRegistration) SetAppServiceUrl(url string) {
	a.Url = url
}

func (a *AppServiceRegistration) SetID(id string) {
	a.Id = id
}

func (a *AppServiceRegistration) GetID() string {
	return a.Id
}

func (a *AppServiceRegistration) SetProtocols(protocols []string) {
	a.Protocols = protocols
}

func (a *AppServiceRegistration) GetProtocols() []string {
	return a.Protocols
}

func (a *AppServiceRegistration) SetHomeserverToken (token string) {
	a.HsToken = token
}

func (a *AppServiceRegistration) GetHomeserverToken() string {
	return a.HsToken
}

func (a *AppServiceRegistration) SetAppServiceToken(token string) {
	a.AsToken = token
}

func (a *AppServiceRegistration) GetAppServiceToken() string {
	return a.AsToken
}

func (a *AppServiceRegistration) SetSenderLocalpart(localpart string) {
	a.SenderLocalpart = localpart
}

func (a *AppServiceRegistration) GetSenderLocalpart() string {
	return a.SenderLocalpart
}

func (a *AppServiceRegistration) SetRateLimited(isRateLimited bool) {
	a.RateLimited = isRateLimited
}

func (a *AppServiceRegistration) AddRegexPattern(NSType string, regexString string, exclusive bool) error {
	switch NSType {
	case "users":
		regex, err := regexp.Compile(regexString)
		if err != nil {
			return err
		}
		regexObjectStruct := RegexObject{Exclusive: exclusive, Regex: regex}
		a.Namespaces.Users = append(a.Namespaces.Users, regexObjectStruct)
	case "aliases":
		regex, err := regexp.Compile(regexString)
		if err != nil {
			return err
		}
		regexObjectStruct := RegexObject{Exclusive: exclusive, Regex: regex}
		a.Namespaces.Aliases = append(a.Namespaces.Aliases, regexObjectStruct)
	case "rooms":
		regex, err := regexp.Compile(regexString)
		if err != nil {
			return err
		}
		regexObjectStruct := RegexObject{Exclusive: exclusive, Regex: regex}
		a.Namespaces.Rooms = append(a.Namespaces.Rooms, regexObjectStruct)
	default:
		log.Panicln("'NSType' must be 'users', 'rooms' or 'aliases'")
	}
	return nil
}

func (a *AppServiceRegistration) OutputAsYaml(filename string) error {
	data, DataErr := a.getOutput(filename)
	if DataErr != nil {
		return DataErr
	}

	f, CreateErr := os.Create(filename)
	if CreateErr != nil {
		return CreateErr
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	_, WriteErr := w.Write(data)
	if WriteErr != nil {
		return WriteErr
	}

	w.Flush()

	return nil
}

func (a *AppServiceRegistration) getOutput(filename string) ([]byte, error) {
	if a.Id == "" || a.HsToken == "" || a.AsToken == "" || a.Url == "" || a.SenderLocalpart == "" {
		log.Fatalln("Missing required field(s): id, hsToken, asToken, url, senderLocalpart")
	}

	data, err := yaml.Marshal(a)
	if err != nil {
		return nil, err
	}

	return data, nil
}