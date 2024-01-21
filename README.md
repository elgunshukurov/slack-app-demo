# Slack Bot with Interactive Confirmation

A simple Slack bot written in Go that utilizes the `github.com/slack-go/slack` and `github.com/slack-go/slack/socketmode` packages. The bot sends a confirmation message to a specified channel, containing "Yes" and "No" buttons. It handles interactive events, particularly button clicks, and acknowledges the user's response.

## Prerequisites
Before running the bot, make sure you have the following:

- Slack App Token (`SLACK_APP_TOKEN`): Set this environment variable with your Slack app token, prefixed with "xapp-".
- Slack Bot Token (`SLACK_BOT_TOKEN`): Set this environment variable with your Slack bot token, prefixed with "xoxb-".

## Installation
1. **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/your-repository.git
    cd your-repository
    ```

2. **Build and run the bot:**

    ```bash
    go build
    ./your-repository
    ```

## Usage
The bot sends a confirmation message to the Slack channel specified in the code (`#demo`). Users can click the "Yes" or "No" buttons to interact with the bot.

## Interactive Events Handling
The bot handles interactive events, such as button clicks. The `middlewareInteractive` function processes these events and acknowledges the user's interaction.

## Customization
Feel free to customize the bot to suit your needs. Update the channel, modify the confirmation message, or extend the event handling logic in the `middlewareInteractive` function.

```go
// Example: Change the channel where the confirmation message is sent
api.SendMessage("#your-custom-channel", message)
```

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement:

1. Open an [issue](https://github.com/your-username/your-repository/issues) to report problems or propose new features.
2. Create a [pull request](https://github.com/your-username/your-repository/pulls) to contribute code changes.

Please ensure that your contributions align with the project's goals and follow the [code of conduct](CODE_OF_CONDUCT.md).


## Note
Ensure that your Slack app has the necessary permissions, and make sure the bot is added to the specified channel in your Slack workspace.

Happy coding! 🤖✨
