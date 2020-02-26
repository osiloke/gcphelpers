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
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/xfrr/goffmpeg/transcoder"
)

func TestExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Exists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transcodeFile(t *testing.T) {
	f := downloadTestFile("https://file-examples.com/wp-content/uploads/2017/11/file_example_WAV_1MG.wav", "wav")
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()
	type args struct {
		path        string
		destination string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test transcoding a single file",
			args{f.Name(), f.Name() + "_trd.mp4"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTrans, err := transcodeFile(tt.args.path, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("transcodeFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTrans == nil {
				t.Errorf("transcodeFile() = %v, is empty", gotTrans)
			}
		})
	}
}

func TestPreviewAudioFile(t *testing.T) {
	f := downloadTestFile("https://file-examples.com/wp-content/uploads/2017/11/file_example_WAV_1MG.wav", "wav")
	i := downloadTestFile("https://picsum.photos/200/300", "jpg")
	defer func() {
		f.Close()
		os.Remove(f.Name())
		i.Close()
		os.Remove(i.Name())
		os.Remove("./preview.mp4")
	}()
	type args struct {
		start           int
		end             int
		contentPath     string
		destinationPath string
		image           []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test 25 secs preview",
			args{0, 25, f.Name(), "./preview.mp4", []string{i.Name()}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PreviewAudioFile(tt.args.start, tt.args.end, tt.args.contentPath, tt.args.destinationPath, tt.args.image...)
			if (err != nil) != tt.wantErr {
				t.Errorf("PreviewAudioFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			<-got.Run(false)
			stream := getAudioStream(got)
			want := tt.args.end - tt.args.start
			gotDuration, _ := strconv.ParseFloat(stream.Duration, 32)
			if !(int(gotDuration) > want) {
				t.Errorf("PreviewAudioFile() got = %v, want %v", gotDuration, want)
			}
		})
	}
}

func TestTranscodeFile(t *testing.T) {
	type args struct {
		contentPath     string
		destinationPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *transcoder.Transcoder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TranscodeFile(tt.args.contentPath, tt.args.destinationPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("TranscodeFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TranscodeFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
