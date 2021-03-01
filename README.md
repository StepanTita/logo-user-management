# geo-provider

**REQUESTS:** <br>
* /geo/sources <br>
Response: 
  ```json
  {
    "message": ["locations","devices"] # list of all available sources
  } 
  ```
  
* /geo/data/{source}?count=1&offset=3
Response: 
  ```json
    {
        "message": [
            {
                "_DATE": "2021-02-02",
                "_TIMESTAMP": "2021-02-02 04:05:06 UTC",
                "address": "0x2bAC4BB15eDFE376145342ac0f945862eBAB8029",
                "apps": "",
                "geocash_version": "2.3.8",
                "locale": "en",
                "model": "iPhone SE",
                "os": "iOS",
                "time": "2021-02-02 01:04:18 UTC",
                "uuid": "33DE13AC-F5A8-4E62-AC5A-A3127CB9F2CB",
                "version": "14.3"
            }
        ]
    }
```
# logo-user-management
