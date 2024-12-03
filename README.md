# Anime Player CLI

**Anime Player CLI** is a simple command-line tool to stream anime directly from your terminal. It fetches stream URLs from the [amvstr.me](https://amvstr.me) API and plays anime episodes using `mpv` media player. 

This tool provides two modes:
1. **Default mode**: Plays the anime episode in your default media player window.
2. **Terminal mode**: Streams anime directly within the terminal using sixel graphics.

---

## Features

- Fetch anime streams from the web.
- Play episodes in your default media player.
- Watch anime directly in your terminal using sixel graphics.

---

## Requirements

- **`mpv` Media Player**: This tool uses `mpv` for streaming.
- **Snap (optional)**: You can install `mpv` using Snap on Ubuntu.

---

## Installation of `mpv` (Ubuntu)

1. Ensure Snap is installed and up-to-date:
   ```bash
   sudo apt update
   sudo apt install snapd
   ```
2. Install `mpv` using snap
   ```
   sudo snap install mpv
   ```

---

## How to Use

1. Fork and clone the repo
2. Build the project using 
    ```bash
    go build -o anime
    ```
3. Move it to system's PATH
    ```bash
    sudo mv anime /usr/bin/
    ```
4. Use the following syntax to play an anime episode:

   ```bash
   anime [anime-name] [episode-num]
   ```

   Example:
   ```bash
   anime naruto 5
   ```

5. To stream anime directly in your terminal, use the `--in-terminal` flag:
   ```bash
   anime [anime-name] [episode-num] --in-terminal
   ```


---

## Example Output

### Playing Normally:
```bash
Playing Naruto episode 5
```
A window with the anime episode will open.

### Playing in Terminal:
```bash
Playing Naruto episode 5
```
The anime will play directly within your terminal.

---

## Known Issues

- Make sure the anime name and episode number are correct, as the API will not fetch invalid entries.
- Ensure that `mpv` supports sixel graphics if you plan to use the `--in-terminal` mode.

---

## License

This project is licensed under the MIT License. Feel free to modify and distribute it.

---

## Contributing

If you want to contribute or report bugs, feel free to submit an issue or a pull request on the project repository.

---

Enjoy your anime streaming! 🎥✨
