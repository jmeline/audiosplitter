# audiosplitter
Splitting audio using FFmpeg from a text file

Note: Requires (ffmpeg)[https://ffmpeg.org/download.html] to be installed

```
$ go build
$ ./audiosplitter audiofile.m4a tracks
```

### Tracks

audiosplitter requires a file that lays out how you would like to split the audio.

comma separated format: start,end,track name
```
00:00,03:30,track1
03:30,07:20,track2
07:20,12:30,track3
12:30,17:08,track4
17:08,22:42,track5
```

Track names can have spaces in them. This program will put underscores inplace of spaces. Special characters such as & can cause problems.
