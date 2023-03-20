# Netnod DNSNODE API PROXY

This API proxy can be used to proxy Netnod's DNSNODE API to endcustomers. This allows end customers to add/remove zones automatically.

Each endcustomer gets its own Authorization token which enforces customer specific settings:
 - dnsnode end customer id
 - product
 - master servers
 - tsig key


# API implementation status

| Method | Endpoint                | Status        | Comment                                     |
|--------|-------------------------|---------------|---------------------------------------------|
| *      | /tsig/                  | not supported | endcustomers are not allowed to modify TSIG |
| GET    | /zone/                  | implemented   |                                             |
| GET    | /zone/<zone-name>       | implemented   |                                             |
| PUT    | /zone/<zone-name>       | not supported |                                             |
| POST   | /zone/<zone-name>       | implemented   |                                             |
| PATCH  | /zone/<zone-name>       | not supported |                                             |
| DELETE | /zone/<zone-name>       | implemented   |                                             |
| GET    | /status/<zone-name>     | implemented   |                                             |
| GET    | /statistics/graph/<zone-name> | implemented |   only available for premium anycast    |
| GET    | /anomalies/serial/<zone-name>  | implemented   |  can contain inactive instances      |
| GET    | /product/               | not supported |  not planned for endcustomers               |

# configuration example with comments

config.json:

```json 
{
    "dnsnodetoken": "xxxxxxxxxxxxxxxxxxxxx", // the upstream dnsnode API token
    "customerconfigs": {
        "3fb303c89207ddbfbf71fb4299fe6374d7adb298d56f43e5d2e1760b2dd1b00b27f16d3e39ebde4ca23109e9dd158b84e1a03bbba0c1b4a7fb586e3e0e6e6918":{ // sha512 hash of endcustomer's individual access token
            "allowedipranges": [ // list of ips or ip ranges that are allowed to use this token
                "192.168.1.0/24",
                "::1"
            ],
            "endcustomer":"customer1", // name at dnsnode
            "forcedmasters": [ // list of master servers that are forced to be used for this customer
                {
                    "ip": "10.0.0.1",
                    "tsig": "netnod-example-example1.com"
                },
                {
                    "ip": "10.0.0.2",
                    "tsig": "netnod-example-example2.com"
                },
            ],
            "forcedproduct": "se-standard-anycast-a", // product that is forced to be used for this customer
            "maxzones": 50 // max number of zones that can be active at the same time for this customer
        },
        [...] // more customerconfigs

    }
}
```
