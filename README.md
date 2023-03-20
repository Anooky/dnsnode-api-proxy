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

