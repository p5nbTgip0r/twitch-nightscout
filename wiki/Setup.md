This page gives a detailed walk-through of how to set up the bot.

## Configuration: Twitch
For starters, we need to configure the Twitch-specific details. This will be two things: the login details which the chatbot will use, and the name of the Twitch channel(s) the bot should listen for Nightscout commands on.

The bot was designed to allow running it on a separate Twitch user, which is why we need to specify the channel and login names separately. Although it is possible, running the bot on a separate user isn't in the scope of this tutorial. We're just going to host it on your main Twitch account for the sake of simplicity.

1. Open the generated `config.yaml` file in a text editor of some kind: [VS Code](https://code.visualstudio.com/), [Notepad++](https://notepad-plus-plus.org/), or even standard Notepad
2. Replace `my_twitch_username` under `twitch:` and `channels:` with your Twitch username.
   <details>
   <summary>GIF</summary>
   
   > ![GIF: Replacing `my_twitch_username` under the `twitch` and `channels` nodes in the config file](https://user-images.githubusercontent.com/32445075/181119814-7ea77070-493e-417c-863b-a78f87c2b4a4.gif)
   </details>
3. [Generate an OAuth password here](https://twitchapps.com/tmi/), then copy the password (be sure to include the `oauth` part at the beginning of it) into the configuration file under `oauth`.
   <details>
   <summary>GIF</summary>

   > ![GIF: Copying the OAuth generated at the twitchapps.com website and pasting it into the `oauth` key in the config file](https://user-images.githubusercontent.com/32445075/181119873-18a9944a-6610-47ff-a0f4-d93d4c5b980f.gif)
   </details>

At this point, it should look something like this:
<details>
<summary>Image</summary>

> ![Image: Showing the changes made so far. Notably, the `username` and `oauth` keys have changed under `twitch`, and the `my_twitch_username` that was under the `channels` list is also changed](https://user-images.githubusercontent.com/32445075/181112804-127037a8-2dae-4b66-b6cd-0339701fda87.png)
</details>

## Configuration: Nightscout
Now that we've entered our Twitch details, we can begin configuring the Nightscout details.

1. Replace `https://your-nightscout.herokuapp.com/` with your Nightscout URL in the `nightscout:` section under `channels`
2. *If your Nightscout requires a token or API secret to read the data*: Uncomment the `token:` line, and replace `yourtoken-1111111111111111` with the token.
   <details>
   <summary>GIF</summary>
   
   > ![GIF: Removing the `# ` before `token:` and replacing the example value with the access token from the Nightscout admin page](https://user-images.githubusercontent.com/32445075/181120773-a3303cdb-d9a6-4207-a576-182e2273cc16.gif)
   </details>

## Final Steps
Now, your config should look similar to this:
<details>
<summary>Image</summary>

> ![Image: Showing the entire config with all of the previous changes applied. Since the image shown at the end of the Twitch section, the `nightscout` key under the channel has changed. `token` is no longer commented, and the `url` has changed to the user's Nightscout URL](https://user-images.githubusercontent.com/32445075/181115795-852a8098-b34f-43a0-bed4-bf55396aed7a.png)
</details>

Save the file and run the program again. If a console appears and says:
<details open>
<summary>Image</summary>

> ![Image: A console window showing two log messages: "Initializing bot.." and "Connected"](https://user-images.githubusercontent.com/32445075/181116463-9de70901-b7f3-44d6-aa87-e81bc57b4530.png)
</details>

then the bot should be working! Go to your Twitch channel and type `!ns` in chat:
<details open>
<summary>Image</summary>

> ![Image: In a Twitch chat, a viewer says `!ns` and the streamer (running this bot) responds with diabetes data](https://user-images.githubusercontent.com/32445075/181109463-b493f53b-d318-4b27-8e3f-ac90a3701602.png)
</details>