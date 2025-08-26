# Go Shortener

![GOLANG_BADGE](https://img.shields.io/badge/Golang-%2300FFFF?style=for-the-badge&logo=go)
![AWS_BADGE](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)

This project is a modern URL shortener API built using Golang with a serverless cloud-native architecture. The API uses AWS API Gateway for routing, AWS Lambda for serverless compute, DynamoDB for fast NoSQL storage, and S3 for deployment artifacts. The project leverages SAM (Serverless Application Model) for infrastructure as code and automated deployment, following serverless best practices for cost-effectiveness and automatic scaling.

---

## Features

- Modern serverless URL shortener API built with Golang
- Cloud-native architecture using AWS services (Lambda, API Gateway, DynamoDB)
- Infrastructure as Code with AWS SAM (Serverless Application Model)
- Automatic scaling and cost-effective serverless deployment
- Fast NoSQL storage with DynamoDB for URL mappings
- RESTful API endpoints for URL shortening and redirection
- Integrated Swagger documentation for API testing and exploration
- Clean architecture with proper separation of concerns
- Error handling and logging for production reliability
- Support for custom short codes and URL validation

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/gustavobarez/goshortener.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Makefile Commands

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Run the application without generating API documentation.
- `make run-with-docs`: Generate the API documentation using Swag, then run the application.
- `make build`: Build the application and create an executable file named `gopportunities`.
- `make test`: Run tests for all packages in the project.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Remove the `gopportunities` executable and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```sh
make run
```

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [AWS Lambda](https://aws.amazon.com/lambda/) for serverless compute
- [AWS API Gateway](https://aws.amazon.com/api-gateway/) for HTTP routing
- [Amazon DynamoDB](https://aws.amazon.com/dynamodb/) for NoSQL storage
- [Amazon S3](https://aws.amazon.com/s3/) for deployment artifacts storage
- [AWS SAM](https://aws.amazon.com/serverless/sam/) for infrastructure as code
- [Swagger](https://swagger.io/) for API documentation and testing

## Usage

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting job opportunities. The API can be accessed at `http://localhost:$PORT/swagger/index.html`.

Default $PORT if not provided=3000.

## Contributing

To contribute to this project, please follow these guidelines:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them using Conventional Commits
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

---

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

## Credits

This project was created by [Gustavo Barez](https://github.com/gustavobarez).
