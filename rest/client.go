package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)
// the volume path will also be default when client never mention the path that has to be for the plugin 
const (
	volumePath 	 = "/tmp/mntdir1"
	volumeCreatePath = "/tmp/mntdir1"
	volumeStopPath   = "/tmp/mntsir1"
)

type peer struct {
	ID	string `json:"id"`
	NAME	string `json:"name"`
	Status 	string `json:"status"`
}

type volume struct {
	 Name		string `json:"name"`
	 UUID		string `json:"uuid"`
	 Type		string `json:"type"`
	 Status  	string `json:"status"`
	 NumBricks	int    `json:"num_bricks"`
	 Distribute	int    `json:"distribute"`
	 Striple	int    `json:"stripe"`
	 Replica	int    `json:"replica"`
	 Transport	string `json:"transport"`
}

type reaponse struct {
	ok bool `json:"ok"`
	Err string `json:"error,omitempty"`
}

type peerResponse struct {
	Data []peer `json:"data,omitempty"`
	response
}

type volumeResponse struct {
	Data []volume `json:"data,omitempty"`
	response
}

type Client struct {
	addr string
	base string
}//this struct will send http client that sends request to the ReSET_api


func NewClient(addt,base string) *Client {
	return &Client{addr, base}
}// to create a new client with base n address 

func (r Client) volumeExist(name string) (bool,error) {
	vols, err :=r.volumes()
	if err != nil {
		return false,err
	}

	for _, v := range vols {
 		if v.name == name {
			return true,nil
		}
	}
	return false,nil
}

func (r Client) volumes() ([]volume,error) {
	u := fmt.Sprintf("%s%s",r.addr,volumespath)

	res,err :=http.Get(u)
	if err != nil {
		return nil,err
	}

	var d volumeResponse
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return nil, err
	}

	if !d.ok {
		return nil, fmt.Errorf(d.Err)
	}
	return d.Data, nil
}

// createvolume create a new volume with the gien name in the cluster

func (r Client ) CreateVolume(name string, peers []string)error{
	u :=fmt.sprintf("%s%s", r.addr, fmt.Sprintf(volumeStopPath, name ))

	req, err := http.NewRequest("PUT", u, nil)
	if err != nil {
		return err
	}

	return responseCheck(resp)
}

// stopVolume stops the volume woth the given name in the cluster
func (r Client) StopVolume(name string) error {
 	u :=fmt.Sprintf("%s%s", r.addr, fmt.Sprintf(volumeStopPath, name))

	req, err := http.NewRequest("PUT", u,nil)
	if err != nil {
		return err
	}

	resp, err :=http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return responsecheck(resp)
}
// this will check the response of the json file 
func responseCheck(resp *http.Response) error {
	var p response
	if err := json.NewDecoder(resp.body).Decode(&p); err != nil {
		return err
	}

	if !p.ok {
		return fmt.Errorf(p.Err)
	}
	return nil
}





