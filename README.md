# notifier

Get instantly notified about the important stuff on your machine, like a failed pipeline or a production bug.

Notifier is the lightweight app you don't notice, unless it is important.


# Installation
Download the .zip file and unzip it on your machine.

For macOS, double click the Notifier.dmg and drag the Notifier.app into your Applications directory.
For windows and linux, please reach out to me if you are interested in a working version.


# Configuration
### Register ngrok Account

1. Create an Account with ngrok: https://dashboard.ngrok.com
2. Save your auth token: https://dashboard.ngrok.com/get-started/your-authtoken
3. Create and save a constant endpoint: https://dashboard.ngrok.com/cloud-edge/endpoints

### Configuring App

Open the application and enter your app url and auth token


# Using Notifier

1. Create webhooks in the app
2. Copy the url
3. Trigger the url with a POST request from anywhere, i.e from a cli interface like so

```bash
    curl -X POST https://<your-url>.ngrok-free.app/notifier\?name\=<your-webhook-name>
```
