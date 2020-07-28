# remote-space-button
Send basic keypresses to windows machine

# Why?
This project enables user to control movie players via smartphone

## Implemented functions
Each button sends specific sequence to currently active (focused) application on windows.
Current buttons:
- play/pause - space button, used in vlc, netflix, youtube, other
- S - skip intro on netflix
- F - fullscreen toggle on netflix and youtube
- FS - both of above functions at one click
- M - mute toggle

## How to compile (convert source to actual program files)
If you do not like to install any new software or are unskilled in that kind of things, I could provide pre-made binaries for you.

### Build go web server ###
Install golang build toolchain (easy) and `go build remote-pause.go`

Executable built in step above acts as HTTP server to be used via WIFI on smartphone.

If desired add resulting executable to autostart, otherwise start it once at begining of each movie day.

Most probably you will need to adjust firewall rules to allow access to port 20133 (my machine just gently asked if I would like to allow it).

### Build C app ###
Install your preffered C toolchain (fast and easy in case of MinGW).

`gcc winclick.c -o winclick.exe`

Both executable files make together whole application.

# How to use ###
Check your local IP, to do so type: `ipconfig` in cmd.

Use your mobile device (or any device in your local network) and play/pause/etc your movies without moving your body more than necessary.

Go to website:

http://four-number-ip-from-step-above:20133

eg

http://192.168.1.22:20133

and enjoy!
