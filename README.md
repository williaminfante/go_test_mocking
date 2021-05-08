# Go Test Mocking

The accompanying code repo for the [Golang Testing: Mocking and Error Handling](https://williaminfante.medium.com/golang-testing-mocking-and-error-handling-fbfe7f6008b9) article.

This repository contains:

1. [maintenance package](./src/maintenance/maintenance.go) for the acting third-party dependency package
2. [maintenance_test package](./src/maintenance/maintenance_test.go) for the related tests for the mainatenance package
3. [vehicle package](./src/vehicle/vehicle.go) for the target or in-development package
4. [vehicle_test package](./src/vehicle/vehicle_test.go) for the related tests for the vehicle package
5. [main package](./src/main/main.go)

## Table of Contents

- [Go Test Mocking](#go-test-mocking)
  - [Table of Contents](#table-of-contents)
  - [Background](#background)
  - [Usage](#usage)
  - [Related Link](#related-link)
  - [Contributing](#contributing)

## Background

The accompanying code repo for the [Golang Testing: Mocking and Error Handling](https://williaminfante.medium.com/golang-testing-mocking-and-error-handling-fbfe7f6008b9) article.

There are two branches used for the demo:
- The (`main` branch) contains mocks, interfaces, and receiver functions where vehicle package is decoupled from the maintenance package.
https://github.com/williaminfante/go_test_mocking
- The (`no-mocks` branch) does NOT consider mocks, interfaces, and receiver functions where vehicle package is coupled with the maintenance package. https://github.com/williaminfante/go_test_mocking/tree/no-mocks Please note that this is not the recommended way and was only created for learning purposes.


## Usage

This requires a minimum version of Go 1.16 and testify 1.7.0.

You can also just execute:
```
go mod tidy
```

To run all tests and optionally add `-v` for verbose, you can just run the following the base repo path:
```
go test ./... -v
```
Or go to the folder like the `maintenance` folder and run
```
go test -v
```


To check coverage to the desired output like `cover.out` in the `coverage` for all tests, simply:
```
go test ./... -v -short -coverprofile coverage/cover.out
```

To create an html-formatted coverage document:
```
go tool cover -html=coverage/cover.out -o coverage/cover.html
```

## Related Link

Recommended article [Golang testing with TDD](https://williaminfante.medium.com/golang-testing-with-tdd-e548d8be776) before reading the accompanying article for this code repo.

## Contributing

Please don't hesitate to [raise an issue](https://github.com/williaminfante/go_test_mocking/issues/new) or submit a PR to improve this project. Thanks!
