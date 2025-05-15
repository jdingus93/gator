# Gator - Database Aggregator Tool

Greetings and thank you for checking out my aggregator that I made with the help of Boot.dev!

## Prerequisites

To run this aggregator tool, you'll need:
- Go
- PostgreSQL

## Installation

### Installing Go

If you're new to coding or not very tech-savvy, the simplest way to install Go is:
1. Visit [go.dev/doc/install](https://go.dev/doc/install)
2. Click on the download for your operating system (Windows, Mac, Linux)
3. Follow the installation instructions

For Linux users familiar with the terminal, you can run:
```bash
sudo apt-get update && sudo apt-get -y install golang

Verify your version of Go using:
go version

### Installing the aggregator

You can install the aggregator by using:
go install github.com/jdingus93/gator@latest

## Configure 

To make sure that everything is setup properly you'll need to create a file named .gatorconfig.json and insert the following information:

{
"db_url": "postgres://username:password@localhost:5432/database_name",
"current_user_name": "your_username"
}

1. Make sure to replace USERNAME and PASSWORD with your PostgresSQL credentials
2. Change DATABASE_NAME to your database name\
3. Set YOUR_USERNAME to your preferred username for the aggregator

## Usage

After installing and getting setup you can run various commands like:

## Usage

After installation and configuration, you can run the following commands:

```bash
# Create a new user account
gator register

# Log in to your account
gator login

# Reset the database (caution: this will erase data)
gator reset

# List all registered users
gator users

# Run aggregation of feeds
gator agg

# Add a new RSS feed to the system
gator addfeed [feed_url]

# List all available feeds
gator feeds

# Follow a feed (must be logged in)
gator follow [feed_id]

# List all feeds you're following
gator following

# Unfollow a feed
gator unfollow [feed_follow_id]

# Browse posts from feeds you follow
gator browse