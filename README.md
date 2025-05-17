# Weather Application

This is a simple Go application that consumes the [api.weather.gov](https://api.weather.gov) API to fetch weather data based on a given location.

## Project Structure

```
Weather
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── api
│   │   └── weather.go   # Functions to interact with the weather API
│   └── models
│       └── weather.go   # Data structures for weather information
├── go.mod                # Module definition file
├── go.sum                # Dependency checksums
└── README.md             # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd Weather
   ```

2. Install the necessary dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

## Usage

Once the application is running, you can make requests to the defined endpoints to fetch weather data. Refer to the API documentation for specific usage examples.

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.