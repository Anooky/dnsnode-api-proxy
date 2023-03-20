This API proxy can be used to proxy Netnod's DNSNODE API to endcustomers. 

Each endcustomer gets its own Authorization token which is locked to:
 - a static end customer id
 - optional fixed product
 - optional fixed list of master servers
 - optional fixed list of tsig keys 

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

