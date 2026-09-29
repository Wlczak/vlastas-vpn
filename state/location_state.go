package state

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	h "github.com/Wlczak/vlastas-vpn/helpers"
)

func GetServerLocationService() *ServerLocationService {
	return GetServerState().LocationService
}

func (service *ServerLocationService) GetLocationList() []MullvadServerLocation {
	if time.Now().Unix()-1000*60 > service.lastFetchedAt {
		service.locationList = fetchMullvadLocationList()
	}
	return service.locationList
}

func (service *ServerLocationService) SelectFirstAvailableLocation() error {
	locationList := service.GetLocationList()
	if len(locationList) == 0 {
		return errors.New("no locations available")
	}
	service.CurrentLocation = locationList[0]
	return nil
}

func (service *ServerLocationService) SetServerLocationByHostname(newHostname string) {
	for _, mullvadLocationListItem := range service.GetLocationList() {
		if mullvadLocationListItem.Hostname == newHostname {
			service.CurrentLocation = mullvadLocationListItem
		}
	}
}

func fetchMullvadLocationList() []MullvadServerLocation {
	url := "https://api.mullvad.net/www/relays/all" // TODO: Move to constants envetually
	resp, err := http.Get(url)
	h.PanicChckErr(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	h.PanicChckErr(err)

	mullvadServerLocationList := []MullvadServerLocation{}
	err = json.Unmarshal(body, &mullvadServerLocationList)
	h.PanicChckErr(err)

	return mullvadServerLocationList
}
