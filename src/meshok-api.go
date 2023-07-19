package meshokapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type MeshokAPI struct {
	baseURL string
	token   string
}

func NewMeshokAPI(token string) *MeshokAPI {
	return &MeshokAPI{
		baseURL: "https://api.meshok.net/sAPIv1/",
		token:   token,
	}
}

func (m *MeshokAPI) sendRequest(method string, params url.Values) (map[string]interface{}, error) {
	url := m.baseURL + method
	headers := http.Header{
		"Authorization": []string{"Bearer " + m.token},
	}

	resp, err := http.PostForm(url, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (m *MeshokAPI) getItemList() (map[string]interface{}, error) {
	return m.sendRequest("getItemList", url.Values{})
}

func (m *MeshokAPI) getFinishedItemList() (map[string]interface{}, error) {
	return m.sendRequest("getFinishedItemList", url.Values{})
}

func (m *MeshokAPI) getUnsoldFinishedItemList() (map[string]interface{}, error) {
	return m.sendRequest("getUnsoldFinishedItemList", url.Values{})
}

func (m *MeshokAPI) getSoldFinishedItemList() (map[string]interface{}, error) {
	return m.sendRequest("getSoldFinishedItemList", url.Values{})
}

func (m *MeshokAPI) getItemInfo(id string) (map[string]interface{}, error) {
	params := url.Values{"id": []string{id}}
	return m.sendRequest("getItemInfo", params)
}

func (m *MeshokAPI) getAccountInfo() (map[string]interface{}, error) {
	return m.sendRequest("getAccountInfo", url.Values{})
}

func (m *MeshokAPI) getCommonDescriptionList() (map[string]interface{}, error) {
	return m.sendRequest("getCommonDescriptionList", url.Values{})
}

func (m *MeshokAPI) getSubCategory(id string) (map[string]interface{}, error) {
	params := url.Values{"id": []string{id}}
	return m.sendRequest("getSubCategory", params)
}

func (m *MeshokAPI) getCategoryInfo(id string) (map[string]interface{}, error) {
	params := url.Values{"id": []string{id}}
	return m.sendRequest("getCategoryInfo", params)
}

func (m *MeshokAPI) getCurencyList() (map[string]interface{}, error) {
	return m.sendRequest("getCurencyList", url.Values{})
}

func (m *MeshokAPI) getCountryList() (map[string]interface{}, error) {
	return m.sendRequest("getCountryList", url.Values{})
}

func (m *MeshokAPI) getCitiesList(id string) (map[string]interface{}, error) {
	params := url.Values{"id": []string{id}}
	return m.sendRequest("getCitiesList", params)
}

func (m *MeshokAPI) stopSale(id string) (map[string]interface{}, error) {
	params := url.Values{"id": []string{id}}
	return m.sendRequest("stopSale", params)
}

func (m *MeshokAPI) relistItem(id string) (map[string]interface{}, error) {
	params := url.Values{"id": []string{id}}
	return m.sendRequest("relistItem", params)
}

func (m *MeshokAPI) deleteItem(id string) (map[string]interface{}, error) {
	params := url.Values{"id": []string{id}}
	return m.sendRequest("deleteItem", params)
}

func (m *MeshokAPI) listItem(params url.Values) (map[string]interface{}, error) {
	return m.sendRequest("listItem", params)
}

func (m *MeshokAPI) updateItem(params url.Values) (map[string]interface{}, error) {
	return m.sendRequest("updateItem", params)
}
