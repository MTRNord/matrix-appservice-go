package registration

import (
	"github.com/MTRNord/matrix-appservice-go/utils"
	"log"
	"regexp"
)

type RegexObject struct {
	exclusive bool
	regex *regexp.Regexp
}

type Namespaces struct {
	users []RegexObject
	aliases []RegexObject
	rooms []RegexObject
}

// Something is the structure we work with
type AppServiceRegistration struct {
	url string
	id string
	hsToken string
	asToken string
	senderLocalpart string
	rateLimited bool
	Namespaces
	protocols []string
	cachedRegex string
}

// NewSomething create new instance of Something
func NewAppServiceRegistration(appServiceUrl string) AppServiceRegistration {
	AppServiceRegistrationStruct := AppServiceRegistration {
		url: appServiceUrl,
		rateLimited: true,
		}
	return AppServiceRegistrationStruct
}

func GenerateToken() string {
	return utils.RandomString(32)
}

func (a *AppServiceRegistration) SetAppServiceUrl(url string) {
	a.url = url
}

func (a *AppServiceRegistration) SetID(id string) {
	a.id = id
}

func (a *AppServiceRegistration) GetID() string {
	return a.id
}

func (a *AppServiceRegistration) SetProtocols(protocols []string) {
	a.protocols = protocols
}

func (a *AppServiceRegistration) GetProtocols() []string {
	return a.protocols
}

func (a *AppServiceRegistration) SetHomeserverToken (token string) {
	a.hsToken = token
}

func (a *AppServiceRegistration) GetHomeserverToken() string {
	return a.hsToken
}

func (a *AppServiceRegistration) SetAppServiceToken(token string) {
	a.asToken = token
}

func (a *AppServiceRegistration) GetAppServiceToken() string {
	return a.asToken
}

func (a *AppServiceRegistration) SetSenderLocalpart(localpart string) {
	a.senderLocalpart = localpart
}

func (a *AppServiceRegistration) GetSenderLocalpart() string {
	return a.senderLocalpart
}

func (a *AppServiceRegistration) SetRateLimited(isRateLimited bool) {
	a.rateLimited = isRateLimited
}

func (a *AppServiceRegistration) AddRegexPattern(NSType string, regexString string, exclusive bool) error {
	switch NSType {
	case "users":
		regex, err := regexp.Compile(regexString)
		if err != nil {
			return err
		}
		regexObjectStruct := RegexObject{exclusive: exclusive, regex: regex}
		a.Namespaces.users = append(a.Namespaces.users, regexObjectStruct)
	case "aliases":
		regex, err := regexp.Compile(regexString)
		if err != nil {
			return err
		}
		regexObjectStruct := RegexObject{exclusive: exclusive, regex: regex}
		a.Namespaces.aliases = append(a.Namespaces.aliases, regexObjectStruct)
	case "rooms":
		regex, err := regexp.Compile(regexString)
		if err != nil {
			return err
		}
		regexObjectStruct := RegexObject{exclusive: exclusive, regex: regex}
		a.Namespaces.rooms = append(a.Namespaces.rooms, regexObjectStruct)
	default:
		log.Panicln("'NSType' must be 'users', 'rooms' or 'aliases'")
	}
	return nil
}
