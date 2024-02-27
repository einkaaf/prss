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
# Contribute

Thank you for considering contributing to this project! Your involvement is crucial for making it better.

## Ways to Contribute

1. **Bug Reports:** If you encounter any issues while using the project, please [open a new issue](https://github.com/einkaaf/prss/issues/new) and provide detailed information about the problem, including steps to reproduce it.

2. **Feature Requests:** Have a great idea for a new feature? [Open an issue](https://github.com/einkaaf/prss/issues/new) to discuss and share your thoughts on potential enhancements.

3. **Code Contributions:** Feel free to fork the repository, make improvements, and submit a pull request. Please make sure to follow the existing coding style and include relevant documentation. We appreciate your effort to enhance the project.

4. **Documentation:** Help us improve the project's documentation by fixing typos, adding clarifications, or suggesting improvements. Documentation is crucial for the project's success, and your contributions are valuable.

5. **Community Feedback:** Engage with other users, share your experiences, and help answer questions in the [Discussions](https://github.com/einkaaf/prss/discussions) section. Your insights can make a significant impact on the project's development.

## Getting Started

1. **Fork the Repository:** Click the "Fork" button on the top-right corner of this page to create your copy of the repository.

2. **Clone Your Fork:** Clone your forked repository to your local machine using the following command:
   ```
   git clone https://github.com/einkaaf/prss.git
   ```

3. **Create a Branch:** Create a new branch for your contributions. This helps keep your changes isolated and makes the review process smoother.
   ```
   git checkout -b feature-or-fix-branch
   ```

4. **Make Changes:** Implement your changes, following the coding style and guidelines.

5. **Commit Changes:** Commit your changes with a descriptive commit message.
   ```
   git commit -m "Your informative commit message"
   ```

6. **Push Changes:** Push your changes to your fork on GitHub.
   ```
   git push origin feature-or-fix-branch
   ```

7. **Open a Pull Request:** Go to the [Pull Requests]( https://github.com/einkaaf/prss/pulls) section of the original repository and open a new pull request. Provide details about your changes and follow the template, if available.

Thank you, Saeed Karimi, for your invaluable contributions that have greatly enriched and strengthened our project.

Thank you for your contributions!

## License
This PRSS (Persian RSS) Integration is licensed under the MIT License. See the LICENSE file for details.
