package otp

import (
	"bytes"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/require"
)

type panicWriter struct{}

func (panicWriter) Write([]byte) (int, error) {
	panic("write called")
}

func TestNewWriterDoesNotUseStreams(t *testing.T) {
	require.NotNil(t, NewWriter(panicWriter{}, panicReader{}))
}

func TestWriterUsesShortPRNGReadsAcrossCalls(t *testing.T) {
	data := []byte{65, 66, 67, 68, 69}
	key := []byte{42, 1, 7, 2, 3}

	var output bytes.Buffer

	writer := NewWriter(&output, iotest.OneByteReader(bytes.NewReader(key)))

	n, err := writer.Write(data[:2])
	require.NoError(t, err)
	require.Equal(t, 2, n)
	n, err = writer.Write(data[2:])
	require.NoError(t, err)
	require.Equal(t, 3, n)
	require.Equal(t, []byte{107, 67, 68, 70, 70}, output.Bytes())
	require.Equal(t, []byte{65, 66, 67, 68, 69}, data)
}
