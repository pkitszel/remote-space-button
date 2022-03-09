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

## Download binaries
If your preffer to just download working solution, download following TWO .exe files:

https://github.com/pkitszel/remote-space-button/blob/master/remote-space-button.exe

https://github.com/pkitszel/remote-space-button/blob/master/winclick.exe

and place them in the same directory on the machine that plays movies, eg your Desktop.

Then just go to the `How to use` section below

## How to compile (convert source to actual program files)

### Build go web server ###
Install golang build toolchain (easy), go to your local checkout of this repo, and invoke: `go build`

Executable built in step above acts as HTTP server to be used via WIFI on smartphone.

### Build C app ###
Install your preffered C toolchain (fast and easy in case of MinGW).

`gcc src-c/winclick.c -o winclick.exe`

Both executable files make together whole application.

# How to use ###
Start `remote-space-button.exe` program.

Most probably you will need to adjust firewall rules to allow access to port 20133 (my machine just gently asked if I would like to allow it).

If desired add resulting executable to autostart, otherwise start it once at begining of each movie day.

Use your mobile device (or any device in your local network) and play/pause/etc your movies without moving your body more than necessary.
To do so, go to the website printed by server app (but use your remote device!)


http://your-ip:20133

eg

http://192.168.1.22:20133

and enjoy!
