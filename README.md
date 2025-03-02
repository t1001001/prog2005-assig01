# prog2005-assignment-01

## Country Information Service
**Author:** Tobias Nguyen - tobiasng@stud.ntnu.no

The Country Information Service is a RESTful API for provides country-related information by retrieving information from two API endpoints.
- **RestCountries API:** https://restcountries.com/v3.1
- **CountriesNow API:** https://countriesnow.space/api/v0.1

*Please note that these URLs are the actual API endpoints.*
*This Assignment strictly invokes a self-hosted version of both endpoints.*

## How to use the service?
To run the service, simply run the `main.go` file

`go run cmd/country-information-service/main.go `

## The API endpoints
The service provides three endpoints:

### Info Enpoint
`info/{two_letter_country_code}{limit=n}`

This endpoint returns general information for a given country.
- **two_letter_country_code** is the corressponding 2-letter country ISO code.
- **limit** is the number of cities that are listed in the response where **n** is the limit.
*This parameter is optional.*

Example request:
- info/de
- info/de?limit=10
---
### Population Endpoint
`population/{two_letter_country_code}{limit={startYear-endYear}}`

This endpoint returns historical population information for a given country.
- **two_letter_country_code** is the corressponding 2-letter country ISO code.
- **limit** contrains the population history to values between a **startYear** and an **endYear**.
*This parameter is optional.*

Example request:
- population/de
- population/de?limit=2010-2015
---
### Status Endpoint
`status/`

This endpoint indicates the availability of the individual services. Furthermore, it also returns the uptime of the whole service.

# Testing
To ensure the functionality of the service, tests have been implemented for the handlers.

The tests are located in the `handlers` package and can be run using the following ways:

- To execute all the tests, please run `go test ./...`.

- If a more detailed output when running the tests is desired, please run `go test -v ./...`.

# Deployment
This service is also deployed on Render.

https://prog2005-assignment-01-dx3w.onrender.com

Everytime an update is pushed, it will automatically deploy to Render and the NTNU GitLab instance.

For insight, look into the .yml file which is located in `.github/workflows/deployment.yml`.