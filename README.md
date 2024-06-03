# Website Monitoring Tool

This is a simple command-line application written in Go (Golang) for monitoring the availability of websites. It allows users to specify the delay between checks and the number of monitorings to perform.

## Technologies Used

- **Go (Golang)**: The programming language used to develop the application.
- **HTTP Package**: Used for making HTTP requests to the specified websites.
- **OS Package**: Used for file operations, such as reading from `sites.txt` and logging to `log.txt`.
- **Time Package**: Used for introducing delays between monitorings.
- **Log Package**: Used for logging errors and status information.

## Purpose

The primary purpose of this project is to learn how to manipulate files and handle HTTP requests in Go. It serves as a practical exercise to understand concepts such as reading from files, making HTTP requests, error handling, and logging in Go applications.

## Usage

1. **Setup**: List the websites to be monitored in the file `sites.txt`, with each URL on a new line.
2. **Run**: Execute the application and follow the on-screen instructions to start monitoring.
3. **Monitor**: The application will periodically check the specified websites for availability and log the results to `log.txt`.
4. **View Logs**: Use the application to display the logs stored in `log.txt`.

## Sites for Simulation

The `sites.txt` file should contain URLs representing different websites with varying statuses. These URLs can be from personal websites, public APIs, or any online resource that returns different HTTP status codes. For example:

- http://example.com
- http://example.com/notfound
- http://example.com/servererror
- ...

Ensure that the URLs in `sites.txt` are accessible and return the expected status codes to simulate different scenarios accurately.

## Contributing

Contributions are welcome! If you have any suggestions, improvements, or bug fixes, feel free to open an issue or create a pull request.
