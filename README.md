# go-klaviyo-api

A Klaviyo API Golang Wrapper for the [klaviyo API](https://developers.klaviyo.com/en/reference/api_overview)

This package implements the API revision `2024-05-15` as the default

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

  // List Klaviyo profiles
  profiles, _, err := client.Profiles.GetProfiles(nil)
```

Use a specific API revision.

```go
  config := klaviyo.ClientConfig{
    RestEndpointRevision: "2024-02-15.pre",
  }

  client := klaviyo.NewWithConfig(config)

  // List Klaviyo profiles
  profiles, _, err := client.Profiles.GetProfiles(nil)
```

Use a specific API revision for a specific service

```go
  client := klaviyo.New()

  service.client.Webhooks.SetServiceRevision("2024-05-15.pre")
  service.client.Profiles.SetServiceRevision("2023-07-15")

  // List Klaviyo accounts using library default revision. (defaultRestAPIRevision = "2024-05-15")
  accounts, _, err := s.client.Accounts.GetAccounts(nil)

  // List Klaviyo webhooks using revision '2024-05-15.pre'.
  hooks, _, err := s.client.Webhooks.GetWebhooks(nil)

  // List Klaviyo profiles using revision '2023-07-15'.
  profiles, _, err := s.client.Profiles.GetProfiles(nil)
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

# Examples

### Events

**List events**
```go
  events, _, err := s.client.Events.GetEvents(nil)
```

**Get event by ID**
```go
  id := "5ckD924B3fj"

  event, _, err := s.client.Events.GetEventByID(id, nil)
```

**Get events by profile ID**
```go
  // Create query filter and set 'profile_id'
  filter := klaviyo.QueryFilter{}
  filter.CreateEqualsFilter("profile_id", "01HZ2BE77YSX0T85PXKZAGFPWB")

  // Create query params and set filter and includes field.
  params := s.client.Events.Query().NewGetEvents()
  params.Filter(filter)
  params.Include([]string{"metric"})

  // Do request
  event, _, err := s.client.Events.GetEvents(params)
```

**Create an event for a profile by email**
```go
  // Create event properties
  eventProperties := make(map[string]interface{})
  eventProperties["event_data_1"] = "value1"
  eventProperties["event_data_2"] = "value2"

  // Create event profile
  eventProfile := &klaviyo.Profile{
    Attributes: &klaviyo.ProfileAttributes{
      Email: "john.doe@example.com",
    },
  }

  // Create event Card and set values
  createEvent := klaviyo.CreateEventCard{}
  createEvent.SetEventMetric("Event name", "")
  createEvent.SetEventProperties(eventProperties)
  createEvent.SetEventProfile(eventProfile)

  // Do request
  resp, err := s.client.Events.CreateEvent(&createEvent)
```

### Profiles

**Search a profile by Email**
```go
  // Create request params filter
  filter := klaviyo.QueryFilter{}
  filter.CreateEqualsFilter("email", "john.doe@example.com")

  // Create profiles query params and set filter
  params := s.client.Profiles.Query().NewGetProfiles()
  params.Filter(filter)

  // Do request
  profiles, _,  err := s.client.Profiles.GetProfiles(params)
```

**Get a profile by ID**
```go
  id := "01HWYQ7EZ3S2ZFSA9G3EGKKBVV"

  profiles, _, err := s.client.Profiles.GetProfileByID(id, nil)
```

**List profiles with subscription statuses**
```go
  params := s.client.Profiles.Query().NewGetProfiles()
  params.SetProfileAdditionalFields([]string{"subscriptions"})

  profiles, _, err := s.client.Profiles.GetProfiles(nil)
```

**Get lists associated with a profile (with the 'name' field included)**
```go
  id := "01HWYQ7EZ3S2ZFSA9G3EGKKBVV"

  params := s.client.Profiles.Query().NewGetProfileLists()
  params.SetListFields([]string{"name"})

  profiles, _, err := s.client.Profiles.GetProfileLists(id, params)
```

**Get segments associated with a profile**
```go
  id := "01HWWDZSHWG5J12PPK2HFG9VH9"

  profiles, _, err := s.client.Profiles.GetProfileSegments(id, nil)
```
