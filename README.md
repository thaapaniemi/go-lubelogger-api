# Lubelogger API Bindings for Go

This repository provides Go bindings for the Lubelogger API, allowing developers to interact with the Lubelogger service using Go.

## Notes
**Note:** This project is currently under development. Stability is not guaranteed.

Required lubelogger version: v1.4.6 due the [usage of culture-invariant API](https://github.com/hargata/lubelog/issues/895). With older versions __vehicles__ and __vehicleinfo__ bindings are not working correctly.

This is an early development version, there will be rough edges, conversion errors and breaking changes.


## Installation
Required go version: 1.11

To install the library, use `go get`:

```bash
go get github.com/thaapaniemi/go-lubelogger-api
```

## Basic Usage
```go
import (
    "github.com/thaapaniemi/go-lubelogger-api"
)

func main(){
	c := lubelogger.NewClient("https://demo.lubelogger.com", "test", "1234")
	ctx := context.Background()
	vehicleId := int64(1)

	rr, err := odometer.GetRecords(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}

	for _, r := range rr {
		fmt.Printf("Odometer %d: %s\n", r.Odometer, r.Date.Format("2006-01-02"))
	}

	r := odometer.Odometer{
		Date:            time.Now(),
		InitialOdometer: 299000,
		Odometer:        299001,
		Notes:           "test",
	}

	err = r.Add(ctx, c, vehicleId)
	if err != nil {
		panic(err)
	}
}
```

## Functions
```go
lubelogger:func NewClient(endpoint, username, password string) client.Client

odometer:func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]OdometerRecord, error)
odometer:func (o OdometerRecord) Add(ctx context.Context, c client.Client, vehicleId int64) error
odometer:func (o OdometerRecord) Update(ctx context.Context, c client.Client) error
odometer:func (o OdometerRecord) Delete(ctx context.Context, c client.Client) error

reminders:func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]Reminder, error)
reminders:func SendReminderEmails(ctx context.Context, c client.Client, urgencies []Urgency) error

vehicleinfo:func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]VehicleInfo, error)
vehicles:func GetRecords(ctx context.Context, c client.Client) ([]VehicleData, error)

servicerecord:func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]ServiceRecord, error)
servicerecord:func (o ServiceRecord) Add(ctx context.Context, c client.Client, vehicleId int64) error
servicerecord:func (o ServiceRecord) Update(ctx context.Context, c client.Client) error
servicerecord:func (o ServiceRecord) Delete(ctx context.Context, c client.Client) error

repairrecords:func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]RepairRecord, error)
repairrecords:func (o RepairRecord) Add(ctx context.Context, c client.Client, vehicleId int64) error
repairrecords:func (o RepairRecord) Update(ctx context.Context, c client.Client) error
repairrecords:func (o RepairRecord) Delete(ctx context.Context, c client.Client) error

upgraderecords:func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]UpgradeRecord, error)
upgraderecords:func (o UpgradeRecord) Add(ctx context.Context, c client.Client, vehicleId int64) error
upgraderecords:func (o UpgradeRecord) Update(ctx context.Context, c client.Client) error
upgraderecords:func (o UpgradeRecord) Delete(ctx context.Context, c client.Client) error

taxrecords:func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]TaxRecord, error)
taxrecords:func (o TaxRecord) Add(ctx context.Context, c client.Client, vehicleId int64) error
taxrecords:func (o TaxRecord) Update(ctx context.Context, c client.Client) error
taxrecords:func (o TaxRecord) Delete(ctx context.Context, c client.Client) error

gasrecords:func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]GasRecord, error)
gasrecords:func (o GasRecord) Add(ctx context.Context, c client.Client, vehicleId int64) error
gasrecords:func (o GasRecord) Update(ctx context.Context, c client.Client) error
gasrecords:func (o GasRecord) Delete(ctx context.Context, c client.Client) error

calendar:func GetCalendar(ctx context.Context, c client.Client) (string, error)

document:func (o Document) Upload(ctx context.Context, c client.Client) error

root:func MakeBackup(ctx context.Context, c client.Client) ([]byte, error)
root:func Cleanup(ctx context.Context, c client.Client) ([]byte, error)

client:func (c *Client) HttpClient(newClient *http.Client)
```

## Custom HTTP client
Current version uses Go standardlibrary http.DefaultClient as http client. This can be overrid with client.HttpClient function
```
httpClient := &http.Client{
    Timeout: 10 * time.Second,
}

c := lubelogger.NewClient(...)
c.HttpClient(httpClient)
```

## Debug logs
Debug logging can be enabled with
```
import "github.com/thaapaniemi/go-lubelogger-api/debuglog"
...
debuglog.Enabled = true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests.

1. Fork the repository.
2. Create a new feature branch.
3. Commit your changes.
4. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Links
* [Lubelogger](https://lubelogger.com/)
* [Lubelogger API](https://docs.lubelogger.com/Advanced/API)
* [Lubelogger API postman collection](https://github.com/hargata/lubelog_scripts/blob/main/misc/LubeLogger.postman_collection.json)
