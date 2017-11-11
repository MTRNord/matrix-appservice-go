package registration

import (
	"bufio"
	"encoding/json"
	"github.com/MTRNord/matrix-appservice-go/utils"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"regexp"
)

type RegexObject struct {
	Exclusive bool           `yaml:"exclusive" json:"exclusive"`
	Regex     *regexp.Regexp `yaml:"regex" json:"regex"`
}

type Namespaces struct {
	Users   []RegexObject `yaml:",flow" json:"users"`
	Aliases []RegexObject `yaml:",flow" json:"aliases"`
	Rooms   []RegexObject `yaml:",flow" json:"rooms"`
}

// Something is the structure we work with
type AppServiceRegistration struct {
	Url             string `yaml:"url" json:"url"`
	Id              string `yaml:"id" json:"id"`
	HsToken         string `yaml:"hs_token" json:"hs_token"`
	AsToken         string `yaml:"as_token" json:"as_token"`
	SenderLocalpart string `yaml:"sender_localpart" json:"sender_localpart"`
	RateLimited     bool   `yaml:"rate_limited" json:"rate_limited"`
	Namespaces      `yaml:"namespaces" json:"namespaces"`
	Protocols       []string `yaml:"protocols" json:"protocols"`
}

// NewSomething create new instance of Something
func NewAppServiceRegistration(appServiceUrl string) *AppServiceRegistration {
	AppServiceRegistrationStruct := AppServiceRegistration{
		Url:         appServiceUrl,
		RateLimited: true,
	}
	return &AppServiceRegistrationStruct
}

func NewFromJson(data []byte) *AppServiceRegistration {
	AppServiceRegistrationStruct := AppServiceRegistration{}
	json.Unmarshal(data, &AppServiceRegistrationStruct)
	return &AppServiceRegistrationStruct
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

func (a *AppServiceRegistration) SetHomeserverToken(token string) {
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

func (a *AppServiceRegistration) isUserMatch(userId string, onlyExclusive bool) bool {
	return a.isMatch(a.Namespaces.Users, userId, onlyExclusive)
}

func (a *AppServiceRegistration) isAliasMatch(alias string, onlyExclusive bool) bool {
	return a.isMatch(a.Namespaces.Aliases, alias, onlyExclusive)
}

func (a *AppServiceRegistration) isRoomMatch(roomId string, onlyExclusive bool) bool {
	return a.isMatch(a.Namespaces.Rooms, roomId, onlyExclusive)
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

func (a *AppServiceRegistration) isMatch(regexList []RegexObject, sample string, onlyExclusive bool) bool {
	for _, regex := range regexList {
		if regex.Regex.MatchString(sample) {
			if onlyExclusive && !regex.Exclusive {
				continue
			}
			return true
		}
	}
	return false
}
