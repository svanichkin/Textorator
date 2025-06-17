# Go OpenAI Backend Service

A backend service written in Go that utilizes the OpenAI API for text transformation and generation tasks. It features a REST API for interaction and a local caching mechanism to improve performance and reduce API calls.

## Features

*   **OpenAI Integration**: Connects to OpenAI (GPT-4, gpt-4o) for advanced text processing.
*   **Text Transformation**: Modifies input text based on predefined instructions.
*   **Text Generation**: Generates new text content based on user prompts.
*   **REST API**: Exposes endpoints (`/transform`, `/generate`, `/assist`) for easy integration.
*   **Caching**: Implements a local file-based cache for API responses to optimize speed and reduce costs.
*   **Configuration**: Uses an external `config.ini` file for easy setup of server parameters and API keys.

## Configuration

A `config.ini` file is required in the root directory to run the application. This file stores essential configuration parameters.

**Example `config.ini` structure:**

```ini
host = localhost
port = 8080
openai = YOUR_OPENAI_API_KEY
cache = ./cache/
```

**Note:** Replace `YOUR_OPENAI_API_KEY` with your actual OpenAI API key.

## API Endpoints

The service provides the following REST API endpoints:

*   **`POST /transform`**:
    *   Accepts JSON payload: `{"text": "your input text"}`
    *   Transforms the input text using OpenAI's GPT-4 model with a specific system prompt designed for precise modifications.
    *   Returns JSON payload: `{"transformed_text": "output text"}`

*   **`POST /generate`**:
    *   Accepts JSON payload: `{"text": "your input prompt"}`
    *   Generates new text content using OpenAI's gpt-4o model based on the provided input.
    *   Returns JSON payload: `{"generated_text": "output text"}`

*   **`POST /assist`**:
    *   Accepts JSON payload: `{"text": "your query or instruction"}`
    *   Provides assistance using OpenAI. (The specific behavior of this endpoint, including the model used and the nature of assistance, would require inspecting `rest/handlers.go`.)
    *   Returns JSON payload with the assistant's response.

## How to Run

1.  **Install Go**: Ensure that Go (version 1.x or later) is installed on your system.
2.  **Create `config.ini`**: Create a `config.ini` file in the root directory of the project with the structure specified in the "Configuration" section. Remember to add your OpenAI API key.
3.  **Download Dependencies**: Open a terminal in the project's root directory and run:
    ```bash
    go mod tidy
    ```
    Alternatively, if `go mod tidy` doesn't fetch all dependencies (older Go versions), you might need:
    ```bash
    go get
    ```
4.  **Run the Application**: Start the service using:
    ```bash
    go run main.go
    ```
5.  The server will start and listen on the host and port specified in your `config.ini` file (e.g., `http://localhost:8080`).

## Project Structure

```
/
├── main.go         # Entry point, initializes and starts the service
├── conf/           # Configuration loading (config.ini)
│   └── conf.go
├── data/           # Caching logic for API responses
│   └── cache.go
├── open/           # Handles interaction with the OpenAI API
│   └── open.go
├── rest/           # Defines REST API endpoints and their handlers
│   ├── handlers.go
│   └── routes.go
├── tick/           # Utility for periodic tasks (e.g., cache cleanup)
│   └── tick.go
├── config.ini      # (Example, not in version control) Runtime configuration
└── README.md       # This file
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.