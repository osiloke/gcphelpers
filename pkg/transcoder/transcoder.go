// Copyright © 2020 Osiloke Harold Emoekpere <me@osiloke.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package transcoder handles transcoding files
package transcoder

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/apex/log"
	"github.com/xfrr/goffmpeg/models"
	"github.com/xfrr/goffmpeg/transcoder"
)

// Exists check if oath exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func transcodeFile(path, destination string) (trans *transcoder.Transcoder, err error) {
	// Create new instance of transcoder
	trans = new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	log.Debugf("Transcode %s to %s", path, destination)
	err = trans.Initialize(path, destination)
	if err != nil {
		return nil, err
	}
	trans.MediaFile().SetSkipVideo(true)
	return
}

// GetAudioStream get first Audio stream
func GetAudioStream(tf *transcoder.Transcoder) *models.Streams {
	for _, s := range tf.MediaFile().Metadata().Streams {
		if s.CodecType == "audio" {
			return &s
		}
	}
	return nil
}

// DurationInParts get duration as xx:xx:xx
func DurationInParts(d float64) string {
	s := int(d)
	hours := 0
	minutes := s / 60
	seconds := s % 60
	if minutes > 60 {
		hours = minutes / 60
		minutes = minutes % 60
	}
	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

// PreviewAudioFile creates a preview of an audio file
func PreviewAudioFile(start, end int, contentPath, destinationPath string, image ...string) (*transcoder.Transcoder, error) {
	tf := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	log.Debugf("Transcode %d-%d of %s to %s", start, end, contentPath, destinationPath)
	err := tf.Initialize(contentPath, destinationPath)
	if err != nil {
		return nil, err
	}
	s := GetAudioStream(tf)
	duration, err := strconv.Atoi(strings.Split(s.Duration, ".")[0])
	if err != nil {
		return nil, err
	}
	if duration < end {
		end = duration - 1
	}
	tf.MediaFile().SetSeekTime(fmt.Sprintf("%d", start))
	tf.MediaFile().SetDuration(fmt.Sprintf("%d", end))
	tf.MediaFile().SetVideoCodec("libx264")
	tf.MediaFile().SetVideoProfile("scale=320:240")

	// if len(image) > 0 {
	// 	tf.MediaFile().Set(image[0])
	// }
	return tf, err
}

// TranscodeFile transcode a file to another format
func TranscodeFile(contentPath, destinationPath string) (*transcoder.Transcoder, error) {
	tf, err := transcodeFile(contentPath, destinationPath)
	return tf, err
}
