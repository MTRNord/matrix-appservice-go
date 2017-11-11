package registration

import "testing"

func TestSettingData(t *testing.T) {
	data := NewAppServiceRegistration("localhost")
	if data.Url != "localhost" {
		t.Error("AppServiceUrl did not get set on init!")
	}

	AsToken := GenerateToken()
	data.SetAppServiceToken(AsToken)
	if data.AsToken != AsToken {
		t.Error("AsToken did not get set!")
	}

	HsToken := GenerateToken()
	data.SetHomeserverToken(HsToken)
	if data.HsToken != HsToken {
		t.Error("HsToken did not get set!")
	}

	ID := GenerateToken()
	data.SetID(ID)
	if data.Id != ID {
		t.Error("ID did not get set!")
	}

	data.SetAppServiceUrl("localhost1")
	if data.Url != "localhost1" {
		t.Error("AppServiceUrl did not get changed!")
	}

	data.SetSenderLocalpart("bot")
	if data.SenderLocalpart != "bot" {
		t.Error("SenderLocalpart did not get changed!")
	}
}