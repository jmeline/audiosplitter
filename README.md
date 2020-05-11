# audiosplitter
Splitting audio using FFmpeg from a text file

Note:
  * Requires [ffmpeg](https://ffmpeg.org/download.html) to be installed
  * Currently only works with m4a files.

```
$ go build
$ ./audiosplitter audiofile.m4a tracks
```

### Tracks

audiosplitter requires a file that lays out how you would like to split the audio.

comma separated format: start,track name
```
00:00,track1
03:30,track2
07:20,track3
12:30,track4
17:08,track5
```

internally, track1 will start from 00:00 to 03:30.

Track names can have spaces in them. This program will put underscores inplace of spaces. Special characters such as & can cause problems.
