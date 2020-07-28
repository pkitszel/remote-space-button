#define WINVER 0x0500
#include <windows.h>

int winclick(char key) {
	int ret = 0;

	// This structure will be used to create the keyboard
	// input event.
	INPUT ip;

	// Set up a generic keyboard event.
	ip.type = INPUT_KEYBOARD;
	ip.ki.wScan = 0; // hardware scan code for key
	ip.ki.time = 0;
	ip.ki.dwExtraInfo = 0;

	// Press the key
	ip.ki.wVk = key; // virtual-key code for the  key
	ip.ki.dwFlags = 0; // 0 for key press
	ret += 1 != SendInput(1, &ip, sizeof(INPUT));

	// Release the key
	ip.ki.dwFlags = KEYEVENTF_KEYUP; // KEYEVENTF_KEYUP for key release
	ret += 1 != SendInput(1, &ip, sizeof(INPUT));

	return ret;
}
    
int main(int argc, char **argv) {
	if (argc != 2)
		return -1;

	int ret = 0;
	for (char *c = argv[1]; *c; ++c)
		ret += winclick(*c);
	return ret;
}
