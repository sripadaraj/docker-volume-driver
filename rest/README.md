## ReST api


> The rest will create a new client/volume , stop/check with user response 


### usage 
 
 ## 1. using method one view source code ![HERE](https://github.com/sripadaraj/ReST_api/blob/master/client.go)
> use the below package and create a new api client using the https portal
```go
 package main

import (
  "log"
  docker-volume-driver_api "github.com/sripadaraj/ReSET_api"
)

func main() {
    client := docker-volume-driver_api.NewrpcClient("http://yourapiserver:<port>", "user", "password")
    req := &docker-volume-driver_api.CreateVolumeRequest{
        Name:              "MyVolume",
        RootUserID:        "root",
        RootGroupID:       "root",
        ConfigurationName: "base",
    }
    volume_uuid, err := client.CreateVolume(req)
    if err != nil {
        log.Fatalf("Error:", err)
    }

    log.Printf("%s", volume_uuid)
}
 ``` 
## 2. Using method two view source code ![HERE](https://github.com/sripadaraj/ReST_api/blob/master/rest/client.go)

- Newclient () --> creates a new client 
  ```go
  func NewClient(addr, base string) *Client {
        return &Client{addr, base}
}


- Volumeexist () --> To check whether volume exists (or) not 
 ```go
 func (r Client) VolumeExist(name string) (bool, error) {
        vols, err := r.volumes()
        if err != nil {
                return false, err
        }

        for _, v := range vols {
                if v.Name == name {
                        return true, nil
                }
        }

        return false, nil
}
```
- volume () --> to check and compare the volume names with the local host

- CreateVolume () --> to stop the volume 
 ```go 
    
    ```
    
    ``` code will be updated soon ```

