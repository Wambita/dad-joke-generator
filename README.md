# Dad Joke Generator

A simple web application built with Go that fetches and displays random dad jokes from an external API. The jokes are displayed on a clean, styled HTML page, and users can fetch a new joke with the click of a button.

## Project Structure

```bash
/myproject
│
├── /api
│   └── api.go
│
├── /thandlers
│   └── ihandlers.go
│
├── /static
│   └── styles.css
│
├── /templates
│   └── index.html
│
└── go.mod
│
└── LICENSE
│
└── main.go
│
└── README.md
```


- **static/**: Contains static assets like CSS files.
- **templates/**: Contains HTML templates used by the application.
- **handlers/**: Contains HTTP handler functions that manage routing and render pages.
- **logic/**: Contains core application logic, including the function to fetch jokes from the API.
- **main.go**: The entry point of the application, sets up the server and routes.

## Installation

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/Wambita/dad-joke-generator.git
    cd dad-joke-generator
    ```

2. **Install Go:**

    Ensure Go is installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

3. **Run the Application:**

    ```bash
    go run main.go
    ```

4. **Access the Application:**

    Open your browser and navigate to `http://localhost:8080` to view the Dad Joke Generator.

## Usage

- The application fetches a random dad joke from the [icanhazdadjoke API](https://icanhazdadjoke.com/api).
- Click the "Get Another Joke" button to fetch and display a new joke.

## Project Details

- **Language:** Go
- **API:** [icanhazdadjoke API](https://icanhazdadjoke.com/api)
- **Dependencies:** No external dependencies beyond the Go standard library.

## Contributing

If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to [icanhazdadjoke](https://icanhazdadjoke.com) for providing the API used in this project.
