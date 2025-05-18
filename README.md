# Weather Application

This is a simple Go application that consumes the [open-meteo](https://api.open-meteo.com/) API to fetch weather data based on a given location.

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
   go run cmd/main.go -lat=<latitude> -lon<longitude>
   ```

## Usage

Pass your latitude and logitude and check the current weather conditions.

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.
