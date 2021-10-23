# GO Chat Server

Chat Server is a server written in GO that provides telnet chat functionality with the port :10000.
Every connected client will receive text written from others client.

Private project by Lorenzo Saraniti

## Prerequisites

To run the project you should have installed:

[Docker](https://www.docker.com)

[Telnet](https://it.wikipedia.org/wiki/Telnet) 

Please verify that your OS have Telnet preinstalled.  If you've a Mac you can [install it with Brew in Mac](https://brewinstall.org/how-to-install-telnet-on-mac-using-brew/)

[make](https://www.gnu.org/software/make/), you could [install it with Brew in Mac](https://formulae.brew.sh/formula/make)

## Run

To run the server:

```bash
make start
```

Docker will expose the port 10000.

Please check for any active firewall or port blocking in the OS or your software.

## Usage

To use the server in localhost, you can connect via telnet:

```bash
telnet localhost 10000
```

Type a message and that message will be forwarded to all connected clients.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
No license provided.