# Watched Seen Read

Inspired by Steven Soderbergh's yearly Seen, Read blog post, I set out to create a small web app I can use to track my own list for the year.

## Technologies Used

Go\
Vue\
TypeScript\
PostgreSQL\

## How It Works

Using the app is pretty straight forward, you use the form to enter what you've watched, seen, or read, the date it was completed, and what type of entry it is,
then hit Submit. The data gets saved to the PostgreSQL database, and the Entry List fetches table and displays the updated data. There are filter buttons to filter
the list down to a specific entry type.

## Installation and Setup

### Prerequisites

Make sure you have the following installed:
Go\
Node.js\
npm\
Git\

You can verify your installations with:

```
go version
node --version
npm --version
git --version
```

### Clone the Repository

```
git clone https://github.com/hardiing/watched-seen-read.git
cd watched-seen-read
```

### Backend Setup

Install the Go dependencies:

```
go mod download
```

Create a local environment file:

```
cp .env.example .env
```

Update .env with your local configuration.

Start the Go backend:

```
go run .
```

The API will typically be available at:

```
http://localhost:8080
```

### Frontend Setup

Install the Vue dependencies:

```
cd frontend
npm install
```

Start the Vue development server:

```
npm run dev
```

The frontend will typically be available at:

```
http://localhost:5173
```

### Database Setup

The backend expects a PostgreSQL connection string in the DB_URL environment variable.

Create a .env file in the project root and add the following:

```
DB_URL=postgres://username:password@localhost:5432/watched_seen_read
```

DB_URL is required for the application to connect to PostgreSQL.
