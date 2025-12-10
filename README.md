# Capitolina

Capitolina (or "Capitu") is a Telegram bot written in Go that provides utilities for finding book download links and generating summaries. It integrates with the Google Books API for metadata and file links, and Google Gemini AI for content generation.

## Environment Variables

The application requires the following environment variables to be set in a `.env` file in the root directory:

- `BOT_TOKEN`: The Telegram Bot API token.
- `BOOK_API_KEY`: API Key for Google Books.
- `GEN_AI_KEY`: API Key for Google Gemini (GenAI).

## Running the Application

### Option 1: Using Docker (Recommended)

You can run the application in a containerized environment using Docker Compose.

1.  Ensure Docker and Docker Compose are installed.
2.  Create the `.env` file with the required keys.
3.  Build and start the container in detached mode:

```bash
docker compose up --build -d
```

## Commands Reference

### /download

**Syntax:** `/download <book_title> [author_name]`

Retrieves direct download links for the specified book.

- **Parameters:**
  - `book_title` (required): The name of the book to search for.
  - `author_name` (optional): The name of the author to refine the search.
- **Behavior:**
  1.  Queries Google Books API for available PDF or EPUB tokens.
  2.  If no links are found, attempts to find the book on OpenLibrary (fallback).
  3.  Returns a list of direct links or an error message if unavailable.

### /synopsis

**Syntax:** `/synopsis <book_title> [author_name]`

Retrieves the official synopsis or description of the book.

- **Parameters:**
  - `book_title` (required): The name of the book.
  - `author_name` (optional): The author's name.
- **Behavior:**
  1.  Queries Google Books API for the volume description.
  2.  If a description exists, it is returned immediately.
  3.  If the description is missing or empty, the system automatically falls back to `/aisynopsis` to generate a summary using AI.

### /aisynopsis

**Syntax:** `/aisynopsis <book_title> [author_name]`

Forces the generation of a book summary using Artificial Intelligence, bypassing existing metadata.

- **Parameters:**
  - `book_title` (required): The name of the book.
- **Behavior:**
  - Sends a prompt to the Google Gemini 2.5 Flash model to generate a concise, 50-word summary suitable for a book cover.

### TODO

1. Test G-zip compression on googleBooks requests (as suggested in the performance documentation )
2. Include openLibrary search
3. Include download links from Project Gutenberg
