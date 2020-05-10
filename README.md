# audiosplitter
Splitting audio using FFmpeg from a text file

Note: Requires ffmpeg to be installed

```
./audiosplitter audiofile.m4a tracks
```

### Tracks file format

comma separated format: start,end,track name
```
00:00,03:30,track1
03:30,07:20,track2
07:20,12:30,track3
12:30,17:08,track4
17:08,22:42,track5
```
