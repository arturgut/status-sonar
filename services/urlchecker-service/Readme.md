## Description
URL checker allows to validate HTTP status and time taken to complete request for a given URL. Output is handled by net/http package and reports its results in Prometheus and json formats.

---

## Installation 
```bash
git clone https://github.com/arturgut/urlchecker
cd urlchecker
git build *.go
./main
```


---
## API

#### '/metrics' reports in Prometheus format
```bash
curl http://localhost:8091/metrics
```
```bash
# HELP Version: 0.1 Alpha - https://github.com/arturgut/urlchecker
# HELP url_checker. Label: HTTP Response code. Value: Request duration in Ms
url_checker{ url='google.com', http_status_code=200 } 232
url_checker{ url='https://google.com', http_status_code=200 } 56
url_checker{ url='http://facebook.com', http_status_code=200 } 130
```

### The following api endpoint are avaiable: 
> #### /api/add - expects that URL to be passed via GET. 
  ```bash
    $ curl http://localhost:8091/api/add?http://testsite.local
  ```
  ```json
    {
      "google.com": { 
          "url": "http://google.com",
          "responseCode": 200,
          "durationInMs": 232
      },
      "https://fake.local": {
          "url": "https://fake.local",
          "responseCode": 404,
          "durationInMs": 0
      }
    }        

  ```
> #### /api/remove - removes map element using URL to be passed via GET
  ```bash
  $ curl http://localhost:8091/api/remove?http://google.com
  ```
  ```json
  {
      "http://testsite.local": {
          "url": "http://testsite.local",
          "responseCode": 404,
          "durationInMs": 0
      },
      "https://fake.local": {
          "url": "https://fake.local",
          "responseCode": 404,
          "durationInMs": 0
      }
    }        
  ```
> #### /api/list - Lists all map elements 
  
  ```bash
  curl http://localhost:8091/api/list
  ```
  ```json
  {
      "http://testsite.local": {
          "url": "http://testsite.local",
          "responseCode": 404,
          "durationInMs": 0
      },
      "https://fake.local": {
          "url": "https://fake.local",
          "responseCode": 404,
          "durationInMs": 0
      }
    }    
  ```

---

### Configuration 
The configuration is stored in a single `config.yaml` file
```
server: 
  port: 8091
client: 
  skipSSL: true
  timeout: 30 # http.client timeout
  period: 5 # How often to scan a URL. In seconds.
urls:
  - https://google.com
  - http://facebook.com
  - https://amazon.com
  - http://ebay.com
  - http://fake.local # This is a fake one. Handy for testing
```

Configuration with environment variables:
```
yaml:"loglevel" -------> envconfig:"URL_CHECKER_LOGLEVEL"
yaml:"port" -------> envconfig:"URL_CHECKER_SERVER_PORT"
yaml:"timeout" -------> envconfig:"URL_CHECKER_TIMEOUT"
yaml:"period" -------> envconfig:"URL_CHECKER_PERIOD"
```

---

### Roadmap

> ##### 0.1 - Initial release

* *[DONE]* - use Go routines to periodically check URL
* *[DONE]* - Add http.client timeout 
* *[DONE]* - Add dedicated log library
* *[DONE]* - Add support for environment variables for port, log level, period
* *[DONE]* - add /check endpoint to allow dynamically check URL. Eg. /check?url=http:\/\/google.com
* *[DONE]* - allow dynmically manipulate list of URL's with API. Eg. /api/add/{url} /api/remove/{url}
* *[DONE]* - HTTP response should be in JSON 
* *[DONE]* - add /api/list endpoint

> ##### 0.2 - CI/CD 

* *[TODO]* - add Dockerfile, push image to registry with commit SHA
* *[DONE]* - add Makefile
* *[TODO]* - add Jenkinsfile
* *[TODO]* - add tests - https://github.com/stretchr/testify

> ##### 0.3 - Status Sonar prep work
* *[DONE]* - Output as JSON array 
* *[DONE]* - Add support for CORS




