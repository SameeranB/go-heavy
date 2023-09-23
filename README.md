# go-heavy

A K6 Alternative To Perform Load Testing Written In Golang

---

## Why

There is a need for load testing tools that are meant to be used in tech organisations.

* The tool must be easy to setup as a part of the infrastructure and easy to use by developers.
* It should support a directory structure that is oriented around how teams organise their services.
* It should be easy to write tests and share them across teams.
* It should provide a way to track service specific performance over time.
* It should be open source and free to use. Allowing organisations to add to and edit the tool as they see fit.

---

## What

The goal for Go-Heavy is to provide a set of packages that can be used to write load tests in Golang. The packages will
provide a way to write tests that are easy to read and write. The packages will also provide a way to run tests in
parallel and report the results in a way that is easy to understand.

* A set of packages to be used to write load tests in Golang.
* A CLI tool to run the tests and report the results.

---

## How

This is how we expect the tool to be used. But to reiterate, the tool will be made as flexible as possible to allow
organisations to use it in a way that suits them.

There is an example directory structure in the repo that shows how the tool can be used.

```
-- example
|-- utils
|-- common
|-- services
|   |-- service1
|   |   |-- worflow1.go
|   |   |-- workflow2.go
|   |   |-- .heavy-config.go
|   |   |-- .env
|   |   |-- env.default
```

* The `utils` and `common` directories contains helper functions that can be used across services.
* The `services` directory contains directories for each service.
* Each service directory contains the tests for that service.
* Each service directory contains a `.heavy-config.go` file that contains the configuration for the service.
* Each service directory contains a `.env` file that contains the environment variables for the service.
* Each service directory contains a `env.default` file that contains the default environment variables for the service.

Once tests have been written, the CLI tool can be used to run individual tests or all tests in a service. The CLI tool
will also provide a way to run tests in parallel and report the results in a way that is easy to understand.

There are four important concepts about how the tests are organised:

* Test
* Workflow
* Step
* Request

### Test

A single test, is a single request by the user to perform an action. This action could be running a workflow, or running
all the workflows in a service.

### Workflow

A workflow is a set of steps that are performed in a sequence.
These workflows are setup as individual files in the service directory.

### Step

A step is a set of code that achieve is a single task. A step can be as simple as a single request, or as complex as a
set of requests that are performed in a sequence.

### Request

A request is a single HTTP request that is performed as a part of a workflow.

---

## Features

### CLI

* Run an individual workflow.
* Run all workflows in a service.
* Allow to choose the number of unique users. (Optional)
* Allow to choose one of the following:
    * Allow to choose the number of users running in parallel.
    * Allow to set an WPS (Workflows per second).
* Allow to choose how long to run the workflow for.
* After a test is run, the CLI tool should provide a summary of the results.
* After a test is run, the CLI tool should allow the user to export the results to a file.

### Packages

* A custom HTTP client that tracks metrics and logs automatically.
* An interface to define each test.
    * Init Function
    * Main Body
    * Cleanup Function
    * Path to the environment variables file.

### Reports

#### CLI Report

* Summary
    * Test ID
    * Total run time.
    * Number of Successful Workflow runs.
    * Number of Failed Workflow runs.

* Workflow Name
    * Slowest run time.
    * Fastest run time.
    * Average run time.
    * Total number of runs.
    * Average number of runs per second.
    * Status Code Distribution
        * 200
        * 400
        * 500
        * 300
        * 100
    * Latency Histogram
        * P10
        * P25
        * P50
        * P75
        * P90
        * P95
        * P99
        * P100
    * Latency Distribution
        * Auto generated buckets - Count of requests in each bucket.

#### Complete Report

* Summary
    * Test ID
    * Total run time.
    * Number of Successful Workflow runs.
    * Number of Failed Workflow runs.

* Workflow Name
    * Slowest run time.
    * Fastest run time.
    * Average run time.
    * Total number of runs.
    * Average number of runs per second.
    * Status Code Distribution
        * 200
        * 400
        * 500
        * 300
        * 100
    * Latency Histogram
        * P10
        * P25
        * P50
        * P75
        * P90
        * P95
        * P99
        * P100
    * Latency Distribution
        * Auto generated buckets - Count of requests in each bucket.

    * <details>
        <summary>Steps</summary>

        * Step Name
            * Slowest run time.
            * Fastest run time.
            * Average run time.
            * Total number of runs.
            * Average number of runs per second.
            * Status Code Distribution
                * 200
                * 400
                * 500
                * 300
                * 100
            * Latency Histogram
                * P10
                * P25
                * P50
                * P75
                * P90
                * P95
                * P99
                * P100
            * Latency Distribution
                * Auto generated buckets - Count of requests in each bucket.

            * <details>
                <summary>Requests</summary>

                * Request Endpoint and Method
                    * Maximum latency.
                    * Minimum latency.
                    * Average latency.
                    * Total number of requests.
                    * Average number of requests per second.
                    * Status Code Distribution
                        * 200
                        * 400
                        * 500
                        * 300
                        * 100
                    * Latency Histogram
                        * P10
                        * P25
                        * P50
                        * P75
                        * P90
                        * P95
                        * P99
                        * P100
                    * Latency Distribution
                        * Auto generated buckets - Count of requests in each bucket.
            </details>
    </details>

### Configuration

### Environment Variables

### Parallelism

### Docker

### CI/CD

---

## Installation

### CLI Tool

### Packages

---

## Contributing

### Code of Conduct

### Contributing Guide

### License

---

## Roadmap

### Version 1.0.0

* A CLI tool that supports the following:
    * Run an individual workflow.
    * Run all workflows in a service.
    * Allow to choose the number of unique users. (Optional)
    * Allow to choose one of the following:
        * Allow to choose the number of users running in parallel.
        * Allow to set an WPS (Workflows per second).
    * Allow to choose how long to run the workflow for.

* Packages:
    * A custom HTTP client that tracks metrics and logs automatically.
    * An interface to define each test.
        * Init Function
        * Main Body
        * Cleanup Function
        * Path to the environment variables file.

### Version 2.0.0

* CLI:
    * Import OpenAPI spec and generate tests.
    * Auto generate directory structure
    * A ramp-up function (optional)
    * Export logs to a file.
        * Unique Request ID - which is also set in the headers
        * Timestamp
        * Request
        * Response
        * Status Code
    * Stored test-run-configurations.
    * Diff two tests' performance.
        * A summary of the differences in metrics.
        * Every response that is different between the two tests.
    * After a test is run, the CLI tool should allow the user to dig through the results by providing a mongosh like
      interface.

* WebUI

