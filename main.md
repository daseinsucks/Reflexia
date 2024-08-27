# main

## Documentation

This code sets up a Telegram bot using the `go-telegram-bot-api` library. It initializes the bot, loads environment variables, and handles incoming updates. The bot interacts with a local AI endpoint and a database to manage user data.

First, the code loads environment variables using `godotenv.Load()`. It retrieves the OpenAI API key (which is actually a local key for localai) and the Telegram bot token from the environment variables.

Next, the code initializes a new Telegram bot using the retrieved token. It also retrieves the admin ID and GPT key from the environment variables.

The code then creates a map to store admin data, including the admin ID and GPT key.

The code initializes a database and a command handler. The database is used to store user data, and the command handler is responsible for processing incoming commands.

The code then starts a loop to handle incoming updates from the Telegram API. For each update, it checks if the user is new. If the user is new, the code creates a new entry in the database.

Finally, the code starts a goroutine to handle updates and process commands. The goroutine listens for updates from the Telegram API and passes them to the command handler.

In summary, this code sets up a Telegram bot that interacts with a local AI endpoint and a database to manage user data. It handles incoming updates, processes commands, and creates new database entries for new users.



