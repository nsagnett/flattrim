# flattrim
Make flatmap keeping base value type in Golang

# Get library
```
go get -u github.com/nsagnett/flattrim
```
# How to use

```
import "github.com/nsagnett/flattrim"
```
```
flattrimizer := flattrim.NewFlattrimizer(flattrim.KEEPCASE) /// or LOWERCASE to apply a lower case on result keys
flatmap := flattrimizer.Flatten({someJSONData}) // where {someJSONData} is map[string]interface{}
```

The difference with some others librairies is __flattrim__ keeps base value types.
