# BotPasswordGenerator

BotPasswordGenerator is a simple Telegram bot that generates random passwords based on the length specified by the user. 

The bot makes use of the telegram-bot-api library (https://go-telegram-bot-api.dev/).  


## Prerequisites
**`bot_password_key`**: This is a required environment variable that storage the telegram bot key.
  
To create the environment variable, use the following command (on Linux/macOS):
```bash
export bot_password_key="<your_secret_key>"
```

Replace <your_secret_key> with the telegram bot key.

## How to Use

Once the bot is running, you can interact with it via Telegram.  
Send the bot a message with the length of the password which must be greater than 0 and less than or equal to 100.


