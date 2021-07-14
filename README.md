This repository contains generated golang code for:
* Event Family: https://desp.kroger.com/event-family/004eaaec-d8a6-350a-becf-9e552255f630
* Version: v1.0.0

To use this in your go client, add the following requirements to your go.mod file.

```
module github.com/krogertechnology/my-go-client-application

go 1.13

require (
	github.com/actgardner/gogen-avro/v9 v9.0.0 // indirect
	github.com/rhammonds1-kr/desp-catalog-go-subscription-subscriptionscore v1.0.0
	...
```