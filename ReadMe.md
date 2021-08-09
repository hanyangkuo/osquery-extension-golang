# Osquery Extension - Golang.
Testing with osquery version: 4.6.0.2, golang version: 1.16.7
## Build extensions
- Download libs: `go mod download`
- Build extension execute file: 
`go build -o script_example.exe main.go`
- Build extension execute file: 
`go build -o script_example.exe -ldflags "-s -w" main.go`

For details, refer [Optimizing the size of the Go binary](https://prog.world/optimizing-the-size-of-the-go-binary/)

## Using osquery extension plugins
- copy configuration files to osquery
```
# Copy files
Copy-Item .\osquery_conf\extensions.load -Destination 'C:\Program Files\osquery\extensions.load'
Copy-Item .\osquery_conf\osquery.conf -Destination 'C:\Program Files\osquery\osquery.conf'
Copy-Item .\osquery_conf\osquery.flags -Destination 'C:\Program Files\osquery\osquery.flags'
New-Item -ItemType directory -Path 'C:\Program Files\osquery\Extensions'
Copy-Item .\script_example.exe -Destination 'C:\Program Files\osquery\Extensions\script_example.exe'
Copy-Item .\scripts -Destination 'C:\Program Files\osquery' -Recurse
```
- switch to osquery location
cd "C:\Program Files\osquery"

- grant execute file permission
```
icacls .\Extensions /setowner Administrators /t
icacls .\Extensions /grant Administrators:f /t
icacls .\Extensions /inheritance:r /t
icacls .\Extensions /inheritance:d /t
```

![](.\images\osquery_grant.png)

- Test extension plugins with osqueryi
```
.\osqueryi --flagfile .\osquery.flags
select * from script_example;
select * from registry_example;
.exit
```
![](.\images\osquery_registry.png)

- Restart service osqueryd 
```
.\manage-osqueryd.ps1 -stop
.\manage-osqueryd.ps1 -start
```
osqueryd with create log under `C:\Program Files\osquery\log`
![](.\images\osquery_loglocation.png)

~ ERROR appeared ~
server.RegisterPlugin (failed) -> osqueryd (created socket) -> osqueryd schedule query (failed) -> server.RegisterPlugin (re-register) -> osqueryd schedule query (success)

![](.\images\osquery_log.png)

[how to fixed] Changed time.sleep to ensure extension server registerPlugin after socket has been created by osqueryd
osqueryd (created socket) -> server.RegisterPlugin (success)  -> osqueryd schedule query (success)

![](.\images\golang_sleep.png)


## Reference
- [OSQuery Official](https://osquery.io/downloads/official/4.6.0)
- [OSQuery Docs - Developing osquery Extensions](https://osquery.readthedocs.io/en/stable/deployment/extensions/#extensions-binary-permissions)
- [Optimizing the size of the Go binary](https://prog.world/optimizing-the-size-of-the-go-binary/)

###### Tags `osquery` `golang`

