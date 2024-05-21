# go-klaviyo-api

A Klaviyo API Golang Wrapper for the [klaviyo API](https://developers.klaviyo.com/en/reference/api_overview)

This package implements the API revision `2024-02-15` as the default

# Install

```console
$ go get github.com/dinistavares/go-klaviyo-api
```

# Usage

Import the Klaviyo package.

```go
import "github.com/dinistavares/go-klaviyo-api"
```

Create a new Klaviyo Client and use the provided services.

```go
  client := klaviyo.New()

  // List Klaviyo preferences
  profiles, _, err := client.Profiles.GetProfiles(nil)
```

Use a specific API revision.

```go
  config := klaviyo.ClientConfig{
    RestEndpointRevision: "2024-02-15.pre",
  }

  client := klaviyo.NewWithConfig(config)

  // List Klaviyo preferences
  profiles, _, err := client.Profiles.GetProfiles(nil)
```

## Authenticate

```go
import (
  "github.com/dinistavares/go-klaviyo-api"
)

func main(){
  accessToken := "xxxxxxx"

  client := klaviyo.New()
  client.Authenticate(accessToken)

  // List Klaviyo preferences
  profiles, _, err := client.Profiles.GetProfiles(nil)
}
```
