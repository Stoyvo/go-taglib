package taglib

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	name := "/tmp/test1.mp3"
	os.Remove(name)
	f := Create(name, TypeMPEG)
	if f == nil {
		t.Fatal("expected to open file, but got a NULL")
	}
	defer f.Close()
	defer os.Remove(name)
	if !f.Save() {
		t.Error("failed to save created file")
	}
}
func TestFile(t *testing.T) {
	f := Open("testdata/id3v1.mp3")
	if f == nil {
		t.Fatal("expected to open file, but got a NULL")
	}
	defer f.Close()

	if !f.Valid() {
		t.Fatal("file is not valid")
	}

	exTags := Tags{Title: "Title", Artist: "Artist", Album: "Album", Comment: "", Genre: "Blues", Year: 0, Track: 0}
	tags := f.GetTags()
	if tags.Title != exTags.Title {
		t.Errorf("expected %q; got %q", exTags.Title, tags.Title)
	}
	if tags.Artist != exTags.Artist {
		t.Errorf("expected %q; got %q", exTags.Artist, tags.Artist)
	}
	if tags.Album != exTags.Album {
		t.Errorf("expected %q; got %q", exTags.Album, tags.Album)
	}
	if tags.Genre != exTags.Genre {
		t.Errorf("expected %q; got %q", exTags.Genre, tags.Genre)
	}

	exProp := Properties{Length: 1, Bitrate: 34, Samplerate: 44100, Channels: 2}
	p := f.GetProperties()
	if p.Length != exProp.Length {
		t.Errorf("expected %q; got %q", exProp.Length, p.Length)
	}
	if p.Bitrate != exProp.Bitrate {
		t.Errorf("expected %q; got %q", exProp.Bitrate, p.Bitrate)
	}
	if p.Samplerate != exProp.Samplerate {
		t.Errorf("expected %q; got %q", exProp.Samplerate, p.Samplerate)
	}
	if p.Channels != exProp.Channels {
		t.Errorf("expected %q; got %q", exProp.Channels, p.Channels)
	}
}
