package tartools

import (
	"archive/tar"
	"compress/gzip"
	"io"
)

func WrapInTar(reader io.Reader, name string) (io.ReadCloser, error) {
	out, w := io.Pipe()

	go func() {
		defer w.Close()

		gw := gzip.NewWriter(w)
		defer gw.Close()

		tw := tar.NewWriter(gw)
		defer tw.Close()

		file, err := io.ReadAll(reader)
		if err != nil {
			return
		}

		if err := AddFile(tw, name, file, 0600); err != nil {
			return
		}
	}()

	return out, nil
}

func AddFile(tarWriter *tar.Writer, filePath string, fileContent []byte, mode int64) error {
	header := &tar.Header{
		Name: filePath,
		Mode: mode,
		Size: int64(len(fileContent)),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	_, err := tarWriter.Write(fileContent)

	return err
}