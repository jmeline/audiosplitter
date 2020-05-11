package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func createOutputDirectory(audioInput string) string {
	outputDir := strings.Replace(audioInput, ".m4a", "", -1)
	_, err := os.Stat(outputDir)
	if os.IsExist(err) {
		os.RemoveAll(outputDir)
	}
	os.Mkdir(outputDir, os.FileMode(0755))
	return outputDir
}

func buildTracksWithTimestampRanges(tracks [][]string) []Track {
	var previousTimestamp string
	var timestamps []Track
	for i, line := range tracks {
		if previousTimestamp != "" {
			var t Track
			t.startTime = previousTimestamp
			t.endTime = line[0]
			timestamps = append(timestamps, t)
		}

		previousTimestamp = line[0]

		// We're finished and need to capture the last timestamp
		if i == len(tracks)-1 {
			t := Track{
				startTime: line[0],
				endTime:   "24:00:00",
			}

			timestamps = append(timestamps, t)
		}
	}
	return timestamps
}

type Track struct {
	startTime string
	endTime   string
	trackName string
}

func collectTracksFromFile(tracksFile string) ([]Track, error) {
	file, err := os.Open(tracksFile)
	s := bufio.NewScanner(file)
	defer file.Close()

	var tracks [][]string
	for s.Scan() {
		line := strings.Split(s.Text(), ",")
		tracks = append(tracks, line)
	}

	tracksWithTimestamps := buildTracksWithTimestampRanges(tracks)
	for index := range tracksWithTimestamps {
		tracksWithTimestamps[index].trackName = tracks[index][1]
	}
	return tracksWithTimestamps, err
}

func createCommands(tracks []Track, outputDir string, audioInput string) []string {
	commandString := "ffmpeg -i %s -y -acodec copy -ss %s -to %s %s"
	var commands []string

	for i, track := range tracks {
		trackTitle := strings.Replace(track.trackName, " ", "_", -1)
		filename := fmt.Sprintf("%s/%d_%s.m4a", outputDir, i, trackTitle)
		command := fmt.Sprintf(commandString,
			audioInput,
			track.startTime,
			track.endTime,
			filename)
		commands = append(commands, command)
	}
	return commands
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough Arguments: ./audiosplitter audio.m4a tracks")
		os.Exit(1)
	}
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	audioInput := argsWithoutProg[0]
	tracksFile := argsWithoutProg[1]

	outputDir := createOutputDirectory(audioInput)
	tracks, err := collectTracksFromFile(tracksFile)
	if err != nil {
		fmt.Println("Encountered an error while collecting tracks from file", err)
		os.Exit(1)
	}

	commands := createCommands(tracks, outputDir, audioInput)
	var runningCommands []*exec.Cmd
	for _, command := range commands {
		args := strings.Fields(command)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Start()
		runningCommands = append(runningCommands, cmd)
	}

	for _, cmd := range runningCommands {
		fmt.Println(cmd)
		cmd.Wait()
	}
}
