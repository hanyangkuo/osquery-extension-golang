# Hook Powershell Script on Osquery Extension Table.

Testing with osquery version: 4.6.0.2, golang version: 1.16.7
`go mod download`
`go build -o script_example.exe main.go`

```
# Copy files
Copy-Item .\osquery_conf\extensions.load -Destination 'C:\Program Files\osquery\extensions.load'
Copy-Item .\osquery_conf\osquery.conf -Destination 'C:\Program Files\osquery\osquery.conf'
Copy-Item .\osquery_conf\osquery.flags -Destination 'C:\Program Files\osquery\osquery.flags'
New-Item -ItemType directory -Path 'C:\Program Files\osquery\Extensions'
Copy-Item .\script_example.exe -Destination 'C:\Program Files\osquery\Extensions\script_example.exe'
Copy-Item .\script -Destination 'C:\Program Files\osquery' -Recurse

cd "C:\Program Files\osquery"

# grant permission
icacls .\Extensions /setowner Administrators /t
icacls .\Extensions /grant Administrators:f /t
icacls .\Extensions /inheritance:r /t
icacls .\Extensions /inheritance:d /t

# Test with osqueryi
.\osqueryi --flagfile .\osquery.flags
select * from script_example;
.exit

# Restart service osqueryd 
.\manage-osqueryd.ps1 -stop
.\manage-osqueryd.ps1 -start
```


## Reference
- [OSQuery Official](https://osquery.io/downloads/official/4.6.0)
- [OSQuery Docs - Developing osquery Extensions](https://osquery.readthedocs.io/en/stable/deployment/extensions/#extensions-binary-permissions)
###### Tags `osquery` `golang`

