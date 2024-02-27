A Go program using the Gin framework to create an API endpoint that fetches an RSS feed from Some Persian News (see them in below) Website and returns it as JSON.

# Rss API Integration

This project focuses on web scraping and crawling news articles from various Iranian news websites. It utilizes Golang for efficient web crawling and JSON data generation. The primary goal is to provide a tool that gathers up-to-date news content and sends it in a structured JSON format to individuals or systems that require this information.

| Category | Web Link | Reading Type |
|  :---:  |     :---:      |          :---: |
| Tech   | [Zoomit](https://www.zoomit.ir/)     | RSS    |
| Tech    | [Digiato](http://www.digiato.com)     | RSS     |
| All     | [Tasnim](https://www.tasnimnews.com)      | RSS     |
| All     | [Tabnak](https://www.tabnak.ir)      | RSS     |
| All     | [Yjc](https://www.yjc.ir)     | RSS     |
| Gaming & Movie     | [Zoomg](https://www.zoomg.ir)     | RSS     |
| Car     | [KhodroBank](https://www.khodrobank.com/)     | RSS     |
| Car Price     | [Car](https://www.Car.ir/)     | Direct Req    |

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)

## Installation

To run the Prss API Integration locally, make sure you have Go installed. Clone this repository and execute the following command:

```bash
go run main.go
```
## Docker Installation
To run the Prss API Integration via docker, make sure you have docker installed and execute the following command:

```bash
> git clone https://github.com/einkaaf/prss
```
After Cloning Project just build docker file :
```bash
> docker build -t press .
```
And for the last step , just run container :
```bash
> docker run -d -p 8080:8080 press
```


## Usage
- Once the application is running, access the API endpoint at http://localhost:8080/api/zoomit to retrieve the zoomit RSS feed in JSON format.
- Once the application is running, access the API endpoint at http://localhost:8080/api/digiato to retrieve the digiato RSS feed in JSON format.
- Once the application is running, access the API endpoint at http://localhost:8080/api/tabnak to retrieve the tabnak RSS feed in JSON format.
- Once the application is running, access the API endpoint at http://localhost:8080/api/tasnim to retrieve the tasnim RSS feed in JSON format.
- Once the application is running, access the API endpoint at http://localhost:8080/api/yjc to retrieve the yjc RSS feed in JSON format.
- Once the application is running, access the API endpoint at http://localhost:8080/api/zoomg to retrieve the zoomg RSS feed in JSON format.
- Once the application is running, access the API endpoint at http://localhost:8080/api/khodrobank to retrieve the khodrobank RSS feed in JSON format.
- Once the application is running, access the API endpoint at http://localhost:8080/api/carprice to retrieve the Car.IR RSS feed in JSON format.

## API Endpoints


> GET /api/zoomit |
Fetches the Zoomit RSS feed and returns it as JSON.

> GET /api/digiato |
Fetches the Digiato RSS feed and returns it as JSON.

> GET /api/tabnak |
Fetches the tabnak RSS feed and returns it as JSON.

> GET /api/tasnim |
Fetches the tasnim RSS feed and returns it as JSON.

> GET /api/yjc |
Fetches the yjc RSS feed and returns it as JSON.

> GET /api/zoomg |
Fetches the zoomg RSS feed and returns it as JSON.

> GET /api/khodrobank |
Fetches the khodrobank RSS feed and returns it as JSON.

> GET /api/carprice |
Fetches the Car.IR RSS feed and returns it as JSON.

## Error Handling
The API provides structured error responses for various scenarios:
- Failed RSS Feed Retrieval
```json
{
  "error": "Failed to fetch the RSS feed"
}
```
- Unexpected Status Code 
```json
{
  "error": "Unexpected status code: 502"
}
```
- Invalid XML Format
```json
{
  "error": "Invalid XML format"
}
```
- Failed JSON Conversion
```json
{
  "error": "Failed to convert data to JSON"
}
```
## Contributing
If you'd like to contribute to this project, feel free to open an issue or submit a pull request. Please follow the Contribution Guidelines.

## License
This PRSS (Persian RSS) Integration is licensed under the MIT License. See the LICENSE file for details.
