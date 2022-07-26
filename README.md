# twitch-nightscout

*twitch-nightscout* is a Twitch chatbot which allows users in your chat to get a snapshot of your continuous glucose monitor data from a [Nightscout](https://github.com/nightscout/cgm-remote-monitor) instance via a command.

<div style="text-align: center;"><img src="https://user-images.githubusercontent.com/32445075/181109463-b493f53b-d318-4b27-8e3f-ac90a3701602.png" alt="Preview of the `!ns` command provided by this bot"/></div>

## Installation

Pre-built executables are available in the [releases](https://github.com/p5nbTgip0r/twitch-nightscout/releases/latest) on GitHub. Extract the file anywhere you want, but keep in mind that a configuration file will be generated in the directory of the program. 

Alternatively, if you have Go installed, you can build the program from source:
```bash
git clone https://github.com/p5nbTgip0r/twitch-nightscout.git
cd twitch-nightscout
go build
```

Or, run it directly:
```bash
go run main.go
```

## Usage

Run the executable and an example configuration file will be created. Plenty of documentation is provided in the example file, but if you just want to get it running, you can follow the [steps outlined in the wiki](https://github.com/p5nbTgip0r/twitch-nightscout/wiki/Setup).

## Contributing

Pull requests and issues are welcome.