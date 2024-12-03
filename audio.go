package otto

import (
	"bytes"
	"fmt"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func mergeWAV(inputFiles []string) ([]byte, error) {
	if len(inputFiles) == 1 {
		singleFile, err := audioFS.ReadFile(inputFiles[0])
		if err != nil {
			return nil, fmt.Errorf("failed to read single WAV file: %s", err)
		}
		return singleFile, nil
	}

	var combinedBuffer *audio.IntBuffer
	for i, inputFile := range inputFiles {
		_, err := audioFS.Open(inputFile)
		if err != nil {
			inputFile = "token/empty.wav"
		}
		wavFile, err := audioFS.ReadFile(inputFile)
		if err != nil {
			continue
		}
		decoder := wav.NewDecoder(bytes.NewReader(wavFile))
		if !decoder.IsValidFile() {
			return nil, fmt.Errorf("invalid WAV file: %s", inputFile)
		}
		buffer, err := decoder.FullPCMBuffer()
		if err != nil {
			return nil, fmt.Errorf("failed to read WAV data from %s: %v", inputFile, err)
		}
		if i == 0 {
			combinedBuffer = &audio.IntBuffer{
				Data:           buffer.Data,
				Format:         buffer.Format,
				SourceBitDepth: buffer.SourceBitDepth,
			}
		} else {
			if buffer.Format.SampleRate != combinedBuffer.Format.SampleRate ||
				buffer.Format.NumChannels != combinedBuffer.Format.NumChannels {
				return nil, fmt.Errorf("audio format of %s does not match the first file", inputFile)
			}
			combinedBuffer.Data = append(combinedBuffer.Data, buffer.Data...)
		}
	}

	outBuffer := NewFileBuffer(nil)
	sampleRate := combinedBuffer.Format.SampleRate
	numChannels := combinedBuffer.Format.NumChannels
	encoder := wav.NewEncoder(&outBuffer, sampleRate, 16, numChannels, 1)
	err := encoder.Write(combinedBuffer)
	if err != nil {
		return nil, fmt.Errorf("failed to write WAV data: %v", err)
	}
	err = encoder.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close encoder: %v", err)
	}
	return outBuffer.Bytes(), nil
}
