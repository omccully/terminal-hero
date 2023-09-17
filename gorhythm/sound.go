package main

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/vorbis"
)

type songSounds struct {
	guitar     beep.StreamSeeker
	song       beep.StreamSeeker
	bass       beep.StreamSeeker
	songFormat beep.Format
}

func openAudioFile(filePath string) (beep.StreamSeeker, beep.Format, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, beep.Format{}, err
	}
	streamer, format, err := vorbis.Decode(file)
	if err != nil {
		return nil, beep.Format{}, err
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()
	bufferedStreamer := buffer.Streamer(0, buffer.Len())
	return bufferedStreamer, format, nil
}

func closeSoundStreams(songSounds songSounds) {
	guitarCLoser := songSounds.guitar.(beep.StreamSeekCloser)
	if guitarCLoser != nil {
		guitarCLoser.Close()
	}

	songCLoser := songSounds.song.(beep.StreamSeekCloser)
	if songCLoser != nil {
		songCLoser.Close()
	}

	bassCloser := songSounds.bass.(beep.StreamSeekCloser)
	if bassCloser != nil {
		bassCloser.Close()
	}
}