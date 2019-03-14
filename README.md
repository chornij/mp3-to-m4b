# mp3-to-m4b
Tool for converting audiobook from separated mp3 files to single mp4 (iTunes audiobook format) file
Required installed [ffmpeg](https://ffmpeg.org).

## Example
```bash
ls ~/Downloads/audiobook
```
> 01-01-01.mp3 01-01-02.mp3 01-02-01.mp3 01-03-01.mp3 02-01-01.mp3 02-02-01.mp3

```bash
./mp3-to-m4b ~/Downloads/audiobook
```
```bash
ls -alh | grep m4b
```
> audiobook.m4b